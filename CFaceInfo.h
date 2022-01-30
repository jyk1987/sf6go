#pragma once

#include "CStruct.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct SeetaFaceInfo
    {
        SeetaRect pos;
        float score;
    } SeetaFaceInfo;

    typedef struct SeetaFaceInfoArray
    {
        struct SeetaFaceInfo *data;
        int size;
    } SeetaFaceInfoArray;

#ifdef __cplusplus
}
#endif
