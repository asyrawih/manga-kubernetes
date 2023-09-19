package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/mocks"
	"github.com/asyrawih/manga/pkg/password"
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
		t.Run(tt.name, func(_ *testing.T) {
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
				assert.Equal(t, err.Error(), "Password Not Match")
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
		t.Run(tt.name, func(_ *testing.T) {
			tt.beforeFunc(tt.args, tt.fields)
		})
	}
}

func TestUserService_DoGetUsers(t *testing.T) {
	ur := mocks.NewUserRepository(t)

	ur.On("GetUsers").Return([]*domain.User{
		{
			ID:       "1",
			Name:     "any",
			Username: "any",
			Email:    "any",
			Password: "any",
		},
	}, nil)

	us := &UserService{
		userRepo: ur,
		config:   &config.Config{},
	}

	u, err := us.DoGetUsers()
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserService_DoGetUser(t *testing.T) {
	ur := mocks.NewUserRepository(t)

	ur.On("GetUser", "hanan").Return(&domain.User{ID: "1", Name: "any", Username: "any", Email: "any", Password: "any"}, nil)

	us := &UserService{
		userRepo: ur,
		config:   &config.Config{},
	}

	u, err := us.DoGetUser("hanan")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserService_DoDeleteUser(t *testing.T) {
	ur := mocks.NewUserRepository(t)

	ur.On("GetUserById", "1").Return(&domain.User{
		ID:       "1",
		Name:     "any",
		Username: "any",
		Email:    "any",
		Password: "any",
	}, nil)

	ur.On("DeleteUser", "1").Return(nil)

	us := &UserService{
		userRepo: ur,
		config:   &config.Config{},
	}

	err := us.DoDeleteUser("1")
	assert.NoError(t, err)
}

func TestUserService_DoDeleteUser_WithNoUser(t *testing.T) {
	ur := mocks.NewUserRepository(t)

	// Why Delete Method No Mocked ? cus after this function has called the function has return
	// So the delete method not executed
	ur.On("GetUserById", "1").Return(nil, errors.New("User Not Found"))

	us := &UserService{
		userRepo: ur,
		config:   &config.Config{},
	}

	err := us.DoDeleteUser("1")
	assert.Error(t, err)
}

func TestNewUserServie(t *testing.T) {
	ur := mocks.NewUserRepository(t)
	us := &UserService{
		userRepo: ur,
		config:   &config.Config{},
	}
	assert.NotNil(t, us)

	us2 := NewUserServie(ur, us.config)
	assert.NotNil(t, us2)
}
