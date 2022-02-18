package main

import (
	"log"
	"time"

	"github.com/jyk1987/sf6go"
)

func main() {
	// sf6go.TestCStruct()
	// sf6go.TestFaceDetector()
	// sf6go.TestFaceLandmarker()

	sf6go.InitModelPath("/var/sf6/models")

	fd := sf6go.NewFaceDetector()
	defer fd.Close()

	imageData, err := sf6go.NewSeetaImageDataFromFile("duo6.jpeg")
	if err != nil {
		log.Panic(err)
	}

	start := time.Now()
	begin := start

	faces := fd.Detect(imageData)
	log.Println("检测人脸", len(faces), "个耗时:", time.Since(start))

	fl := sf6go.NewFaceLandmarker(sf6go.ModelType_light)
	defer fl.Close()
	fr := sf6go.NewFaceRecognizer(sf6go.ModelType_light)
	defer fr.Close()
	fas := sf6go.NewFaceAntiSpoofing_v2()
	defer fas.Close()
	md := sf6go.NewMaskDetector()
	defer md.Close()
	for i := 0; i < len(faces); i++ {
		log.Printf("===========识别人脸%v===========", i)
		postion := faces[i].Postion
		start = time.Now()
		isMask := md.Detect(imageData, postion)
		log.Println("口罩检测:", isMask, "耗时:", time.Since(start))
		start = time.Now()
		pointInfo := fl.Mark(imageData, postion)
		log.Println("特征定位耗时:", time.Since(start))
		start = time.Now()
		success, features := fr.Extract(imageData, pointInfo)
		log.Println("特征提取", success, len(features), "耗时:", time.Since(start))
		start = time.Now()
		status := fas.Predict(imageData, postion, pointInfo)
		log.Println("活体检测", status, "耗时:", time.Since(start))
	}
	log.Println("单帧总耗时:", time.Since(begin))
}
