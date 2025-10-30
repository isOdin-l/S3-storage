package main

import (
	"github.com/isOdin-l/S3-storage/api/handler"
	"github.com/isOdin-l/S3-storage/internal/databases/postgresql"
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

	// Postgres
	metadataDb, err := postgresql.NewPostgresDB(&postgresql.Config{
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		Username: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		Database: viper.GetString("POSTGRES_DB"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize metadata-db: %s", err.Error())
	}
	defer metadataDb.Close()

	//MinIO

	// repository
	repository := repository.NewRepository()

	// service
	service := service.NewService(repository)

	// handler
	handler := handler.NewHandler(service)

	// router
	router := router.NewRouter(handler)

	// server
	server := server.New()
	if err := server.Run(viper.GetString("SERVER_PORT"), router); err != nil {
		logrus.Fatalf("error while running server %s", err.Error())
	}
}

func init_config() error {
	if err := godotenv.Load("./configs/server.env", "./configs/s3-storage.env", "./configs/metadata-storage.env"); err != nil {
		return err
	}

	viper.AutomaticEnv()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	return nil
}
