package service

import "Network/internal/entites/dto"

type INodeService interface {
	Create(data dto.Node) (int, error)
	GetAll() ([]dto.Node, error)
	GetById(nodeId int) (dto.Node, error)
	Delete(nodeId int) error
	Update(nodeId int, data dto.Node) error
}