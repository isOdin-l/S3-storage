package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/isOdin-l/S3-storage/internal/repository"
	"github.com/isOdin-l/S3-storage/internal/types"
)

type StorageService struct {
	repo repository.StorageRepositoryInterace
}

func NewStorageService(repo repository.StorageRepositoryInterace) *StorageService {
	return &StorageService{repo: repo}
}

func (s *StorageService) Upload(medCerfFile *multipart.File, fileHeader *multipart.FileHeader, studentInfo *types.StudentInfo) error {
	studentFullName := fmt.Sprintf("%s %s %s", studentInfo.Surname, studentInfo.Name, studentInfo.Patronymic)
	studentFolder := fmt.Sprintf("%s/%s", studentInfo.GroupName, studentFullName)

	newFileName := fmt.Sprintf("%s/%s%s", studentFolder, uuid.New().String(), filepath.Ext(fileHeader.Filename))
	return s.repo.Upload(medCerfFile, newFileName, fileHeader)
}
