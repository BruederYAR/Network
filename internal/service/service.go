package service

import (
	"Network/internal/entites/dto"
	"crypto/rsa"
)

type INodeService interface {
	Create(data dto.Node) (int, error)
	Add(data dto.Node) (int, error)
	IfExist(data dto.Node) (bool, error)
	GetAll() ([]dto.Node, error)
	GetAllIds() ([]dto.Ids, error)
	GetById(nodeId []byte) (dto.Node, error)
	GetPublickeyById(nodeId []byte) (rsa.PublicKey, error)
	FindByName(name string) (dto.Node, error)
	FindByAddress(address string) (dto.Node, error)
	IsAddressExist(address string) (bool, error)
	Delete(nodeId []byte) (int, error)
	Update(data dto.Node) (int, error)
}
