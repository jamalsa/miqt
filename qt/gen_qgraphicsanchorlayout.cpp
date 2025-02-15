#include <QEvent>
#include <QGraphicsAnchor>
#include <QGraphicsAnchorLayout>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QMetaObject>
#include <QObject>
#include <QRectF>
#include <QSizeF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qgraphicsanchorlayout.h>
#include "gen_qgraphicsanchorlayout.h"
#include "_cgo_export.h"

QMetaObject* QGraphicsAnchor_MetaObject(const QGraphicsAnchor* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGraphicsAnchor_Metacast(QGraphicsAnchor* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGraphicsAnchor_Tr(const char* s) {
	QString _ret = QGraphicsAnchor::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsAnchor_TrUtf8(const char* s) {
	QString _ret = QGraphicsAnchor::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsAnchor_SetSpacing(QGraphicsAnchor* self, double spacing) {
	self->setSpacing(static_cast<qreal>(spacing));
}

void QGraphicsAnchor_UnsetSpacing(QGraphicsAnchor* self) {
	self->unsetSpacing();
}

double QGraphicsAnchor_Spacing(const QGraphicsAnchor* self) {
	qreal _ret = self->spacing();
	return static_cast<double>(_ret);
}

void QGraphicsAnchor_SetSizePolicy(QGraphicsAnchor* self, int policy) {
	self->setSizePolicy(static_cast<QSizePolicy::Policy>(policy));
}

int QGraphicsAnchor_SizePolicy(const QGraphicsAnchor* self) {
	QSizePolicy::Policy _ret = self->sizePolicy();
	return static_cast<int>(_ret);
}

struct miqt_string QGraphicsAnchor_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsAnchor::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsAnchor_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsAnchor::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsAnchor_TrUtf82(const char* s, const char* c) {
	QString _ret = QGraphicsAnchor::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsAnchor_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QGraphicsAnchor::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsAnchor_Delete(QGraphicsAnchor* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QGraphicsAnchor*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQGraphicsAnchorLayout : public virtual QGraphicsAnchorLayout {
public:

	MiqtVirtualQGraphicsAnchorLayout(): QGraphicsAnchorLayout() {};
	MiqtVirtualQGraphicsAnchorLayout(QGraphicsLayoutItem* parent): QGraphicsAnchorLayout(parent) {};

	virtual ~MiqtVirtualQGraphicsAnchorLayout() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__RemoveAt = 0;

	// Subclass to allow providing a Go implementation
	virtual void removeAt(int index) override {
		if (handle__RemoveAt == 0) {
			QGraphicsAnchorLayout::removeAt(index);
			return;
		}
		
		int sigval1 = index;

		miqt_exec_callback_QGraphicsAnchorLayout_RemoveAt(this, handle__RemoveAt, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_RemoveAt(int index) {

		QGraphicsAnchorLayout::removeAt(static_cast<int>(index));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetGeometry = 0;

	// Subclass to allow providing a Go implementation
	virtual void setGeometry(const QRectF& rect) override {
		if (handle__SetGeometry == 0) {
			QGraphicsAnchorLayout::setGeometry(rect);
			return;
		}
		
		const QRectF& rect_ret = rect;
		// Cast returned reference into pointer
		QRectF* sigval1 = const_cast<QRectF*>(&rect_ret);

		miqt_exec_callback_QGraphicsAnchorLayout_SetGeometry(this, handle__SetGeometry, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetGeometry(QRectF* rect) {

		QGraphicsAnchorLayout::setGeometry(*rect);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Count = 0;

	// Subclass to allow providing a Go implementation
	virtual int count() const override {
		if (handle__Count == 0) {
			return QGraphicsAnchorLayout::count();
		}
		

		int callback_return_value = miqt_exec_callback_QGraphicsAnchorLayout_Count(const_cast<MiqtVirtualQGraphicsAnchorLayout*>(this), handle__Count);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Count() const {

		return QGraphicsAnchorLayout::count();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ItemAt = 0;

	// Subclass to allow providing a Go implementation
	virtual QGraphicsLayoutItem* itemAt(int index) const override {
		if (handle__ItemAt == 0) {
			return QGraphicsAnchorLayout::itemAt(index);
		}
		
		int sigval1 = index;

		QGraphicsLayoutItem* callback_return_value = miqt_exec_callback_QGraphicsAnchorLayout_ItemAt(const_cast<MiqtVirtualQGraphicsAnchorLayout*>(this), handle__ItemAt, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QGraphicsLayoutItem* virtualbase_ItemAt(int index) const {

		return QGraphicsAnchorLayout::itemAt(static_cast<int>(index));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Invalidate = 0;

	// Subclass to allow providing a Go implementation
	virtual void invalidate() override {
		if (handle__Invalidate == 0) {
			QGraphicsAnchorLayout::invalidate();
			return;
		}
		

		miqt_exec_callback_QGraphicsAnchorLayout_Invalidate(this, handle__Invalidate);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Invalidate() {

		QGraphicsAnchorLayout::invalidate();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		if (handle__SizeHint == 0) {
			return QGraphicsAnchorLayout::sizeHint(which, constraint);
		}
		
		Qt::SizeHint which_ret = which;
		int sigval1 = static_cast<int>(which_ret);
		const QSizeF& constraint_ret = constraint;
		// Cast returned reference into pointer
		QSizeF* sigval2 = const_cast<QSizeF*>(&constraint_ret);

		QSizeF* callback_return_value = miqt_exec_callback_QGraphicsAnchorLayout_SizeHint(const_cast<MiqtVirtualQGraphicsAnchorLayout*>(this), handle__SizeHint, sigval1, sigval2);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSizeF* virtualbase_SizeHint(int which, QSizeF* constraint) const {

		return new QSizeF(QGraphicsAnchorLayout::sizeHint(static_cast<Qt::SizeHint>(which), *constraint));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__GetContentsMargins = 0;

	// Subclass to allow providing a Go implementation
	virtual void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		if (handle__GetContentsMargins == 0) {
			QGraphicsAnchorLayout::getContentsMargins(left, top, right, bottom);
			return;
		}
		
		qreal* left_ret = left;
		double* sigval1 = static_cast<double*>(left_ret);
		qreal* top_ret = top;
		double* sigval2 = static_cast<double*>(top_ret);
		qreal* right_ret = right;
		double* sigval3 = static_cast<double*>(right_ret);
		qreal* bottom_ret = bottom;
		double* sigval4 = static_cast<double*>(bottom_ret);

		miqt_exec_callback_QGraphicsAnchorLayout_GetContentsMargins(const_cast<MiqtVirtualQGraphicsAnchorLayout*>(this), handle__GetContentsMargins, sigval1, sigval2, sigval3, sigval4);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_GetContentsMargins(double* left, double* top, double* right, double* bottom) const {

		QGraphicsAnchorLayout::getContentsMargins(static_cast<qreal*>(left), static_cast<qreal*>(top), static_cast<qreal*>(right), static_cast<qreal*>(bottom));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__UpdateGeometry = 0;

	// Subclass to allow providing a Go implementation
	virtual void updateGeometry() override {
		if (handle__UpdateGeometry == 0) {
			QGraphicsAnchorLayout::updateGeometry();
			return;
		}
		

		miqt_exec_callback_QGraphicsAnchorLayout_UpdateGeometry(this, handle__UpdateGeometry);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_UpdateGeometry() {

		QGraphicsAnchorLayout::updateGeometry();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WidgetEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void widgetEvent(QEvent* e) override {
		if (handle__WidgetEvent == 0) {
			QGraphicsAnchorLayout::widgetEvent(e);
			return;
		}
		
		QEvent* sigval1 = e;

		miqt_exec_callback_QGraphicsAnchorLayout_WidgetEvent(this, handle__WidgetEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_WidgetEvent(QEvent* e) {

		QGraphicsAnchorLayout::widgetEvent(e);

	}

};

void QGraphicsAnchorLayout_new(QGraphicsAnchorLayout** outptr_QGraphicsAnchorLayout, QGraphicsLayout** outptr_QGraphicsLayout, QGraphicsLayoutItem** outptr_QGraphicsLayoutItem) {
	MiqtVirtualQGraphicsAnchorLayout* ret = new MiqtVirtualQGraphicsAnchorLayout();
	*outptr_QGraphicsAnchorLayout = ret;
	*outptr_QGraphicsLayout = static_cast<QGraphicsLayout*>(ret);
	*outptr_QGraphicsLayoutItem = static_cast<QGraphicsLayoutItem*>(ret);
}

void QGraphicsAnchorLayout_new2(QGraphicsLayoutItem* parent, QGraphicsAnchorLayout** outptr_QGraphicsAnchorLayout, QGraphicsLayout** outptr_QGraphicsLayout, QGraphicsLayoutItem** outptr_QGraphicsLayoutItem) {
	MiqtVirtualQGraphicsAnchorLayout* ret = new MiqtVirtualQGraphicsAnchorLayout(parent);
	*outptr_QGraphicsAnchorLayout = ret;
	*outptr_QGraphicsLayout = static_cast<QGraphicsLayout*>(ret);
	*outptr_QGraphicsLayoutItem = static_cast<QGraphicsLayoutItem*>(ret);
}

QGraphicsAnchor* QGraphicsAnchorLayout_AddAnchor(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstEdge, QGraphicsLayoutItem* secondItem, int secondEdge) {
	return self->addAnchor(firstItem, static_cast<Qt::AnchorPoint>(firstEdge), secondItem, static_cast<Qt::AnchorPoint>(secondEdge));
}

QGraphicsAnchor* QGraphicsAnchorLayout_Anchor(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstEdge, QGraphicsLayoutItem* secondItem, int secondEdge) {
	return self->anchor(firstItem, static_cast<Qt::AnchorPoint>(firstEdge), secondItem, static_cast<Qt::AnchorPoint>(secondEdge));
}

void QGraphicsAnchorLayout_AddCornerAnchors(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstCorner, QGraphicsLayoutItem* secondItem, int secondCorner) {
	self->addCornerAnchors(firstItem, static_cast<Qt::Corner>(firstCorner), secondItem, static_cast<Qt::Corner>(secondCorner));
}

void QGraphicsAnchorLayout_AddAnchors(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, QGraphicsLayoutItem* secondItem) {
	self->addAnchors(firstItem, secondItem);
}

void QGraphicsAnchorLayout_SetHorizontalSpacing(QGraphicsAnchorLayout* self, double spacing) {
	self->setHorizontalSpacing(static_cast<qreal>(spacing));
}

void QGraphicsAnchorLayout_SetVerticalSpacing(QGraphicsAnchorLayout* self, double spacing) {
	self->setVerticalSpacing(static_cast<qreal>(spacing));
}

void QGraphicsAnchorLayout_SetSpacing(QGraphicsAnchorLayout* self, double spacing) {
	self->setSpacing(static_cast<qreal>(spacing));
}

double QGraphicsAnchorLayout_HorizontalSpacing(const QGraphicsAnchorLayout* self) {
	qreal _ret = self->horizontalSpacing();
	return static_cast<double>(_ret);
}

double QGraphicsAnchorLayout_VerticalSpacing(const QGraphicsAnchorLayout* self) {
	qreal _ret = self->verticalSpacing();
	return static_cast<double>(_ret);
}

void QGraphicsAnchorLayout_RemoveAt(QGraphicsAnchorLayout* self, int index) {
	self->removeAt(static_cast<int>(index));
}

void QGraphicsAnchorLayout_SetGeometry(QGraphicsAnchorLayout* self, QRectF* rect) {
	self->setGeometry(*rect);
}

int QGraphicsAnchorLayout_Count(const QGraphicsAnchorLayout* self) {
	return self->count();
}

QGraphicsLayoutItem* QGraphicsAnchorLayout_ItemAt(const QGraphicsAnchorLayout* self, int index) {
	return self->itemAt(static_cast<int>(index));
}

void QGraphicsAnchorLayout_Invalidate(QGraphicsAnchorLayout* self) {
	self->invalidate();
}

void QGraphicsAnchorLayout_AddAnchors3(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, QGraphicsLayoutItem* secondItem, int orientations) {
	self->addAnchors(firstItem, secondItem, static_cast<Qt::Orientations>(orientations));
}

void QGraphicsAnchorLayout_override_virtual_RemoveAt(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__RemoveAt = slot;
}

void QGraphicsAnchorLayout_virtualbase_RemoveAt(void* self, int index) {
	( (MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_RemoveAt(index);
}

void QGraphicsAnchorLayout_override_virtual_SetGeometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__SetGeometry = slot;
}

void QGraphicsAnchorLayout_virtualbase_SetGeometry(void* self, QRectF* rect) {
	( (MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_SetGeometry(rect);
}

void QGraphicsAnchorLayout_override_virtual_Count(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__Count = slot;
}

int QGraphicsAnchorLayout_virtualbase_Count(const void* self) {
	return ( (const MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_Count();
}

void QGraphicsAnchorLayout_override_virtual_ItemAt(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__ItemAt = slot;
}

QGraphicsLayoutItem* QGraphicsAnchorLayout_virtualbase_ItemAt(const void* self, int index) {
	return ( (const MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_ItemAt(index);
}

void QGraphicsAnchorLayout_override_virtual_Invalidate(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__Invalidate = slot;
}

void QGraphicsAnchorLayout_virtualbase_Invalidate(void* self) {
	( (MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_Invalidate();
}

void QGraphicsAnchorLayout_override_virtual_SizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__SizeHint = slot;
}

QSizeF* QGraphicsAnchorLayout_virtualbase_SizeHint(const void* self, int which, QSizeF* constraint) {
	return ( (const MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_SizeHint(which, constraint);
}

void QGraphicsAnchorLayout_override_virtual_GetContentsMargins(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__GetContentsMargins = slot;
}

void QGraphicsAnchorLayout_virtualbase_GetContentsMargins(const void* self, double* left, double* top, double* right, double* bottom) {
	( (const MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_GetContentsMargins(left, top, right, bottom);
}

void QGraphicsAnchorLayout_override_virtual_UpdateGeometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__UpdateGeometry = slot;
}

void QGraphicsAnchorLayout_virtualbase_UpdateGeometry(void* self) {
	( (MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_UpdateGeometry();
}

void QGraphicsAnchorLayout_override_virtual_WidgetEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( (QGraphicsAnchorLayout*)(self) )->handle__WidgetEvent = slot;
}

void QGraphicsAnchorLayout_virtualbase_WidgetEvent(void* self, QEvent* e) {
	( (MiqtVirtualQGraphicsAnchorLayout*)(self) )->virtualbase_WidgetEvent(e);
}

void QGraphicsAnchorLayout_Delete(QGraphicsAnchorLayout* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsAnchorLayout*>( self );
	} else {
		delete self;
	}
}

