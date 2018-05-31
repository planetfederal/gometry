package gpkg

import (
	"time"
)

var contentsTableName = "gpkg_contents"

//Contents is a representation of a gpgk_contents row
type Content struct {
	TableNam               string                 `gorm:"column:table_name;type:TEXT;not null;PRIMARY_KEY`
	DataType               string                 `gorm:"type:TEXT;not null"`
	Identifier             string                 `gorm:"type:TEXT;unique"`
	Description            string                 `gorm:"type:TEXT;DEFAULT ''"`
	LastChange             time.Time              `gorm:"type:DATETIME";not null;DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')`
	MinX                   float64                `gorm:"type:DOUBLE"`
	MinY                   float64                `gorm:"type:DOUBLE"`
	MaxX                   float64                `gorm:"type:DOUBLE"`
	MaxY                   float64                `gorm:"type:DOUBLE"`
	SrsId                  int32                  `gorm:"type:INTEGER"`
	SpatialReferenceSystem SpatialReferenceSystem `gorm:"column:fk_gc_r_srs_id"`
}

func (Content) TableName() string {
	return "gpkg_contents"
}

func (g *Geopackage) AddContent(c Content) error {
	return g.db.Create(c).Error
}

func (g *Geopackage) ModifyContent(c Content) error {
	return g.db.Save(&c).Error
}

func (g *Geopackage) FindContent(tableName string) Content {
	var c Content
	g.db.First(&c, tableName)
	return c
}

func (g *Geopackage) RemoveContent(tableName string) error {
	return g.db.Delete(&Content{TableNam: tableName}).Error
}
