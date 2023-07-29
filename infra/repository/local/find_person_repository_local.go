package personRepositoryLocal

import (
	"errors"
	personEntity "projetoGo/entity/person"
	"strings"
)

func FindAll() ([]personEntity.Person, error) {
	return People, nil
}

func FindByName(name string) (personEntity.Person, error) {
	for _, p := range People {

		if strings.ToLower(p.GetName()) == strings.ToLower(name) {
			return p, nil
		}
	}
	return personEntity.Person{}, errors.New("FindByName: person not found")
}

func FindByCpf(cpf string) (personEntity.Person, error) {
	for _, p := range People {
		if p.GetCpf() == cpf {
			return p, nil
		}
	}
	return personEntity.Person{}, errors.New("FindByCpf: person not found")
}
