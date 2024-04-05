package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ResponseData struct {
	data      interface{}
	err       error
	errors    []map[string]interface{}
	messageID string
}

func HealtyResponse(ctx *gin.Context) {
	message := Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "HEALTY",
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"message": message,
		"error":   nil,
		"errors":  nil,
		"data":    gin.H{},
	})
}

func OkResponse(ctx *gin.Context, responseData ResponseData) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"message": responseData.messageID,
		"error":   nil,
		"errors":  nil,
		"data":    responseData.data,
	})
}

func OkWithErrorsResponse(ctx *gin.Context, responseData ResponseData) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success With Errors",
		"message": responseData.messageID,
		"error":   nil,
		"errors":  responseData.errors,
		"data":    responseData.data,
	})
}

func BadRequestResponse(ctx *gin.Context, responseData ResponseData) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"status":  "Bad Request",
		"message": responseData.messageID,
		"error":   responseData.err.Error(),
		"errors":  nil,
		"data":    gin.H{},
	})
}

func InternalServerErrorResponse(ctx *gin.Context, responseData ResponseData) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"status":  "Internal Server Error",
		"message": responseData.messageID,
		"error":   responseData.err.Error(),
		"errors":  nil,
		"data":    gin.H{},
	})
}

func ValidationErrorResponse(ctx *gin.Context, responseData ResponseData) {
	var errors []map[string]interface{}
	for _, err := range responseData.err.(validator.ValidationErrors) {
		errors = append(errors, map[string]interface{}{
			"field": err.Field(),
			"error": err.Tag(),
		})
	}

	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"code":    http.StatusUnprocessableEntity,
		"status":  "Validation Error",
		"message": responseData.messageID,
		"error":   responseData.err.Error(),
		"errors":  errors,
		"data":    gin.H{},
	})
}
