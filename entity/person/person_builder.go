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

func (p *personBuilder) BuildDto() *PersonDTO {
	return &PersonDTO{
		ID:       p.person.ID,
		Name:     p.person.name,
		Birthday: p.person.birthday,
		Genre:    p.person.genre,
		Cpf:      p.person.cpf,
		Address:  p.person.address,
	}
}
