package gpkg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeometryColumn(t *testing.T) {
	g, err := NewGeopackage("/tmp/geometrycolumn_test.db")
	if err != nil {
		t.Error(err)
	}

	gc := GeometryColumn{
		TableNam:         "test_tablename",
		ColumnName:       "test_columnname",
		GeometryTypeName: "test_geometrytype",
		SrsId:            4326,
	}

	if err := g.AddGeometryColumn(gc); err != nil {
		assert.Error(t, err)
	}

	gc2 := g.FindGeometryColumn(gc.TableNam)
	assert.Equal(t, gc.TableNam, gc2.TableNam)
	assert.Equal(t, gc.ColumnName, gc2.ColumnName)
	assert.Equal(t, gc.GeometryTypeName, gc2.GeometryTypeName)
	assert.Equal(t, gc.SrsId, gc2.SrsId)

	gc2.GeometryTypeName = "POINT"
	if err = g.ModifyGeometryColumn(gc2); err != nil {
		assert.Error(t, err)
	}

	gc3 := g.FindGeometryColumn(gc2.TableNam)
	assert.Equal(t, gc2.GeometryTypeName, gc3.GeometryTypeName)
	if err := g.RemoveGeometryColumn(gc3.TableNam); err != nil {
		assert.Error(t, err)
	}

	gc4 := g.FindGeometryColumn(gc2.TableNam)
	assert.Equal(t, "", gc4.TableNam)

	g.Close()
	os.Remove("/tmp/geometrycolumn_test.db")
}
