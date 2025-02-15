#include <QIconEngine>
#include <QIconEnginePlugin>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qiconengineplugin.h>
#include "gen_qiconengineplugin.h"
#include "_cgo_export.h"

QMetaObject* QIconEnginePlugin_MetaObject(const QIconEnginePlugin* self) {
	return (QMetaObject*) self->metaObject();
}

void* QIconEnginePlugin_Metacast(QIconEnginePlugin* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QIconEnginePlugin_Tr(const char* s) {
	QString _ret = QIconEnginePlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QIconEngine* QIconEnginePlugin_Create(QIconEnginePlugin* self, struct miqt_string filename) {
	QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return self->create(filename_QString);
}

struct miqt_string QIconEnginePlugin_Tr2(const char* s, const char* c) {
	QString _ret = QIconEnginePlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QIconEnginePlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QIconEnginePlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QIconEnginePlugin_Delete(QIconEnginePlugin* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QIconEnginePlugin*>( self );
	} else {
		delete self;
	}
}

