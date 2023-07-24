package builder

type Director interface {
	Construct()
	SetBuilder(bp BuildProcess)
}

type directorImpl struct {
	builder BuildProcess
}

func (d *directorImpl) Construct() {
	d.builder.SetWheels().SetSeats().SetStructure()
}

func (d *directorImpl) SetBuilder(bp BuildProcess) {
	d.builder = bp
}

func NewDirector() Director {
	return &directorImpl{}
}
