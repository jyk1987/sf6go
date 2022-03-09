package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -ltennis -lSeetaAuthorize
// #include <stdlib.h>
// #include "CStruct.h"
// #include "CFaceInfo.h"
// #include "CTrackingFaceInfo.h"
import "C"
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"reflect"
	"unsafe"
)

// _model_base_path 模型基础路径
var _model_base_path string

// 初始化模型目录
func InitModelPath(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return fmt.Errorf("%v不是目录", path)
	}
	_model_base_path = path
	return nil
}

// ModelType 模型类型
type ModelType uint8

const (
	ModelType_default ModelType = iota // 默认模型，68特征点
	ModelType_light                    // 轻量级模型，5特征点
	ModelType_mask                     // 口罩模型，5特征点
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

func (s *SeetaImageData) GetImage() image.Image {
	width := s.GetWidth()
	height := s.GetHeight()
	channel := s.GetChannels()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	index := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.SetRGBA(x, y, color.RGBA{
				B: uint8(s.cdata[index]),
				G: uint8(s.cdata[index+1]),
				R: uint8(s.cdata[index+2]),
			})
			index += channel
		}
	}
	return img
}

func (s *SeetaImageData) CutFace(rect *SeetaRect) image.Image {
	rx := rect.GetX()
	ry := rect.GetY()
	width := rect.GetWidth()
	height := rect.GetHeight()
	channel := s.GetChannels()
	if rx-(height-width)/2 >= 0 {
		rx = rx - (height-width)/2
		width = height
	} else if rx > 0 {
		width += rx * 2
		rx = 0
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	originalWidth := s.GetWidth()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := (ry+y)*originalWidth*channel + (rx+x)*channel
			img.SetRGBA(x, y, color.RGBA{
				B: uint8(s.cdata[index]),
				G: uint8(s.cdata[index+1]),
				R: uint8(s.cdata[index+2]),
			})
		}
	}
	return img
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

func NewSeetaImageDataFromCStruct(cstruct C.struct_SeetaImageData) *SeetaImageData {
	imageData := &SeetaImageData{
		_ptr: cstruct,
	}
	var clist []C.uchar
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&clist))
	arrayLen := int(imageData.GetWidth() * imageData.GetHeight() * imageData.GetChannels())
	sliceHeader.Cap = arrayLen
	sliceHeader.Len = arrayLen
	sliceHeader.Data = uintptr(unsafe.Pointer(cstruct.data))
	var cdata []C.uchar = make([]C.uchar, arrayLen)
	for i := 0; i < arrayLen; i++ {
		cdata[i] = clist[i]
	}
	imageData.cdata = cdata
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

// SeetaTrackingFaceInfo 人脸追踪结果信息
type SeetaTrackingFaceInfo struct {
	Postion  *SeetaRect
	Score    float32
	Frame_NO int
	PID      int
	Step     int
}

func NewSeetaTrackingFaceInfo(seetaTrackingFaceInfo C.struct_SeetaTrackingFaceInfo) *SeetaTrackingFaceInfo {
	return &SeetaTrackingFaceInfo{
		Postion:  newSeetaRect(seetaTrackingFaceInfo.pos),
		Score:    float32(seetaTrackingFaceInfo.score),
		Frame_NO: int(seetaTrackingFaceInfo.frame_no),
		PID:      int(seetaTrackingFaceInfo.PID),
		Step:     int(seetaTrackingFaceInfo.step),
	}
}

// SeetaPointInfo 人脸特征点信息
type SeetaPointInfo struct {
	PointCount int // 特征点数
	Points     []C.struct_SeetaPointF
	Masks      []int
}

func NewSeetaPointInfo(pointCount int) *SeetaPointInfo {
	return &SeetaPointInfo{
		PointCount: pointCount,
		Points:     make([]C.struct_SeetaPointF, pointCount),
		Masks:      make([]int, pointCount),
	}
}

func (s *SeetaPointInfo) getCSeetaPointFArray() *C.struct_SeetaPointF {
	return &s.Points[0]
}

// Mask 是否佩戴口罩
func (s *SeetaPointInfo) Mask() bool {
	maskCount := 0
	for i := 2; i < len(s.Masks); i++ {
		maskCount += s.Masks[i]
	}
	return maskCount >= 2
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
