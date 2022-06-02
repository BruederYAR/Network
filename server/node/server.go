package node

import (
	"Network/internal/entites"
	"Network/internal/entites/dto"
	"Network/pkg/crypt"
	"Network/pkg/protocol"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func handlerServer(node *Node) { //Запуск сервера
	listen, err := net.Listen("tcp", node.Address.Port) //Слушаем определенный порт
	if err != nil {                                     //если есть ошибки вызываем панику
		panic("listen err")
	}
	defer listen.Close() //Ошибок нет - закрываем прослушку
	for {
		conn, err := listen.Accept() //Принимаем подключение
		if err != nil {              //При ошибке выходим из цикла и заного начинаем слушать
			break
		}
		go handleConnection(node, conn) //читаем данные
	}
}

func handleConnection(node *Node, conn net.Conn) { //Читаем данные
	defer conn.Close()
	var (
		buffer  = make([]byte, 512)
		message []byte
		pack    entites.Packege
	)
	for {
		lenght, err := conn.Read(buffer) //Читаем всё в буфер
		if err != nil {
			break
		}

		message = append(message[:], buffer[:lenght]...) //Записываем только до длины, чтобы убрать мусор
	}

	pack, err := protocol.ToPackege(message) //Переводим в пакет
	if err != nil {                          //При ошибке метод закрываеться
		return
	}

	if pack.Name == node.Name { //Если вызвали сами себя, то выключаем метод
		return
	}

	node.Service.Create(dto.Node{ Address: pack.From, Name: pack.Name, PublicKey: pack.PublicKey.N.String()})
	node.ConnectTo([]string{pack.From}, pack.Name, pack.PublicKey) //записываем того, кто отослал данные

	WorkingWithData(node, pack)
}

func WorkingWithData(node *Node, pack entites.Packege) {
	switch pack.Title {
	case node.Titles[1]: //date
		switch pack.Type {
		case node.Types[0]:
			message := crypt.RSA_OAEP_Decrypt(pack.Date, node.PrivateKey)
			fmt.Println(string(message)) //Выводим данные
		}

	case node.Titles[0]: //Рукопожатие handshake
		var handShake entites.HandShake
		json.Unmarshal(pack.Date, &handShake) //Забираем список узлов

		if handShake.Status { //Если начало рукопожатия
			node.HandShake(pack.From, false) //Отправляем узлы обратно
		}

		for _, local_node := range handShake.Nodes { //Добавляем узлы в локальные список
			if node.Connections[local_node.Address] == nil && local_node.Name != node.Name { //Если узел, который нам прислали был не известен, то выполняем рукопожатие с ним
				node.HandShake(local_node.Address, true)
			}
			node.Connections[local_node.Address] = &entites.NodeInfo{Name: local_node.Name, PublicKey: local_node.PublicKey}
		}
	}
}

func (node *Node) ConnectTo(addresses []string, name string, publickey rsa.PublicKey) { //Добавление в список подключений
	for _, addr := range addresses {
		if addr == "" {
			node.Logger.LogError(errors.New("empty address"))
		} else {
			node.Connections[addr] = &entites.NodeInfo{Name: name, PublicKey: publickey}
		}
	}
}
