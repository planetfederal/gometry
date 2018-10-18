/*
MIT License

Copyright (c) 2018 Boundless

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
