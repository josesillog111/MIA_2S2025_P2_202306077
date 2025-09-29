package ext2

import (
	dmanager "backend/disk_manager"
	file "backend/ext2_structs"
	rep "backend/report_manager"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	DIR_TYPE  byte = 0
	FILE_TYPE byte = 1
)

type EXT2 struct {
	Partition        dmanager.MountedPartition
	SBManager        *SBManager
	FileManager      FileManager
	InodeManager     InodeManager
	BlockManager     BlockManager
	DirectoryManager DirectoryManager
	LoggedUser       file.User // usuario actualmente logueado
}

func NewEXT2(partition dmanager.MountedPartition) *EXT2 {
	sbm := NewSBManager(partition)
	sbm.UpdateLastMountedDate()
	sbm.UpdateMountedCount()

	return &EXT2{
		Partition:        partition,
		SBManager:        sbm,
		FileManager:      *NewFileManager(partition, sbm),
		InodeManager:     *NewInodeManager(partition, sbm),
		BlockManager:     *NewBlockManager(partition, sbm),
		DirectoryManager: *NewDirectoryManager(partition, sbm),
		LoggedUser:       *file.NewUser(),
	}
}

/*	Funciones privadas	*/
func (e *EXT2) clearPartition() error {
	partition_start := e.Partition.Partition.Part_start
	partition_size := e.Partition.Partition.Part_size

	disk, err := os.OpenFile(e.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("ClearPartition: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	// Crear buffer de ceros del tamaño de la partición
	zeros := make([]byte, partition_size)

	if _, err := disk.WriteAt(zeros, partition_start); err != nil {
		return fmt.Errorf("ClearPartition: Error limpiando partición: %v", err)
	}

	return nil
}

// getUserInfo returns the UID and GID for a given username by reading /users.txt
func (e *EXT2) getUserInfo(username string) (file.User, error) {
	var info file.User
	parentInode, fileName, err := e.DirectoryManager.ResolvePath("/users.txt")
	if err != nil {
		return info, fmt.Errorf("GetUserInfo: Error resolviendo /users.txt: %v", err)
	}
	fileID, err := e.DirectoryManager.FindEntry(parentInode, fileName)
	if err != nil {
		return info, fmt.Errorf("GetUserInfo: /users.txt no encontrado: %v", err)
	}
	content, err := e.FileManager.ReadFile(fileID)
	if err != nil {
		return info, fmt.Errorf("GetUserInfo: Error leyendo /users.txt: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	var userGroup string
	found := false
	for _, line := range lines {
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		// Usuario → formato: id, U, name, group, password
		if len(parts) == 5 && parts[1] == "U" && parts[2] == username {
			fmt.Sscanf(parts[0], "%d", &info.I_uid)
			info.Name = parts[2]
			userGroup = parts[3]
			found = true
		}
		// Grupo → formato: id, G, name
		if len(parts) == 3 && parts[1] == "G" && parts[2] == userGroup {
			fmt.Sscanf(parts[0], "%d", &info.I_gid)
		}
	}
	if !found {
		return info, fmt.Errorf("GetUserInfo: Usuario '%s' no encontrado", username)
	}
	return info, nil
}

func (e *EXT2) hasPermission(inode file.Inode, mode string) bool {
	// Obtener usuario logueado
	if !e.isLogged() {
		return false
	}
	currentUser := e.LoggedUser
	userInfo, err := e.getUserInfo(currentUser.Name)
	if err != nil {
		return false
	}

	// Caso dueño
	if userInfo.I_uid == inode.I_uid {
		return e.checkBits(uint16(inode.I_permissions)>>6, mode) // dueño
	}
	// Caso grupo
	if userInfo.I_gid == inode.I_gid {
		return e.checkBits(uint16((inode.I_permissions>>3)&7), mode) // grupo
	}
	// Caso otros
	return e.checkBits(uint16(inode.I_permissions&7), mode)
}

func (e *EXT2) checkBits(bits uint16, mode string) bool {
	switch mode {
	case "r":
		return bits&4 != 0
	case "w":
		return bits&2 != 0
	case "x":
		return bits&1 != 0
	}
	return false
}

func (e *EXT2) isLogged() bool {
	return e.LoggedUser != (file.User{})
}

func (e *EXT2) isUserRoot() bool {
	return e.LoggedUser.Name == "root"
}

/*

	Implementación de FileSystem Interface

*/

// mkfs
func (e *EXT2) Mkfs(id string, typeFormat string) error {
	// 1. Limpiar partición

	if err := e.clearPartition(); err != nil {
		return fmt.Errorf("Mkfs: Error limpiando partición: %v", err)
	}

	// 2. Crear superbloque
	sb, _, _ := e.SBManager.NewSuperBlock(
		e.Partition.Partition.Part_start,
		e.Partition.Partition.Part_size,
		[10]byte{'E', 'X', 'T', '2', 0, 0, 0, 0, 0, 0},
	)
	e.SBManager.SB = *sb
	if err := e.SBManager.WriteSuperBlock(*sb); err != nil {
		return fmt.Errorf("Mkfs: Error escribiendo superbloque: %v", err)
	}

	// 3. Inicializar bitmaps
	if err := e.InodeManager.InitInodeBitmap(sb.S_bm_inode_start, sb.S_inodes_count, sb.S_first_ino); err != nil {
		return fmt.Errorf("Mkfs: Error inicializando bitmap de inodos: %v", err)
	}
	if err := e.BlockManager.InitBlockBitmap(sb.S_bm_block_start, sb.S_blocks_count); err != nil {
		return fmt.Errorf("Mkfs: Error inicializando bitmap de bloques: %v", err)
	}

	// 4. Crear directorio raíz
	if err := e.Mkdir("/", false); err != nil {
		return fmt.Errorf("Mkfs: Error creando directorio raíz: %v", err)
	}

	// 5. Crear archivo /users.txt
	tmpPath := "/tmp/users_init.txt"
	initContent := []byte("1, G, root\n1, U, root, root, 123\n")
	if err := os.WriteFile(tmpPath, initContent, 0644); err != nil {
		return fmt.Errorf("Mkfs: Error creando archivo temporal users_init: %v", err)
	}
	defer os.Remove(tmpPath) // limpiar después

	if err := e.Mkfile("/users.txt", false, 0, tmpPath); err != nil {
		return fmt.Errorf("Mkfs: Error creando /users.txt: %v", err)
	}

	return nil
}

// cat
func (e *EXT2) Cat(paths ...string) (string, error) {
	if !e.isLogged() {
		return "", fmt.Errorf("Cat: No hay ningún usuario logueado")
	}

	var result []string

	for _, p := range paths {
		virtualPath := path.Clean("/" + strings.Trim(p, "/"))

		// Resolver archivo
		parentID, name, err := e.DirectoryManager.ResolvePath(virtualPath)
		if err != nil {
			return "", fmt.Errorf("Cat: No se pudo resolver ruta '%s': %v", virtualPath, err)
		}
		fileID, err := e.DirectoryManager.FindEntry(parentID, name)
		if err != nil {
			return "", fmt.Errorf("Cat: Archivo no encontrado '%s'", virtualPath)
		}

		// Leer inodo
		ino, err := e.InodeManager.ReadInode(fileID)
		if err != nil {
			return "", fmt.Errorf("Cat: Error leyendo inodo de '%s': %v", virtualPath, err)
		}

		// Validar permisos de lectura
		if !e.hasPermission(ino, "r") {
			return "", fmt.Errorf("Cat: Permiso denegado para leer '%s'", virtualPath)
		}

		// Leer contenido
		content, err := e.FileManager.ReadFile(fileID)
		if err != nil {
			return "", fmt.Errorf("Cat: Error leyendo archivo '%s': %v", virtualPath, err)
		}

		result = append(result, string(content))
	}

	// Concatenar contenidos con salto de línea entre archivos
	return strings.Join(result, "\n"), nil
}

// login
func (e *EXT2) Login(username string, password string) error {
	// 1. Resolver /users.txt
	parentInode, fileName, err := e.DirectoryManager.ResolvePath("/users.txt")
	if err != nil {
		return fmt.Errorf("Login: Error resolviendo /users.txt: %v", err)
	}

	// Obtener inodo del archivo
	inodeID, err := e.DirectoryManager.FindEntry(parentInode, fileName)
	if err != nil {
		return fmt.Errorf("Login: /users.txt no encontrado: %v", err)
	}

	// Leer contenido
	content, err := e.FileManager.ReadFile(inodeID)
	if err != nil {
		return fmt.Errorf("Login: Error leyendo /users.txt: %v", err)
	}

	usersContent := string(content)

	// 2. Parsear usuarios válidos
	type UserInfo struct {
		Name     string
		Password string
		Group    string
	}
	var users []UserInfo

	for _, line := range strings.Split(usersContent, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		// Usuario → formato: id, U, name, group, password
		if len(parts) == 5 && parts[1] == "U" {
			users = append(users, UserInfo{
				Name:     parts[2],
				Group:    parts[3],
				Password: parts[4],
			})
		}
	}

	for _, user := range users {
		if user.Name == username && user.Password == password {
			userInfo, err := e.getUserInfo(username)
			if err != nil {
				return fmt.Errorf("Login: Error obteniendo información del usuario: %v", err)
			}
			e.LoggedUser = userInfo
			return nil
		}
	}

	return fmt.Errorf("Login: Autenticación fallida para usuario '%s'", username)
}

// logout
func (e *EXT2) Logout() error {
	if !e.isLogged() {
		return fmt.Errorf("Logout: No hay ningún usuario logueado")
	}
	e.LoggedUser = file.User{}
	return nil
}

// Crear grupo | Loguado como root
func (e *EXT2) Mkgrp(name string) error {
	if !e.isUserRoot() {
		return fmt.Errorf("Mkgrp: Operación no permitida, solo el usuario root puede crear grupos")
	}
	if !e.isLogged() {
		return fmt.Errorf("Mkgrp: No hay ningún usuario logueado")
	}
	inode, _, err := e.FileManager.DirectoryManager.ResolvePath("/users.txt")
	if err != nil {
		return fmt.Errorf("Mkgrp: Error resolviendo /users.txt: %v", err)
	}
	fileID, err := e.FileManager.DirectoryManager.FindEntry(inode, "users.txt")
	if err != nil {
		return fmt.Errorf("Mkgrp: /users.txt no encontrado: %v", err)
	}

	contentBytes, err := e.FileManager.ReadFile(fileID)
	if err != nil {
		return fmt.Errorf("Mkgrp: Error leyendo /users.txt: %v", err)
	}

	lines := strings.Split(string(contentBytes), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) >= 3 && parts[1] == "G" && parts[2] == name {
			return fmt.Errorf("Mkgrp: Grupo '%s' ya existe", name)
		}
	}

	// Generar id (simple: 1 + número de líneas existentes)
	newID := len(lines) + 1
	newLine := fmt.Sprintf("%d, G, %s", newID, name)
	lines = append(lines, newLine)
	newContent := strings.Join(lines, "\n")

	if err := e.FileManager.UpdateFile(fileID, []byte(newContent)); err != nil {
		return fmt.Errorf("Mkgrp: Error actualizando /users.txt: %v", err)
	}

	return nil
}

// Eliminar grupo | Loguado como root
func (e *EXT2) Rmgrp(name string) error {
	if !e.isUserRoot() {
		return fmt.Errorf("Rmgrp: Operación no permitida, solo el usuario root puede eliminar grupos")
	}
	if !e.isLogged() {
		return fmt.Errorf("Rmgrp: No hay ningún usuario logueado")
	}
	inode, _, err := e.FileManager.DirectoryManager.ResolvePath("/users.txt")
	if err != nil {
		return fmt.Errorf("Rmgrp: Error resolviendo /users.txt: %v", err)
	}
	fileID, err := e.FileManager.DirectoryManager.FindEntry(inode, "users.txt")
	if err != nil {
		return fmt.Errorf("Rmgrp: /users.txt no encontrado: %v", err)
	}

	contentBytes, err := e.FileManager.ReadFile(fileID)
	if err != nil {
		return fmt.Errorf("Rmgrp: Error leyendo /users.txt: %v", err)
	}

	lines := strings.Split(string(contentBytes), "\n")
	var newLines []string
	found := false
	for _, line := range lines {
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) >= 3 {
			if parts[1] == "G" && parts[2] == name {
				found = true
				continue // eliminar grupo
			}
			if parts[1] == "U" && parts[3] == name {
				continue // eliminar usuarios del grupo
			}
		}
		if strings.TrimSpace(line) != "" {
			newLines = append(newLines, line)
		}
	}

	if !found {
		return fmt.Errorf("Rmgrp: Grupo '%s' no encontrado", name)
	}

	newContent := strings.Join(newLines, "\n")
	if err := e.FileManager.UpdateFile(fileID, []byte(newContent)); err != nil {
		return fmt.Errorf("Rmgrp: Error actualizando /users.txt: %v", err)
	}

	return nil
}

// Crear usuario | Logueado como root
func (e *EXT2) Mkusr(name, pass, grp string) error {
	if !e.isUserRoot() {
		return fmt.Errorf("Mkusr: Operación no permitida, solo el usuario root puede crear usuarios")
	}
	if !e.isLogged() {
		return fmt.Errorf("Mkusr: No hay ningún usuario logueado")
	}
	inode, _, err := e.FileManager.DirectoryManager.ResolvePath("/users.txt")
	if err != nil {
		return fmt.Errorf("Mkusr: Error resolviendo /users.txt: %v", err)
	}
	fileID, err := e.FileManager.DirectoryManager.FindEntry(inode, "users.txt")
	if err != nil {
		return fmt.Errorf("Mkusr: /users.txt no encontrado: %v", err)
	}

	contentBytes, err := e.FileManager.ReadFile(fileID)
	if err != nil {
		return fmt.Errorf("Mkusr: Error leyendo /users.txt: %v", err)
	}

	lines := strings.Split(string(contentBytes), "\n")
	groupExists := false
	for _, line := range lines {
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) >= 3 && parts[1] == "G" && parts[2] == grp {
			groupExists = true
		}
		if len(parts) >= 5 && parts[1] == "U" && parts[2] == name {
			return fmt.Errorf("Mkusr: usuario '%s' ya existe", name)
		}
	}

	if !groupExists {
		return fmt.Errorf("Mkusr: grupo '%s' no existe", grp)
	}

	newID := len(lines) + 1
	newLine := fmt.Sprintf("%d, U, %s, %s, %s", newID, name, grp, pass) // formato consistente
	lines = append(lines, newLine)
	newContent := strings.Join(lines, "\n")

	if err := e.FileManager.UpdateFile(fileID, []byte(newContent)); err != nil {
		return fmt.Errorf("Mkusr: error actualizando /users.txt: %v", err)
	}

	return nil
}

// Eliminar usuario | Logueado como root
func (e *EXT2) Rmusr(name string) error {
	if !e.isUserRoot() {
		return fmt.Errorf("Rmusr: Operación no permitida, solo el usuario root puede eliminar usuarios")
	}
	if !e.isLogged() {
		return fmt.Errorf("Rmusr: No hay ningún usuario logueado")
	}
	inode, _, err := e.FileManager.DirectoryManager.ResolvePath("/users.txt")
	if err != nil {
		return fmt.Errorf("Rmusr: Error resolviendo /users.txt: %v", err)
	}
	fileID, err := e.FileManager.DirectoryManager.FindEntry(inode, "users.txt")
	if err != nil {
		return fmt.Errorf("Rmusr: /users.txt no encontrado: %v", err)
	}

	contentBytes, err := e.FileManager.ReadFile(fileID)
	if err != nil {
		return fmt.Errorf("Rmusr: Error leyendo /users.txt: %v", err)
	}

	lines := strings.Split(string(contentBytes), "\n")
	var newLines []string
	found := false
	for _, line := range lines {
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) >= 5 && parts[1] == "U" && parts[2] == name {
			found = true
			continue
		}
		if strings.TrimSpace(line) != "" {
			newLines = append(newLines, line)
		}
	}

	if !found {
		return fmt.Errorf("Rmusr: Usuario '%s' no encontrado", name)
	}

	newContent := strings.Join(newLines, "\n")
	if err := e.FileManager.UpdateFile(fileID, []byte(newContent)); err != nil {
		return fmt.Errorf("Rmusr: Error actualizando /users.txt: %v", err)
	}

	return nil
}

// Cambiar grupo de usuario | Logueado como root
func (e *EXT2) Chgrp(name string, newgrp string) error {
	if !e.isUserRoot() {
		return fmt.Errorf("Chgrp: Operación no permitida, solo el usuario root puede cambiar grupos")
	}
	if !e.isLogged() {
		return fmt.Errorf("Chgrp: No hay ningún usuario logueado")
	}

	// 1. Obtener inodo de /users.txt
	parentInode, fileName, err := e.FileManager.DirectoryManager.ResolvePath("/users.txt")
	if err != nil {
		return fmt.Errorf("Chgrp: Error resolviendo /users.txt: %v", err)
	}
	fileID, err := e.FileManager.DirectoryManager.FindEntry(parentInode, fileName)
	if err != nil {
		return fmt.Errorf("Chgrp: /users.txt no encontrado: %v", err)
	}

	// 2. Leer contenido
	contentBytes, err := e.FileManager.ReadFile(fileID)
	if err != nil {
		return fmt.Errorf("Chgrp: Error leyendo /users.txt: %v", err)
	}
	lines := strings.Split(string(contentBytes), "\n")

	// 3. Validar existencia de usuario y grupo
	groupExists := false
	userFound := false
	for _, line := range lines {
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) >= 3 && parts[1] == "G" && parts[2] == newgrp {
			groupExists = true
		}
	}
	if !groupExists {
		return fmt.Errorf("Chgrp: El grupo '%s' no existe", newgrp)
	}

	// 4. Reescribir cambiando grupo del usuario
	var newLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		if len(parts) == 5 && parts[1] == "U" && parts[2] == name {
			userFound = true
			// Usuario → formato: id, U, name, group, pass
			parts[3] = newgrp
			line = fmt.Sprintf("%s, %s, %s, %s, %s", parts[0], parts[1], parts[2], parts[3], parts[4])
		}
		newLines = append(newLines, line)
	}

	if !userFound {
		return fmt.Errorf("Chgrp: Usuario '%s' no encontrado", name)
	}

	// 5. Guardar de nuevo en /users.txt
	newContent := strings.Join(newLines, "\n")
	if err := e.FileManager.UpdateFile(fileID, []byte(newContent)); err != nil {
		return fmt.Errorf("Chgrp: Error actualizando /users.txt: %v", err)
	}

	return nil
}

// mkfile
func (e *EXT2) Mkfile(virtualPath string, recursive bool, size int64, content string) error {
	// Normalizar la ruta
	virtualPath = path.Clean("/" + strings.Trim(virtualPath, "/"))

	isUsersFile := virtualPath == "/users.txt"

	if !isUsersFile && !e.isLogged() {
		return fmt.Errorf("Mkfile: No hay ningún usuario logueado")
	}
	if virtualPath == "/" {
		return fmt.Errorf("Mkfile: No se puede crear archivo en '/'")
	}

	// Resolver padre
	parentID, name, err := e.DirectoryManager.ResolvePath(virtualPath)
	if err != nil {
		if recursive {
			if err := e.Mkdir(path.Dir(virtualPath), true); err != nil {
				return fmt.Errorf("Mkfile: Error creando directorios intermedios: %v", err)
			}
			parentID, name, err = e.DirectoryManager.ResolvePath(virtualPath)
			if err != nil {
				return fmt.Errorf("Mkfile: Error resolviendo ruta '%s': %v", virtualPath, err)
			}
		} else {
			return fmt.Errorf("Mkfile: Directorio padre no existe, use -r para crearlo")
		}
	}

	// Verificar permisos de escritura en padre
	if !isUsersFile {
		parentInode, err := e.InodeManager.ReadInode(parentID)
		if err != nil {
			return fmt.Errorf("Mkfile: Error leyendo inodo del padre: %v", err)
		}
		if !e.hasPermission(parentInode, "w") {
			return fmt.Errorf("Mkfile: Permiso denegado en directorio padre")
		}
	}

	// Comprobar existencia previa
	entries, err := e.DirectoryManager.ReadDirectory(parentID)
	if err != nil {
		return fmt.Errorf("Mkfile: Error leyendo directorio padre: %v", err)
	}
	for _, entry := range entries {
		existingName := strings.Trim(string(entry.B_name[:]), "\x00")
		if existingName == name {
			return fmt.Errorf("Mkfile: El archivo '%s' ya existe", name)
		}
	}

	// --- Preparar contenido ---
	var contentBytes []byte
	if content != "" {
		// Cargar desde archivo externo
		data, err := os.ReadFile(content)
		if err != nil {
			return fmt.Errorf("Mkfile: Error leyendo archivo externo '%s': %v", content, err)
		}
		contentBytes = data
	} else {
		if size < 0 {
			return fmt.Errorf("Mkfile: Tamaño inválido (%d)", size)
		}
		if size > 0 {
			digits := []byte("0123456789")
			for int64(len(contentBytes)) < size {
				remaining := size - int64(len(contentBytes))
				if remaining >= int64(len(digits)) {
					contentBytes = append(contentBytes, digits...)
				} else {
					contentBytes = append(contentBytes, digits[:remaining]...)
				}
			}
		}
	}

	// Determinar propietario y permisos
	ownerName := "root"
	defaultPerms := uint16(0o600)
	if !isUsersFile {
		ownerName = e.LoggedUser.Name
		defaultPerms = 0o664
	}

	// Crear archivo
	id, err := e.FileManager.CreateFile(parentID, name, contentBytes, ownerName, defaultPerms)
	if err != nil {
		return fmt.Errorf("Mkfile: Error creando archivo: %v", err)
	}

	// Ajustar UID/GID/permisos en inodo
	ino, err := e.InodeManager.ReadInode(id)
	if err != nil {
		return fmt.Errorf("Mkfile: Error leyendo inodo: %v", err)
	}
	if isUsersFile {
		ino.I_uid = 1
		ino.I_gid = 1
		ino.I_permissions = 0o600
	} else {
		ino.I_uid = e.LoggedUser.I_uid
		ino.I_gid = e.LoggedUser.I_gid
		ino.I_permissions = defaultPerms
	}
	if err := e.InodeManager.UpdateInode(id, ino); err != nil {
		return fmt.Errorf("Mkfile: Error actualizando inodo: %v", err)
	}

	return nil
}

// mkdir
func (e *EXT2) Mkdir(virtualPath string, p bool) error {
	virtualPath = path.Clean("/" + strings.Trim(virtualPath, "/"))

	// Caso raíz
	if virtualPath == "/" {

		rootInodeID, rootInode, err := e.InodeManager.CreateInode(
			DIR_TYPE,
			int64(e.SBManager.SB.S_block_size),
			e.LoggedUser.Name,
			0o755, // permisos rwxr-xr-x para el root
		)
		if err != nil {
			return fmt.Errorf("Mkdir (/): Error creando inodo raíz: %v", err)
		}

		// Inicializar bloque con "." y ".."
		var dir file.Directory
		for i := range dir.B_content {
			dir.B_content[i].B_inodo = -1
		}
		copy(dir.B_content[0].B_name[:], []byte("."))
		dir.B_content[0].B_inodo = rootInodeID
		copy(dir.B_content[1].B_name[:], []byte(".."))
		dir.B_content[1].B_inodo = rootInodeID

		db, err := e.DirectoryManager.structToBytes(&dir)
		if err != nil {
			_ = e.InodeManager.DeleteInode(rootInodeID)
			return fmt.Errorf("Mkdir (/): Error serializando directorio raíz: %v", err)
		}
		blockID, err := e.BlockManager.CreateBlock(db)
		if err != nil {
			_ = e.InodeManager.DeleteInode(rootInodeID)
			return fmt.Errorf("Mkdir (/): No se pudo crear bloque raíz: %v", err)
		}
		rootInode.I_block[0] = blockID
		rootInode.I_size = e.SBManager.SB.S_block_size
		if err := e.InodeManager.UpdateInode(rootInodeID, rootInode); err != nil {
			_ = e.InodeManager.DeleteInode(rootInodeID)
			return fmt.Errorf("Mkdir (/): Error actualizando inodo raíz: %v", err)
		}
		return nil
	}

	if !e.isLogged() {
		return fmt.Errorf("Mkdir: No hay ningún usuario logueado")
	}
	// Caso general
	parentID, name, err := e.DirectoryManager.ResolvePath(virtualPath)
	if err != nil {
		if !p {
			return fmt.Errorf("Mkdir: No se pudo resolver ruta: %v", err)
		}
		// Crear directorios intermedios si p == true
		if err := e.Mkdir(path.Dir(virtualPath), true); err != nil {
			return fmt.Errorf("Mkdir: Error creando directorios intermedios: %v", err)
		}
		parentID, name, err = e.DirectoryManager.ResolvePath(virtualPath)
		if err != nil {
			return fmt.Errorf("Mkdir: Error resolviendo ruta '%s': %v", virtualPath, err)
		}
	}

	// Validar permisos de escritura en el padre
	parentInode, err := e.InodeManager.ReadInode(parentID)
	if err != nil {
		return fmt.Errorf("Mkdir: Error leyendo inodo del padre: %v", err)
	}
	if !e.hasPermission(parentInode, "w") {
		return fmt.Errorf("Mkdir: Permiso denegado. El usuario '%s' no puede escribir en el directorio padre", e.LoggedUser.Name)
	}

	// Comprobar si ya existe
	entries, err := e.DirectoryManager.ReadDirectory(parentID)
	if err != nil {
		return fmt.Errorf("Mkdir: Error leyendo directorio padre: %v", err)
	}
	for _, entry := range entries {
		existingName := strings.Trim(string(entry.B_name[:]), "\x00")
		if existingName == name {
			return fmt.Errorf("Mkdir: El directorio '%s' ya existe en '%s'", name, path.Dir(virtualPath))
		}
	}

	// Crear inodo para el nuevo directorio con permisos 775
	newInodeID, newInode, err := e.InodeManager.CreateInode(
		DIR_TYPE,
		int64(e.SBManager.SB.S_block_size),
		e.LoggedUser.Name,
		0o775, // rwxrwxr-x
	)
	if err != nil {
		return fmt.Errorf("Mkdir: Error creando inodo: %v", err)
	}

	// Bloque con "." y ".."
	var dirStruct file.Directory
	for i := range dirStruct.B_content {
		dirStruct.B_content[i].B_inodo = -1
	}
	copy(dirStruct.B_content[0].B_name[:], []byte("."))
	dirStruct.B_content[0].B_inodo = newInodeID
	copy(dirStruct.B_content[1].B_name[:], []byte(".."))
	dirStruct.B_content[1].B_inodo = parentID

	db, err := e.DirectoryManager.structToBytes(&dirStruct)
	if err != nil {
		_ = e.InodeManager.DeleteInode(newInodeID)
		return fmt.Errorf("Mkdir: Error serializando nuevo directorio: %v", err)
	}
	blockID, err := e.BlockManager.CreateBlock(db)
	if err != nil {
		_ = e.InodeManager.DeleteInode(newInodeID)
		return fmt.Errorf("Mkdir: No se pudo crear bloque del nuevo directorio: %v", err)
	}

	newInode.I_block[0] = blockID
	newInode.I_size = e.SBManager.SB.S_block_size
	if err := e.InodeManager.UpdateInode(newInodeID, newInode); err != nil {
		_ = e.InodeManager.DeleteInode(newInodeID)
		return fmt.Errorf("Mkdir: Error actualizando inodo nuevo: %v", err)
	}

	if err := e.DirectoryManager.AddEntry(parentID, name, newInodeID); err != nil {
		_ = e.BlockManager.DeleteBlock(blockID)
		return fmt.Errorf("Mkdir: Error añadiendo entrada al padre: %v", err)
	}

	return nil
}

func (e *EXT2) Remove(path string) error {
	return nil
}

func (e *EXT2) Edit(path string, contenido string) error {
	return nil
}

func (e *EXT2) Rename(path string, name string) error {
	return nil
}

func (e *EXT2) Copy(path string, dest string) error {
	return nil
}

func (e *EXT2) Move(path string, dest string) error {
	return nil
}

func (e *EXT2) Find(path string, name string) (string, error) {
	return "no implementado", nil
}

func (e *EXT2) Chown(path string, user string, recursive bool) error {
	return nil
}

func (e *EXT2) Chmod(path string, ugo int64, recursive bool) error {
	return nil
}

func (e *EXT2) Recovery() error {
	fmt.Println("EXT2 no cuenta con soporte para recuperación de datos")
	return nil
}

func (e *EXT2) Loss() error {
	fmt.Println("EXT2 no cuenta con soporte para pérdida de datos")
	return nil
}

func (e *EXT2) Journaling() error {
	fmt.Println("EXT2 no cuenta con soporte para journaling")
	return nil
}

// rep
func (e *EXT2) Rep(name string, path string, path_file_ls string, format string) error {
	// Sacar extensión del archivo de salida
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))
	if ext == "" {
		return fmt.Errorf("Rep: no se pudo determinar extensión para '%s'", path)
	}

	// Instanciar generador con los managers reales de este FS
	var generator rep.ReportSysGenerator = NewReportSysManager(
		e.Partition,
		e.SBManager,
		e.FileManager,
		e.InodeManager,
		e.BlockManager,
		e.DirectoryManager,
	)

	if name == "ls" {
		generator.Ls(path, path_file_ls, format)
		return nil
	}

	if name == "file" {
		generator.File(path, path_file_ls, format)
		return nil
	}

	// Mapa de reportes soportados
	reporters := map[string]func(string, string) error{
		"inode":    generator.Inode,
		"block":    generator.Block,
		"bm_inode": generator.BitmapInode,
		"bm_block": generator.BitmapBlock,
		"tree":     generator.Tree,
		"sb":       generator.SuperBlock,
	}

	if fn, ok := reporters[name]; ok {
		if err := fn(path, ext); err != nil {
			return fmt.Errorf("Rep: error ejecutando reporte '%s': %v", name, err)
		}
		return nil
	}

	return fmt.Errorf("Rep: reporte '%s' no soportado", name)
}
