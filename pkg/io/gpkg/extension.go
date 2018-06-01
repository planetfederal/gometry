package gpkg

type Extension struct {
	TableNam      string `gorm:"column:table_name;type:TEXT"`
	ColumnName    string `gorm:"type:TEXT"`
	ExtensionName string `gorm:"type:TEXT;not null"`
	Definition    string `gorm:"type:TEXT;not null"`
	Scope         string `gorm:"type:TEXT;not null"`
}

func (Extension) TableName() string {
	return "gpkg_extensions"
}

func (g *Geopackage) AddExtension(ext Extension) error {
	if err := g.db.Create(&ext).Error; err != nil {
		return err
	}
	return nil
}

func (g *Geopackage) ModifyExtension(ext Extension) error {
	return g.db.Model(&ext).Where("extension_name = ?", ext.ExtensionName).Updates(ext).Error
}

func (g *Geopackage) FindExtension(extName string) Extension {
	var ext Extension
	g.db.First(&ext, "extension_name = ?", extName)
	return ext
}

func (g *Geopackage) RemoveExtension(extName string) error {
	return g.db.Where("extension_name = ?", extName).Delete(Extension{}).Error
}
