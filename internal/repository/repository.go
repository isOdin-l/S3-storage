package repository

type StorageRepositoryInterace interface {
	Upload()
	Download()
}

type Repository struct {
	StorageRepositoryInterace
}

func NewRepository() *Repository {
	return &Repository{
		StorageRepositoryInterace: NewStorageRepository(),
	}
}
