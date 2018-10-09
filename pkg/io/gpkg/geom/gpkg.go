package geom

import "github.com/boundlessgeo/gometry/pkg/io/wkb"

type GeoPackageBinaryHeader struct {
	magic    []byte
	version  byte
	flags    byte
	srsId    int32
	envelope []float64
}

type StandardGeoPackageBinary struct {
	header   GeoPackageBinaryHeader
	geometry wkb.WKBGeometry
}
