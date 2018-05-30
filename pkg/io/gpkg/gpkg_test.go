package gpkg

import "testing"

func TestGpkgInit(t *testing.T) {
	g, err := NewGeopackage("/tmp/gpkg_test.db")
	if err != nil {
		t.Error(err)
	}
	g.Close()
}
