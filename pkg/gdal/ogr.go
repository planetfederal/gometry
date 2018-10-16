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
