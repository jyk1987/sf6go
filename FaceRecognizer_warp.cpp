#include "FaceRecognizer.h"
#include "FaceRecognizer_warp.h"

#include <iostream>

facerecognizer *facerecognizer_new(char *model)
{
    facerecognizer *fr = (facerecognizer *)calloc(1, sizeof(facerecognizer));
    try
    {
        // 声明模型配置
        seeta::ModelSetting setting;
        setting.device = SeetaDevice::SEETA_DEVICE_CPU;
        // 增加模型路径
        setting.append(model);
        // 构造特征检测器C++对象
        seeta::FaceRecognizer *cls = new seeta::FaceRecognizer(setting);
        // 保存人脸识别器对象指针
        fr->cls = (void *)cls;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fr;
}

void facerecognizer_setProperty(facerecognizer *fr, int property, double value)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    cls->set(seeta::FaceRecognizer::Property(property), value);
}

double facerecognizer_getProperty(facerecognizer *fr, int property)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    return cls->get(seeta::FaceRecognizer::Property(property));
}

int facerecognizer_GetCropFaceWidthV2(facerecognizer *fr)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    return cls->GetCropFaceWidthV2();
}
int facerecognizer_GetCropFaceHeightV2(facerecognizer *fr)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    return cls->GetCropFaceHeightV2();
}
int facerecognizer_GetCropFaceChannelsV2(facerecognizer *fr)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    return cls->GetCropFaceChannelsV2();
}

int facerecognizer_Extract(facerecognizer *fr, const SeetaImageData image, const SeetaPointF *points, float *features)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    return cls->Extract(image, points, features);
}

int facerecognizer_GetExtractFeatureSize(facerecognizer *fr)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    return cls->GetExtractFeatureSize();
}

float facerecognizer_CalculateSimilarity(facerecognizer *fr, const float *features1, const float *features2)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    return cls->CalculateSimilarity(features1, features2);
}

SeetaImageData facerecognizer_CropFaceV2(facerecognizer *fr, const SeetaImageData image, const SeetaPointF *points)
{
    seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
    SeetaImageData face;
    face.width = cls->GetCropFaceWidthV2();
    face.height = cls->GetCropFaceHeightV2();
    face.channels = cls->GetCropFaceChannelsV2();
    unsigned char data[face.width * face.height * face.channels];
    face.data = &data[0];
    cls->CropFaceV2(image, points, face);
    return face;
}

void facerecognizer_free(facerecognizer *fr)
{
    if (fr)
    {
        if (fr->cls)
        {
            seeta::FaceRecognizer *cls = (seeta::FaceRecognizer *)fr->cls;
            delete cls;
            fr->cls = nullptr;
        }
        free(fr);
    }
}