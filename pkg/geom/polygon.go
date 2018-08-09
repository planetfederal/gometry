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
