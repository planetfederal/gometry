package gpkg

import "testing"

func TestContents(t *testing.T) {
	g, err := NewGeopackage("/tmp/contents.db")
	if err != nil {
		t.Error(err)
	}

	g.Close()
}
