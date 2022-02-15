#pragma once

#include "CStruct.h"
#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct maskdetector
    {
        void *cls;
    } maskdetector;

    maskdetector *maskdetector_new(char *model);
    void maskdetector_free(maskdetector *md);
    int maskdetector_detect(maskdetector *md, const SeetaImageData image, const SeetaRect face, float *score);

#ifdef __cplusplus
}
#endif