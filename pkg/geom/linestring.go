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
