package service

import (
	"mime/multipart"

	"github.com/isOdin-l/S3-storage/internal/repository"
)

type StorageServiceInterface interface {
	Upload(medCerfFile *multipart.File, fileHeader *multipart.FileHeader) error
	Download()
}

type Service struct {
	StorageServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		StorageServiceInterface: NewStorageService(repo.StorageRepositoryInterace),
	}
}
