package domain

import (
	"github.com/go-playground/validator"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Users []User

type CreateUser struct {
	Name     string `json:"name,omitempty"     validate:"required"`
	Username string `json:"username,omitempty" validate:"required"`
	Email    string `json:"email,omitempty"    validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UserValidator struct {
	Validate validator.Validate
}

type UserLogin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserLoginResponse struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}
