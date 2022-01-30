#pragma once

#include "CStruct.h"
#include "CFaceInfo.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct facedetector
    {
        void *cls;
    } facedetector;

    facedetector *newFaceDetector(char *model);

    SeetaFaceInfoArray detect(facedetector *fd, SeetaImageData image);

#ifdef __cplusplus
}
#endif