#include "Struct.h"
#include "MaskDetector_warp.h"
#include "MaskDetector.h"
#include <iostream>

maskdetector *maskdetector_new(char *model)
{
    maskdetector *fl = (maskdetector *)calloc(1, sizeof(maskdetector));
    try
    {
        // 声明模型配置
        seeta::ModelSetting setting;
        setting.device = SeetaDevice::SEETA_DEVICE_CPU;
        // 增加模型路径
        setting.append(model);
        // 构造特征检测器C++对象
        seeta::MaskDetector *cls = new seeta::MaskDetector(setting);
        // 保存人脸识别器对象指针
        fl->cls = (void *)cls;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fl;
}
void maskdetector_free(maskdetector *md)
{
    if (md)
    {
        if (md->cls)
        {
            seeta::MaskDetector *cls = (seeta::MaskDetector *)md->cls;
            delete cls;
            md->cls = nullptr;
        }
        free(md);
    }
}
int maskdetector_detect(maskdetector *md, const SeetaImageData image, const SeetaRect face, float *score)
{
    seeta::MaskDetector *cls = (seeta::MaskDetector *)md->cls;
    return cls->detect(image, face, score);
}