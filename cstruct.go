package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -ltennis -lSeetaAuthorize -lSeetaFaceDetector600
// #include <stdlib.h>
// #include "CStruct.h"
import "C"
import (
	"log"
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
func (s *SeetaImageData) SetData(data []uint8) {
	// TODO: 完成数据转换和设置
}
func (s *SeetaImageData) Free() {
	// ptr := unsafe.Pointer(s.ptr)
	// C.free(ptr)
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

// SeetaModelSetting 模型配置数据结构
type SeetaModelSetting struct {
	ptr *C.struct_SeetaModelSetting
}

func NewSeetaModelSetting(models []string) {
	var setting C.struct_SeetaModelSetting
	setting.device = C.SEETA_DEVICE_AUTO
	setting.id = 0
	// TODO:设置模型

}

func Test() {
	a := NewSeetaImageData(320, 160, 3)
	// defer a.Free()
	// sid.width = 123
	log.Println(a.ptr)

}
