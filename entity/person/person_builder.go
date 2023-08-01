package personEntity

import (
	addressEntity "projetoGo/entity/address"
	structValidator "projetoGo/infra/validator"
	"time"
)

type personBuilder struct {
	person *Person
}

func NewPersonBuilder(name string, cpf string) *personBuilder {
	return &personBuilder{
		person: &Person{
			Name: name,
			Cpf:  cpf,
		},
	}
}

func (p *personBuilder) WithBirthday(birthday time.Time) *personBuilder {
	p.person.Birthday = birthday
	return p
}

func (p *personBuilder) WithGenre(genre string) *personBuilder {
	p.person.Genre = genre
	return p
}

func (p *personBuilder) WithAddress(address *addressEntity.Address) *personBuilder {
	p.person.Address = address
	return p
}

func (p *personBuilder) Build() (*Person, []error) {
	return structValidator.Validate(p.person)
}

func (p *personBuilder) BuildDto() (*PersonDTO, []error) {
	// 	_, err := structValidator.Validate(&p.person)

	// 	if err != nil {
	// 		return new(PersonDTO), err
	// 	}

	return &PersonDTO{
		Model:    p.person.Model,
		Name:     p.person.Name,
		Birthday: p.person.Birthday,
		Genre:    p.person.Genre,
		Cpf:      p.person.Cpf,
		Address:  p.person.Address,
	}, nil
}
