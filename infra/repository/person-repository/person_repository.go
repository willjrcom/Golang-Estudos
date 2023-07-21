package person_repository

import (
	"errors"
	personEntity "projetoGo/entity/person"
	"strconv"
)

var people = []personEntity.Person{}

// Populate database
func init() {
	for i := 0; i < 10; i++ {
		person := personEntity.NewPersonBuilder("name"+strconv.Itoa(i), "cpf"+strconv.Itoa(i)).WithBirthday("birthday" + strconv.Itoa(i)).WithGenre("genre" + strconv.Itoa(i)).Build()
		people = append(people, *person)
	}
}

func FindByName(name string) (personEntity.Person, error) {
	for _, p := range people {
		if p.GetName() == name {
			return p, nil
		}
	}
	return personEntity.Person{}, errors.New("FindByName: person not found")
}

func FindByCpf(cpf string) (personEntity.Person, error) {
	for _, p := range people {
		if p.GetCpf() == cpf {
			return p, nil
		}
	}
	return personEntity.Person{}, errors.New("GetCpf: person not found")
}
