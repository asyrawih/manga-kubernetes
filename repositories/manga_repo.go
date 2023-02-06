package repositories

import (
	"context"
	"database/sql"

	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/rs/zerolog/log"
)

type MangaRepo struct {
	db *sql.DB
}

// NewMangaRepo function  
func NewMangaRepo(db *sql.DB) *MangaRepo {
	return &MangaRepo{
		db: db,
	}
}

const getAll = "SELECT id , title , thumb , author , publisher , year_published , status , genre FROM manga m limit 100;"

// Get All Manga
func (ma *MangaRepo) GetAll() (*domain.Mangas, error) {
	m := make(domain.Mangas, 0)

	r, err := ma.db.QueryContext(context.Background(), getAll)
	if err != nil {
		return nil, err
	}

	for r.Next() {
		var manga domain.Manga
		if err := r.Scan(&manga.Id, &manga.Title, &manga.Thumb, &manga.Author, &manga.Publisher, &manga.YearPublished, &manga.Status, &manga.Genre); err != nil {
			log.Err(err).Caller().Msg("")
		}
		m = append(m, manga)
	}

	return &m, nil

}

const createQuery = "INSERT INTO manga (title, thumb, author, publisher, year_published, status, genre, create_by) VALUES( ?, ?, ?, ?, ?, ?, ?, ?); "

// Create Manga
func (ma *MangaRepo) Create(in *domain.UpdateRequest) error {
	_, err := ma.db.ExecContext(context.Background(), createQuery, in.Title, in.Thumb, in.Author, in.Publisher, in.YearPublished, in.Status, in.Genre, in.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

const updateQuery = "UPDATE manga SET title=?, thumb=?, author=?, publisher=?, year_published=?, status=?, genre=?, created_by=? WHERE id=?; "

// Update The Manga
func (ma *MangaRepo) Update(id int, in *domain.UpdateRequest) error {
	_, err := ma.db.ExecContext(context.Background(), updateQuery, in.Title, in.Thumb, in.Author, in.Publisher, in.YearPublished, in.Status, in.Genre, in.CreatedBy, id)
	if err != nil {
		return err
	}
	return nil
}

// Get manga By Id
func (ma *MangaRepo) GetById(id string) (*domain.Manga, error) {
	panic("not implemented") // TODO: Implement
}

// Get Manga By Author
func (ma *MangaRepo) GetByAuthor(author string) (*domain.Mangas, error) {
	panic("not implemented") // TODO: Implement
}

// Search Manga by limit them 100 page
func (ma *MangaRepo) Search(title string) (*domain.Mangas, error) {
	panic("not implemented") // TODO: Implement
}

// Delete The Manga
func (ma *MangaRepo) Delete(mangaId string) error {
	panic("not implemented") // TODO: Implement
}
