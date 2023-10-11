package handler

import (
	"net/http"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/server/response"
	"otus-social-network/internal/validator"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input dto.SignUpInput
	validator := validator.BuildValidator(input)

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validator.DecryptErrors(err).(map[string]interface{}))
		return
	}

	uuid, err := h.AuthService.CreateUser(&input)

	if err != nil {
		response.HttpErrorResponse(c, err)
		return
	}

	response.Created(c, uuid)
}

func (h *Handler) SignIn(c *gin.Context) {
	var input dto.SignInInput

	g := galidator.G()
	validator := g.Validator(input)

	if err := c.BindJSON(&input); err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, validator.DecryptErrors(err).(map[string]interface{}))
		return
	}

	token, err := h.AuthService.GenerateToken(input.Email, input.Password)

	if err != nil {
		response.HttpErrorResponse(c, err)
		return
	}

	response.Ok(c, map[string]interface{}{"token": token})
}
