package sf6go

// #cgo CXXFLAGS: -std=c++11 -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -lSeetaMaskDetector200
// #include <stdlib.h>
// #include "MaskDetector_warp.h"
import "C"
import (
	"path/filepath"
	"unsafe"
)

const (
	MaskDetector_score float32 = 0.5
	MaskDetector_model string  = "mask_detector.csta"
)

type MaskDetector struct {
	ptr *C.struct_maskdetector
}

func NewMaskDetector() *MaskDetector {
	cs := C.CString(filepath.Join(_model_base_path, MaskDetector_model))
	defer C.free(unsafe.Pointer(cs))
	return &MaskDetector{
		ptr: C.maskdetector_new(cs),
	}
}

func (s *MaskDetector) Close() {
	C.maskdetector_free(s.ptr)
}

func (s *MaskDetector) Detect(img *SeetaImageData, postion *SeetaRect) bool {
	score := C.float(0.0)
	result := C.maskdetector_detect(s.ptr, img.getCStruct(), postion.getCStruct(), &score)
	return int(result) == 1 && float32(score) > MaskDetector_score
}
