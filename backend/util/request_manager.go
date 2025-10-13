package util

import (
	dmanager "backend/disk_manager"
	ext2 "backend/ext2_manager"
	ext3 "backend/ext3_manager"
	fsys "backend/file_system"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type RequestManager struct {
	DiskManager           *dmanager.DiskManager
	ListMountedPartitions **dmanager.MountedPartitionList
	IdMountedAndLogued    *string /* ID de la partición montada y del usuario logueado */
}

func NewRequestManager(diskManager *dmanager.DiskManager, listMountedPartitions **dmanager.MountedPartitionList, idMountedAndLogued *string) *RequestManager {
	return &RequestManager{
		DiskManager:           diskManager,
		ListMountedPartitions: listMountedPartitions,
		IdMountedAndLogued:    idMountedAndLogued,
	}
}

// helper para respuestas JSON
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// helper para errores
func jsonError(w http.ResponseWriter, status int, err error) {
	jsonResponse(w, status, map[string]string{"error": err.Error()})
}

// Crea una instancia del sistema de archivos basado en el tipo y la partición montada
func (v *RequestManager) NewFileSystem(fsType string, partitionId string) (fsys.FileSystem, error) {
	switch strings.ToLower(fsType) {
	case "2fs":
		partition := (*v.ListMountedPartitions).GetPartitionById(partitionId)
		if partition == nil {
			return nil, fmt.Errorf("NewFileSystem: la partición con id '%s' no está montada", partitionId)
		}
		return ext2.NewEXT2(*partition), nil
	case "3fs":
		partition := (*v.ListMountedPartitions).GetPartitionById(partitionId)
		if partition == nil {
			return nil, fmt.Errorf("NewFileSystem: la partición con id '%s' no está montada", partitionId)
		}
		return ext3.NewEXT3(*partition), nil
	default:
		return nil, fmt.Errorf("NewFileSystem: Sistema de archivos '%s' no soportado", fsType)
	}
}

// GET /disk/list
func (rm *RequestManager) ListDisksHandler(w http.ResponseWriter, r *http.Request) {
	// Define the path to the disks.json file relative to where the program is executed.
	regPath := "disks.json"

	// 1. Read the contents of the file.
	data, err := os.ReadFile(regPath)

	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Disk registry file not found.", http.StatusNotFound)
			return
		}

		fmt.Printf("Error reading disk registry file: %v\n", err)
		http.Error(w, "Error reading disk registry.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}

// GET /disk/partitions?diskPath=/ruta/del/disco
func (rm *RequestManager) ListPartitionsHandler(w http.ResponseWriter, r *http.Request) {
	diskPath := r.URL.Query().Get("diskPath")
	if diskPath == "" {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("ListPartitionsHandler: Debe proporcionar 'diskPath'"))
		return
	}

	partitions, ebrs, err := rm.DiskManager.PartitionManager.ListPartitions(diskPath)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	// Convertir las particiones MBR a una forma legible
	type PartitionInfo struct {
		Status      byte   `json:"status"`
		Type        string `json:"type"`
		Fit         string `json:"fit"`
		Start       int64  `json:"start"`
		Size        int64  `json:"size"`
		Name        string `json:"name"`
		Correlative int64  `json:"correlative"`
		ID          string `json:"id"`
		Usage       int64  `json:"usage"`
	}

	type EBRInfo struct {
		Status byte   `json:"status"`
		Fit    string `json:"fit"`
		Start  int64  `json:"start"`
		Size   int64  `json:"size"`
		Next   int64  `json:"next"`
		Name   string `json:"name"`
		Usage  int64  `json:"usage"`
	}

	mbr, err := rm.DiskManager.MbrManager.ReadMBR(diskPath)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, fmt.Errorf("ListPartitionsHandler: error al leer el MBR para calcular uso: %v", err))
		return
	}

	// Serializar particiones principales (MBR)
	var partitionsInfo []PartitionInfo
	for _, p := range partitions {
		if p.Part_size == 0 {
			continue // ignorar particiones no usadas
		}

		usage := (p.Part_size * 100) / mbr.Mbr_tamano

		partitionsInfo = append(partitionsInfo, PartitionInfo{
			Status:      p.Part_status,
			Type:        string(p.Part_type),
			Fit:         string(p.Part_fit),
			Start:       p.Part_start,
			Size:        p.Part_size,
			Name:        string(bytes.Trim(p.Part_name[:], "\x00")),
			Correlative: p.Part_correlative,
			ID:          string(bytes.Trim(p.Part_id[:], "\x00")),
			Usage:       usage,
		})
	}

	// Serializar EBRs (particiones lógicas)
	var ebrInfo []EBRInfo
	for _, e := range ebrs {
		usage := (e.Part_size * 100) / mbr.Mbr_tamano

		ebrInfo = append(ebrInfo, EBRInfo{
			Status: e.Part_mount,
			Fit:    string(e.Part_fit),
			Start:  e.Part_start,
			Size:   e.Part_size,
			Next:   e.Part_next,
			Name:   string(bytes.Trim(e.Part_name[:], "\x00")),
			Usage:  usage,
		})
	}

	// Enviar respuesta JSON combinada
	response := map[string]interface{}{
		"mbr_partitions": partitionsInfo,
		"ebr_partitions": ebrInfo,
	}

	jsonResponse(w, http.StatusOK, response)
}

// GET /disk/checkfs?diskPath=/ruta/disco&partitionName=nombre_particion
func (rm *RequestManager) CheckFSHandler(w http.ResponseWriter, r *http.Request) {
	diskPath := r.URL.Query().Get("diskPath")
	partitionName := r.URL.Query().Get("partitionName")

	if diskPath == "" {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("CheckFSHandler: Debe proporcionar 'diskPath'"))
		return
	}

	list, err := rm.DiskManager.Mount(**rm.ListMountedPartitions, diskPath, partitionName)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("CheckFSHandler: Error al montar la partición: %v", err))
		return
	}

	*rm.ListMountedPartitions = &list

	// Buscar partición recién montada
	mounted := (*rm.ListMountedPartitions).GetPartitionByPathAndName(diskPath, partitionName)
	if mounted == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("CheckFSHandler: La partición montada no fue encontrada después de montar"))
		return
	}

	// Revisar si ya tiene FS
	hasFS, err := mounted.HasFileSystem()
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("CheckFSHandler: Error al verificar el sistema de archivos: %v", err))
		return
	}

	list, err = rm.DiskManager.Unmount(**rm.ListMountedPartitions, mounted.Id)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("CheckFSHandler: Error al desmontar la partición con ID %s: %v", mounted.Id, err))
		return
	}

	*rm.ListMountedPartitions = &list

	if hasFS {
		jsonResponse(w, http.StatusOK, map[string]interface{}{
			"hasFS": hasFS,
		})
	} else {
		jsonResponse(w, http.StatusOK, map[string]interface{}{
			"hasFS": hasFS,
		})
	}
}

// GET /disk/mounted?diskPath=/ruta/del/disco&partitionName=nombre_particion
func (rm *RequestManager) MountDiskHandler(w http.ResponseWriter, r *http.Request) {
	var diskPath = r.URL.Query().Get("diskPath")
	var partitionName = r.URL.Query().Get("partitionName")

	if diskPath == "" || partitionName == "" {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("MountDiskHandler: Debe proporcionar 'diskPath' y 'partitionName'"))
		return
	}

	// Montar partición
	partition, err := rm.DiskManager.Mount(**rm.ListMountedPartitions, diskPath, partitionName)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	// Actualizar lista
	*rm.ListMountedPartitions = &partition

	// Buscar partición recién montada
	mounted := (*rm.ListMountedPartitions).GetPartitionByPathAndName(diskPath, partitionName)
	if mounted == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("MountDiskHandler: La partición montada no fue encontrada después de montar"))
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{
		"message":     "Partición montada exitosamente",
		"partitionId": mounted.Id,
	})
}

// POST /api/fs/login
func (rm *RequestManager) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PartitionId string `json:"partitionId"`
		FsType      string `json:"fsType"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("LoginHandler: JSON inválido"))
		return
	}

	fs, err := rm.NewFileSystem(req.FsType, req.PartitionId)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	part := (*rm.ListMountedPartitions).GetPartitionById(req.PartitionId)
	if part == nil { // debería ser imposible
		jsonError(w, http.StatusInternalServerError, fmt.Errorf("LoginHandler: la partición con id '%s' no está montada", req.PartitionId))
		return
	}
	part.FileSystem = fs

	*rm.IdMountedAndLogued = req.PartitionId
	jsonResponse(w, http.StatusOK, map[string]string{"message": "Login exitoso"})
}

// GET /api/fs/list?path=/ruta
func (rm *RequestManager) ListHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")

	if path == "" {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("ListHandler: Debe proporcionar 'path'"))
		return
	}

	partPtr := (*rm.ListMountedPartitions).GetPartitionById(*rm.IdMountedAndLogued)
	if partPtr == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("ListHandler: La partición con id '%s' no está montada", *rm.IdMountedAndLogued))
		return
	}
	if partPtr.FileSystem == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("ListHandler: La partición con id '%s' no tiene un sistema de archivos asociado", *rm.IdMountedAndLogued))
		return
	}

	files, err := partPtr.FileSystem.List(path)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	jsonResponse(w, http.StatusOK, files)
}

// GET /api/fs/read?path=/ruta/archivo.txt
func (rm *RequestManager) ReadHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("ReadHandler: Debe proporcionar 'path'"))
		return
	}

	part := (*rm.ListMountedPartitions).GetPartitionById(*rm.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("ReadHandler: La partición con id '%s' no está montada", *rm.IdMountedAndLogued))
		return
	}

	content, err := part.FileSystem.Cat(path)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"content": content})
}

// POST /api/fs/mkdir
func (rm *RequestManager) MkdirHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Path string `json:"path"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("MkdirHandler: JSON inválido"))
		return
	}

	part := (*rm.ListMountedPartitions).GetPartitionById(*rm.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("MkdirHandler: La partición con id '%s' no está montada", *rm.IdMountedAndLogued))
		return
	}

	err := part.FileSystem.Mkdir(req.Path, true)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": "Directorio creado"})
}

// POST /api/fs/create
func (rm *RequestManager) MkfileHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("MkfileHandler: JSON inválido"))
		return
	}

	part := (*rm.ListMountedPartitions).GetPartitionById(*rm.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("MkfileHandler: La partición con id '%s' no está montada", *rm.IdMountedAndLogued))
		return
	}

	err := part.FileSystem.Mkfile(req.Path, true, 0, req.Content)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": "Archivo creado"})
}

// DELETE /api/fs/delete?path=/ruta
func (rm *RequestManager) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("DeleteHandler: Debe proporcionar 'path'"))
		return
	}
	part := (*rm.ListMountedPartitions).GetPartitionById(*rm.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("DeleteHandler: La partición con id '%s' no está montada", *rm.IdMountedAndLogued))
		return
	}

	err := part.FileSystem.Remove(path) // suponiendo que implementas Remove() en EXT3
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"message": "Eliminado correctamente"})
}
