package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/asyrawih/manga/internal/core/domain"
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

const createQuery = "INSERT INTO manga (title, thumb, author, publisher, year_published, status, genre, created_by) VALUES( ?, ?, ?, ?, ?, ?, ?, ?); "

// Create Manga
func (ma *MangaRepo) Create(in *domain.CreateRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	errChan := make(chan error, 1)
	done := make(chan struct{})
	go func() {
		result, err := ma.db.ExecContext(
			ctx,
			createQuery,
			in.Title,
			in.Thumb,
			in.Author,
			in.Publisher,
			in.YearPublished,
			in.Status,
			in.Genre,
			in.CreatedBy,
		)
		if err != nil {
			errChan <- err
		}

		i, err := result.RowsAffected()
		if err != nil {
			errChan <- err
		}

		log.Info().Msgf("affected row %d", i)

		done <- struct{}{}
	}()

	for {
		select {
		case err := <-errChan:
			log.Err(err).Caller().Msg("")
			return err
		default:
			<-done
			return nil
		}
	}
}

const updateQuery = "UPDATE manga SET title=?, thumb=?, author=?, publisher=?, year_published=?, status=?, genre=?, created_by=? WHERE id=?; "

// Update The Manga
func (ma *MangaRepo) Update(id int, in *domain.UpdateRequest) error {
	Start := time.Now()
	errChan := make(chan error, 1)
	done := make(chan struct{})
	go func() {
		_, err := ma.db.ExecContext(
			context.Background(),
			updateQuery,
			in.Title,
			in.Thumb,
			in.Author,
			in.Publisher,
			in.YearPublished,
			in.Status,
			in.Genre,
			in.CreatedBy,
			id,
		)
		if err != nil {
			errChan <- err
		}
		done <- struct{}{}
	}()

	for {
		select {
		case someErr := <-errChan:
			fmt.Println(someErr.Error())
			break
		default:
			<-done
			fmt.Println(time.Since(Start).Abs().String())
			return nil
		}
	}
}

// Get manga By Id
func (ma *MangaRepo) GetById(id string) (*domain.Manga, error) {
	manga := new(domain.Manga)
	const query = "SELECT * from manga m WHERE m.id =?"
	r := ma.db.QueryRow(query, id)
	if err := r.Scan(
		&manga.Id,
		&manga.Title,
		&manga.Thumb,
		&manga.Author,
		&manga.Publisher,
		&manga.YearPublished,
		&manga.Status,
		&manga.Genre,
		&manga.CreatedBy,
	); err != nil {
		return nil, err
	}
	return manga, nil
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
