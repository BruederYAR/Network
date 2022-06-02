package input

import (
	"Network/pkg/logs"
	"errors"
	"net"
	"runtime"
	"strings"
)

func LocalIpAddress(logger logs.ILogger) (ipv4 string, ipv6 string) {

	// Получаем все доступные сетевые интерфейсы
	interfaces, err := net.Interfaces()
	if err != nil {
		logger.LogError(errors.New("interfaces not found: "+ err.Error()))
	}

	for _, interf := range interfaces {
		// Список адресов для каждого сетевого интерфейса
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}

		if len(addrs) < 1 {
			continue
		}

		if runtime.GOOS == "windows" {
			if !strings.Contains((strings.Split(addrs[1].String(), "/"))[0], "192.168") {
				continue
			}
			return (strings.Split(addrs[1].String(), "/"))[0], (strings.Split(addrs[0].String(), "/"))[0]
		}
		if runtime.GOOS == "linux" {
			if !strings.Contains((strings.Split(addrs[0].String(), "/"))[0], "192.168") {
				continue
			}
			return (strings.Split(addrs[0].String(), "/"))[0], (strings.Split(addrs[1].String(), "/"))[0]
		}
	}
	return
}
