package gpkg

import (
	"os"
	"testing"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

var GeopackageTestFileName = "/tmp/gpkg_test.db"

type TestTable struct {
	ID   int           `gorm:"PRIMARYKEY"`
	Geom geom.Geometry `gorm:"type:GEOMETRY"`
}

func (t TestTable) TableName() string {
	return "test_table"
}

func (t TestTable) GeomColumnName() string {
	return "geom"
}

func (t TestTable) SrsId() int32 {
	return 4326
}

func TestGpkgInit(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
	}
	tt := TestTable{}
	g.AddFeatureTable(tt)
	g.Close()
	os.Remove(GeopackageTestFileName)
}
