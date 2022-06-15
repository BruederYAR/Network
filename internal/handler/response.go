package handler

import (
	"Network/pkg/logs"
	"errors"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(logger logs.ILogger, context *gin.Context, statusCode int, message string) {
	logger.LogError(errors.New(message))
	context.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}