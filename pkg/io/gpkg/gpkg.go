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
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Geopackage struct {
	db *gorm.DB
}

type GeopackageContent interface {
	TableName() string
	GeomColumnName() string
	SrsId() int32
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

func (g *Geopackage) AddFeatureTable(model GeopackageContent) error {
	err := g.db.AutoMigrate(model).Error
	if err != nil {
		return err
	}
	return g.addContentTable(model, "features")
}

func (g *Geopackage) AddTileTable(model GeopackageContent) error {
	err := g.db.AutoMigrate(model).Error
	if err != nil {
		return err
	}
	return g.addContentTable(model, "tiles")
}

func (g *Geopackage) addContentTable(model GeopackageContent, dataType string) error {
	c := Content{
		TableNam:   model.TableName(),
		DataType:   dataType,
		Identifier: model.TableName(),
		LastChange: time.Now(),
		SrsId:      model.SrsId(),
	}
	return g.AddContent(c)
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
