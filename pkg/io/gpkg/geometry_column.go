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

func (g *Geopackage) AddGeometryColumn(gc GeometryColumn) error {
	return g.db.Create(&gc).Error
}

func (g *Geopackage) ModifyGeometryColumn(gc GeometryColumn) error {
	return g.db.Model(&gc).Where("table_name = ?", gc.TableNam).Updates(gc).Error
}

func (g *Geopackage) FindGeometryColumn(tableName string) GeometryColumn {
	var gc GeometryColumn
	g.db.First(&gc, "table_name = ?", tableName)
	return gc
}

func (g *Geopackage) RemoveGeometryColumn(tableName string) error {
	return g.db.Where("table_name = ?", tableName).Delete(GeometryColumn{}).Error
}
