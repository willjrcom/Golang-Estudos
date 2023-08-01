package personEntity

import (
	addressEntity "projetoGo/entity/address"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var people = []Person{
	{
		Name:     "William",
		Birthday: time.Date(2000, time.January, 26, 0, 0, 0, 0, time.UTC),
	},
	{
		Name:     "Duda",
		Birthday: time.Date(2005, time.October, 24, 0, 0, 0, 0, time.UTC),
	},
	{
		Name:     "Ana",
		Birthday: time.Date(2013, time.October, 20, 0, 0, 0, 0, time.UTC),
	},
}

func TestNameValid(t *testing.T) {
	assert := assert.New(t)
	address := addressEntity.NewAddressBuilder().WithStreet("Rua Piedade").WithNumber(226).WithCity("Sorocaba").WithState("SP").WithCountry("BR").Build()
	person, _ := NewPersonBuilder("William", "436.377.998-55").WithAddress(address).BuildDto()
	assert.Equal("William", person.Name)
}

func TestIsChild(t *testing.T) {
	assert := assert.New(t)
	isChild, _ := people[0].isChild()
	assert.Equal(false, isChild)
}

func TestIsAdult(t *testing.T) {
	assert := assert.New(t)
	isAdult, _ := people[1].isAdult()
	assert.Equal(false, isAdult)
}

func TestIsTeen(t *testing.T) {
	assert := assert.New(t)
	isTeen, _ := people[2].isTeen()
	assert.Equal(false, isTeen)
}
