package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type multiLineString struct {
	Typ        geoJSONType          `json:"type"`
	Coordinate [][]*geom.Coordinate `json:"coordinates"`
}

func marshalMultiLineString(l *geom.MultiLineString) ([]byte, error) {
	coords := make([][]*geom.Coordinate, len(l.LineStrings))
	for i, v := range l.LineStrings {
		coords[i] = v.Coordinates
	}
	return json.Marshal(&multiLineString{multiLineStringType, coords})
}
