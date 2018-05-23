package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineString(t *testing.T) {
	coords := [][3]float64{{0.0, 0.0, 0.0}, {1.0, 1.0, 1.0}, {2.0, 2.0, 2.0}}
	ls := NewLineStringFromArray(coords)
	for i, c := range ls.Coordinates {
		assert.Equal(t, c.X, coords[i][0])
		assert.Equal(t, c.Y, coords[i][0])
		assert.Equal(t, c.Z, coords[i][0])
	}
}
