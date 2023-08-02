package personService

import (
	"net/url"
	personEntity "projetoGo/entity/person"
	repositoryInterface "projetoGo/entity/repository"
	"projetoGo/infra/presenter"
)

type Service struct {
	Repository repositoryInterface.Repository[personEntity.Person]
}

func NewService(repository repositoryInterface.Repository[personEntity.Person]) *Service {
	return &Service{Repository: repository}
}

func (s *Service) FindAll() ([]*personEntity.Person, error) {
	return s.Repository.FindAll()
}

func (s *Service) RegisterPerson(values url.Values) error {
	person, err := presenter.UrlValuesToPerson(values)

	if err != nil {
		return err
	}

	return s.Repository.Save(person)
}

func (s *Service) FindBy(values url.Values) ([]*personEntity.Person, error) {
	person, err := presenter.UrlValuesToPerson(values)

	if err != nil {
		return []*personEntity.Person{}, err
	}

	return s.Repository.FindBy(person)
}

func (s *Service) DeletePerson(values url.Values) error {
	person, err := presenter.UrlValuesToPerson(values)

	if err != nil {
		return err
	}

	return s.Repository.DeleteBy(person)
}
