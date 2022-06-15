package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) getAllNodes(context *gin.Context) {
	data, err := handler.Node.Service.GetAll()

	if err != nil {
		newErrorResponse(handler.Node.Logger, context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, data)
}
