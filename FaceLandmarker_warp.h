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

#ifdef __cplusplus
}
#endif