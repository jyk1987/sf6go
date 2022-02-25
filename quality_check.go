package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -ltennis -lSeetaAuthorize -lSeetaQualityAssessor300
// #include <stdlib.h>
// #include "CStruct.h"
// #include "QualityCheck_warp.h"
import "C"

type QualityLevel uint8

const (
	QualityLevel_LOW    QualityLevel = 0
	QualityLevel_MEDIUM QualityLevel = 1
	QualityLevel_HIGH   QualityLevel = 2
)

type QualityResult struct {
	Level QualityLevel
	Score float32
}

type QualityCheck struct {
	ptr *C.struct_qualitycheck
}

// NewQualityRule 创建质量检测器
func NewQualityCheck() *QualityCheck {
	qr := &QualityCheck{
		ptr: C.qualitycheck_new(),
	}
	qr.SetBrightnessValues(
		QualityCheckBrightnessDefaultValues[0],
		QualityCheckBrightnessDefaultValues[1],
		QualityCheckBrightnessDefaultValues[2],
		QualityCheckBrightnessDefaultValues[3],
	)
	return qr
}

func (s *QualityCheck) Close() {
	C.qualitycheck_free(s.ptr)
}

// CheckBrightness 检测亮度
func (s *QualityCheck) CheckBrightness(img *SeetaImageData, postion *SeetaRect, points *SeetaPointInfo) *QualityResult {
	var cresult C.struct_CQualityResult = C.qualitycheck_CheckBrightness(
		s.ptr, img.getCStruct(),
		postion.getCStruct(),
		points.getCSeetaPointFArray(),
		C.int(points.PointCount),
	)
	result := &QualityResult{
		Score: float32(cresult.score),
		Level: QualityLevel(cresult.level),
	}
	return result
}

// QualityRuleBrightnessDefaultValues默认亮度阈值
var QualityCheckBrightnessDefaultValues []float32 = []float32{70, 100, 320, 230}

// SetBrightnessValues 设置亮度阈值
func (s *QualityCheck) SetBrightnessValues(v0, v1, v2, v3 float32) {
	C.qualitycheck_SetBrightnessValues(s.ptr, C.float(v0), C.float(v1), C.float(v2), C.float(v3))
}
