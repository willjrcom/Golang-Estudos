package personEntity

import (
	"fmt"
	addressEntity "projetoGo/entity/address"
	"testing"
	"time"
)

var people = []Person{
	{
		name:     "William",
		birthday: time.Date(2000, time.January, 26, 0, 0, 0, 0, time.UTC),
	},
	{
		name:     "Duda",
		birthday: time.Date(2005, time.October, 24, 0, 0, 0, 0, time.UTC),
	},
	{
		name:     "Ana",
		birthday: time.Date(2013, time.October, 20, 0, 0, 0, 0, time.UTC),
	},
}

func TestNameValid(t *testing.T) {
	address := addressEntity.NewAddressBuilder().WithStreet("Rua Piedade").WithNumber(226).WithCity("Sorocaba").WithState("SP").WithCountry("BR").Build()
	person, err := NewPersonBuilder("William", "436.377.998-55").WithAddress(address).Build()

	fmt.Println(person)

	if err != nil {
		fmt.Println(err)
		t.Errorf(err[0].Error())
	}
}

func TestIsChild(t *testing.T) {
	if isChild, _ := people[0].isChild(); isChild {
		t.Errorf("%v, isChild() = %v, want %v", people[0].name, isChild, false)
	}
}

func TestIsAdult(t *testing.T) {
	if isAdult, _ := people[1].isAdult(); isAdult {
		t.Errorf("%v, isAdult() = %v, want %v", people[1].name, isAdult, false)
	}
}

func TestIsTeen(t *testing.T) {
	if isTeen, _ := people[2].isTeen(); isTeen {
		t.Errorf("%v, isTeen() = %v, want %v", people[2].name, isTeen, false)
	}
}
