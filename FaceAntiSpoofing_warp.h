#pragma once

#include "CStruct.h"
#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct faceantispoofing
    {
        void *cls;
    } faceantispoofing;

    faceantispoofing *faceantispoofing_new(char *firstModel);
    faceantispoofing *faceantispoofing_new_v2(char *firstModel, char *secondModel);
    void faceantispoofing_free(faceantispoofing *fas);
    int faceantispoofing_predict(faceantispoofing *fas, const SeetaImageData image, const SeetaRect face, const SeetaPointF *points);

#ifdef __cplusplus
}
#endif