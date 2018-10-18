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
	"testing"

	"github.com/boundlessgeo/gometry/pkg/geom"
	"github.com/stretchr/testify/assert"
)

func TestPointJSON(t *testing.T) {
	p := geom.NewPointFromArray([3]float64{0.0, 0.0, 0.0})
	bs, _ := Marshal(p)
	assert.Equal(t, `{"type":"Point","coordinates":[0,0,0]}`, string(bs))
}

func TestPolygonJSON(t *testing.T) {
	p := geom.NewPolygonFromArray([][][3]float64{{{0, 0, 0}, {0, 1, 0}, {1, 1, 0}, {1, 0, 0}, {0, 0, 0}}})
	bs, _ := Marshal(p)
	assert.Equal(t, `{"type":"Polygon","coordinates":[[[0,0,0],[0,1,0],[1,1,0],[1,0,0],[0,0,0]]]}`, string(bs))
}

func TestLineStringJSON(t *testing.T) {
	l := geom.NewLineStringFromArray([][3]float64{{0, 0, 0}, {1, 1, 0}, {2, 2, 0}})
	bs, _ := Marshal(l)
	assert.Equal(t, `{"type":"LineString","coordinates":[[0,0,0],[1,1,0],[2,2,0]]}`, string(bs))
}

func TestMultiPointJSON(t *testing.T) {
	mp := geom.NewMultiPointFromArray([][3]float64{{1, 1, 0}, {2, 2, 3}})
	bs, _ := Marshal(mp)
	assert.Equal(t, `{"type":"MultiPoint","coordinates":[[1,1,0],[2,2,3]]}`, string(bs))
}

func TestMultiLineStringJSON(t *testing.T) {
	ml := geom.NewMultiLineString([][][3]float64{{{1, 1, 2}, {2, 2.3}, {3, 3, 3}}, {{4, 4, 5}, {5, 5, 6}, {7, 7, 7}}})
	bs, _ := Marshal(ml)
	assert.Equal(t, `{"type":"MultiLineString","coordinates":[[[1,1,2],[2,2.3,0],[3,3,3]],[[4,4,5],[5,5,6],[7,7,7]]]}`, string(bs))
}

func TestMultiPolygonJSON(t *testing.T) {
	p1 := [][][3]float64{{{0, 0, 0}, {0, 1, 0}, {1, 1, 0}, {1, 0, 0}, {0, 0, 0}}}
	p2 := [][][3]float64{{{0, 0, 0}, {0, 2, 0}, {2, 2, 0}, {2, 0, 0}, {0, 0, 0}}}

	mp := geom.NewMultiPolygonFromArray([][][][3]float64{p1, p2})
	bs, _ := Marshal(mp)
	assert.Equal(t, `{"type":"MultiPolygon","coordinates":[[[[0,0,0],[0,1,0],[1,1,0],[1,0,0],[0,0,0]]],[[[0,0,0],[0,2,0],[2,2,0],[2,0,0],[0,0,0]]]]}`, string(bs))
}

func TestFeatureJSON(t *testing.T) {
	pt := geom.NewPointFromArray([3]float64{0.0, 0.0, 0.0})
	pr := map[string]interface{}{"foo": 2}
	f := &Feature{ID: "1", Properties: pr, Geometry: pt}
	bs, _ := Marshal(f)
	assert.Equal(t, `{"type":"Feature","id":"1","geometry":{"type":"Point","coordinates":[0,0,0]},"properties":{"foo":2}}`, string(bs))
}

func TestFeatureCollectionJSON(t *testing.T) {
	pt := geom.NewPointFromArray([3]float64{0.0, 0.0, 0.0})
	pr := map[string]interface{}{"foo": 2}
	f := Feature{ID: "1", Properties: pr, Geometry: pt}
	pt2 := geom.NewPointFromArray([3]float64{0.0, 0.0, 0.0})
	pr2 := map[string]interface{}{"boo": 3}
	f2 := Feature{ID: "2", Properties: pr2, Geometry: pt2}
	fs := &FeatureCollection{[]*Feature{&f, &f2}}
	bs, _ := Marshal(fs)
	assert.Equal(t, `{"type":"FeatureCollection","features":[{"type":"Feature","id":"1","geometry":{"type":"Point","coordinates":[0,0,0]},"properties":{"foo":2}},{"type":"Feature","id":"2","geometry":{"type":"Point","coordinates":[0,0,0]},"properties":{"boo":3}}]}`, string(bs))
}
