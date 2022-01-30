package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -ltennis -lSeetaAuthorize -lSeetaFaceDetector600
// #include <stdlib.h>
// #include "FaceDetector.c.h"
import "C"

import (
	"unsafe"
)

type FaceDetector struct {
	ptr *C.struct_facedetector
}

func newFaceDetector(model string) *FaceDetector {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	return &FaceDetector{
		ptr: C.newFaceDetector(cs),
	}
}

func (s *FaceDetector) Free() {
	C.free(unsafe.Pointer(s.ptr))
}

func TestFaceDetector() {
	model := "/var/sf6/models/face_detector.csta"
	fd := newFaceDetector(model)
	defer fd.Free()
}
