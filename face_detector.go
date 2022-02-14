package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSeetaFaceDetector600
// #include <stdlib.h>
// #include "FaceDetector_warp.h"
import "C"

import (
	"path/filepath"
	"reflect"
	"unsafe"
)

type FaceDetectorProperty int

const (
	FaceDetector_PROPERTY_MIN_FACE_SIZE    FaceDetectorProperty = 0
	FaceDetector_PROPERTY_THRESHOLD        FaceDetectorProperty = 1
	FaceDetector_PROPERTY_MAX_IMAGE_WIDTH  FaceDetectorProperty = 2
	FaceDetector_PROPERTY_MAX_IMAGE_HEIGHT FaceDetectorProperty = 3
	FaceDetector_PROPERTY_NUMBER_THREADS   FaceDetectorProperty = 4
	FaceDetector_PROPERTY_ARM_CPU_MODE     FaceDetectorProperty = 0x101
)

type FaceDetector struct {
	ptr *C.struct_facedetector
}

const (
	_NewFaceDetector_model = "face_detector.csta"
)

// NewFaceDetector 创建一个人脸检测器
func NewFaceDetector() *FaceDetector {
	cs := C.CString(filepath.Join(_model_base_path, _NewFaceDetector_model))
	defer C.free(unsafe.Pointer(cs))
	fd := &FaceDetector{
		ptr: C.faceDetector_new(cs),
	}
	fd.SetProperty(FaceDetector_PROPERTY_NUMBER_THREADS, 1)
	return fd
}

func (s *FaceDetector) SetProperty(property FaceDetectorProperty, value float64) {
	C.facedetector_setProperty(s.ptr, C.int(property), C.double(value))
}

func (s *FaceDetector) GetProperty(property FaceDetectorProperty) float64 {
	return float64(C.facedetector_getProperty(s.ptr, C.int(property)))
}

func (s *FaceDetector) Detect(imageData *SeetaImageData) []*SeetaFaceInfo {
	var result C.struct_SeetaFaceInfoArray = C.facedetector_detect(s.ptr, imageData.getCStruct())
	var clist []C.struct_SeetaFaceInfo
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&clist))
	arrayLen := int(result.size)
	sliceHeader.Cap = arrayLen
	sliceHeader.Len = arrayLen
	sliceHeader.Data = uintptr(unsafe.Pointer(result.data))

	faceInfoList := make([]*SeetaFaceInfo, arrayLen)
	for i := 0; i < arrayLen; i++ {
		faceInfoList[i] = NewSeetaFaceInfo(clist[i])
	}
	return faceInfoList
}

func (s *FaceDetector) Close() {
	C.facedetector_free(s.ptr)
}
