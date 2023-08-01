package presenter

import (
	"errors"
	"net/url"
	addressEntity "projetoGo/entity/address"
	personEntity "projetoGo/entity/person"
	utils "projetoGo/infra/utils/treino"
)

func UrlValuesToPerson(values url.Values) (*personEntity.Person, error) {
	if values == nil {
		return nil, errors.New("Não foi encontrado nenhum valor válido!")
	}

	name := values.Get("name")
	cpf := values.Get("cpf")

	if name == "" || cpf == "" {
		return nil, errors.New("Nome ou cpf está vazio!")
	}

	address := addressEntity.NewAddressBuilder().WithStreet("Rua Piedade").WithNumber(226).WithCity("Sorocaba").WithState("SP").WithCountry("BR").Build()

	builder := personEntity.NewPersonBuilder(name, cpf)
	builder.WithBirthday(utils.StringtoDate(values.Get("birthday")))
	builder.WithGenre(values.Get("genre"))
	builder.WithAddress(address)
	person, err := builder.Build()

	if err != nil {
		return nil, err[0]
	}

	return person, nil
}
