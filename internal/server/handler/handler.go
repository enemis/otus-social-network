package handler

import (
	"otus-social-network/internal/service"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type Handler struct {
	AuthService service.AuthService
	UserService service.UserService
}

func NewHandler(authService service.AuthService, userService service.UserService) *Handler {
	return &Handler{AuthService: authService, UserService: userService}
}
