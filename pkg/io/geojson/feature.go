package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type Feature struct {
	ID         string `json:"id"`
	Geometry   geom.Geometry
	Properties map[string]interface{}
}

type FeatureCollection struct {
	Features []*Feature
}

func marshalFeature(f *Feature) ([]byte, error) {
	gj := []byte(`{"type":"Feature","id":"`)
	gj = append(gj, f.ID...)
	gj = append(gj, `","geometry":`...)
	geom, _ := Marshal(f.Geometry)
	gj = append(gj, geom...)
	gj = append(gj, `,"properties":`...)
	prop, _ := json.Marshal(f.Properties)
	gj = append(gj, prop...)
	gj = append(gj, `}`...)
	return gj, nil
}

func marshalFeatureCollection(f *FeatureCollection) ([]byte, error) {
	gj := []byte(`{"type":"FeatureCollection",`)
	gj = append(gj, `"features":[`...)
	for i, feat := range f.Features {
		if i != 0 {
			gj = append(gj, `,`...)
		}
		fb, _ := marshalFeature(feat)
		gj = append(gj, fb...)
	}
	gj = append(gj, `]}`...)
	return gj, nil
}
