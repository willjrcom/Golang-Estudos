package peopleEntity

import addressEntity "projetoGo/entity/address"

type peopleBuilder struct {
	people
}

func NewPeopleBuilder(name string, cpf string) *peopleBuilder {
	return &peopleBuilder{
		people: people{
			name: name,
			cpf:  cpf,
		},
	}
}

func (p *peopleBuilder) WithBirthday(birthday string) *peopleBuilder {
	p.people.birthday = birthday
	return p
}

func (p *peopleBuilder) WithGenre(genre string) *peopleBuilder {
	p.genre = genre
	return p
}

func (p *peopleBuilder) WithAddress(address *addressEntity.Address) *peopleBuilder {
	p.address = *address
	return p
}

func (p *peopleBuilder) Build() *people {
	return &p.people
}
