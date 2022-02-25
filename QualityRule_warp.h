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

    typedef struct qualityrule
    {
        void *brightness_cls;
    } qualityrule;

    qualityrule *qualityrule_new();
    CQualityResult qualityrule_CheckBrightness(qualityrule *qr,
                                               const SeetaImageData image,
                                               const SeetaRect face,
                                               const SeetaPointF *points,
                                               const int32_t N);
    void qualityrule_SetBrightnessValues(qualityrule *qr, float v0, float v1, float v2, float v3);

    void qualityrule_free(qualityrule *qr);
    // int maskdetector_detect(maskdetector *md, const SeetaImageData image, const SeetaRect face, float *score);

#ifdef __cplusplus
}
#endif