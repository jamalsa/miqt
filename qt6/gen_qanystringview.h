#pragma once
#ifndef MIQT_QT6_GEN_QANYSTRINGVIEW_H
#define MIQT_QT6_GEN_QANYSTRINGVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAnyStringView;
class QByteArray;
class QChar;
#else
typedef struct QAnyStringView QAnyStringView;
typedef struct QByteArray QByteArray;
typedef struct QChar QChar;
#endif

void QAnyStringView_new(QAnyStringView** outptr_QAnyStringView);
void QAnyStringView_new2(struct miqt_string str, QAnyStringView** outptr_QAnyStringView);
void QAnyStringView_new3(struct miqt_string str, QAnyStringView** outptr_QAnyStringView);
void QAnyStringView_new4(QChar* c, QAnyStringView** outptr_QAnyStringView);
void QAnyStringView_new5(QAnyStringView* param1, QAnyStringView** outptr_QAnyStringView);
struct miqt_string QAnyStringView_ToString(const QAnyStringView* self);
ptrdiff_t QAnyStringView_Size(const QAnyStringView* self);
const void* QAnyStringView_Data(const QAnyStringView* self);
int QAnyStringView_Compare(QAnyStringView* lhs, QAnyStringView* rhs);
bool QAnyStringView_Equal(QAnyStringView* lhs, QAnyStringView* rhs);
QChar* QAnyStringView_Front(const QAnyStringView* self);
QChar* QAnyStringView_Back(const QAnyStringView* self);
bool QAnyStringView_Empty(const QAnyStringView* self);
ptrdiff_t QAnyStringView_SizeBytes(const QAnyStringView* self);
bool QAnyStringView_IsNull(const QAnyStringView* self);
bool QAnyStringView_IsEmpty(const QAnyStringView* self);
ptrdiff_t QAnyStringView_Length(const QAnyStringView* self);
int QAnyStringView_Compare3(QAnyStringView* lhs, QAnyStringView* rhs, int cs);
void QAnyStringView_Delete(QAnyStringView* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
