package repository

import (
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type StorageRepositoryInterace interface {
	Upload(medCerfFile *multipart.File, newFileName string, fileHeader *multipart.FileHeader) error
}

type Repository struct {
	StorageRepositoryInterace
}

func NewRepository(s3Storage *minio.Client) *Repository {
	return &Repository{
		StorageRepositoryInterace: NewStorageRepository(s3Storage),
	}
}
