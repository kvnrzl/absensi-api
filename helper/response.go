package helper

import (
	"net/http"

	"absensi-api.com/model"
	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, model.ResponseWithoutData{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseUnprocessableEntity(c *gin.Context, message string) {
	c.JSON(http.StatusUnprocessableEntity, model.ResponseWithoutData{
		Status:  http.StatusText(http.StatusUnprocessableEntity),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, model.ResponseWithoutData{
		Status:  http.StatusText(http.StatusNotFound),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseInternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, model.ResponseWithoutData{
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, model.ResponseWithoutData{
		Status:  http.StatusText(http.StatusUnauthorized),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.ResponseWithData{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success",
		Data:    data,
	})
}

func ResponseCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, model.ResponseWithData{
		Status:  http.StatusText(http.StatusCreated),
		Message: "Success",
		Data:    data,
	})
}
