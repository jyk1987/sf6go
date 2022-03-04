package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSeetaFaceLandmarker600
// #include <stdlib.h>
// #include "FaceLandmarker_warp.h"
import "C"
import (
	"path/filepath"
	"unsafe"
)

type FaceLandmarker struct {
	ptr        *C.struct_facelandmarker
	PointCount int
	FaceType   ModelType
}

var _FaceLandmarker_model = map[ModelType]string{
	ModelType_default: "face_landmarker_pts68.csta",
	ModelType_light:   "face_landmarker_pts5.csta",
	ModelType_mask:    "face_landmarker_mask_pts5.csta",
}

// NewFaceLandmarker 创建人脸特征定位器
func NewFaceLandmarker(modelType ModelType) *FaceLandmarker {
	model := filepath.Join(_model_base_path, _FaceLandmarker_model[modelType])
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	fl := &FaceLandmarker{
		ptr:      C.faceLandmarker_new(cs),
		FaceType: modelType,
	}
	fl.PointCount = fl.number()
	return fl
}

// Number 获取当前模型的特征点数
func (s *FaceLandmarker) number() int {
	return int(C.facelandmarker_number(s.ptr))
}

// Mark_Mask 检测特征点和遮挡情况
// 返回的特征点遮挡信息只有在使用face_landmarker_mask_pts5.csta模型时有效
func (s *FaceLandmarker) Mark(img *SeetaImageData, postion *SeetaRect) *SeetaPointInfo {
	pointInfo := NewSeetaPointInfo(s.PointCount)
	img.Reset()
	cmask := make([]C.int, s.PointCount)
	C.facelandmarker_mark_mask(s.ptr, img.getCStruct(), postion.getCStruct(), &pointInfo.Points[0], &cmask[0])
	for i := 0; i < s.PointCount; i++ {
		pointInfo.Masks[i] = int(cmask[i]) == 1
	}
	return pointInfo
}

// func (s *FaceLandmarker) Mark(img *SeetaImageData, postion *SeetaRect) *SeetaPointInfo {
// 	pointInfo := NewSeetaPointInfo(s.PointCount)
// 	img.Reset()
// 	C.facelandmarker_mark(s.ptr, img.getCStruct(), postion.getCStruct(), &pointInfo.Points[0])
// 	return pointInfo
// }

func (s *FaceLandmarker) Close() {
	C.facelandmarker_free(s.ptr)
}
