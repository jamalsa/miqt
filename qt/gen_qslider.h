#pragma once
#ifndef MIQT_QT_GEN_QSLIDER_H
#define MIQT_QT_GEN_QSLIDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractSlider;
class QEvent;
class QKeyEvent;
class QMetaObject;
class QMouseEvent;
class QObject;
class QPaintDevice;
class QPaintEvent;
class QSize;
class QSlider;
class QTimerEvent;
class QWheelEvent;
class QWidget;
#else
typedef struct QAbstractSlider QAbstractSlider;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QSize QSize;
typedef struct QSlider QSlider;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

void QSlider_new(QWidget* parent, QSlider** outptr_QSlider, QAbstractSlider** outptr_QAbstractSlider, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
void QSlider_new2(QSlider** outptr_QSlider, QAbstractSlider** outptr_QAbstractSlider, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
void QSlider_new3(int orientation, QSlider** outptr_QSlider, QAbstractSlider** outptr_QAbstractSlider, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
void QSlider_new4(int orientation, QWidget* parent, QSlider** outptr_QSlider, QAbstractSlider** outptr_QAbstractSlider, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
QMetaObject* QSlider_MetaObject(const QSlider* self);
void* QSlider_Metacast(QSlider* self, const char* param1);
struct miqt_string QSlider_Tr(const char* s);
struct miqt_string QSlider_TrUtf8(const char* s);
QSize* QSlider_SizeHint(const QSlider* self);
QSize* QSlider_MinimumSizeHint(const QSlider* self);
void QSlider_SetTickPosition(QSlider* self, int position);
int QSlider_TickPosition(const QSlider* self);
void QSlider_SetTickInterval(QSlider* self, int ti);
int QSlider_TickInterval(const QSlider* self);
bool QSlider_Event(QSlider* self, QEvent* event);
void QSlider_PaintEvent(QSlider* self, QPaintEvent* ev);
void QSlider_MousePressEvent(QSlider* self, QMouseEvent* ev);
void QSlider_MouseReleaseEvent(QSlider* self, QMouseEvent* ev);
void QSlider_MouseMoveEvent(QSlider* self, QMouseEvent* ev);
struct miqt_string QSlider_Tr2(const char* s, const char* c);
struct miqt_string QSlider_Tr3(const char* s, const char* c, int n);
struct miqt_string QSlider_TrUtf82(const char* s, const char* c);
struct miqt_string QSlider_TrUtf83(const char* s, const char* c, int n);
void QSlider_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QSlider_virtualbase_SizeHint(const void* self);
void QSlider_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QSlider_virtualbase_MinimumSizeHint(const void* self);
void QSlider_override_virtual_Event(void* self, intptr_t slot);
bool QSlider_virtualbase_Event(void* self, QEvent* event);
void QSlider_override_virtual_PaintEvent(void* self, intptr_t slot);
void QSlider_virtualbase_PaintEvent(void* self, QPaintEvent* ev);
void QSlider_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QSlider_virtualbase_MousePressEvent(void* self, QMouseEvent* ev);
void QSlider_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QSlider_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* ev);
void QSlider_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QSlider_virtualbase_MouseMoveEvent(void* self, QMouseEvent* ev);
void QSlider_override_virtual_SliderChange(void* self, intptr_t slot);
void QSlider_virtualbase_SliderChange(void* self, int change);
void QSlider_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QSlider_virtualbase_KeyPressEvent(void* self, QKeyEvent* ev);
void QSlider_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSlider_virtualbase_TimerEvent(void* self, QTimerEvent* param1);
void QSlider_override_virtual_WheelEvent(void* self, intptr_t slot);
void QSlider_virtualbase_WheelEvent(void* self, QWheelEvent* e);
void QSlider_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QSlider_virtualbase_ChangeEvent(void* self, QEvent* e);
void QSlider_Delete(QSlider* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
