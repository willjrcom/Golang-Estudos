package database

import (
	addressEntity "projetoGo/entity/address"
	animalEntity "projetoGo/entity/animal"
	personEntity "projetoGo/entity/person"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	// Executar migração para criar a tabela de pessoas
	err = db.AutoMigrate(&personEntity.Person{}, &addressEntity.Address{}, &animalEntity.Animal{})
	if err != nil {
		panic("Falha ao executar migração: " + err.Error())
	}

	// people := []*personEntity.Person{}
	// db.Find(&people)

	// if len(people) != 0 {
	// 	for _, person := range people {
	// 		fmt.Println(person)
	// 	}
	// }

	return db
}
