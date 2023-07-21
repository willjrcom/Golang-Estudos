package peopleEntity

import addressEntity "projetoGo/entity/address"

type people struct {
	id       int
	name     string
	birthday string
	genre    string
	cpf      string
	address  addressEntity.Address
}

func (p *people) BirthdayIsgreaterThan(birthday string) bool {
	return p.birthday > birthday
}

func (p *people) changeAdress(adress *addressEntity.Address) (err error) {
	err = nil
	p.address = *adress
	return
}

func (p *people) GetAdress() *addressEntity.Address {
	return &p.address
}
