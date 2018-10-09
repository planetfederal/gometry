package gdal

// #cgo pkg-config: gdal
// #include "goodle.h"
import "C"

import "log"

//Info takes a filename can calls GDAL Info using CGo
func Info(filename string) string {
	val := C.CString(filename)
	out, err := C.Info(val)
	if err != nil {
		log.Fatal(err)
	}
	return C.GoString(out)
}
