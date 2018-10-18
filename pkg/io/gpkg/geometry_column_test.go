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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeometryColumn(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
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
}
