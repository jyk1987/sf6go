#include "Struct.h"
#include "QualityStructure.h"
#include "QualityOfBrightness.h"
#include "QualityOfClarity.h"
#include "QualityCheck_warp.h"
#include "QualityOfIntegrity.h"
#include <iostream>

CQualityResult CQualityResult_new(seeta::QualityResult *result)
{
    CQualityResult cresult;
    cresult.score = result->score;
    cresult.level = CQualityLevel((int)result->level);
    result = nullptr;
    return cresult;
}

qualitycheck *qualitycheck_new()
{
    qualitycheck *qr = (qualitycheck *)calloc(1, sizeof(qualitycheck));
    // try
    // {
    //     seeta::QualityOfBrightness *brightness = new seeta::QualityOfBrightness();
    //     qr->brightness_cls = (void *)brightness;
    // }
    // catch (const std::exception &e)
    // {
    //     std::cerr << e.what() << '\n';
    // }
    return qr;
}

// 亮度检测
CQualityResult qualitycheck_CheckBrightness(
    qualitycheck *qr, const SeetaImageData image, const SeetaRect face,
    const SeetaPointF *points, const int32_t N)
{
    seeta::QualityOfBrightness *cls;
    if (!qr->brightness_cls)
    {
        cls = new seeta::QualityOfBrightness();
        qr->brightness_cls = (void *)cls;
    }
    else
    {
        cls = (seeta::QualityOfBrightness *)qr->brightness_cls;
    }

    auto result = cls->check(image, face, points, N);
    return CQualityResult_new(&result);
}

void qualitycheck_SetBrightnessValues(qualitycheck *qr, float v0, float v1, float v2, float v3)
{
    seeta::QualityOfBrightness *cls = new seeta::QualityOfBrightness(v0, v1, v2, v3);
    if (qr->brightness_cls)
    {
        delete (seeta::QualityOfBrightness *)qr->brightness_cls;
        qr->brightness_cls = nullptr;
    }
    qr->brightness_cls = (void *)cls;
}

// 清晰度检测
CQualityResult qualitycheck_CheckClarity(qualitycheck *qr,
                                         const SeetaImageData image,
                                         const SeetaRect face,
                                         const SeetaPointF *points,
                                         const int32_t N)
{
    seeta::QualityOfClarity *cls;
    if (!qr->clarity_cls)
    {
        cls = new seeta::QualityOfClarity();
        qr->clarity_cls = (void *)cls;
    }
    else
    {
        cls = (seeta::QualityOfClarity *)qr->clarity_cls;
    }

    auto result = cls->check(image, face, points, N);
    return CQualityResult_new(&result);
}
void qualitycheck_SetClarityValues(qualitycheck *qr, float low, float height)
{
    seeta::QualityOfClarity *cls = new seeta::QualityOfClarity(low, height);
    if (qr->clarity_cls)
    {
        delete (seeta::QualityOfClarity *)qr->clarity_cls;
        qr->clarity_cls = nullptr;
    }
    qr->clarity_cls = (void *)cls;
}

// 完整度检测
CQualityResult qualitycheck_CheckIntegrity(qualitycheck *qr,
                                           const SeetaImageData image,
                                           const SeetaRect face,
                                           const SeetaPointF *points,
                                           const int32_t N)
{
    seeta::QualityOfIntegrity *cls;
    if (!qr->integrity_cls)
    {
        cls = new seeta::QualityOfIntegrity();
        qr->integrity_cls = (void *)cls;
    }
    else
    {
        cls = (seeta::QualityOfIntegrity *)qr->integrity_cls;
    }

    auto result = cls->check(image, face, points, N);
    return CQualityResult_new(&result);
}
void qualitycheck_SetIntegrityValues(qualitycheck *qr, float low, float height)
{
    seeta::QualityOfIntegrity *cls = new seeta::QualityOfIntegrity(low, height);
    if (qr->integrity_cls)
    {
        delete (seeta::QualityOfIntegrity *)qr->integrity_cls;
        qr->integrity_cls = nullptr;
    }
    qr->integrity_cls = (void *)cls;
}

void qualitycheck_free(qualitycheck *qr)
{
    if (qr)
    {
        if (qr->brightness_cls)
        {
            delete (seeta::QualityOfBrightness *)qr->brightness_cls;
            qr->brightness_cls = nullptr;
        }
        if (qr->clarity_cls)
        {
            delete (seeta::QualityOfClarity *)qr->clarity_cls;
            qr->clarity_cls = nullptr;
        }

        free(qr);
    }
}