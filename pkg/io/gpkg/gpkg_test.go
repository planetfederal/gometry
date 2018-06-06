package gpkg

import (
	"os"
	"testing"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

var GeopackageTestFileName = "/tmp/gpkg_test.db"

type TestFeatureTable struct {
	ID   int           `gorm:"PRIMARYKEY"`
	Geom geom.Geometry `gorm:"type:GEOMETRY"`
}

func (t TestFeatureTable) TableName() string {
	return "test_table"
}

func (t TestFeatureTable) GeomColumnName() string {
	return "geom"
}

func (t TestFeatureTable) SrsId() int32 {
	return 4326
}

type TestRasterTable struct {
	ID   int           `gorm:"PRIMARYKEY"`
	Geom geom.Geometry `gorm:"type:GEOMETRY"`
}

func (t TestRasterTable) TableName() string {
	return "test_raster_table"
}

func (t TestRasterTable) GeomColumnName() string {
	return "geom"
}

func (t TestRasterTable) SrsId() int32 {
	return 4326
}

func TestGpkgInit(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
	}
	tt := TestFeatureTable{}
	g.AddFeatureTable(tt)

	tr := TestRasterTable{}
	g.AddTileTable(tr)

	g.Close()
	os.Remove(GeopackageTestFileName)
}
