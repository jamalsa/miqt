#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTTRANSITION_H
#define MIQT_QT_GEN_QABSTRACTTRANSITION_H

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
class QAbstractState;
class QAbstractTransition;
class QEvent;
class QMetaObject;
class QObject;
class QState;
class QStateMachine;
#else
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QAbstractState QAbstractState;
typedef struct QAbstractTransition QAbstractTransition;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QState QState;
typedef struct QStateMachine QStateMachine;
#endif

QMetaObject* QAbstractTransition_MetaObject(const QAbstractTransition* self);
void* QAbstractTransition_Metacast(QAbstractTransition* self, const char* param1);
struct miqt_string QAbstractTransition_Tr(const char* s);
struct miqt_string QAbstractTransition_TrUtf8(const char* s);
QState* QAbstractTransition_SourceState(const QAbstractTransition* self);
QAbstractState* QAbstractTransition_TargetState(const QAbstractTransition* self);
void QAbstractTransition_SetTargetState(QAbstractTransition* self, QAbstractState* target);
struct miqt_array /* of QAbstractState* */  QAbstractTransition_TargetStates(const QAbstractTransition* self);
void QAbstractTransition_SetTargetStates(QAbstractTransition* self, struct miqt_array /* of QAbstractState* */  targets);
int QAbstractTransition_TransitionType(const QAbstractTransition* self);
void QAbstractTransition_SetTransitionType(QAbstractTransition* self, int typeVal);
QStateMachine* QAbstractTransition_Machine(const QAbstractTransition* self);
void QAbstractTransition_AddAnimation(QAbstractTransition* self, QAbstractAnimation* animation);
void QAbstractTransition_RemoveAnimation(QAbstractTransition* self, QAbstractAnimation* animation);
struct miqt_array /* of QAbstractAnimation* */  QAbstractTransition_Animations(const QAbstractTransition* self);
bool QAbstractTransition_EventTest(QAbstractTransition* self, QEvent* event);
void QAbstractTransition_OnTransition(QAbstractTransition* self, QEvent* event);
bool QAbstractTransition_Event(QAbstractTransition* self, QEvent* e);
struct miqt_string QAbstractTransition_Tr2(const char* s, const char* c);
struct miqt_string QAbstractTransition_Tr3(const char* s, const char* c, int n);
struct miqt_string QAbstractTransition_TrUtf82(const char* s, const char* c);
struct miqt_string QAbstractTransition_TrUtf83(const char* s, const char* c, int n);
void QAbstractTransition_Delete(QAbstractTransition* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
