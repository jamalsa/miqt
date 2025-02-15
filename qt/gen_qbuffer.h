#pragma once
#ifndef MIQT_QT_GEN_QBUFFER_H
#define MIQT_QT_GEN_QBUFFER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBuffer;
class QByteArray;
class QIODevice;
class QMetaMethod;
class QMetaObject;
class QObject;
#else
typedef struct QBuffer QBuffer;
typedef struct QByteArray QByteArray;
typedef struct QIODevice QIODevice;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
#endif

void QBuffer_new(QBuffer** outptr_QBuffer, QIODevice** outptr_QIODevice, QObject** outptr_QObject);
void QBuffer_new2(QObject* parent, QBuffer** outptr_QBuffer, QIODevice** outptr_QIODevice, QObject** outptr_QObject);
QMetaObject* QBuffer_MetaObject(const QBuffer* self);
void* QBuffer_Metacast(QBuffer* self, const char* param1);
struct miqt_string QBuffer_Tr(const char* s);
struct miqt_string QBuffer_TrUtf8(const char* s);
struct miqt_string QBuffer_Buffer(QBuffer* self);
struct miqt_string QBuffer_Buffer2(const QBuffer* self);
void QBuffer_SetData(QBuffer* self, struct miqt_string data);
void QBuffer_SetData2(QBuffer* self, const char* data, int lenVal);
struct miqt_string QBuffer_Data(const QBuffer* self);
bool QBuffer_Open(QBuffer* self, int openMode);
void QBuffer_Close(QBuffer* self);
long long QBuffer_Size(const QBuffer* self);
long long QBuffer_Pos(const QBuffer* self);
bool QBuffer_Seek(QBuffer* self, long long off);
bool QBuffer_AtEnd(const QBuffer* self);
bool QBuffer_CanReadLine(const QBuffer* self);
void QBuffer_ConnectNotify(QBuffer* self, QMetaMethod* param1);
void QBuffer_DisconnectNotify(QBuffer* self, QMetaMethod* param1);
long long QBuffer_ReadData(QBuffer* self, char* data, long long maxlen);
long long QBuffer_WriteData(QBuffer* self, const char* data, long long lenVal);
struct miqt_string QBuffer_Tr2(const char* s, const char* c);
struct miqt_string QBuffer_Tr3(const char* s, const char* c, int n);
struct miqt_string QBuffer_TrUtf82(const char* s, const char* c);
struct miqt_string QBuffer_TrUtf83(const char* s, const char* c, int n);
void QBuffer_override_virtual_Open(void* self, intptr_t slot);
bool QBuffer_virtualbase_Open(void* self, int openMode);
void QBuffer_override_virtual_Close(void* self, intptr_t slot);
void QBuffer_virtualbase_Close(void* self);
void QBuffer_override_virtual_Size(void* self, intptr_t slot);
long long QBuffer_virtualbase_Size(const void* self);
void QBuffer_override_virtual_Pos(void* self, intptr_t slot);
long long QBuffer_virtualbase_Pos(const void* self);
void QBuffer_override_virtual_Seek(void* self, intptr_t slot);
bool QBuffer_virtualbase_Seek(void* self, long long off);
void QBuffer_override_virtual_AtEnd(void* self, intptr_t slot);
bool QBuffer_virtualbase_AtEnd(const void* self);
void QBuffer_override_virtual_CanReadLine(void* self, intptr_t slot);
bool QBuffer_virtualbase_CanReadLine(const void* self);
void QBuffer_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QBuffer_virtualbase_ConnectNotify(void* self, QMetaMethod* param1);
void QBuffer_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QBuffer_virtualbase_DisconnectNotify(void* self, QMetaMethod* param1);
void QBuffer_override_virtual_ReadData(void* self, intptr_t slot);
long long QBuffer_virtualbase_ReadData(void* self, char* data, long long maxlen);
void QBuffer_override_virtual_WriteData(void* self, intptr_t slot);
long long QBuffer_virtualbase_WriteData(void* self, const char* data, long long lenVal);
void QBuffer_override_virtual_IsSequential(void* self, intptr_t slot);
bool QBuffer_virtualbase_IsSequential(const void* self);
void QBuffer_override_virtual_Reset(void* self, intptr_t slot);
bool QBuffer_virtualbase_Reset(void* self);
void QBuffer_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QBuffer_virtualbase_BytesAvailable(const void* self);
void QBuffer_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QBuffer_virtualbase_BytesToWrite(const void* self);
void QBuffer_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QBuffer_virtualbase_WaitForReadyRead(void* self, int msecs);
void QBuffer_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QBuffer_virtualbase_WaitForBytesWritten(void* self, int msecs);
void QBuffer_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QBuffer_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
void QBuffer_Delete(QBuffer* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
