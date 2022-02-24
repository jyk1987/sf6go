#include "FaceTracker.h"
#include "FaceTracker_warp.h"
#include "CTrackingFaceInfo.h"

#include <iostream>

facetracker *facetracker_new(char *model, int video_width, int video_height)
{
    facetracker *ft = (facetracker *)calloc(1, sizeof(facetracker));
    try
    {
        // 声明模型配置
        seeta::ModelSetting setting;
        setting.device = SeetaDevice::SEETA_DEVICE_CPU;
        // 增加模型路径
        setting.append(model);
        // 构造特征检测器C++对象
        seeta::FaceTracker *cls = new seeta::FaceTracker(setting, video_width, video_height);
        // 保存人脸识别器对象指针
        ft->cls = (void *)cls;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return ft;
}
void facetracker_free(facetracker *ft)
{
    if (ft)
    {
        if (ft->cls)
        {
            seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
            delete cls;
            ft->cls = nullptr;
        }
        free(ft);
    }
}

SeetaTrackingFaceInfoArray facetracker_Track(facetracker *ft, const SeetaImageData image)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    return cls->Track(image);
}

void facetracker_SetMinFaceSize(facetracker *ft, int size)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    cls->SetMinFaceSize(size);
}

int facetracker_GetMinFaceSize(facetracker *ft)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    return cls->GetMinFaceSize();
}

void facetracker_SetThreshold(facetracker *ft, float thresh)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    cls->SetThreshold(thresh);
}

float facetracker_GetThreshold(facetracker *ft)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    return cls->GetThreshold();
}

void facetracker_SetVideoStable(facetracker *ft, int stable)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    cls->SetVideoStable(stable);
}
int facetracker_GetVideoStable(facetracker *ft)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    return cls->GetVideoStable();
}
void facetracker_SetSingleCalculationThreads(facetracker *ft, int num)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    cls->SetSingleCalculationThreads(num);
}

void facetracker_SetInterval(facetracker *ft, int interval)
{
    seeta::FaceTracker *cls = (seeta::FaceTracker *)ft->cls;
    cls->SetInterval(interval);
}