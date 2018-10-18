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

type Polygon struct {
	Shell *LinearRing
	Holes []*LinearRing
}

func NewPolygonFromArray(coords [][][3]float64) *Polygon {
	shell := NewLinearRing(coords[0])
	holes := make([]*LinearRing, len(coords[1:]))
	for i, v := range coords[1:] {
		holes[i] = NewLinearRing(v)
	}
	return &Polygon{shell, holes}
}

func NewPolygonFromBBOX(minLat, minLon, maxLat, maxLon float64) *Polygon {
	return NewPolygonFromArray([][][3]float64{{{minLon, minLat, 0.0}, {minLon, maxLat, 0.0}, {maxLon, maxLat, 0.0}, {maxLon, minLat, 0.0}}})
}

func (p Polygon) Envelope() *Envelope {
	e := &Envelope{}
	for _, c := range p.Shell.Coordinates {
		e.ExpandToInclude(c.X, c.Y)
	}
	for _, h := range p.Holes {
		for _, c := range h.Coordinates {
			e.ExpandToInclude(c.X, c.Y)
		}
	}
	return e
}

func (p Polygon) Center() *Point {
	x, y := 0.0, 0.0
	for _, c := range p.Shell.Coordinates {
		x += c.X
		y += c.Y
	}
	//Subtract one because the start and stop vertex it in there twice
	cnt := float64(len(p.Shell.Coordinates) - 1)
	return NewPointFromArray([3]float64{x / cnt, y / cnt, 0.0})
}

func (_ Polygon) GeometryType() string {
	return "Polygon"
}

func (_ Polygon) Area() float64 {
	return 0.0
}
