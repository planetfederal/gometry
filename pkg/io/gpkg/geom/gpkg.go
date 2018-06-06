package geom

type GeoPackageBinaryHeader struct {
	magic    []byte
	version  byte
	flags    byte
	srsId    int32
	envelope []float64
}

type StandardGeoPackageBinary struct {
	header   GeoPackageBinaryHeader
	geometry WKBGeometry
}
