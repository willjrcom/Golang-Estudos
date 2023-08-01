package personEntity

import (
	"encoding/json"
	"errors"
	"math"
	addressEntity "projetoGo/entity/address"
	animalEntity "projetoGo/entity/animal"
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name      string    `validate:"required,min=5"`
	Birthday  time.Time `validate:"required"`
	Genre     string    `validate:"required"`
	Cpf       string    `validate:"required,cpf"`
	AddressID uint
	Address   *addressEntity.Address `gorm:"foreignkey:AddressID" validate:"required"`
	Animals   []*animalEntity.Animal
}

func (p *Person) GetId() uint {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetCpf() string {
	return p.Cpf
}

func (p *Person) GetAdress() *addressEntity.Address {
	return p.Address
}

func (p *Person) BirthdayIsgreaterThan(birthday time.Time) bool {
	return p.Birthday.GoString() > birthday.GoString()
}

func (p *Person) changeAdress(adress *addressEntity.Address) {
	p.Address = adress
}

func (p *Person) AdoptAnimal(animal *animalEntity.Animal) {
	p.Animals = append(p.Animals, animal)
}

func (p *Person) DonateAnimal(animalToRemove *animalEntity.Animal) error {
	for index, animal := range p.Animals {
		if animal == animalToRemove {
			p.Animals = append(p.Animals[:index], p.Animals[index+1:]...)
			return nil
		}
	}

	return errors.New("DonateAnimal() - Animal not found")
}

func (p *Person) isChild() (bool, error) {
	diff := p.Birthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years < 12 {
		return true, nil
	}

	return false, nil
}

func (p *Person) isTeen() (bool, error) {
	diff := p.Birthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years >= 12 && years < 18 {
		return true, nil
	}

	return false, nil
}

func (p *Person) isAdult() (bool, error) {
	diff := p.Birthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years >= 18 {
		return true, nil
	}

	return false, nil
}

func (p *Person) GetJson() ([]byte, error) {
	PersonDTO := PersonDTO{
		Model:    p.Model,
		Name:     p.Name,
		Birthday: p.Birthday,
		Genre:    p.Genre,
		Cpf:      p.Cpf,
		Address:  p.Address,
	}
	return json.Marshal(PersonDTO)
}
