package gpkg

import (
	"database/sql"
	"errors"
	"fmt"
)

type Geopackage struct {
	db *sql.DB
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
