package service

import "github.com/isOdin-l/S3-storage/internal/repository"

type StorageServiceInterface interface {
	Upload()
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
