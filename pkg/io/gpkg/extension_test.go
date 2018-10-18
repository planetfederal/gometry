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

	"github.com/stretchr/testify/assert"
)

func TestExtension(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
		return
	}

	ext := Extension{
		TableNam:      "test_tablename",
		ColumnName:    "test_columnname",
		ExtensionName: "test_extensionname",
		Definition:    "test_definition",
		Scope:         "test_scope",
	}

	if err := g.AddExtension(ext); err != nil {
		assert.Error(t, err)
	}

	ext2 := g.FindExtension(ext.ExtensionName)
	assert.Equal(t, "test_tablename", ext2.TableNam)
	assert.Equal(t, "test_columnname", ext2.ColumnName)
	assert.Equal(t, "test_extensionname", ext2.ExtensionName)

	ext2.Definition = "Test Definition"
	if err = g.ModifyExtension(ext2); err != nil {
		assert.Error(t, err)
	}

	ext3 := g.FindExtension(ext2.ExtensionName)
	assert.Equal(t, ext2.Definition, ext3.Definition)
	if err = g.RemoveExtension(ext3.ExtensionName); err != nil {
		assert.Error(t, err)
	}

	ext4 := g.FindExtension(ext2.ExtensionName)
	assert.Equal(t, "", ext4.TableNam)
	g.Close()
	os.Remove("/tmp/extension_test.db")
}
