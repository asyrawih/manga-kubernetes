package mysql

import (
	"context"
	"database/sql"

	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/utils"
)

type ChapterRepository struct {
	db *sql.DB
}

// NewChapterRepository function
func NewChapterRepository(db *sql.DB) *ChapterRepository {
	return &ChapterRepository{
		db: db,
	}
}

// Get All Chapters
func (ch *ChapterRepository) GetChapters(mangaId string, args domain.QueryArgs) (chapters *domain.Chapters, err error) {
	chaps := make(domain.Chapters, 0)
	const query = `SELECT * from chapters c WHERE c.manga_id = ?  `
	var mainQuery string
	mainQuery = query

	if args.OrderBy != "" {
		f := utils.WithOrderBy(mainQuery)
		s := f("c.id", domain.Desc)
		mainQuery = s
	}

	if args.Limit != "" && args.Offset != "" {
		f := utils.WithLimit(mainQuery)
		s := f(args.Limit, args.Offset)
		mainQuery = s
	}

	ctx := context.Background()
	r, err := ch.db.QueryContext(ctx, mainQuery, mangaId)
	if err != nil {
		return nil, err
	}

	for r.Next() {
		c := new(domain.Chapter)
		var content string
		if err := r.Scan(&c.Id, &c.MangaId, &c.ChapterNumber, &c.Title, &content); err != nil {
			return nil, err
		}

		c.Images = append(c.Images, domain.Content(content))

		chaps = append(chaps, *c)
	}

	chapters = &chaps
	return
}

// Get One Chapter
func (ch *ChapterRepository) ReadChapter(id string) (chap *domain.Chapter, err error) {
	panic("not implemented") // TODO: Implement
}

// Create Chapter
func (ch *ChapterRepository) CreateChapter(in *domain.CreateChapterRequest) error {
	panic("not implemented") // TODO: Implement
}

// Update Chapters
func (ch *ChapterRepository) UpdateChapters(in *domain.UpdateChapterRequest) error {
	panic("not implemented") // TODO: Implement
}
