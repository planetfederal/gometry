package gpkg

import (
	"time"
)

var contentsTableName = "gpkg_contents"

//Contents is a representation of a gpgk_contents row
type Content struct {
	TableNam    string    `gorm:"column:table_name;type:TEXT;not null;PRIMARY_KEY`
	DataType    string    `gorm:"type:TEXT;not null"`
	Identifier  string    `gorm:"type:TEXT;unique"`
	Description string    `gorm:"type:TEXT;DEFAULT ''"`
	LastChange  time.Time `gorm:"type:DATETIME";not null;DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')`
	MinX        float64   `gorm:"type:DOUBLE"`
	MinY        float64   `gorm:"type:DOUBLE"`
	MaxX        float64   `gorm:"type:DOUBLE"`
	MaxY        float64   `gorm:"type:DOUBLE"`
	SrsId       int32     `gorm:"type:INTEGER"`
}

func (Content) TableName() string {
	return "gpkg_contents"
}

// func (g *Geopackage) Contents() []Content {
// 	qry := fmt.Sprintf("SELECT * FROM %s", contentsTableName)
// 	var contents []Content
// 	rows, err := g.db.Query(qry)
// 	if err != nil {
// 		log.Printf(err.Error())
// 		return contents
// 	}
// 	for rows.Next() {
// 		var content Content
// 		err := rows.Scan(&content)
// 		if err != nil {
// 			log.Printf(err.Error())
// 			break
// 		}
// 		contents = append(contents, content)
// 	}
// 	return contents
// }

// func (g *Geopackage) Content(tableName string) Content {
// 	qry := fmt.Sprintf("SELECT * FROM %s WHERE table_name = %s", contentsTableName, tableName)
// 	var content Content
// 	row := g.db.QueryRow(qry)
// 	err := row.Scan(&content)
// 	if err != nil {
// 		log.Print(err.Error())
// 		return

// 	return content
// }
