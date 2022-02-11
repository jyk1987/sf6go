package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSeetaFaceLandmarker600
// #include <stdlib.h>
// #include "FaceLandmarker_warp.h"
import "C"
import (
	"log"
	"unsafe"
)

type FaceLandmarker struct {
	ptr        *C.struct_facelandmarker
	pointCount int
}

func NewFaceLandmarker(model string) *FaceLandmarker {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	fl := &FaceLandmarker{
		ptr: C.newFaceLandmarker(cs),
	}
	fl.pointCount = fl.Number()
	return fl
}

// Number 获取当前模型的特征点数
func (s *FaceLandmarker) Number() int {
	return int(C.facelandmarker_number(s.ptr))
}

func (s *FaceLandmarker) Mark(image *SeetaImageData, faceInfo SeetaFaceInfo) *SeetaPointInfo {
	pointInfo := NewSeetaPointInfo(s.pointCount)
	image.Reset()
	cmask := make([]C.int, s.pointCount)
	C.facelandmarker_mark(s.ptr, image.ptr, faceInfo.Postion.ptr, &pointInfo.Points[0], &cmask[0])
	// TODO 换转cint to go bool
	for i := 0; i < s.pointCount; i++ {
		pointInfo.Masks[i] = int(cmask[i]) == 1
	}
	// log.Println(pointInfo)
	return pointInfo
}

func (s *FaceLandmarker) Close() {
	C.facelandmarker_free(s.ptr)
}

func TestFaceLandmarker() {
	model := "/var/sf6/models/face_landmarker_pts5.csta"
	fl := NewFaceLandmarker(model)
	log.Println(fl.Number())
}
