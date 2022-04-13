package sf6go

// #cgo CXXFLAGS: -std=c++11 -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -lSeetaFaceTracking600
// #include <stdlib.h>
// #include "FaceTracker_warp.h"
// #include "CTrackingFaceInfo.h"
import "C"
import (
	"path/filepath"
	"reflect"
	"unsafe"
)

// FaceTracker 人脸追踪器
type FaceTracker struct {
	ptr *C.struct_facetracker
}

// NewFaceTracker 创建一个人脸追踪器
func NewFaceTracker(width, height int) *FaceTracker {
	cs := C.CString(filepath.Join(_model_base_path, _FaceDetector_model))
	defer C.free(unsafe.Pointer(cs))
	ft := &FaceTracker{
		ptr: C.facetracker_new(cs, C.int(width), C.int(height)),
	}
	ft.SetThreads(1)
	return ft
}

func (s *FaceTracker) Track(img *SeetaImageData) []*SeetaTrackingFaceInfo {
	var result C.struct_SeetaTrackingFaceInfoArray = C.facetracker_Track(s.ptr, img.getCStruct())
	var clist []C.struct_SeetaTrackingFaceInfo
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&clist))
	arrayLen := int(result.size)
	sliceHeader.Cap = arrayLen
	sliceHeader.Len = arrayLen
	sliceHeader.Data = uintptr(unsafe.Pointer(result.data))

	faceInfoList := make([]*SeetaTrackingFaceInfo, arrayLen)
	for i := 0; i < arrayLen; i++ {
		faceInfoList[i] = NewSeetaTrackingFaceInfo(clist[i])
	}
	// TODO: c free
	return faceInfoList
}

func (s *FaceTracker) Close() {
	C.facetracker_free(s.ptr)
}

func (s *FaceTracker) SetMinFaceSize(size int) {
	C.facetracker_SetMinFaceSize(s.ptr, C.int(size))
}

func (s *FaceTracker) GetMinFaceSize() int {
	return int(C.facetracker_GetMinFaceSize(s.ptr))
}

func (s *FaceTracker) SetThreshold(thresh float32) {
	C.facetracker_SetThreshold(s.ptr, C.float(thresh))
}

func (s *FaceTracker) GetThreshold() float32 {
	return float32(C.facetracker_GetThreshold(s.ptr))
}

func (s *FaceTracker) SetVideoStable(stable bool) {
	var cstable_int C.int = 0
	if stable {
		cstable_int = 1
	}
	C.facetracker_SetVideoStable(s.ptr, cstable_int)
}

func (s *FaceTracker) GetVideoStable() bool {
	return int(C.facetracker_GetVideoStable(s.ptr)) == 1
}

func (s *FaceTracker) SetThreads(num int) {
	C.facetracker_SetSingleCalculationThreads(s.ptr, C.int(num))
}

func (s *FaceTracker) SetInterval(interval int) {
	C.facetracker_SetInterval(s.ptr, C.int(interval))
}

func (s *FaceTracker) Reset() {
	C.facetracker_Reset(s.ptr)
}
