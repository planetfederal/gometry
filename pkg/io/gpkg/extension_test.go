package gpkg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtension(t *testing.T) {
	g, err := NewGeopackage("/tmp/extension_test.db")
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

	g.db.Create(&ext)

	var ext2 Extension
	g.db.Where("table_name = ?", ext.TableNam).First(&ext2)
	assert.Equal(t, "test_tablename", ext2.TableNam)
	assert.Equal(t, "test_columnname", ext2.ColumnName)
	assert.Equal(t, "test_extensionname", ext2.ExtensionName)

	ext2.Definition = "Test Definition"
	g.db.Model(&ext2).Where("table_name = ?", ext2.TableNam).Updates(ext2)
	if g.db.Error != nil {
		t.Error(g.db.Error)
		return
	}

	var ext3 Extension
	g.db.Where("table_name = ?", ext2.TableNam).First(&ext3)
	assert.Equal(t, ext2.Definition, ext3.Definition)

	g.db.Delete(&ext3)

	var ext4 Extension
	g.db.Where("table_name = ?", ext2.TableNam).First(&ext4)
	assert.Equal(t, "", ext4.TableNam)

	g.Close()
	os.Remove("/tmp/extension_test.db")
}
