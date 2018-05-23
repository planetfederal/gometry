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
