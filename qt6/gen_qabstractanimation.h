#pragma once
#ifndef MIQT_QT6_GEN_QABSTRACTANIMATION_H
#define MIQT_QT6_GEN_QABSTRACTANIMATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractAnimation;
class QAnimationDriver;
class QAnimationGroup;
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTimerEvent;
#else
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QAnimationDriver QAnimationDriver;
typedef struct QAnimationGroup QAnimationGroup;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
#endif

QMetaObject* QAbstractAnimation_MetaObject(const QAbstractAnimation* self);
void* QAbstractAnimation_Metacast(QAbstractAnimation* self, const char* param1);
struct miqt_string QAbstractAnimation_Tr(const char* s);
int QAbstractAnimation_State(const QAbstractAnimation* self);
QAnimationGroup* QAbstractAnimation_Group(const QAbstractAnimation* self);
int QAbstractAnimation_Direction(const QAbstractAnimation* self);
void QAbstractAnimation_SetDirection(QAbstractAnimation* self, int direction);
int QAbstractAnimation_CurrentTime(const QAbstractAnimation* self);
int QAbstractAnimation_CurrentLoopTime(const QAbstractAnimation* self);
int QAbstractAnimation_LoopCount(const QAbstractAnimation* self);
void QAbstractAnimation_SetLoopCount(QAbstractAnimation* self, int loopCount);
int QAbstractAnimation_CurrentLoop(const QAbstractAnimation* self);
int QAbstractAnimation_Duration(const QAbstractAnimation* self);
int QAbstractAnimation_TotalDuration(const QAbstractAnimation* self);
void QAbstractAnimation_Finished(QAbstractAnimation* self);
void QAbstractAnimation_connect_Finished(QAbstractAnimation* self, intptr_t slot);
void QAbstractAnimation_StateChanged(QAbstractAnimation* self, int newState, int oldState);
void QAbstractAnimation_connect_StateChanged(QAbstractAnimation* self, intptr_t slot);
void QAbstractAnimation_CurrentLoopChanged(QAbstractAnimation* self, int currentLoop);
void QAbstractAnimation_connect_CurrentLoopChanged(QAbstractAnimation* self, intptr_t slot);
void QAbstractAnimation_DirectionChanged(QAbstractAnimation* self, int param1);
void QAbstractAnimation_connect_DirectionChanged(QAbstractAnimation* self, intptr_t slot);
void QAbstractAnimation_Start(QAbstractAnimation* self);
void QAbstractAnimation_Pause(QAbstractAnimation* self);
void QAbstractAnimation_Resume(QAbstractAnimation* self);
void QAbstractAnimation_SetPaused(QAbstractAnimation* self, bool paused);
void QAbstractAnimation_Stop(QAbstractAnimation* self);
void QAbstractAnimation_SetCurrentTime(QAbstractAnimation* self, int msecs);
bool QAbstractAnimation_Event(QAbstractAnimation* self, QEvent* event);
void QAbstractAnimation_UpdateCurrentTime(QAbstractAnimation* self, int currentTime);
void QAbstractAnimation_UpdateState(QAbstractAnimation* self, int newState, int oldState);
void QAbstractAnimation_UpdateDirection(QAbstractAnimation* self, int direction);
struct miqt_string QAbstractAnimation_Tr2(const char* s, const char* c);
struct miqt_string QAbstractAnimation_Tr3(const char* s, const char* c, int n);
void QAbstractAnimation_Start1(QAbstractAnimation* self, int policy);
void QAbstractAnimation_Delete(QAbstractAnimation* self, bool isSubclass);

void QAnimationDriver_new(QAnimationDriver** outptr_QAnimationDriver, QObject** outptr_QObject);
void QAnimationDriver_new2(QObject* parent, QAnimationDriver** outptr_QAnimationDriver, QObject** outptr_QObject);
QMetaObject* QAnimationDriver_MetaObject(const QAnimationDriver* self);
void* QAnimationDriver_Metacast(QAnimationDriver* self, const char* param1);
struct miqt_string QAnimationDriver_Tr(const char* s);
void QAnimationDriver_Advance(QAnimationDriver* self);
void QAnimationDriver_Install(QAnimationDriver* self);
void QAnimationDriver_Uninstall(QAnimationDriver* self);
bool QAnimationDriver_IsRunning(const QAnimationDriver* self);
long long QAnimationDriver_Elapsed(const QAnimationDriver* self);
void QAnimationDriver_Started(QAnimationDriver* self);
void QAnimationDriver_connect_Started(QAnimationDriver* self, intptr_t slot);
void QAnimationDriver_Stopped(QAnimationDriver* self);
void QAnimationDriver_connect_Stopped(QAnimationDriver* self, intptr_t slot);
void QAnimationDriver_Start(QAnimationDriver* self);
void QAnimationDriver_Stop(QAnimationDriver* self);
struct miqt_string QAnimationDriver_Tr2(const char* s, const char* c);
struct miqt_string QAnimationDriver_Tr3(const char* s, const char* c, int n);
void QAnimationDriver_override_virtual_Advance(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_Advance(void* self);
void QAnimationDriver_override_virtual_Elapsed(void* self, intptr_t slot);
long long QAnimationDriver_virtualbase_Elapsed(const void* self);
void QAnimationDriver_override_virtual_Start(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_Start(void* self);
void QAnimationDriver_override_virtual_Stop(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_Stop(void* self);
void QAnimationDriver_override_virtual_Event(void* self, intptr_t slot);
bool QAnimationDriver_virtualbase_Event(void* self, QEvent* event);
void QAnimationDriver_override_virtual_EventFilter(void* self, intptr_t slot);
bool QAnimationDriver_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
void QAnimationDriver_override_virtual_TimerEvent(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_TimerEvent(void* self, QTimerEvent* event);
void QAnimationDriver_override_virtual_ChildEvent(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_ChildEvent(void* self, QChildEvent* event);
void QAnimationDriver_override_virtual_CustomEvent(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_CustomEvent(void* self, QEvent* event);
void QAnimationDriver_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
void QAnimationDriver_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QAnimationDriver_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QAnimationDriver_Delete(QAnimationDriver* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
