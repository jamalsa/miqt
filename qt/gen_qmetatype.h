#ifndef GEN_QMETATYPE_H
#define GEN_QMETATYPE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QDataStream;
class QDebug;
class QMetaObject;
class QMetaType;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtMetaTypePrivate__QAssociativeIterableImpl)
typedef QtMetaTypePrivate::QAssociativeIterableImpl QtMetaTypePrivate__QAssociativeIterableImpl;
#else
class QtMetaTypePrivate__QAssociativeIterableImpl;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtMetaTypePrivate__QPairVariantInterfaceImpl)
typedef QtMetaTypePrivate::QPairVariantInterfaceImpl QtMetaTypePrivate__QPairVariantInterfaceImpl;
#else
class QtMetaTypePrivate__QPairVariantInterfaceImpl;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtMetaTypePrivate__QSequentialIterableImpl)
typedef QtMetaTypePrivate::QSequentialIterableImpl QtMetaTypePrivate__QSequentialIterableImpl;
#else
class QtMetaTypePrivate__QSequentialIterableImpl;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtMetaTypePrivate__VariantData)
typedef QtMetaTypePrivate::VariantData QtMetaTypePrivate__VariantData;
#else
class QtMetaTypePrivate__VariantData;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtMetaTypePrivate__VectorBoolElements)
typedef QtMetaTypePrivate::VectorBoolElements QtMetaTypePrivate__VectorBoolElements;
#else
class QtMetaTypePrivate__VectorBoolElements;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtPrivate__AbstractComparatorFunction)
typedef QtPrivate::AbstractComparatorFunction QtPrivate__AbstractComparatorFunction;
#else
class QtPrivate__AbstractComparatorFunction;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtPrivate__AbstractConverterFunction)
typedef QtPrivate::AbstractConverterFunction QtPrivate__AbstractConverterFunction;
#else
class QtPrivate__AbstractConverterFunction;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtPrivate__AbstractDebugStreamFunction)
typedef QtPrivate::AbstractDebugStreamFunction QtPrivate__AbstractDebugStreamFunction;
#else
class QtPrivate__AbstractDebugStreamFunction;
#endif
#else
typedef struct QByteArray QByteArray;
typedef struct QDataStream QDataStream;
typedef struct QDebug QDebug;
typedef struct QMetaObject QMetaObject;
typedef struct QMetaType QMetaType;
typedef struct QtMetaTypePrivate__QAssociativeIterableImpl QtMetaTypePrivate__QAssociativeIterableImpl;
typedef struct QtMetaTypePrivate__QPairVariantInterfaceImpl QtMetaTypePrivate__QPairVariantInterfaceImpl;
typedef struct QtMetaTypePrivate__QSequentialIterableImpl QtMetaTypePrivate__QSequentialIterableImpl;
typedef struct QtMetaTypePrivate__VariantData QtMetaTypePrivate__VariantData;
typedef struct QtMetaTypePrivate__VectorBoolElements QtMetaTypePrivate__VectorBoolElements;
typedef struct QtPrivate__AbstractComparatorFunction QtPrivate__AbstractComparatorFunction;
typedef struct QtPrivate__AbstractConverterFunction QtPrivate__AbstractConverterFunction;
typedef struct QtPrivate__AbstractDebugStreamFunction QtPrivate__AbstractDebugStreamFunction;
#endif

QtPrivate__AbstractDebugStreamFunction* QtPrivate__AbstractDebugStreamFunction_new();
void QtPrivate__AbstractDebugStreamFunction_Delete(QtPrivate__AbstractDebugStreamFunction* self);

QtPrivate__AbstractComparatorFunction* QtPrivate__AbstractComparatorFunction_new();
void QtPrivate__AbstractComparatorFunction_Delete(QtPrivate__AbstractComparatorFunction* self);

QtPrivate__AbstractConverterFunction* QtPrivate__AbstractConverterFunction_new();
void QtPrivate__AbstractConverterFunction_Delete(QtPrivate__AbstractConverterFunction* self);

QMetaType* QMetaType_new();
QMetaType* QMetaType_new2(const int typeVal);
bool QMetaType_UnregisterType(int typeVal);
int QMetaType_RegisterTypedef(const char* typeName, int aliasId);
int QMetaType_RegisterNormalizedTypedef(QByteArray* normalizedTypeName, int aliasId);
int QMetaType_Type(const char* typeName);
int QMetaType_TypeWithTypeName(QByteArray* typeName);
const char* QMetaType_TypeName(int typeVal);
int QMetaType_SizeOf(int typeVal);
int QMetaType_TypeFlags(int typeVal);
QMetaObject* QMetaType_MetaObjectForType(int typeVal);
bool QMetaType_IsRegistered(int typeVal);
void* QMetaType_Create(int typeVal);
void QMetaType_Destroy(int typeVal, void* data);
void* QMetaType_Construct(int typeVal, void* where, const void* copyVal);
void QMetaType_Destruct(int typeVal, void* where);
bool QMetaType_Save(QDataStream* stream, int typeVal, const void* data);
bool QMetaType_Load(QDataStream* stream, int typeVal, void* data);
bool QMetaType_IsValid(const QMetaType* self);
bool QMetaType_IsRegistered2(const QMetaType* self);
int QMetaType_Id(const QMetaType* self);
int QMetaType_SizeOf2(const QMetaType* self);
int QMetaType_Flags(const QMetaType* self);
QMetaObject* QMetaType_MetaObject(const QMetaType* self);
QByteArray* QMetaType_Name(const QMetaType* self);
void* QMetaType_Create2(const QMetaType* self);
void QMetaType_DestroyWithData(const QMetaType* self, void* data);
void* QMetaType_ConstructWithWhere(const QMetaType* self, void* where);
void QMetaType_DestructWithData(const QMetaType* self, void* data);
bool QMetaType_HasRegisteredComparators(int typeId);
bool QMetaType_HasRegisteredDebugStreamOperator(int typeId);
bool QMetaType_Convert(const void* from, int fromTypeId, void* to, int toTypeId);
bool QMetaType_Compare(const void* lhs, const void* rhs, int typeId, int* result);
bool QMetaType_Equals(const void* lhs, const void* rhs, int typeId, int* result);
bool QMetaType_DebugStream(QDebug* dbg, const void* rhs, int typeId);
bool QMetaType_HasRegisteredConverterFunction(int fromTypeId, int toTypeId);
void* QMetaType_Create22(int typeVal, const void* copyVal);
void* QMetaType_Create1(const QMetaType* self, const void* copyVal);
void* QMetaType_Construct2(const QMetaType* self, void* where, const void* copyVal);
void QMetaType_Delete(QMetaType* self);

QtMetaTypePrivate__VariantData* QtMetaTypePrivate__VariantData_new(const int metaTypeId_, const void* data_, const unsigned int flags_);
QtMetaTypePrivate__VariantData* QtMetaTypePrivate__VariantData_new2(QtMetaTypePrivate__VariantData* other);
void QtMetaTypePrivate__VariantData_Delete(QtMetaTypePrivate__VariantData* self);

void QtMetaTypePrivate__VectorBoolElements_Delete(QtMetaTypePrivate__VectorBoolElements* self);

QtMetaTypePrivate__QSequentialIterableImpl* QtMetaTypePrivate__QSequentialIterableImpl_new();
QtMetaTypePrivate__QSequentialIterableImpl* QtMetaTypePrivate__QSequentialIterableImpl_new2(QtMetaTypePrivate__QSequentialIterableImpl* param1);
int QtMetaTypePrivate__QSequentialIterableImpl_IteratorCapabilities(QtMetaTypePrivate__QSequentialIterableImpl* self);
unsigned int QtMetaTypePrivate__QSequentialIterableImpl_Revision(QtMetaTypePrivate__QSequentialIterableImpl* self);
unsigned int QtMetaTypePrivate__QSequentialIterableImpl_ContainerCapabilities(QtMetaTypePrivate__QSequentialIterableImpl* self);
void QtMetaTypePrivate__QSequentialIterableImpl_MoveToBegin(QtMetaTypePrivate__QSequentialIterableImpl* self);
void QtMetaTypePrivate__QSequentialIterableImpl_MoveToEnd(QtMetaTypePrivate__QSequentialIterableImpl* self);
bool QtMetaTypePrivate__QSequentialIterableImpl_Equal(const QtMetaTypePrivate__QSequentialIterableImpl* self, QtMetaTypePrivate__QSequentialIterableImpl* other);
QtMetaTypePrivate__QSequentialIterableImpl* QtMetaTypePrivate__QSequentialIterableImpl_Advance(QtMetaTypePrivate__QSequentialIterableImpl* self, int i);
void QtMetaTypePrivate__QSequentialIterableImpl_Append(QtMetaTypePrivate__QSequentialIterableImpl* self, const void* newElement);
QtMetaTypePrivate__VariantData* QtMetaTypePrivate__QSequentialIterableImpl_GetCurrent(const QtMetaTypePrivate__QSequentialIterableImpl* self);
QtMetaTypePrivate__VariantData* QtMetaTypePrivate__QSequentialIterableImpl_At(const QtMetaTypePrivate__QSequentialIterableImpl* self, int idx);
int QtMetaTypePrivate__QSequentialIterableImpl_Size(const QtMetaTypePrivate__QSequentialIterableImpl* self);
void QtMetaTypePrivate__QSequentialIterableImpl_DestroyIter(QtMetaTypePrivate__QSequentialIterableImpl* self);
void QtMetaTypePrivate__QSequentialIterableImpl_Copy(QtMetaTypePrivate__QSequentialIterableImpl* self, QtMetaTypePrivate__QSequentialIterableImpl* other);
void QtMetaTypePrivate__QSequentialIterableImpl_OperatorAssign(QtMetaTypePrivate__QSequentialIterableImpl* self, QtMetaTypePrivate__QSequentialIterableImpl* param1);
void QtMetaTypePrivate__QSequentialIterableImpl_Delete(QtMetaTypePrivate__QSequentialIterableImpl* self);

QtMetaTypePrivate__QAssociativeIterableImpl* QtMetaTypePrivate__QAssociativeIterableImpl_new();
QtMetaTypePrivate__QAssociativeIterableImpl* QtMetaTypePrivate__QAssociativeIterableImpl_new2(QtMetaTypePrivate__QAssociativeIterableImpl* param1);
void QtMetaTypePrivate__QAssociativeIterableImpl_Begin(QtMetaTypePrivate__QAssociativeIterableImpl* self);
void QtMetaTypePrivate__QAssociativeIterableImpl_End(QtMetaTypePrivate__QAssociativeIterableImpl* self);
bool QtMetaTypePrivate__QAssociativeIterableImpl_Equal(const QtMetaTypePrivate__QAssociativeIterableImpl* self, QtMetaTypePrivate__QAssociativeIterableImpl* other);
QtMetaTypePrivate__QAssociativeIterableImpl* QtMetaTypePrivate__QAssociativeIterableImpl_Advance(QtMetaTypePrivate__QAssociativeIterableImpl* self, int i);
void QtMetaTypePrivate__QAssociativeIterableImpl_DestroyIter(QtMetaTypePrivate__QAssociativeIterableImpl* self);
QtMetaTypePrivate__VariantData* QtMetaTypePrivate__QAssociativeIterableImpl_GetCurrentKey(const QtMetaTypePrivate__QAssociativeIterableImpl* self);
QtMetaTypePrivate__VariantData* QtMetaTypePrivate__QAssociativeIterableImpl_GetCurrentValue(const QtMetaTypePrivate__QAssociativeIterableImpl* self);
void QtMetaTypePrivate__QAssociativeIterableImpl_Find(QtMetaTypePrivate__QAssociativeIterableImpl* self, QtMetaTypePrivate__VariantData* key);
int QtMetaTypePrivate__QAssociativeIterableImpl_Size(const QtMetaTypePrivate__QAssociativeIterableImpl* self);
void QtMetaTypePrivate__QAssociativeIterableImpl_Copy(QtMetaTypePrivate__QAssociativeIterableImpl* self, QtMetaTypePrivate__QAssociativeIterableImpl* other);
void QtMetaTypePrivate__QAssociativeIterableImpl_OperatorAssign(QtMetaTypePrivate__QAssociativeIterableImpl* self, QtMetaTypePrivate__QAssociativeIterableImpl* param1);
void QtMetaTypePrivate__QAssociativeIterableImpl_Delete(QtMetaTypePrivate__QAssociativeIterableImpl* self);

QtMetaTypePrivate__QPairVariantInterfaceImpl* QtMetaTypePrivate__QPairVariantInterfaceImpl_new();
QtMetaTypePrivate__QPairVariantInterfaceImpl* QtMetaTypePrivate__QPairVariantInterfaceImpl_new2(QtMetaTypePrivate__QPairVariantInterfaceImpl* param1);
QtMetaTypePrivate__VariantData* QtMetaTypePrivate__QPairVariantInterfaceImpl_First(const QtMetaTypePrivate__QPairVariantInterfaceImpl* self);
QtMetaTypePrivate__VariantData* QtMetaTypePrivate__QPairVariantInterfaceImpl_Second(const QtMetaTypePrivate__QPairVariantInterfaceImpl* self);
void QtMetaTypePrivate__QPairVariantInterfaceImpl_Delete(QtMetaTypePrivate__QPairVariantInterfaceImpl* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
