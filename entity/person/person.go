package personEntity

import addressEntity "projetoGo/entity/address"

type Person struct {
	id       int
	name     string
	birthday string
	genre    string
	cpf      string
	address  addressEntity.Address
}

func (p *Person) BirthdayIsgreaterThan(birthday string) bool {
	return p.birthday > birthday
}

func (p *Person) changeAdress(adress *addressEntity.Address) (err error) {
	err = nil
	p.address = *adress
	return
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
