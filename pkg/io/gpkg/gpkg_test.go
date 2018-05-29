package gpkg

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitGeopackage(t *testing.T) {
	g, err := NewGeopackage("./test.db")
	if err != nil {
		t.Error(err)
		return
	}
	defer g.db.Close()

	err = g.initializeSpatialMetadata()
	if err != nil {
		t.Error(err)
		return
	}

	err = g.validateSchema()
	if err != nil {
		t.Error(err)
		return
	}
}
