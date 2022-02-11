#pragma once

#include "CStruct.h"

#ifdef __cplusplus
extern "C"
{
#endif
    typedef struct facelandmarker
    {
        void *cls;
    } facelandmarker;

    facelandmarker *newFaceLandmarker(char *model);
    void facelandmarker_free(facelandmarker *fl);
    int facelandmarker_number(facelandmarker *fl);
    void facelandmarker_mark(facelandmarker *fl, const SeetaImageData image, const SeetaRect face, SeetaPointF *points, int32_t *mask);
#ifdef __cplusplus
}
#endif