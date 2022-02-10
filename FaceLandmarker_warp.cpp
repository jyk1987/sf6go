#include "FaceLandmarker.h"
#include "FaceLandmarker_warp.h"

#include <iostream>

facelandmarker *newFaceLandmarker(char *model)
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
        seeta::FaceLandmarker *cppfl = new seeta::FaceLandmarker(setting);
        // 保存人脸识别器对象指针
        fl->cls = (void *)cppfl;
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

void facelandmarker_free(facelandmarker *fl)
{
    if (fl)
    {
        if (fl->cls)
        {
            seeta::FaceLandmarker *cls = (seeta::FaceLandmarker *)fl->cls;
            delete cls;
            fl->cls = NULL;
        }
        free(fl);
    }
}