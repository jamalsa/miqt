#include <QGenericPlugin>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include "qgenericplugin.h"
#include "gen_qgenericplugin.h"
#include "_cgo_export.h"

QMetaObject* QGenericPlugin_MetaObject(const QGenericPlugin* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QGenericPlugin_Tr(const char* s) {
	QString _ret = QGenericPlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGenericPlugin_TrUtf8(const char* s) {
	QString _ret = QGenericPlugin::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QObject* QGenericPlugin_Create(QGenericPlugin* self, struct miqt_string* name, struct miqt_string* spec) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	QString spec_QString = QString::fromUtf8(&spec->data, spec->len);
	return self->create(name_QString, spec_QString);
}

struct miqt_string* QGenericPlugin_Tr2(const char* s, const char* c) {
	QString _ret = QGenericPlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGenericPlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGenericPlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGenericPlugin_TrUtf82(const char* s, const char* c) {
	QString _ret = QGenericPlugin::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGenericPlugin_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QGenericPlugin::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QGenericPlugin_Delete(QGenericPlugin* self) {
	delete self;
}

