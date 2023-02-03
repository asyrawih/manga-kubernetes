package services

import "github.com/asyrawih/manga/internal/core/domain"

type UserService struct {
}

func NewUserServie() *UserService {
	return &UserService{}
}

// Get On User
func (us *UserService) DoGetUser(id string) (*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

// Get All User
func (us *UserService) DoGetUsers() ([]*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

// Create An User
func (us *UserService) DoCreateUser(in *domain.CreateUser) (bool, error) {
	panic("not implemented") // TODO: Implement
}

// Delete An User
func (us *UserService) DoDeleteUser(id string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

// Login Using Username And Password
func (us *UserService) Login(username string, password string) (*domain.User, error) {
	panic("not implemented") // TODO: Implement
}
