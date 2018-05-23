package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type multiPolygon struct {
	Typ        geoJSONType            `json:"type"`
	Coordinate [][][]*geom.Coordinate `json:"coordinates"`
}

func marshalMultiPolygon(p *geom.MultiPolygon) ([]byte, error) {
	coords := make([][][]*geom.Coordinate, len(p.Polygons))
	for i, v := range p.Polygons {
		coords[i] = make([][]*geom.Coordinate, 1+len(v.Holes))
		coords[i][0] = v.Shell.Coordinates
		for j, h := range v.Holes {
			coords[i][j+1] = h.Coordinates
		}
	}
	return json.Marshal(&multiPolygon{multiPolygonType, coords})
}
