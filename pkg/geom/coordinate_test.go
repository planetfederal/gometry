package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	c := &Coordinate{1, 1, 1}
	c2 := &Coordinate{1, 1, 1}
	assert.Equal(t, c.Equals2d(c2), true, "Should Equal2d")
	assert.Equal(t, c.Equals3d(c2), true, "Should Equal3d")
}

func TestDistance(t *testing.T) {
	c := &Coordinate{1, 1, 1}
	c2 := &Coordinate{1, 1, 1}
	assert.Equal(t, c.EuclideanDistance(c2), float64(0.0), "Distance 2d should be 0")
	assert.Equal(t, c.EuclideanDistance3d(c2), float64(0.0), "Distance 3d should be 0")

	c3 := &Coordinate{0, 0, 0}
	c4 := &Coordinate{3, 4, 0}

	assert.Equal(t, c3.EuclideanDistance(c4), float64(5.0), "Distance should be 5")
	assert.Equal(t, c3.EuclideanDistance3d(c4), float64(5.0), "Distance should be 5 in 3d")

	c5 := &Coordinate{0, 0, 0}
	c6 := &Coordinate{0.3, 0.4, 0}

	assert.Equal(t, c5.EuclideanDistance(c6), float64(0.5), "Distance should be 5")
	assert.Equal(t, c5.EuclideanDistance3d(c6), float64(0.5), "Distance should be 5 in 3d")
}
