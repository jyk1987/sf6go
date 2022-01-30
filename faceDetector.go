package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -ltennis -lSeetaAuthorize -lSeetaFaceDetector600
// #include <stdlib.h>
// #include "FaceDetector.c.h"
import "C"

import (
	"log"
	"time"
	"unsafe"

	"gocv.io/x/gocv"
)

type FaceDetector struct {
	ptr *C.struct_facedetector
}

func NewFaceDetector(model string) *FaceDetector {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))
	return &FaceDetector{
		ptr: C.newFaceDetector(cs),
	}
}

func (s *FaceDetector) detect(imageData *SeetaImageData) {
	var result = C.detect(s.ptr, *imageData.ptr)
	// TODO: 解析返回结构
	log.Println(result)
}

func (s *FaceDetector) Close() {
	C.free(unsafe.Pointer(s.ptr))
}

func TestFaceDetector() {
	model := "/var/sf6/models/face_detector.csta"
	fd := NewFaceDetector(model)
	img := gocv.IMRead("duo6.jpeg", gocv.IMReadColor)
	defer img.Close()
	imageData := NewSeetaImageData(img.Cols(), img.Rows(), img.Channels())
	defer imageData.Close()
	err := imageData.SetMat(&img)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < 10; i++ {
		start := time.Now()
		fd.detect(imageData)
		log.Println("耗时:", time.Since(start))
	}
	defer fd.Close()
}
