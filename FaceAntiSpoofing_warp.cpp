#include "Struct.h"
#include "FaceAntiSpoofing.h"
#include "FaceAntiSpoofing_warp.h"
#include <iostream>
faceantispoofing *faceantispoofing_new(char *firstModel)
{
    // 分配一个人脸识别器结构内存
    faceantispoofing *fas = (faceantispoofing *)calloc(1, sizeof(faceantispoofing));
    try
    {
        // 声明模型配置
        seeta::ModelSetting setting;
        setting.device = SeetaDevice::SEETA_DEVICE_CPU;
        // 增加模型路径
        setting.append(firstModel);
        // 构造一个人脸识别器C++对象
        seeta::FaceAntiSpoofing *cppfas = new seeta::FaceAntiSpoofing(setting);
        // 保存人脸识别器对象指针
        fas->cls = (void *)cppfas;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fas;
}
faceantispoofing *faceantispoofing_new_v2(char *firstModel, char *secondModel)
{
    // 分配一个人脸识别器结构内存
    faceantispoofing *fas = (faceantispoofing *)calloc(1, sizeof(faceantispoofing));
    try
    {
        // 声明模型配置
        seeta::ModelSetting setting;
        setting.device = SeetaDevice::SEETA_DEVICE_CPU;
        // 增加模型路径
        setting.append(firstModel);
        setting.append(secondModel);
        // 构造一个人脸识别器C++对象
        seeta::FaceAntiSpoofing *cppfas = new seeta::FaceAntiSpoofing(setting);
        // 保存人脸识别器对象指针
        fas->cls = (void *)cppfas;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fas;
}
void faceantispoofing_free(faceantispoofing *fas)
{
    if (fas)
    {
        if (fas->cls)
        {
            seeta::FaceAntiSpoofing *cls = (seeta::FaceAntiSpoofing *)fas->cls;
            delete cls;
            fas->cls = nullptr;
        }
        free(fas);
    }
}
int faceantispoofing_predict(faceantispoofing *fas, const SeetaImageData image, const SeetaRect face, const SeetaPointF *points)
{
    seeta::FaceAntiSpoofing *cls = (seeta::FaceAntiSpoofing *)fas->cls;
    return cls->Predict(image, face, points);
}