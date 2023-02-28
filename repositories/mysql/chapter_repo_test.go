package mysql

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/utils"
	"github.com/stretchr/testify/assert"
)

func TestChapterRepository_GetChapters(t *testing.T) {

	type args struct {
		mangaId string
		args    domain.QueryArgs
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		beforeFunc func(args)
	}{
		{
			name: "should return chapters",
			args: args{
				mangaId: "1",
				args:    domain.QueryArgs{},
			},
			wantErr: false,
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				defer db.Close()

				query := "SELECT * from chapters c WHERE c.manga_id = ?"
				rows := sqlmock.NewRows([]string{"id", "manga_id", "chapter_number", "title", "content"})

				rows.AddRow("1", "1", "1", "Solo Leveling", "more content")

				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

				ch := &ChapterRepository{
					db: db,
				}
				gotChapters, err := ch.GetChapters(args.mangaId, args.args)
				assert.NoError(t, err)
				assert.NotNil(t, gotChapters)
			},
		},

		{
			name: "should return chapters error cus the query not valid",
			args: args{
				mangaId: "1",
				args:    domain.QueryArgs{},
			},
			wantErr: false,
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				defer db.Close()

				query := "SELECT * from chapters c WHERE c.manga_id = ?"
				rows := sqlmock.NewRows([]string{"id", "manga_id", "chapter_number", "title", "content"})

				rows.AddRow("1", "1", "1", "Solo Leveling", "more content")

				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sqlmock.ErrCancelled)

				ch := &ChapterRepository{
					db: db,
				}
				gotChapters, err := ch.GetChapters(args.mangaId, args.args)
				assert.Error(t, err)
				assert.Nil(t, gotChapters)
			},
		},
		{
			name: "should oke with order by",
			args: args{
				mangaId: "1",
				args: domain.QueryArgs{
					Limit:   "",
					Offset:  "",
					OrderBy: domain.OrderBy(domain.Desc),
				},
			},
			wantErr: false,
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				defer db.Close()

				query := "SELECT * from chapters c WHERE c.manga_id = ?"
				f := utils.WithOrderBy(query)
				withOrderBy := f("c.id", string(args.args.OrderBy))
				rows := sqlmock.NewRows([]string{"id", "manga_id", "chapter_number", "title", "content"})

				rows.AddRow("1", "1", "1", "Solo Leveling", "more content")

				mock.ExpectQuery(regexp.QuoteMeta(withOrderBy)).WillReturnRows(rows)

				ch := &ChapterRepository{
					db: db,
				}
				gotChapters, err := ch.GetChapters(args.mangaId, args.args)
				assert.NoError(t, err)
				assert.NotNil(t, gotChapters)
			},
		},
		{
			name: "should oke with with limit",
			args: args{
				mangaId: "1",
				args: domain.QueryArgs{
					Limit:  "1",
					Offset: "15",
				},
			},
			wantErr: false,
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				defer db.Close()

				query := "SELECT * from chapters c WHERE c.manga_id = ?"
				f := utils.WithLimit(query)
				withLimit := f(args.args.Limit, args.args.Offset)

				rows := sqlmock.NewRows([]string{"id", "manga_id", "chapter_number", "title", "content"})

				rows.AddRow("1", "1", "1", "Solo Leveling", "more content")
				rows.AddRow("2", "1", "1", "Solo Leveling", "more content")

				mock.ExpectQuery(regexp.QuoteMeta(withLimit)).WillReturnRows(rows)

				ch := &ChapterRepository{
					db: db,
				}
				gotChapters, err := ch.GetChapters(args.mangaId, args.args)
				assert.NoError(t, err)
				assert.NotNil(t, gotChapters)
			},
		},
		{
			name: "should error when scan the row has error",
			args: args{
				mangaId: "1",
				args: domain.QueryArgs{
					Limit:  "1",
					Offset: "15",
				},
			},
			wantErr: false,
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				defer db.Close()

				query := "SELECT * from chapters c WHERE c.manga_id = ?"
				f := utils.WithLimit(query)
				withLimit := f(args.args.Limit, args.args.Offset)

				rows := sqlmock.NewRows([]string{"id", "manga_id", "chapter_number", "title", "content"})

				rows.AddRow("1", "1", "1", "Solo Leveling", nil)

				mock.ExpectQuery(regexp.QuoteMeta(withLimit)).WillReturnRows(rows)

				ch := &ChapterRepository{
					db: db,
				}
				gotChapters, err := ch.GetChapters(args.mangaId, args.args)
				assert.Error(t, err)
				assert.Nil(t, gotChapters)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeFunc(tt.args)
		})
	}
}

func TestChapterRepository_ReadChapter(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name       string
		args       args
		beforeFunc func(args args)
	}{
		{
			name: "should ok get detail chapter",
			args: args{
				id: "1",
			},
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				query := "SELECT * from chapters c where c.id  = ?"
				// id|manga_id|chapter_number|title|content|
				rows := sqlmock.NewRows([]string{"id", "manga_id", "chapter_number", "title", "content"})
				rows.AddRow("1", "1", "1", "test", `https://someimage.com/apa.png`)

				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

				assert.NoError(t, err)
				ch := &ChapterRepository{
					db: db,
				}
				gotChap, err := ch.ReadChapter(args.id)
				t.Log(string(gotChap.Images))
				assert.NoError(t, err)
				assert.NotNil(t, gotChap)
			},
		},

		{
			name: "should return an error if got scan error",
			args: args{
				id: "1",
			},
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				query := "SELECT * from chapters c where c.id  = ?"
				// id|manga_id|chapter_number|title|content|
				rows := sqlmock.NewRows([]string{"id", "manga_id", "chapter_number", "title", "content"})
				rows.AddRow("1", nil, nil, "test", "https://someimage.com/apa.png")

				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

				assert.NoError(t, err)
				ch := &ChapterRepository{
					db: db,
				}
				gotChap, err := ch.ReadChapter(args.id)
				assert.Error(t, err)
				assert.Nil(t, gotChap)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeFunc(tt.args)
		})
	}
}

func TestChapterRepository_CreateChapter(t *testing.T) {
	type args struct {
		in *domain.CreateChapterRequest
	}
	tests := []struct {
		name       string
		args       args
		beforeFunc func(args args)
	}{
		{
			name: "should oke insert the chapter",
			args: args{
				in: &domain.CreateChapterRequest{
					MangaId:       "1",
					Title:         "Nisa Manga Test",
					ChapterNumber: 1,
					Images:        []string{"<img src='someurl'/>"},
				},
			},
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)

				const query = "INSERT INTO chapters (manga_id, chapter_number, title, content) VALUES(?, ?, ?, ?)"

				mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))

				ch := &ChapterRepository{
					db: db,
				}

				err = ch.CreateChapter(args.in)
				assert.NoError(t, err)

			},
		},
		{
			name: "should error invalid type insertion",
			args: args{
				in: &domain.CreateChapterRequest{
					MangaId:       "1",
					Title:         "Nisa Manga Test",
					ChapterNumber: 1,
					Images:        []string{"<img src='someurl'/>"},
				},
			},
			beforeFunc: func(args args) {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)

				const query = "INSERT INTO chapters (manga_id, chapter_number, title, content) VALUES(?, ?, ?, ?)"

				mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(errors.New("some error"))

				ch := &ChapterRepository{
					db: db,
				}

				err = ch.CreateChapter(args.in)
				assert.Error(t, err)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeFunc(tt.args)
		})
	}
}
