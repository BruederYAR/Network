package protocol

import (
	"Network/internal/entites"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func HandShakeToJson(nodes map[string]*entites.NodeInfo, status bool) ([]byte, error) {
	var handShake = entites.HandShake{} //Создание списка адресов для рукопожатия
	for addr := range nodes {
		handShake.Nodes = append(handShake.Nodes, entites.Node{Address: addr, Name: nodes[addr].Name, PublicKey: nodes[addr].PublicKey})
	}
	handShake.Status = status

	return json.Marshal(handShake)
}

func ToByteArray(pack entites.Packege) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(pack)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func ToPackege(message []byte) (entites.Packege, error) {
	var buffer bytes.Buffer
	var pack entites.Packege

	buffer.Write(message)
	dec := gob.NewDecoder(&buffer)

	err := dec.Decode(&pack)
	if err != nil {
		return entites.Packege{}, err
	}

	return pack, nil
}