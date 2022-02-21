#include "FaceLandmarker.h"
#include "FaceLandmarker_warp.h"

#include <iostream>

facelandmarker *faceLandmarker_new(char *model)
{
    facelandmarker *fl = (facelandmarker *)calloc(1, sizeof(facelandmarker));
    try
    {
        // 声明模型配置
        seeta::ModelSetting setting;
        setting.device = SeetaDevice::SEETA_DEVICE_CPU;
        // 增加模型路径
        setting.append(model);
        // 构造特征检测器C++对象
        seeta::FaceLandmarker *cls = new seeta::FaceLandmarker(setting);
        // 保存人脸识别器对象指针
        fl->cls = (void *)cls;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fl;
}

int facelandmarker_number(facelandmarker *fl)
{
    seeta::FaceLandmarker *cls = (seeta::FaceLandmarker *)fl->cls;
    return cls->number();
}

void facelandmarker_mark(facelandmarker *fl, const SeetaImageData image, const SeetaRect face, SeetaPointF *points)
{
    seeta::FaceLandmarker *cls = (seeta::FaceLandmarker *)fl->cls;
    cls->mark(image, face, points);
}

// 检测特征点和遮挡情况
void facelandmarker_mark_mask(facelandmarker *fl, const SeetaImageData image, const SeetaRect face, SeetaPointF *points, int32_t *mask)
{
    seeta::FaceLandmarker *cls = (seeta::FaceLandmarker *)fl->cls;
    cls->mark(image, face, points, mask);
}

void facelandmarker_free(facelandmarker *fl)
{
    if (fl)
    {
        if (fl->cls)
        {
            seeta::FaceLandmarker *cls = (seeta::FaceLandmarker *)fl->cls;
            delete cls;
            fl->cls = nullptr;
        }
        free(fl);
    }
}