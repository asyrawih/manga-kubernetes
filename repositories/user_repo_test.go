package repositories

import (
	"database/sql"
	"testing"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/pkg/dbconn"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_CreateUser(t *testing.T) {

	c := config.LoadConfig("../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)

	defer cleanUp(db, t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		in *domain.CreateUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should oke",
			fields: fields{
				db: db,
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
		{
			name: "duplicate request",
			fields: fields{
				db: db,
			},
			args: args{
				in: &domain.CreateUser{
					Name:     "hanan",
					Username: "hanan",
					Email:    "hasyrawi@gmail.com",
					Password: "awiroot123",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserRepo{
				db: tt.fields.db,
			}
			err := us.CreateUser(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}

func cleanUp(db *sql.DB, t *testing.T) {
	_, err := db.Exec("TRUNCATE TABLE users")
	assert.NoError(t, err)
}

func TestUserRepo_GetUser(t *testing.T) {

	c := config.LoadConfig("../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)

	defer cleanUp(db, t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "should oke get user by username",
			fields: fields{
				db: db,
			},
			args: args{
				username: "hanan",
			},
			want:    &domain.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserRepo{
				db: tt.fields.db,
			}
			err := us.CreateUser(&domain.CreateUser{
				Name:     "hanan",
				Username: "hanan",
				Email:    "hasyrawi@gmail.com",
				Password: "awiroot123",
			})

			assert.NoError(t, err)

			got, err := us.GetUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestUserRepo_Login(t *testing.T) {
	c := config.LoadConfig("../config/config.json")
	db, err := dbconn.NewMySQLDB(c)
	assert.NoError(t, err)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.UserLogin
		wantErr bool
	}{
		{
			name: "get user by username with return username and password",
			fields: fields{
				db: db,
			},
			args: args{
				username: "asyrawi",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserRepo{
				db: tt.fields.db,
			}
			got, err := us.Login(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
