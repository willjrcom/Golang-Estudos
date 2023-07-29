package personUsecases

import (
	personEntity "projetoGo/entity/person"
	repositoryInterface "projetoGo/entity/repository"
)

type Service struct {
	Repository repositoryInterface.Repository[personEntity.Person]
}

func (s *Service) FindAll() ([]*personEntity.Person, error) {
	return s.Repository.FindAll()
}

func (s *Service) FindBy() ([]*personEntity.Person, error) {
	people, err := s.Repository.FindAll()
	return people, err
}

func (s *Service) RegisterPerson(newPerson *personEntity.Person) {
	s.Repository.Save(newPerson)
	// addressBuilder := addressEntity.AddressBuilder{}
	// address := addressBuilder.WithStreet("Rua Piedade").WithNumber(226).WithCity("Sorocaba").WithState("SP").WithCountry("BR").Build()
	// person, err := personEntity.NewPersonBuilder("William", "436.377.998-55").WithAddress(address).Build()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func (s *Service) DeletePerson(params string) {
	// !!! Criar regra do where
	s.Repository.DeleteBy(params, []interface{}{"william"})
}
