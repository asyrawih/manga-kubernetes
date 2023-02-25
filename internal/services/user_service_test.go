package services

import (
	"errors"
	"testing"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/mocks"
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
