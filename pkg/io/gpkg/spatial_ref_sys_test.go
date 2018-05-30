package gpkg

import "testing"

func TestSpatialRefSys(t *testing.T) {
	g, err := NewGeopackage("/tmp/srs_test.db")
	if err != nil {
		t.Error(err)
	}
	g.Close()
}
