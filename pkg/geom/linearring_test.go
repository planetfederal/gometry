package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearRing(t *testing.T) {
	coords := [][3]float64{{0.0, 0.0, 0.0}, {0.0, 1.0, 0.0}, {1.0, 1.0, 0.0}}
	lr := NewLinearRing(coords)
	//Should close the ring and make a non matching first and last point append first to
	//the end an match it.
	assert.Equal(t, lr.Coordinates[0].Equals3d(lr.Coordinates[3]), true)
}
