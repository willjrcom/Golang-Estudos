package personRepository

import (
	personEntity "projetoGo/entity/person"
)

func InsertPerson(person *personEntity.Person) []personEntity.Person {
	People = append(People, *person)
	return People
}
