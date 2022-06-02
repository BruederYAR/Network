package node

import (
	"Network/internal/entites"
	"Network/pkg/crypt"
	"Network/pkg/protocol"
	"errors"
	"fmt"
	"net"
	"strings"
)

func (node *Node) recipientSearch(To string) (string, error){
	if strings.Contains(To, ":") { //Если адрес
		if node.Connections[To] == nil { //Если адреса нет в списке адресов, выполняем рукопожатия
			node.HandShake(To, true)
			return "", errors.New("[err] Unknown address, completed handshake try again")
		}
		return To, nil
	} else { //Если имя
		for key, item := range node.Connections { //Поиск адреса по имени
			if item != nil{
				if To == item.Name {
					return key, nil
				}
			} 
		}
		if To == "" {
			return "", errors.New("Couldn't find the recipient " + To)
		}
	}
	return "", errors.New("Couldn't find the recipient " + To)
}


func (node *Node) HandShake(address string, status bool) error { //Рукопожатие при первом подключении
	var new_pack = entites.Packege{
		From:      node.Address.IP + node.Address.Port,
		To:        address,
		Name:      node.Name,
		PublicKey: node.PublicKey,
		Type:      node.Types[1],
		Date:      []byte{},
		Title:     node.Titles[0],
	}

	if !status {
		node.Logger.LogInfo(fmt.Sprintf("HandShake from %s to %s", new_pack.From , new_pack.To))
	}

	new_pack.Date, _ = protocol.HandShakeToJson(node.Connections, status) //Статус нужен для того, чтобы определять кто начал рукопожатие, иначе сеть будет постоянно их слать

	return node.Send(&new_pack)
}

func (node *Node) SendMessageTo(To string, message []byte) (error) {
	var new_pack = entites.Packege{
		From:      node.Address.IP + node.Address.Port,
		Name:      node.Name,
		PublicKey: node.PublicKey,
		Title:     node.Titles[1],
		Type:      node.Types[0],
	}

	if strings.Contains(To, ":") { //Если адрес
		if node.Connections[To] == nil { //Если адреса нет в списке адресов, выполняем рукопожатия
			node.HandShake(To, true)
			return errors.New("Unknown address, completed handshake try again")
		}
		new_pack.To = To
	} else { //Если имя
		for key, item := range node.Connections { //Поиск адреса по имени
			if item != nil{
				if To == item.Name {
					new_pack.To = key
				}
			} 
		}
		if new_pack.To == "" {
			return errors.New("Couldn't find the recipient" + To)
		}
	}

	new_pack.Date = crypt.RSA_OAEP_Encrypt(message, node.Connections[new_pack.To].PublicKey)

	return node.Send(&new_pack)
}

func (node *Node) SendMessageToAll(message []byte) error { //Отправка сообщений всем
	var new_pack = entites.Packege{
		From:      node.Address.IP + node.Address.Port,
		Name:      node.Name,
		PublicKey: node.PublicKey,
		Title:     node.Titles[1],
		Type:      node.Types[0],
	}
	for addr := range node.Connections { //Переборам отправляем сообщение
		new_pack.To = addr
		new_pack.Date = crypt.RSA_OAEP_Encrypt(message, node.Connections[new_pack.To].PublicKey)
		return node.Send(&new_pack)
	}
	return nil
}

func (node *Node) Send(pack *entites.Packege) error { //Отправление данных конкретному пользователю
	conn, err := net.Dial("tcp", pack.To) //Подключаемся
	if err != nil {                       //Если подключение не прошло, забываем о узле
		delete(node.Connections, pack.To)
		return errors.New("Connection error to "+ pack.To)
	}
	defer conn.Close()

	byte_array, _ := protocol.ToByteArray(*pack)
	conn.Write(byte_array) //Отправляем
	return nil
}