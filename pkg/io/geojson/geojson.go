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
	"errors"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type geoJSONType string

const (
	pointType             geoJSONType = "Point"
	polygonType           geoJSONType = "Polygon"
	lineStringType        geoJSONType = "LineString"
	multiPointType        geoJSONType = "MultiPoint"
	multiPolygonType      geoJSONType = "MultiPolygon"
	multiLineStringType   geoJSONType = "MultiLineString"
	featureType           geoJSONType = "Feature"
	featureCollectionType geoJSONType = "FeatureCollection"
)

//Marshal will serialize geometries as geojson UTF-8 byte arrays
func Marshal(g interface{}) ([]byte, error) {
	switch g.(type) {
	case *geom.Point:
		return marshalPoint(g.(*geom.Point))
	case *geom.Polygon:
		return marshalPolygon(g.(*geom.Polygon))
	case *geom.LineString:
		return marshalLineString(g.(*geom.LineString))
	case *geom.MultiPoint:
		return marshalMultiPoint(g.(*geom.MultiPoint))
	case *geom.MultiPolygon:
		return marshalMultiPolygon(g.(*geom.MultiPolygon))
	case *geom.MultiLineString:
		return marshalMultiLineString(g.(*geom.MultiLineString))
	case *Feature:
		return marshalFeature(g.(*Feature))
	case *FeatureCollection:
		return marshalFeatureCollection(g.(*FeatureCollection))
	default:
		return nil, errors.New("Type not supported")
	}
}
