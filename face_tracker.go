package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSeetaFaceTracking600
// #include <stdlib.h>
// #include "FaceTracker_warp.h"
// #include "CTrackingFaceInfo.h"
import "C"

type FaceTracker struct {
	ptr *C.struct_facetracker
}

func NewFaceTracker() {}
