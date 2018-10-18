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

type LineString struct {
	Coordinates []*Coordinate
}

func NewLineStringFromArray(pts [][3]float64) *LineString {
	coords := make([]*Coordinate, len(pts))
	for i, v := range pts {
		coords[i] = NewCoordinate(v)
	}
	return &LineString{coords}
}

func (l LineString) Envelope() *Envelope {
	e := &Envelope{}
	for _, c := range l.Coordinates {
		e.ExpandToInclude(c.X, c.Y)
	}
	return e
}

func (l LineString) Center() *Point {
	x, y := 0.0, 0.0
	for _, c := range l.Coordinates {
		x += c.X
		y += c.Y
	}
	//Subtract one because the start and stop vertex it in there twice
	cnt := float64(len(l.Coordinates))
	return NewPointFromArray([3]float64{x / cnt, y / cnt, 0.0})
}

func (_ LineString) GeometryType() string {
	return "LineString"
}

func (_ LineString) Area() float64 {
	return 0.0
}
