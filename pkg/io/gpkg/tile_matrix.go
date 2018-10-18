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

type TileMatrixSet struct {
	TableNam string  `gorm:"column:table_name;type:TEXT;not null;PRIMARY_KEY`
	SrsId    int32   `gorm:"type:INTEGER;not null"`
	MinX     float64 `gorm:"type:DOUBLE;not null"`
	MinY     float64 `gorm:"type:DOUBLE;not null"`
	MaxX     float64 `gorm:"type:DOUBLE;not null"`
	MaxY     float64 `gorm:"type:DOUBLE;not null"`
}

func (TileMatrixSet) TableName() string {
	return "gpkg_tile_matrix_set"
}

func (g *Geopackage) AddTileMatrixSet(tms TileMatrixSet) error {
	return g.db.Create(&tms).Error
}

func (g *Geopackage) ModifyTileMatrixSet(tms TileMatrixSet) error {
	return g.db.Save(&tms).Error
}

func (g *Geopackage) FindTileMatrixSet(tableName string) TileMatrixSet {
	var tms TileMatrixSet
	g.db.First(&tms, tableName)
	return tms
}

func (g *Geopackage) RemoveTileMatrixSet(tableName string) error {
	return g.db.Delete(TileMatrixSet{TableNam: tableName}).Error
}

type TileMatrix struct {
	TableNam     string  `gorm:"column:table_name;type:TEXT;not null"`
	ZoomLevel    int32   `gorm:"type:INTEGER;not null"`
	MatrixWidth  int32   `gorm:"type:INTEGER;not null"`
	MatrixHeight int32   `gorm:"type:INTEGER;not null"`
	TileWidth    int32   `gorm:"type:INTEGER;not null"`
	TileHeight   int32   `gorm:"type:INTEGER;not null"`
	PixelXSize   float64 `gorm:"type:DOUBLE;not null"`
	PixelYSize   float64 `gorm:"type:DOUBLE;not null"`
}

func (TileMatrix) TableName() string {
	return "gpkg_tile_matrix"
}

func (g *Geopackage) AddTileMatrix(tm TileMatrix) error {
	return g.db.Create(&tm).Error
}

func (g *Geopackage) ModifyTileMatrix(tm TileMatrix) error {
	return g.db.Model(&tm).Where("table_name = ?", tm.TableNam).Updates(&tm).Error
}

func (g *Geopackage) FindTileMatrix(tableName string) TileMatrix {
	var tm TileMatrix
	g.db.First(&tm, "table_name = ?", tableName)
	return tm
}

func (g *Geopackage) RemoveTileMatrix(tableName string) error {
	return g.db.Where("table_name = ?", tableName).Delete(TileMatrix{}).Error
}
