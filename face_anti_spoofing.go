package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSeetaFaceAntiSpoofingX600
// #include <stdlib.h>
// #include "FaceAntiSpoofing_warp.h"
import "C"
import (
	"path/filepath"
	"unsafe"
)

type FaceAntiSpoofingStatus int

const (
	FaceAntiSpoofingStatus_REAL      = 0 ///< 真实人脸
	FaceAntiSpoofingStatus_SPOOF     = 1 ///< 攻击人脸（假人脸）
	FaceAntiSpoofingStatus_FUZZY     = 2 ///< 无法判断（人脸成像质量不好）
	FaceAntiSpoofingStatus_DETECTING = 3 ///< 正在检测
)

type FaceAntiSpoofing struct {
	ptr *C.struct_faceantispoofing
}

var FaceAntiSpoofing_model = []string{"fas_first.csta", "fas_second.csta"}

// NewFaceAntiSpoofing 创建局部活体检测器,速度快
func NewFaceAntiSpoofing() *FaceAntiSpoofing {
	first := C.CString(filepath.Join(_model_base_path, FaceAntiSpoofing_model[0]))
	defer C.free(unsafe.Pointer(first))
	fd := &FaceAntiSpoofing{
		ptr: C.faceantispoofing_new(first),
	}
	return fd
}

// NewFaceAntiSpoofing_v2 创建全局活体检测器，速度慢
func NewFaceAntiSpoofing_v2() *FaceAntiSpoofing {
	first := C.CString(filepath.Join(_model_base_path, FaceAntiSpoofing_model[0]))
	second := C.CString(filepath.Join(_model_base_path, FaceAntiSpoofing_model[1]))
	defer func() {
		C.free(unsafe.Pointer(first))
		C.free(unsafe.Pointer(second))
	}()
	fd := &FaceAntiSpoofing{
		ptr: C.faceantispoofing_new_v2(first, second),
	}
	return fd
}

func (s *FaceAntiSpoofing) Close() {
	C.faceantispoofing_free(s.ptr)
}

func (s *FaceAntiSpoofing) Predict(image *SeetaImageData, postion *SeetaRect, pointInfo *SeetaPointInfo) FaceAntiSpoofingStatus {
	status := C.faceantispoofing_predict(s.ptr, image.getCStruct(), postion.getCStruct(), pointInfo.getCSeetaPointFArray())
	return FaceAntiSpoofingStatus(int(status))
}
