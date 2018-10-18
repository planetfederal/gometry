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
