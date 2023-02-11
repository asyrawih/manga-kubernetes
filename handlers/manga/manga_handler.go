package manga

import (
	"fmt"
	"net/http"

	"github.com/asyrawih/manga/config"
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
	// userMiddleware := middleware.AuthMiddleware(h.config.Key)
	mangaGroup := e.Group("/v1/api/manga")
	mangaGroup.POST("/", h.CreateManga)
}

func (h *MangaHttpHandler) CreateManga(e echo.Context) error {
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

	if err := h.mangaService.DoCreate(mangaRequest); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, mangaRequest)
}
