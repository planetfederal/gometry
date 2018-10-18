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

type LinearRing struct {
	Coordinates []*Coordinate
}

func NewLinearRing(pts [][3]float64) *LinearRing {
	if len(pts) < 3 {
		return nil
	}
	coords := make([]*Coordinate, len(pts))
	for i, v := range pts {
		coords[i] = NewCoordinate(v)
	}
	first := coords[0]
	last := coords[len(coords)-1]
	if first.X != last.X || first.Y != last.Y {
		coords = append(coords, first)
	}
	return &LinearRing{coords}
}

func (l *LinearRing) AsArray() [][3]float64 {
	arr := make([][3]float64, len(l.Coordinates))
	for i, v := range l.Coordinates {
		arr[i] = v.AsArray()
	}
	return arr
}
