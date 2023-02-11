package users

import (
	"net/http"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/handlers/middleware"
	"github.com/asyrawih/manga/internal/core/domain"
	"github.com/asyrawih/manga/internal/ports"
	"github.com/asyrawih/manga/pkg/validation"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type HttpHandler struct {
	userService ports.UserService
	config      *config.Config
}

// NewHttpHandler function  
//
// Create User By returning User Instance
func NewHttpHandler(userServie ports.UserService, config *config.Config) *HttpHandler {
	return &HttpHandler{
		userService: userServie,
		config:      config,
	}
}

type ValidatinResponse struct {
	Field      string
	Validation string
}

func (h *HttpHandler) Routes(e *echo.Echo) {
	userMiddleware := middleware.AuthMiddleware(h.config.Key)

	userGroup := e.Group("/v1/api/user")
	userGroup.POST("/", h.CreateUser)
	userGroup.POST("/login", h.Login)

	// Restrict Route
	userGroup.GET("/", h.GetUsers, userMiddleware)
	userGroup.GET("/:username", h.GetUser, userMiddleware)
	userGroup.DELETE("/:id", h.DeleteUser, userMiddleware)
}

func (h *HttpHandler) CreateUser(e echo.Context) error {
	var useRequest *domain.CreateUser

	if err := e.Bind(&useRequest); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := e.Validate(useRequest); err != nil {
		vr := validation.ValidateMessage(err)
		return e.JSON(http.StatusBadRequest, vr)
	}

	if err := h.userService.DoCreateUser(useRequest); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, "success create user")
}

func (h *HttpHandler) GetUser(e echo.Context) error {
	username := e.Param("username")
	u, err := h.userService.DoGetUser(username)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, u)
}

func (h *HttpHandler) GetUsers(e echo.Context) error {
	u, err := h.userService.DoGetUsers()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, u)
}

func (h *HttpHandler) DeleteUser(e echo.Context) error {
	id := e.Param("id")
	err := h.userService.DoDeleteUser(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusNoContent, "success delete")
}

func (h *HttpHandler) Login(e echo.Context) error {
	var userLogin domain.UserLogin

	if err := e.Bind(&userLogin); err != nil {
		log.Err(err).Caller().Msg("")
	}

	ulr, err := h.userService.DoLogin(userLogin.Username, userLogin.Password)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, err.Error())
	}

	return e.JSON(http.StatusOK, ulr)
}
