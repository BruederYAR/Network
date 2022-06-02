package entites

import "crypto/rsa"

type Node struct { //Узел адрес|Имя
	Address   string
	PublicKey rsa.PublicKey
	Name      string
}