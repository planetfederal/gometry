package geom

type MultiPolygon struct {
	Polygons []*Polygon
}

func NewMultiPolygonFromArray(coords [][][][3]float64) *MultiPolygon {
	ps := make([]*Polygon, len(coords))
	for i, v := range coords {
		ps[i] = NewPolygonFromArray(v)
	}
	return &MultiPolygon{ps}
}
