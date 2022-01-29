#include "FaceDetector.h"
#include "FaceDetector.c.h"

#include <iostream>
#include <vector>

// using namespace seeta::SEETA_FACE_DETECTOR_NAMESPACE_VERSION;

facedetector *newFaceDetector(SeetaModelSetting setting)
{
    facedetector *fd = (facedetector *)calloc(1, sizeof(facedetector));
    try
    {
        seeta::ModelSetting set;
        set.append("face_detector.csta");

        seeta::FaceDetector *cppfd = new seeta::FaceDetector(set);
        fd->fd = (void *)cppfd;
    }
    catch (const std::exception &e)
    {
        std::cerr << e.what() << '\n';
    }
    return fd;
}
