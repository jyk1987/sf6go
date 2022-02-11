#pragma once

#include "CStruct.h"
#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct facerecognizer
    {
        void *cls;
    } facerecognizer;

    facerecognizer *facerecognizer_new(char *model);

#ifdef __cplusplus
}
#endif