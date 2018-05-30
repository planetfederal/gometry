package gpkg

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Geopackage struct {
	db *gorm.DB
}

func NewGeopackage(filename string) (*Geopackage, error) {
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Content{})
	db.AutoMigrate(&GeometryColumn{})
	db.AutoMigrate(&TileMatrix{})
	db.AutoMigrate(&TileMatrixSet{})
	db.AutoMigrate(&Extension{})
	db.AutoMigrate(&SpatialReferenceSystem{})

	return &Geopackage{db}, nil
}

// func (g *Geopackage) initializeSpatialMetadata() error {
// 	qry := "SELECT InitSpatialMetadata();"

// 	if row := g.db.QueryRow(qry); row == nil {
// 		return errors.New("Unable to run InitSpatialMetadata()")
// 	}
// 	return nil
// }

// func (g *Geopackage) validateSchema() error {
// 	qry := "SELECT CheckSpatialMetadata();"

// 	if row := g.db.QueryRow(qry); row == nil {
// 		return errors.New("Unable to CheckSpatialMetadata()")
// 	}
// 	return nil
// }

// func (g *Geopackage) createSpatialIndex(tableName, geomColumnName, pkName string) error {
// 	qry := fmt.Sprintf("SELECT CreateSpatialIndex('%s', '%s', '%s')", tableName, geomColumnName, pkName)
// 	row := g.db.QueryRow(qry)
// 	if row == nil {
// 		return errors.New(fmt.Sprintf("Unable to create spatial index for %s:%s:%s", tableName, geomColumnName, pkName))
// 	}
// 	return nil
// }

func (g *Geopackage) Close() error {
	return g.db.Close()
}
