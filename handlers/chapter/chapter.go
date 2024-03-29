package chapter

import (
	"github.com/labstack/echo/v4"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
)

type ChapterHTTPHandler struct {
	ports.ChapterService
	*config.Config
}

func NewHTTPHandler(chapterService ports.ChapterService, config *config.Config) *ChapterHTTPHandler {
	return &ChapterHTTPHandler{
		ChapterService: chapterService,
		Config:         config,
	}
}

func (h *ChapterHTTPHandler) Routes(e *echo.Echo) {
	chapterGroup := e.Group("/v1/api/chapter")
	chapterGroup.GET("/:mangaID", h.GetChapter)
}

func (c *ChapterHTTPHandler) GetChapter(e echo.Context) error {
	s := e.Param("mangaID")
	limit := e.QueryParam("limit")
	offset := e.QueryParam("page")

	args := domain.QueryArgs{
		Limit:   limit,
		Offset:  offset,
		OrderBy: "DESC",
	}

	chapters, err := c.DoGetChapters(s, args)
	if err != nil {
		return e.JSON(500, err.Error())
	}
	return e.JSON(200, chapters)
}
