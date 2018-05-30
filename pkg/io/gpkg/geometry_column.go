package gpkg

type GeometryColumn struct {
	TableNam         string `gorm:"column:table_name;type:TEXT;not null"`
	ColumnName       string `gorm:"type:TEXT;not null"`
	GeometryTypeName string `gorm:"type:TEXT;not null"`
	SrsId            int32  `gorm:"type:INTEGER;not null"`
	Z                int8   `gorm:"type:TINYINT;nt null"`
	M                int8   `gorm:"type:TINYINT;not null`
}

func (GeometryColumn) TableName() string {
	return "gpkg_geometry_columns"
}
