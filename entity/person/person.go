package personEntity

import (
	"encoding/json"
	"math"
	addressEntity "projetoGo/entity/address"
	"time"
)

type Person struct {
	ID       string
	name     string
	birthday string
	genre    string
	cpf      string
	address  addressEntity.Address
}

func (p *Person) GetAdress() *addressEntity.Address {
	return &p.address
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetCpf() string {
	return p.cpf
}

func (p *Person) BirthdayIsgreaterThan(birthday string) bool {
	return p.birthday > birthday
}

func (p *Person) changeAdress(adress *addressEntity.Address) (err error) {
	err = nil
	p.address = *adress
	return
}

func (p *Person) isChild() (bool, error) {
	dateBirthday, err := time.Parse("2006-01-02", p.birthday)

	if err != nil {
		return false, err
	}
	diff := dateBirthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years < 12 {
		return true, nil
	}

	return false, nil
}

func (p *Person) isTeen() (bool, error) {
	dateBirthday, err := time.Parse("2006-01-02", p.birthday)

	if err != nil {
		return false, err
	}
	diff := dateBirthday.Sub(time.Now())
	years := math.Abs(float64(diff.Hours() / 24 / 365))

	if years >= 12 && years < 18 {
		return true, nil
	}

	return false, nil
}

func (p *Person) isAdult() (bool, error) {
	dateBirthday, err := time.Parse("2006-01-02", p.birthday)

	if err != nil {
		return false, err
	}
	diff := dateBirthday.Sub(time.Now())
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
