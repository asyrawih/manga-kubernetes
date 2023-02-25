package services

import (
	"errors"
	"time"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/asyrawih/manga/pkg/password"
	"github.com/o1egl/paseto"
)

type UserService struct {
	userRepo ports.UserRepository
	config   *config.Config
}

func NewUserServie(userRepo ports.UserRepository, config *config.Config) *UserService {
	return &UserService{
		userRepo: userRepo,
		config:   config,
	}
}

func (us *UserService) DoCreateUser(in *domain.CreateUser) error {
	return us.userRepo.CreateUser(in)
}

func (us *UserService) DoGetUsers() ([]*domain.User, error) {
	return us.userRepo.GetUsers()
}

func (us *UserService) DoGetUser(username string) (*domain.User, error) {
	return us.userRepo.GetUser(username)
}

// DoDeleteUser method  
func (us *UserService) DoDeleteUser(id string) error {
	u, err := us.userRepo.GetUserById(id)
	if err != nil {
		return err
	}
	return us.userRepo.DeleteUser(u.Id)
}

func (us *UserService) DoLogin(username string, pass string) (*domain.UserLoginResponse, error) {
	ul, err := us.userRepo.Login(username)
	if err != nil {
		return nil, err
	}

	if match := password.CheckPasswordHash(pass, ul.Password); !match {
		return nil, errors.New("Password Not Match")
	}

	symmetricKey := []byte(us.config.Key)
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   ul.Username,
		Issuer:     "manga",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	// Add custom claim    to the token
	jsonToken.Set("username", ul.Username)
	// Encrypt data
	token, _ := paseto.NewV2().Encrypt(symmetricKey, jsonToken, nil)

	return &domain.UserLoginResponse{Message: "Success Login", Token: token}, nil
}
