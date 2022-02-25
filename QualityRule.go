package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -ltennis -lSeetaAuthorize -lSeetaQualityAssessor300
// #include <stdlib.h>
// #include "CStruct.h"
// #include "QualityRule_warp.h"
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

type QualityRule struct {
	ptr *C.struct_qualityrule
}

func NewQualityRule() *QualityRule {
	return &QualityRule{
		ptr: C.qualityrule_new(),
	}
}

func (s *QualityRule) Close() {
	C.qualityrule_free(s.ptr)
}

func (s *QualityRule) CheckBrightness(img *SeetaImageData, postion *SeetaRect, points *SeetaPointInfo) *QualityResult {
	var cresult C.struct_CQualityResult = C.qualityrule_CheckBrightness(
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

func (s *QualityRule) SetBrightnessValues(v0, v1, v2, v3 float32) {
	C.qualityrule_SetBrightnessValues(s.ptr, C.float(v0), C.float(v1), C.float(v2), C.float(v3))
}
