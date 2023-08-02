package personRepositoryGorm

import (
	"gorm.io/gorm"
)

type PersonRepositoryImpl struct {
	Db *gorm.DB
}

func NewPersonRepositoryImpl(db *gorm.DB) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{Db: db}
}
