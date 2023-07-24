package personEntity

import (
	addressEntity "projetoGo/entity/address"
)

type PersonDTO struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Birthday string                `json:"birthday"`
	Genre    string                `json:"genre"`
	Cpf      string                `json:"cpf"`
	Address  addressEntity.Address `json:"address"`
}
