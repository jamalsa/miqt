#pragma once
#ifndef MIQT_QT6_GEN_QGESTURE_H
#define MIQT_QT6_GEN_QGESTURE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QGesture;
class QGestureEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QPanGesture;
class QPinchGesture;
class QPointF;
class QSwipeGesture;
class QTapAndHoldGesture;
class QTapGesture;
class QTimerEvent;
class QWidget;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QGesture QGesture;
typedef struct QGestureEvent QGestureEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPanGesture QPanGesture;
typedef struct QPinchGesture QPinchGesture;
typedef struct QPointF QPointF;
typedef struct QSwipeGesture QSwipeGesture;
typedef struct QTapAndHoldGesture QTapAndHoldGesture;
typedef struct QTapGesture QTapGesture;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
#endif

void QGesture_new(QGesture** outptr_QGesture, QObject** outptr_QObject);
void QGesture_new2(QObject* parent, QGesture** outptr_QGesture, QObject** outptr_QObject);
QMetaObject* QGesture_MetaObject(const QGesture* self);
void* QGesture_Metacast(QGesture* self, const char* param1);
struct miqt_string QGesture_Tr(const char* s);
int QGesture_GestureType(const QGesture* self);
int QGesture_State(const QGesture* self);
QPointF* QGesture_HotSpot(const QGesture* self);
void QGesture_SetHotSpot(QGesture* self, QPointF* value);
bool QGesture_HasHotSpot(const QGesture* self);
void QGesture_UnsetHotSpot(QGesture* self);
void QGesture_SetGestureCancelPolicy(QGesture* self, int policy);
int QGesture_GestureCancelPolicy(const QGesture* self);
struct miqt_string QGesture_Tr2(const char* s, const char* c);
struct miqt_string QGesture_Tr3(const char* s, const char* c, int n);
void QGesture_override_virtual_Event(void* self, intptr_t slot);
bool QGesture_virtualbase_Event(void* self, QEvent* event);
void QGesture_override_virtual_EventFilter(void* self, intptr_t slot);
bool QGesture_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
void QGesture_override_virtual_TimerEvent(void* self, intptr_t slot);
void QGesture_virtualbase_TimerEvent(void* self, QTimerEvent* event);
void QGesture_override_virtual_ChildEvent(void* self, intptr_t slot);
void QGesture_virtualbase_ChildEvent(void* self, QChildEvent* event);
void QGesture_override_virtual_CustomEvent(void* self, intptr_t slot);
void QGesture_virtualbase_CustomEvent(void* self, QEvent* event);
void QGesture_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QGesture_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
void QGesture_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QGesture_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QGesture_Delete(QGesture* self, bool isSubclass);

void QPanGesture_new(QPanGesture** outptr_QPanGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
void QPanGesture_new2(QObject* parent, QPanGesture** outptr_QPanGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
QMetaObject* QPanGesture_MetaObject(const QPanGesture* self);
void* QPanGesture_Metacast(QPanGesture* self, const char* param1);
struct miqt_string QPanGesture_Tr(const char* s);
QPointF* QPanGesture_LastOffset(const QPanGesture* self);
QPointF* QPanGesture_Offset(const QPanGesture* self);
QPointF* QPanGesture_Delta(const QPanGesture* self);
double QPanGesture_Acceleration(const QPanGesture* self);
void QPanGesture_SetLastOffset(QPanGesture* self, QPointF* value);
void QPanGesture_SetOffset(QPanGesture* self, QPointF* value);
void QPanGesture_SetAcceleration(QPanGesture* self, double value);
struct miqt_string QPanGesture_Tr2(const char* s, const char* c);
struct miqt_string QPanGesture_Tr3(const char* s, const char* c, int n);
void QPanGesture_Delete(QPanGesture* self, bool isSubclass);

void QPinchGesture_new(QPinchGesture** outptr_QPinchGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
void QPinchGesture_new2(QObject* parent, QPinchGesture** outptr_QPinchGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
QMetaObject* QPinchGesture_MetaObject(const QPinchGesture* self);
void* QPinchGesture_Metacast(QPinchGesture* self, const char* param1);
struct miqt_string QPinchGesture_Tr(const char* s);
int QPinchGesture_TotalChangeFlags(const QPinchGesture* self);
void QPinchGesture_SetTotalChangeFlags(QPinchGesture* self, int value);
int QPinchGesture_ChangeFlags(const QPinchGesture* self);
void QPinchGesture_SetChangeFlags(QPinchGesture* self, int value);
QPointF* QPinchGesture_StartCenterPoint(const QPinchGesture* self);
QPointF* QPinchGesture_LastCenterPoint(const QPinchGesture* self);
QPointF* QPinchGesture_CenterPoint(const QPinchGesture* self);
void QPinchGesture_SetStartCenterPoint(QPinchGesture* self, QPointF* value);
void QPinchGesture_SetLastCenterPoint(QPinchGesture* self, QPointF* value);
void QPinchGesture_SetCenterPoint(QPinchGesture* self, QPointF* value);
double QPinchGesture_TotalScaleFactor(const QPinchGesture* self);
double QPinchGesture_LastScaleFactor(const QPinchGesture* self);
double QPinchGesture_ScaleFactor(const QPinchGesture* self);
void QPinchGesture_SetTotalScaleFactor(QPinchGesture* self, double value);
void QPinchGesture_SetLastScaleFactor(QPinchGesture* self, double value);
void QPinchGesture_SetScaleFactor(QPinchGesture* self, double value);
double QPinchGesture_TotalRotationAngle(const QPinchGesture* self);
double QPinchGesture_LastRotationAngle(const QPinchGesture* self);
double QPinchGesture_RotationAngle(const QPinchGesture* self);
void QPinchGesture_SetTotalRotationAngle(QPinchGesture* self, double value);
void QPinchGesture_SetLastRotationAngle(QPinchGesture* self, double value);
void QPinchGesture_SetRotationAngle(QPinchGesture* self, double value);
struct miqt_string QPinchGesture_Tr2(const char* s, const char* c);
struct miqt_string QPinchGesture_Tr3(const char* s, const char* c, int n);
void QPinchGesture_Delete(QPinchGesture* self, bool isSubclass);

void QSwipeGesture_new(QSwipeGesture** outptr_QSwipeGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
void QSwipeGesture_new2(QObject* parent, QSwipeGesture** outptr_QSwipeGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
QMetaObject* QSwipeGesture_MetaObject(const QSwipeGesture* self);
void* QSwipeGesture_Metacast(QSwipeGesture* self, const char* param1);
struct miqt_string QSwipeGesture_Tr(const char* s);
int QSwipeGesture_HorizontalDirection(const QSwipeGesture* self);
int QSwipeGesture_VerticalDirection(const QSwipeGesture* self);
double QSwipeGesture_SwipeAngle(const QSwipeGesture* self);
void QSwipeGesture_SetSwipeAngle(QSwipeGesture* self, double value);
struct miqt_string QSwipeGesture_Tr2(const char* s, const char* c);
struct miqt_string QSwipeGesture_Tr3(const char* s, const char* c, int n);
void QSwipeGesture_Delete(QSwipeGesture* self, bool isSubclass);

void QTapGesture_new(QTapGesture** outptr_QTapGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
void QTapGesture_new2(QObject* parent, QTapGesture** outptr_QTapGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
QMetaObject* QTapGesture_MetaObject(const QTapGesture* self);
void* QTapGesture_Metacast(QTapGesture* self, const char* param1);
struct miqt_string QTapGesture_Tr(const char* s);
QPointF* QTapGesture_Position(const QTapGesture* self);
void QTapGesture_SetPosition(QTapGesture* self, QPointF* pos);
struct miqt_string QTapGesture_Tr2(const char* s, const char* c);
struct miqt_string QTapGesture_Tr3(const char* s, const char* c, int n);
void QTapGesture_Delete(QTapGesture* self, bool isSubclass);

void QTapAndHoldGesture_new(QTapAndHoldGesture** outptr_QTapAndHoldGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
void QTapAndHoldGesture_new2(QObject* parent, QTapAndHoldGesture** outptr_QTapAndHoldGesture, QGesture** outptr_QGesture, QObject** outptr_QObject);
QMetaObject* QTapAndHoldGesture_MetaObject(const QTapAndHoldGesture* self);
void* QTapAndHoldGesture_Metacast(QTapAndHoldGesture* self, const char* param1);
struct miqt_string QTapAndHoldGesture_Tr(const char* s);
QPointF* QTapAndHoldGesture_Position(const QTapAndHoldGesture* self);
void QTapAndHoldGesture_SetPosition(QTapAndHoldGesture* self, QPointF* pos);
void QTapAndHoldGesture_SetTimeout(int msecs);
int QTapAndHoldGesture_Timeout();
struct miqt_string QTapAndHoldGesture_Tr2(const char* s, const char* c);
struct miqt_string QTapAndHoldGesture_Tr3(const char* s, const char* c, int n);
void QTapAndHoldGesture_Delete(QTapAndHoldGesture* self, bool isSubclass);

void QGestureEvent_new(struct miqt_array /* of QGesture* */  gestures, QGestureEvent** outptr_QGestureEvent, QEvent** outptr_QEvent);
void QGestureEvent_new2(QGestureEvent* param1, QGestureEvent** outptr_QGestureEvent, QEvent** outptr_QEvent);
struct miqt_array /* of QGesture* */  QGestureEvent_Gestures(const QGestureEvent* self);
QGesture* QGestureEvent_Gesture(const QGestureEvent* self, int typeVal);
struct miqt_array /* of QGesture* */  QGestureEvent_ActiveGestures(const QGestureEvent* self);
struct miqt_array /* of QGesture* */  QGestureEvent_CanceledGestures(const QGestureEvent* self);
void QGestureEvent_SetAccepted(QGestureEvent* self, QGesture* param1, bool param2);
void QGestureEvent_Accept(QGestureEvent* self, QGesture* param1);
void QGestureEvent_Ignore(QGestureEvent* self, QGesture* param1);
bool QGestureEvent_IsAccepted(const QGestureEvent* self, QGesture* param1);
void QGestureEvent_SetAccepted2(QGestureEvent* self, int param1, bool param2);
void QGestureEvent_AcceptWithQtGestureType(QGestureEvent* self, int param1);
void QGestureEvent_IgnoreWithQtGestureType(QGestureEvent* self, int param1);
bool QGestureEvent_IsAcceptedWithQtGestureType(const QGestureEvent* self, int param1);
void QGestureEvent_SetWidget(QGestureEvent* self, QWidget* widget);
QWidget* QGestureEvent_Widget(const QGestureEvent* self);
QPointF* QGestureEvent_MapToGraphicsScene(const QGestureEvent* self, QPointF* gesturePoint);
void QGestureEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QGestureEvent_virtualbase_SetAccepted(void* self, bool accepted);
void QGestureEvent_override_virtual_Clone(void* self, intptr_t slot);
QEvent* QGestureEvent_virtualbase_Clone(const void* self);
void QGestureEvent_Delete(QGestureEvent* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
