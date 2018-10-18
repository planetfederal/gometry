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
