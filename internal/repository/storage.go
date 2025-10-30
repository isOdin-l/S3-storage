package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
)

type StorageRepository struct {
	s3 *minio.Client
	db *pgxpool.Pool
}

func NewStorageRepository(s3Storage *minio.Client, db *pgxpool.Pool) *StorageRepository {
	return &StorageRepository{s3: s3Storage, db: db}
}

func (r *StorageRepository) Upload() {

}

func (r *StorageRepository) Download() {

}
