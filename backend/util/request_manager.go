package util

import (
	dmanager "backend/disk_manager"
	ext2 "backend/ext2_manager"
	ext3 "backend/ext3_manager"
	fsys "backend/file_system"
	"encoding/json"
	"fmt"
	"net/http"
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

// GET /api/fs/list?path=/ruta
func (rm *RequestManager) ListHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")

	if path == "" {
		jsonError(w, http.StatusBadRequest, fmt.Errorf("ListHandler: Debe proporcionar 'path'"))
		return
	}

	fmt.Println("ID MOUNTED AND LOGUED:", *rm.IdMountedAndLogued)

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
