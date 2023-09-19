package manga

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/handlers/middleware"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/asyrawih/manga/pkg/validation"
)

type MangaHTTPHandler struct {
	mangaService ports.MangaService
	config       *config.Config
}

var (
	ValidStatus = []string{"Publish", "Draft"}
	ValidGenre  = []string{"Manga", "Manhwa", "Manhua"}
)

func NewHTTPHandler(mangaService ports.MangaService, config *config.Config) *MangaHTTPHandler {
	return &MangaHTTPHandler{
		mangaService: mangaService,
		config:       config,
	}
}

func (h *MangaHTTPHandler) Routes(e *echo.Echo) {
	userMiddleware := middleware.AuthMiddleware(h.config.Key)
	mangaGroup := e.Group("/v1/api/manga")
	mangaGroup.POST("/", h.Create, userMiddleware)
	mangaGroup.PUT("/", h.Update, userMiddleware)
	mangaGroup.DELETE("/:mangaID", h.Delete, userMiddleware)
	mangaGroup.GET("/:mangaID", h.GetByID)
	mangaGroup.GET("/all", h.GetAll)
	mangaGroup.GET("/author/:author_name", h.GetByAuthor)
	mangaGroup.GET("/search", h.Search)
}

// Get All Manga
func (ma *MangaHTTPHandler) GetAll(e echo.Context) error {
	m, err := ma.mangaService.DoGetAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, m)
}

// Create Manga
func (ma *MangaHTTPHandler) Create(e echo.Context) error {
	var mangaRequest *domain.CreateRequest

	if err := e.Bind(&mangaRequest); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := e.Validate(mangaRequest); err != nil {
		vr := validation.ValidateMessage(err)
		return e.JSON(http.StatusBadRequest, vr)
	}

	checkStatus := func(a []string, b string) bool {
		for _, val := range a {
			if val == b {
				return true
			}
		}
		return false
	}

	if b := checkStatus(ValidStatus, mangaRequest.Status); !b {
		return e.JSON(http.StatusBadRequest, "Status Must Publish, or Draft")
	}

	if b := checkStatus(ValidGenre, string(mangaRequest.Genre)); !b {
		return e.JSON(http.StatusBadRequest, fmt.Sprintf("Must Valid Genre %+v", ValidGenre))
	}

	if err := ma.mangaService.DoCreate(mangaRequest); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, mangaRequest)
}

// Update The Manga
func (ma *MangaHTTPHandler) Update(e echo.Context) error {
	return e.JSON(http.StatusOK, "oke")
}

// Get manga By Id
func (ma *MangaHTTPHandler) GetByID(e echo.Context) error {
	s := e.Param("mangaID")
	m, err := ma.mangaService.DoGetByID(s)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, m)
}

// Get Manga By Author
func (ma *MangaHTTPHandler) GetByAuthor(e echo.Context) error {
	s := e.Param("author")
	mangas, err := ma.mangaService.DoGetByAuthor(s)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, mangas)
}

// Search Manga by limit them 100 page
func (ma *MangaHTTPHandler) Search(e echo.Context) error {
	s := e.QueryParam("title")
	mangas, err := ma.mangaService.DoSearch(s)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, mangas)
}

// Delete The Manga
func (ma *MangaHTTPHandler) Delete(e echo.Context) error {
	s := e.Param("mangaID")
	if err := ma.mangaService.DoDelete(s); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusNoContent, "Success")
}
