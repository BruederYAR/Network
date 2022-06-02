package entites

import "crypto/rsa"

type Packege struct {
	To        string
	From      string
	Title     string
	Name      string
	PublicKey rsa.PublicKey
	Type      string
	Date      []byte
}
