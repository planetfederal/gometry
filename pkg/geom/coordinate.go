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
	"encoding/json"
	"math"
)

var eps = 0.00000001

//Coordinate is a 3 dimensional point in space
type Coordinate struct {
	X, Y, Z float64
}

func New2dCoordinate(a [2]float64) *Coordinate {
	return &Coordinate{a[0], a[1], 0}
}

//NewCoordinate takes a 1 dimensional array of [X,Y] or [X,Y,Z]
func NewCoordinate(a [3]float64) *Coordinate {
	return &Coordinate{a[0], a[1], a[2]}
}

//Equals2d is a comparison of the X,Y values of two *Coordinates
func (c *Coordinate) Equals2d(c2 *Coordinate) bool {
	if math.Abs(c.X-c2.X) < eps &&
		math.Abs(c.Y-c2.Y) < eps {
		return true
	}
	return false
}

func (c *Coordinate) Equals3d(c2 *Coordinate) bool {
	if c.Equals2d(c2) &&
		math.Abs(c.Z-c2.Z) < eps {
		return true
	}
	return false
}

func (c *Coordinate) EuclideanDistance(c2 *Coordinate) float64 {
	dx := c.X - c2.X
	dy := c.Y - c2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (c *Coordinate) EuclideanDistance3d(c2 *Coordinate) float64 {
	dx := c.X - c2.X
	dy := c.Y - c2.Y
	dz := c.Z - c2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (c *Coordinate) AsArray() [3]float64 {
	return [3]float64{c.X, c.Y, c.Z}
}

func (c Coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal((&c).AsArray())
}
