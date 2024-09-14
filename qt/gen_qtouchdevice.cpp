#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTouchDevice>
#include "qtouchdevice.h"
#include "gen_qtouchdevice.h"
#include "_cgo_export.h"

QTouchDevice* QTouchDevice_new() {
	return new QTouchDevice();
}

struct miqt_array* QTouchDevice_Devices() {
	QList<const QTouchDevice*> _ret = QTouchDevice::devices();
	// Convert QList<> from C++ memory to manually-managed C memory
	const QTouchDevice** _arr = static_cast<const QTouchDevice**>(malloc(sizeof(const QTouchDevice*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = const_cast<QTouchDevice*>(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string* QTouchDevice_Name(const QTouchDevice* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

uintptr_t QTouchDevice_Type(const QTouchDevice* self) {
	QTouchDevice::DeviceType _ret = self->type();
	return static_cast<uintptr_t>(_ret);
}

int QTouchDevice_Capabilities(const QTouchDevice* self) {
	QTouchDevice::Capabilities _ret = self->capabilities();
	return static_cast<int>(_ret);
}

int QTouchDevice_MaximumTouchPoints(const QTouchDevice* self) {
	return self->maximumTouchPoints();
}

void QTouchDevice_SetName(QTouchDevice* self, struct miqt_string* name) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	self->setName(name_QString);
}

void QTouchDevice_SetType(QTouchDevice* self, uintptr_t devType) {
	self->setType(static_cast<QTouchDevice::DeviceType>(devType));
}

void QTouchDevice_SetCapabilities(QTouchDevice* self, int caps) {
	self->setCapabilities(static_cast<QTouchDevice::Capabilities>(caps));
}

void QTouchDevice_SetMaximumTouchPoints(QTouchDevice* self, int max) {
	self->setMaximumTouchPoints(static_cast<int>(max));
}

void QTouchDevice_Delete(QTouchDevice* self) {
	delete self;
}

