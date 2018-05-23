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
