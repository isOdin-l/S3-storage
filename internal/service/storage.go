package service

import "github.com/isOdin-l/S3-storage/internal/repository"

type StorageService struct {
	repo repository.StorageRepositoryInterace
}

func NewStorageService(repo repository.StorageRepositoryInterace) *StorageService {
	return &StorageService{repo: repo}
}

func (s *StorageService) Upload() {

}

func (s *StorageService) Download() {

}
