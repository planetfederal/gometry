package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolygonCreate(t *testing.T) {
	shell := [][3]float64{{0.0, 0.0, 0.0}, {0.0, 10.0, 0.0}, {10.0, 10.0, 0.0}, {10.0, 0.0, 0.0}}
	hole1 := [][3]float64{{2.0, 2.0, 0.0}, {2.0, 3.0, 0.0}, {3.0, 3.0, 0.0}, {3.0, 2.0, 0.0}}
	hole2 := [][3]float64{{4.0, 4.0, 0.0}, {4.0, 5.0, 0.0}, {5.0, 5.0, 0.0}, {5.0, 4.0, 0.0}}
	coords := [][][3]float64{shell, hole1, hole2}
	p := NewPolygonFromArray(coords)
	assert.Equal(t, 5, len(p.Shell.Coordinates))
	assert.Equal(t, 2, len(p.Holes))
	assert.Equal(t, [3]float64{0.0, 0.0, 0.0}, p.Shell.Coordinates[0].AsArray())
	assert.Equal(t, [3]float64{2.0, 2.0, 0.0}, p.Holes[0].Coordinates[0].AsArray())
}

func TestPolygonCreateFromBBOX(t *testing.T) {
	shell := [][3]float64{{0.0, 0.0, 0.0}, {0.0, 10.0, 0.0}, {10.0, 10.0, 0.0}, {10.0, 0.0, 0.0}}
	hole1 := [][3]float64{{2.0, 2.0, 0.0}, {2.0, 3.0, 0.0}, {3.0, 3.0, 0.0}, {3.0, 2.0, 0.0}}
	hole2 := [][3]float64{{4.0, 4.0, 0.0}, {4.0, 5.0, 0.0}, {5.0, 5.0, 0.0}, {5.0, 4.0, 0.0}}
	coords := [][][3]float64{shell, hole1, hole2}
	p := NewPolygonFromArray(coords)
	assert.Equal(t, 5, len(p.Shell.Coordinates))
	assert.Equal(t, 2, len(p.Holes))
	assert.Equal(t, [3]float64{0.0, 0.0, 0.0}, p.Shell.Coordinates[0].AsArray())
	assert.Equal(t, [3]float64{2.0, 2.0, 0.0}, p.Holes[0].Coordinates[0].AsArray())
}
