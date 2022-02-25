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
        void *brightness_cls; // 亮度检测器
        void *clarity_cls;    // 清晰度检测器
        void *integrity_cls;  // 完整度检测器
    } qualitycheck;

    qualitycheck *qualitycheck_new();
    // 检测亮度
    CQualityResult qualitycheck_CheckBrightness(qualitycheck *qr,
                                                const SeetaImageData image,
                                                const SeetaRect face,
                                                const SeetaPointF *points,
                                                const int32_t N);
    void qualitycheck_SetBrightnessValues(qualitycheck *qr, float v0, float v1, float v2, float v3);

    // 检测清晰度
    CQualityResult qualitycheck_CheckClarity(qualitycheck *qr,
                                             const SeetaImageData image,
                                             const SeetaRect face,
                                             const SeetaPointF *points,
                                             const int32_t N);
    void qualitycheck_SetClarityValues(qualitycheck *qr, float low, float height);

    // 完整度检测
    CQualityResult qualitycheck_CheckIntegrity(qualitycheck *qr,
                                               const SeetaImageData image,
                                               const SeetaRect face,
                                               const SeetaPointF *points,
                                               const int32_t N);
    void qualitycheck_SetIntegrityValues(qualitycheck *qr, float low, float height);

    void qualitycheck_free(qualitycheck *qr);
    // int maskdetector_detect(maskdetector *md, const SeetaImageData image, const SeetaRect face, float *score);

#ifdef __cplusplus
}
#endif