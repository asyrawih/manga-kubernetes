package services

import (
	"testing"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/asyrawih/manga/pkg/dbconn"
	"github.com/asyrawih/manga/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMangaService_DoCreate(t *testing.T) {
	c := config.LoadConfig("../../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)
	mangaRepo := repositories.NewMangaRepo(db)

	type fields struct {
		mangaRepo ports.MangaRepository
		config    *config.Config
	}
	type args struct {
		in *domain.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should oke create manga from service",
			fields: fields{
				mangaRepo: mangaRepo,
				config:    c,
			},
			args: args{
				in: &domain.CreateRequest{
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
			ma := &MangaService{
				mangaRepo: tt.fields.mangaRepo,
				config:    tt.fields.config,
			}
			if err := ma.DoCreate(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("MangaService.DoCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
