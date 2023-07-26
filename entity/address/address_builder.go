package addressEntity

type AddressBuilder struct {
	address *Address
}

func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{address: new(Address)}
}

func (b *AddressBuilder) WithStreet(street string) *AddressBuilder {
	b.address.street = street
	return b
}

func (b *AddressBuilder) WithNumber(number int) *AddressBuilder {
	b.address.number = number
	return b
}

func (b *AddressBuilder) WithCity(city string) *AddressBuilder {
	b.address.city = city
	return b
}
func (b *AddressBuilder) WithState(state string) *AddressBuilder {
	b.address.state = state
	return b
}
func (b *AddressBuilder) WithCountry(country string) *AddressBuilder {
	b.address.country = country
	return b
}

func (b *AddressBuilder) Build() *Address {
	return b.address
}
