package addressEntity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street  string
	Number  int
	City    string
	State   string
	Country string
}
