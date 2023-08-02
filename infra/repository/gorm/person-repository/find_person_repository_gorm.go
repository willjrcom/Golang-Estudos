package personRepositoryGorm

import (
	personEntity "projetoGo/entity/person"
)

func (pr *PersonRepositoryImpl) FindAll() ([]*personEntity.Person, error) {
	people := []*personEntity.Person{}
	err := pr.Db.Find(&people).Error

	if err != nil {
		return []*personEntity.Person{}, err
	}
	return people, nil
}

func (pr *PersonRepositoryImpl) FindById(id uint) (*personEntity.Person, error) {
	person := personEntity.Person{}
	err := pr.Db.Where("id = ?", id).Preload("Address").Preload("Animals").First(&person).Error

	if err != nil {
		return &personEntity.Person{}, err
	}

	return &person, nil
}

func (pr *PersonRepositoryImpl) FindByName(name string) ([]*personEntity.Person, error) {
	people := []*personEntity.Person{}
	err := pr.Db.Where("name = ?", name).Preload("Address").Preload("Animals").Find(&people).Error

	if err != nil {
		return []*personEntity.Person{}, err
	}

	return people, nil
}

func (pr *PersonRepositoryImpl) FindBy(modelConditions *personEntity.Person) ([]*personEntity.Person, error) {
	people := []*personEntity.Person{}

	err := pr.Db.Where(modelConditions).Preload("Address").Preload("Animals").Find(&people).Error

	if err != nil {
		return []*personEntity.Person{}, err
	}

	return people, nil
}
