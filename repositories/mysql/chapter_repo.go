package mysql

import (
	"context"
	"database/sql"
	"strings"

	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/utils"
)

type ChapterRepository struct {
	db *sql.DB
}

// NewChapterRepository function
func NewChapterRepo(db *sql.DB) *ChapterRepository {
	return &ChapterRepository{
		db: db,
	}
}

// Get All Chapters
func (ch *ChapterRepository) GetChapters(mangaId string, args domain.QueryArgs) (chapters *[]domain.Chapter, err error) {
	chaps := make([]domain.Chapter, 0)
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

		c.Images = domain.Content(content)

		chaps = append(chaps, *c)
	}

	chapters = &chaps
	return
}

// Get One Chapter
func (ch *ChapterRepository) ReadChapter(id string) (*domain.Chapter, error) {
	c := new(domain.Chapter)
	query := "SELECT * from chapters c where c.id  = ?"

	ctx := context.Background()

	r := ch.db.QueryRowContext(ctx, query, id)

	var content string

	if err := r.Scan(&c.Id, &c.MangaId, &c.ChapterNumber, &c.Title, &content); err != nil {
		return nil, err
	}

	c.Images = domain.Content(content)

	return c, nil

}

// Create Chapter
func (ch *ChapterRepository) CreateChapter(in *domain.CreateChapterRequest) error {
	const query = "INSERT INTO chapters (manga_id, chapter_number, title, content) VALUES(?, ?, ?, ?)"

	ctx := context.Background()

	imageStrings := make([]string, len(in.Images))

	for i, image := range in.Images {
		imageStrings[i] = string(image)
	}

	imageString := strings.Join(imageStrings, ",")

	_, err := ch.db.ExecContext(ctx, query, in.MangaId, in.ChapterNumber, in.Title, &imageString)
	if err != nil {
		return err
	}
	return nil

}

// Update Chapters
func (ch *ChapterRepository) UpdateChapters(in *domain.UpdateChapterRequest) error {
	panic("not implemented") // TODO: Implement
}
