package protocol

import (
	"Network/internal/entites"
	"Network/internal/entites/dto"
	"bytes"
	"crypto/rsa"
	"encoding/gob"
	"encoding/json"
)

func HandShakeToJson(nodes []dto.Node) ([]byte, error) {
	var handShake = entites.HandShake{Nodes: nodes} //Создание списка адресов для рукопожатия
	return json.Marshal(handShake)
}

func HandShakeIdsToJson(ids []dto.Ids, status bool) ([]byte, error) {
	var handShake = entites.HandShakeIds{Ids: ids, Status: status} //Создание списка адресов для рукопожатия
	return json.Marshal(handShake)
}

func BytesToPublicKey(b []byte) (rsa.PublicKey, error) {
	var buffer bytes.Buffer
	var key rsa.PublicKey

	buffer.Write(b)
	dec := gob.NewDecoder(&buffer)

	err := dec.Decode(&key)
	if err != nil {
		return rsa.PublicKey{}, err
	}

	return key, nil
}

func PublicKeyToBytes(key rsa.PublicKey) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func BytesToPrivateKey(b []byte) (rsa.PrivateKey, error) {
	var buffer bytes.Buffer
	var key rsa.PrivateKey

	buffer.Write(b)
	dec := gob.NewDecoder(&buffer)

	err := dec.Decode(&key)
	if err != nil {
		return rsa.PrivateKey{}, err
	}

	return key, nil
}

func PrivateKeyToBytes(key rsa.PrivateKey) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
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
