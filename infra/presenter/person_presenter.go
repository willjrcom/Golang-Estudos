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

	address := addressEntity.NewAddressBuilder().WithStreet("Rua Piedade").WithNumber(226).WithCity("Sorocaba").WithState("SP").WithCountry("BR").Build()

	person := &personEntity.Person{
		Name:     values.Get("name"),
		Cpf:      values.Get("cpf"),
		Address:  address,
		Birthday: utils.StringtoDate(values.Get("birthday")),
		Genre:    values.Get("genre"),
	}

	return person, nil
}
