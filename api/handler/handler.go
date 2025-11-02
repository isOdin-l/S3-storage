package handler

import (
	"net/http"

	"github.com/isOdin-l/S3-storage/internal/service"
)

type StorageHandlerInterface interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	StorageHandlerInterface
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		StorageHandlerInterface: NewStorageHandler(service.StorageServiceInterface),
	}
}
