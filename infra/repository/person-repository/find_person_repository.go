package personRepository

import (
	"errors"
	"fmt"
	personEntity "projetoGo/entity/person"
	"strings"
)

var People = []personEntity.Person{}

// Populate database
func init() {
	People = append(People, *personEntity.NewPersonBuilder("William", "436.377.998-55").WithBirthday("26/01/2000").WithGenre("Masculino").Build())
	People = append(People, *personEntity.NewPersonBuilder("Duda", "436.377.998-50").WithBirthday("24/10/2005").WithGenre("Feminino").Build())
	People = append(People, *personEntity.NewPersonBuilder("Ana", "436.377.998-51").WithBirthday("20/10/2013").WithGenre("Feminino").Build())
	People = append(People, *personEntity.NewPersonBuilder("Nicolas", "436.377.998-52").WithBirthday("08/10/2004").WithGenre("Masculino").Build())
}

func FindAll() []personEntity.Person {
	return People
}

func FindByName(name string) (personEntity.Person, error) {
	for _, p := range People {
		fmt.Println(p.GetName() + " == " + name)
		if strings.ToLower(p.GetName()) == name {
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
	return personEntity.Person{}, errors.New("GetCpf: person not found")
}
