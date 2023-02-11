package manga

import (
	"fmt"
	"net/http"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/handlers/middleware"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/asyrawih/manga/pkg/validation"
	echo "github.com/labstack/echo/v4"
)

type MangaHttpHandler struct {
	mangaService ports.MangaService
	config       *config.Config
}

var (
	ValidStatus = []string{"Publish", "Draft"}
	ValidGenre  = []string{"Manga", "Manhwa", "Manhua"}
)

func NewHttpHandler(mangaService ports.MangaService, config *config.Config) *MangaHttpHandler {
	return &MangaHttpHandler{
		mangaService: mangaService,
		config:       config,
	}
}

func (h *MangaHttpHandler) Routes(e *echo.Echo) {
	userMiddleware := middleware.AuthMiddleware(h.config.Key)
	mangaGroup := e.Group("/v1/api/manga")
	mangaGroup.POST("/", h.Create, userMiddleware)
	mangaGroup.PUT("/", h.Update, userMiddleware)
	mangaGroup.DELETE("/:mangaID", h.Delete, userMiddleware)
	mangaGroup.GET("/:mangaID", h.GetById)
	mangaGroup.GET("/all", h.GetAll)
	mangaGroup.GET("/author/:author_name", h.GetByAuthor)
	mangaGroup.GET("/search", h.Search)
}

// Get All Manga
func (ma *MangaHttpHandler) GetAll(e echo.Context) error {
	m, err := ma.mangaService.DoGetAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, m)
}

// Create Manga
func (ma *MangaHttpHandler) Create(e echo.Context) error {
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
func (ma *MangaHttpHandler) Update(e echo.Context) error {
	return e.JSON(http.StatusOK, "oke")
}

// Get manga By Id
func (ma *MangaHttpHandler) GetById(e echo.Context) error {
	s := e.Param("mangaID")
	m, err := ma.mangaService.DoGetByID(s)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, m)
}

// Get Manga By Author
func (ma *MangaHttpHandler) GetByAuthor(e echo.Context) error {
	return e.JSON(http.StatusOK, "oke")
}

// Search Manga by limit them 100 page
func (ma *MangaHttpHandler) Search(e echo.Context) error {
	return e.JSON(http.StatusOK, "oke")
}

// Delete The Manga
func (ma *MangaHttpHandler) Delete(e echo.Context) error {
	s := e.Param("mangaID")
	if err := ma.mangaService.DoDelete(s); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusNoContent, "Success")
}
