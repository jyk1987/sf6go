#include "FaceDetector.h"
#include "FaceDetector_warp.h"

#include <iostream>
#include <vector>

// using namespace seeta::SEETA_FACE_DETECTOR_NAMESPACE_VERSION;

facedetector *newFaceDetector(char *model)
{
    facedetector *fd = (facedetector *)calloc(1, sizeof(facedetector));
    try
    {
        seeta::ModelSetting setting;
        setting.append(model);
        seeta::FaceDetector *cppfd = new seeta::FaceDetector(setting);
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
