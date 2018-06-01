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
