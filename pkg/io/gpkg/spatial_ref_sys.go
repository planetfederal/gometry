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
