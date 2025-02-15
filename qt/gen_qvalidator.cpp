#include <QDoubleValidator>
#include <QIntValidator>
#include <QLocale>
#include <QMetaObject>
#include <QObject>
#include <QRegExp>
#include <QRegExpValidator>
#include <QRegularExpression>
#include <QRegularExpressionValidator>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QValidator>
#include <qvalidator.h>
#include "gen_qvalidator.h"
#include "_cgo_export.h"

QMetaObject* QValidator_MetaObject(const QValidator* self) {
	return (QMetaObject*) self->metaObject();
}

void* QValidator_Metacast(QValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QValidator_Tr(const char* s) {
	QString _ret = QValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QValidator_TrUtf8(const char* s) {
	QString _ret = QValidator::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QValidator_SetLocale(QValidator* self, QLocale* locale) {
	self->setLocale(*locale);
}

QLocale* QValidator_Locale(const QValidator* self) {
	return new QLocale(self->locale());
}

int QValidator_Validate(const QValidator* self, struct miqt_string param1, int* param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QValidator::State _ret = self->validate(param1_QString, static_cast<int&>(*param2));
	return static_cast<int>(_ret);
}

void QValidator_Fixup(const QValidator* self, struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	self->fixup(param1_QString);
}

void QValidator_Changed(QValidator* self) {
	self->changed();
}

void QValidator_connect_Changed(QValidator* self, intptr_t slot) {
	QValidator::connect(self, static_cast<void (QValidator::*)()>(&QValidator::changed), self, [=]() {
		miqt_exec_callback_QValidator_Changed(slot);
	});
}

struct miqt_string QValidator_Tr2(const char* s, const char* c) {
	QString _ret = QValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QValidator_TrUtf82(const char* s, const char* c) {
	QString _ret = QValidator::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QValidator_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QValidator::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QValidator_Delete(QValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QValidator*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQIntValidator : public virtual QIntValidator {
public:

	MiqtVirtualQIntValidator(): QIntValidator() {};
	MiqtVirtualQIntValidator(int bottom, int top): QIntValidator(bottom, top) {};
	MiqtVirtualQIntValidator(QObject* parent): QIntValidator(parent) {};
	MiqtVirtualQIntValidator(int bottom, int top, QObject* parent): QIntValidator(bottom, top, parent) {};

	virtual ~MiqtVirtualQIntValidator() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Validate = 0;

	// Subclass to allow providing a Go implementation
	virtual QValidator::State validate(QString& param1, int& param2) const override {
		if (handle__Validate == 0) {
			return QIntValidator::validate(param1, param2);
		}
		
		QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;
		int* sigval2 = &param2;

		int callback_return_value = miqt_exec_callback_QIntValidator_Validate(const_cast<MiqtVirtualQIntValidator*>(this), handle__Validate, sigval1, sigval2);

		return static_cast<QValidator::State>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Validate(struct miqt_string param1, int* param2) const {
		QString param1_QString = QString::fromUtf8(param1.data, param1.len);

		QValidator::State _ret = QIntValidator::validate(param1_QString, static_cast<int&>(*param2));
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Fixup = 0;

	// Subclass to allow providing a Go implementation
	virtual void fixup(QString& input) const override {
		if (handle__Fixup == 0) {
			QIntValidator::fixup(input);
			return;
		}
		
		QString input_ret = input;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray input_b = input_ret.toUtf8();
		struct miqt_string input_ms;
		input_ms.len = input_b.length();
		input_ms.data = static_cast<char*>(malloc(input_ms.len));
		memcpy(input_ms.data, input_b.data(), input_ms.len);
		struct miqt_string sigval1 = input_ms;

		miqt_exec_callback_QIntValidator_Fixup(const_cast<MiqtVirtualQIntValidator*>(this), handle__Fixup, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Fixup(struct miqt_string input) const {
		QString input_QString = QString::fromUtf8(input.data, input.len);

		QIntValidator::fixup(input_QString);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetRange = 0;

	// Subclass to allow providing a Go implementation
	virtual void setRange(int bottom, int top) override {
		if (handle__SetRange == 0) {
			QIntValidator::setRange(bottom, top);
			return;
		}
		
		int sigval1 = bottom;
		int sigval2 = top;

		miqt_exec_callback_QIntValidator_SetRange(this, handle__SetRange, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetRange(int bottom, int top) {

		QIntValidator::setRange(static_cast<int>(bottom), static_cast<int>(top));

	}

};

void QIntValidator_new(QIntValidator** outptr_QIntValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQIntValidator* ret = new MiqtVirtualQIntValidator();
	*outptr_QIntValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QIntValidator_new2(int bottom, int top, QIntValidator** outptr_QIntValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQIntValidator* ret = new MiqtVirtualQIntValidator(static_cast<int>(bottom), static_cast<int>(top));
	*outptr_QIntValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QIntValidator_new3(QObject* parent, QIntValidator** outptr_QIntValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQIntValidator* ret = new MiqtVirtualQIntValidator(parent);
	*outptr_QIntValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QIntValidator_new4(int bottom, int top, QObject* parent, QIntValidator** outptr_QIntValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQIntValidator* ret = new MiqtVirtualQIntValidator(static_cast<int>(bottom), static_cast<int>(top), parent);
	*outptr_QIntValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QIntValidator_MetaObject(const QIntValidator* self) {
	return (QMetaObject*) self->metaObject();
}

void* QIntValidator_Metacast(QIntValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QIntValidator_Tr(const char* s) {
	QString _ret = QIntValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QIntValidator_TrUtf8(const char* s) {
	QString _ret = QIntValidator::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QIntValidator_Validate(const QIntValidator* self, struct miqt_string param1, int* param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QValidator::State _ret = self->validate(param1_QString, static_cast<int&>(*param2));
	return static_cast<int>(_ret);
}

void QIntValidator_Fixup(const QIntValidator* self, struct miqt_string input) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	self->fixup(input_QString);
}

void QIntValidator_SetBottom(QIntValidator* self, int bottom) {
	self->setBottom(static_cast<int>(bottom));
}

void QIntValidator_SetTop(QIntValidator* self, int top) {
	self->setTop(static_cast<int>(top));
}

void QIntValidator_SetRange(QIntValidator* self, int bottom, int top) {
	self->setRange(static_cast<int>(bottom), static_cast<int>(top));
}

int QIntValidator_Bottom(const QIntValidator* self) {
	return self->bottom();
}

int QIntValidator_Top(const QIntValidator* self) {
	return self->top();
}

void QIntValidator_BottomChanged(QIntValidator* self, int bottom) {
	self->bottomChanged(static_cast<int>(bottom));
}

void QIntValidator_connect_BottomChanged(QIntValidator* self, intptr_t slot) {
	MiqtVirtualQIntValidator::connect(self, static_cast<void (QIntValidator::*)(int)>(&QIntValidator::bottomChanged), self, [=](int bottom) {
		int sigval1 = bottom;
		miqt_exec_callback_QIntValidator_BottomChanged(slot, sigval1);
	});
}

void QIntValidator_TopChanged(QIntValidator* self, int top) {
	self->topChanged(static_cast<int>(top));
}

void QIntValidator_connect_TopChanged(QIntValidator* self, intptr_t slot) {
	MiqtVirtualQIntValidator::connect(self, static_cast<void (QIntValidator::*)(int)>(&QIntValidator::topChanged), self, [=](int top) {
		int sigval1 = top;
		miqt_exec_callback_QIntValidator_TopChanged(slot, sigval1);
	});
}

struct miqt_string QIntValidator_Tr2(const char* s, const char* c) {
	QString _ret = QIntValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QIntValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QIntValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QIntValidator_TrUtf82(const char* s, const char* c) {
	QString _ret = QIntValidator::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QIntValidator_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QIntValidator::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QIntValidator_override_virtual_Validate(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIntValidator*>( (QIntValidator*)(self) )->handle__Validate = slot;
}

int QIntValidator_virtualbase_Validate(const void* self, struct miqt_string param1, int* param2) {
	return ( (const MiqtVirtualQIntValidator*)(self) )->virtualbase_Validate(param1, param2);
}

void QIntValidator_override_virtual_Fixup(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIntValidator*>( (QIntValidator*)(self) )->handle__Fixup = slot;
}

void QIntValidator_virtualbase_Fixup(const void* self, struct miqt_string input) {
	( (const MiqtVirtualQIntValidator*)(self) )->virtualbase_Fixup(input);
}

void QIntValidator_override_virtual_SetRange(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIntValidator*>( (QIntValidator*)(self) )->handle__SetRange = slot;
}

void QIntValidator_virtualbase_SetRange(void* self, int bottom, int top) {
	( (MiqtVirtualQIntValidator*)(self) )->virtualbase_SetRange(bottom, top);
}

void QIntValidator_Delete(QIntValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQIntValidator*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQDoubleValidator : public virtual QDoubleValidator {
public:

	MiqtVirtualQDoubleValidator(): QDoubleValidator() {};
	MiqtVirtualQDoubleValidator(double bottom, double top, int decimals): QDoubleValidator(bottom, top, decimals) {};
	MiqtVirtualQDoubleValidator(QObject* parent): QDoubleValidator(parent) {};
	MiqtVirtualQDoubleValidator(double bottom, double top, int decimals, QObject* parent): QDoubleValidator(bottom, top, decimals, parent) {};

	virtual ~MiqtVirtualQDoubleValidator() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Validate = 0;

	// Subclass to allow providing a Go implementation
	virtual QValidator::State validate(QString& param1, int& param2) const override {
		if (handle__Validate == 0) {
			return QDoubleValidator::validate(param1, param2);
		}
		
		QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;
		int* sigval2 = &param2;

		int callback_return_value = miqt_exec_callback_QDoubleValidator_Validate(const_cast<MiqtVirtualQDoubleValidator*>(this), handle__Validate, sigval1, sigval2);

		return static_cast<QValidator::State>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Validate(struct miqt_string param1, int* param2) const {
		QString param1_QString = QString::fromUtf8(param1.data, param1.len);

		QValidator::State _ret = QDoubleValidator::validate(param1_QString, static_cast<int&>(*param2));
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetRange = 0;

	// Subclass to allow providing a Go implementation
	virtual void setRange(double bottom, double top, int decimals) override {
		if (handle__SetRange == 0) {
			QDoubleValidator::setRange(bottom, top, decimals);
			return;
		}
		
		double sigval1 = bottom;
		double sigval2 = top;
		int sigval3 = decimals;

		miqt_exec_callback_QDoubleValidator_SetRange(this, handle__SetRange, sigval1, sigval2, sigval3);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetRange(double bottom, double top, int decimals) {

		QDoubleValidator::setRange(static_cast<double>(bottom), static_cast<double>(top), static_cast<int>(decimals));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Fixup = 0;

	// Subclass to allow providing a Go implementation
	virtual void fixup(QString& param1) const override {
		if (handle__Fixup == 0) {
			QDoubleValidator::fixup(param1);
			return;
		}
		
		QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;

		miqt_exec_callback_QDoubleValidator_Fixup(const_cast<MiqtVirtualQDoubleValidator*>(this), handle__Fixup, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Fixup(struct miqt_string param1) const {
		QString param1_QString = QString::fromUtf8(param1.data, param1.len);

		QDoubleValidator::fixup(param1_QString);

	}

};

void QDoubleValidator_new(QDoubleValidator** outptr_QDoubleValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQDoubleValidator* ret = new MiqtVirtualQDoubleValidator();
	*outptr_QDoubleValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QDoubleValidator_new2(double bottom, double top, int decimals, QDoubleValidator** outptr_QDoubleValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQDoubleValidator* ret = new MiqtVirtualQDoubleValidator(static_cast<double>(bottom), static_cast<double>(top), static_cast<int>(decimals));
	*outptr_QDoubleValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QDoubleValidator_new3(QObject* parent, QDoubleValidator** outptr_QDoubleValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQDoubleValidator* ret = new MiqtVirtualQDoubleValidator(parent);
	*outptr_QDoubleValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QDoubleValidator_new4(double bottom, double top, int decimals, QObject* parent, QDoubleValidator** outptr_QDoubleValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQDoubleValidator* ret = new MiqtVirtualQDoubleValidator(static_cast<double>(bottom), static_cast<double>(top), static_cast<int>(decimals), parent);
	*outptr_QDoubleValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QDoubleValidator_MetaObject(const QDoubleValidator* self) {
	return (QMetaObject*) self->metaObject();
}

void* QDoubleValidator_Metacast(QDoubleValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QDoubleValidator_Tr(const char* s) {
	QString _ret = QDoubleValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDoubleValidator_TrUtf8(const char* s) {
	QString _ret = QDoubleValidator::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QDoubleValidator_Validate(const QDoubleValidator* self, struct miqt_string param1, int* param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QValidator::State _ret = self->validate(param1_QString, static_cast<int&>(*param2));
	return static_cast<int>(_ret);
}

void QDoubleValidator_SetRange(QDoubleValidator* self, double bottom, double top, int decimals) {
	self->setRange(static_cast<double>(bottom), static_cast<double>(top), static_cast<int>(decimals));
}

void QDoubleValidator_SetBottom(QDoubleValidator* self, double bottom) {
	self->setBottom(static_cast<double>(bottom));
}

void QDoubleValidator_SetTop(QDoubleValidator* self, double top) {
	self->setTop(static_cast<double>(top));
}

void QDoubleValidator_SetDecimals(QDoubleValidator* self, int decimals) {
	self->setDecimals(static_cast<int>(decimals));
}

void QDoubleValidator_SetNotation(QDoubleValidator* self, int notation) {
	self->setNotation(static_cast<QDoubleValidator::Notation>(notation));
}

double QDoubleValidator_Bottom(const QDoubleValidator* self) {
	return self->bottom();
}

double QDoubleValidator_Top(const QDoubleValidator* self) {
	return self->top();
}

int QDoubleValidator_Decimals(const QDoubleValidator* self) {
	return self->decimals();
}

int QDoubleValidator_Notation(const QDoubleValidator* self) {
	QDoubleValidator::Notation _ret = self->notation();
	return static_cast<int>(_ret);
}

void QDoubleValidator_BottomChanged(QDoubleValidator* self, double bottom) {
	self->bottomChanged(static_cast<double>(bottom));
}

void QDoubleValidator_connect_BottomChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(double)>(&QDoubleValidator::bottomChanged), self, [=](double bottom) {
		double sigval1 = bottom;
		miqt_exec_callback_QDoubleValidator_BottomChanged(slot, sigval1);
	});
}

void QDoubleValidator_TopChanged(QDoubleValidator* self, double top) {
	self->topChanged(static_cast<double>(top));
}

void QDoubleValidator_connect_TopChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(double)>(&QDoubleValidator::topChanged), self, [=](double top) {
		double sigval1 = top;
		miqt_exec_callback_QDoubleValidator_TopChanged(slot, sigval1);
	});
}

void QDoubleValidator_DecimalsChanged(QDoubleValidator* self, int decimals) {
	self->decimalsChanged(static_cast<int>(decimals));
}

void QDoubleValidator_connect_DecimalsChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(int)>(&QDoubleValidator::decimalsChanged), self, [=](int decimals) {
		int sigval1 = decimals;
		miqt_exec_callback_QDoubleValidator_DecimalsChanged(slot, sigval1);
	});
}

void QDoubleValidator_NotationChanged(QDoubleValidator* self, int notation) {
	self->notationChanged(static_cast<QDoubleValidator::Notation>(notation));
}

void QDoubleValidator_connect_NotationChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(QDoubleValidator::Notation)>(&QDoubleValidator::notationChanged), self, [=](QDoubleValidator::Notation notation) {
		QDoubleValidator::Notation notation_ret = notation;
		int sigval1 = static_cast<int>(notation_ret);
		miqt_exec_callback_QDoubleValidator_NotationChanged(slot, sigval1);
	});
}

struct miqt_string QDoubleValidator_Tr2(const char* s, const char* c) {
	QString _ret = QDoubleValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDoubleValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDoubleValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDoubleValidator_TrUtf82(const char* s, const char* c) {
	QString _ret = QDoubleValidator::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDoubleValidator_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QDoubleValidator::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QDoubleValidator_override_virtual_Validate(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDoubleValidator*>( (QDoubleValidator*)(self) )->handle__Validate = slot;
}

int QDoubleValidator_virtualbase_Validate(const void* self, struct miqt_string param1, int* param2) {
	return ( (const MiqtVirtualQDoubleValidator*)(self) )->virtualbase_Validate(param1, param2);
}

void QDoubleValidator_override_virtual_SetRange(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDoubleValidator*>( (QDoubleValidator*)(self) )->handle__SetRange = slot;
}

void QDoubleValidator_virtualbase_SetRange(void* self, double bottom, double top, int decimals) {
	( (MiqtVirtualQDoubleValidator*)(self) )->virtualbase_SetRange(bottom, top, decimals);
}

void QDoubleValidator_override_virtual_Fixup(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDoubleValidator*>( (QDoubleValidator*)(self) )->handle__Fixup = slot;
}

void QDoubleValidator_virtualbase_Fixup(const void* self, struct miqt_string param1) {
	( (const MiqtVirtualQDoubleValidator*)(self) )->virtualbase_Fixup(param1);
}

void QDoubleValidator_Delete(QDoubleValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDoubleValidator*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQRegExpValidator : public virtual QRegExpValidator {
public:

	MiqtVirtualQRegExpValidator(): QRegExpValidator() {};
	MiqtVirtualQRegExpValidator(const QRegExp& rx): QRegExpValidator(rx) {};
	MiqtVirtualQRegExpValidator(QObject* parent): QRegExpValidator(parent) {};
	MiqtVirtualQRegExpValidator(const QRegExp& rx, QObject* parent): QRegExpValidator(rx, parent) {};

	virtual ~MiqtVirtualQRegExpValidator() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Validate = 0;

	// Subclass to allow providing a Go implementation
	virtual QValidator::State validate(QString& input, int& pos) const override {
		if (handle__Validate == 0) {
			return QRegExpValidator::validate(input, pos);
		}
		
		QString input_ret = input;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray input_b = input_ret.toUtf8();
		struct miqt_string input_ms;
		input_ms.len = input_b.length();
		input_ms.data = static_cast<char*>(malloc(input_ms.len));
		memcpy(input_ms.data, input_b.data(), input_ms.len);
		struct miqt_string sigval1 = input_ms;
		int* sigval2 = &pos;

		int callback_return_value = miqt_exec_callback_QRegExpValidator_Validate(const_cast<MiqtVirtualQRegExpValidator*>(this), handle__Validate, sigval1, sigval2);

		return static_cast<QValidator::State>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Validate(struct miqt_string input, int* pos) const {
		QString input_QString = QString::fromUtf8(input.data, input.len);

		QValidator::State _ret = QRegExpValidator::validate(input_QString, static_cast<int&>(*pos));
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Fixup = 0;

	// Subclass to allow providing a Go implementation
	virtual void fixup(QString& param1) const override {
		if (handle__Fixup == 0) {
			QRegExpValidator::fixup(param1);
			return;
		}
		
		QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;

		miqt_exec_callback_QRegExpValidator_Fixup(const_cast<MiqtVirtualQRegExpValidator*>(this), handle__Fixup, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Fixup(struct miqt_string param1) const {
		QString param1_QString = QString::fromUtf8(param1.data, param1.len);

		QRegExpValidator::fixup(param1_QString);

	}

};

void QRegExpValidator_new(QRegExpValidator** outptr_QRegExpValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegExpValidator* ret = new MiqtVirtualQRegExpValidator();
	*outptr_QRegExpValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QRegExpValidator_new2(QRegExp* rx, QRegExpValidator** outptr_QRegExpValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegExpValidator* ret = new MiqtVirtualQRegExpValidator(*rx);
	*outptr_QRegExpValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QRegExpValidator_new3(QObject* parent, QRegExpValidator** outptr_QRegExpValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegExpValidator* ret = new MiqtVirtualQRegExpValidator(parent);
	*outptr_QRegExpValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QRegExpValidator_new4(QRegExp* rx, QObject* parent, QRegExpValidator** outptr_QRegExpValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegExpValidator* ret = new MiqtVirtualQRegExpValidator(*rx, parent);
	*outptr_QRegExpValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QRegExpValidator_MetaObject(const QRegExpValidator* self) {
	return (QMetaObject*) self->metaObject();
}

void* QRegExpValidator_Metacast(QRegExpValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QRegExpValidator_Tr(const char* s) {
	QString _ret = QRegExpValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegExpValidator_TrUtf8(const char* s) {
	QString _ret = QRegExpValidator::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QRegExpValidator_Validate(const QRegExpValidator* self, struct miqt_string input, int* pos) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	QValidator::State _ret = self->validate(input_QString, static_cast<int&>(*pos));
	return static_cast<int>(_ret);
}

void QRegExpValidator_SetRegExp(QRegExpValidator* self, QRegExp* rx) {
	self->setRegExp(*rx);
}

QRegExp* QRegExpValidator_RegExp(const QRegExpValidator* self) {
	const QRegExp& _ret = self->regExp();
	// Cast returned reference into pointer
	return const_cast<QRegExp*>(&_ret);
}

void QRegExpValidator_RegExpChanged(QRegExpValidator* self, QRegExp* regExp) {
	self->regExpChanged(*regExp);
}

void QRegExpValidator_connect_RegExpChanged(QRegExpValidator* self, intptr_t slot) {
	MiqtVirtualQRegExpValidator::connect(self, static_cast<void (QRegExpValidator::*)(const QRegExp&)>(&QRegExpValidator::regExpChanged), self, [=](const QRegExp& regExp) {
		const QRegExp& regExp_ret = regExp;
		// Cast returned reference into pointer
		QRegExp* sigval1 = const_cast<QRegExp*>(&regExp_ret);
		miqt_exec_callback_QRegExpValidator_RegExpChanged(slot, sigval1);
	});
}

struct miqt_string QRegExpValidator_Tr2(const char* s, const char* c) {
	QString _ret = QRegExpValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegExpValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRegExpValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegExpValidator_TrUtf82(const char* s, const char* c) {
	QString _ret = QRegExpValidator::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegExpValidator_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QRegExpValidator::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QRegExpValidator_override_virtual_Validate(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRegExpValidator*>( (QRegExpValidator*)(self) )->handle__Validate = slot;
}

int QRegExpValidator_virtualbase_Validate(const void* self, struct miqt_string input, int* pos) {
	return ( (const MiqtVirtualQRegExpValidator*)(self) )->virtualbase_Validate(input, pos);
}

void QRegExpValidator_override_virtual_Fixup(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRegExpValidator*>( (QRegExpValidator*)(self) )->handle__Fixup = slot;
}

void QRegExpValidator_virtualbase_Fixup(const void* self, struct miqt_string param1) {
	( (const MiqtVirtualQRegExpValidator*)(self) )->virtualbase_Fixup(param1);
}

void QRegExpValidator_Delete(QRegExpValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRegExpValidator*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQRegularExpressionValidator : public virtual QRegularExpressionValidator {
public:

	MiqtVirtualQRegularExpressionValidator(): QRegularExpressionValidator() {};
	MiqtVirtualQRegularExpressionValidator(const QRegularExpression& re): QRegularExpressionValidator(re) {};
	MiqtVirtualQRegularExpressionValidator(QObject* parent): QRegularExpressionValidator(parent) {};
	MiqtVirtualQRegularExpressionValidator(const QRegularExpression& re, QObject* parent): QRegularExpressionValidator(re, parent) {};

	virtual ~MiqtVirtualQRegularExpressionValidator() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Validate = 0;

	// Subclass to allow providing a Go implementation
	virtual QValidator::State validate(QString& input, int& pos) const override {
		if (handle__Validate == 0) {
			return QRegularExpressionValidator::validate(input, pos);
		}
		
		QString input_ret = input;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray input_b = input_ret.toUtf8();
		struct miqt_string input_ms;
		input_ms.len = input_b.length();
		input_ms.data = static_cast<char*>(malloc(input_ms.len));
		memcpy(input_ms.data, input_b.data(), input_ms.len);
		struct miqt_string sigval1 = input_ms;
		int* sigval2 = &pos;

		int callback_return_value = miqt_exec_callback_QRegularExpressionValidator_Validate(const_cast<MiqtVirtualQRegularExpressionValidator*>(this), handle__Validate, sigval1, sigval2);

		return static_cast<QValidator::State>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Validate(struct miqt_string input, int* pos) const {
		QString input_QString = QString::fromUtf8(input.data, input.len);

		QValidator::State _ret = QRegularExpressionValidator::validate(input_QString, static_cast<int&>(*pos));
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Fixup = 0;

	// Subclass to allow providing a Go implementation
	virtual void fixup(QString& param1) const override {
		if (handle__Fixup == 0) {
			QRegularExpressionValidator::fixup(param1);
			return;
		}
		
		QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;

		miqt_exec_callback_QRegularExpressionValidator_Fixup(const_cast<MiqtVirtualQRegularExpressionValidator*>(this), handle__Fixup, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Fixup(struct miqt_string param1) const {
		QString param1_QString = QString::fromUtf8(param1.data, param1.len);

		QRegularExpressionValidator::fixup(param1_QString);

	}

};

void QRegularExpressionValidator_new(QRegularExpressionValidator** outptr_QRegularExpressionValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegularExpressionValidator* ret = new MiqtVirtualQRegularExpressionValidator();
	*outptr_QRegularExpressionValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QRegularExpressionValidator_new2(QRegularExpression* re, QRegularExpressionValidator** outptr_QRegularExpressionValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegularExpressionValidator* ret = new MiqtVirtualQRegularExpressionValidator(*re);
	*outptr_QRegularExpressionValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QRegularExpressionValidator_new3(QObject* parent, QRegularExpressionValidator** outptr_QRegularExpressionValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegularExpressionValidator* ret = new MiqtVirtualQRegularExpressionValidator(parent);
	*outptr_QRegularExpressionValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QRegularExpressionValidator_new4(QRegularExpression* re, QObject* parent, QRegularExpressionValidator** outptr_QRegularExpressionValidator, QValidator** outptr_QValidator, QObject** outptr_QObject) {
	MiqtVirtualQRegularExpressionValidator* ret = new MiqtVirtualQRegularExpressionValidator(*re, parent);
	*outptr_QRegularExpressionValidator = ret;
	*outptr_QValidator = static_cast<QValidator*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QRegularExpressionValidator_MetaObject(const QRegularExpressionValidator* self) {
	return (QMetaObject*) self->metaObject();
}

void* QRegularExpressionValidator_Metacast(QRegularExpressionValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QRegularExpressionValidator_Tr(const char* s) {
	QString _ret = QRegularExpressionValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegularExpressionValidator_TrUtf8(const char* s) {
	QString _ret = QRegularExpressionValidator::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QRegularExpressionValidator_Validate(const QRegularExpressionValidator* self, struct miqt_string input, int* pos) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	QValidator::State _ret = self->validate(input_QString, static_cast<int&>(*pos));
	return static_cast<int>(_ret);
}

QRegularExpression* QRegularExpressionValidator_RegularExpression(const QRegularExpressionValidator* self) {
	return new QRegularExpression(self->regularExpression());
}

void QRegularExpressionValidator_SetRegularExpression(QRegularExpressionValidator* self, QRegularExpression* re) {
	self->setRegularExpression(*re);
}

void QRegularExpressionValidator_RegularExpressionChanged(QRegularExpressionValidator* self, QRegularExpression* re) {
	self->regularExpressionChanged(*re);
}

void QRegularExpressionValidator_connect_RegularExpressionChanged(QRegularExpressionValidator* self, intptr_t slot) {
	MiqtVirtualQRegularExpressionValidator::connect(self, static_cast<void (QRegularExpressionValidator::*)(const QRegularExpression&)>(&QRegularExpressionValidator::regularExpressionChanged), self, [=](const QRegularExpression& re) {
		const QRegularExpression& re_ret = re;
		// Cast returned reference into pointer
		QRegularExpression* sigval1 = const_cast<QRegularExpression*>(&re_ret);
		miqt_exec_callback_QRegularExpressionValidator_RegularExpressionChanged(slot, sigval1);
	});
}

struct miqt_string QRegularExpressionValidator_Tr2(const char* s, const char* c) {
	QString _ret = QRegularExpressionValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegularExpressionValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRegularExpressionValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegularExpressionValidator_TrUtf82(const char* s, const char* c) {
	QString _ret = QRegularExpressionValidator::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegularExpressionValidator_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QRegularExpressionValidator::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QRegularExpressionValidator_override_virtual_Validate(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRegularExpressionValidator*>( (QRegularExpressionValidator*)(self) )->handle__Validate = slot;
}

int QRegularExpressionValidator_virtualbase_Validate(const void* self, struct miqt_string input, int* pos) {
	return ( (const MiqtVirtualQRegularExpressionValidator*)(self) )->virtualbase_Validate(input, pos);
}

void QRegularExpressionValidator_override_virtual_Fixup(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRegularExpressionValidator*>( (QRegularExpressionValidator*)(self) )->handle__Fixup = slot;
}

void QRegularExpressionValidator_virtualbase_Fixup(const void* self, struct miqt_string param1) {
	( (const MiqtVirtualQRegularExpressionValidator*)(self) )->virtualbase_Fixup(param1);
}

void QRegularExpressionValidator_Delete(QRegularExpressionValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRegularExpressionValidator*>( self );
	} else {
		delete self;
	}
}

