package response

import (
	"net/http"
	"otus-social-network/internal/app_error"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Created(c *gin.Context, id interface{}) {
	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
}

func Ok(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func Unauthorised(c *gin.Context, message string, originalError string) {
	logrus.Error(originalError)
	ErrorResponse(c, http.StatusUnauthorized, map[string]interface{}{"error": message})
}

func ErrorResponse(c *gin.Context, statusCode int, errors map[string]interface{}) {
	logrus.Error(errors)
	c.AbortWithStatusJSON(statusCode, map[string]interface{}{"errors": errors})
}

func HttpErrorResponse(c *gin.Context, httpError *app_error.HttpError) {
	errors := map[string]interface{}{httpError.Field(): httpError.Error()}
	logrus.Error(httpError.OriginalError())
	ErrorResponse(c, httpError.Status(), errors)
}

func InternalServerError(c *gin.Context, err error) {
	logrus.Error(err)
	httpError := app_error.NewInternalServerError(err)
	errors := map[string]interface{}{httpError.Field(): httpError.Error()}
	c.AbortWithStatusJSON(httpError.Status(), map[string]interface{}{"errors": errors})
}
