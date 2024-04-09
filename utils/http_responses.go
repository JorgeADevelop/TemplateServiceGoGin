package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseData struct {
	Data        interface{}
	Err         error
	Errors      []map[string]interface{}
	Meta        map[string]interface{}
	MessageID   string
	MessageMeta map[string]interface{}
}

func HealtyResponse(ctx *gin.Context) {
	languaje := getLanguaje(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"message": getMessage("HEALTHY", nil, languaje),
		"error":   nil,
		"errors":  nil,
		"data":    gin.H{},
		"meta":    gin.H{},
	})
}

func OkResponse(ctx *gin.Context, responseData ResponseData) {
	languaje := getLanguaje(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"message": getMessage(responseData.MessageID, responseData.MessageMeta, languaje),
		"error":   nil,
		"errors":  nil,
		"data":    responseData.Data,
		"meta":    responseData.Meta,
	})
}

func OkWithErrorsResponse(ctx *gin.Context, responseData ResponseData) {
	languaje := getLanguaje(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success With Errors",
		"message": getMessage(responseData.MessageID, responseData.MessageMeta, languaje),
		"error":   nil,
		"errors":  responseData.Errors,
		"data":    responseData.Data,
		"meta":    responseData.Meta,
	})
}

func BadRequestResponse(ctx *gin.Context, responseData ResponseData) {
	languaje := getLanguaje(ctx)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"status":  "Bad Request",
		"message": getMessage(responseData.MessageID, responseData.MessageMeta, languaje),
		"error":   responseData.Err.Error(),
		"errors":  nil,
		"data":    gin.H{},
		"meta":    gin.H{},
	})
}

func BadRequestResponseWithErrors(ctx *gin.Context, responseData ResponseData) {
	languaje := getLanguaje(ctx)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"status":  "Bad Request",
		"message": getMessage(responseData.MessageID, responseData.MessageMeta, languaje),
		"error":   responseData.Err.Error(),
		"errors":  responseData.Errors,
		"data":    gin.H{},
		"meta":    gin.H{},
	})
}

func InternalServerErrorResponse(ctx *gin.Context, responseData ResponseData) {
	languaje := getLanguaje(ctx)

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"status":  "Internal Server Error",
		"message": getMessage(responseData.MessageID, responseData.MessageMeta, languaje),
		"error":   responseData.Err.Error(),
		"errors":  nil,
		"data":    gin.H{},
		"meta":    gin.H{},
	})
}

func ValidationErrorResponse(ctx *gin.Context, responseData ResponseData) {
	languaje := getLanguaje(ctx)

	var errors []map[string]interface{}
	for _, err := range responseData.Err.(validator.ValidationErrors) {
		errors = append(errors, map[string]interface{}{
			"field": err.Field(),
			"error": err.Tag(),
		})
	}

	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"code":    http.StatusUnprocessableEntity,
		"status":  "Validation Error",
		"message": getMessage(responseData.MessageID, responseData.MessageMeta, languaje),
		"error":   responseData.Err.Error(),
		"errors":  errors,
		"data":    gin.H{},
		"meta":    responseData.Meta,
	})
}

func UnauthorizedResponse(ctx *gin.Context, responseData ResponseData) {
	languaje := getLanguaje(ctx)

	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"status":  "Unauthorized",
		"message": getMessage(responseData.MessageID, responseData.MessageMeta, languaje),
		"error":   responseData.Err.Error(),
		"errors":  nil,
		"data":    gin.H{},
		"meta":    responseData.Meta,
	})
}
