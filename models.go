package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -ltennis -lSeetaAuthorize
// #include <stdlib.h>
// #include "CStruct.h"
// #include "CFaceInfo.h"
import "C"
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

// SeetaImageData 图像数据结构
type SeetaImageData struct {
	_ptr  C.struct_SeetaImageData
	cdata []C.uchar //此数据最终将指针交给c处理，此数据时为了方式数据逃逸，方便go释放内存
}

func (s *SeetaImageData) GetWidth() int {
	return int(s._ptr.width)
}
func (s *SeetaImageData) GetHeight() int {
	return int(s._ptr.height)
}
func (s *SeetaImageData) GetChannels() int {
	return int(s._ptr.channels)
}

func (s *SeetaImageData) GetData() []uint8 {
	size := len(s.cdata)
	data := make([]uint8, size)
	for i := 0; i < size; i++ {
		data[i] = uint8(s.cdata[i])
	}
	return data
}

func (s *SeetaImageData) getCStruct() C.struct_SeetaImageData {
	s.Reset()
	return s._ptr
}

func (s *SeetaImageData) SetUint8(data []uint8) error {
	if len(s.cdata) != len(data) {
		return fmt.Errorf("设置的数据与初始化的大小不符")
	}
	for i, v := range data {
		s.cdata[i] = C.uchar(v)
	}
	return nil
}

func (s *SeetaImageData) Reset() {
	s._ptr.data = &s.cdata[0]
}

func (s *SeetaImageData) Close() {
	// C.free(unsafe.Pointer(&s.cdata))
	// C.free(unsafe.Pointer(&s.ptr))
	// C.free(unsafe.Pointer(&s.ptr.width))
	// C.free(unsafe.Pointer(&s.ptr.height))
	// C.free(unsafe.Pointer(&s.ptr.channels))
}

// NewSeetaImageData 创建一个sf6图片（不包含数据）
// 创建后需要通过SetUint8设置数据，一般用于opencv的mat转换
func NewSeetaImageData(width, height, channels int) *SeetaImageData {
	imageData := &SeetaImageData{
		cdata: make([]C.uchar, width*height*channels),
		_ptr: C.struct_SeetaImageData{
			width:    C.int(width),
			height:   C.int(height),
			channels: C.int(channels),
		},
	}
	imageData._ptr.data = &imageData.cdata[0]
	return imageData
}

// NewSeetaImageDataFromBase64 通过base64图片数据创建sf6图片数据
func NewSeetaImageDataFromBase64(data string) (*SeetaImageData, error) {

	buffer, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(buffer)

	img, foramt, err := image.Decode(reader)
	if err != nil {
		log.Println(foramt)
		return nil, err
	}
	return NewSeetaImageDataFromImage(img), nil
}

// NewSeetaImageDataFromFile 根据文件路径创建sf6图片数据
func NewSeetaImageDataFromFile(filePath string) (*SeetaImageData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, foramt, err := image.Decode(file)
	if err != nil {
		log.Println(foramt)
		return nil, err
	}
	return NewSeetaImageDataFromImage(img), nil
}

// NewSeetaImageDataFromImage 根据go 内置image数据创建sf6图片数据
func NewSeetaImageDataFromImage(img image.Image) *SeetaImageData {
	rect := img.Bounds()
	width := rect.Dx()
	height := rect.Dy()
	channels := 3

	imageData := &SeetaImageData{
		cdata: make([]C.uchar, width*height*channels),
		_ptr: C.struct_SeetaImageData{
			width:    C.int(width),
			height:   C.int(height),
			channels: C.int(channels),
		},
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := img.At(rect.Min.X+x, rect.Min.Y+y)
			r, g, b, _ := c.RGBA()
			offset := y*width*channels + x*channels
			imageData.cdata[offset] = C.uchar(b >> 8)
			imageData.cdata[offset+1] = C.uchar(g >> 8)
			imageData.cdata[offset+2] = C.uchar(r >> 8)
		}
	}

	imageData._ptr.data = &imageData.cdata[0]
	return imageData
}

// SeetaRect 人脸位置信息
type SeetaRect struct {
	_ptr C.struct_SeetaRect
}

func newSeetaRect(seetaRect C.struct_SeetaRect) *SeetaRect {
	return &SeetaRect{
		_ptr: seetaRect,
	}
}

func (s *SeetaRect) getCStruct() C.struct_SeetaRect {
	return s._ptr
}

func (s *SeetaRect) GetX() int {
	return int(s._ptr.x)
}
func (s *SeetaRect) GetY() int {
	return int(s._ptr.y)
}
func (s *SeetaRect) GetWidth() int {
	return int(s._ptr.width)
}
func (s *SeetaRect) GetHeight() int {
	return int(s._ptr.height)
}

type SeetaFaceInfo struct {
	Postion *SeetaRect
	Score   float32
}

func NewSeetaFaceInfo(seetaFaceInfo C.struct_SeetaFaceInfo) *SeetaFaceInfo {
	return &SeetaFaceInfo{
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

func (s *SeetaPointInfo) getCSeetaPointFArray() *C.struct_SeetaPointF {
	return &s.Points[0]
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
	log.Println(a.getCStruct())

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
