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
        setting.device = SeetaDevice::SEETA_DEVICE_CPU;
        // 增加模型路径
        setting.append(model);
        // 构造一个人脸识别器C++对象
        seeta::FaceDetector *cppfd = new seeta::FaceDetector(setting);
        // 设置识别器使用的线程数
        // cppfd->set(seeta::FaceDetector::PROPERTY_NUMBER_THREADS, 1);
        // 保存人脸识别其对象指针
        fd->cls = (void *)cppfd;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fd;
}

SeetaFaceInfoArray facedetector_detect(facedetector *fd, SeetaImageData image)
{
    seeta::FaceDetector *cls = (seeta::FaceDetector *)fd->cls;
    return cls->detect(image);
}

void facedetector_setProperty(facedetector *fd, int property, double value)
{
    seeta::FaceDetector *cls = (seeta::FaceDetector *)fd->cls;
    switch (property)
    {
    case 0:
        cls->set(seeta::FaceDetector::PROPERTY_MIN_FACE_SIZE, value);
        break;
    case 1:
        cls->set(seeta::FaceDetector::PROPERTY_THRESHOLD, value);
        break;
    case 2:
        cls->set(seeta::FaceDetector::PROPERTY_MAX_IMAGE_WIDTH, value);
        break;
    case 3:
        cls->set(seeta::FaceDetector::PROPERTY_MAX_IMAGE_HEIGHT, value);
        break;
    case 4:
        cls->set(seeta::FaceDetector::PROPERTY_NUMBER_THREADS, value);
        break;
    case 0x101:
        cls->set(seeta::FaceDetector::PROPERTY_ARM_CPU_MODE, value);
        break;
    default:
        break;
    }
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
