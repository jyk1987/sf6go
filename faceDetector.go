package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -ltennis -lSeetaAuthorize -lSeetaFaceDetector600
// #include <stdlib.h>
// #include "FaceDetector_warp.h"
import "C"

import (
	"log"
	"reflect"
	"sync"
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

func (s *FaceDetector) Detect(imageData *SeetaImageData) []SeetaFaceInfo {
	var result C.struct_SeetaFaceInfoArray = C.detect(s.ptr, *imageData.ptr)
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

	img := gocv.IMRead("duo6.jpeg", gocv.IMReadColor)
	defer img.Close()
	imageData := NewSeetaImageData(img.Cols(), img.Rows(), img.Channels())
	defer imageData.Close()
	err := imageData.SetMat(&img)
	if err != nil {
		log.Panic(err)
	}

	var wait sync.WaitGroup

	for i := 0; i < 1; i++ {
		wait.Add(1)
		go func() {
			fd := NewFaceDetector(model)
			defer fd.Close()
			for j := 0; j < 1; j++ {
				start := time.Now()
				faces := fd.Detect(imageData)
				log.Println("检测人脸", len(faces), "耗时:", time.Since(start))
			}
			wait.Done()
		}()
	}
	wait.Wait()

}
