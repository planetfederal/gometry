package gdal

import (
	"testing"
)

func TestGdalInfo(t *testing.T) {
	assert.Equals(t, "foo", Info("foo"))
}
