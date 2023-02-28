package ports

import "github.com/asyrawih/manga/internal/core/domain"

type ChapterRepository interface {
	// Get All Chapters
	GetChapters(mangaId string, args domain.QueryArgs) (chapters *domain.Chapters, err error)
	// Get One Chapter
	ReadChapter(id string) (chap *domain.Chapter, err error)
	// Create Chapter
	CreateChapter(in *domain.CreateChapterRequest) error
	// Update Chapters
	UpdateChapters(in *domain.UpdateChapterRequest) error
}

type ChapterService interface {
	// Get All Chapters
	DoGetChapters(mangaId string, args domain.QueryArgs) (chapters *domain.Chapters, err error)
	// Get One Chapter
	DoReadChapter(id string) (chap *domain.Chapter, err error)
	// Create Chapter
	DoCreateChapter(in *domain.CreateChapterRequest) error
	// Update Chapters
	DoUpdateChapters(in *domain.UpdateChapterRequest) error
}
