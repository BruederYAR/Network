package handler

import (
	"Network/server/node"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Node node.Node
}

func NewHandler(node node.Node) *Handler {
	return &Handler{Node: node}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api", CORSMiddleware())
	{
		nodes := api.Group("/nodes")
		{
			nodes.GET("/", handler.getAllNodes)
		}
		control := api.Group("/control")
		{
			control.POST("/send", handler.sendMessage)
			control.GET("/connect/:address", handler.connectTo)
		}
	}

	return router
}
