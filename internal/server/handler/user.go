package handler

import (
	"net/http"
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/server/response"
	"otus-social-network/internal/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (h *Handler) UserPage(c *gin.Context) {
	var input dto.UserId
	validator := validator.BuildValidator(input)

	if err := c.ShouldBindUri(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validator.DecryptErrors(err).(map[string]interface{}))
		return
	}

	//https://github.com/gin-gonic/gin/issues/2423
	userUUID, err := uuid.Parse(input.Id)
	if err != nil {
		response.HttpErrorResponse(c, app_error.NewHttpError(err, "Error user id", "id", http.StatusBadRequest))
	}

	user, httpError := h.UserService.GetUserById(userUUID)

	logrus.Debug(user)
	logrus.Debugln(httpError)
	if httpError != nil {
		logrus.Debugln("not nil")
		response.HttpErrorResponse(c, httpError)
		return
	}

	response.Ok(c, user)
}
