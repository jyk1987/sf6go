package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -ltennis -lSeetaAuthorize -lSeetaFaceDetector600
// #include <stdlib.h>
// #include "CStruct.h"
// #include "CFaceInfo.h"
import "C"
import (
	"log"

	"gocv.io/x/gocv"
)

// SeetaImageData 图像数据结构
type SeetaImageData struct {
	ptr *C.struct_SeetaImageData
}

func (s *SeetaImageData) GetWidth() int {
	return int(s.ptr.width)
}
func (s *SeetaImageData) SetWidth(width int) {
	s.ptr.width = C.int(width)
}
func (s *SeetaImageData) GetHeight() int {
	return int(s.ptr.height)
}
func (s *SeetaImageData) SetHeight(height int) {
	s.ptr.height = C.int(height)
}
func (s *SeetaImageData) GetChannels() int {
	return int(s.ptr.channels)
}
func (s *SeetaImageData) SetChannels(channels int) {
	s.ptr.channels = C.int(channels)
}

func (s *SeetaImageData) GetData() []uint8 {
	// TODO: 完成数据获取
	return nil
}
func (s *SeetaImageData) SetMat(mat *gocv.Mat) error {
	data, err := mat.DataPtrUint8()
	if err != nil {
		return err
	}
	cdata := make([]C.uchar, len(data))
	for i, v := range data {

		cdata[i] = C.uchar(v)
		// log.Println(cdata[i], ":", v)
	}
	s.ptr.data = &cdata[0]
	return nil
}
func (s *SeetaImageData) Close() {
	// C.free(unsafe.Pointer(s.ptr))
	// C.free(unsafe.Pointer(&s.ptr.width))
	// C.free(unsafe.Pointer(&s.ptr.height))
	// C.free(unsafe.Pointer(&s.ptr.channels))
}

func NewSeetaImageData(width, height, channels int) *SeetaImageData {
	return &SeetaImageData{
		ptr: &C.struct_SeetaImageData{
			width:    C.int(width),
			height:   C.int(height),
			channels: C.int(channels),
		},
	}
}

type SeetaRect struct {
	ptr C.SeetaRect
}

func newSeetaRect(seetaRect C.SeetaRect) SeetaRect {
	return SeetaRect{
		ptr: seetaRect,
	}
}

func (s *SeetaRect) GetX() int {
	return int(s.ptr.x)
}
func (s *SeetaRect) GetY() int {
	return int(s.ptr.y)
}
func (s *SeetaRect) GetWidth() int {
	return int(s.ptr.width)
}
func (s *SeetaRect) GetHeight() int {
	return int(s.ptr.height)
}

type SeetaFaceInfo struct {
	Postion SeetaRect
	Score   float32
}

func NewSeetaFaceInfo(seetaFaceInfo C.struct_SeetaFaceInfo) SeetaFaceInfo {
	return SeetaFaceInfo{
		Postion: newSeetaRect(seetaFaceInfo.pos),
		Score:   float32(seetaFaceInfo.score),
	}
}

// SeetaModelSetting 模型配置数据结构
// type SeetaModelSetting struct {
// 	ptr *C.struct_SeetaModelSetting
// }

// func NewSeetaModelSetting(model string) *SeetaModelSetting {
// 	var setting C.struct_SeetaModelSetting
// 	setting.device = C.SEETA_DEVICE_AUTO
// 	setting.id = 0
// 	css := make([]*C.char, len(models))
// 	for i, v := range models {
// 		cs := C.CString(v)
// 		defer C.free(unsafe.Pointer(cs))
// 		css[i] = cs
// 	}
// 	setting.model = &css[0]
// 	return &SeetaModelSetting{
// 		ptr: &setting,
// 	}
// }

func TestCStruct() {
	a := NewSeetaImageData(320, 160, 3)
	defer a.Close()
	log.Println(a.ptr)

}

// GoStrings 将字符串数组转换成go的字符串数组
// func GoStrings(argc C.int, argv **C.char) []string {

// 	length := int(argc)
// 	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
// 	gostrings := make([]string, length)
// 	for i, s := range tmpslice {
// 		gostrings[i] = C.GoString(s)
// 	}
// 	return gostrings
// }
