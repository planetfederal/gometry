package geom

type MultiLineString struct {
	LineStrings []*LineString
}

func NewMultiLineString(coords [][][3]float64) *MultiLineString {
	ls := make([]*LineString, len(coords))
	for i, v := range coords {
		ls[i] = NewLineStringFromArray(v)
	}
	return &MultiLineString{ls}
}
