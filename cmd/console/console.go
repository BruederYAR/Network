package main

import (
	sqlite_repository "Network/internal/repository/sqlite"
	"Network/internal/service"
	ui "Network/internal/ui/console_ui"
	"Network/pkg/db"
	"Network/pkg/input"
	"Network/pkg/logs"
	"Network/server/node"
	"errors"
	"fmt"
	"os"
)

func main() {
	//Test args 
	input.TestArgs(os.Args)

	//create folder
	err := input.CreateDirectory(os.Args[2])
	if err != nil {
		panic(errors.New("folder create exection:" + err.Error()))
	}

	//create logger
	logger, err := logs.NewLogger("./cache/"+os.Args[2]+"/", os.Args[2])
	if err != nil {
		panic(errors.New("log file create exection" + err.Error()))
	} else {
		logger.LogInfo("Logger created")
	}

	//create config
	config := input.NewConfigByConsole(logger, os.Args)
	logger.LogInfo("Config created")

	//create db
	err = db.CreateDB(config.Connect)
	if err != nil {
		logger.LogError(errors.New("create db exception: " + err.Error()))
	} else {
		logger.LogInfo("Database created")
	}

	repos := sqlite_repository.NewNodeRepository(*config, logger)
	serv := service.NewNodeService(logger, repos)

	//create node
	n, err := node.NewNode(logger, *config, serv)
	if err != nil {
		panic("node dont start")
	} else {
		logger.LogInfo("Node created")
	}

	//serv.Create(dto.Node{Address: "123.123", Name: "Test", PublicKey: "qwer" })
	// res, err := serv.GetAll()
	// for i := 0; i < len(res); i++  {
	// 	fmt.Println(fmt.Sprintf("ID: %s | PublicKey: %s | Address: %s | Name: %s", res[i].NodeId, res[i].PublicKey, res[i].Address, res[i].Name ))
	// }
	//id := res[0].NodeId
	//res2, err := serv.GetById(id)
	//fmt.Println(res2)

	res, err := serv.GetAllIds()
	for i := 0; i < len(res); i++ {
		fmt.Println(fmt.Sprintf("ID: %s ", res[i].Id))
	}

	n.Run(ui.ConsoleClient)
}
