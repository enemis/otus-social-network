package router

import "github.com/gin-gonic/gin"

func (router *Router) initFriendRoutes(authorizedGroup *gin.RouterGroup) {
	user := authorizedGroup.Group("/friends")
	{
		user.PUT("/", router.handler.AddFriend)
		user.DELETE("/", router.handler.RemoveFriend)
	}
}
