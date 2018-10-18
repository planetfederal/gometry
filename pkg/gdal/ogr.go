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

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: gdal
// #cgo pkg-config: json-c
// #include "ogr.h"
import "C"
import "encoding/json"

type OGRInfoResponse struct {
	Driver       string `json:"driver"`
	File         string `json:"file"`
	LayerName    string `json:"layer_name"`
	Geometry     string `json:"geometry"`
	FeatureCount int    `json:"feature_count"`
	Fid          string `json:"fid_column"`
}

func OGRInfo(filename, layer string) (*OGRInfoResponse, error) {
	fname := C.CString(filename)
	lname := C.CString(layer)
	out, err := C.OGRInfoJSONString(fname, lname)
	if err != nil {
		return nil, err
	}
	jsonStr := C.GoString(out)
	var oir OGRInfoResponse
	err = json.Unmarshal([]byte(jsonStr), &oir)
	if err != nil {
		return nil, err
	}

	return &oir, nil
}
