#pragma once

#include "CStruct.h"
#include "CTrackingFaceInfo.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct facetracker
    {
        void *cls;
    } facetracker;

    facetracker *facetracker_new(char *model, int video_width, int video_height);
    void facetracker_free(facetracker *ft);

    SeetaTrackingFaceInfoArray facetracker_Track(facetracker *ft, const SeetaImageData image);

    void facetracker_SetMinFaceSize(facetracker *ft, int size);

    int facetracker_GetMinFaceSize(facetracker *ft);

    void facetracker_SetThreshold(facetracker *ft, float thresh);

    float facetracker_GetThreshold(facetracker *ft);

    void facetracker_SetVideoStable(facetracker *ft, int stable);
    int facetracker_GetVideoStable(facetracker *ft);

    void facetracker_SetSingleCalculationThreads(facetracker *ft, int num);

    void facetracker_SetInterval(facetracker *ft, int interval);
    void facetracker_Reset(facetracker *ft);
#ifdef __cplusplus
}
#endif