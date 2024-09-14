#ifndef GEN_QSTRINGVIEW_H
#define GEN_QSTRINGVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "binding.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QChar;
class QStringView;
#else
typedef struct QByteArray QByteArray;
typedef struct QChar QChar;
typedef struct QStringView QStringView;
#endif

QStringView* QStringView_new();
struct miqt_string* QStringView_ToString(const QStringView* self);
size_t QStringView_Size(const QStringView* self);
QChar* QStringView_OperatorSubscript(const QStringView* self, size_t n);
QByteArray* QStringView_ToLatin1(const QStringView* self);
QByteArray* QStringView_ToUtf8(const QStringView* self);
QByteArray* QStringView_ToLocal8Bit(const QStringView* self);
struct miqt_array* QStringView_ToUcs4(const QStringView* self);
QChar* QStringView_At(const QStringView* self, size_t n);
void QStringView_Truncate(QStringView* self, size_t n);
void QStringView_Chop(QStringView* self, size_t n);
int QStringView_CompareWithQChar(const QStringView* self, QChar* c);
int QStringView_Compare2(const QStringView* self, QChar* c, uintptr_t cs);
bool QStringView_StartsWithWithQChar(const QStringView* self, QChar* c);
bool QStringView_StartsWith2(const QStringView* self, QChar* c, uintptr_t cs);
bool QStringView_EndsWithWithQChar(const QStringView* self, QChar* c);
bool QStringView_EndsWith2(const QStringView* self, QChar* c, uintptr_t cs);
size_t QStringView_IndexOf(const QStringView* self, QChar* c);
bool QStringView_Contains(const QStringView* self, QChar* c);
size_t QStringView_Count(const QStringView* self, QChar* c);
size_t QStringView_LastIndexOf(const QStringView* self, QChar* c);
bool QStringView_IsRightToLeft(const QStringView* self);
bool QStringView_IsValidUtf16(const QStringView* self);
int16_t QStringView_ToShort(const QStringView* self);
uint16_t QStringView_ToUShort(const QStringView* self);
int QStringView_ToInt(const QStringView* self);
unsigned int QStringView_ToUInt(const QStringView* self);
long QStringView_ToLong(const QStringView* self);
unsigned long QStringView_ToULong(const QStringView* self);
int64_t QStringView_ToLongLong(const QStringView* self);
uint64_t QStringView_ToULongLong(const QStringView* self);
float QStringView_ToFloat(const QStringView* self);
double QStringView_ToDouble(const QStringView* self);
bool QStringView_Empty(const QStringView* self);
QChar* QStringView_Front(const QStringView* self);
QChar* QStringView_Back(const QStringView* self);
bool QStringView_IsNull(const QStringView* self);
bool QStringView_IsEmpty(const QStringView* self);
int QStringView_Length(const QStringView* self);
QChar* QStringView_First(const QStringView* self);
QChar* QStringView_Last(const QStringView* self);
size_t QStringView_IndexOf2(const QStringView* self, QChar* c, size_t from);
size_t QStringView_IndexOf3(const QStringView* self, QChar* c, size_t from, uintptr_t cs);
bool QStringView_Contains2(const QStringView* self, QChar* c, uintptr_t cs);
size_t QStringView_Count2(const QStringView* self, QChar* c, uintptr_t cs);
size_t QStringView_LastIndexOf2(const QStringView* self, QChar* c, size_t from);
size_t QStringView_LastIndexOf3(const QStringView* self, QChar* c, size_t from, uintptr_t cs);
int16_t QStringView_ToShort1(const QStringView* self, bool* ok);
int16_t QStringView_ToShort2(const QStringView* self, bool* ok, int base);
uint16_t QStringView_ToUShort1(const QStringView* self, bool* ok);
uint16_t QStringView_ToUShort2(const QStringView* self, bool* ok, int base);
int QStringView_ToInt1(const QStringView* self, bool* ok);
int QStringView_ToInt2(const QStringView* self, bool* ok, int base);
unsigned int QStringView_ToUInt1(const QStringView* self, bool* ok);
unsigned int QStringView_ToUInt2(const QStringView* self, bool* ok, int base);
long QStringView_ToLong1(const QStringView* self, bool* ok);
long QStringView_ToLong2(const QStringView* self, bool* ok, int base);
unsigned long QStringView_ToULong1(const QStringView* self, bool* ok);
unsigned long QStringView_ToULong2(const QStringView* self, bool* ok, int base);
int64_t QStringView_ToLongLong1(const QStringView* self, bool* ok);
int64_t QStringView_ToLongLong2(const QStringView* self, bool* ok, int base);
uint64_t QStringView_ToULongLong1(const QStringView* self, bool* ok);
uint64_t QStringView_ToULongLong2(const QStringView* self, bool* ok, int base);
float QStringView_ToFloat1(const QStringView* self, bool* ok);
double QStringView_ToDouble1(const QStringView* self, bool* ok);
void QStringView_Delete(QStringView* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
