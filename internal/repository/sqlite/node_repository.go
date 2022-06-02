package sqlite_repository

import (
	"Network/internal/entites/dto"
	"Network/pkg/input"
	"Network/pkg/logs"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type NodeRepository struct {
	config input.Congig
	logger logs.ILogger
}

func (nodeRepository *NodeRepository) Create(data dto.Node) (int, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("INSERT INTO Node (Address, PublicKey, Name) VALUES (?, ?, ?)")
	result, err := statement.Exec(data.Address, data.PublicKey, data.Name)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return 0, err
	}

	number, _ := result.RowsAffected()
	return int(number), nil
}

func (nodeRepository *NodeRepository) GetAll() ([]dto.Node, error) {
	return []dto.Node{}, nil
}

func (nodeRepository *NodeRepository) GetById(nodeId int) (dto.Node, error){
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("SELECT NodeId, Address, PublicKey, Name FROM Node WHERE NodeId = (?)")
	rows, err := statement.Query(nodeId)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	
	result := nodeRepository.rowsToNodes(*rows)

	if len(result) == 0{
		nodeRepository.logger.LogWarning(fmt.Sprintf("nodeid: %d does not exist", nodeId))
		return dto.Node{}, errors.New(fmt.Sprintf("nodeid: %d does not exist", nodeId))
	}

	return result[0], nil
}

func (nodeRepository *NodeRepository) Delete(nodeId int) error{
	return nil
}

func (nodeRepository *NodeRepository) Update(nodeId int, node dto.Node) error {
	return nil
}

func NewNodeRepository(cfg input.Congig, logger logs.ILogger) *NodeRepository {
	nodeRepository := &NodeRepository{
		config: cfg,
		logger: logger,
	}
	return nodeRepository
}

func (nodeRepository *NodeRepository) rowsToNodes(rows sql.Rows) []dto.Node{
	results := []dto.Node{}
	for rows.Next(){
        n := dto.Node{}
        err := rows.Scan(&n.NodeId, &n.Address, &n.PublicKey, &n.Name)
        if err != nil{
			nodeRepository.logger.LogWarning(err.Error())
            continue
        }
        results = append(results, n)
    }

	return results
}
