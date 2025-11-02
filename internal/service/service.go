package service

import (
	"mime/multipart"

	"github.com/isOdin-l/S3-storage/internal/repository"
	"github.com/isOdin-l/S3-storage/internal/types"
)

type StorageServiceInterface interface {
	Upload(medCerfFile *multipart.File, fileHeader *multipart.FileHeader, studentInfo *types.StudentInfo) error
}

type Service struct {
	StorageServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		StorageServiceInterface: NewStorageService(repo.StorageRepositoryInterace),
	}
}
