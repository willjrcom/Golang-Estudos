package personRepositoryLocal

import (
	personEntity "projetoGo/entity/person"
	"time"
)

var People = []personEntity.Person{}

// Populate database
func init() {
	t1 := time.Date(2000, time.January, 26, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2005, time.October, 24, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2013, time.October, 20, 0, 0, 0, 0, time.UTC)
	t4 := time.Date(2005, time.October, 8, 0, 0, 0, 0, time.UTC)

	person, _ := personEntity.NewPersonBuilder("William", "436.377.998-55").WithBirthday(t1).WithGenre("Masculino").Build()
	People = append(People, *person)

	person, _ = personEntity.NewPersonBuilder("Duda", "436.377.998-50").WithBirthday(t2).WithGenre("Feminino").Build()
	People = append(People, *person)

	person, _ = personEntity.NewPersonBuilder("Ana", "436.377.998-51").WithBirthday(t3).WithGenre("Feminino").Build()
	People = append(People, *person)

	person, _ = personEntity.NewPersonBuilder("Nicolas", "436.377.998-52").WithBirthday(t4).WithGenre("Masculino").Build()
	People = append(People, *person)
}
