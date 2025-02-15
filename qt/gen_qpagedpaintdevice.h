#pragma once
#ifndef MIQT_QT_GEN_QPAGEDPAINTDEVICE_H
#define MIQT_QT_GEN_QPAGEDPAINTDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMarginsF;
class QPageLayout;
class QPageSize;
class QPagedPaintDevice;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPagedPaintDevice__Margins)
typedef QPagedPaintDevice::Margins QPagedPaintDevice__Margins;
#else
class QPagedPaintDevice__Margins;
#endif
class QPaintDevice;
class QSizeF;
#else
typedef struct QMarginsF QMarginsF;
typedef struct QPageLayout QPageLayout;
typedef struct QPageSize QPageSize;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPagedPaintDevice__Margins QPagedPaintDevice__Margins;
typedef struct QPaintDevice QPaintDevice;
typedef struct QSizeF QSizeF;
#endif

bool QPagedPaintDevice_NewPage(QPagedPaintDevice* self);
bool QPagedPaintDevice_SetPageLayout(QPagedPaintDevice* self, QPageLayout* pageLayout);
bool QPagedPaintDevice_SetPageSize(QPagedPaintDevice* self, QPageSize* pageSize);
bool QPagedPaintDevice_SetPageOrientation(QPagedPaintDevice* self, int orientation);
bool QPagedPaintDevice_SetPageMargins(QPagedPaintDevice* self, QMarginsF* margins);
bool QPagedPaintDevice_SetPageMargins2(QPagedPaintDevice* self, QMarginsF* margins, int units);
QPageLayout* QPagedPaintDevice_PageLayout(const QPagedPaintDevice* self);
void QPagedPaintDevice_SetPageSizeWithSize(QPagedPaintDevice* self, int size);
int QPagedPaintDevice_PageSize(const QPagedPaintDevice* self);
void QPagedPaintDevice_SetPageSizeMM(QPagedPaintDevice* self, QSizeF* size);
QSizeF* QPagedPaintDevice_PageSizeMM(const QPagedPaintDevice* self);
void QPagedPaintDevice_SetMargins(QPagedPaintDevice* self, QPagedPaintDevice__Margins* margins);
QPagedPaintDevice__Margins* QPagedPaintDevice_Margins(const QPagedPaintDevice* self);
void QPagedPaintDevice_Delete(QPagedPaintDevice* self, bool isSubclass);

void QPagedPaintDevice__Margins_Delete(QPagedPaintDevice__Margins* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
