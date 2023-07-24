package fileManager

import (
	"encoding/json"
	personEntity "projetoGo/entity/person"
	"testing"
)

func TestWriteData(t *testing.T) {
	people := []personEntity.PersonDTO{}
	people = append(people, *personEntity.NewPersonBuilder("William", "436.377.998-55").WithBirthday("26/01/2000").WithGenre("Masculino").BuildDto())
	people = append(people, *personEntity.NewPersonBuilder("Duda", "436.377.998-50").WithBirthday("24/10/2005").WithGenre("Feminino").BuildDto())
	people = append(people, *personEntity.NewPersonBuilder("Ana", "436.377.998-51").WithBirthday("20/10/2013").WithGenre("Feminino").BuildDto())
	people = append(people, *personEntity.NewPersonBuilder("Nicolas", "436.377.998-52").WithBirthday("08/10/2004").WithGenre("Masculino").BuildDto())

	// Obj to json
	dataJson, _ := json.Marshal(people)
	err := WriteDataFile("people.json", dataJson)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestReadData(t *testing.T) {
	_, err := ReadDataFile("people.json")
	if err != nil {
		t.Errorf(err.Error())
	}
}
