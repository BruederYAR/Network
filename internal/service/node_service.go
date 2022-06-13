package service

import (
	"Network/internal/entites/dto"
	"Network/internal/repository"
	"Network/pkg/logs"
	"Network/pkg/protocol"
	"crypto/rsa"
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

func (nodeService *NodeService) Add(data dto.Node) (int, error) {
	return nodeService.repository.Add(data)
}

func (nodeService *NodeService) IfExist(data dto.Node) (bool, error) {
	count, err := nodeService.repository.CountExist(data)

	if count == 0 || err != nil {
		return false, err
	}

	return true, nil
}

func (nodeService *NodeService) GetAll() ([]dto.Node, error) {
	return nodeService.repository.GetAll()
}

func (nodeService *NodeService) GetAllIds() ([]dto.Ids, error) {
	return nodeService.repository.GetAllIds()
}

func (nodeService *NodeService) GetById(nodeId []byte) (dto.Node, error) {
	return nodeService.repository.GetById(nodeId)
}

func (nodeService *NodeService) FindByName(name string) (dto.Node, error) {
	return nodeService.repository.FindByName(name)
}

func (nodeService *NodeService) Delete(nodeId []byte) (int, error) {
	return nodeService.repository.Delete(nodeId)
}

func (nodeService *NodeService) Update(data dto.Node) (int, error) {
	return nodeService.repository.Update(data)
}

func (nodeService *NodeService) GetPublickeyById(nodeId []byte) (rsa.PublicKey, error) {
	key, err := nodeService.repository.GetPublickeyById(nodeId)
	if err != nil {
		return rsa.PublicKey{}, err
	}

	return protocol.BytesToPublicKey(key)
}

func (nodeService *NodeService) FindByAddress(address string) (dto.Node, error) {
	return nodeService.repository.FindByAddress(address)
}

func (nodeService *NodeService) IsAddressExist(address string) (bool, error) {
	nodes, err := nodeService.repository.FindByAddress(address)
	if err != nil {
		return false, err
	}
	if len(nodes.Name) <= 0 {
		return false, nil
	}

	return true, nil
}
