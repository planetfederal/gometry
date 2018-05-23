package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type multiPoint struct {
	Typ         geoJSONType        `json:"type"`
	Coordinates []*geom.Coordinate `json:"coordinates"`
}

func marshalMultiPoint(p *geom.MultiPoint) ([]byte, error) {
	coords := make([]*geom.Coordinate, len(p.Points))
	for i, v := range p.Points {
		coords[i] = v.Coordinate
	}
	return json.Marshal(&multiPoint{multiPointType, coords})
}
