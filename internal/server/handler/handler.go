package handler

import (
	"errors"
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/model"
	"otus-social-network/internal/server/middleware"
	"otus-social-network/internal/service"

	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type Handler struct {
	AuthService   service.AuthService
	UserService   service.UserService
	FriendService service.FriendService
	PostService   service.PostService
}

func NewHandler(
	authService service.AuthService,
	userService service.UserService,
	friendService service.FriendService,
	postService service.PostService,
) *Handler {
	return &Handler{AuthService: authService, UserService: userService, FriendService: friendService, PostService: postService}
}

func (h *Handler) getUserFromContext(c *gin.Context) (*model.User, error) {
	var err error
	usr, exists := c.Get(middleware.UserContext)

	if !exists {
		err = app_error.NewInternalServerError(errors.New("User must be present in context"))
	}

	user, ok := usr.(*model.User)
	if !ok {
		err = app_error.NewInternalServerError(errors.New("Context must contain model.User"))
	}

	return user, err
}
