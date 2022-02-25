package main

import (
	"log"
	"time"

	"github.com/jyk1987/sf6go"
)

func main() {
	sf6go.InitModelPath("/var/sf6/models")
	standard_Test()
	facetracker_Test()

}

func facetracker_Test() {
	log.Println("人脸追踪测试开始:", time.Now())
	imageData, err := sf6go.NewSeetaImageDataFromFile("duo6.jpeg")
	if err != nil {
		log.Panic(err)
	}
	log.Println(imageData.GetWidth(), "*", imageData.GetHeight())
	ft := sf6go.NewFaceTracker(imageData.GetWidth(), imageData.GetHeight())
	ft.SetInterval(10)
	log.Println("MinFaceSize:", ft.GetMinFaceSize())
	log.Println("Threshold:", ft.GetThreshold())
	log.Println("VideoStable:", ft.GetVideoStable())
	defer ft.Close()
	for i := 0; i < 2; i++ {
		log.Println("---------------")
		t := time.Now()
		faces := ft.Track(imageData)
		faceCount := len(faces)
		log.Printf("追踪人脸%v个,耗时:%v", faceCount, time.Since(t))

		for j := 0; j < faceCount; j++ {
			face := faces[j]
			log.Printf("Postion:%v,PID:%v,Frame_NO:%v", face.Postion, face.PID, face.Frame_NO)
		}
	}

	log.Println("人脸追踪测试结束:", time.Now())
}

func standard_Test() {
	log.Println("标准测试开始:", time.Now())
	// 人脸检测器
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
	// 人脸特征定位器
	fl := sf6go.NewFaceLandmarker(sf6go.ModelType_light)
	defer fl.Close()
	// 人脸特征提取器
	fr := sf6go.NewFaceRecognizer(sf6go.ModelType_light)
	defer fr.Close()
	// 活体检测器（全局）
	fas := sf6go.NewFaceAntiSpoofing_v2()
	defer fas.Close()
	// 口罩检测器
	md := sf6go.NewMaskDetector()
	defer md.Close()
	// 质量评估器
	qr := sf6go.NewQualityCheck()
	defer qr.Close()
	// 如果使用默认值，一下参数可以不设置
	qr.SetBrightnessValues(70, 100, 210, 230)
	qr.SetClarityValues(0.1, 0.2)
	qr.SetIntegrityValues(10, 1.5)
	for i := 0; i < len(faces); i++ {
		log.Println("---------------------------------------")
		postion := faces[i].Postion
		log.Printf("识别人脸%v,x:%v,y:%v,width:%v,height:%v", i,
			postion.GetX(), postion.GetY(), postion.GetWidth(), postion.GetHeight(),
		)
		start = time.Now()
		isMask := md.Detect(imageData, postion)
		log.Println("口罩检测:", isMask, "耗时:", time.Since(start))
		start = time.Now()
		pointInfo := fl.Mark(imageData, postion)
		log.Println("特征定位耗时:", time.Since(start))
		start = time.Now()
		brightness := qr.CheckBrightness(imageData, postion, pointInfo)
		log.Printf("亮度:%v,检测耗时:%v", brightness.Level, time.Since(start))
		start = time.Now()
		clarity := qr.CheckClarity(imageData, postion, pointInfo)
		log.Printf("清晰度:%v,检测耗时:%v", clarity.Level, time.Since(start))
		start = time.Now()
		integrity := qr.CheckIntegrity(imageData, postion, pointInfo)
		log.Printf("完整度:%v,检测耗时:%v", integrity.Level, time.Since(start))
		start = time.Now()
		pose := qr.CheckPose(imageData, postion, pointInfo)
		log.Printf("姿态:%v,可信度:%v,检测耗时:%v", pose.Level, pose.Score, time.Since(start))
		start = time.Now()
		success, features := fr.Extract(imageData, pointInfo)
		log.Println("特征提取", success, len(features), "耗时:", time.Since(start))
		start = time.Now()
		status := fas.Predict(imageData, postion, pointInfo)
		log.Println("活体检测", status, "耗时:", time.Since(start))
	}
	log.Println("单帧总耗时:", time.Since(begin))
	log.Println("标准测试结束:", time.Now())
}
