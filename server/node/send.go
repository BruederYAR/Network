package node

import (
	"Network/internal/entites"
	"Network/internal/entites/dto"
	"Network/pkg/crypt"
	"Network/pkg/protocol"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
)

func (node *Node) HandShake(address string, handshake entites.HandShake) error { //Рукопожатие при первом подключении
	var new_pack = entites.Packege{
		From:      node.Address.IP + node.Address.Port,
		To:        address,
		Name:      node.Name,
		PublicKey: node.PublicKey,
		Type:      node.Types[1],
		Title:     node.Titles[0],
	}

	new_pack.Date, _ = json.Marshal(handshake)

	var test entites.HandShake
	json.Unmarshal(new_pack.Date, &test)

	node.Logger.LogInfo(fmt.Sprintf("Send handshake to %s", new_pack.To))
	return node.Send(&new_pack)
}

func (node *Node) HandShakeIds(address string, status bool) error {
	var new_pack = entites.Packege{
		From:      node.Address.IP + node.Address.Port,
		To:        address,
		Name:      node.Name,
		PublicKey: node.PublicKey,
		Type:      node.Types[1],
		Date:      []byte{},
		Title:     node.Titles[-1],
	}

	ids, err := node.Service.GetAllIds()
	if err != nil {
		node.Logger.LogWarning(err.Error())
		return err
	}

	new_pack.Date, err = protocol.HandShakeIdsToJson(ids, status)
	if err != nil {
		node.Logger.LogWarning(err.Error())
		return err
	}

	node.Logger.LogInfo(fmt.Sprintf("Send handshake ids to %s", new_pack.To))
	return node.Send(&new_pack)
}

func (node *Node) SendMessageTo(To string, message []byte) error {
	var new_pack = entites.Packege{
		From:      node.Address.IP + node.Address.Port,
		Name:      node.Name,
		PublicKey: node.PublicKey,
		Title:     node.Titles[1],
		Type:      node.Types[0],
	}

	var result dto.Node
	if strings.Contains(To, ":") { //Если адрес
		r, err := node.Service.FindByAddress(To)
		if err != nil {
			return errors.New("Find address exception:" + err.Error())
		}
		new_pack.To = To
		result = r
	} else { //Если имя
		res, err := node.Service.FindByName(To)
		if err != nil {
			return err
		}
		result = res
		new_pack.To = result.Address
		if new_pack.To == "" {
			return errors.New("Couldn't find the recipient" + To)
		}
	}

	publicKey, err := node.Service.GetPublickeyById(result.NodeId)
	if err != nil {
		return errors.New("Public key cannot be read")
	}

	new_pack.Date = crypt.RSA_OAEP_Encrypt(message, publicKey)

	return node.Send(&new_pack)
}

func (node *Node) Send(pack *entites.Packege) error { //Отправление данных конкретному пользователю
	conn, err := net.Dial("tcp", pack.To) //Подключаемся
	if err != nil {                       //Если подключение не прошло, забываем о узле
		//node.Service.Delete((node.Service.FindByName(pack.To)))
		return errors.New("Connection error to " + pack.To)
	}
	defer conn.Close()

	byte_array, _ := protocol.ToByteArray(*pack)
	conn.Write(byte_array) //Отправляем
	return nil
}
