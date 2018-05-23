package geom

//Geometry PlaceHolder
type Geometry interface {
	Envelope() *Envelope
	Center() *Point
	GeometryType() string
	Area() float64
}
