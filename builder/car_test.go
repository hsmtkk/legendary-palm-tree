package builder_test

import (
	"testing"

	"github.com/hsmtkk/legendary-palm-tree/builder"
	"github.com/stretchr/testify/assert"
)

func TestCar(t *testing.T) {
	director := builder.NewDirector()
	carBuilder := builder.NewCarBuilder()
	director.SetBuilder(carBuilder)
	director.Construct()
	car := carBuilder.Build()
	assert.Equal(t, 4, car.Wheels)
	assert.Equal(t, "Car", car.Structure)
	assert.Equal(t, 5, car.Seats)
}
