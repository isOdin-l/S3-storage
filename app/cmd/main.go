package main

import (
	"context"
	"net/http"

	"github.com/isOdin-l/S3-storage/api/handler"
	"github.com/isOdin-l/S3-storage/internal/databases/minio_storage"
	"github.com/isOdin-l/S3-storage/internal/repository"
	"github.com/isOdin-l/S3-storage/internal/router"
	"github.com/isOdin-l/S3-storage/internal/server"
	"github.com/isOdin-l/S3-storage/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// configuration
	init_config()

	//MinIO
	s3storage, err := minio_storage.NewMinioDB(&minio_storage.S3Config{
		Endpoint:  viper.GetString("S3_ENDPOINT"),
		AccessKey: viper.GetString("S3_ACCESS_KEY"),
		SecretKey: viper.GetString("S3_SECRET_KEY"),
		Port:      viper.GetString("S3_PORT"),
		Region:    viper.GetString("S3_REGION"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize s3-storage: %s", err.Error())
	}

	// repository
	repository := repository.NewRepository(s3storage)

	// service
	service := service.NewService(repository)

	// handler
	handler := handler.NewHandler(service)

	// router
	router := router.NewRouter(handler)

	// server
	logrus.Info("Server starting...")
	server := server.New()
	go func() {
		if err := server.Run(viper.GetString("SERVER_PORT"), router); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error while running server %s", err.Error())
		}
	}()

	server.GracefulShutdown(context.Background())
}

func init_config() error {
	if err := godotenv.Load("./configs/server.env", "./configs/s3-storage.env", "./configs/metadata-storage.env"); err != nil {
		return err
	}

	viper.AutomaticEnv()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	return nil
}
