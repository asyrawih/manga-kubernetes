package services

import (
	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
)

type MangaService struct {
	mangaRepo ports.MangaRepository
	config    *config.Config
}

// Create Manga by accept *domain.CreateRequest as arguments
func (ma *MangaService) DoCreate(in *domain.CreateRequest) error {
	return ma.mangaRepo.Create(in)
}

// Get All Manga
func (ma *MangaService) DoGetAll() (*domain.Mangas, error) {
	panic("not implemented") // TODO: Implement
}

func NewMangaService(mangaRepo ports.MangaRepository, config *config.Config) *MangaService {
	return &MangaService{
		mangaRepo: mangaRepo,
		config:    config,
	}
}
