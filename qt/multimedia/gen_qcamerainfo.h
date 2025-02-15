#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QCAMERAINFO_H
#define MIQT_QT_MULTIMEDIA_GEN_QCAMERAINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QCamera;
class QCameraInfo;
#else
typedef struct QByteArray QByteArray;
typedef struct QCamera QCamera;
typedef struct QCameraInfo QCameraInfo;
#endif

void QCameraInfo_new(QCameraInfo** outptr_QCameraInfo);
void QCameraInfo_new2(QCamera* camera, QCameraInfo** outptr_QCameraInfo);
void QCameraInfo_new3(QCameraInfo* other, QCameraInfo** outptr_QCameraInfo);
void QCameraInfo_new4(struct miqt_string name, QCameraInfo** outptr_QCameraInfo);
void QCameraInfo_OperatorAssign(QCameraInfo* self, QCameraInfo* other);
bool QCameraInfo_OperatorEqual(const QCameraInfo* self, QCameraInfo* other);
bool QCameraInfo_OperatorNotEqual(const QCameraInfo* self, QCameraInfo* other);
bool QCameraInfo_IsNull(const QCameraInfo* self);
struct miqt_string QCameraInfo_DeviceName(const QCameraInfo* self);
struct miqt_string QCameraInfo_Description(const QCameraInfo* self);
int QCameraInfo_Position(const QCameraInfo* self);
int QCameraInfo_Orientation(const QCameraInfo* self);
QCameraInfo* QCameraInfo_DefaultCamera();
struct miqt_array /* of QCameraInfo* */  QCameraInfo_AvailableCameras();
struct miqt_array /* of QCameraInfo* */  QCameraInfo_AvailableCameras1(int position);
void QCameraInfo_Delete(QCameraInfo* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
