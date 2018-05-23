package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type polygon struct {
	Typ        geoJSONType          `json:"type"`
	Coordinate [][]*geom.Coordinate `json:"coordinates"`
}

func marshalPolygon(p *geom.Polygon) ([]byte, error) {
	coordss := make([][]*geom.Coordinate, 1)
	coordss[0] = p.Shell.Coordinates
	return json.Marshal(&polygon{polygonType, coordss})
}
