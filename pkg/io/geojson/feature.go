/*
MIT License

Copyright (c) 2018 Boundless

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
