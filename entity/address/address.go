package addressEntity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	street  string
	number  int
	city    string
	state   string
	country string
}
