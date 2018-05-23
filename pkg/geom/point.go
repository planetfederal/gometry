package geom

type Point struct {
	Coordinate *Coordinate
}

func NewPointFromArray(coords [3]float64) *Point {
	c := NewCoordinate(coords)
	return &Point{c}
}

func (p Point) Envelope() *Envelope {
	return NewEnvelopeFromCoords(p.Coordinate, p.Coordinate)
}

func (p Point) Center() *Point {
	return &p
}

func (_ Point) GeometryType() string {
	return "Point"
}

func (_ Point) Area() float64 {
	return 0.0
}
