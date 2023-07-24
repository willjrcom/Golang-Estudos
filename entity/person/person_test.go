package personEntity

import "testing"

var people = []Person{
	{
		name:     "William",
		birthday: "2000-01-26",
	},
	{
		name:     "Duda",
		birthday: "2005-10-24",
	},
	{
		name:     "Ana",
		birthday: "2013-10-20",
	},
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
