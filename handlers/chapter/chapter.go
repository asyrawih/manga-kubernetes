package chapter

import (
	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/labstack/echo/v4"
)

type ChapterHttpHandler struct {
	ports.ChapterService
	*config.Config
}

func NewHttpHandler(chapterService ports.ChapterService, config *config.Config) *ChapterHttpHandler {
	return &ChapterHttpHandler{
		ChapterService: chapterService,
		Config:         config,
	}
}

func (h *ChapterHttpHandler) Routes(e *echo.Echo) {
	chapterGroup := e.Group("/v1/api/chapter")
	chapterGroup.GET("/:mangaID", h.GetChapter)

}

func (c *ChapterHttpHandler) GetChapter(e echo.Context) error {
	s := e.Param("mangaID")
	limit := e.QueryParam("limit")
	offset := e.QueryParam("offset")

	args := domain.QueryArgs{
		Limit:   limit,
		Offset:  offset,
		OrderBy: "DESC",
	}

	chapters, err := c.ChapterService.DoGetChapters(s, args)

	if err != nil {
		return e.JSON(500, err.Error())
	}
	return e.JSON(200, chapters)
}
