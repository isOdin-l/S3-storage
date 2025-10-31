package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/isOdin-l/S3-storage/internal/repository"
)

type StorageService struct {
	repo repository.StorageRepositoryInterace
}

func NewStorageService(repo repository.StorageRepositoryInterace) *StorageService {
	return &StorageService{repo: repo}
}

func (s *StorageService) Upload(medCerfFile *multipart.File, fileHeader *multipart.FileHeader) error {
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(fileHeader.Filename))
	return s.repo.Upload(medCerfFile, newFileName, fileHeader)
}

func (s *StorageService) Download() {

}
