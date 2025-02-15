#include <QMetaObject>
#include <QObject>
#include <QPicture>
#include <QPictureFormatPlugin>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qpictureformatplugin.h>
#include "gen_qpictureformatplugin.h"
#include "_cgo_export.h"

QMetaObject* QPictureFormatPlugin_MetaObject(const QPictureFormatPlugin* self) {
	return (QMetaObject*) self->metaObject();
}

void* QPictureFormatPlugin_Metacast(QPictureFormatPlugin* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QPictureFormatPlugin_Tr(const char* s) {
	QString _ret = QPictureFormatPlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPictureFormatPlugin_TrUtf8(const char* s) {
	QString _ret = QPictureFormatPlugin::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QPictureFormatPlugin_LoadPicture(QPictureFormatPlugin* self, struct miqt_string format, struct miqt_string filename, QPicture* pic) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return self->loadPicture(format_QString, filename_QString, pic);
}

bool QPictureFormatPlugin_SavePicture(QPictureFormatPlugin* self, struct miqt_string format, struct miqt_string filename, QPicture* pic) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return self->savePicture(format_QString, filename_QString, *pic);
}

bool QPictureFormatPlugin_InstallIOHandler(QPictureFormatPlugin* self, struct miqt_string format) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return self->installIOHandler(format_QString);
}

struct miqt_string QPictureFormatPlugin_Tr2(const char* s, const char* c) {
	QString _ret = QPictureFormatPlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPictureFormatPlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPictureFormatPlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPictureFormatPlugin_TrUtf82(const char* s, const char* c) {
	QString _ret = QPictureFormatPlugin::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPictureFormatPlugin_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QPictureFormatPlugin::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPictureFormatPlugin_Delete(QPictureFormatPlugin* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPictureFormatPlugin*>( self );
	} else {
		delete self;
	}
}

