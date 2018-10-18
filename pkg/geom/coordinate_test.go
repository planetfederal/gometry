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
