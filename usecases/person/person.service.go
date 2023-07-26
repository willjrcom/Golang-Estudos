package personUsecases

import (
	"fmt"
	addressEntity "projetoGo/entity/address"
	personEntity "projetoGo/entity/person"
)

func RegisterPerson() {
	addressBuilder := addressEntity.AddressBuilder{}
	address := addressBuilder.WithStreet("Rua Piedade").WithNumber(226).WithCity("Sorocaba").WithState("SP").WithCountry("BR").Build()
	person, err := personEntity.NewPersonBuilder("William", "436.377.998-55").WithAddress(address).Build()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)
}
