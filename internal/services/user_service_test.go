package services

import (
	"errors"
	"testing"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/mocks"
	"github.com/asyrawih/manga/pkg/password"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_DoCreateUser(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		in *domain.CreateUser
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		beforeFunc func(args args, fields fields)
	}{
		{
			name: "should oke create user",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				in: &domain.CreateUser{
					Name:     "hanan",
					Username: "hanan",
					Email:    "hanan@asyrawih.id",
					Password: "test123",
				},
			},
			beforeFunc: func(args args, fields fields) {

				ur := mocks.NewUserRepository(t)

				ur.On("CreateUser", args.in).Return(nil)

				us := &UserService{
					userRepo: ur,
					config:   fields.config,
				}

				err := us.DoCreateUser(args.in)
				assert.NoError(t, err)
				ur.AssertExpectations(t)

			},
		},
		{
			name: "should error if got some error",
			fields: fields{
				config: &config.Config{},
			},
			args: args{
				in: nil,
			},
			beforeFunc: func(args args, fields fields) {
				ur := mocks.NewUserRepository(t)
				ur.On("CreateUser", mock.Anything).Return(errors.New("some error"))

				us := &UserService{
					userRepo: ur,
					config:   fields.config,
				}

				err := us.DoCreateUser(args.in)
				assert.Error(t, err)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeFunc(tt.args, tt.fields)
		})
	}
}

func TestUserService_DoLogin(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		username string
		pass     string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *domain.UserLoginResponse
		wantErr    bool
		beforeFunc func(args args, fields fields)
	}{
		{
			name: "should oke login",
			fields: fields{
				config: &config.Config{
					Mysql: config.Mysql{},
					Key:   "a6GPsUW7BmOA63givx7ykMhfLf5fMwGP",
				},
			},
			args: args{
				username: "hanan",
				pass:     "$2a$14$EBbFxkn/pppj/sc8UrU/bOlWE0sGLvj2tLGo4XQX/STHjhyww08XO",
			},
			want:    &domain.UserLoginResponse{},
			wantErr: false,
			beforeFunc: func(args args, fields fields) {
				ur := mocks.NewUserRepository(t)

				s, err := password.HashPassword(args.pass)
				assert.NoError(t, err)

				ur.On("Login", args.username).Return(&domain.UserLogin{
					Username: "hanan",
					Password: s,
				}, nil)

				us := &UserService{
					userRepo: ur,
					config:   fields.config,
				}

				got, err := us.DoLogin(args.username, args.pass)
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.NotNil(t, got.Message)
				assert.NotNil(t, got.Token)
			},
		},
		{
			name: "password not match",
			fields: fields{
				config: &config.Config{
					Mysql: config.Mysql{},
					Key:   "a6GPsUW7BmOA63givx7ykMhfLf5fMwGP",
				},
			},
			args: args{
				username: "hanan",
				pass:     "$2a$14$EBbFxkn/pppj/sc8UrU/bOlWE0sGLvj2tLGo4XQX/STHjhyww08XO",
			},
			want:    &domain.UserLoginResponse{},
			wantErr: false,
			beforeFunc: func(args args, fields fields) {
				ur := mocks.NewUserRepository(t)

				ur.On("Login", args.username).Return(&domain.UserLogin{
					Username: "hanan",
					Password: "wrong_password",
				}, nil)

				us := &UserService{
					userRepo: ur,
					config:   fields.config,
				}

				got, err := us.DoLogin(args.username, args.pass)
				assert.Error(t, err)
				assert.Nil(t, got)
			},
		},
		{
			name: "call login but got an error",
			fields: fields{
				config: &config.Config{
					Mysql: config.Mysql{},
					Key:   "a6GPsUW7BmOA63givx7ykMhfLf5fMwGP",
				},
			},
			args: args{
				username: "hanan",
				pass:     "$2a$14$EBbFxkn/pppj/sc8UrU/bOlWE0sGLvj2tLGo4XQX/STHjhyww08XO",
			},
			want:    &domain.UserLoginResponse{},
			wantErr: false,
			beforeFunc: func(args args, fields fields) {
				ur := mocks.NewUserRepository(t)

				ur.On("Login", args.username).Return(nil, errors.New("error call login"))

				us := &UserService{
					userRepo: ur,
					config:   fields.config,
				}

				got, err := us.DoLogin(args.username, args.pass)
				assert.Error(t, err)
				assert.Nil(t, got)
			},
		},
		{
			name: "paseto key wrong ",
			fields: fields{
				config: &config.Config{
					Mysql: config.Mysql{},
					Key:   "",
				},
			},
			args: args{
				username: "hanan",
				pass:     "$2a$14$EBbFxkn/pppj/sc8UrU/bOlWE0sGLvj2tLGo4XQX/STHjhyww08XO",
			},
			want:    &domain.UserLoginResponse{},
			wantErr: false,
			beforeFunc: func(args args, fields fields) {
				ur := mocks.NewUserRepository(t)

				s, err := password.HashPassword(args.pass)
				assert.NoError(t, err)

				ur.On("Login", args.username).Return(&domain.UserLogin{
					Username: "hanan",
					Password: s,
				}, nil)

				us := &UserService{
					userRepo: ur,
					config:   fields.config,
				}

				got, err := us.DoLogin(args.username, args.pass)
				assert.NoError(t, err)
				assert.Equal(t, got.Token, "")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeFunc(tt.args, tt.fields)
		})
	}
}
