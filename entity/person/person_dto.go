package personEntity

import (
	addressEntity "projetoGo/entity/address"
	animalEntity "projetoGo/entity/animal"
	"time"

	"gorm.io/gorm"
)

type PersonDTO struct {
	gorm.Model
	Name     string                 `json:"name"`
	Birthday time.Time              `json:"birthday"`
	Genre    string                 `json:"genre"`
	Cpf      string                 `json:"cpf"`
	Address  *addressEntity.Address `json:"address"`
	animals  []*animalEntity.Animal
}
