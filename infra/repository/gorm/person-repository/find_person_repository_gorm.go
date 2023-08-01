package personRepositoryGorm

import (
	personEntity "projetoGo/entity/person"
)

func (pr *PersonRepository) FindAll() ([]*personEntity.Person, error) {
	people := []*personEntity.Person{}
	err := pr.Db.Find(&people).Error

	if err != nil {
		return []*personEntity.Person{}, err
	}
	return people, nil
}

func (pr *PersonRepository) FindById(id uint) (*personEntity.Person, error) {
	person := personEntity.Person{}
	tx := pr.Db.Where("id = ?", id).Preload("Address").Preload("Animals").First(&person)

	if tx != nil {
		return &personEntity.Person{}, tx.Error
	}

	return &person, nil
}

func (pr *PersonRepository) FindByName(name string) (*personEntity.Person, error) {
	person := personEntity.Person{}
	tx := pr.Db.Where("name = ?", name).Preload("Address").Preload("Animals").First(&person)

	if tx != nil {
		return &personEntity.Person{}, tx.Error
	}

	return &person, nil
}

func (pr *PersonRepository) FindBy(modelConditions *personEntity.Person) ([]*personEntity.Person, error) {
	people := []*personEntity.Person{}

	err := pr.Db.Where(modelConditions).Preload("Address").Preload("Animals").Find(&people).Error

	if err != nil {
		return []*personEntity.Person{}, err
	}

	return people, nil
}
