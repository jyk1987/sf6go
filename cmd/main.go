package main

import (
	"log"
	"time"

	"git.gyb3.cn/kuaibang/sf6go"
	"gocv.io/x/gocv"
)

func main() {
	// sf6go.TestCStruct()
	// sf6go.TestFaceDetector()
	// sf6go.TestFaceLandmarker()
	fd := sf6go.NewFaceDetector("/var/sf6/models/face_detector.csta")
	fd.SetProperty(sf6go.FaceDetector_PROPERTY_NUMBER_THREADS, 1)
	img := gocv.IMRead("duo6.jpeg", gocv.IMReadColor)
	imageData := sf6go.NewSeetaImageData(img.Cols(), img.Rows(), img.Channels())
	err := imageData.SetMat(&img)
	img.Close()
	if err != nil {
		log.Println(err)
	}
	start := time.Now()
	faces := fd.Detect(imageData)
	log.Println("检测人脸", len(faces), "耗时:", time.Since(start))
	fl := sf6go.NewFaceLandmarker("/var/sf6/models/face_landmarker_pts5.csta")
	for i := 0; i < len(faces); i++ {
		start := time.Now()
		_ = fl.Mark(imageData, faces[i])
		log.Println("特征定位", i, "耗时:", time.Since(start))
	}
}
