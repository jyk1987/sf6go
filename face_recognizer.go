package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSeetaFaceRecognizer610
// #include <stdlib.h>
// #include "FaceRecognizer_warp.h"
import "C"
import "unsafe"

type FaceRecognizerProperty int

const (
	FaceRecognizer_PROPERTY_NUMBER_THREADS FaceRecognizerProperty = 4
	FaceRecognizer_PROPERTY_ARM_CPU_MODE   FaceRecognizerProperty = 5
)

type FaceRecognizer struct {
	ptr         *C.struct_facerecognizer
	FeatureSize int
}

// NewFaceRecognizer 创建一个人脸识别器
func NewFaceRecognizer(model string) *FaceRecognizer {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	fr := &FaceRecognizer{
		ptr: C.facerecognizer_new(cs),
	}
	fr.SetProperty(FaceRecognizer_PROPERTY_NUMBER_THREADS, 1)
	fr.FeatureSize = fr.getExtractFeatureSize()
	return fr
}

func (s *FaceRecognizer) SetProperty(property FaceRecognizerProperty, value float64) {
	C.facerecognizer_setProperty(s.ptr, C.int(property), C.double(value))
}

func (s *FaceRecognizer) GetProperty(property FaceRecognizerProperty) float64 {
	return float64(C.facerecognizer_getProperty(s.ptr, C.int(property)))
}

func (s *FaceRecognizer) GetCropFaceWidthV2() int {
	return int(C.facerecognizer_GetCropFaceWidthV2(s.ptr))
}
func (s *FaceRecognizer) GetCropFaceHeightV2() int {
	return int(C.facerecognizer_GetCropFaceHeightV2(s.ptr))
}
func (s *FaceRecognizer) GetCropFaceChannelsV2() int {
	return int(C.facerecognizer_GetCropFaceChannelsV2(s.ptr))
}

// GetExtractFeatureSize 获取当前模型的特征数量
func (s *FaceRecognizer) getExtractFeatureSize() int {
	return int(C.facerecognizer_GetExtractFeatureSize(s.ptr))
}

// Extract 提取人脸特征,从完整图像中提取人脸特征数据
// 返回值 bool代表提取是否成功
// 返回值 []float32为特征数据
func (s *FaceRecognizer) Extract(imageData *SeetaImageData, pointInfo *SeetaPointInfo) (bool, []float32) {
	cfeatures := make([]C.float, s.FeatureSize)
	success := int(C.facerecognizer_Extract(s.ptr, imageData.getCStruct(), pointInfo.getCSeetaPointFArray(), &cfeatures[0])) == 1
	if success {
		features := make([]float32, s.FeatureSize)
		for i := 0; i < s.FeatureSize; i++ {
			features[i] = float32(cfeatures[i])
		}
		return success, features
	}
	return success, nil
}

func (s *FaceRecognizer) Close() {
	C.facerecognizer_free(s.ptr)
}