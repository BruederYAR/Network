package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) sendMessage(context *gin.Context) {
	type connect struct {
		Address string `json:"address"`
		Message string `json:"message"`
	}

	var cn connect
	if err := context.BindJSON(&cn); err != nil {
		newErrorResponse(handler.Node.Logger, context, http.StatusBadRequest, err.Error())
	}

	err := handler.Node.SendMessageTo(cn.Address, []byte(cn.Message))
	if err != nil {
		newErrorResponse(handler.Node.Logger, context, http.StatusBadRequest, err.Error())
		return
	}
}

func (handler *Handler) connectTo(context *gin.Context) {
	address := context.Param("address")

	err := handler.Node.HandShakeIds(address, true)

	if err != nil {
		newErrorResponse(handler.Node.Logger, context, http.StatusBadRequest, "invalid address param")
		return
	}

	context.JSON(http.StatusOK, "connect to "+address)
}
