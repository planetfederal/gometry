package geoserver

import (
	"testing"

	"github.com/boundlessgeo/gometry/pkg/geoserver/openapi"
)

func TestClient(t *testing.T) {
	client := openapi.NewAPIClient(nil)
	t.Fail()
}
