package handlers

import "github.com/labstack/echo/v4"

type HttpHandler struct {
}

// NewHttpHandler function  
//
// Create User By returning User Instance
func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) CreateUser(e echo.Context) error {
	return nil
}
