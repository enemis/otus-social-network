package handler

import (
	"fmt"
	"net/http"
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/model"
	"otus-social-network/internal/server/response"
	"otus-social-network/internal/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) AddFriend(c *gin.Context) {
	friend, err := h.resolveFriendFromRequest(c)
	if err != nil {
		return
	}

	user, err := h.getUserFromContext(c)

	if err != nil {
		return
	}

	httpError := h.FriendService.AddFriend(user, friend)
	if httpError != nil {
		response.HttpErrorResponse(c, httpError)
		return
	}

	response.OkWithMessage(c, fmt.Sprintf("User %s %s now is in your friends list", friend.Surname, friend.Name))
}

func (h *Handler) RemoveFriend(c *gin.Context) {
	friend, err := h.resolveFriendFromRequest(c)
	if err != nil {
		return
	}

	user, err := h.getUserFromContext(c)

	if err != nil {
		return
	}

	err = h.FriendService.RemoveFriend(user, friend)

	response.OkWithMessage(c, fmt.Sprintf("User %s %s has been removed from your friends list", friend.Surname, friend.Name))
}

func (h *Handler) resolveFriendFromRequest(c *gin.Context) (*model.User, error) {
	var input dto.FriendId
	validator := validator.BuildValidator(input)

	if err := c.ShouldBind(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validator.DecryptErrors(err).(map[string]interface{}))
		return nil, err
	}

	userUUID, err := uuid.Parse(input.Id)
	if err != nil {
		response.HttpErrorResponse(c, app_error.NewHttpError(err, "Error user id", "id", http.StatusBadRequest))
	}

	user, httpError := h.UserService.GetUserById(userUUID)

	if httpError != nil {
		response.HttpErrorResponse(c, httpError)
		return nil, err
	}

	return user, nil
}
