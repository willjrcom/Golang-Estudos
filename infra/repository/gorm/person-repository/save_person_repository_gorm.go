package personRepositoryGorm

import (
	personEntity "projetoGo/entity/person"
)

func (pr *PersonRepository) Save(person *personEntity.Person) error {
	err := pr.Db.Save(person.Address).Error

	if err != nil {
		return err
	}

	person.AddressID = person.Address.ID
	return pr.Db.Save(person).Error
}
