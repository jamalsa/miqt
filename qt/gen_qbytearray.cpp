#include <QByteArray>
#define WORKAROUND_INNER_CLASS_DEFINITION_QByteArray__FromBase64Result
#include <QByteArrayDataPtr>
#include <QByteRef>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include "qbytearray.h"
#include "gen_qbytearray.h"
#include "_cgo_export.h"

QByteArrayDataPtr* QByteArrayDataPtr_new() {
	return new QByteArrayDataPtr();
}

QByteArrayDataPtr* QByteArrayDataPtr_new2(QByteArrayDataPtr* param1) {
	return new QByteArrayDataPtr(*param1);
}

void QByteArrayDataPtr_Delete(QByteArrayDataPtr* self) {
	delete self;
}

QByteArray* QByteArray_new() {
	return new QByteArray();
}

QByteArray* QByteArray_new2(const char* param1) {
	return new QByteArray(param1);
}

QByteArray* QByteArray_new3(int size, char c) {
	return new QByteArray(static_cast<int>(size), static_cast<char>(c));
}

QByteArray* QByteArray_new4(int size, uintptr_t param2) {
	return new QByteArray(static_cast<int>(size), static_cast<Qt::Initialization>(param2));
}

QByteArray* QByteArray_new5(QByteArray* param1) {
	return new QByteArray(*param1);
}

QByteArray* QByteArray_new6(QByteArrayDataPtr* dd) {
	return new QByteArray(*dd);
}

QByteArray* QByteArray_new7(const char* param1, int size) {
	return new QByteArray(param1, static_cast<int>(size));
}

void QByteArray_OperatorAssign(QByteArray* self, QByteArray* param1) {
	self->operator=(*param1);
}

void QByteArray_OperatorAssignWithStr(QByteArray* self, const char* str) {
	self->operator=(str);
}

void QByteArray_Swap(QByteArray* self, QByteArray* other) {
	self->swap(*other);
}

int QByteArray_Size(const QByteArray* self) {
	return self->size();
}

bool QByteArray_IsEmpty(const QByteArray* self) {
	return self->isEmpty();
}

void QByteArray_Resize(QByteArray* self, int size) {
	self->resize(static_cast<int>(size));
}

QByteArray* QByteArray_Fill(QByteArray* self, char c) {
	QByteArray& _ret = self->fill(static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

int QByteArray_Capacity(const QByteArray* self) {
	return self->capacity();
}

void QByteArray_Reserve(QByteArray* self, int size) {
	self->reserve(static_cast<int>(size));
}

void QByteArray_Squeeze(QByteArray* self) {
	self->squeeze();
}

char* QByteArray_Data(QByteArray* self) {
	return self->data();
}

const char* QByteArray_Data2(const QByteArray* self) {
	return (const char*) self->data();
}

const char* QByteArray_ConstData(const QByteArray* self) {
	return (const char*) self->constData();
}

void QByteArray_Detach(QByteArray* self) {
	self->detach();
}

bool QByteArray_IsDetached(const QByteArray* self) {
	return self->isDetached();
}

bool QByteArray_IsSharedWith(const QByteArray* self, QByteArray* other) {
	return self->isSharedWith(*other);
}

void QByteArray_Clear(QByteArray* self) {
	self->clear();
}

char QByteArray_At(const QByteArray* self, int i) {
	return self->at(static_cast<int>(i));
}

char QByteArray_OperatorSubscript(const QByteArray* self, int i) {
	return self->operator[](static_cast<int>(i));
}

char QByteArray_OperatorSubscriptWithUint(const QByteArray* self, unsigned int i) {
	return self->operator[](static_cast<uint>(i));
}

QByteRef* QByteArray_OperatorSubscriptWithInt(QByteArray* self, int i) {
	QByteRef _ret = self->operator[](static_cast<int>(i));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteRef*>(new QByteRef(_ret));
}

QByteRef* QByteArray_OperatorSubscript2(QByteArray* self, unsigned int i) {
	QByteRef _ret = self->operator[](static_cast<uint>(i));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteRef*>(new QByteRef(_ret));
}

char QByteArray_Front(const QByteArray* self) {
	return self->front();
}

QByteRef* QByteArray_Front2(QByteArray* self) {
	QByteRef _ret = self->front();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteRef*>(new QByteRef(_ret));
}

char QByteArray_Back(const QByteArray* self) {
	return self->back();
}

QByteRef* QByteArray_Back2(QByteArray* self) {
	QByteRef _ret = self->back();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteRef*>(new QByteRef(_ret));
}

int QByteArray_IndexOf(const QByteArray* self, char c) {
	return self->indexOf(static_cast<char>(c));
}

int QByteArray_IndexOfWithChar(const QByteArray* self, const char* c) {
	return self->indexOf(c);
}

int QByteArray_IndexOfWithQByteArray(const QByteArray* self, QByteArray* a) {
	return self->indexOf(*a);
}

int QByteArray_LastIndexOf(const QByteArray* self, char c) {
	return self->lastIndexOf(static_cast<char>(c));
}

int QByteArray_LastIndexOfWithChar(const QByteArray* self, const char* c) {
	return self->lastIndexOf(c);
}

int QByteArray_LastIndexOfWithQByteArray(const QByteArray* self, QByteArray* a) {
	return self->lastIndexOf(*a);
}

bool QByteArray_Contains(const QByteArray* self, char c) {
	return self->contains(static_cast<char>(c));
}

bool QByteArray_ContainsWithChar(const QByteArray* self, const char* a) {
	return self->contains(a);
}

bool QByteArray_ContainsWithQByteArray(const QByteArray* self, QByteArray* a) {
	return self->contains(*a);
}

int QByteArray_Count(const QByteArray* self, char c) {
	return self->count(static_cast<char>(c));
}

int QByteArray_CountWithChar(const QByteArray* self, const char* a) {
	return self->count(a);
}

int QByteArray_CountWithQByteArray(const QByteArray* self, QByteArray* a) {
	return self->count(*a);
}

int QByteArray_Compare(const QByteArray* self, const char* c) {
	return self->compare(c);
}

int QByteArray_CompareWithQByteArray(const QByteArray* self, QByteArray* a) {
	return self->compare(*a);
}

QByteArray* QByteArray_Left(const QByteArray* self, int lenVal) {
	QByteArray _ret = self->left(static_cast<int>(lenVal));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Right(const QByteArray* self, int lenVal) {
	QByteArray _ret = self->right(static_cast<int>(lenVal));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Mid(const QByteArray* self, int index) {
	QByteArray _ret = self->mid(static_cast<int>(index));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Chopped(const QByteArray* self, int lenVal) {
	QByteArray _ret = self->chopped(static_cast<int>(lenVal));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

bool QByteArray_StartsWith(const QByteArray* self, QByteArray* a) {
	return self->startsWith(*a);
}

bool QByteArray_StartsWithWithChar(const QByteArray* self, char c) {
	return self->startsWith(static_cast<char>(c));
}

bool QByteArray_StartsWith2(const QByteArray* self, const char* c) {
	return self->startsWith(c);
}

bool QByteArray_EndsWith(const QByteArray* self, QByteArray* a) {
	return self->endsWith(*a);
}

bool QByteArray_EndsWithWithChar(const QByteArray* self, char c) {
	return self->endsWith(static_cast<char>(c));
}

bool QByteArray_EndsWith2(const QByteArray* self, const char* c) {
	return self->endsWith(c);
}

bool QByteArray_IsUpper(const QByteArray* self) {
	return self->isUpper();
}

bool QByteArray_IsLower(const QByteArray* self) {
	return self->isLower();
}

void QByteArray_Truncate(QByteArray* self, int pos) {
	self->truncate(static_cast<int>(pos));
}

void QByteArray_Chop(QByteArray* self, int n) {
	self->chop(static_cast<int>(n));
}

QByteArray* QByteArray_ToLower(const QByteArray* self) {
	QByteArray _ret = self->toLower();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_ToUpper(const QByteArray* self) {
	QByteArray _ret = self->toUpper();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Trimmed(const QByteArray* self) {
	QByteArray _ret = self->trimmed();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Simplified(const QByteArray* self) {
	QByteArray _ret = self->simplified();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_LeftJustified(const QByteArray* self, int width) {
	QByteArray _ret = self->leftJustified(static_cast<int>(width));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_RightJustified(const QByteArray* self, int width) {
	QByteArray _ret = self->rightJustified(static_cast<int>(width));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Prepend(QByteArray* self, char c) {
	QByteArray& _ret = self->prepend(static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Prepend2(QByteArray* self, int count, char c) {
	QByteArray& _ret = self->prepend(static_cast<int>(count), static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_PrependWithChar(QByteArray* self, const char* s) {
	QByteArray& _ret = self->prepend(s);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Prepend3(QByteArray* self, const char* s, int lenVal) {
	QByteArray& _ret = self->prepend(s, static_cast<int>(lenVal));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_PrependWithQByteArray(QByteArray* self, QByteArray* a) {
	QByteArray& _ret = self->prepend(*a);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Append(QByteArray* self, char c) {
	QByteArray& _ret = self->append(static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Append2(QByteArray* self, int count, char c) {
	QByteArray& _ret = self->append(static_cast<int>(count), static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_AppendWithChar(QByteArray* self, const char* s) {
	QByteArray& _ret = self->append(s);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Append3(QByteArray* self, const char* s, int lenVal) {
	QByteArray& _ret = self->append(s, static_cast<int>(lenVal));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_AppendWithQByteArray(QByteArray* self, QByteArray* a) {
	QByteArray& _ret = self->append(*a);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Insert(QByteArray* self, int i, char c) {
	QByteArray& _ret = self->insert(static_cast<int>(i), static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Insert2(QByteArray* self, int i, int count, char c) {
	QByteArray& _ret = self->insert(static_cast<int>(i), static_cast<int>(count), static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Insert3(QByteArray* self, int i, const char* s) {
	QByteArray& _ret = self->insert(static_cast<int>(i), s);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Insert4(QByteArray* self, int i, const char* s, int lenVal) {
	QByteArray& _ret = self->insert(static_cast<int>(i), s, static_cast<int>(lenVal));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Insert5(QByteArray* self, int i, QByteArray* a) {
	QByteArray& _ret = self->insert(static_cast<int>(i), *a);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Remove(QByteArray* self, int index, int lenVal) {
	QByteArray& _ret = self->remove(static_cast<int>(index), static_cast<int>(lenVal));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace(QByteArray* self, int index, int lenVal, const char* s) {
	QByteArray& _ret = self->replace(static_cast<int>(index), static_cast<int>(lenVal), s);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace2(QByteArray* self, int index, int lenVal, const char* s, int alen) {
	QByteArray& _ret = self->replace(static_cast<int>(index), static_cast<int>(lenVal), s, static_cast<int>(alen));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace3(QByteArray* self, int index, int lenVal, QByteArray* s) {
	QByteArray& _ret = self->replace(static_cast<int>(index), static_cast<int>(lenVal), *s);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace4(QByteArray* self, char before, const char* after) {
	QByteArray& _ret = self->replace(static_cast<char>(before), after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace5(QByteArray* self, char before, QByteArray* after) {
	QByteArray& _ret = self->replace(static_cast<char>(before), *after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace6(QByteArray* self, const char* before, const char* after) {
	QByteArray& _ret = self->replace(before, after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace7(QByteArray* self, const char* before, int bsize, const char* after, int asize) {
	QByteArray& _ret = self->replace(before, static_cast<int>(bsize), after, static_cast<int>(asize));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace8(QByteArray* self, QByteArray* before, QByteArray* after) {
	QByteArray& _ret = self->replace(*before, *after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace9(QByteArray* self, QByteArray* before, const char* after) {
	QByteArray& _ret = self->replace(*before, after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace10(QByteArray* self, const char* before, QByteArray* after) {
	QByteArray& _ret = self->replace(before, *after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace11(QByteArray* self, char before, char after) {
	QByteArray& _ret = self->replace(static_cast<char>(before), static_cast<char>(after));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_OperatorPlusAssign(QByteArray* self, char c) {
	QByteArray& _ret = self->operator+=(static_cast<char>(c));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_OperatorPlusAssignWithChar(QByteArray* self, const char* s) {
	QByteArray& _ret = self->operator+=(s);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_OperatorPlusAssignWithQByteArray(QByteArray* self, QByteArray* a) {
	QByteArray& _ret = self->operator+=(*a);
	// Cast returned reference into pointer
	return &_ret;
}

struct miqt_array* QByteArray_Split(const QByteArray* self, char sep) {
	QList<QByteArray> _ret = self->split(static_cast<char>(sep));
	// Convert QList<> from C++ memory to manually-managed C memory of copy-constructed pointers
	QByteArray** _arr = static_cast<QByteArray**>(malloc(sizeof(QByteArray**) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QByteArray(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

QByteArray* QByteArray_Repeated(const QByteArray* self, int times) {
	QByteArray _ret = self->repeated(static_cast<int>(times));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_AppendWithQString(QByteArray* self, struct miqt_string* s) {
	QString s_QString = QString::fromUtf8(&s->data, s->len);
	QByteArray& _ret = self->append(s_QString);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Insert6(QByteArray* self, int i, struct miqt_string* s) {
	QString s_QString = QString::fromUtf8(&s->data, s->len);
	QByteArray& _ret = self->insert(static_cast<int>(i), s_QString);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace12(QByteArray* self, struct miqt_string* before, const char* after) {
	QString before_QString = QString::fromUtf8(&before->data, before->len);
	QByteArray& _ret = self->replace(before_QString, after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace13(QByteArray* self, char c, struct miqt_string* after) {
	QString after_QString = QString::fromUtf8(&after->data, after->len);
	QByteArray& _ret = self->replace(static_cast<char>(c), after_QString);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Replace14(QByteArray* self, struct miqt_string* before, QByteArray* after) {
	QString before_QString = QString::fromUtf8(&before->data, before->len);
	QByteArray& _ret = self->replace(before_QString, *after);
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_OperatorPlusAssignWithQString(QByteArray* self, struct miqt_string* s) {
	QString s_QString = QString::fromUtf8(&s->data, s->len);
	QByteArray& _ret = self->operator+=(s_QString);
	// Cast returned reference into pointer
	return &_ret;
}

int QByteArray_IndexOfWithQString(const QByteArray* self, struct miqt_string* s) {
	QString s_QString = QString::fromUtf8(&s->data, s->len);
	return self->indexOf(s_QString);
}

int QByteArray_LastIndexOfWithQString(const QByteArray* self, struct miqt_string* s) {
	QString s_QString = QString::fromUtf8(&s->data, s->len);
	return self->lastIndexOf(s_QString);
}

bool QByteArray_OperatorEqual(const QByteArray* self, struct miqt_string* s2) {
	QString s2_QString = QString::fromUtf8(&s2->data, s2->len);
	return self->operator==(s2_QString);
}

bool QByteArray_OperatorNotEqual(const QByteArray* self, struct miqt_string* s2) {
	QString s2_QString = QString::fromUtf8(&s2->data, s2->len);
	return self->operator!=(s2_QString);
}

bool QByteArray_OperatorLesser(const QByteArray* self, struct miqt_string* s2) {
	QString s2_QString = QString::fromUtf8(&s2->data, s2->len);
	return self->operator<(s2_QString);
}

bool QByteArray_OperatorGreater(const QByteArray* self, struct miqt_string* s2) {
	QString s2_QString = QString::fromUtf8(&s2->data, s2->len);
	return self->operator>(s2_QString);
}

bool QByteArray_OperatorLesserOrEqual(const QByteArray* self, struct miqt_string* s2) {
	QString s2_QString = QString::fromUtf8(&s2->data, s2->len);
	return self->operator<=(s2_QString);
}

bool QByteArray_OperatorGreaterOrEqual(const QByteArray* self, struct miqt_string* s2) {
	QString s2_QString = QString::fromUtf8(&s2->data, s2->len);
	return self->operator>=(s2_QString);
}

int16_t QByteArray_ToShort(const QByteArray* self) {
	return self->toShort();
}

uint16_t QByteArray_ToUShort(const QByteArray* self) {
	return self->toUShort();
}

int QByteArray_ToInt(const QByteArray* self) {
	return self->toInt();
}

unsigned int QByteArray_ToUInt(const QByteArray* self) {
	return self->toUInt();
}

long QByteArray_ToLong(const QByteArray* self) {
	return self->toLong();
}

unsigned long QByteArray_ToULong(const QByteArray* self) {
	return self->toULong();
}

int64_t QByteArray_ToLongLong(const QByteArray* self) {
	return self->toLongLong();
}

uint64_t QByteArray_ToULongLong(const QByteArray* self) {
	return self->toULongLong();
}

float QByteArray_ToFloat(const QByteArray* self) {
	return self->toFloat();
}

double QByteArray_ToDouble(const QByteArray* self) {
	return self->toDouble();
}

QByteArray* QByteArray_ToBase64(const QByteArray* self, int options) {
	QByteArray _ret = self->toBase64(static_cast<QByteArray::Base64Options>(options));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_ToBase642(const QByteArray* self) {
	QByteArray _ret = self->toBase64();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_ToHex(const QByteArray* self) {
	QByteArray _ret = self->toHex();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_ToHexWithSeparator(const QByteArray* self, char separator) {
	QByteArray _ret = self->toHex(static_cast<char>(separator));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_ToPercentEncoding(const QByteArray* self) {
	QByteArray _ret = self->toPercentEncoding();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_SetNum(QByteArray* self, int16_t param1) {
	QByteArray& _ret = self->setNum(static_cast<short>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNumWithUshort(QByteArray* self, uint16_t param1) {
	QByteArray& _ret = self->setNum(static_cast<ushort>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNumWithInt(QByteArray* self, int param1) {
	QByteArray& _ret = self->setNum(static_cast<int>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNumWithUint(QByteArray* self, unsigned int param1) {
	QByteArray& _ret = self->setNum(static_cast<uint>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNumWithQlonglong(QByteArray* self, long long param1) {
	QByteArray& _ret = self->setNum(static_cast<qint64>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNumWithQulonglong(QByteArray* self, unsigned long long param1) {
	QByteArray& _ret = self->setNum(static_cast<quint64>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNumWithFloat(QByteArray* self, float param1) {
	QByteArray& _ret = self->setNum(static_cast<float>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNumWithDouble(QByteArray* self, double param1) {
	QByteArray& _ret = self->setNum(static_cast<double>(param1));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetRawData(QByteArray* self, const char* a, unsigned int n) {
	QByteArray& _ret = self->setRawData(a, static_cast<uint>(n));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Number(int param1) {
	QByteArray _ret = QByteArray::number(static_cast<int>(param1));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_NumberWithUint(unsigned int param1) {
	QByteArray _ret = QByteArray::number(static_cast<uint>(param1));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_NumberWithQlonglong(long long param1) {
	QByteArray _ret = QByteArray::number(static_cast<qint64>(param1));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_NumberWithQulonglong(unsigned long long param1) {
	QByteArray _ret = QByteArray::number(static_cast<quint64>(param1));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_NumberWithDouble(double param1) {
	QByteArray _ret = QByteArray::number(static_cast<double>(param1));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_FromRawData(const char* param1, int size) {
	QByteArray _ret = QByteArray::fromRawData(param1, static_cast<int>(size));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray__FromBase64Result* QByteArray_FromBase64Encoding(QByteArray* base64) {
	QByteArray::FromBase64Result _ret = QByteArray::fromBase64Encoding(*base64);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray::FromBase64Result*>(new QByteArray::FromBase64Result(_ret));
}

QByteArray* QByteArray_FromBase64(QByteArray* base64, int options) {
	QByteArray _ret = QByteArray::fromBase64(*base64, static_cast<QByteArray::Base64Options>(options));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_FromBase64WithBase64(QByteArray* base64) {
	QByteArray _ret = QByteArray::fromBase64(*base64);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_FromHex(QByteArray* hexEncoded) {
	QByteArray _ret = QByteArray::fromHex(*hexEncoded);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_FromPercentEncoding(QByteArray* pctEncoded) {
	QByteArray _ret = QByteArray::fromPercentEncoding(*pctEncoded);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

char* QByteArray_Begin(QByteArray* self) {
	return self->begin();
}

const char* QByteArray_Begin2(const QByteArray* self) {
	return (const char*) self->begin();
}

const char* QByteArray_Cbegin(const QByteArray* self) {
	return (const char*) self->cbegin();
}

const char* QByteArray_ConstBegin(const QByteArray* self) {
	return (const char*) self->constBegin();
}

char* QByteArray_End(QByteArray* self) {
	return self->end();
}

const char* QByteArray_End2(const QByteArray* self) {
	return (const char*) self->end();
}

const char* QByteArray_Cend(const QByteArray* self) {
	return (const char*) self->cend();
}

const char* QByteArray_ConstEnd(const QByteArray* self) {
	return (const char*) self->constEnd();
}

void QByteArray_PushBack(QByteArray* self, char c) {
	self->push_back(static_cast<char>(c));
}

void QByteArray_PushBackWithChar(QByteArray* self, const char* c) {
	self->push_back(c);
}

void QByteArray_PushBackWithQByteArray(QByteArray* self, QByteArray* a) {
	self->push_back(*a);
}

void QByteArray_PushFront(QByteArray* self, char c) {
	self->push_front(static_cast<char>(c));
}

void QByteArray_PushFrontWithChar(QByteArray* self, const char* c) {
	self->push_front(c);
}

void QByteArray_PushFrontWithQByteArray(QByteArray* self, QByteArray* a) {
	self->push_front(*a);
}

void QByteArray_ShrinkToFit(QByteArray* self) {
	self->shrink_to_fit();
}

int QByteArray_Count2(const QByteArray* self) {
	return self->count();
}

int QByteArray_Length(const QByteArray* self) {
	return self->length();
}

bool QByteArray_IsNull(const QByteArray* self) {
	return self->isNull();
}

QByteArray* QByteArray_Fill2(QByteArray* self, char c, int size) {
	QByteArray& _ret = self->fill(static_cast<char>(c), static_cast<int>(size));
	// Cast returned reference into pointer
	return &_ret;
}

int QByteArray_IndexOf2(const QByteArray* self, char c, int from) {
	return self->indexOf(static_cast<char>(c), static_cast<int>(from));
}

int QByteArray_IndexOf22(const QByteArray* self, const char* c, int from) {
	return self->indexOf(c, static_cast<int>(from));
}

int QByteArray_IndexOf23(const QByteArray* self, QByteArray* a, int from) {
	return self->indexOf(*a, static_cast<int>(from));
}

int QByteArray_LastIndexOf2(const QByteArray* self, char c, int from) {
	return self->lastIndexOf(static_cast<char>(c), static_cast<int>(from));
}

int QByteArray_LastIndexOf22(const QByteArray* self, const char* c, int from) {
	return self->lastIndexOf(c, static_cast<int>(from));
}

int QByteArray_LastIndexOf23(const QByteArray* self, QByteArray* a, int from) {
	return self->lastIndexOf(*a, static_cast<int>(from));
}

int QByteArray_Compare2(const QByteArray* self, const char* c, uintptr_t cs) {
	return self->compare(c, static_cast<Qt::CaseSensitivity>(cs));
}

int QByteArray_Compare22(const QByteArray* self, QByteArray* a, uintptr_t cs) {
	return self->compare(*a, static_cast<Qt::CaseSensitivity>(cs));
}

QByteArray* QByteArray_Mid2(const QByteArray* self, int index, int lenVal) {
	QByteArray _ret = self->mid(static_cast<int>(index), static_cast<int>(lenVal));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_LeftJustified2(const QByteArray* self, int width, char fill) {
	QByteArray _ret = self->leftJustified(static_cast<int>(width), static_cast<char>(fill));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_LeftJustified3(const QByteArray* self, int width, char fill, bool truncate) {
	QByteArray _ret = self->leftJustified(static_cast<int>(width), static_cast<char>(fill), truncate);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_RightJustified2(const QByteArray* self, int width, char fill) {
	QByteArray _ret = self->rightJustified(static_cast<int>(width), static_cast<char>(fill));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_RightJustified3(const QByteArray* self, int width, char fill, bool truncate) {
	QByteArray _ret = self->rightJustified(static_cast<int>(width), static_cast<char>(fill), truncate);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

int QByteArray_IndexOf24(const QByteArray* self, struct miqt_string* s, int from) {
	QString s_QString = QString::fromUtf8(&s->data, s->len);
	return self->indexOf(s_QString, static_cast<int>(from));
}

int QByteArray_LastIndexOf24(const QByteArray* self, struct miqt_string* s, int from) {
	QString s_QString = QString::fromUtf8(&s->data, s->len);
	return self->lastIndexOf(s_QString, static_cast<int>(from));
}

int16_t QByteArray_ToShort1(const QByteArray* self, bool* ok) {
	return self->toShort(ok);
}

int16_t QByteArray_ToShort2(const QByteArray* self, bool* ok, int base) {
	return self->toShort(ok, static_cast<int>(base));
}

uint16_t QByteArray_ToUShort1(const QByteArray* self, bool* ok) {
	return self->toUShort(ok);
}

uint16_t QByteArray_ToUShort2(const QByteArray* self, bool* ok, int base) {
	return self->toUShort(ok, static_cast<int>(base));
}

int QByteArray_ToInt1(const QByteArray* self, bool* ok) {
	return self->toInt(ok);
}

int QByteArray_ToInt2(const QByteArray* self, bool* ok, int base) {
	return self->toInt(ok, static_cast<int>(base));
}

unsigned int QByteArray_ToUInt1(const QByteArray* self, bool* ok) {
	return self->toUInt(ok);
}

unsigned int QByteArray_ToUInt2(const QByteArray* self, bool* ok, int base) {
	return self->toUInt(ok, static_cast<int>(base));
}

long QByteArray_ToLong1(const QByteArray* self, bool* ok) {
	return self->toLong(ok);
}

long QByteArray_ToLong2(const QByteArray* self, bool* ok, int base) {
	return self->toLong(ok, static_cast<int>(base));
}

unsigned long QByteArray_ToULong1(const QByteArray* self, bool* ok) {
	return self->toULong(ok);
}

unsigned long QByteArray_ToULong2(const QByteArray* self, bool* ok, int base) {
	return self->toULong(ok, static_cast<int>(base));
}

int64_t QByteArray_ToLongLong1(const QByteArray* self, bool* ok) {
	return self->toLongLong(ok);
}

int64_t QByteArray_ToLongLong2(const QByteArray* self, bool* ok, int base) {
	return self->toLongLong(ok, static_cast<int>(base));
}

uint64_t QByteArray_ToULongLong1(const QByteArray* self, bool* ok) {
	return self->toULongLong(ok);
}

uint64_t QByteArray_ToULongLong2(const QByteArray* self, bool* ok, int base) {
	return self->toULongLong(ok, static_cast<int>(base));
}

float QByteArray_ToFloat1(const QByteArray* self, bool* ok) {
	return self->toFloat(ok);
}

double QByteArray_ToDouble1(const QByteArray* self, bool* ok) {
	return self->toDouble(ok);
}

QByteArray* QByteArray_ToPercentEncoding1(const QByteArray* self, QByteArray* exclude) {
	QByteArray _ret = self->toPercentEncoding(*exclude);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_ToPercentEncoding2(const QByteArray* self, QByteArray* exclude, QByteArray* include) {
	QByteArray _ret = self->toPercentEncoding(*exclude, *include);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_ToPercentEncoding3(const QByteArray* self, QByteArray* exclude, QByteArray* include, char percent) {
	QByteArray _ret = self->toPercentEncoding(*exclude, *include, static_cast<char>(percent));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_SetNum2(QByteArray* self, int16_t param1, int base) {
	QByteArray& _ret = self->setNum(static_cast<short>(param1), static_cast<int>(base));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum22(QByteArray* self, uint16_t param1, int base) {
	QByteArray& _ret = self->setNum(static_cast<ushort>(param1), static_cast<int>(base));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum23(QByteArray* self, int param1, int base) {
	QByteArray& _ret = self->setNum(static_cast<int>(param1), static_cast<int>(base));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum24(QByteArray* self, unsigned int param1, int base) {
	QByteArray& _ret = self->setNum(static_cast<uint>(param1), static_cast<int>(base));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum25(QByteArray* self, long long param1, int base) {
	QByteArray& _ret = self->setNum(static_cast<qint64>(param1), static_cast<int>(base));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum26(QByteArray* self, unsigned long long param1, int base) {
	QByteArray& _ret = self->setNum(static_cast<quint64>(param1), static_cast<int>(base));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum27(QByteArray* self, float param1, char f) {
	QByteArray& _ret = self->setNum(static_cast<float>(param1), static_cast<char>(f));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum3(QByteArray* self, float param1, char f, int prec) {
	QByteArray& _ret = self->setNum(static_cast<float>(param1), static_cast<char>(f), static_cast<int>(prec));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum28(QByteArray* self, double param1, char f) {
	QByteArray& _ret = self->setNum(static_cast<double>(param1), static_cast<char>(f));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_SetNum32(QByteArray* self, double param1, char f, int prec) {
	QByteArray& _ret = self->setNum(static_cast<double>(param1), static_cast<char>(f), static_cast<int>(prec));
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray_Number2(int param1, int base) {
	QByteArray _ret = QByteArray::number(static_cast<int>(param1), static_cast<int>(base));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Number22(unsigned int param1, int base) {
	QByteArray _ret = QByteArray::number(static_cast<uint>(param1), static_cast<int>(base));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Number23(long long param1, int base) {
	QByteArray _ret = QByteArray::number(static_cast<qint64>(param1), static_cast<int>(base));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Number24(unsigned long long param1, int base) {
	QByteArray _ret = QByteArray::number(static_cast<quint64>(param1), static_cast<int>(base));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Number25(double param1, char f) {
	QByteArray _ret = QByteArray::number(static_cast<double>(param1), static_cast<char>(f));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray* QByteArray_Number3(double param1, char f, int prec) {
	QByteArray _ret = QByteArray::number(static_cast<double>(param1), static_cast<char>(f), static_cast<int>(prec));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

QByteArray__FromBase64Result* QByteArray_FromBase64Encoding2(QByteArray* base64, int options) {
	QByteArray::FromBase64Result _ret = QByteArray::fromBase64Encoding(*base64, static_cast<QByteArray::Base64Options>(options));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray::FromBase64Result*>(new QByteArray::FromBase64Result(_ret));
}

QByteArray* QByteArray_FromPercentEncoding2(QByteArray* pctEncoded, char percent) {
	QByteArray _ret = QByteArray::fromPercentEncoding(*pctEncoded, static_cast<char>(percent));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

void QByteArray_Delete(QByteArray* self) {
	delete self;
}

QByteRef* QByteRef_new(QByteRef* param1) {
	return new QByteRef(*param1);
}

void QByteRef_OperatorAssign(QByteRef* self, char c) {
	self->operator=(static_cast<char>(c));
}

void QByteRef_OperatorAssignWithQByteRef(QByteRef* self, QByteRef* c) {
	self->operator=(*c);
}

bool QByteRef_OperatorEqual(const QByteRef* self, char c) {
	return self->operator==(static_cast<char>(c));
}

bool QByteRef_OperatorNotEqual(const QByteRef* self, char c) {
	return self->operator!=(static_cast<char>(c));
}

bool QByteRef_OperatorGreater(const QByteRef* self, char c) {
	return self->operator>(static_cast<char>(c));
}

bool QByteRef_OperatorGreaterOrEqual(const QByteRef* self, char c) {
	return self->operator>=(static_cast<char>(c));
}

bool QByteRef_OperatorLesser(const QByteRef* self, char c) {
	return self->operator<(static_cast<char>(c));
}

bool QByteRef_OperatorLesserOrEqual(const QByteRef* self, char c) {
	return self->operator<=(static_cast<char>(c));
}

void QByteRef_Delete(QByteRef* self) {
	delete self;
}

QByteArray__FromBase64Result* QByteArray__FromBase64Result_new(QByteArray__FromBase64Result* param1) {
	return new QByteArray::FromBase64Result(*param1);
}

void QByteArray__FromBase64Result_Swap(QByteArray__FromBase64Result* self, QByteArray__FromBase64Result* other) {
	self->swap(*other);
}

QByteArray* QByteArray__FromBase64Result_OperatorMultiply(QByteArray__FromBase64Result* self) {
	QByteArray& _ret = self->operator*();
	// Cast returned reference into pointer
	return &_ret;
}

QByteArray* QByteArray__FromBase64Result_OperatorMultiply2(const QByteArray__FromBase64Result* self) {
	const QByteArray& _ret = self->operator*();
	// Cast returned reference into pointer
	return const_cast<QByteArray*>(&_ret);
}

void QByteArray__FromBase64Result_Delete(QByteArray__FromBase64Result* self) {
	delete self;
}

