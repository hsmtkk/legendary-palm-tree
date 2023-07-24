package builder_test

import (
	"testing"

	"github.com/hsmtkk/legendary-palm-tree/builder"
	"github.com/stretchr/testify/assert"
)

func TestBike(t *testing.T) {
	director := builder.NewDirector()
	bikeBuilder := builder.NewBikeBuilder()
	director.SetBuilder(bikeBuilder)
	director.Construct()
	bike := bikeBuilder.Build()
	assert.Equal(t, 2, bike.Wheels)
	assert.Equal(t, "Bike", bike.Structure)
	assert.Equal(t, 1, bike.Seats)
}
