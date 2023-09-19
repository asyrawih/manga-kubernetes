package services

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/asyrawih/manga/mocks"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	retensionCode := m.Run()
	os.Exit(retensionCode)
}

func TestMangaService_DoCreate(t *testing.T) {
	mangaMock := mocks.NewMangaRepository(t)

	type fields struct {
		mangaRepo ports.MangaRepository
		config    *config.Config
	}
	type args struct {
		in *domain.CreateRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantSomeErrr error
	}{
		{
			name: "should error happen",
			fields: fields{
				mangaRepo: mangaMock,
				config:    &config.Config{},
			},
			args: args{
				in: &domain.CreateRequest{},
			},
			wantErr:      true,
			wantSomeErrr: errors.New("some"),
		},
		{
			name: "should not error",
			fields: fields{
				mangaRepo: mangaMock,
				config:    &config.Config{},
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
			wantErr:      false,
			wantSomeErrr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := &MangaService{
				mangaRepo: tt.fields.mangaRepo,
				config:    tt.fields.config,
			}
			mangaMock.On("Create", tt.args.in).Return(tt.wantSomeErrr)
			if err := ma.DoCreate(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("MangaService.DoCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMangaService_DoGetAll(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	tests := []struct {
		name       string
		fields     fields
		want       *domain.Mangas
		wantErr    bool
		beforeFunc func(want *domain.Mangas, f fields)
	}{
		{
			name: "should oke",
			fields: fields{
				config: &config.Config{},
			},
			want: &domain.Mangas{
				{
					ID:            "1",
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
			beforeFunc: func(want *domain.Mangas, f fields) {
				mangaMock := mocks.NewMangaRepository(t)
				ma := &MangaService{
					mangaRepo: mangaMock,
					config:    f.config,
				}
				mangaMock.On("GetAll").Return(want, nil)

				got, err := ma.DoGetAll()
				assert.NoError(t, err)
				assert.NotNil(t, got)
			},
		},

		{
			name: "should not oke and mangas will be return nil",
			fields: fields{
				config: &config.Config{},
			},
			want:    &domain.Mangas{},
			wantErr: true,
			beforeFunc: func(_ *domain.Mangas, f fields) {
				mangaMock := mocks.NewMangaRepository(t)
				ma := &MangaService{
					mangaRepo: mangaMock,
					config:    f.config,
				}
				mangaMock.On("GetAll").Return(nil, errors.New("some"))

				got, err := ma.DoGetAll()
				assert.Nil(t, got)
				assert.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.beforeFunc(tt.want, tt.fields)
		})
	}
}

func TestMangaService_DoUpdate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		id int
		in *domain.UpdateRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		beforeFunc func(args args, f fields)
	}{
		{
			name: "if have any issue when update into database should error",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				id: 1,
				in: &domain.UpdateRequest{
					Title:         "any",
					Thumb:         "any",
					Genre:         "any",
					Author:        "any",
					Publisher:     "any",
					YearPublished: "any",
					Status:        "any",
					CreatedBy:     "any",
				},
			},
			wantErr: true,
			beforeFunc: func(args args, f fields) {
				mangaMock := mocks.NewMangaRepository(t)
				ma := &MangaService{
					mangaRepo: mangaMock,
					config:    f.config,
				}

				mangaMock.On("Update", mock.Anything, args.in).Return(errors.New("Error Happen From Database"))
				err := ma.DoUpdate(args.id, args.in)
				assert.Error(t, err)
			},
		},
		{
			name: "should oke if call this DoUpdate",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				id: 1,
				in: &domain.UpdateRequest{
					Title:         "any",
					Thumb:         "any",
					Genre:         "any",
					Author:        "any",
					Publisher:     "any",
					YearPublished: "any",
					Status:        "any",
					CreatedBy:     "any",
				},
			},
			wantErr: false,
			beforeFunc: func(args args, f fields) {
				mangaMock := mocks.NewMangaRepository(t)
				ma := &MangaService{
					mangaRepo: mangaMock,
					config:    f.config,
				}

				mangaMock.On("Update", mock.AnythingOfTypeArgument("int"), args.in).Return(nil)
				err := ma.DoUpdate(args.id, args.in)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.beforeFunc(tt.args, tt.fields)
		})
	}
}

func TestMangaService_DoGetByID(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		id string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *domain.Manga
		beforeFunc func(args args, f fields, want *domain.Manga)
	}{
		{
			name: "should oke",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				id: "1",
			},
			want: &domain.Manga{
				ID:            "1",
				Title:         "any",
				Thumb:         "any",
				Genre:         "any",
				Author:        "any",
				Publisher:     "any",
				YearPublished: "any",
				Status:        "any",
				CreatedBy:     "any",
			},
			beforeFunc: func(args args, f fields, want *domain.Manga) {
				mr := mocks.NewMangaRepository(t)
				mr.On("GetById", mock.AnythingOfType("string")).Return(want, nil)
				ma := &MangaService{
					mangaRepo: mr,
					config:    f.config,
				}
				got, err := ma.DoGetByID(args.id)
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, got.ID, "1")
			},
		},

		{
			name: "should return an error",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				id: "1",
			},
			want: nil,
			beforeFunc: func(args args, f fields, want *domain.Manga) {
				mr := mocks.NewMangaRepository(t)
				mr.On("GetById", mock.AnythingOfType("string")).Return(want, errors.New("some error will happen in here"))
				ma := &MangaService{
					mangaRepo: mr,
					config:    f.config,
				}
				got, err := ma.DoGetByID(args.id)
				assert.Error(t, err)
				assert.Nil(t, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.beforeFunc(tt.args, tt.fields, tt.want)
		})
	}
}

func TestMangaService_DoGetByAuthor(t *testing.T) {
	mangas := GenerateRandomManga()

	zerolog.SetGlobalLevel(zerolog.Disabled)
	type fields struct {
		config *config.Config
	}
	type args struct {
		author string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *domain.Mangas
		beforeFunc func(args args, f fields, want *domain.Mangas)
	}{
		{
			name:   "should get manga",
			fields: fields{},
			args: args{
				author: "any",
			},
			want: mangas,
			beforeFunc: func(args args, f fields, want *domain.Mangas) {
				mr := mocks.NewMangaRepository(t)

				// Program the the behavior of mock
				mr.On("GetByAuthor", mock.Anything).Return(want, nil)

				ma := &MangaService{
					mangaRepo: mr,
					config:    f.config,
				}

				got, err := ma.DoGetByAuthor(args.author)
				assert.NoError(t, err)
				assert.NotNil(t, got)

				for _, val := range *got {
					assert.Equal(t, val.Author, "any")
				}
			},
		},

		{
			name: "error must should handled",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				author: "any",
			},
			want: &domain.Mangas{},
			beforeFunc: func(args args, f fields, want *domain.Mangas) {
				mr := mocks.NewMangaRepository(t)

				// Program the the behavior of mock
				mr.On("GetByAuthor", mock.Anything).Return(want, errors.New("error happan"))

				ma := &MangaService{
					mangaRepo: mr,
					config:    f.config,
				}

				got, err := ma.DoGetByAuthor(args.author)
				assert.Error(t, err)
				assert.Nil(t, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.beforeFunc(tt.args, tt.fields, tt.want)
		})
	}
}

func GenerateRandomManga() *domain.Mangas {
	mangas := &domain.Mangas{
		{
			ID:            "1",
			Title:         "any",
			Thumb:         "any",
			Genre:         "any",
			Author:        "any",
			Publisher:     "any",
			YearPublished: "any",
			Status:        "any",
			CreatedBy:     "any",
		},
		{
			ID:            "2",
			Title:         "any2",
			Thumb:         "any2",
			Genre:         "any2",
			Author:        "any",
			Publisher:     "any2",
			YearPublished: "any2",
			Status:        "any2",
			CreatedBy:     "any2",
		},
	}
	return mangas
}

func TestMangaService_DoSearch(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		title string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *domain.Mangas
		wantErr    bool
		beforeFunc func(args args, f fields, want *domain.Mangas, wantErr bool)
	}{
		{
			name: "should oke Search Manga",
			fields: fields{
				config: &config.Config{},
			},
			args:    args{},
			want:    GenerateRandomManga(),
			wantErr: false,
			beforeFunc: func(args args, f fields, want *domain.Mangas, _ bool) {
				mr := mocks.NewMangaRepository(t)

				mr.On("Search", mock.AnythingOfType("string")).Return(want, nil)

				ma := &MangaService{
					mangaRepo: mr,
					config:    f.config,
				}
				got, err := ma.DoSearch(args.title)
				assert.NoError(t, err)
				assert.NotNil(t, got)
			},
		},
		{
			name: "should error if not data found in result",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				title: "not found",
			},
			want:    GenerateRandomManga(),
			wantErr: true,
			beforeFunc: func(args args, f fields, _ *domain.Mangas, _ bool) {
				mr := mocks.NewMangaRepository(t)

				mr.On("Search", mock.AnythingOfType("string")).Return(nil, errors.New("sql not result set on row"))

				ma := &MangaService{
					mangaRepo: mr,
					config:    f.config,
				}
				got, err := ma.DoSearch(args.title)
				assert.Error(t, err)
				assert.Nil(t, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.beforeFunc(tt.args, tt.fields, tt.want, tt.wantErr)
		})
	}
}

func TestMangaService_DoDelete(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		mangaID string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		beforeFunc func(arg args, field fields)
	}{
		{
			name: "should oke",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				mangaID: "1",
			},
			beforeFunc: func(arg args, field fields) {
				mr := mocks.NewMangaRepository(t)
				ma := &MangaService{
					mangaRepo: mr,
					config:    field.config,
				}

				mr.On("Delete", mock.AnythingOfType("string")).Return(nil)

				err := ma.DoDelete(arg.mangaID)
				assert.NoError(t, err)
			},
		},

		{
			name: "should oke",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				mangaID: "1",
			},
			beforeFunc: func(arg args, field fields) {
				mr := mocks.NewMangaRepository(t)
				ma := &MangaService{
					mangaRepo: mr,
					config:    field.config,
				}
				mr.On("Delete", mock.AnythingOfType("string")).Return(errors.New("some error happen"))

				err := ma.DoDelete(arg.mangaID)

				assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.beforeFunc(tt.args, tt.fields)
		})
	}
}

func TestNewMangaService(t *testing.T) {
	mr := mocks.NewMangaRepository(t)
	type args struct {
		mangaRepo ports.MangaRepository
		config    *config.Config
	}
	tests := []struct {
		name string
		args args
		want *MangaService
	}{
		{
			name: "oke create new intance",
			args: args{
				mangaRepo: mr,
				config:    &config.Config{},
			},
			want: &MangaService{
				mangaRepo: mr,
				config:    &config.Config{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMangaService(tt.args.mangaRepo, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMangaService() = %v, want %v", got, tt.want)
			}
		})
	}
}
