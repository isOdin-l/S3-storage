package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/isOdin-l/S3-storage/internal/service"
)

type StorageHandler struct {
	service service.StorageServiceInterface
}

func NewStorageHandler(service service.StorageServiceInterface) *StorageHandler {
	return &StorageHandler{service: service}
}

func (h *StorageHandler) Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("upload_file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = h.service.Upload(&file, fileHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, "Upload successfully done")
}

func (h *StorageHandler) Download(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "download")
}
