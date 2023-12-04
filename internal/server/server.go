package server

import (
	"otus-social-network/internal/config"
	"otus-social-network/internal/server/handler"
	"otus-social-network/internal/server/router"
	"otus-social-network/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config        *config.Config
	authService   service.AuthService
	userService   service.UserService
	friendService service.FriendService
}

func NewServer(config *config.Config, authService service.AuthService, userService service.UserService, friendService service.FriendService) *Server {
	return &Server{config: config, authService: authService, userService: userService, friendService: friendService}
}

func (s *Server) Run() error {
	ginEngine := gin.Default()
	handler := handler.NewHandler(s.authService, s.userService, s.friendService)
	router.NewRouter(handler, ginEngine, s.authService, s.userService)

	return ginEngine.Run()
}
