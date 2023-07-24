package person_repository

import (
	"errors"
	personEntity "projetoGo/entity/person"
)

var people = []personEntity.Person{}

// Populate database
func init() {
	people = append(people, *personEntity.NewPersonBuilder("William", "436.377.998-55").WithBirthday("26/01/2000").WithGenre("Masculino").Build())
	people = append(people, *personEntity.NewPersonBuilder("Duda", "436.377.998-50").WithBirthday("24/10/2005").WithGenre("Feminino").Build())
	people = append(people, *personEntity.NewPersonBuilder("Ana", "436.377.998-51").WithBirthday("20/10/2013").WithGenre("Feminino").Build())
	people = append(people, *personEntity.NewPersonBuilder("Nicolas", "436.377.998-52").WithBirthday("08/10/2004").WithGenre("Masculino").Build())
}

func FindAll() []personEntity.Person {
	return people
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
