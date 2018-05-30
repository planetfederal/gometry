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

type TileMatrix struct {
	TableNam     string  `gorm:"type:TEXT;not null"`
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
