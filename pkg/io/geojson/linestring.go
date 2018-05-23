package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type lineString struct {
	Typ        geoJSONType        `json:"type"`
	Coordinate []*geom.Coordinate `json:"coordinates"`
}

func marshalLineString(l *geom.LineString) ([]byte, error) {
	return json.Marshal(&lineString{lineStringType, l.Coordinates})
}
