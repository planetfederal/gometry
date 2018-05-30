package gpkg

type Extension struct {
	TableNam      string `gorm:"type:TEXT"`
	ColumnName    string `gorm:"type:TEXT"`
	ExtensionName string `gorm:"type:TEXT;not null"`
	Definition    string `gorm:"type:TEXT;not null"`
	Scope         string `gorm:"type:TEXT;not null"`
}

func (Extension) TableName() string {
	return "gpkg_extensions"
}
