#include <QByteArray>
#include <QList>
#include <QMetaObject>
#include <QMimeData>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <QVariant>
#include "qmimedata.h"
#include "gen_qmimedata.h"
#include "_cgo_export.h"

QMimeData* QMimeData_new() {
	return new QMimeData();
}

QMetaObject* QMimeData_MetaObject(const QMimeData* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QMimeData_Tr(const char* s) {
	QString _ret = QMimeData::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QMimeData_TrUtf8(const char* s) {
	QString _ret = QMimeData::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_array* QMimeData_Urls(const QMimeData* self) {
	QList<QUrl> _ret = self->urls();
	// Convert QList<> from C++ memory to manually-managed C memory of copy-constructed pointers
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl**) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QMimeData_SetUrls(QMimeData* self, struct miqt_array* /* of QUrl */ urls) {
	QList<QUrl> urls_QList;
	urls_QList.reserve(urls->len);
	QUrl** urls_arr = static_cast<QUrl**>(urls->data);
	for(size_t i = 0; i < urls->len; ++i) {
		urls_QList.push_back(*(urls_arr[i]));
	}
	self->setUrls(urls_QList);
}

bool QMimeData_HasUrls(const QMimeData* self) {
	return self->hasUrls();
}

struct miqt_string* QMimeData_Text(const QMimeData* self) {
	QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QMimeData_SetText(QMimeData* self, struct miqt_string* text) {
	QString text_QString = QString::fromUtf8(&text->data, text->len);
	self->setText(text_QString);
}

bool QMimeData_HasText(const QMimeData* self) {
	return self->hasText();
}

struct miqt_string* QMimeData_Html(const QMimeData* self) {
	QString _ret = self->html();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QMimeData_SetHtml(QMimeData* self, struct miqt_string* html) {
	QString html_QString = QString::fromUtf8(&html->data, html->len);
	self->setHtml(html_QString);
}

bool QMimeData_HasHtml(const QMimeData* self) {
	return self->hasHtml();
}

QVariant* QMimeData_ImageData(const QMimeData* self) {
	QVariant _ret = self->imageData();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(_ret));
}

void QMimeData_SetImageData(QMimeData* self, QVariant* image) {
	self->setImageData(*image);
}

bool QMimeData_HasImage(const QMimeData* self) {
	return self->hasImage();
}

QVariant* QMimeData_ColorData(const QMimeData* self) {
	QVariant _ret = self->colorData();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(_ret));
}

void QMimeData_SetColorData(QMimeData* self, QVariant* color) {
	self->setColorData(*color);
}

bool QMimeData_HasColor(const QMimeData* self) {
	return self->hasColor();
}

QByteArray* QMimeData_Data(const QMimeData* self, struct miqt_string* mimetype) {
	QString mimetype_QString = QString::fromUtf8(&mimetype->data, mimetype->len);
	QByteArray _ret = self->data(mimetype_QString);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(_ret));
}

void QMimeData_SetData(QMimeData* self, struct miqt_string* mimetype, QByteArray* data) {
	QString mimetype_QString = QString::fromUtf8(&mimetype->data, mimetype->len);
	self->setData(mimetype_QString, *data);
}

void QMimeData_RemoveFormat(QMimeData* self, struct miqt_string* mimetype) {
	QString mimetype_QString = QString::fromUtf8(&mimetype->data, mimetype->len);
	self->removeFormat(mimetype_QString);
}

bool QMimeData_HasFormat(const QMimeData* self, struct miqt_string* mimetype) {
	QString mimetype_QString = QString::fromUtf8(&mimetype->data, mimetype->len);
	return self->hasFormat(mimetype_QString);
}

struct miqt_array* QMimeData_Formats(const QMimeData* self) {
	QStringList _ret = self->formats();
	// Convert QStringList from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QMimeData_Clear(QMimeData* self) {
	self->clear();
}

struct miqt_string* QMimeData_Tr2(const char* s, const char* c) {
	QString _ret = QMimeData::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QMimeData_Tr3(const char* s, const char* c, int n) {
	QString _ret = QMimeData::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QMimeData_TrUtf82(const char* s, const char* c) {
	QString _ret = QMimeData::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QMimeData_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QMimeData::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QMimeData_Delete(QMimeData* self) {
	delete self;
}

