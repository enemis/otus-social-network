package router

import "github.com/gin-gonic/gin"

func (router *Router) initAuthRoutes(authorizedGroup *gin.RouterGroup) {
	auth := authorizedGroup.Group("/auth")
	{
		auth.POST("/sign-in", router.handler.SignIn)
		auth.POST("/sign-up", router.handler.SignUp)
	}
}
