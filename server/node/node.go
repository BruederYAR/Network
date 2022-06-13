package node

import (
	"Network/internal/entites/dto"
	"Network/internal/service"
	"Network/pkg/input"
	"Network/pkg/logs"
	"Network/pkg/protocol"
	"crypto/rand"
	"crypto/rsa"
	"os"
)

type Address struct {
	IP   string
	Port string
}

type Node struct {
	Titles     map[int]string
	Types      map[int]string
	Address    Address
	Name       string         //Имя узла
	PrivateKey rsa.PrivateKey //Приватный ключ для rsa
	PublicKey  rsa.PublicKey  //Публичный ключ для rsa
	Logger     logs.ILogger
	Config     input.Congig
	Service    service.INodeService
}

func NewNode(logger logs.ILogger, cfg input.Congig, service service.INodeService) (*Node, error) {

	//PrivateKey
	privateKeyPath := cfg.CachePath + "privatekey"

	var PrivateKey *rsa.PrivateKey
	if !input.FileExist(privateKeyPath) {
		key, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			logger.LogPanic("private key generate exeption: " + err.Error())
		}
		input.SavePrivateKeyToFile(privateKeyPath, *key)
		PrivateKey = key
	} else {
		key, err := input.LoadPrivateKeyOnFile(privateKeyPath)
		if err != nil {
			logger.LogPanic("private key read exeption: " + err.Error())
		}
		PrivateKey = key
	}

	newnode := &Node{
		Titles:     map[int]string{-1: "handshakeids", 0: "handshake", 1: "date", 2: "modcmd", 3: "cmd"},
		Types:      map[int]string{0: "string", 1: "json"},
		Name:       os.Args[2],
		PrivateKey: *PrivateKey,
		PublicKey:  PrivateKey.PublicKey,
		Logger:     logger,
		Config:     cfg,
		Service:    service,
	}

	switch cfg.Address.Ip {
	case "":
		newnode.Address = Address{IP: cfg.Address.Ipv4, Port: ":" + cfg.Address.Port}
	case "ipv4":
		newnode.Address = Address{IP: cfg.Address.Ipv4, Port: ":" + cfg.Address.Port}
	case "ipv6":
		newnode.Address = Address{IP: cfg.Address.Ipv6, Port: ":" + cfg.Address.Port}
	case "null":
		newnode.Address = Address{IP: "", Port: ":" + cfg.Address.Port}
	default:
		newnode.Address = Address{IP: cfg.Address.Ip, Port: ":" + cfg.Address.Port}
	}

	//Adding information about yourself to the database
	key, err := protocol.PublicKeyToBytes(newnode.PublicKey)
	if err != nil {
		logger.LogPanic("It is impossible to pack a public key")
	}

	localNode := dto.Node{Address: newnode.Address.IP + newnode.Address.Port, Name: newnode.Name, PublicKey: key}
	flag, _ := service.IfExist(localNode)
	if !flag {
		service.Create(localNode)
	}

	return newnode, nil
}

func (node *Node) Run(handleClient func(*Node)) { //Выполняется запуск как сервер, так и клиента
	go handlerServer(node)
	handleClient(node)
}
