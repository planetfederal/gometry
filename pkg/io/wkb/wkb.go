package wkb

type WKBGeometry struct {
}

const (
	typePoint           uint32 = 1
	typeLineString      uint32 = 2
	typePolygon         uint32 = 3
	typeMultiPoint      uint32 = 4
	typeMultiLineString uint32 = 5
	typeMultiPolygon    uint32 = 6
	typeCollection      uint32 = 7
)
