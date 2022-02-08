package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -lSeetaFaceLandmarker600
// #include <stdlib.h>
// #include "FaceLandmarker_warp.h"
import "C"
import "unsafe"

type FaceLandmarker struct {
	ptr *C.struct_facelandmarker
}

func NewFaceLandmarker(model string) *FaceLandmarker {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	return &FaceLandmarker{
		ptr: C.newFaceLandmarker(cs),
	}
}

func (s *FaceLandmarker) Close() {
	C.facelandmarker_free(s.ptr)
}
