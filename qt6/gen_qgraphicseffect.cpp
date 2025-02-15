#include <QBrush>
#include <QColor>
#include <QGraphicsBlurEffect>
#include <QGraphicsColorizeEffect>
#include <QGraphicsDropShadowEffect>
#include <QGraphicsEffect>
#include <QGraphicsOpacityEffect>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QPointF>
#include <QRectF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qgraphicseffect.h>
#include "gen_qgraphicseffect.h"
#include "_cgo_export.h"

QMetaObject* QGraphicsEffect_MetaObject(const QGraphicsEffect* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGraphicsEffect_Metacast(QGraphicsEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGraphicsEffect_Tr(const char* s) {
	QString _ret = QGraphicsEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QRectF* QGraphicsEffect_BoundingRectFor(const QGraphicsEffect* self, QRectF* sourceRect) {
	return new QRectF(self->boundingRectFor(*sourceRect));
}

QRectF* QGraphicsEffect_BoundingRect(const QGraphicsEffect* self) {
	return new QRectF(self->boundingRect());
}

bool QGraphicsEffect_IsEnabled(const QGraphicsEffect* self) {
	return self->isEnabled();
}

void QGraphicsEffect_SetEnabled(QGraphicsEffect* self, bool enable) {
	self->setEnabled(enable);
}

void QGraphicsEffect_Update(QGraphicsEffect* self) {
	self->update();
}

void QGraphicsEffect_EnabledChanged(QGraphicsEffect* self, bool enabled) {
	self->enabledChanged(enabled);
}

void QGraphicsEffect_connect_EnabledChanged(QGraphicsEffect* self, intptr_t slot) {
	QGraphicsEffect::connect(self, static_cast<void (QGraphicsEffect::*)(bool)>(&QGraphicsEffect::enabledChanged), self, [=](bool enabled) {
		bool sigval1 = enabled;
		miqt_exec_callback_QGraphicsEffect_EnabledChanged(slot, sigval1);
	});
}

struct miqt_string QGraphicsEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsEffect_Delete(QGraphicsEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QGraphicsEffect*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQGraphicsColorizeEffect : public virtual QGraphicsColorizeEffect {
public:

	MiqtVirtualQGraphicsColorizeEffect(): QGraphicsColorizeEffect() {};
	MiqtVirtualQGraphicsColorizeEffect(QObject* parent): QGraphicsColorizeEffect(parent) {};

	virtual ~MiqtVirtualQGraphicsColorizeEffect() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Draw = 0;

	// Subclass to allow providing a Go implementation
	virtual void draw(QPainter* painter) override {
		if (handle__Draw == 0) {
			QGraphicsColorizeEffect::draw(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QGraphicsColorizeEffect_Draw(this, handle__Draw, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Draw(QPainter* painter) {

		QGraphicsColorizeEffect::draw(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BoundingRectFor = 0;

	// Subclass to allow providing a Go implementation
	virtual QRectF boundingRectFor(const QRectF& sourceRect) const override {
		if (handle__BoundingRectFor == 0) {
			return QGraphicsColorizeEffect::boundingRectFor(sourceRect);
		}
		
		const QRectF& sourceRect_ret = sourceRect;
		// Cast returned reference into pointer
		QRectF* sigval1 = const_cast<QRectF*>(&sourceRect_ret);

		QRectF* callback_return_value = miqt_exec_callback_QGraphicsColorizeEffect_BoundingRectFor(const_cast<MiqtVirtualQGraphicsColorizeEffect*>(this), handle__BoundingRectFor, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRectF* virtualbase_BoundingRectFor(QRectF* sourceRect) const {

		return new QRectF(QGraphicsColorizeEffect::boundingRectFor(*sourceRect));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SourceChanged = 0;

	// Subclass to allow providing a Go implementation
	virtual void sourceChanged(QGraphicsEffect::ChangeFlags flags) override {
		if (handle__SourceChanged == 0) {
			QGraphicsColorizeEffect::sourceChanged(flags);
			return;
		}
		
		QGraphicsEffect::ChangeFlags flags_ret = flags;
		int sigval1 = static_cast<int>(flags_ret);

		miqt_exec_callback_QGraphicsColorizeEffect_SourceChanged(this, handle__SourceChanged, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SourceChanged(int flags) {

		QGraphicsColorizeEffect::sourceChanged(static_cast<QGraphicsEffect::ChangeFlags>(flags));

	}

};

void QGraphicsColorizeEffect_new(QGraphicsColorizeEffect** outptr_QGraphicsColorizeEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsColorizeEffect* ret = new MiqtVirtualQGraphicsColorizeEffect();
	*outptr_QGraphicsColorizeEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QGraphicsColorizeEffect_new2(QObject* parent, QGraphicsColorizeEffect** outptr_QGraphicsColorizeEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsColorizeEffect* ret = new MiqtVirtualQGraphicsColorizeEffect(parent);
	*outptr_QGraphicsColorizeEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QGraphicsColorizeEffect_MetaObject(const QGraphicsColorizeEffect* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGraphicsColorizeEffect_Metacast(QGraphicsColorizeEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGraphicsColorizeEffect_Tr(const char* s) {
	QString _ret = QGraphicsColorizeEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QColor* QGraphicsColorizeEffect_Color(const QGraphicsColorizeEffect* self) {
	return new QColor(self->color());
}

double QGraphicsColorizeEffect_Strength(const QGraphicsColorizeEffect* self) {
	qreal _ret = self->strength();
	return static_cast<double>(_ret);
}

void QGraphicsColorizeEffect_SetColor(QGraphicsColorizeEffect* self, QColor* c) {
	self->setColor(*c);
}

void QGraphicsColorizeEffect_SetStrength(QGraphicsColorizeEffect* self, double strength) {
	self->setStrength(static_cast<qreal>(strength));
}

void QGraphicsColorizeEffect_ColorChanged(QGraphicsColorizeEffect* self, QColor* color) {
	self->colorChanged(*color);
}

void QGraphicsColorizeEffect_connect_ColorChanged(QGraphicsColorizeEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsColorizeEffect::connect(self, static_cast<void (QGraphicsColorizeEffect::*)(const QColor&)>(&QGraphicsColorizeEffect::colorChanged), self, [=](const QColor& color) {
		const QColor& color_ret = color;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&color_ret);
		miqt_exec_callback_QGraphicsColorizeEffect_ColorChanged(slot, sigval1);
	});
}

void QGraphicsColorizeEffect_StrengthChanged(QGraphicsColorizeEffect* self, double strength) {
	self->strengthChanged(static_cast<qreal>(strength));
}

void QGraphicsColorizeEffect_connect_StrengthChanged(QGraphicsColorizeEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsColorizeEffect::connect(self, static_cast<void (QGraphicsColorizeEffect::*)(qreal)>(&QGraphicsColorizeEffect::strengthChanged), self, [=](qreal strength) {
		qreal strength_ret = strength;
		double sigval1 = static_cast<double>(strength_ret);
		miqt_exec_callback_QGraphicsColorizeEffect_StrengthChanged(slot, sigval1);
	});
}

struct miqt_string QGraphicsColorizeEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsColorizeEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsColorizeEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsColorizeEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsColorizeEffect_override_virtual_Draw(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsColorizeEffect*>( (QGraphicsColorizeEffect*)(self) )->handle__Draw = slot;
}

void QGraphicsColorizeEffect_virtualbase_Draw(void* self, QPainter* painter) {
	( (MiqtVirtualQGraphicsColorizeEffect*)(self) )->virtualbase_Draw(painter);
}

void QGraphicsColorizeEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsColorizeEffect*>( (QGraphicsColorizeEffect*)(self) )->handle__BoundingRectFor = slot;
}

QRectF* QGraphicsColorizeEffect_virtualbase_BoundingRectFor(const void* self, QRectF* sourceRect) {
	return ( (const MiqtVirtualQGraphicsColorizeEffect*)(self) )->virtualbase_BoundingRectFor(sourceRect);
}

void QGraphicsColorizeEffect_override_virtual_SourceChanged(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsColorizeEffect*>( (QGraphicsColorizeEffect*)(self) )->handle__SourceChanged = slot;
}

void QGraphicsColorizeEffect_virtualbase_SourceChanged(void* self, int flags) {
	( (MiqtVirtualQGraphicsColorizeEffect*)(self) )->virtualbase_SourceChanged(flags);
}

void QGraphicsColorizeEffect_Delete(QGraphicsColorizeEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsColorizeEffect*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQGraphicsBlurEffect : public virtual QGraphicsBlurEffect {
public:

	MiqtVirtualQGraphicsBlurEffect(): QGraphicsBlurEffect() {};
	MiqtVirtualQGraphicsBlurEffect(QObject* parent): QGraphicsBlurEffect(parent) {};

	virtual ~MiqtVirtualQGraphicsBlurEffect() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BoundingRectFor = 0;

	// Subclass to allow providing a Go implementation
	virtual QRectF boundingRectFor(const QRectF& rect) const override {
		if (handle__BoundingRectFor == 0) {
			return QGraphicsBlurEffect::boundingRectFor(rect);
		}
		
		const QRectF& rect_ret = rect;
		// Cast returned reference into pointer
		QRectF* sigval1 = const_cast<QRectF*>(&rect_ret);

		QRectF* callback_return_value = miqt_exec_callback_QGraphicsBlurEffect_BoundingRectFor(const_cast<MiqtVirtualQGraphicsBlurEffect*>(this), handle__BoundingRectFor, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRectF* virtualbase_BoundingRectFor(QRectF* rect) const {

		return new QRectF(QGraphicsBlurEffect::boundingRectFor(*rect));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Draw = 0;

	// Subclass to allow providing a Go implementation
	virtual void draw(QPainter* painter) override {
		if (handle__Draw == 0) {
			QGraphicsBlurEffect::draw(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QGraphicsBlurEffect_Draw(this, handle__Draw, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Draw(QPainter* painter) {

		QGraphicsBlurEffect::draw(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SourceChanged = 0;

	// Subclass to allow providing a Go implementation
	virtual void sourceChanged(QGraphicsEffect::ChangeFlags flags) override {
		if (handle__SourceChanged == 0) {
			QGraphicsBlurEffect::sourceChanged(flags);
			return;
		}
		
		QGraphicsEffect::ChangeFlags flags_ret = flags;
		int sigval1 = static_cast<int>(flags_ret);

		miqt_exec_callback_QGraphicsBlurEffect_SourceChanged(this, handle__SourceChanged, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SourceChanged(int flags) {

		QGraphicsBlurEffect::sourceChanged(static_cast<QGraphicsEffect::ChangeFlags>(flags));

	}

};

void QGraphicsBlurEffect_new(QGraphicsBlurEffect** outptr_QGraphicsBlurEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsBlurEffect* ret = new MiqtVirtualQGraphicsBlurEffect();
	*outptr_QGraphicsBlurEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QGraphicsBlurEffect_new2(QObject* parent, QGraphicsBlurEffect** outptr_QGraphicsBlurEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsBlurEffect* ret = new MiqtVirtualQGraphicsBlurEffect(parent);
	*outptr_QGraphicsBlurEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QGraphicsBlurEffect_MetaObject(const QGraphicsBlurEffect* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGraphicsBlurEffect_Metacast(QGraphicsBlurEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGraphicsBlurEffect_Tr(const char* s) {
	QString _ret = QGraphicsBlurEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QRectF* QGraphicsBlurEffect_BoundingRectFor(const QGraphicsBlurEffect* self, QRectF* rect) {
	return new QRectF(self->boundingRectFor(*rect));
}

double QGraphicsBlurEffect_BlurRadius(const QGraphicsBlurEffect* self) {
	qreal _ret = self->blurRadius();
	return static_cast<double>(_ret);
}

int QGraphicsBlurEffect_BlurHints(const QGraphicsBlurEffect* self) {
	QGraphicsBlurEffect::BlurHints _ret = self->blurHints();
	return static_cast<int>(_ret);
}

void QGraphicsBlurEffect_SetBlurRadius(QGraphicsBlurEffect* self, double blurRadius) {
	self->setBlurRadius(static_cast<qreal>(blurRadius));
}

void QGraphicsBlurEffect_SetBlurHints(QGraphicsBlurEffect* self, int hints) {
	self->setBlurHints(static_cast<QGraphicsBlurEffect::BlurHints>(hints));
}

void QGraphicsBlurEffect_BlurRadiusChanged(QGraphicsBlurEffect* self, double blurRadius) {
	self->blurRadiusChanged(static_cast<qreal>(blurRadius));
}

void QGraphicsBlurEffect_connect_BlurRadiusChanged(QGraphicsBlurEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsBlurEffect::connect(self, static_cast<void (QGraphicsBlurEffect::*)(qreal)>(&QGraphicsBlurEffect::blurRadiusChanged), self, [=](qreal blurRadius) {
		qreal blurRadius_ret = blurRadius;
		double sigval1 = static_cast<double>(blurRadius_ret);
		miqt_exec_callback_QGraphicsBlurEffect_BlurRadiusChanged(slot, sigval1);
	});
}

void QGraphicsBlurEffect_BlurHintsChanged(QGraphicsBlurEffect* self, int hints) {
	self->blurHintsChanged(static_cast<QGraphicsBlurEffect::BlurHints>(hints));
}

void QGraphicsBlurEffect_connect_BlurHintsChanged(QGraphicsBlurEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsBlurEffect::connect(self, static_cast<void (QGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), self, [=](QGraphicsBlurEffect::BlurHints hints) {
		QGraphicsBlurEffect::BlurHints hints_ret = hints;
		int sigval1 = static_cast<int>(hints_ret);
		miqt_exec_callback_QGraphicsBlurEffect_BlurHintsChanged(slot, sigval1);
	});
}

struct miqt_string QGraphicsBlurEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsBlurEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsBlurEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsBlurEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsBlurEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsBlurEffect*>( (QGraphicsBlurEffect*)(self) )->handle__BoundingRectFor = slot;
}

QRectF* QGraphicsBlurEffect_virtualbase_BoundingRectFor(const void* self, QRectF* rect) {
	return ( (const MiqtVirtualQGraphicsBlurEffect*)(self) )->virtualbase_BoundingRectFor(rect);
}

void QGraphicsBlurEffect_override_virtual_Draw(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsBlurEffect*>( (QGraphicsBlurEffect*)(self) )->handle__Draw = slot;
}

void QGraphicsBlurEffect_virtualbase_Draw(void* self, QPainter* painter) {
	( (MiqtVirtualQGraphicsBlurEffect*)(self) )->virtualbase_Draw(painter);
}

void QGraphicsBlurEffect_override_virtual_SourceChanged(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsBlurEffect*>( (QGraphicsBlurEffect*)(self) )->handle__SourceChanged = slot;
}

void QGraphicsBlurEffect_virtualbase_SourceChanged(void* self, int flags) {
	( (MiqtVirtualQGraphicsBlurEffect*)(self) )->virtualbase_SourceChanged(flags);
}

void QGraphicsBlurEffect_Delete(QGraphicsBlurEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsBlurEffect*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQGraphicsDropShadowEffect : public virtual QGraphicsDropShadowEffect {
public:

	MiqtVirtualQGraphicsDropShadowEffect(): QGraphicsDropShadowEffect() {};
	MiqtVirtualQGraphicsDropShadowEffect(QObject* parent): QGraphicsDropShadowEffect(parent) {};

	virtual ~MiqtVirtualQGraphicsDropShadowEffect() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BoundingRectFor = 0;

	// Subclass to allow providing a Go implementation
	virtual QRectF boundingRectFor(const QRectF& rect) const override {
		if (handle__BoundingRectFor == 0) {
			return QGraphicsDropShadowEffect::boundingRectFor(rect);
		}
		
		const QRectF& rect_ret = rect;
		// Cast returned reference into pointer
		QRectF* sigval1 = const_cast<QRectF*>(&rect_ret);

		QRectF* callback_return_value = miqt_exec_callback_QGraphicsDropShadowEffect_BoundingRectFor(const_cast<MiqtVirtualQGraphicsDropShadowEffect*>(this), handle__BoundingRectFor, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRectF* virtualbase_BoundingRectFor(QRectF* rect) const {

		return new QRectF(QGraphicsDropShadowEffect::boundingRectFor(*rect));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Draw = 0;

	// Subclass to allow providing a Go implementation
	virtual void draw(QPainter* painter) override {
		if (handle__Draw == 0) {
			QGraphicsDropShadowEffect::draw(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QGraphicsDropShadowEffect_Draw(this, handle__Draw, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Draw(QPainter* painter) {

		QGraphicsDropShadowEffect::draw(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SourceChanged = 0;

	// Subclass to allow providing a Go implementation
	virtual void sourceChanged(QGraphicsEffect::ChangeFlags flags) override {
		if (handle__SourceChanged == 0) {
			QGraphicsDropShadowEffect::sourceChanged(flags);
			return;
		}
		
		QGraphicsEffect::ChangeFlags flags_ret = flags;
		int sigval1 = static_cast<int>(flags_ret);

		miqt_exec_callback_QGraphicsDropShadowEffect_SourceChanged(this, handle__SourceChanged, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SourceChanged(int flags) {

		QGraphicsDropShadowEffect::sourceChanged(static_cast<QGraphicsEffect::ChangeFlags>(flags));

	}

};

void QGraphicsDropShadowEffect_new(QGraphicsDropShadowEffect** outptr_QGraphicsDropShadowEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsDropShadowEffect* ret = new MiqtVirtualQGraphicsDropShadowEffect();
	*outptr_QGraphicsDropShadowEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QGraphicsDropShadowEffect_new2(QObject* parent, QGraphicsDropShadowEffect** outptr_QGraphicsDropShadowEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsDropShadowEffect* ret = new MiqtVirtualQGraphicsDropShadowEffect(parent);
	*outptr_QGraphicsDropShadowEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QGraphicsDropShadowEffect_MetaObject(const QGraphicsDropShadowEffect* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGraphicsDropShadowEffect_Metacast(QGraphicsDropShadowEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGraphicsDropShadowEffect_Tr(const char* s) {
	QString _ret = QGraphicsDropShadowEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QRectF* QGraphicsDropShadowEffect_BoundingRectFor(const QGraphicsDropShadowEffect* self, QRectF* rect) {
	return new QRectF(self->boundingRectFor(*rect));
}

QPointF* QGraphicsDropShadowEffect_Offset(const QGraphicsDropShadowEffect* self) {
	return new QPointF(self->offset());
}

double QGraphicsDropShadowEffect_XOffset(const QGraphicsDropShadowEffect* self) {
	qreal _ret = self->xOffset();
	return static_cast<double>(_ret);
}

double QGraphicsDropShadowEffect_YOffset(const QGraphicsDropShadowEffect* self) {
	qreal _ret = self->yOffset();
	return static_cast<double>(_ret);
}

double QGraphicsDropShadowEffect_BlurRadius(const QGraphicsDropShadowEffect* self) {
	qreal _ret = self->blurRadius();
	return static_cast<double>(_ret);
}

QColor* QGraphicsDropShadowEffect_Color(const QGraphicsDropShadowEffect* self) {
	return new QColor(self->color());
}

void QGraphicsDropShadowEffect_SetOffset(QGraphicsDropShadowEffect* self, QPointF* ofs) {
	self->setOffset(*ofs);
}

void QGraphicsDropShadowEffect_SetOffset2(QGraphicsDropShadowEffect* self, double dx, double dy) {
	self->setOffset(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

void QGraphicsDropShadowEffect_SetOffsetWithQreal(QGraphicsDropShadowEffect* self, double d) {
	self->setOffset(static_cast<qreal>(d));
}

void QGraphicsDropShadowEffect_SetXOffset(QGraphicsDropShadowEffect* self, double dx) {
	self->setXOffset(static_cast<qreal>(dx));
}

void QGraphicsDropShadowEffect_SetYOffset(QGraphicsDropShadowEffect* self, double dy) {
	self->setYOffset(static_cast<qreal>(dy));
}

void QGraphicsDropShadowEffect_SetBlurRadius(QGraphicsDropShadowEffect* self, double blurRadius) {
	self->setBlurRadius(static_cast<qreal>(blurRadius));
}

void QGraphicsDropShadowEffect_SetColor(QGraphicsDropShadowEffect* self, QColor* color) {
	self->setColor(*color);
}

void QGraphicsDropShadowEffect_OffsetChanged(QGraphicsDropShadowEffect* self, QPointF* offset) {
	self->offsetChanged(*offset);
}

void QGraphicsDropShadowEffect_connect_OffsetChanged(QGraphicsDropShadowEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsDropShadowEffect::connect(self, static_cast<void (QGraphicsDropShadowEffect::*)(const QPointF&)>(&QGraphicsDropShadowEffect::offsetChanged), self, [=](const QPointF& offset) {
		const QPointF& offset_ret = offset;
		// Cast returned reference into pointer
		QPointF* sigval1 = const_cast<QPointF*>(&offset_ret);
		miqt_exec_callback_QGraphicsDropShadowEffect_OffsetChanged(slot, sigval1);
	});
}

void QGraphicsDropShadowEffect_BlurRadiusChanged(QGraphicsDropShadowEffect* self, double blurRadius) {
	self->blurRadiusChanged(static_cast<qreal>(blurRadius));
}

void QGraphicsDropShadowEffect_connect_BlurRadiusChanged(QGraphicsDropShadowEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsDropShadowEffect::connect(self, static_cast<void (QGraphicsDropShadowEffect::*)(qreal)>(&QGraphicsDropShadowEffect::blurRadiusChanged), self, [=](qreal blurRadius) {
		qreal blurRadius_ret = blurRadius;
		double sigval1 = static_cast<double>(blurRadius_ret);
		miqt_exec_callback_QGraphicsDropShadowEffect_BlurRadiusChanged(slot, sigval1);
	});
}

void QGraphicsDropShadowEffect_ColorChanged(QGraphicsDropShadowEffect* self, QColor* color) {
	self->colorChanged(*color);
}

void QGraphicsDropShadowEffect_connect_ColorChanged(QGraphicsDropShadowEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsDropShadowEffect::connect(self, static_cast<void (QGraphicsDropShadowEffect::*)(const QColor&)>(&QGraphicsDropShadowEffect::colorChanged), self, [=](const QColor& color) {
		const QColor& color_ret = color;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&color_ret);
		miqt_exec_callback_QGraphicsDropShadowEffect_ColorChanged(slot, sigval1);
	});
}

struct miqt_string QGraphicsDropShadowEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsDropShadowEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsDropShadowEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsDropShadowEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsDropShadowEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsDropShadowEffect*>( (QGraphicsDropShadowEffect*)(self) )->handle__BoundingRectFor = slot;
}

QRectF* QGraphicsDropShadowEffect_virtualbase_BoundingRectFor(const void* self, QRectF* rect) {
	return ( (const MiqtVirtualQGraphicsDropShadowEffect*)(self) )->virtualbase_BoundingRectFor(rect);
}

void QGraphicsDropShadowEffect_override_virtual_Draw(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsDropShadowEffect*>( (QGraphicsDropShadowEffect*)(self) )->handle__Draw = slot;
}

void QGraphicsDropShadowEffect_virtualbase_Draw(void* self, QPainter* painter) {
	( (MiqtVirtualQGraphicsDropShadowEffect*)(self) )->virtualbase_Draw(painter);
}

void QGraphicsDropShadowEffect_override_virtual_SourceChanged(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsDropShadowEffect*>( (QGraphicsDropShadowEffect*)(self) )->handle__SourceChanged = slot;
}

void QGraphicsDropShadowEffect_virtualbase_SourceChanged(void* self, int flags) {
	( (MiqtVirtualQGraphicsDropShadowEffect*)(self) )->virtualbase_SourceChanged(flags);
}

void QGraphicsDropShadowEffect_Delete(QGraphicsDropShadowEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsDropShadowEffect*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQGraphicsOpacityEffect : public virtual QGraphicsOpacityEffect {
public:

	MiqtVirtualQGraphicsOpacityEffect(): QGraphicsOpacityEffect() {};
	MiqtVirtualQGraphicsOpacityEffect(QObject* parent): QGraphicsOpacityEffect(parent) {};

	virtual ~MiqtVirtualQGraphicsOpacityEffect() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Draw = 0;

	// Subclass to allow providing a Go implementation
	virtual void draw(QPainter* painter) override {
		if (handle__Draw == 0) {
			QGraphicsOpacityEffect::draw(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QGraphicsOpacityEffect_Draw(this, handle__Draw, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Draw(QPainter* painter) {

		QGraphicsOpacityEffect::draw(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BoundingRectFor = 0;

	// Subclass to allow providing a Go implementation
	virtual QRectF boundingRectFor(const QRectF& sourceRect) const override {
		if (handle__BoundingRectFor == 0) {
			return QGraphicsOpacityEffect::boundingRectFor(sourceRect);
		}
		
		const QRectF& sourceRect_ret = sourceRect;
		// Cast returned reference into pointer
		QRectF* sigval1 = const_cast<QRectF*>(&sourceRect_ret);

		QRectF* callback_return_value = miqt_exec_callback_QGraphicsOpacityEffect_BoundingRectFor(const_cast<MiqtVirtualQGraphicsOpacityEffect*>(this), handle__BoundingRectFor, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRectF* virtualbase_BoundingRectFor(QRectF* sourceRect) const {

		return new QRectF(QGraphicsOpacityEffect::boundingRectFor(*sourceRect));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SourceChanged = 0;

	// Subclass to allow providing a Go implementation
	virtual void sourceChanged(QGraphicsEffect::ChangeFlags flags) override {
		if (handle__SourceChanged == 0) {
			QGraphicsOpacityEffect::sourceChanged(flags);
			return;
		}
		
		QGraphicsEffect::ChangeFlags flags_ret = flags;
		int sigval1 = static_cast<int>(flags_ret);

		miqt_exec_callback_QGraphicsOpacityEffect_SourceChanged(this, handle__SourceChanged, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SourceChanged(int flags) {

		QGraphicsOpacityEffect::sourceChanged(static_cast<QGraphicsEffect::ChangeFlags>(flags));

	}

};

void QGraphicsOpacityEffect_new(QGraphicsOpacityEffect** outptr_QGraphicsOpacityEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsOpacityEffect* ret = new MiqtVirtualQGraphicsOpacityEffect();
	*outptr_QGraphicsOpacityEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QGraphicsOpacityEffect_new2(QObject* parent, QGraphicsOpacityEffect** outptr_QGraphicsOpacityEffect, QGraphicsEffect** outptr_QGraphicsEffect, QObject** outptr_QObject) {
	MiqtVirtualQGraphicsOpacityEffect* ret = new MiqtVirtualQGraphicsOpacityEffect(parent);
	*outptr_QGraphicsOpacityEffect = ret;
	*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QGraphicsOpacityEffect_MetaObject(const QGraphicsOpacityEffect* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGraphicsOpacityEffect_Metacast(QGraphicsOpacityEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGraphicsOpacityEffect_Tr(const char* s) {
	QString _ret = QGraphicsOpacityEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

double QGraphicsOpacityEffect_Opacity(const QGraphicsOpacityEffect* self) {
	qreal _ret = self->opacity();
	return static_cast<double>(_ret);
}

QBrush* QGraphicsOpacityEffect_OpacityMask(const QGraphicsOpacityEffect* self) {
	return new QBrush(self->opacityMask());
}

void QGraphicsOpacityEffect_SetOpacity(QGraphicsOpacityEffect* self, double opacity) {
	self->setOpacity(static_cast<qreal>(opacity));
}

void QGraphicsOpacityEffect_SetOpacityMask(QGraphicsOpacityEffect* self, QBrush* mask) {
	self->setOpacityMask(*mask);
}

void QGraphicsOpacityEffect_OpacityChanged(QGraphicsOpacityEffect* self, double opacity) {
	self->opacityChanged(static_cast<qreal>(opacity));
}

void QGraphicsOpacityEffect_connect_OpacityChanged(QGraphicsOpacityEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsOpacityEffect::connect(self, static_cast<void (QGraphicsOpacityEffect::*)(qreal)>(&QGraphicsOpacityEffect::opacityChanged), self, [=](qreal opacity) {
		qreal opacity_ret = opacity;
		double sigval1 = static_cast<double>(opacity_ret);
		miqt_exec_callback_QGraphicsOpacityEffect_OpacityChanged(slot, sigval1);
	});
}

void QGraphicsOpacityEffect_OpacityMaskChanged(QGraphicsOpacityEffect* self, QBrush* mask) {
	self->opacityMaskChanged(*mask);
}

void QGraphicsOpacityEffect_connect_OpacityMaskChanged(QGraphicsOpacityEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsOpacityEffect::connect(self, static_cast<void (QGraphicsOpacityEffect::*)(const QBrush&)>(&QGraphicsOpacityEffect::opacityMaskChanged), self, [=](const QBrush& mask) {
		const QBrush& mask_ret = mask;
		// Cast returned reference into pointer
		QBrush* sigval1 = const_cast<QBrush*>(&mask_ret);
		miqt_exec_callback_QGraphicsOpacityEffect_OpacityMaskChanged(slot, sigval1);
	});
}

struct miqt_string QGraphicsOpacityEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsOpacityEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsOpacityEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsOpacityEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsOpacityEffect_override_virtual_Draw(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsOpacityEffect*>( (QGraphicsOpacityEffect*)(self) )->handle__Draw = slot;
}

void QGraphicsOpacityEffect_virtualbase_Draw(void* self, QPainter* painter) {
	( (MiqtVirtualQGraphicsOpacityEffect*)(self) )->virtualbase_Draw(painter);
}

void QGraphicsOpacityEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsOpacityEffect*>( (QGraphicsOpacityEffect*)(self) )->handle__BoundingRectFor = slot;
}

QRectF* QGraphicsOpacityEffect_virtualbase_BoundingRectFor(const void* self, QRectF* sourceRect) {
	return ( (const MiqtVirtualQGraphicsOpacityEffect*)(self) )->virtualbase_BoundingRectFor(sourceRect);
}

void QGraphicsOpacityEffect_override_virtual_SourceChanged(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsOpacityEffect*>( (QGraphicsOpacityEffect*)(self) )->handle__SourceChanged = slot;
}

void QGraphicsOpacityEffect_virtualbase_SourceChanged(void* self, int flags) {
	( (MiqtVirtualQGraphicsOpacityEffect*)(self) )->virtualbase_SourceChanged(flags);
}

void QGraphicsOpacityEffect_Delete(QGraphicsOpacityEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsOpacityEffect*>( self );
	} else {
		delete self;
	}
}

