#pragma once
#include "FaceDetector.h"

#ifdef __cplusplus
extern "C"
{
#endif
    struct facedetector
    {
        void *fd;
    };

    facedetector *newFaceDetector(SeetaModelSetting setting);

#ifdef __cplusplus
}
#endif