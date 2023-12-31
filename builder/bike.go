package builder

type BikeBuilder struct {
	v VehicleProduct
}

func NewBikeBuilder() BuildProcess {
	return &BikeBuilder{}
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}
