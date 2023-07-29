package animalEntity

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	name        string
	class       string
	isDangerous bool
	isDomestic  bool
}
