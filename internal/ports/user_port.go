package ports

import "github.com/asyrawih/manga/internal/core/domain"

type UserRepository interface {
	// Get On User
	GetUser(id string) (*domain.User, error)
	// Get All User
	GetUsers() ([]*domain.User, error)
	// Create An User
	CreateUser(in *domain.CreateUser) (bool, error)
	// Delete An User
	DeleteUser(id string) (bool, error)
}

type UserService interface {
	// Get On User
	DoGetUser(id string) (*domain.User, error)
	// Get All User
	DoGetUsers() ([]*domain.User, error)
	// Create An User
	DoCreateUser(in *domain.CreateUser) (bool, error)
	// Delete An User
	DoDeleteUser(id string) (bool, error)

	// Login Using Username And Password
	Login(username, password string) (*domain.User, error)
}
