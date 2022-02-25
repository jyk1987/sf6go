#pragma once

#include "CStruct.h"
#ifdef __cplusplus
extern "C"
{
#endif

    typedef enum CQualityLevel
    {
        LOW = 0,
        MEDIUM = 1,
        HIGH = 2,
    } CQualityLevel;

    typedef struct CQualityResult
    {
        CQualityLevel level; ///< quality level
        float score;         ///< greater means better, no range limit
    } CQualityResult;

    typedef struct qualitycheck
    {
        void *brightness_cls;
    } qualitycheck;

    qualitycheck *qualitycheck_new();
    CQualityResult qualitycheck_CheckBrightness(qualitycheck *qr,
                                                const SeetaImageData image,
                                                const SeetaRect face,
                                                const SeetaPointF *points,
                                                const int32_t N);
    void qualitycheck_SetBrightnessValues(qualitycheck *qr, float v0, float v1, float v2, float v3);

    void qualitycheck_free(qualitycheck *qr);
    // int maskdetector_detect(maskdetector *md, const SeetaImageData image, const SeetaRect face, float *score);

#ifdef __cplusplus
}
#endif