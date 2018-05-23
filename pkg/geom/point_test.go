package geom

import "testing"
import "github.com/stretchr/testify/assert"

func TestPoint(t *testing.T) {
	p := NewPointFromArray([3]float64{1.0, 2.0, 0.0})
	assert.Equal(t, p.Coordinate.X, 1.0)
	assert.Equal(t, p.Coordinate.Y, 2.0)
}
