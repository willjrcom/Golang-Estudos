package personRepositoryGorm

import (
	"gorm.io/gorm"
)

type PersonRepository struct {
	Db *gorm.DB
}
