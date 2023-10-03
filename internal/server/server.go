package server

import (
	"otus-social-network/internal/config"
	"otus-social-network/internal/server/handler"
	"otus-social-network/internal/server/router"
	"otus-social-network/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config      *config.Config
	authService service.AuthService
	userService service.UserService
}

func NewServer(config *config.Config, authService service.AuthService, userService service.UserService) *Server {
	return &Server{config: config, authService: authService, userService: userService}
}

func (s *Server) Run() error {
	ginEngine := gin.Default()
	handler := handler.NewHandler(s.authService, s.userService)
	router.NewRouter(handler, ginEngine, s.authService, s.userService)

	return ginEngine.Run()
}
