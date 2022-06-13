package input

import (
	"Network/pkg/logs"
	"strings"
)

const pathToConnect = "./cache/"

type Congig struct {
	Name      string
	Connect   string
	CachePath string
	Address   *Address
}

type Address struct {
	Ip   string
	Ipv4 string
	Ipv6 string
	Port string
}

func TestArgs(args []string) {
	if len(args[1]) == 0 {
		panic("please write your address")
	}
	splited := strings.Split(args[1], ":")
	if len(splited) != 2 {
		panic("wrong address")
	}

	if len(args[2]) == 0 {
		panic("please write your name")
	}
}

func NewConfigByConsole(logger logs.ILogger, args []string) *Congig {
	splited := strings.Split(args[1], ":")

	config := &Congig{
		Name:    args[2],
		CachePath: pathToConnect + args[2] + "/",
	}
	config.Connect = config.CachePath + args[2] + ".db"

	ipv4, ipv6 := LocalIpAddress(logger)
	config.Address = &Address{
		Ip:   splited[0],
		Ipv4: ipv4,
		Ipv6: ipv6,
		Port: splited[1],
	}

	return config
}
