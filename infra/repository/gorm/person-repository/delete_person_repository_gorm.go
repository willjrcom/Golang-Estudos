package personRepositoryGorm

import (
	personEntity "projetoGo/entity/person"
)

func (pr *PersonRepository) Delete(obj *personEntity.Person) error {
	return pr.Db.Delete(&personEntity.Person{}).Error
}

func (pr *PersonRepository) DeleteBy(conditions string, values []interface{}) error {
	return pr.Db.Where(conditions, values...).Delete(&personEntity.Person{}).Error
}

func (pr *PersonRepository) DeleteById(id uint) error {
	return pr.Db.Where("ID = ?", id).Delete(&personEntity.Person{}).Error
}
