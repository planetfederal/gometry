package gdal

// #cgo pkg-config: gdal
// #include "goodle.h"
import "C"

import (
	"encoding/json"
)

type GDALInfoResponse struct {
	Description       string `json:"description"`
	DriverShortName   string
	DriverLongName    string
	Files             []string
	Size              []int32
	CoordinateSystem  map[string]string
	GeoTransform      []float32
	Metadata          map[string]interface{}
	CornerCoordinates map[string]interface{}
}

//Info takes a filename can calls GDAL Info using CGo
func Info(filename string) (*GDALInfoResponse, error) {
	val := C.CString(filename)
	out, err := C.Info(val)
	if err != nil {
		return nil, err
	}
	jsonStr := C.GoString(out)
	var gir GDALInfoResponse
	err = json.Unmarshal([]byte(jsonStr), &gir)
	if err != nil {
		return nil, err
	}
	return &gir, nil
}
