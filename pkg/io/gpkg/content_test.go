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
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContents(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
	}

	c := Content{
		TableNam:    "test_tablename",
		DataType:    "test_datatype",
		Identifier:  "test_identifier",
		Description: "test_description",
		LastChange:  time.Now(),
		MinX:        -180.0,
		MaxX:        180.0,
		MinY:        -90.0,
		MaxY:        90.0,
		SrsId:       4326,
	}

	if err := g.AddContent(c); err != nil {
		assert.Error(t, err)
	}

	c2 := g.FindContent("test_tablename")
	assert.Equal(t, "test_tablename", c2.TableNam)
	assert.Equal(t, "test_datatype", c2.DataType)
	assert.Equal(t, "test_identifier", c2.Identifier)
	assert.Equal(t, "test_description", c2.Description)

	c2.Description = "Test Description"
	if err = g.ModifyContent(c2); err != nil {
		assert.Error(t, err)
	}

	c3 := g.FindContent(c2.TableNam)
	assert.Equal(t, c3.Description, c2.Description)
	if err = g.RemoveContent(c3.TableNam); err != nil {
		assert.Error(t, err)
	}

	c4 := g.FindContent(c2.TableNam)
	assert.Equal(t, "", c4.TableNam)
	g.Close()
}
