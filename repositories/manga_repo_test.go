package repositories

import (
	"database/sql"
	"testing"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/pkg/dbconn"
	"github.com/stretchr/testify/assert"
)

func TestMangaRepo_GetAll(t *testing.T) {
	c := config.LoadConfig("../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *domain.Mangas
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				db: db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := &MangaRepo{
				db: tt.fields.db,
			}
			got, err := ma.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("MangaRepo.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			t.Log(got)
		})
	}
}

func TestMangaRepo_Create(t *testing.T) {

	c := config.LoadConfig("../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		in *domain.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				db: db,
			},
			args: args{
				in: &domain.UpdateRequest{
					Title:         "test",
					Thumb:         "test",
					Genre:         "Manga",
					Author:        "test",
					Publisher:     "test",
					YearPublished: "test",
					Status:        "Publish",
					CreatedBy:     "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := &MangaRepo{
				db: tt.fields.db,
			}
			if err := ma.Create(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("MangaRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMangaRepo_Update(t *testing.T) {
	c := config.LoadConfig("../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int
		in *domain.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should oke update manga",
			fields: fields{
				db: db,
			},
			args: args{
				id: 1,
				in: &domain.UpdateRequest{
					Title:         "Test",
					Thumb:         "Test",
					Genre:         "Manga",
					Author:        "Test",
					Publisher:     "Test",
					YearPublished: "Test",
					Status:        "Publish",
					CreatedBy:     "Test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := &MangaRepo{
				db: tt.fields.db,
			}
			if err := ma.Update(tt.args.id, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("MangaRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
