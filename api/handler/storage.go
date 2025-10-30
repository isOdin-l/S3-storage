package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/isOdin-l/S3-storage/internal/service"
)

type StorageHandler struct {
	s service.StorageServiceInterface
}

func NewStorageHandler(service service.StorageServiceInterface) *StorageHandler {
	return &StorageHandler{s: service}
}

func (h *StorageHandler) Upload(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "/upload")
}

func (h *StorageHandler) Download(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "download")
}
