#pragma once

#include "CStruct.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct facerecognizer
    {
        void *cls;
    } facerecognizer;

    facerecognizer *facerecognizer_new(char *model);
    void facerecognizer_free(facerecognizer *fr);

    void facerecognizer_setProperty(facerecognizer *fr, int property, double value);

    double facerecognizer_getProperty(facerecognizer *fr, int property);

    int facerecognizer_GetCropFaceWidthV2(facerecognizer *fr);
    int facerecognizer_GetCropFaceHeightV2(facerecognizer *fr);
    int facerecognizer_GetCropFaceChannelsV2(facerecognizer *fr);
    int facerecognizer_Extract(facerecognizer *fr, const SeetaImageData image, const SeetaPointF *points, float *features);
    int facerecognizer_GetExtractFeatureSize(facerecognizer *fr);

    float facerecognizer_CalculateSimilarity(facerecognizer *fr, const float *features1, const float *features2);

    SeetaImageData facerecognizer_CropFaceV2(facerecognizer *fr, const SeetaImageData image, const SeetaPointF *points);

#ifdef __cplusplus
}
#endif