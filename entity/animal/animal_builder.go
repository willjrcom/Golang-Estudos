package animalEntity

type AnimalBuilder struct {
	animal *Animal
}

func newAnimalBuilder(name, class string) *AnimalBuilder {
	return &AnimalBuilder{animal: &Animal{name: name, class: class}}
}

func (a *AnimalBuilder) WithIsDangerous(isDangerous bool) *AnimalBuilder {
	a.animal.isDangerous = isDangerous
	return a
}

func (a *AnimalBuilder) WithIsDomestic(isDomestic bool) *AnimalBuilder {
	a.animal.isDomestic = isDomestic
	return a
}

func (a *AnimalBuilder) Build() *Animal {
	return a.animal
}
