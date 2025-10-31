package repository

import (
	"context"
	"mime/multipart"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
)

type StorageRepository struct {
	s3 *minio.Client
	db *pgxpool.Pool
}

func NewStorageRepository(s3Storage *minio.Client, db *pgxpool.Pool) *StorageRepository {
	return &StorageRepository{s3: s3Storage, db: db}
}

func (r *StorageRepository) Upload(medCerfFile *multipart.File, newFileName string, fileHeader *multipart.FileHeader) error {
	_, err := r.s3.PutObject(context.Background(), viper.GetString("S3_BUCKET"), newFileName, *medCerfFile, fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: fileHeader.Header.Get("Content-Type"),
		})

	return err
}

func (r *StorageRepository) Download() {

}
