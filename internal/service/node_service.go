package service

import (
	"Network/internal/entites/dto"
	"Network/internal/repository"
	"Network/pkg/logs"
)

type NodeService struct {
	logger     logs.ILogger
	repository repository.INodesReposytory
}

func NewNodeService(logger logs.ILogger, reposytory repository.INodesReposytory) *NodeService {
	service := &NodeService{
		logger:     logger,
		repository: reposytory,
	}
	return service
}

func (nodeService *NodeService) Create(data dto.Node) (int, error) {
	return nodeService.repository.Create(data)
}

func (nodeService *NodeService) GetAll() ([]dto.Node, error){
	return nodeService.repository.GetAll()
}

func (nodeService *NodeService) GetById(nodeId int) (dto.Node, error) {
	return nodeService.repository.GetById(nodeId)
}

func (nodeService *NodeService) Delete(nodeId int) error {
	return nodeService.repository.Delete(nodeId)
}

func (nodeService *NodeService) Update(nodeId int, data dto.Node) error {
	return nodeService.repository.Update(nodeId, data)
}
