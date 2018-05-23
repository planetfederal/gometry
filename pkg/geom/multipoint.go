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
