package router

import "github.com/gin-gonic/gin"

func (router *Router) initUserRoutes(authorizedGroup *gin.RouterGroup) {
	user := authorizedGroup.Group("/users")
	{
		user.GET("/search", router.handler.FindUsers)
		user.GET(":id", router.handler.UserPage)
	}
}
