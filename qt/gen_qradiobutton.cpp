#include <QAbstractButton>
#include <QEvent>
#include <QFocusEvent>
#include <QKeyEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QRadioButton>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <QWidget>
#include <qradiobutton.h>
#include "gen_qradiobutton.h"
#include "_cgo_export.h"

class MiqtVirtualQRadioButton : public virtual QRadioButton {
public:

	MiqtVirtualQRadioButton(QWidget* parent): QRadioButton(parent) {};
	MiqtVirtualQRadioButton(): QRadioButton() {};
	MiqtVirtualQRadioButton(const QString& text): QRadioButton(text) {};
	MiqtVirtualQRadioButton(const QString& text, QWidget* parent): QRadioButton(text, parent) {};

	virtual ~MiqtVirtualQRadioButton() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QRadioButton::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QRadioButton_SizeHint(const_cast<MiqtVirtualQRadioButton*>(this), handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QRadioButton::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSizeHint() const override {
		if (handle__MinimumSizeHint == 0) {
			return QRadioButton::minimumSizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QRadioButton_MinimumSizeHint(const_cast<MiqtVirtualQRadioButton*>(this), handle__MinimumSizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSizeHint() const {

		return new QSize(QRadioButton::minimumSizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* e) override {
		if (handle__Event == 0) {
			return QRadioButton::event(e);
		}
		
		QEvent* sigval1 = e;

		bool callback_return_value = miqt_exec_callback_QRadioButton_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* e) {

		return QRadioButton::event(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HitButton = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hitButton(const QPoint& param1) const override {
		if (handle__HitButton == 0) {
			return QRadioButton::hitButton(param1);
		}
		
		const QPoint& param1_ret = param1;
		// Cast returned reference into pointer
		QPoint* sigval1 = const_cast<QPoint*>(&param1_ret);

		bool callback_return_value = miqt_exec_callback_QRadioButton_HitButton(const_cast<MiqtVirtualQRadioButton*>(this), handle__HitButton, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HitButton(QPoint* param1) const {

		return QRadioButton::hitButton(*param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void paintEvent(QPaintEvent* param1) override {
		if (handle__PaintEvent == 0) {
			QRadioButton::paintEvent(param1);
			return;
		}
		
		QPaintEvent* sigval1 = param1;

		miqt_exec_callback_QRadioButton_PaintEvent(this, handle__PaintEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_PaintEvent(QPaintEvent* param1) {

		QRadioButton::paintEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseMoveEvent(QMouseEvent* param1) override {
		if (handle__MouseMoveEvent == 0) {
			QRadioButton::mouseMoveEvent(param1);
			return;
		}
		
		QMouseEvent* sigval1 = param1;

		miqt_exec_callback_QRadioButton_MouseMoveEvent(this, handle__MouseMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseMoveEvent(QMouseEvent* param1) {

		QRadioButton::mouseMoveEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CheckStateSet = 0;

	// Subclass to allow providing a Go implementation
	virtual void checkStateSet() override {
		if (handle__CheckStateSet == 0) {
			QRadioButton::checkStateSet();
			return;
		}
		

		miqt_exec_callback_QRadioButton_CheckStateSet(this, handle__CheckStateSet);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CheckStateSet() {

		QRadioButton::checkStateSet();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__NextCheckState = 0;

	// Subclass to allow providing a Go implementation
	virtual void nextCheckState() override {
		if (handle__NextCheckState == 0) {
			QRadioButton::nextCheckState();
			return;
		}
		

		miqt_exec_callback_QRadioButton_NextCheckState(this, handle__NextCheckState);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_NextCheckState() {

		QRadioButton::nextCheckState();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyPressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyPressEvent(QKeyEvent* e) override {
		if (handle__KeyPressEvent == 0) {
			QRadioButton::keyPressEvent(e);
			return;
		}
		
		QKeyEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_KeyPressEvent(this, handle__KeyPressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyPressEvent(QKeyEvent* e) {

		QRadioButton::keyPressEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyReleaseEvent(QKeyEvent* e) override {
		if (handle__KeyReleaseEvent == 0) {
			QRadioButton::keyReleaseEvent(e);
			return;
		}
		
		QKeyEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_KeyReleaseEvent(this, handle__KeyReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyReleaseEvent(QKeyEvent* e) {

		QRadioButton::keyReleaseEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MousePressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mousePressEvent(QMouseEvent* e) override {
		if (handle__MousePressEvent == 0) {
			QRadioButton::mousePressEvent(e);
			return;
		}
		
		QMouseEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_MousePressEvent(this, handle__MousePressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MousePressEvent(QMouseEvent* e) {

		QRadioButton::mousePressEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseReleaseEvent(QMouseEvent* e) override {
		if (handle__MouseReleaseEvent == 0) {
			QRadioButton::mouseReleaseEvent(e);
			return;
		}
		
		QMouseEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_MouseReleaseEvent(this, handle__MouseReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseReleaseEvent(QMouseEvent* e) {

		QRadioButton::mouseReleaseEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusInEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusInEvent(QFocusEvent* e) override {
		if (handle__FocusInEvent == 0) {
			QRadioButton::focusInEvent(e);
			return;
		}
		
		QFocusEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_FocusInEvent(this, handle__FocusInEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusInEvent(QFocusEvent* e) {

		QRadioButton::focusInEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusOutEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusOutEvent(QFocusEvent* e) override {
		if (handle__FocusOutEvent == 0) {
			QRadioButton::focusOutEvent(e);
			return;
		}
		
		QFocusEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_FocusOutEvent(this, handle__FocusOutEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusOutEvent(QFocusEvent* e) {

		QRadioButton::focusOutEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChangeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void changeEvent(QEvent* e) override {
		if (handle__ChangeEvent == 0) {
			QRadioButton::changeEvent(e);
			return;
		}
		
		QEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_ChangeEvent(this, handle__ChangeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChangeEvent(QEvent* e) {

		QRadioButton::changeEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* e) override {
		if (handle__TimerEvent == 0) {
			QRadioButton::timerEvent(e);
			return;
		}
		
		QTimerEvent* sigval1 = e;

		miqt_exec_callback_QRadioButton_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* e) {

		QRadioButton::timerEvent(e);

	}

};

void QRadioButton_new(QWidget* parent, QRadioButton** outptr_QRadioButton, QAbstractButton** outptr_QAbstractButton, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice) {
	MiqtVirtualQRadioButton* ret = new MiqtVirtualQRadioButton(parent);
	*outptr_QRadioButton = ret;
	*outptr_QAbstractButton = static_cast<QAbstractButton*>(ret);
	*outptr_QWidget = static_cast<QWidget*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
	*outptr_QPaintDevice = static_cast<QPaintDevice*>(ret);
}

void QRadioButton_new2(QRadioButton** outptr_QRadioButton, QAbstractButton** outptr_QAbstractButton, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice) {
	MiqtVirtualQRadioButton* ret = new MiqtVirtualQRadioButton();
	*outptr_QRadioButton = ret;
	*outptr_QAbstractButton = static_cast<QAbstractButton*>(ret);
	*outptr_QWidget = static_cast<QWidget*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
	*outptr_QPaintDevice = static_cast<QPaintDevice*>(ret);
}

void QRadioButton_new3(struct miqt_string text, QRadioButton** outptr_QRadioButton, QAbstractButton** outptr_QAbstractButton, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	MiqtVirtualQRadioButton* ret = new MiqtVirtualQRadioButton(text_QString);
	*outptr_QRadioButton = ret;
	*outptr_QAbstractButton = static_cast<QAbstractButton*>(ret);
	*outptr_QWidget = static_cast<QWidget*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
	*outptr_QPaintDevice = static_cast<QPaintDevice*>(ret);
}

void QRadioButton_new4(struct miqt_string text, QWidget* parent, QRadioButton** outptr_QRadioButton, QAbstractButton** outptr_QAbstractButton, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	MiqtVirtualQRadioButton* ret = new MiqtVirtualQRadioButton(text_QString, parent);
	*outptr_QRadioButton = ret;
	*outptr_QAbstractButton = static_cast<QAbstractButton*>(ret);
	*outptr_QWidget = static_cast<QWidget*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
	*outptr_QPaintDevice = static_cast<QPaintDevice*>(ret);
}

QMetaObject* QRadioButton_MetaObject(const QRadioButton* self) {
	return (QMetaObject*) self->metaObject();
}

void* QRadioButton_Metacast(QRadioButton* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QRadioButton_Tr(const char* s) {
	QString _ret = QRadioButton::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRadioButton_TrUtf8(const char* s) {
	QString _ret = QRadioButton::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QSize* QRadioButton_SizeHint(const QRadioButton* self) {
	return new QSize(self->sizeHint());
}

QSize* QRadioButton_MinimumSizeHint(const QRadioButton* self) {
	return new QSize(self->minimumSizeHint());
}

struct miqt_string QRadioButton_Tr2(const char* s, const char* c) {
	QString _ret = QRadioButton::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRadioButton_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRadioButton::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRadioButton_TrUtf82(const char* s, const char* c) {
	QString _ret = QRadioButton::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRadioButton_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QRadioButton::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QRadioButton_override_virtual_SizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__SizeHint = slot;
}

QSize* QRadioButton_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQRadioButton*)(self) )->virtualbase_SizeHint();
}

void QRadioButton_override_virtual_MinimumSizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__MinimumSizeHint = slot;
}

QSize* QRadioButton_virtualbase_MinimumSizeHint(const void* self) {
	return ( (const MiqtVirtualQRadioButton*)(self) )->virtualbase_MinimumSizeHint();
}

void QRadioButton_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__Event = slot;
}

bool QRadioButton_virtualbase_Event(void* self, QEvent* e) {
	return ( (MiqtVirtualQRadioButton*)(self) )->virtualbase_Event(e);
}

void QRadioButton_override_virtual_HitButton(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__HitButton = slot;
}

bool QRadioButton_virtualbase_HitButton(const void* self, QPoint* param1) {
	return ( (const MiqtVirtualQRadioButton*)(self) )->virtualbase_HitButton(param1);
}

void QRadioButton_override_virtual_PaintEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__PaintEvent = slot;
}

void QRadioButton_virtualbase_PaintEvent(void* self, QPaintEvent* param1) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_PaintEvent(param1);
}

void QRadioButton_override_virtual_MouseMoveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__MouseMoveEvent = slot;
}

void QRadioButton_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_MouseMoveEvent(param1);
}

void QRadioButton_override_virtual_CheckStateSet(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__CheckStateSet = slot;
}

void QRadioButton_virtualbase_CheckStateSet(void* self) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_CheckStateSet();
}

void QRadioButton_override_virtual_NextCheckState(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__NextCheckState = slot;
}

void QRadioButton_virtualbase_NextCheckState(void* self) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_NextCheckState();
}

void QRadioButton_override_virtual_KeyPressEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__KeyPressEvent = slot;
}

void QRadioButton_virtualbase_KeyPressEvent(void* self, QKeyEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_KeyPressEvent(e);
}

void QRadioButton_override_virtual_KeyReleaseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__KeyReleaseEvent = slot;
}

void QRadioButton_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_KeyReleaseEvent(e);
}

void QRadioButton_override_virtual_MousePressEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__MousePressEvent = slot;
}

void QRadioButton_virtualbase_MousePressEvent(void* self, QMouseEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_MousePressEvent(e);
}

void QRadioButton_override_virtual_MouseReleaseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__MouseReleaseEvent = slot;
}

void QRadioButton_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_MouseReleaseEvent(e);
}

void QRadioButton_override_virtual_FocusInEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__FocusInEvent = slot;
}

void QRadioButton_virtualbase_FocusInEvent(void* self, QFocusEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_FocusInEvent(e);
}

void QRadioButton_override_virtual_FocusOutEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__FocusOutEvent = slot;
}

void QRadioButton_virtualbase_FocusOutEvent(void* self, QFocusEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_FocusOutEvent(e);
}

void QRadioButton_override_virtual_ChangeEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__ChangeEvent = slot;
}

void QRadioButton_virtualbase_ChangeEvent(void* self, QEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_ChangeEvent(e);
}

void QRadioButton_override_virtual_TimerEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__TimerEvent = slot;
}

void QRadioButton_virtualbase_TimerEvent(void* self, QTimerEvent* e) {
	( (MiqtVirtualQRadioButton*)(self) )->virtualbase_TimerEvent(e);
}

void QRadioButton_Delete(QRadioButton* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRadioButton*>( self );
	} else {
		delete self;
	}
}

