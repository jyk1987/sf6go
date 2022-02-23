#pragma once

#include "CStruct.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct SeetaTrackingFaceInfo
    {
        SeetaRect pos;
        float score;

        int frame_no;
        int PID;
        int step;
    } SeetaTrackingFaceInfo;

    typedef struct SeetaTrackingFaceInfoArray
    {
        struct SeetaTrackingFaceInfo *data;
        int size;
    } SeetaTrackingFaceInfoArray;

#ifdef __cplusplus
}
#endif
