package services

import (
	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/rs/zerolog/log"
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
		log.Err(err).Caller().Msg("")
		return err
	}
	return nil
}

// Get All Manga
func (ma *MangaService) DoGetAll() (*domain.Mangas, error) {
	m, err := ma.mangaRepo.GetAll()
	if err != nil {
		log.Err(err).Caller().Msg("")
		return nil, err
	}
	return m, nil
}

// Update Manga
func (ma *MangaService) DoUpdate(id int, in *domain.UpdateRequest) error {
	if err := ma.mangaRepo.Update(id, in); err != nil {
		log.Err(err).Caller().Msg("")
		return err
	}
	return nil
}

// Get By Manga
func (ma *MangaService) DoGetByID(id string) (*domain.Manga, error) {
	m, err := ma.mangaRepo.GetById(id)
	if err != nil {
		log.Err(err).Caller().Msg("")
		return nil, err
	}
	return m, nil
}

// Get By Manga Author
func (ma *MangaService) DoGetByAuthor(author string) (*domain.Mangas, error) {
	mangas, err := ma.mangaRepo.GetByAuthor(author)
	if err != nil {
		log.Err(err).Caller().Msg("")
		return nil, err
	}
	return mangas, nil
}

// Get By Manga Title
func (ma *MangaService) DoSearch(title string) (*domain.Mangas, error) {
	mangas, err := ma.mangaRepo.Search(title)
	if err != nil {
		log.Err(err).Caller().Msg("")
		return nil, err
	}
	return mangas, nil
}

// Delete Manga
func (ma *MangaService) DoDelete(mangaID string) error {
	err := ma.mangaRepo.Delete(mangaID)
	if err != nil {
		log.Err(err).Caller().Msg("")
		return err
	}
	return nil
}
