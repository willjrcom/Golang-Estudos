package animalEntity

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	Name        string
	Class       string
	IsDangerous bool
	IsDomestic  bool
	PersonID    uint
}
