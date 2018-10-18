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

type MultiPoint struct {
	Points []*Point
}

func NewMultiPointFromArray(coords [][3]float64) *MultiPoint {
	cs := make([]*Point, len(coords))
	for i, v := range coords {
		cs[i] = NewPointFromArray(v)
	}
	return &MultiPoint{cs}
}

func (p MultiPoint) Envelope() *Envelope {
	e := &Envelope{}

	for _, v := range p.Points {
		e.ExpandToInclude(v.Coordinate.X, v.Coordinate.Y)
	}

	return e
}

func (p MultiPoint) Center() *Point {
	x, y := 0.0, 0.0
	for _, v := range p.Points {
		x += v.Coordinate.X
		y += v.Coordinate.Y
	}
	cnt := float64(len(p.Points))
	return NewPointFromArray([3]float64{x / cnt, y / cnt, 0.0})
}

func (_ MultiPoint) GeometryType() string {
	return "MultiPoint"
}

func (_ MultiPoint) Area() float64 {
	return 0.0
}
