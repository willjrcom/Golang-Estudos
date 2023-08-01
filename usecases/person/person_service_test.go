package personService

import (
	"net/url"
	"projetoGo/infra/database"
	personRepositoryGorm "projetoGo/infra/repository/gorm/person-repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterPerson(t *testing.T) {
	assert := assert.New(t)

	var Service = Service{
		Repository: &personRepositoryGorm.PersonRepository{
			Db: database.NewDb(),
		},
	}
	values := url.Values{
		"name":     {"William TestRegisterPerson"},
		"cpf":      {"436.377.998-55"},
		"birthday": {"2010-10-10"},
		"genre":    {"Masc"},
	}

	err := Service.RegisterPerson(values)
	assert.Equal(nil, err)
}
