package dto

type Node struct {
	NodeId    []byte `json:"nodeid"`
	Address   string `json:"address"`
	PublicKey []byte `json:"publickey"`
	Name      string `json:"name"`
}
