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
	PointCount int
}

// NewFaceLandmarker 创建人脸特征定位器
func NewFaceLandmarker(model string) *FaceLandmarker {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	fl := &FaceLandmarker{
		ptr: C.faceLandmarker_new(cs),
	}
	fl.PointCount = fl.number()
	return fl
}

// Number 获取当前模型的特征点数
func (s *FaceLandmarker) number() int {
	return int(C.facelandmarker_number(s.ptr))
}

func (s *FaceLandmarker) Mark(image *SeetaImageData, postion *SeetaRect) *SeetaPointInfo {
	pointInfo := NewSeetaPointInfo(s.PointCount)
	image.Reset()
	cmask := make([]C.int, s.PointCount)
	C.facelandmarker_mark(s.ptr, image.getCStruct(), postion.getCStruct(), &pointInfo.Points[0], &cmask[0])
	for i := 0; i < s.PointCount; i++ {
		pointInfo.Masks[i] = int(cmask[i]) == 1
	}
	return pointInfo
}

func (s *FaceLandmarker) Close() {
	C.facelandmarker_free(s.ptr)
}

func TestFaceLandmarker() {
	model := "/var/sf6/models/face_landmarker_pts5.csta"
	fl := NewFaceLandmarker(model)
	log.Println(fl.PointCount)
}
