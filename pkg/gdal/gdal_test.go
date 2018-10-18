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
package gdal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGdalInfo(t *testing.T) {
	fname := "/vsis3/landsat-pds/L8/210/244/LC82102442014232LGN00/LC82102442014232LGN00_B1.TIF"
	gir, err := GDALInfo(fname)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, fname, gir.Description)
	assert.Equal(t, "GTiff", gir.DriverShortName)
	assert.Equal(t, "GeoTIFF", gir.DriverLongName)
	assert.NotEmpty(t, gir.Files)
}
