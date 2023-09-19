package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/handlers/chapter"
	"github.com/asyrawih/manga/handlers/manga"
	"github.com/asyrawih/manga/handlers/users"
	"github.com/asyrawih/manga/internal/services"
	repositories "github.com/asyrawih/manga/repositories/mysql"
)

type HTTPService struct {
	echo   *echo.Echo
	config *config.Config
	db     *sql.DB
}

func NewHTTPService(echo *echo.Echo, config *config.Config, db *sql.DB) *HTTPService {
	return &HTTPService{
		echo:   echo,
		config: config,
		db:     db,
	}
}

func (h *HTTPService) Run(port string) error {
	h.echo.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "oke")
	})

	// User
	ur := repositories.NewUserRepo(h.db)
	us := services.NewUserServie(ur, h.config)
	userHandler := users.NewHTTPHandler(us, h.config)
	userHandler.Routes(h.echo)

	// Manga
	mangaRepo := repositories.NewMangaRepo(h.db)
	mangaService := services.NewMangaService(mangaRepo, h.config)
	mangaHandler := manga.NewHTTPHandler(mangaService, h.config)
	mangaHandler.Routes(h.echo)

	// Chapter
	chapterRepo := repositories.NewChapterRepo(h.db)
	chapterService := services.NewChapterService(chapterRepo)
	chapterHandler := chapter.NewHTTPHandler(chapterService, h.config)
	chapterHandler.Routes(h.echo)

	return h.echo.Start(port)
}
