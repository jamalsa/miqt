#ifndef GEN_QCBORSTREAMWRITER_H
#define GEN_QCBORSTREAMWRITER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QCborStreamWriter;
class QIODevice;
#else
typedef struct QByteArray QByteArray;
typedef struct QCborStreamWriter QCborStreamWriter;
typedef struct QIODevice QIODevice;
#endif

QCborStreamWriter* QCborStreamWriter_new(QIODevice* device);
QCborStreamWriter* QCborStreamWriter_new2(QByteArray* data);
void QCborStreamWriter_SetDevice(QCborStreamWriter* self, QIODevice* device);
QIODevice* QCborStreamWriter_Device(QCborStreamWriter* self);
void QCborStreamWriter_Append(QCborStreamWriter* self, unsigned long long u);
void QCborStreamWriter_AppendWithQint64(QCborStreamWriter* self, long long i);
void QCborStreamWriter_AppendWithBa(QCborStreamWriter* self, QByteArray* ba);
void QCborStreamWriter_AppendWithFloat(QCborStreamWriter* self, float f);
void QCborStreamWriter_AppendWithDouble(QCborStreamWriter* self, double d);
void QCborStreamWriter_AppendByteString(QCborStreamWriter* self, const char* data, size_t lenVal);
void QCborStreamWriter_AppendTextString(QCborStreamWriter* self, const char* utf8, size_t lenVal);
void QCborStreamWriter_AppendWithBool(QCborStreamWriter* self, bool b);
void QCborStreamWriter_AppendNull(QCborStreamWriter* self);
void QCborStreamWriter_AppendUndefined(QCborStreamWriter* self);
void QCborStreamWriter_AppendWithInt(QCborStreamWriter* self, int i);
void QCborStreamWriter_AppendWithUint(QCborStreamWriter* self, unsigned int u);
void QCborStreamWriter_Append4(QCborStreamWriter* self, const char* str);
void QCborStreamWriter_StartArray(QCborStreamWriter* self);
void QCborStreamWriter_StartArrayWithCount(QCborStreamWriter* self, unsigned long long count);
bool QCborStreamWriter_EndArray(QCborStreamWriter* self);
void QCborStreamWriter_StartMap(QCborStreamWriter* self);
void QCborStreamWriter_StartMapWithCount(QCborStreamWriter* self, unsigned long long count);
bool QCborStreamWriter_EndMap(QCborStreamWriter* self);
void QCborStreamWriter_Append22(QCborStreamWriter* self, const char* str, size_t size);
void QCborStreamWriter_Delete(QCborStreamWriter* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
