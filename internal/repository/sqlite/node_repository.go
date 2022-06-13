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

func (nodeRepository *NodeRepository) Add(data dto.Node) (int, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("INSERT INTO Node (NodeId, Address, PublicKey, Name) VALUES (?, ?, ?, ?)")
	result, err := statement.Exec(data.NodeId, data.Address, data.PublicKey, data.Name)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return 0, err
	}

	number, _ := result.RowsAffected()
	return int(number), nil
}

func (nodeRepository *NodeRepository) CountExist(data dto.Node) (int, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return 0, err
	}
	defer database.Close()

	var result int
	err = database.QueryRow("SELECT Count(*) FROM Node WHERE Address = ? and Name = ? GROUP BY NodeId, Address, PublicKey, Name", data.Address, data.Name).Scan(&result)

	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return 0, err
	}

	return result, nil
}

func (nodeRepository *NodeRepository) GetAll() ([]dto.Node, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("SELECT NodeId, Address, PublicKey, Name FROM Node")
	rows, err := statement.Query()
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return nil, err
	}

	result := nodeRepository.rowsToNodes(*rows)

	return result, nil
}

func (nodeRepository *NodeRepository) GetAllIds() ([]dto.Ids, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("SELECT NodeId FROM Node")
	rows, err := statement.Query()
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return nil, err
	}

	result := nodeRepository.rowsToIds(*rows)

	return result, nil
}

func (nodeRepository *NodeRepository) FindByAddress(address string) (dto.Node, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	var result dto.Node
	err = database.QueryRow("SELECT NodeId, Address, PublicKey, Name FROM Node WHERE Address = ?", address).Scan(&result.NodeId, &result.Address, &result.PublicKey, &result.Name)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return dto.Node{}, err
	}

	return result, nil
}

func (nodeRepository *NodeRepository) GetById(nodeId []byte) (dto.Node, error) {
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

	if len(result) == 0 {
		nodeRepository.logger.LogWarning(fmt.Sprintf("nodeid: %d does not exist", nodeId))
		return dto.Node{}, errors.New(fmt.Sprintf("nodeid: %d does not exist", nodeId))
	}

	return result[0], nil
}

func (nodeRepository *NodeRepository) FindByName(name string) (dto.Node, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	//Like \"%$1%\"
	rows, err := database.Query("SELECT NodeId, Address, PublicKey, Name FROM Node WHERE Name = ?", name)
	if err != nil {
		return dto.Node{}, err
	}
	result := nodeRepository.rowsToNodes(*rows)

	if len(result) == 0 {
		return dto.Node{}, errors.New(fmt.Sprintf("Name: %s does not exist", name))
	}

	return result[0], nil
}

func (nodeRepository *NodeRepository) Delete(nodeId []byte) (int, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("DELETE FROM Node WHERE NodeId = ?")
	result, err := statement.Exec(nodeId)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return 0, err
	}

	number, _ := result.RowsAffected()
	return int(number), nil
}

func (nodeRepository *NodeRepository) Update(node dto.Node) (int, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("UPDATE Node SET Address = ?, PublicKey = ?, Name = ? WHERE NodeId = ?;")
	result, err := statement.Exec(node.Address, node.PublicKey, node.Name, node.NodeId)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return 0, err
	}

	number, _ := result.RowsAffected()
	return int(number), nil
}

func (nodeRepository *NodeRepository) GetPublickeyById(nodeId []byte) ([]byte, error) {
	database, err := sql.Open("sqlite3", nodeRepository.config.Connect)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
	}
	defer database.Close()

	statement, _ := database.Prepare("SELECT PublicKey FROM Node WHERE NodeId = ?")
	var key []byte
	err = statement.QueryRow(nodeId).Scan(&key)
	if err != nil {
		nodeRepository.logger.LogWarning(err.Error())
		return nil, err
	}

	return key, nil
}

func NewNodeRepository(cfg input.Congig, logger logs.ILogger) *NodeRepository {
	nodeRepository := &NodeRepository{
		config: cfg,
		logger: logger,
	}
	return nodeRepository
}

func (nodeRepository *NodeRepository) rowsToNodes(rows sql.Rows) []dto.Node {
	results := []dto.Node{}
	for rows.Next() {
		n := dto.Node{}
		err := rows.Scan(&n.NodeId, &n.Address, &n.PublicKey, &n.Name)
		if err != nil {
			nodeRepository.logger.LogWarning(err.Error())
			continue
		}
		results = append(results, n)
	}

	return results
}

func (nodeRepository *NodeRepository) rowsToIds(rows sql.Rows) []dto.Ids {
	results := []dto.Ids{}

	for rows.Next() {
		n := dto.Ids{}
		err := rows.Scan(&n.Id)
		if err != nil {
			nodeRepository.logger.LogWarning(err.Error())
			continue
		}
		results = append(results, n)
	}

	return results
}
