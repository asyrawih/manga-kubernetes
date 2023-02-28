package chapter

import (
	"github.com/asyrawih/manga/config"
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

}
