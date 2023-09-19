package mysql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/pkg/password"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// Get On User
func (us *UserRepo) GetUserByID(id string) (*domain.User, error) {
	ctx := context.Background()

	var user domain.User
	const query = `SELECT id, username, email, name FROM users u where u.id= ?;`
	r := us.db.QueryRowContext(ctx, query, id)
	if err := r.Scan(&user.ID, &user.Username, &user.Email, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}

// Get On User

const getUser = `SELECT id, username, email, name FROM users u where u.username = ?;`

func (us *UserRepo) GetUser(username string) (*domain.User, error) {
	ctx := context.Background()

	var user domain.User

	r := us.db.QueryRowContext(ctx, getUser, username)
	if err := r.Scan(&user.ID, &user.Username, &user.Email, &user.Name); err != nil {
		return nil, err
	}

	return &user, nil
}

const getUsers = "SELECT id, username, email, name FROM users LIMIT 100;"

// Get All User
func (us *UserRepo) GetUsers() ([]*domain.User, error) {
	ctx := context.Background()

	var users []*domain.User

	r, err := us.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}

	defer r.Close()

	for r.Next() {
		var user domain.User
		if err := r.Scan(&user.ID, &user.Username, &user.Email, &user.Name); err != nil {
			log.Err(err).Caller().Msg("")
		}
		users = append(users, &user)
	}

	return users, nil
}

// Create An User
func (us *UserRepo) CreateUser(in *domain.CreateUser) error {
	ctx := context.Background()

	const query = "INSERT INTO users (Id, username, name, email, password) VALUES(?, ?, ?, ?, ?);"
	Id := uuid.New().String()

	s, err := password.HashPassword(in.Password)
	if err != nil {
		return err
	}

	result, err := us.db.ExecContext(ctx, query, Id, in.Username, in.Name, in.Email, s)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Info().Msgf("Row Affected %d", affected)

	return nil
}

func (us *UserRepo) DeleteUser(id string) error {
	const query = "DELETE FROM users WHERE Id= ?;"
	r, err := us.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

const loginQuery = "SELECT username , password from users where users.username =?"

func (us *UserRepo) Login(username string) (*domain.UserLogin, error) {
	ctx := context.Background()

	var userLogin domain.UserLogin

	r := us.db.QueryRowContext(ctx, loginQuery, username)

	if err := r.Scan(&userLogin.Username, &userLogin.Password); err != nil {
		return nil, err
	}

	return &userLogin, nil
}
