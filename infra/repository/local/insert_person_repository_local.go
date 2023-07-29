package personRepositoryLocal

import (
	personEntity "projetoGo/entity/person"
)

func Create(person *personEntity.Person) {
	People = append(People, *person)
}
