package main

import (
	"Network/pkg/input"
	"Network/internal/repository/sqlite"
	"Network/internal/service"
	ui "Network/internal/ui/console_ui"
	"Network/pkg/db"
	"Network/pkg/logs"
	"Network/server/node"
	"errors"
	"os"
)

func main() {
	//create logger
	logger, err := logs.NewLogger("./cache/", os.Args[2])
	if err != nil {
		panic(errors.New("log file create exection" + err.Error()))
	}else{
		logger.LogInfo("Logger created")
	}
	
	//create config
	config := input.NewConfigByConsole(logger, os.Args)
	logger.LogInfo("Config created")

	//create db
	err = db.CreateDB(config.Connect)
	if err !=nil {
		logger.LogError(errors.New("create db exception: " + err.Error()))
	}else{
		logger.LogInfo("Database created")
	}

	repos := sqlite_repository.NewNodeRepository(*config, logger)
	serv := service.NewNodeService(logger, repos)
	serv.GetById(1)
	
	//create node
	n, err := node.NewNode(logger, *config, serv)
	if err != nil {
		panic("node dont start")
	}else{
		logger.LogInfo("Node created")
	}

	n.Run(ui.ConsoleClient)
}
