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
