#pragma once
#ifndef MIQT_QT_GEN_QRUNNABLE_H
#define MIQT_QT_GEN_QRUNNABLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QRunnable;
#else
typedef struct QRunnable QRunnable;
#endif

void QRunnable_Run(QRunnable* self);
bool QRunnable_AutoDelete(const QRunnable* self);
void QRunnable_SetAutoDelete(QRunnable* self, bool _autoDelete);
void QRunnable_OperatorAssign(QRunnable* self, QRunnable* param1);
void QRunnable_Delete(QRunnable* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
