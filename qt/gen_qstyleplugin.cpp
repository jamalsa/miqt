#include <QMetaObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyle>
#include <QStylePlugin>
#include "qstyleplugin.h"
#include "gen_qstyleplugin.h"
#include "_cgo_export.h"

QMetaObject* QStylePlugin_MetaObject(const QStylePlugin* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QStylePlugin_Tr(const char* s) {
	QString _ret = QStylePlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QStylePlugin_TrUtf8(const char* s) {
	QString _ret = QStylePlugin::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QStyle* QStylePlugin_Create(QStylePlugin* self, struct miqt_string* key) {
	QString key_QString = QString::fromUtf8(&key->data, key->len);
	return self->create(key_QString);
}

struct miqt_string* QStylePlugin_Tr2(const char* s, const char* c) {
	QString _ret = QStylePlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QStylePlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStylePlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QStylePlugin_TrUtf82(const char* s, const char* c) {
	QString _ret = QStylePlugin::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QStylePlugin_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QStylePlugin::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QStylePlugin_Delete(QStylePlugin* self) {
	delete self;
}

