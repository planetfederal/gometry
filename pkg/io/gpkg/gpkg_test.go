package gpkg

import "testing"

var GeopackageTestFileName = "/tmp/gpkg_test.db"

func TestGpkgInit(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
	}
	g.Close()
}
