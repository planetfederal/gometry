package gpkg

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Geopackage struct {
	db *sql.DB
}

type GeopackageContent struct {
	name string
}

type FeatureSource struct {
}

type TileSource struct {
}

func NewGeopackage(filename string) (*Geopackage, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error connecting to %s ", filename))
	}
	return &Geopackage{db}, nil
}

func (g *Geopackage) initializeSpatialMetadata() error {
	qry := "SELECT InitSpatialMetadata();"

	if row := g.db.QueryRow(qry); row == nil {
		return errors.New("Unable to run InitSpatialMetadata()")
	}
	return nil
}

func (g *Geopackage) validateSchema() error {
	qry := "SELECT CheckSpatialMetadata();"

	if row := g.db.QueryRow(qry); row == nil {
		return errors.New("Unable to CheckSpatialMetadata()")
	}
	return nil
}

func (g *Geopackage) createSpatialIndex(tableName, geomColumnName, pkName string) error {
	qry := fmt.Sprintf("SELECT CreateSpatialIndex('%s', '%s', '%s')", tableName, geomColumnName, pkName)
	row := g.db.QueryRow(qry)
	if row == nil {
		return errors.New(fmt.Sprintf("Unable to create spatial index for %s:%s:%s", tableName, geomColumnName, pkName))
	}
	return nil
}

func (g *Geopackage) GeoPackageContents() []GeopackageContent {
	qry := "SELECT name FROM gpkg_contents"
	var contents []GeopackageContent
	rows, err := g.db.Query(qry)
	if err != nil {
		return contents
	}
	for rows.Next() {
		var content GeopackageContent
		err := rows.Scan(&content.name)
		if err != nil {
			log.Printf(err.Error())
			break
		}
		contents = append(contents, content)
	}
	return contents
}

func (g *Geopackage) TileSources() {

}

func (g *Geopackage) Close() error {
	return g.db.Close()
}
