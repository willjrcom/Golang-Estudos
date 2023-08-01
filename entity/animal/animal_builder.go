package animalEntity

type AnimalBuilder struct {
	animal *Animal
}

func newAnimalBuilder(name, class string) *AnimalBuilder {
	return &AnimalBuilder{animal: &Animal{Name: name, Class: class}}
}

func (a *AnimalBuilder) WithIsDangerous(isDangerous bool) *AnimalBuilder {
	a.animal.IsDangerous = isDangerous
	return a
}

func (a *AnimalBuilder) WithIsDomestic(isDomestic bool) *AnimalBuilder {
	a.animal.IsDomestic = isDomestic
	return a
}

func (a *AnimalBuilder) Build() *Animal {
	return a.animal
}
