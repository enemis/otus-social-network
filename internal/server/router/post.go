package router

import "github.com/gin-gonic/gin"

func (router *Router) initPostRoutes(authorizedGroup *gin.RouterGroup) {
	user := authorizedGroup.Group("/posts")
	{
		user.GET("/:id", router.handler.AddFriend)
		user.POST("/", router.handler.CreatePost)
		user.PUT("/", router.handler.UpdatePost)
		user.DELETE("/:id", router.handler.DeletePost)
	}
}
