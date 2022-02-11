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

    facedetector *faceDetector_new(char *model);

    SeetaFaceInfoArray facedetector_detect(facedetector *fd, const SeetaImageData image);

    void facedetector_free(facedetector *fd);

    void facedetector_setProperty(facedetector *fd, int property, double value);

    double facedetector_getProperty(facedetector *fd, int property);

#ifdef __cplusplus
}
#endif