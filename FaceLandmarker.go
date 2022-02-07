package sf6go

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -L./lib -lSeetaFaceLandmarker600
// #include <stdlib.h>
// #include "FaceDetector_warp.h"
import "C"
