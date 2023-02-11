package handler

import (
	"database/sql"
	"net/http"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/handlers/users"
	"github.com/asyrawih/manga/internal/services"
	"github.com/asyrawih/manga/repositories"
	"github.com/labstack/echo/v4"
)

type HttpService struct {
	echo   *echo.Echo
	config *config.Config
	db     *sql.DB
}

func NewHttpService(echo *echo.Echo, config *config.Config, db *sql.DB) *HttpService {
	return &HttpService{
		echo:   echo,
		config: config,
		db:     db,
	}
}

func (h *HttpService) Run(port string) error {
	h.echo.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "oke")
	})

	ur := repositories.NewUserRepo(h.db)
	us := services.NewUserServie(ur, h.config)

	// Users
	userHandler := users.NewHttpHandler(us, h.config)
	userHandler.Routes(h.echo)

	return h.echo.Start(port)
}
