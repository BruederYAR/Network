package entites

import "Network/internal/entites/dto"

type HandShake struct { //Информация о узлах при рукопожатии
	Nodes []dto.Node
}

type HandShakeIds struct {
	Ids    []dto.Ids
	Status bool
}
