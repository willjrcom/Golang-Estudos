package personRepositoryGorm

import (
	personEntity "projetoGo/entity/person"
)

func (pr *PersonRepositoryImpl) Delete(obj *personEntity.Person) error {
	return pr.Db.Delete(obj).Error
}

func (pr *PersonRepositoryImpl) DeleteBy(modelConditions *personEntity.Person) error {
	return pr.Db.Where(modelConditions).Select("address").Select("animals").Delete(&personEntity.Person{}).Error
}

func (pr *PersonRepositoryImpl) DeleteById(id uint) error {
	return pr.Db.Where("ID = ?", id).Select("address").Select("animals").Delete(&personEntity.Person{}).Error
}
