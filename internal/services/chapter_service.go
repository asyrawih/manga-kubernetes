package services

import (
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
)

type ChapterService struct {
	chapterRepo ports.ChapterRepository
}

func NewChapterService(chapterRepo ports.ChapterRepository) *ChapterService {
	return &ChapterService{
		chapterRepo: chapterRepo,
	}
}

// Get All Chapters
func (ch *ChapterService) DoGetChapters(mangaID string, args domain.QueryArgs) (chapters *domain.Chapters, err error) {
	c, err := ch.chapterRepo.GetChapters(mangaID, args)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Get One Chapter
func (ch *ChapterService) DoReadChapter(id string) (chap *domain.Chapter, err error) {
	c, err := ch.chapterRepo.ReadChapter(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Create Chapter
func (ch *ChapterService) DoCreateChapter(in *domain.CreateChapterRequest) error {
	return ch.chapterRepo.CreateChapter(in)
}

// Update Chapters
func (ch *ChapterService) DoUpdateChapters(in *domain.UpdateChapterRequest) error {
	return ch.chapterRepo.UpdateChapters(in)
}
