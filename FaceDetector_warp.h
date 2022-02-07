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

    SeetaFaceInfoArray facedetector_detect(facedetector *fd, SeetaImageData image);

    void facedetector_free(facedetector *fd);

    void facedetector_setProperty(facedetector *fd, int property, double value);

#ifdef __cplusplus
}
#endif