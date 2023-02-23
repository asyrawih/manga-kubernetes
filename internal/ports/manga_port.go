package ports

import (
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/labstack/echo/v4"
)

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
	// Update Manga
	DoUpdate(id int, in *domain.UpdateRequest) error
	// Get By Manga
	DoGetByID(id string) (*domain.Manga, error)
	// Get By Manga Author
	DoGetByAuthor(author string) (*domain.Mangas, error)
	// Get By Manga Title
	DoSearch(title string) (*domain.Mangas, error)
	// Delete Manga
	DoDelete(mangaID string) error
}

// MangaRoute interface Not Require for this interface
type MangaRoute interface {
	// Get All Manga
	GetAll(e echo.Context) error
	// Create Manga
	Create(e echo.Context) error
	// Update The Manga
	Update(e echo.Context) error
	// Get manga By Id
	GetById(e echo.Context) error
	// Get Manga By Author
	GetByAuthor(e echo.Context) error
	// Search Manga by limit them 100 page
	Search(e echo.Context) error
	// Delete The Manga
	Delete(e echo.Context) error
}
