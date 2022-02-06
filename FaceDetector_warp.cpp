#include "FaceDetector.h"
#include "FaceDetector_warp.h"

#include <iostream>
#include <vector>

// using namespace seeta::SEETA_FACE_DETECTOR_NAMESPACE_VERSION;

facedetector *newFaceDetector(char *model)
{
    // 分配一个人脸识别器结构内存
    facedetector *fd = (facedetector *)calloc(1, sizeof(facedetector));
    try
    {
        // 声明模型配置
        seeta::ModelSetting setting;
        // 增加模型路径
        setting.append(model);
        // 构造一个人脸识别器C++对象
        seeta::FaceDetector *cppfd = new seeta::FaceDetector(setting);
        // 保存人脸识别其对象指针
        fd->cls = (void *)cppfd;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fd;
}

SeetaFaceInfoArray detect(facedetector *fd, SeetaImageData image)
{
    seeta::FaceDetector *cls = (seeta::FaceDetector *)fd->cls;
    return cls->detect(image);
}

// 释放人脸识别器结构和保存的C++对象的内存
void facedetector_free(facedetector *fd)
{
    if (fd)
    {
        if (fd->cls)
        {
            seeta::FaceDetector *cls = (seeta::FaceDetector *)fd->cls;
            delete cls;
            fd->cls = NULL;
        }
        free(fd);
    }
}
