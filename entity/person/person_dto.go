package personEntity

import (
	addressEntity "projetoGo/entity/address"
	animalEntity "projetoGo/entity/animal"
	"time"
)

type PersonDTO struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Birthday time.Time              `json:"birthday"`
	Genre    string                 `json:"genre"`
	Cpf      string                 `json:"cpf"`
	Address  *addressEntity.Address `json:"address"`
	animals  []*animalEntity.Animal
}
