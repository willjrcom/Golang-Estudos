package entity

type animal struct {
	ID          int
	name        string
	class       string
	isDangerous bool
	isDomestic  bool
}

func NewAnimal() *animal {
	return new(animal)
}
