package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
)

type StorageRepositoryInterace interface {
	Upload()
	Download()
}

type Repository struct {
	StorageRepositoryInterace
}

func NewRepository(s3Storage *minio.Client, db *pgxpool.Pool) *Repository {
	return &Repository{
		StorageRepositoryInterace: NewStorageRepository(s3Storage, db),
	}
}
