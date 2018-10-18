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
