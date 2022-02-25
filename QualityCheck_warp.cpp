#include "Struct.h"
#include "QualityStructure.h"
#include "QualityOfBrightness.h"
#include "QualityCheck_warp.h"
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

void qualitycheck_free(qualitycheck *qr)
{
    if (qr)
    {
        if (qr->brightness_cls)
        {
            delete (seeta::QualityOfBrightness *)qr->brightness_cls;
            qr->brightness_cls = nullptr;
        }

        free(qr);
    }
}