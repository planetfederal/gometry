package gpkg

import "testing"

func TestTileMatrix(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
	}
	g.Close()
}

func TestTileMatrixSet(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
	}
	g.Close()
}
