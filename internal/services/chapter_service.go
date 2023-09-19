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

	var chaps domain.Chapters

	for _, val := range *c {
		chaps = append(chaps, domain.ChapterResponse{
			ID:            val.ID,
			MangaID:       val.MangaID,
			ChapterNumber: val.ChapterNumber,
			Title:         val.Title,
			Images:        string(val.Images),
		})
	}

	chapters = &chaps

	return
}

// Get One Chapter
func (ch *ChapterService) DoReadChapter(id string) (chap *domain.ChapterResponse, err error) {
	c, err := ch.chapterRepo.ReadChapter(id)
	if err != nil {
		return nil, err
	}

	res := &domain.ChapterResponse{
		ID:            c.ID,
		MangaID:       c.MangaID,
		ChapterNumber: c.ChapterNumber,
		Title:         c.Title,
		Images:        string(c.Images),
	}

	return res, nil
}

// Create Chapter
func (ch *ChapterService) DoCreateChapter(in *domain.CreateChapterRequest) error {
	return ch.chapterRepo.CreateChapter(in)
}

// Update Chapters
func (ch *ChapterService) DoUpdateChapters(in *domain.UpdateChapterRequest) error {
	return ch.chapterRepo.UpdateChapters(in)
}
