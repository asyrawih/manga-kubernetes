package repositories

import "github.com/asyrawih/manga/internal/core/domain"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// Get On User
func (us *UserRepo) GetUser(id string) (*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

// Get All User
func (us *UserRepo) GetUsers() ([]*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

// Create An User
func (us *UserRepo) CreateUser(in *domain.CreateUser) (bool, error) {
	panic("not implemented") // TODO: Implement
}

// Delete An User
func (us *UserRepo) DeleteUser(id string) (bool, error) {
	panic("not implemented") // TODO: Implement
}
