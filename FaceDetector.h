//
// Created by kier on 19-4-24.
//

#ifndef INC_SEETA_FACEDETECTOR_H
#define INC_SEETA_FACEDETECTOR_H

#include "Struct.h"
#include "CFaceInfo.h"
#include "SeetaFaceDetectorConfig.h"

namespace seeta {
    namespace SEETA_FACE_DETECTOR_NAMESPACE_VERSION {
        class FaceDetector {
        public:
            using self = FaceDetector;

            enum Property {
                PROPERTY_MIN_FACE_SIZE,
                PROPERTY_THRESHOLD,
                PROPERTY_MAX_IMAGE_WIDTH,
                PROPERTY_MAX_IMAGE_HEIGHT,
                PROPERTY_NUMBER_THREADS,

                PROPERTY_ARM_CPU_MODE = 0x101,
            };

            SEETA_API explicit FaceDetector(const SeetaModelSetting &setting);

            SEETA_API ~FaceDetector();

            SEETA_API explicit FaceDetector(const self *other);

            SEETA_API SeetaFaceInfoArray detect(const SeetaImageData &image) const;
            
            std::vector<SeetaFaceInfo> detect_v2(const SeetaImageData &image) const {
                auto faces = this->detect(image);
                return std::vector<SeetaFaceInfo>(faces.data, faces.data + faces.size);
            }

            SEETA_API void set(Property property, double value);

            SEETA_API double get(Property property) const;

        private:
            FaceDetector(const FaceDetector &) = delete;

            const FaceDetector &operator=(const FaceDetector &) = delete;

        private:
            class Implement;

            Implement *m_impl;
        };
    }
    using namespace SEETA_FACE_DETECTOR_NAMESPACE_VERSION;
}

#endif //INC_SEETA_FACEDETECTOR_H
