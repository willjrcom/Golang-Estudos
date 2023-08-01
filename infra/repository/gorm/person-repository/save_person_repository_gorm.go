package personRepositoryGorm

import (
	personEntity "projetoGo/entity/person"
)

func (pr *PersonRepository) Save(person *personEntity.Person) error {
	return pr.Db.Save(*person).Error
}
