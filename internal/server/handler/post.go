package handler

import (
	"net/http"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/server/response"
	"otus-social-network/internal/validator"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPost(c *gin.Context) {

}

func (h *Handler) CreatePost(c *gin.Context) {
	var input dto.PostInput
	validator := validator.BuildValidator(input)

	if err := c.ShouldBind(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validator.DecryptErrors(err).(map[string]interface{}))
		return
	}

	user, httpError := h.getUserFromContext(c)

	if httpError != nil {
		return
	}

	h.PostService.CreatePost(user, input)
}

func (h *Handler) UpdatePost(c *gin.Context) {

}

func (h *Handler) DeletePost(c *gin.Context) {

}
