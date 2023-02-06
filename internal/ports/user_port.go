package ports

import "github.com/asyrawih/manga/internal/core/domain"

type UserRepository interface {
	Login(username string) (*domain.UserLogin, error)
	// Get On User
	GetUserById(id string) (*domain.User, error)
	// Get On User
	GetUser(username string) (*domain.User, error)
	// Get All User
	GetUsers() ([]*domain.User, error)
	// Create User
	// INSERT INTO users (Id, username, name, email, password) VALUES(?, ?, ?, ?, ?);
	CreateUser(in *domain.CreateUser) error
	// Delete An User
	// DELETE FROM users WHERE Id= ?;
	DeleteUser(id string) error
}

type UserService interface {
	DoCreateUser(in *domain.CreateUser) error
	DoGetUser(username string) (*domain.User, error)
	DoGetUsers() ([]*domain.User, error)
	DoDeleteUser(id string) error
	DoLogin(username string, password string) (*domain.UserLoginResponse, error)
}
