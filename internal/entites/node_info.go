package entites

import "crypto/rsa"

type NodeInfo struct {
	Name      string
	PublicKey rsa.PublicKey
}
