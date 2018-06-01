package gpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSRS(t *testing.T) {
	g, err := NewGeopackage(GeopackageTestFileName)
	if err != nil {
		t.Error(err)
		return
	}
	srs := SpatialReferenceSystem{
		SrsName:                "GCS_WGS_1984",
		SrsID:                  4326,
		Organization:           "EPSG",
		OrganizationCoordsysID: 4326,
		Definition:             `GEOGCS["GCS_WGS_1984",DATUM["D_WGS_1984",SPHEROID["WGS_1984",6378137.0,298.257223563]],PRIMEM["Greenwich",0.0],UNIT["Degree",0.0174532925199433]]"`,
		Description:            "WGS 1984",
	}
	g.db.Create(&srs)

	var srs2 SpatialReferenceSystem
	g.db.First(&srs2, 4326)
	assert.Equal(t, "EPSG", srs2.Organization)
	assert.Equal(t, 4326, srs2.OrganizationCoordsysID)
	assert.Equal(t, "WGS 1984", srs2.Description)

	srs2.Description = "Test Description"
	g.db.Save(&srs2)
	if g.db.Error != nil {
		t.Error(g.db.Error)
		return
	}

	var srs3 SpatialReferenceSystem
	g.db.First(&srs3, 4326)
	assert.Equal(t, srs3.Description, srs2.Description)

	g.db.Delete(&srs3)

	var srs4 SpatialReferenceSystem
	g.db.First(&srs4, 4326)
	assert.Equal(t, 0, srs4.SrsID)

	g.Close()
}
