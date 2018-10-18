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

//Envelope is the bounding box of a geometry
type Envelope struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func NewEnvelope(x1, x2, y1, y2 float64) *Envelope {
	var e Envelope
	if x1 < x2 {
		e.MinX = x1
		e.MaxX = x2
	} else {
		e.MinX = x2
		e.MaxX = x1
	}

	if y1 < y2 {
		e.MinY = y1
		e.MaxY = y2
	} else {
		e.MinY = y2
		e.MaxY = y1
	}

	return &e
}

func NewEnvelopeFromCoords(c1, c2 *Coordinate) *Envelope {
	return NewEnvelope(c1.X, c2.X, c1.Y, c2.Y)
}

func (e *Envelope) GetWidth() float64 {
	return e.MaxX - e.MinX
}

func (e *Envelope) GetHeight() float64 {
	return e.MaxY - e.MinY
}

func (e *Envelope) GetArea() float64 {
	return e.GetWidth() * e.GetHeight()
}

func (e *Envelope) ExpandByDelta(deltaX, deltaY float64) {
	e.MinX -= deltaX
	e.MaxX += deltaX
	e.MinY -= deltaY
	e.MaxY += deltaX
}

func (e *Envelope) ExpandByDistance(distance float64) {
	e.ExpandByDelta(distance, distance)
}

func (e *Envelope) ExpandToInclude(x, y float64) {
	if x < e.MinX {
		e.MinX = x
	}
	if x > e.MaxX {
		e.MaxX = x
	}
	if y < e.MinY {
		e.MinY = y
	}
	if y > e.MaxY {
		e.MaxY = y
	}
}

func (e *Envelope) ExpandToIncludeEnvelope(other *Envelope) {
	if other.MinX < e.MinX {
		e.MinX = other.MinX
	}
	if other.MaxX > e.MaxX {
		e.MaxX = other.MaxX
	}
	if other.MinY < e.MinY {
		e.MinY = other.MinY
	}
	if other.MaxY > e.MaxY {
		e.MaxY = other.MaxY
	}
}

func (e *Envelope) Translate(transX, transY float64) {
	e.MinX += transX
	e.MaxX += transX
	e.MinY += transY
	e.MaxY += transY
}

func (e *Envelope) Center() *Coordinate {
	cx := (e.MinX + e.MaxX) / 2.0
	cy := (e.MinY + e.MinY) / 2.0
	return &Coordinate{X: cx, Y: cy}
}

func (e *Envelope) Intersection(env *Envelope) *Envelope {
	var intMinX, intMinY, intMaxX, intMaxY float64
	if intMinX = env.MinX; e.MinX > env.MinX {
		intMinX = e.MinX
	}
	if intMinY = env.MinY; e.MinY > env.MinY {
		intMinY = e.MinY
	}
	if intMaxX = env.MaxX; e.MaxX < env.MaxX {
		intMaxX = e.MaxX
	}
	if intMaxY = env.MaxY; e.MaxY < env.MaxY {
		intMaxY = e.MaxY
	}
	return &Envelope{MinX: intMinX, MaxX: intMaxX, MinY: intMinY, MaxY: intMaxY}
}

func (e *Envelope) IntersectsEnvelope(other *Envelope) bool {
	return !(other.MinX > e.MaxX ||
		other.MaxX < e.MinX ||
		other.MinY > e.MaxY ||
		other.MaxY < e.MinY)
}

func (e *Envelope) IntersectsCoordinates(a, b *Coordinate) bool {
	var envMinY, envMaxY, envMinX, envMaxX float64
	if envMinX = b.X; a.X < b.X {
		envMinX = a.X
	}

	if envMinX > e.MaxX {
		return false
	}

	if envMaxX = b.X; a.X > b.X {
		envMaxX = a.X
	}

	if envMaxX < e.MinX {
		return false
	}

	if envMinY = b.Y; a.Y < b.Y {
		envMinY = a.Y
	}

	if envMinY > e.MaxY {
		return false
	}

	if envMaxY = b.Y; a.Y > b.Y {
		envMaxY = a.Y
	}

	if envMaxY < e.MinY {
		return false
	}
	return true
}

func (e *Envelope) Intersects(x, y float64) bool {
	return !(x > e.MaxX ||
		x < e.MinX ||
		y > e.MaxY ||
		y < e.MinY)
}

func (e *Envelope) OverlapsEnvelope(other *Envelope) bool {
	return e.IntersectsEnvelope(other)
}

func (e *Envelope) Overlaps(x, y float64) bool {
	return e.Intersects(x, y)
}

func (e *Envelope) ContainsEnvelope(other *Envelope) bool {
	return e.CoversEnvelope(other)
}

func (e *Envelope) ContainsCoordinate(p *Coordinate) bool {
	return e.CoversCoordinate(p)
}

func (e *Envelope) Contains(x, y float64) bool {
	return e.Covers(x, y)
}

func (e *Envelope) Covers(x, y float64) bool {
	return (x >= e.MinX && x <= e.MaxX &&
		y >= e.MinY && y <= e.MaxY)
}

func (e *Envelope) CoversCoordinate(p *Coordinate) bool {
	return e.Covers(p.X, p.Y)
}

func (e *Envelope) CoversEnvelope(other *Envelope) bool {
	return other.MinX >= e.MinX && other.MaxX <= e.MaxX &&
		other.MinY >= e.MinY && other.MaxY <= e.MaxY
}
