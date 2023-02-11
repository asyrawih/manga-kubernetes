package ports

import "github.com/asyrawih/manga/internal/core/domain"

type MangaRepository interface {
	// Get All Manga
	GetAll() (*domain.Mangas, error)
	// Create Manga
	Create(in *domain.CreateRequest) error
	// Update The Manga
	Update(id int, in *domain.UpdateRequest) error
	// Get manga By Id
	GetById(id string) (*domain.Manga, error)
	// Get Manga By Author
	GetByAuthor(author string) (*domain.Mangas, error)
	// Search Manga by limit them 100 page
	Search(title string) (*domain.Mangas, error)
	// Delete The Manga
	Delete(mangaId string) error
}

type MangaService interface {
	// Create Manga by accept *domain.CreateRequest as arguments
	DoCreate(in *domain.CreateRequest) error
	// Get All Manga
	DoGetAll() (*domain.Mangas, error)
}
