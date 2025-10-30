package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/isOdin-l/S3-storage/api/handler"
)

func NewRouter(h *handler.Handler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Post("/upload", h.StorageHandlerInterface.Upload)
	r.Get("/download", h.StorageHandlerInterface.Download)

	return r
}
