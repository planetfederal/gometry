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

type SpatialReferenceSystem struct {
	SrsName                string `gorm:"type:TEXT;not null"`
	SrsID                  int    `gorm:"type:INTEGER;not null;PRIMARY_KEY"`
	Organization           string `gorm:"type:TEXT;not null"`
	OrganizationCoordsysID int    `gorm:"type:INTEGER;not null"`
	Definition             string `gorm:"type:TEXT;not null"`
	Description            string `gorm:"type:TEXT"`
}

func (SpatialReferenceSystem) TableName() string {
	return "gpkg_spatial_ref_sys"
}

func (g *Geopackage) AddSpatialReferenceSystem(srs SpatialReferenceSystem) error {
	return g.db.Create(&srs).Error
}

func (g *Geopackage) ModifySpatialReferenceSystem(srs SpatialReferenceSystem) error {
	return g.db.Save(&srs).Error
}

func (g *Geopackage) FindSpatialReferenceSystem(srsId int) SpatialReferenceSystem {
	var srs SpatialReferenceSystem
	g.db.First(&srs, srsId)
	return srs
}

func (g *Geopackage) RemoveSpatialReferenceSystem(srsId int) error {
	return g.db.Delete(&SpatialReferenceSystem{SrsID: srsId}).Error
}
