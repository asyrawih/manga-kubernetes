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

func NewMangaService(mangaRepo ports.MangaRepository, config *config.Config) *MangaService {
	return &MangaService{
		mangaRepo: mangaRepo,
		config:    config,
	}
}

// Create Manga by accept *domain.CreateRequest as arguments
func (ma *MangaService) DoCreate(in *domain.CreateRequest) error {
	if err := ma.mangaRepo.Create(in); err != nil {
		return err
	}
	return nil
}

// Get All Manga
func (ma *MangaService) DoGetAll() (*domain.Mangas, error) {
	m, err := ma.mangaRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Update Manga
func (ma *MangaService) DoUpdate(id int, in *domain.UpdateRequest) error {
	if err := ma.mangaRepo.Update(id, in); err != nil {
		return err
	}
	return nil
}
