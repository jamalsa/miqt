#include <QLayout>
#include <QLayoutItem>
#include <QRect>
#include <QSize>
#include <QSizePolicy>
#include <QSpacerItem>
#include <QWidget>
#include <QWidgetItem>
#include <QWidgetItemV2>
#include <qlayoutitem.h>
#include "gen_qlayoutitem.h"
#include "_cgo_export.h"

QSize* QLayoutItem_SizeHint(const QLayoutItem* self) {
	return new QSize(self->sizeHint());
}

QSize* QLayoutItem_MinimumSize(const QLayoutItem* self) {
	return new QSize(self->minimumSize());
}

QSize* QLayoutItem_MaximumSize(const QLayoutItem* self) {
	return new QSize(self->maximumSize());
}

int QLayoutItem_ExpandingDirections(const QLayoutItem* self) {
	Qt::Orientations _ret = self->expandingDirections();
	return static_cast<int>(_ret);
}

void QLayoutItem_SetGeometry(QLayoutItem* self, QRect* geometry) {
	self->setGeometry(*geometry);
}

QRect* QLayoutItem_Geometry(const QLayoutItem* self) {
	return new QRect(self->geometry());
}

bool QLayoutItem_IsEmpty(const QLayoutItem* self) {
	return self->isEmpty();
}

bool QLayoutItem_HasHeightForWidth(const QLayoutItem* self) {
	return self->hasHeightForWidth();
}

int QLayoutItem_HeightForWidth(const QLayoutItem* self, int param1) {
	return self->heightForWidth(static_cast<int>(param1));
}

int QLayoutItem_MinimumHeightForWidth(const QLayoutItem* self, int param1) {
	return self->minimumHeightForWidth(static_cast<int>(param1));
}

void QLayoutItem_Invalidate(QLayoutItem* self) {
	self->invalidate();
}

QWidget* QLayoutItem_Widget(const QLayoutItem* self) {
	return self->widget();
}

QLayout* QLayoutItem_Layout(QLayoutItem* self) {
	return self->layout();
}

QSpacerItem* QLayoutItem_SpacerItem(QLayoutItem* self) {
	return self->spacerItem();
}

int QLayoutItem_Alignment(const QLayoutItem* self) {
	Qt::Alignment _ret = self->alignment();
	return static_cast<int>(_ret);
}

void QLayoutItem_SetAlignment(QLayoutItem* self, int a) {
	self->setAlignment(static_cast<Qt::Alignment>(a));
}

int QLayoutItem_ControlTypes(const QLayoutItem* self) {
	QSizePolicy::ControlTypes _ret = self->controlTypes();
	return static_cast<int>(_ret);
}

void QLayoutItem_Delete(QLayoutItem* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QLayoutItem*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQSpacerItem : public virtual QSpacerItem {
public:

	MiqtVirtualQSpacerItem(int w, int h): QSpacerItem(w, h) {};
	MiqtVirtualQSpacerItem(const QSpacerItem& param1): QSpacerItem(param1) {};
	MiqtVirtualQSpacerItem(int w, int h, QSizePolicy::Policy hData): QSpacerItem(w, h, hData) {};
	MiqtVirtualQSpacerItem(int w, int h, QSizePolicy::Policy hData, QSizePolicy::Policy vData): QSpacerItem(w, h, hData, vData) {};

	virtual ~MiqtVirtualQSpacerItem() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QSpacerItem::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QSpacerItem_SizeHint(const_cast<MiqtVirtualQSpacerItem*>(this), handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QSpacerItem::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSize = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSize() const override {
		if (handle__MinimumSize == 0) {
			return QSpacerItem::minimumSize();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QSpacerItem_MinimumSize(const_cast<MiqtVirtualQSpacerItem*>(this), handle__MinimumSize);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSize() const {

		return new QSize(QSpacerItem::minimumSize());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MaximumSize = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize maximumSize() const override {
		if (handle__MaximumSize == 0) {
			return QSpacerItem::maximumSize();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QSpacerItem_MaximumSize(const_cast<MiqtVirtualQSpacerItem*>(this), handle__MaximumSize);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MaximumSize() const {

		return new QSize(QSpacerItem::maximumSize());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ExpandingDirections = 0;

	// Subclass to allow providing a Go implementation
	virtual Qt::Orientations expandingDirections() const override {
		if (handle__ExpandingDirections == 0) {
			return QSpacerItem::expandingDirections();
		}
		

		int callback_return_value = miqt_exec_callback_QSpacerItem_ExpandingDirections(const_cast<MiqtVirtualQSpacerItem*>(this), handle__ExpandingDirections);

		return static_cast<Qt::Orientations>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_ExpandingDirections() const {

		Qt::Orientations _ret = QSpacerItem::expandingDirections();
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IsEmpty = 0;

	// Subclass to allow providing a Go implementation
	virtual bool isEmpty() const override {
		if (handle__IsEmpty == 0) {
			return QSpacerItem::isEmpty();
		}
		

		bool callback_return_value = miqt_exec_callback_QSpacerItem_IsEmpty(const_cast<MiqtVirtualQSpacerItem*>(this), handle__IsEmpty);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_IsEmpty() const {

		return QSpacerItem::isEmpty();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetGeometry = 0;

	// Subclass to allow providing a Go implementation
	virtual void setGeometry(const QRect& geometry) override {
		if (handle__SetGeometry == 0) {
			QSpacerItem::setGeometry(geometry);
			return;
		}
		
		const QRect& geometry_ret = geometry;
		// Cast returned reference into pointer
		QRect* sigval1 = const_cast<QRect*>(&geometry_ret);

		miqt_exec_callback_QSpacerItem_SetGeometry(this, handle__SetGeometry, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetGeometry(QRect* geometry) {

		QSpacerItem::setGeometry(*geometry);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Geometry = 0;

	// Subclass to allow providing a Go implementation
	virtual QRect geometry() const override {
		if (handle__Geometry == 0) {
			return QSpacerItem::geometry();
		}
		

		QRect* callback_return_value = miqt_exec_callback_QSpacerItem_Geometry(const_cast<MiqtVirtualQSpacerItem*>(this), handle__Geometry);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRect* virtualbase_Geometry() const {

		return new QRect(QSpacerItem::geometry());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SpacerItem = 0;

	// Subclass to allow providing a Go implementation
	virtual QSpacerItem* spacerItem() override {
		if (handle__SpacerItem == 0) {
			return QSpacerItem::spacerItem();
		}
		

		QSpacerItem* callback_return_value = miqt_exec_callback_QSpacerItem_SpacerItem(this, handle__SpacerItem);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSpacerItem* virtualbase_SpacerItem() {

		return QSpacerItem::spacerItem();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__HasHeightForWidth == 0) {
			return QSpacerItem::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QSpacerItem_HasHeightForWidth(const_cast<MiqtVirtualQSpacerItem*>(this), handle__HasHeightForWidth);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasHeightForWidth() const {

		return QSpacerItem::hasHeightForWidth();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int param1) const override {
		if (handle__HeightForWidth == 0) {
			return QSpacerItem::heightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QSpacerItem_HeightForWidth(const_cast<MiqtVirtualQSpacerItem*>(this), handle__HeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_HeightForWidth(int param1) const {

		return QSpacerItem::heightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int minimumHeightForWidth(int param1) const override {
		if (handle__MinimumHeightForWidth == 0) {
			return QSpacerItem::minimumHeightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QSpacerItem_MinimumHeightForWidth(const_cast<MiqtVirtualQSpacerItem*>(this), handle__MinimumHeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_MinimumHeightForWidth(int param1) const {

		return QSpacerItem::minimumHeightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Invalidate = 0;

	// Subclass to allow providing a Go implementation
	virtual void invalidate() override {
		if (handle__Invalidate == 0) {
			QSpacerItem::invalidate();
			return;
		}
		

		miqt_exec_callback_QSpacerItem_Invalidate(this, handle__Invalidate);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Invalidate() {

		QSpacerItem::invalidate();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Widget = 0;

	// Subclass to allow providing a Go implementation
	virtual QWidget* widget() const override {
		if (handle__Widget == 0) {
			return QSpacerItem::widget();
		}
		

		QWidget* callback_return_value = miqt_exec_callback_QSpacerItem_Widget(const_cast<MiqtVirtualQSpacerItem*>(this), handle__Widget);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QWidget* virtualbase_Widget() const {

		return QSpacerItem::widget();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Layout = 0;

	// Subclass to allow providing a Go implementation
	virtual QLayout* layout() override {
		if (handle__Layout == 0) {
			return QSpacerItem::layout();
		}
		

		QLayout* callback_return_value = miqt_exec_callback_QSpacerItem_Layout(this, handle__Layout);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QLayout* virtualbase_Layout() {

		return QSpacerItem::layout();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ControlTypes = 0;

	// Subclass to allow providing a Go implementation
	virtual QSizePolicy::ControlTypes controlTypes() const override {
		if (handle__ControlTypes == 0) {
			return QSpacerItem::controlTypes();
		}
		

		int callback_return_value = miqt_exec_callback_QSpacerItem_ControlTypes(const_cast<MiqtVirtualQSpacerItem*>(this), handle__ControlTypes);

		return static_cast<QSizePolicy::ControlTypes>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_ControlTypes() const {

		QSizePolicy::ControlTypes _ret = QSpacerItem::controlTypes();
		return static_cast<int>(_ret);

	}

};

void QSpacerItem_new(int w, int h, QSpacerItem** outptr_QSpacerItem, QLayoutItem** outptr_QLayoutItem) {
	MiqtVirtualQSpacerItem* ret = new MiqtVirtualQSpacerItem(static_cast<int>(w), static_cast<int>(h));
	*outptr_QSpacerItem = ret;
	*outptr_QLayoutItem = static_cast<QLayoutItem*>(ret);
}

void QSpacerItem_new2(QSpacerItem* param1, QSpacerItem** outptr_QSpacerItem, QLayoutItem** outptr_QLayoutItem) {
	MiqtVirtualQSpacerItem* ret = new MiqtVirtualQSpacerItem(*param1);
	*outptr_QSpacerItem = ret;
	*outptr_QLayoutItem = static_cast<QLayoutItem*>(ret);
}

void QSpacerItem_new3(int w, int h, int hData, QSpacerItem** outptr_QSpacerItem, QLayoutItem** outptr_QLayoutItem) {
	MiqtVirtualQSpacerItem* ret = new MiqtVirtualQSpacerItem(static_cast<int>(w), static_cast<int>(h), static_cast<QSizePolicy::Policy>(hData));
	*outptr_QSpacerItem = ret;
	*outptr_QLayoutItem = static_cast<QLayoutItem*>(ret);
}

void QSpacerItem_new4(int w, int h, int hData, int vData, QSpacerItem** outptr_QSpacerItem, QLayoutItem** outptr_QLayoutItem) {
	MiqtVirtualQSpacerItem* ret = new MiqtVirtualQSpacerItem(static_cast<int>(w), static_cast<int>(h), static_cast<QSizePolicy::Policy>(hData), static_cast<QSizePolicy::Policy>(vData));
	*outptr_QSpacerItem = ret;
	*outptr_QLayoutItem = static_cast<QLayoutItem*>(ret);
}

void QSpacerItem_ChangeSize(QSpacerItem* self, int w, int h) {
	self->changeSize(static_cast<int>(w), static_cast<int>(h));
}

QSize* QSpacerItem_SizeHint(const QSpacerItem* self) {
	return new QSize(self->sizeHint());
}

QSize* QSpacerItem_MinimumSize(const QSpacerItem* self) {
	return new QSize(self->minimumSize());
}

QSize* QSpacerItem_MaximumSize(const QSpacerItem* self) {
	return new QSize(self->maximumSize());
}

int QSpacerItem_ExpandingDirections(const QSpacerItem* self) {
	Qt::Orientations _ret = self->expandingDirections();
	return static_cast<int>(_ret);
}

bool QSpacerItem_IsEmpty(const QSpacerItem* self) {
	return self->isEmpty();
}

void QSpacerItem_SetGeometry(QSpacerItem* self, QRect* geometry) {
	self->setGeometry(*geometry);
}

QRect* QSpacerItem_Geometry(const QSpacerItem* self) {
	return new QRect(self->geometry());
}

QSpacerItem* QSpacerItem_SpacerItem(QSpacerItem* self) {
	return self->spacerItem();
}

QSizePolicy* QSpacerItem_SizePolicy(const QSpacerItem* self) {
	return new QSizePolicy(self->sizePolicy());
}

void QSpacerItem_ChangeSize3(QSpacerItem* self, int w, int h, int hData) {
	self->changeSize(static_cast<int>(w), static_cast<int>(h), static_cast<QSizePolicy::Policy>(hData));
}

void QSpacerItem_ChangeSize4(QSpacerItem* self, int w, int h, int hData, int vData) {
	self->changeSize(static_cast<int>(w), static_cast<int>(h), static_cast<QSizePolicy::Policy>(hData), static_cast<QSizePolicy::Policy>(vData));
}

void QSpacerItem_override_virtual_SizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__SizeHint = slot;
}

QSize* QSpacerItem_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_SizeHint();
}

void QSpacerItem_override_virtual_MinimumSize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__MinimumSize = slot;
}

QSize* QSpacerItem_virtualbase_MinimumSize(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_MinimumSize();
}

void QSpacerItem_override_virtual_MaximumSize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__MaximumSize = slot;
}

QSize* QSpacerItem_virtualbase_MaximumSize(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_MaximumSize();
}

void QSpacerItem_override_virtual_ExpandingDirections(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__ExpandingDirections = slot;
}

int QSpacerItem_virtualbase_ExpandingDirections(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_ExpandingDirections();
}

void QSpacerItem_override_virtual_IsEmpty(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__IsEmpty = slot;
}

bool QSpacerItem_virtualbase_IsEmpty(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_IsEmpty();
}

void QSpacerItem_override_virtual_SetGeometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__SetGeometry = slot;
}

void QSpacerItem_virtualbase_SetGeometry(void* self, QRect* geometry) {
	( (MiqtVirtualQSpacerItem*)(self) )->virtualbase_SetGeometry(geometry);
}

void QSpacerItem_override_virtual_Geometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__Geometry = slot;
}

QRect* QSpacerItem_virtualbase_Geometry(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_Geometry();
}

void QSpacerItem_override_virtual_SpacerItem(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__SpacerItem = slot;
}

QSpacerItem* QSpacerItem_virtualbase_SpacerItem(void* self) {
	return ( (MiqtVirtualQSpacerItem*)(self) )->virtualbase_SpacerItem();
}

void QSpacerItem_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__HasHeightForWidth = slot;
}

bool QSpacerItem_virtualbase_HasHeightForWidth(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_HasHeightForWidth();
}

void QSpacerItem_override_virtual_HeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__HeightForWidth = slot;
}

int QSpacerItem_virtualbase_HeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_HeightForWidth(param1);
}

void QSpacerItem_override_virtual_MinimumHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__MinimumHeightForWidth = slot;
}

int QSpacerItem_virtualbase_MinimumHeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_MinimumHeightForWidth(param1);
}

void QSpacerItem_override_virtual_Invalidate(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__Invalidate = slot;
}

void QSpacerItem_virtualbase_Invalidate(void* self) {
	( (MiqtVirtualQSpacerItem*)(self) )->virtualbase_Invalidate();
}

void QSpacerItem_override_virtual_Widget(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__Widget = slot;
}

QWidget* QSpacerItem_virtualbase_Widget(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_Widget();
}

void QSpacerItem_override_virtual_Layout(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__Layout = slot;
}

QLayout* QSpacerItem_virtualbase_Layout(void* self) {
	return ( (MiqtVirtualQSpacerItem*)(self) )->virtualbase_Layout();
}

void QSpacerItem_override_virtual_ControlTypes(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSpacerItem*>( (QSpacerItem*)(self) )->handle__ControlTypes = slot;
}

int QSpacerItem_virtualbase_ControlTypes(const void* self) {
	return ( (const MiqtVirtualQSpacerItem*)(self) )->virtualbase_ControlTypes();
}

void QSpacerItem_Delete(QSpacerItem* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSpacerItem*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQWidgetItem : public virtual QWidgetItem {
public:

	MiqtVirtualQWidgetItem(QWidget* w): QWidgetItem(w) {};

	virtual ~MiqtVirtualQWidgetItem() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QWidgetItem::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QWidgetItem_SizeHint(const_cast<MiqtVirtualQWidgetItem*>(this), handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QWidgetItem::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSize = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSize() const override {
		if (handle__MinimumSize == 0) {
			return QWidgetItem::minimumSize();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QWidgetItem_MinimumSize(const_cast<MiqtVirtualQWidgetItem*>(this), handle__MinimumSize);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSize() const {

		return new QSize(QWidgetItem::minimumSize());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MaximumSize = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize maximumSize() const override {
		if (handle__MaximumSize == 0) {
			return QWidgetItem::maximumSize();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QWidgetItem_MaximumSize(const_cast<MiqtVirtualQWidgetItem*>(this), handle__MaximumSize);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MaximumSize() const {

		return new QSize(QWidgetItem::maximumSize());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ExpandingDirections = 0;

	// Subclass to allow providing a Go implementation
	virtual Qt::Orientations expandingDirections() const override {
		if (handle__ExpandingDirections == 0) {
			return QWidgetItem::expandingDirections();
		}
		

		int callback_return_value = miqt_exec_callback_QWidgetItem_ExpandingDirections(const_cast<MiqtVirtualQWidgetItem*>(this), handle__ExpandingDirections);

		return static_cast<Qt::Orientations>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_ExpandingDirections() const {

		Qt::Orientations _ret = QWidgetItem::expandingDirections();
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IsEmpty = 0;

	// Subclass to allow providing a Go implementation
	virtual bool isEmpty() const override {
		if (handle__IsEmpty == 0) {
			return QWidgetItem::isEmpty();
		}
		

		bool callback_return_value = miqt_exec_callback_QWidgetItem_IsEmpty(const_cast<MiqtVirtualQWidgetItem*>(this), handle__IsEmpty);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_IsEmpty() const {

		return QWidgetItem::isEmpty();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetGeometry = 0;

	// Subclass to allow providing a Go implementation
	virtual void setGeometry(const QRect& geometry) override {
		if (handle__SetGeometry == 0) {
			QWidgetItem::setGeometry(geometry);
			return;
		}
		
		const QRect& geometry_ret = geometry;
		// Cast returned reference into pointer
		QRect* sigval1 = const_cast<QRect*>(&geometry_ret);

		miqt_exec_callback_QWidgetItem_SetGeometry(this, handle__SetGeometry, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetGeometry(QRect* geometry) {

		QWidgetItem::setGeometry(*geometry);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Geometry = 0;

	// Subclass to allow providing a Go implementation
	virtual QRect geometry() const override {
		if (handle__Geometry == 0) {
			return QWidgetItem::geometry();
		}
		

		QRect* callback_return_value = miqt_exec_callback_QWidgetItem_Geometry(const_cast<MiqtVirtualQWidgetItem*>(this), handle__Geometry);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRect* virtualbase_Geometry() const {

		return new QRect(QWidgetItem::geometry());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Widget = 0;

	// Subclass to allow providing a Go implementation
	virtual QWidget* widget() const override {
		if (handle__Widget == 0) {
			return QWidgetItem::widget();
		}
		

		QWidget* callback_return_value = miqt_exec_callback_QWidgetItem_Widget(const_cast<MiqtVirtualQWidgetItem*>(this), handle__Widget);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QWidget* virtualbase_Widget() const {

		return QWidgetItem::widget();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__HasHeightForWidth == 0) {
			return QWidgetItem::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QWidgetItem_HasHeightForWidth(const_cast<MiqtVirtualQWidgetItem*>(this), handle__HasHeightForWidth);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasHeightForWidth() const {

		return QWidgetItem::hasHeightForWidth();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int param1) const override {
		if (handle__HeightForWidth == 0) {
			return QWidgetItem::heightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QWidgetItem_HeightForWidth(const_cast<MiqtVirtualQWidgetItem*>(this), handle__HeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_HeightForWidth(int param1) const {

		return QWidgetItem::heightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int minimumHeightForWidth(int param1) const override {
		if (handle__MinimumHeightForWidth == 0) {
			return QWidgetItem::minimumHeightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QWidgetItem_MinimumHeightForWidth(const_cast<MiqtVirtualQWidgetItem*>(this), handle__MinimumHeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_MinimumHeightForWidth(int param1) const {

		return QWidgetItem::minimumHeightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ControlTypes = 0;

	// Subclass to allow providing a Go implementation
	virtual QSizePolicy::ControlTypes controlTypes() const override {
		if (handle__ControlTypes == 0) {
			return QWidgetItem::controlTypes();
		}
		

		int callback_return_value = miqt_exec_callback_QWidgetItem_ControlTypes(const_cast<MiqtVirtualQWidgetItem*>(this), handle__ControlTypes);

		return static_cast<QSizePolicy::ControlTypes>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_ControlTypes() const {

		QSizePolicy::ControlTypes _ret = QWidgetItem::controlTypes();
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Invalidate = 0;

	// Subclass to allow providing a Go implementation
	virtual void invalidate() override {
		if (handle__Invalidate == 0) {
			QWidgetItem::invalidate();
			return;
		}
		

		miqt_exec_callback_QWidgetItem_Invalidate(this, handle__Invalidate);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Invalidate() {

		QWidgetItem::invalidate();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Layout = 0;

	// Subclass to allow providing a Go implementation
	virtual QLayout* layout() override {
		if (handle__Layout == 0) {
			return QWidgetItem::layout();
		}
		

		QLayout* callback_return_value = miqt_exec_callback_QWidgetItem_Layout(this, handle__Layout);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QLayout* virtualbase_Layout() {

		return QWidgetItem::layout();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SpacerItem = 0;

	// Subclass to allow providing a Go implementation
	virtual QSpacerItem* spacerItem() override {
		if (handle__SpacerItem == 0) {
			return QWidgetItem::spacerItem();
		}
		

		QSpacerItem* callback_return_value = miqt_exec_callback_QWidgetItem_SpacerItem(this, handle__SpacerItem);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSpacerItem* virtualbase_SpacerItem() {

		return QWidgetItem::spacerItem();

	}

};

void QWidgetItem_new(QWidget* w, QWidgetItem** outptr_QWidgetItem, QLayoutItem** outptr_QLayoutItem) {
	MiqtVirtualQWidgetItem* ret = new MiqtVirtualQWidgetItem(w);
	*outptr_QWidgetItem = ret;
	*outptr_QLayoutItem = static_cast<QLayoutItem*>(ret);
}

QSize* QWidgetItem_SizeHint(const QWidgetItem* self) {
	return new QSize(self->sizeHint());
}

QSize* QWidgetItem_MinimumSize(const QWidgetItem* self) {
	return new QSize(self->minimumSize());
}

QSize* QWidgetItem_MaximumSize(const QWidgetItem* self) {
	return new QSize(self->maximumSize());
}

int QWidgetItem_ExpandingDirections(const QWidgetItem* self) {
	Qt::Orientations _ret = self->expandingDirections();
	return static_cast<int>(_ret);
}

bool QWidgetItem_IsEmpty(const QWidgetItem* self) {
	return self->isEmpty();
}

void QWidgetItem_SetGeometry(QWidgetItem* self, QRect* geometry) {
	self->setGeometry(*geometry);
}

QRect* QWidgetItem_Geometry(const QWidgetItem* self) {
	return new QRect(self->geometry());
}

QWidget* QWidgetItem_Widget(const QWidgetItem* self) {
	return self->widget();
}

bool QWidgetItem_HasHeightForWidth(const QWidgetItem* self) {
	return self->hasHeightForWidth();
}

int QWidgetItem_HeightForWidth(const QWidgetItem* self, int param1) {
	return self->heightForWidth(static_cast<int>(param1));
}

int QWidgetItem_MinimumHeightForWidth(const QWidgetItem* self, int param1) {
	return self->minimumHeightForWidth(static_cast<int>(param1));
}

int QWidgetItem_ControlTypes(const QWidgetItem* self) {
	QSizePolicy::ControlTypes _ret = self->controlTypes();
	return static_cast<int>(_ret);
}

void QWidgetItem_override_virtual_SizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__SizeHint = slot;
}

QSize* QWidgetItem_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_SizeHint();
}

void QWidgetItem_override_virtual_MinimumSize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__MinimumSize = slot;
}

QSize* QWidgetItem_virtualbase_MinimumSize(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_MinimumSize();
}

void QWidgetItem_override_virtual_MaximumSize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__MaximumSize = slot;
}

QSize* QWidgetItem_virtualbase_MaximumSize(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_MaximumSize();
}

void QWidgetItem_override_virtual_ExpandingDirections(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__ExpandingDirections = slot;
}

int QWidgetItem_virtualbase_ExpandingDirections(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_ExpandingDirections();
}

void QWidgetItem_override_virtual_IsEmpty(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__IsEmpty = slot;
}

bool QWidgetItem_virtualbase_IsEmpty(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_IsEmpty();
}

void QWidgetItem_override_virtual_SetGeometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__SetGeometry = slot;
}

void QWidgetItem_virtualbase_SetGeometry(void* self, QRect* geometry) {
	( (MiqtVirtualQWidgetItem*)(self) )->virtualbase_SetGeometry(geometry);
}

void QWidgetItem_override_virtual_Geometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__Geometry = slot;
}

QRect* QWidgetItem_virtualbase_Geometry(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_Geometry();
}

void QWidgetItem_override_virtual_Widget(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__Widget = slot;
}

QWidget* QWidgetItem_virtualbase_Widget(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_Widget();
}

void QWidgetItem_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__HasHeightForWidth = slot;
}

bool QWidgetItem_virtualbase_HasHeightForWidth(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_HasHeightForWidth();
}

void QWidgetItem_override_virtual_HeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__HeightForWidth = slot;
}

int QWidgetItem_virtualbase_HeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_HeightForWidth(param1);
}

void QWidgetItem_override_virtual_MinimumHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__MinimumHeightForWidth = slot;
}

int QWidgetItem_virtualbase_MinimumHeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_MinimumHeightForWidth(param1);
}

void QWidgetItem_override_virtual_ControlTypes(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__ControlTypes = slot;
}

int QWidgetItem_virtualbase_ControlTypes(const void* self) {
	return ( (const MiqtVirtualQWidgetItem*)(self) )->virtualbase_ControlTypes();
}

void QWidgetItem_override_virtual_Invalidate(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__Invalidate = slot;
}

void QWidgetItem_virtualbase_Invalidate(void* self) {
	( (MiqtVirtualQWidgetItem*)(self) )->virtualbase_Invalidate();
}

void QWidgetItem_override_virtual_Layout(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__Layout = slot;
}

QLayout* QWidgetItem_virtualbase_Layout(void* self) {
	return ( (MiqtVirtualQWidgetItem*)(self) )->virtualbase_Layout();
}

void QWidgetItem_override_virtual_SpacerItem(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItem*>( (QWidgetItem*)(self) )->handle__SpacerItem = slot;
}

QSpacerItem* QWidgetItem_virtualbase_SpacerItem(void* self) {
	return ( (MiqtVirtualQWidgetItem*)(self) )->virtualbase_SpacerItem();
}

void QWidgetItem_Delete(QWidgetItem* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWidgetItem*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQWidgetItemV2 : public virtual QWidgetItemV2 {
public:

	MiqtVirtualQWidgetItemV2(QWidget* widget): QWidgetItemV2(widget) {};

	virtual ~MiqtVirtualQWidgetItemV2() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QWidgetItemV2::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QWidgetItemV2_SizeHint(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QWidgetItemV2::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSize = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSize() const override {
		if (handle__MinimumSize == 0) {
			return QWidgetItemV2::minimumSize();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QWidgetItemV2_MinimumSize(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__MinimumSize);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSize() const {

		return new QSize(QWidgetItemV2::minimumSize());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MaximumSize = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize maximumSize() const override {
		if (handle__MaximumSize == 0) {
			return QWidgetItemV2::maximumSize();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QWidgetItemV2_MaximumSize(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__MaximumSize);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MaximumSize() const {

		return new QSize(QWidgetItemV2::maximumSize());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int width) const override {
		if (handle__HeightForWidth == 0) {
			return QWidgetItemV2::heightForWidth(width);
		}
		
		int sigval1 = width;

		int callback_return_value = miqt_exec_callback_QWidgetItemV2_HeightForWidth(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__HeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_HeightForWidth(int width) const {

		return QWidgetItemV2::heightForWidth(static_cast<int>(width));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ExpandingDirections = 0;

	// Subclass to allow providing a Go implementation
	virtual Qt::Orientations expandingDirections() const override {
		if (handle__ExpandingDirections == 0) {
			return QWidgetItemV2::expandingDirections();
		}
		

		int callback_return_value = miqt_exec_callback_QWidgetItemV2_ExpandingDirections(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__ExpandingDirections);

		return static_cast<Qt::Orientations>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_ExpandingDirections() const {

		Qt::Orientations _ret = QWidgetItemV2::expandingDirections();
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IsEmpty = 0;

	// Subclass to allow providing a Go implementation
	virtual bool isEmpty() const override {
		if (handle__IsEmpty == 0) {
			return QWidgetItemV2::isEmpty();
		}
		

		bool callback_return_value = miqt_exec_callback_QWidgetItemV2_IsEmpty(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__IsEmpty);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_IsEmpty() const {

		return QWidgetItemV2::isEmpty();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetGeometry = 0;

	// Subclass to allow providing a Go implementation
	virtual void setGeometry(const QRect& geometry) override {
		if (handle__SetGeometry == 0) {
			QWidgetItemV2::setGeometry(geometry);
			return;
		}
		
		const QRect& geometry_ret = geometry;
		// Cast returned reference into pointer
		QRect* sigval1 = const_cast<QRect*>(&geometry_ret);

		miqt_exec_callback_QWidgetItemV2_SetGeometry(this, handle__SetGeometry, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetGeometry(QRect* geometry) {

		QWidgetItemV2::setGeometry(*geometry);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Geometry = 0;

	// Subclass to allow providing a Go implementation
	virtual QRect geometry() const override {
		if (handle__Geometry == 0) {
			return QWidgetItemV2::geometry();
		}
		

		QRect* callback_return_value = miqt_exec_callback_QWidgetItemV2_Geometry(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__Geometry);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRect* virtualbase_Geometry() const {

		return new QRect(QWidgetItemV2::geometry());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Widget = 0;

	// Subclass to allow providing a Go implementation
	virtual QWidget* widget() const override {
		if (handle__Widget == 0) {
			return QWidgetItemV2::widget();
		}
		

		QWidget* callback_return_value = miqt_exec_callback_QWidgetItemV2_Widget(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__Widget);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QWidget* virtualbase_Widget() const {

		return QWidgetItemV2::widget();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__HasHeightForWidth == 0) {
			return QWidgetItemV2::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QWidgetItemV2_HasHeightForWidth(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__HasHeightForWidth);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasHeightForWidth() const {

		return QWidgetItemV2::hasHeightForWidth();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int minimumHeightForWidth(int param1) const override {
		if (handle__MinimumHeightForWidth == 0) {
			return QWidgetItemV2::minimumHeightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QWidgetItemV2_MinimumHeightForWidth(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__MinimumHeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_MinimumHeightForWidth(int param1) const {

		return QWidgetItemV2::minimumHeightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ControlTypes = 0;

	// Subclass to allow providing a Go implementation
	virtual QSizePolicy::ControlTypes controlTypes() const override {
		if (handle__ControlTypes == 0) {
			return QWidgetItemV2::controlTypes();
		}
		

		int callback_return_value = miqt_exec_callback_QWidgetItemV2_ControlTypes(const_cast<MiqtVirtualQWidgetItemV2*>(this), handle__ControlTypes);

		return static_cast<QSizePolicy::ControlTypes>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_ControlTypes() const {

		QSizePolicy::ControlTypes _ret = QWidgetItemV2::controlTypes();
		return static_cast<int>(_ret);

	}

};

void QWidgetItemV2_new(QWidget* widget, QWidgetItemV2** outptr_QWidgetItemV2, QWidgetItem** outptr_QWidgetItem, QLayoutItem** outptr_QLayoutItem) {
	MiqtVirtualQWidgetItemV2* ret = new MiqtVirtualQWidgetItemV2(widget);
	*outptr_QWidgetItemV2 = ret;
	*outptr_QWidgetItem = static_cast<QWidgetItem*>(ret);
	*outptr_QLayoutItem = static_cast<QLayoutItem*>(ret);
}

QSize* QWidgetItemV2_SizeHint(const QWidgetItemV2* self) {
	return new QSize(self->sizeHint());
}

QSize* QWidgetItemV2_MinimumSize(const QWidgetItemV2* self) {
	return new QSize(self->minimumSize());
}

QSize* QWidgetItemV2_MaximumSize(const QWidgetItemV2* self) {
	return new QSize(self->maximumSize());
}

int QWidgetItemV2_HeightForWidth(const QWidgetItemV2* self, int width) {
	return self->heightForWidth(static_cast<int>(width));
}

void QWidgetItemV2_override_virtual_SizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__SizeHint = slot;
}

QSize* QWidgetItemV2_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_SizeHint();
}

void QWidgetItemV2_override_virtual_MinimumSize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__MinimumSize = slot;
}

QSize* QWidgetItemV2_virtualbase_MinimumSize(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_MinimumSize();
}

void QWidgetItemV2_override_virtual_MaximumSize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__MaximumSize = slot;
}

QSize* QWidgetItemV2_virtualbase_MaximumSize(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_MaximumSize();
}

void QWidgetItemV2_override_virtual_HeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__HeightForWidth = slot;
}

int QWidgetItemV2_virtualbase_HeightForWidth(const void* self, int width) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_HeightForWidth(width);
}

void QWidgetItemV2_override_virtual_ExpandingDirections(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__ExpandingDirections = slot;
}

int QWidgetItemV2_virtualbase_ExpandingDirections(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_ExpandingDirections();
}

void QWidgetItemV2_override_virtual_IsEmpty(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__IsEmpty = slot;
}

bool QWidgetItemV2_virtualbase_IsEmpty(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_IsEmpty();
}

void QWidgetItemV2_override_virtual_SetGeometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__SetGeometry = slot;
}

void QWidgetItemV2_virtualbase_SetGeometry(void* self, QRect* geometry) {
	( (MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_SetGeometry(geometry);
}

void QWidgetItemV2_override_virtual_Geometry(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__Geometry = slot;
}

QRect* QWidgetItemV2_virtualbase_Geometry(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_Geometry();
}

void QWidgetItemV2_override_virtual_Widget(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__Widget = slot;
}

QWidget* QWidgetItemV2_virtualbase_Widget(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_Widget();
}

void QWidgetItemV2_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__HasHeightForWidth = slot;
}

bool QWidgetItemV2_virtualbase_HasHeightForWidth(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_HasHeightForWidth();
}

void QWidgetItemV2_override_virtual_MinimumHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__MinimumHeightForWidth = slot;
}

int QWidgetItemV2_virtualbase_MinimumHeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_MinimumHeightForWidth(param1);
}

void QWidgetItemV2_override_virtual_ControlTypes(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetItemV2*>( (QWidgetItemV2*)(self) )->handle__ControlTypes = slot;
}

int QWidgetItemV2_virtualbase_ControlTypes(const void* self) {
	return ( (const MiqtVirtualQWidgetItemV2*)(self) )->virtualbase_ControlTypes();
}

void QWidgetItemV2_Delete(QWidgetItemV2* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWidgetItemV2*>( self );
	} else {
		delete self;
	}
}

