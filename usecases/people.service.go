package usecases

import (
	"fmt"
	addressEntity "projetoGo/entity/address"
	peopleEntity "projetoGo/entity/people"
)

func RegisterPeople() {
	addressBuilder := addressEntity.AddressBuilder{}
	address := addressBuilder.WithStreet("Rua Piedade").WithNumber(226).WithCity("Sorocaba").WithState("SP").WithCountry("BR").Build()
	people := peopleEntity.NewPeopleBuilder("William", "436.377.998-55").WithAddress(address).Build()

	fmt.Println(people)
}
