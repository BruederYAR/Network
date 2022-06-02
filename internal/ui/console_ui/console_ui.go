package ui

import (
	"Network/server/node"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConsoleClient(n *node.Node) { //Клиент
	for {
		message := InputString()
		splited := strings.Split(message, " ") //Берём дынные и разбиваем

		switch splited[0] { //Команды клиента
		case "/exit":
			os.Exit(0)
		case "/help":
			fmt.Println("^/exit", "Выход из программы")
			fmt.Println("^/connect", "Присоеденится к узлу. Ключи: 1)адрес(ip:port)")
			fmt.Println("^/network", "Вывести все присоедененные узлы и собственный адрес")
			fmt.Println("^/m", "Отправить сообщение. Ключи: 1)адрес(ip:port или имя) 2)сообщение")
		case "/connect":
			err := n.HandShake(splited[1], true)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "/network":
			PrintNetwork(n)
		case "/m":
			if len(splited) < 3 {
				fmt.Println("Не верное кол-во аргументов")
				continue
			}
			err := n.SendMessageTo(splited[1], []byte(splited[2]))
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			err := n.SendMessageToAll([]byte(message))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func PrintNetwork(n *node.Node) { //Ввывод всех подключений
	fmt.Println("local address " + n.Address.IP + n.Address.Port)
	for addr := range n.Connections {
		fmt.Println(n.Connections[addr].Name + "|" + addr)
	}
}

func InputString() string { //Чтение с консоли
	msg, _ := bufio.NewReader(os.Stdin).ReadString('\n') //Читаем буфер
	return strings.Replace(msg, "\n", "", -1) + " "      //Убираем переходы на следующую строку и возвращаем сообщение
}
