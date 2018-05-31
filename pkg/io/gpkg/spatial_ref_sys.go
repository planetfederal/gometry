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
