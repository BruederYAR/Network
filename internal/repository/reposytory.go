package repository

import (
	"Network/internal/entites/dto"
)

type INodesReposytory interface {
	Create(data dto.Node) (int, error)
	Add(data dto.Node) (int, error)
	CountExist(data dto.Node) (int, error)
	GetAll() ([]dto.Node, error)
	GetPublickeyById(nodeId []byte) ([]byte, error)
	GetAllIds() ([]dto.Ids, error)
	GetById(nodeId []byte) (dto.Node, error)
	FindByName(name string) (dto.Node, error)
	FindByAddress(address string) (dto.Node, error)
	Delete(nodeId []byte) (int, error)
	Update(data dto.Node) (int, error)
}
