package minio_storage

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Port      string
	Region    string
}

func NewMinioDB(configStorage *S3Config) (*minio.Client, error) {
	fullPathEndpoint := fmt.Sprintf("%s:%s", configStorage.Endpoint, configStorage.Port)

	minioStorage, err := minio.New(fullPathEndpoint, &minio.Options{
		Region: configStorage.Region,
		Creds:  credentials.NewStaticV4(configStorage.AccessKey, configStorage.SecretKey, ""),
	})

	if err != nil {
		return nil, err
	}

	return minioStorage, nil
}
