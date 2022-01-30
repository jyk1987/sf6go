#pragma once

// #include "FaceDetector.h"
#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct facedetector
    {
        void *fd;
    } facedetector;

    facedetector *newFaceDetector(char *model);

#ifdef __cplusplus
}
#endif