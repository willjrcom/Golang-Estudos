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
	tx := pr.Db.Where("id = ?", id).First(&person)

	if tx != nil {
		return &personEntity.Person{}, tx.Error
	}

	return &person, nil
}

func (pr *PersonRepository) FindByName(name string) (*personEntity.Person, error) {
	person := personEntity.Person{}
	tx := pr.Db.Where("name = ?", name).First(&person)

	if tx != nil {
		return &personEntity.Person{}, tx.Error
	}

	return &person, nil
}

func (pr *PersonRepository) FindBy(conditions string, values []interface{}) (*personEntity.Person, error) {
	person := personEntity.Person{}
	tx := pr.Db.Where(conditions, values...).First(&person)

	if tx != nil {
		return &personEntity.Person{}, tx.Error
	}

	return &person, nil
}
