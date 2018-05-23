package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type point struct {
	Typ        geoJSONType      `json:"type"`
	Coordinate *geom.Coordinate `json:"coordinates"`
}

func marshalPoint(p *geom.Point) ([]byte, error) {
	return json.Marshal(&point{pointType, p.Coordinate})
}
