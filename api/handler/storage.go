package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/isOdin-l/S3-storage/internal/service"
	"github.com/isOdin-l/S3-storage/internal/types"
)

type StorageHandler struct {
	service service.StorageServiceInterface
}

func NewStorageHandler(service service.StorageServiceInterface) *StorageHandler {
	return &StorageHandler{service: service}
}

func (h *StorageHandler) Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, errReadFile := r.FormFile("upload_file")
	if errReadFile != nil {
		http.Error(w, errReadFile.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	student := &types.StudentInfo{
		GroupName:  r.FormValue("groupName"),
		Name:       r.FormValue("name"),
		Surname:    r.FormValue("surname"),
		Patronymic: r.FormValue("patronymic"),
	}

	errServer := h.service.Upload(&file, fileHeader, student)
	if errServer != nil {
		http.Error(w, errServer.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, "Upload successfully done")
}
