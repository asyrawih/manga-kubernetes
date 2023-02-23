package services

import (
	"database/sql"
	"testing"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/asyrawih/manga/pkg/dbconn"
	repositories "github.com/asyrawih/manga/repositories/mysql"
	"github.com/stretchr/testify/assert"
)

func TestUserService_DoCreateUser(t *testing.T) {

	c := config.LoadConfig("../../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)

	userRepo := repositories.NewUserRepo(db)

	defer cleanUp(db, t)

	type fields struct {
		userRepo ports.UserRepository
		config   *config.Config
	}
	type args struct {
		in *domain.CreateUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Should Oke Create User From Service Layer",
			fields: fields{
				userRepo: userRepo,
				config:   c,
			},
			args: args{
				in: &domain.CreateUser{
					Name:     "hanan",
					Username: "hanan",
					Email:    "hasyrawi@gmail.com",
					Password: "awiroot123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				userRepo: tt.fields.userRepo,
			}
			err := us.DoCreateUser(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.DoCreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func cleanUp(db *sql.DB, t *testing.T) {
	_, err := db.Exec("TRUNCATE TABLE users")
	assert.NoError(t, err)
}
