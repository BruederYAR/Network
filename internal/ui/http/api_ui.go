package ui

import (
	"Network/internal/handler"
	http_server "Network/server/http"
	"Network/server/node"
)

func ApiServer(n *node.Node) {
	handler := handler.NewHandler(*n)

	server := new(http_server.Server)
	if err := server.Run("7000", handler.InitRoutes()); err != nil {
		n.Logger.LogPanic("error occuped while running http server: " + err.Error())
	}
}
