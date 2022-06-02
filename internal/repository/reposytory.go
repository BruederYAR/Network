package repository

import "Network/internal/entites/dto"

type INodesReposytory interface {
	Create(data dto.Node) (int, error)
	GetAll() ([]dto.Node, error)
	GetById(nodeId int) (dto.Node, error)
	Delete(nodeId int) error
	Update(nodeId int, data dto.Node) error
}

