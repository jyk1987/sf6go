package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -ltennis -lSeetaAuthorize
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
	ptr   C.struct_SeetaImageData
	cdata []C.uchar //此数据最终将指针交给c处理，此数据时为了方式数据逃逸，方便go释放内存
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
	s.cdata = make([]C.uchar, len(data))
	for i, v := range data {
		s.cdata[i] = C.uchar(v)
	}
	s.ptr.data = &s.cdata[0]
	return nil
}

func (s *SeetaImageData) Reset() {
	s.ptr.data = &s.cdata[0]
}

func (s *SeetaImageData) Close() {
	// C.free(unsafe.Pointer(&s.cdata))
	// C.free(unsafe.Pointer(&s.ptr))
	// C.free(unsafe.Pointer(&s.ptr.width))
	// C.free(unsafe.Pointer(&s.ptr.height))
	// C.free(unsafe.Pointer(&s.ptr.channels))
}

func NewSeetaImageData(width, height, channels int) *SeetaImageData {
	return &SeetaImageData{
		ptr: C.struct_SeetaImageData{
			width:    C.int(width),
			height:   C.int(height),
			channels: C.int(channels),
		},
	}
}

// SeetaRect 人脸位置信息
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

// SeetaPointInfo 人脸特征点信息
type SeetaPointInfo struct {
	PointCount int // 特征点数
	Points     []C.struct_SeetaPointF
	Masks      []bool
}

func NewSeetaPointInfo(pointCount int) *SeetaPointInfo {
	return &SeetaPointInfo{
		PointCount: pointCount,
		Points:     make([]C.struct_SeetaPointF, pointCount),
		Masks:      make([]bool, pointCount),
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
