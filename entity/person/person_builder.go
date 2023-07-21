package personEntity

import addressEntity "projetoGo/entity/address"

type personBuilder struct {
	person Person
}

func NewPersonBuilder(name string, cpf string) *personBuilder {
	return &personBuilder{
		person: Person{
			name: name,
			cpf:  cpf,
		},
	}
}

func (p *personBuilder) WithBirthday(birthday string) *personBuilder {
	p.person.birthday = birthday
	return p
}

func (p *personBuilder) WithGenre(genre string) *personBuilder {
	p.person.genre = genre
	return p
}

func (p *personBuilder) WithAddress(address *addressEntity.Address) *personBuilder {
	p.person.address = *address
	return p
}

func (p *personBuilder) Build() *Person {
	return &p.person
}
