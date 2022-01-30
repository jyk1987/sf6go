#include "FaceDetector.h"
#include "FaceDetector.c.h"

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
        fd->fd = (void *)cppfd;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fd;
}
