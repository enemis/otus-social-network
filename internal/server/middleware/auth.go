package middleware

import (
	"fmt"
	"otus-social-network/internal/server/response"
	"otus-social-network/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	AuthorizationHeadder = "Authorization"
	UserContext          = "User"
)

func AuthRequired(authService service.AuthService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(AuthorizationHeadder)

		if header == "" {
			response.Unauthorised(c, fmt.Sprintf("Empty %s header", AuthorizationHeadder), "header empty")
			return
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 {
			response.Unauthorised(c, fmt.Sprintf("Invalid %s header", AuthorizationHeadder), "headers part more than 2")
			return
		}

		userId, err := authService.ParseToken(headerParts[1])

		if err != nil {
			logrus.Debugln("rer")
			response.Unauthorised(c, "Invalid token", err.OriginalError())
			return
		}

		user, err := userService.GetUserById(userId)

		if err != nil {
			response.Unauthorised(c, "Invalid token", err.OriginalError())
			return
		}
		c.Set(UserContext, user)
		c.Next()
	}
}
