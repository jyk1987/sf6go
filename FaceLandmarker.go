package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -lSeetaFaceLandmarker600
// #include <stdlib.h>
// #include "FaceLandmarker_warp.h"
import "C"
import (
	"log"
	"unsafe"
)

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

// Number 获取当前模型的特征点数
func (s *FaceLandmarker) Number() int {
	return int(C.facelandmarker_number(s.ptr))
}

func (s *FaceLandmarker) Close() {
	C.facelandmarker_free(s.ptr)
}

func TestFaceLandmarker() {
	model := "/var/sf6/models/face_landmarker_pts5.csta"
	fl := NewFaceLandmarker(model)
	log.Println(fl.Number())
}
