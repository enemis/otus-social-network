package router

import (
	"otus-social-network/internal/server/handler"
	"otus-social-network/internal/server/middleware"
	"otus-social-network/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler     *handler.Handler
	engine      *gin.Engine
	authService service.AuthService
	userService service.UserService
}

func NewRouter(handler *handler.Handler, engine *gin.Engine, authService service.AuthService, userService service.UserService) *Router {
	router := Router{handler: handler, engine: engine, authService: authService, userService: userService}
	router.initRoutes()
	return &router
}

func (r *Router) initRoutes() {
	unauthorised := r.engine.Group("/")
	authorized := r.engine.Group("/", middleware.AuthRequired(r.authService, r.userService))
	r.initAuthRoutes(unauthorised)
	r.initUserRoutes(authorized)
}
