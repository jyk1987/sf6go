package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSeetaFaceDetector600
// #include <stdlib.h>
// #include "FaceDetector_warp.h"
import "C"

import (
	"log"
	"reflect"
	"time"
	"unsafe"

	"gocv.io/x/gocv"
)

type FaceDetectorProperty int

const (
	FaceDetector_PROPERTY_MIN_FACE_SIZE    FaceDetectorProperty = 0
	FaceDetector_PROPERTY_THRESHOLD        FaceDetectorProperty = 1
	FaceDetector_PROPERTY_MAX_IMAGE_WIDTH  FaceDetectorProperty = 2
	FaceDetector_PROPERTY_MAX_IMAGE_HEIGHT FaceDetectorProperty = 3
	FaceDetector_PROPERTY_NUMBER_THREADS   FaceDetectorProperty = 4
	FaceDetector_PROPERTY_ARM_CPU_MODE     FaceDetectorProperty = 0x101
)

type FaceDetector struct {
	ptr *C.struct_facedetector
}

func NewFaceDetector(model string) *FaceDetector {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	return &FaceDetector{
		ptr: C.faceDetector_new(cs),
	}
}

func (s *FaceDetector) SetProperty(property FaceDetectorProperty, value float64) {
	C.facedetector_setProperty(s.ptr, C.int(property), C.double(value))
}

func (s *FaceDetector) GetProperty(property FaceDetectorProperty) float64 {
	return float64(C.facedetector_getProperty(s.ptr, C.int(property)))
}

func (s *FaceDetector) Detect(imageData *SeetaImageData) []SeetaFaceInfo {
	var result C.struct_SeetaFaceInfoArray = C.facedetector_detect(s.ptr, imageData.ptr)
	var clist []C.struct_SeetaFaceInfo
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&clist))
	arrayLen := int(result.size)
	sliceHeader.Cap = arrayLen
	sliceHeader.Len = arrayLen
	sliceHeader.Data = uintptr(unsafe.Pointer(result.data))

	faceInfoList := make([]SeetaFaceInfo, arrayLen)
	for i := 0; i < arrayLen; i++ {
		faceInfoList[i] = NewSeetaFaceInfo(clist[i])
	}
	return faceInfoList
}

func (s *FaceDetector) Close() {
	C.facedetector_free(s.ptr)
}

func TestFaceDetector() {
	model := "/var/sf6/models/face_detector.csta"
	icount := 4
	imageChan := make(chan *SeetaImageData, icount)
	var work = func() {
		fd := NewFaceDetector(model)
		fd.SetProperty(FaceDetector_PROPERTY_NUMBER_THREADS, 1)
		// log.Println(fd.GetProperty(FaceDetector_PROPERTY_MIN_FACE_SIZE))
		defer fd.Close()
		for {

			img := <-imageChan
			start := time.Now()
			faces := fd.Detect(img)
			log.Println("检测人脸", len(faces), "耗时:", time.Since(start))
		}
	}
	/*
		2个识别器，每个识别器1线程,总线程数2，帧率29.4
		2个识别器，每个识别器2线程,总线程数4，帧率30.3
		3个识别器，每个识别器1线程,总线程数3，帧率32.2
		3个识别器，每个识别器2线程,总线程数6，帧率32.2
		总线程数大于4后有可能长时间高密度工作后造成识别器资源抢夺，造成效能明显下降，
		性能利用率最高是3识别器，每个识别器单线程运行，
		最省资源的是2个识别器，每个识别器单线程运行，但是2*1的方式单帧处理延迟最小，资源占用最小。
	*/
	for i := 0; i < icount; i++ {
		go work()
	}
	begin := time.Now()
	count := 100
	for j := 0; j < count; j++ {
		img := gocv.IMRead("duo6.jpeg", gocv.IMReadColor)
		imageData := NewSeetaImageData(img.Cols(), img.Rows(), img.Channels())
		err := imageData.SetMat(&img)
		img.Close()
		if err != nil {
			log.Panic(err)
		}
		imageChan <- imageData
	}
	mscount := time.Now().UnixMilli()
	log.Println("处理画面", count, "个,用时", time.Since(begin), "fps:", float32(1000)/float32((mscount-begin.UnixMilli())/int64(count)))

}
