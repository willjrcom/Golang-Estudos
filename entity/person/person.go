package personEntity

import (
	"encoding/json"
	"errors"
	"math"
	addressEntity "projetoGo/entity/address"
	animalEntity "projetoGo/entity/animal"
	"time"
)

type Person struct {
	ID       string
	name     string                 `validate:"required,min=5"`
	birthday time.Time              `validate:"required"`
	genre    string                 `validate:"required"`
	cpf      string                 `validate:"required,cpf"`
	address  *addressEntity.Address `validate:"required"`
	animals  []*animalEntity.Animal
}

func (p *Person) GetId() string {
	return p.ID
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetCpf() string {
	return p.cpf
}

func (p *Person) GetAdress() *addressEntity.Address {
	return p.address
}

func (p *Person) BirthdayIsgreaterThan(birthday time.Time) bool {
	return p.birthday.GoString() > birthday.GoString()
}

func (p *Person) changeAdress(adress *addressEntity.Address) {
	p.address = adress
}

func (p *Person) AdoptAnimal(animal *animalEntity.Animal) {
	p.animals = append(p.animals, animal)
}

func (p *Person) DonateAnimal(animalToRemove *animalEntity.Animal) error {
	for index, animal := range p.animals {
		if animal == animalToRemove {
			p.animals = append(p.animals[:index], p.animals[index+1:]...)
			return nil
		}
	}

	return errors.New("DonateAnimal() - Animal not found")
}

func (p *Person) isChild() (bool, error) {
	diff := p.birthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years < 12 {
		return true, nil
	}

	return false, nil
}

func (p *Person) isTeen() (bool, error) {
	diff := p.birthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years >= 12 && years < 18 {
		return true, nil
	}

	return false, nil
}

func (p *Person) isAdult() (bool, error) {
	diff := p.birthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years >= 18 {
		return true, nil
	}

	return false, nil
}

func (p *Person) GetJson() ([]byte, error) {
	PersonDTO := PersonDTO{
		ID:       p.ID,
		Name:     p.name,
		Birthday: p.birthday,
		Genre:    p.genre,
		Cpf:      p.cpf,
		Address:  p.address,
	}
	return json.Marshal(PersonDTO)
}
