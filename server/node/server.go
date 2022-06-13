package node

import (
	"Network/internal/entites"
	"Network/internal/entites/dto"
	"Network/pkg/crypt"
	"Network/pkg/protocol"
	"bytes"
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

	case node.Titles[-1]: //Рукопожатие handshakeids
		var handShake entites.HandShakeIds
		json.Unmarshal(pack.Date, &handShake)

		node.Logger.LogInfo(fmt.Sprintf("Accepted ids from %s", pack.From))

		localIds, err := node.Service.GetAllIds()
		if err != nil {
			node.Logger.LogError(err)
		}
		var sendIds entites.HandShake

		isIdExist := func(val []byte, list []dto.Ids) bool {
			for _, i := range list {
				if bytes.Equal(i.Id, val) {
					return true
				}
			}
			return false
		}

		for _, i := range localIds {
			if !isIdExist(i.Id, handShake.Ids) {
				n, err := node.Service.GetById(i.Id)
				if err != nil {
					node.Logger.LogError(err)
				}
				sendIds.Nodes = append(sendIds.Nodes, n)
			}
		}

		node.Logger.LogInfo(fmt.Sprintf("Send ids to %s. Data: %s", pack.From, sendIds.Nodes))

		if handShake.Status {
			node.HandShakeIds(pack.From, false)
		}

		node.HandShake(pack.From, sendIds)

	case node.Titles[0]: //Рукопожатие handshake
		var handShake entites.HandShake
		json.Unmarshal(pack.Date, &handShake) //Забираем список узлов

		//Добавление новых узлов
		node.Logger.LogInfo(fmt.Sprintf("Accepted handshake from %s", pack.From))
		for _, i := range handShake.Nodes {
			_, err := node.Service.Add(i)
			if err != nil {
				node.Logger.LogError(errors.New(err.Error() + " " + i.Name + " " + i.Address))
			}
		}

		//Рукопожатие с новыми узлами
		for _, i := range handShake.Nodes {
			node.HandShakeIds(i.Address, true)
		}

		node.Logger.LogInfo(fmt.Sprintf("Add %b nodes", len(handShake.Nodes)))
	}
}
