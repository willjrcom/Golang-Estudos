package personUsecases

import (
	"net/url"
	personEntity "projetoGo/entity/person"
	repositoryInterface "projetoGo/entity/repository"
	"projetoGo/infra/presenter"
	utils "projetoGo/infra/utils/treino"
)

type Service struct {
	Repository repositoryInterface.Repository[personEntity.Person]
}

func (s *Service) RegisterPerson(values url.Values) error {
	person, err := presenter.UrlValuesToPerson(values)

	if err != nil {
		return err
	}

	return s.Repository.Save(person)
}

func (s *Service) FindAll() ([]*personEntity.Person, error) {
	return s.Repository.FindAll()
}

func (s *Service) FindBy(values url.Values) ([]*personEntity.Person, error) {
	mapValues := utils.UrlValuesToMap(values)
	conditions := utils.MapToGormConditions(mapValues)
	args := utils.MapKeyToArray(mapValues)
	return s.Repository.FindBy(conditions, args)
}

func (s *Service) DeletePerson(values url.Values) error {
	mapValues := utils.UrlValuesToMap(values)
	conditions := utils.MapToGormConditions(mapValues)
	args := utils.MapKeyToArray(mapValues)
	return s.Repository.DeleteBy(conditions, args)
}
