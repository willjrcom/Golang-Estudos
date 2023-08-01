package addressEntity

type AddressBuilder struct {
	address *Address
}

func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{address: new(Address)}
}

func (b *AddressBuilder) WithStreet(street string) *AddressBuilder {
	b.address.Street = street
	return b
}

func (b *AddressBuilder) WithNumber(number int) *AddressBuilder {
	b.address.Number = number
	return b
}

func (b *AddressBuilder) WithCity(city string) *AddressBuilder {
	b.address.City = city
	return b
}
func (b *AddressBuilder) WithState(state string) *AddressBuilder {
	b.address.State = state
	return b
}
func (b *AddressBuilder) WithCountry(country string) *AddressBuilder {
	b.address.Country = country
	return b
}

func (b *AddressBuilder) Build() *Address {
	return b.address
}
