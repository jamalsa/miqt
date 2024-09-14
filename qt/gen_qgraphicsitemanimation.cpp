#include <QGraphicsItem>
#include <QGraphicsItemAnimation>
#include <QMatrix>
#include <QMetaObject>
#include <QObject>
#include <QPointF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimeLine>
#include <QTransform>
#include "qgraphicsitemanimation.h"
#include "gen_qgraphicsitemanimation.h"
#include "_cgo_export.h"

QGraphicsItemAnimation* QGraphicsItemAnimation_new() {
	return new QGraphicsItemAnimation();
}

QGraphicsItemAnimation* QGraphicsItemAnimation_new2(QObject* parent) {
	return new QGraphicsItemAnimation(parent);
}

QMetaObject* QGraphicsItemAnimation_MetaObject(const QGraphicsItemAnimation* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QGraphicsItemAnimation_Tr(const char* s) {
	QString _ret = QGraphicsItemAnimation::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGraphicsItemAnimation_TrUtf8(const char* s) {
	QString _ret = QGraphicsItemAnimation::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QGraphicsItem* QGraphicsItemAnimation_Item(const QGraphicsItemAnimation* self) {
	return self->item();
}

void QGraphicsItemAnimation_SetItem(QGraphicsItemAnimation* self, QGraphicsItem* item) {
	self->setItem(item);
}

QTimeLine* QGraphicsItemAnimation_TimeLine(const QGraphicsItemAnimation* self) {
	return self->timeLine();
}

void QGraphicsItemAnimation_SetTimeLine(QGraphicsItemAnimation* self, QTimeLine* timeLine) {
	self->setTimeLine(timeLine);
}

QPointF* QGraphicsItemAnimation_PosAt(const QGraphicsItemAnimation* self, double step) {
	QPointF _ret = self->posAt(static_cast<qreal>(step));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPointF*>(new QPointF(_ret));
}

void QGraphicsItemAnimation_SetPosAt(QGraphicsItemAnimation* self, double step, QPointF* pos) {
	self->setPosAt(static_cast<qreal>(step), *pos);
}

QMatrix* QGraphicsItemAnimation_MatrixAt(const QGraphicsItemAnimation* self, double step) {
	QMatrix _ret = self->matrixAt(static_cast<qreal>(step));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QMatrix*>(new QMatrix(_ret));
}

QTransform* QGraphicsItemAnimation_TransformAt(const QGraphicsItemAnimation* self, double step) {
	QTransform _ret = self->transformAt(static_cast<qreal>(step));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTransform*>(new QTransform(_ret));
}

double QGraphicsItemAnimation_RotationAt(const QGraphicsItemAnimation* self, double step) {
	return self->rotationAt(static_cast<qreal>(step));
}

void QGraphicsItemAnimation_SetRotationAt(QGraphicsItemAnimation* self, double step, double angle) {
	self->setRotationAt(static_cast<qreal>(step), static_cast<qreal>(angle));
}

double QGraphicsItemAnimation_XTranslationAt(const QGraphicsItemAnimation* self, double step) {
	return self->xTranslationAt(static_cast<qreal>(step));
}

double QGraphicsItemAnimation_YTranslationAt(const QGraphicsItemAnimation* self, double step) {
	return self->yTranslationAt(static_cast<qreal>(step));
}

void QGraphicsItemAnimation_SetTranslationAt(QGraphicsItemAnimation* self, double step, double dx, double dy) {
	self->setTranslationAt(static_cast<qreal>(step), static_cast<qreal>(dx), static_cast<qreal>(dy));
}

double QGraphicsItemAnimation_VerticalScaleAt(const QGraphicsItemAnimation* self, double step) {
	return self->verticalScaleAt(static_cast<qreal>(step));
}

double QGraphicsItemAnimation_HorizontalScaleAt(const QGraphicsItemAnimation* self, double step) {
	return self->horizontalScaleAt(static_cast<qreal>(step));
}

void QGraphicsItemAnimation_SetScaleAt(QGraphicsItemAnimation* self, double step, double sx, double sy) {
	self->setScaleAt(static_cast<qreal>(step), static_cast<qreal>(sx), static_cast<qreal>(sy));
}

double QGraphicsItemAnimation_VerticalShearAt(const QGraphicsItemAnimation* self, double step) {
	return self->verticalShearAt(static_cast<qreal>(step));
}

double QGraphicsItemAnimation_HorizontalShearAt(const QGraphicsItemAnimation* self, double step) {
	return self->horizontalShearAt(static_cast<qreal>(step));
}

void QGraphicsItemAnimation_SetShearAt(QGraphicsItemAnimation* self, double step, double sh, double sv) {
	self->setShearAt(static_cast<qreal>(step), static_cast<qreal>(sh), static_cast<qreal>(sv));
}

void QGraphicsItemAnimation_Clear(QGraphicsItemAnimation* self) {
	self->clear();
}

void QGraphicsItemAnimation_SetStep(QGraphicsItemAnimation* self, double x) {
	self->setStep(static_cast<qreal>(x));
}

void QGraphicsItemAnimation_Reset(QGraphicsItemAnimation* self) {
	self->reset();
}

struct miqt_string* QGraphicsItemAnimation_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsItemAnimation::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGraphicsItemAnimation_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsItemAnimation::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGraphicsItemAnimation_TrUtf82(const char* s, const char* c) {
	QString _ret = QGraphicsItemAnimation::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QGraphicsItemAnimation_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QGraphicsItemAnimation::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QGraphicsItemAnimation_Delete(QGraphicsItemAnimation* self) {
	delete self;
}

