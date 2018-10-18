package gdal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOgrInfo(t *testing.T) {
	fname := "/vsis3/gometry/wpi_features.gpkg"
	layerName := "WPI"
	oir, err := OGRInfo(fname, layerName)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, layerName, oir.LayerName)
	assert.Equal(t, "GPKG", oir.Driver)
	assert.Equal(t, "gpkg_id", oir.Fid)
	assert.Equal(t, fname, oir.File)
}
