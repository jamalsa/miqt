#include <QMargins>
#include <QMarginsF>
#include <QPoint>
#include <QPointF>
#include <QRect>
#include <QRectF>
#include <QSize>
#include <QSizeF>
#include "qrect.h"
#include "gen_qrect.h"
#include "_cgo_export.h"

QRect* QRect_new() {
	return new QRect();
}

QRect* QRect_new2(QPoint* topleft, QPoint* bottomright) {
	return new QRect(*topleft, *bottomright);
}

QRect* QRect_new3(QPoint* topleft, QSize* size) {
	return new QRect(*topleft, *size);
}

QRect* QRect_new4(int left, int top, int width, int height) {
	return new QRect(static_cast<int>(left), static_cast<int>(top), static_cast<int>(width), static_cast<int>(height));
}

QRect* QRect_new5(QRect* param1) {
	return new QRect(*param1);
}

bool QRect_IsNull(const QRect* self) {
	return self->isNull();
}

bool QRect_IsEmpty(const QRect* self) {
	return self->isEmpty();
}

bool QRect_IsValid(const QRect* self) {
	return self->isValid();
}

int QRect_Left(const QRect* self) {
	return self->left();
}

int QRect_Top(const QRect* self) {
	return self->top();
}

int QRect_Right(const QRect* self) {
	return self->right();
}

int QRect_Bottom(const QRect* self) {
	return self->bottom();
}

QRect* QRect_Normalized(const QRect* self) {
	QRect _ret = self->normalized();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

int QRect_X(const QRect* self) {
	return self->x();
}

int QRect_Y(const QRect* self) {
	return self->y();
}

void QRect_SetLeft(QRect* self, int pos) {
	self->setLeft(static_cast<int>(pos));
}

void QRect_SetTop(QRect* self, int pos) {
	self->setTop(static_cast<int>(pos));
}

void QRect_SetRight(QRect* self, int pos) {
	self->setRight(static_cast<int>(pos));
}

void QRect_SetBottom(QRect* self, int pos) {
	self->setBottom(static_cast<int>(pos));
}

void QRect_SetX(QRect* self, int x) {
	self->setX(static_cast<int>(x));
}

void QRect_SetY(QRect* self, int y) {
	self->setY(static_cast<int>(y));
}

void QRect_SetTopLeft(QRect* self, QPoint* p) {
	self->setTopLeft(*p);
}

void QRect_SetBottomRight(QRect* self, QPoint* p) {
	self->setBottomRight(*p);
}

void QRect_SetTopRight(QRect* self, QPoint* p) {
	self->setTopRight(*p);
}

void QRect_SetBottomLeft(QRect* self, QPoint* p) {
	self->setBottomLeft(*p);
}

QPoint* QRect_TopLeft(const QRect* self) {
	QPoint _ret = self->topLeft();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPoint*>(new QPoint(_ret));
}

QPoint* QRect_BottomRight(const QRect* self) {
	QPoint _ret = self->bottomRight();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPoint*>(new QPoint(_ret));
}

QPoint* QRect_TopRight(const QRect* self) {
	QPoint _ret = self->topRight();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPoint*>(new QPoint(_ret));
}

QPoint* QRect_BottomLeft(const QRect* self) {
	QPoint _ret = self->bottomLeft();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPoint*>(new QPoint(_ret));
}

QPoint* QRect_Center(const QRect* self) {
	QPoint _ret = self->center();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPoint*>(new QPoint(_ret));
}

void QRect_MoveLeft(QRect* self, int pos) {
	self->moveLeft(static_cast<int>(pos));
}

void QRect_MoveTop(QRect* self, int pos) {
	self->moveTop(static_cast<int>(pos));
}

void QRect_MoveRight(QRect* self, int pos) {
	self->moveRight(static_cast<int>(pos));
}

void QRect_MoveBottom(QRect* self, int pos) {
	self->moveBottom(static_cast<int>(pos));
}

void QRect_MoveTopLeft(QRect* self, QPoint* p) {
	self->moveTopLeft(*p);
}

void QRect_MoveBottomRight(QRect* self, QPoint* p) {
	self->moveBottomRight(*p);
}

void QRect_MoveTopRight(QRect* self, QPoint* p) {
	self->moveTopRight(*p);
}

void QRect_MoveBottomLeft(QRect* self, QPoint* p) {
	self->moveBottomLeft(*p);
}

void QRect_MoveCenter(QRect* self, QPoint* p) {
	self->moveCenter(*p);
}

void QRect_Translate(QRect* self, int dx, int dy) {
	self->translate(static_cast<int>(dx), static_cast<int>(dy));
}

void QRect_TranslateWithQPoint(QRect* self, QPoint* p) {
	self->translate(*p);
}

QRect* QRect_Translated(const QRect* self, int dx, int dy) {
	QRect _ret = self->translated(static_cast<int>(dx), static_cast<int>(dy));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QRect* QRect_TranslatedWithQPoint(const QRect* self, QPoint* p) {
	QRect _ret = self->translated(*p);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QRect* QRect_Transposed(const QRect* self) {
	QRect _ret = self->transposed();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

void QRect_MoveTo(QRect* self, int x, int t) {
	self->moveTo(static_cast<int>(x), static_cast<int>(t));
}

void QRect_MoveToWithQPoint(QRect* self, QPoint* p) {
	self->moveTo(*p);
}

void QRect_SetRect(QRect* self, int x, int y, int w, int h) {
	self->setRect(static_cast<int>(x), static_cast<int>(y), static_cast<int>(w), static_cast<int>(h));
}

void QRect_GetRect(const QRect* self, int* x, int* y, int* w, int* h) {
	self->getRect(static_cast<int*>(x), static_cast<int*>(y), static_cast<int*>(w), static_cast<int*>(h));
}

void QRect_SetCoords(QRect* self, int x1, int y1, int x2, int y2) {
	self->setCoords(static_cast<int>(x1), static_cast<int>(y1), static_cast<int>(x2), static_cast<int>(y2));
}

void QRect_GetCoords(const QRect* self, int* x1, int* y1, int* x2, int* y2) {
	self->getCoords(static_cast<int*>(x1), static_cast<int*>(y1), static_cast<int*>(x2), static_cast<int*>(y2));
}

void QRect_Adjust(QRect* self, int x1, int y1, int x2, int y2) {
	self->adjust(static_cast<int>(x1), static_cast<int>(y1), static_cast<int>(x2), static_cast<int>(y2));
}

QRect* QRect_Adjusted(const QRect* self, int x1, int y1, int x2, int y2) {
	QRect _ret = self->adjusted(static_cast<int>(x1), static_cast<int>(y1), static_cast<int>(x2), static_cast<int>(y2));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QSize* QRect_Size(const QRect* self) {
	QSize _ret = self->size();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(_ret));
}

int QRect_Width(const QRect* self) {
	return self->width();
}

int QRect_Height(const QRect* self) {
	return self->height();
}

void QRect_SetWidth(QRect* self, int w) {
	self->setWidth(static_cast<int>(w));
}

void QRect_SetHeight(QRect* self, int h) {
	self->setHeight(static_cast<int>(h));
}

void QRect_SetSize(QRect* self, QSize* s) {
	self->setSize(*s);
}

QRect* QRect_OperatorBitwiseOr(const QRect* self, QRect* r) {
	QRect _ret = self->operator|(*r);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QRect* QRect_OperatorBitwiseAnd(const QRect* self, QRect* r) {
	QRect _ret = self->operator&(*r);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

void QRect_OperatorBitwiseOrAssign(QRect* self, QRect* r) {
	self->operator|=(*r);
}

void QRect_OperatorBitwiseAndAssign(QRect* self, QRect* r) {
	self->operator&=(*r);
}

bool QRect_Contains(const QRect* self, QRect* r) {
	return self->contains(*r);
}

bool QRect_ContainsWithQPoint(const QRect* self, QPoint* p) {
	return self->contains(*p);
}

bool QRect_Contains2(const QRect* self, int x, int y) {
	return self->contains(static_cast<int>(x), static_cast<int>(y));
}

bool QRect_Contains3(const QRect* self, int x, int y, bool proper) {
	return self->contains(static_cast<int>(x), static_cast<int>(y), proper);
}

QRect* QRect_United(const QRect* self, QRect* other) {
	QRect _ret = self->united(*other);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QRect* QRect_Intersected(const QRect* self, QRect* other) {
	QRect _ret = self->intersected(*other);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

bool QRect_Intersects(const QRect* self, QRect* r) {
	return self->intersects(*r);
}

QRect* QRect_MarginsAdded(const QRect* self, QMargins* margins) {
	QRect _ret = self->marginsAdded(*margins);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QRect* QRect_MarginsRemoved(const QRect* self, QMargins* margins) {
	QRect _ret = self->marginsRemoved(*margins);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QRect* QRect_OperatorPlusAssign(QRect* self, QMargins* margins) {
	QRect& _ret = self->operator+=(*margins);
	// Cast returned reference into pointer
	return &_ret;
}

QRect* QRect_OperatorMinusAssign(QRect* self, QMargins* margins) {
	QRect& _ret = self->operator-=(*margins);
	// Cast returned reference into pointer
	return &_ret;
}

bool QRect_Contains22(const QRect* self, QRect* r, bool proper) {
	return self->contains(*r, proper);
}

bool QRect_Contains23(const QRect* self, QPoint* p, bool proper) {
	return self->contains(*p, proper);
}

void QRect_Delete(QRect* self) {
	delete self;
}

QRectF* QRectF_new() {
	return new QRectF();
}

QRectF* QRectF_new2(QPointF* topleft, QSizeF* size) {
	return new QRectF(*topleft, *size);
}

QRectF* QRectF_new3(QPointF* topleft, QPointF* bottomRight) {
	return new QRectF(*topleft, *bottomRight);
}

QRectF* QRectF_new4(double left, double top, double width, double height) {
	return new QRectF(static_cast<qreal>(left), static_cast<qreal>(top), static_cast<qreal>(width), static_cast<qreal>(height));
}

QRectF* QRectF_new5(QRect* rect) {
	return new QRectF(*rect);
}

QRectF* QRectF_new6(QRectF* param1) {
	return new QRectF(*param1);
}

bool QRectF_IsNull(const QRectF* self) {
	return self->isNull();
}

bool QRectF_IsEmpty(const QRectF* self) {
	return self->isEmpty();
}

bool QRectF_IsValid(const QRectF* self) {
	return self->isValid();
}

QRectF* QRectF_Normalized(const QRectF* self) {
	QRectF _ret = self->normalized();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

double QRectF_Left(const QRectF* self) {
	return self->left();
}

double QRectF_Top(const QRectF* self) {
	return self->top();
}

double QRectF_Right(const QRectF* self) {
	return self->right();
}

double QRectF_Bottom(const QRectF* self) {
	return self->bottom();
}

double QRectF_X(const QRectF* self) {
	return self->x();
}

double QRectF_Y(const QRectF* self) {
	return self->y();
}

void QRectF_SetLeft(QRectF* self, double pos) {
	self->setLeft(static_cast<qreal>(pos));
}

void QRectF_SetTop(QRectF* self, double pos) {
	self->setTop(static_cast<qreal>(pos));
}

void QRectF_SetRight(QRectF* self, double pos) {
	self->setRight(static_cast<qreal>(pos));
}

void QRectF_SetBottom(QRectF* self, double pos) {
	self->setBottom(static_cast<qreal>(pos));
}

void QRectF_SetX(QRectF* self, double pos) {
	self->setX(static_cast<qreal>(pos));
}

void QRectF_SetY(QRectF* self, double pos) {
	self->setY(static_cast<qreal>(pos));
}

QPointF* QRectF_TopLeft(const QRectF* self) {
	QPointF _ret = self->topLeft();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPointF*>(new QPointF(_ret));
}

QPointF* QRectF_BottomRight(const QRectF* self) {
	QPointF _ret = self->bottomRight();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPointF*>(new QPointF(_ret));
}

QPointF* QRectF_TopRight(const QRectF* self) {
	QPointF _ret = self->topRight();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPointF*>(new QPointF(_ret));
}

QPointF* QRectF_BottomLeft(const QRectF* self) {
	QPointF _ret = self->bottomLeft();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPointF*>(new QPointF(_ret));
}

QPointF* QRectF_Center(const QRectF* self) {
	QPointF _ret = self->center();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QPointF*>(new QPointF(_ret));
}

void QRectF_SetTopLeft(QRectF* self, QPointF* p) {
	self->setTopLeft(*p);
}

void QRectF_SetBottomRight(QRectF* self, QPointF* p) {
	self->setBottomRight(*p);
}

void QRectF_SetTopRight(QRectF* self, QPointF* p) {
	self->setTopRight(*p);
}

void QRectF_SetBottomLeft(QRectF* self, QPointF* p) {
	self->setBottomLeft(*p);
}

void QRectF_MoveLeft(QRectF* self, double pos) {
	self->moveLeft(static_cast<qreal>(pos));
}

void QRectF_MoveTop(QRectF* self, double pos) {
	self->moveTop(static_cast<qreal>(pos));
}

void QRectF_MoveRight(QRectF* self, double pos) {
	self->moveRight(static_cast<qreal>(pos));
}

void QRectF_MoveBottom(QRectF* self, double pos) {
	self->moveBottom(static_cast<qreal>(pos));
}

void QRectF_MoveTopLeft(QRectF* self, QPointF* p) {
	self->moveTopLeft(*p);
}

void QRectF_MoveBottomRight(QRectF* self, QPointF* p) {
	self->moveBottomRight(*p);
}

void QRectF_MoveTopRight(QRectF* self, QPointF* p) {
	self->moveTopRight(*p);
}

void QRectF_MoveBottomLeft(QRectF* self, QPointF* p) {
	self->moveBottomLeft(*p);
}

void QRectF_MoveCenter(QRectF* self, QPointF* p) {
	self->moveCenter(*p);
}

void QRectF_Translate(QRectF* self, double dx, double dy) {
	self->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

void QRectF_TranslateWithQPointF(QRectF* self, QPointF* p) {
	self->translate(*p);
}

QRectF* QRectF_Translated(const QRectF* self, double dx, double dy) {
	QRectF _ret = self->translated(static_cast<qreal>(dx), static_cast<qreal>(dy));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

QRectF* QRectF_TranslatedWithQPointF(const QRectF* self, QPointF* p) {
	QRectF _ret = self->translated(*p);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

QRectF* QRectF_Transposed(const QRectF* self) {
	QRectF _ret = self->transposed();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

void QRectF_MoveTo(QRectF* self, double x, double y) {
	self->moveTo(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QRectF_MoveToWithQPointF(QRectF* self, QPointF* p) {
	self->moveTo(*p);
}

void QRectF_SetRect(QRectF* self, double x, double y, double w, double h) {
	self->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QRectF_GetRect(const QRectF* self, double* x, double* y, double* w, double* h) {
	self->getRect(static_cast<qreal*>(x), static_cast<qreal*>(y), static_cast<qreal*>(w), static_cast<qreal*>(h));
}

void QRectF_SetCoords(QRectF* self, double x1, double y1, double x2, double y2) {
	self->setCoords(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

void QRectF_GetCoords(const QRectF* self, double* x1, double* y1, double* x2, double* y2) {
	self->getCoords(static_cast<qreal*>(x1), static_cast<qreal*>(y1), static_cast<qreal*>(x2), static_cast<qreal*>(y2));
}

void QRectF_Adjust(QRectF* self, double x1, double y1, double x2, double y2) {
	self->adjust(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

QRectF* QRectF_Adjusted(const QRectF* self, double x1, double y1, double x2, double y2) {
	QRectF _ret = self->adjusted(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

QSizeF* QRectF_Size(const QRectF* self) {
	QSizeF _ret = self->size();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSizeF*>(new QSizeF(_ret));
}

double QRectF_Width(const QRectF* self) {
	return self->width();
}

double QRectF_Height(const QRectF* self) {
	return self->height();
}

void QRectF_SetWidth(QRectF* self, double w) {
	self->setWidth(static_cast<qreal>(w));
}

void QRectF_SetHeight(QRectF* self, double h) {
	self->setHeight(static_cast<qreal>(h));
}

void QRectF_SetSize(QRectF* self, QSizeF* s) {
	self->setSize(*s);
}

QRectF* QRectF_OperatorBitwiseOr(const QRectF* self, QRectF* r) {
	QRectF _ret = self->operator|(*r);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

QRectF* QRectF_OperatorBitwiseAnd(const QRectF* self, QRectF* r) {
	QRectF _ret = self->operator&(*r);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

void QRectF_OperatorBitwiseOrAssign(QRectF* self, QRectF* r) {
	self->operator|=(*r);
}

void QRectF_OperatorBitwiseAndAssign(QRectF* self, QRectF* r) {
	self->operator&=(*r);
}

bool QRectF_Contains(const QRectF* self, QRectF* r) {
	return self->contains(*r);
}

bool QRectF_ContainsWithQPointF(const QRectF* self, QPointF* p) {
	return self->contains(*p);
}

bool QRectF_Contains2(const QRectF* self, double x, double y) {
	return self->contains(static_cast<qreal>(x), static_cast<qreal>(y));
}

QRectF* QRectF_United(const QRectF* self, QRectF* other) {
	QRectF _ret = self->united(*other);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

QRectF* QRectF_Intersected(const QRectF* self, QRectF* other) {
	QRectF _ret = self->intersected(*other);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

bool QRectF_Intersects(const QRectF* self, QRectF* r) {
	return self->intersects(*r);
}

QRectF* QRectF_MarginsAdded(const QRectF* self, QMarginsF* margins) {
	QRectF _ret = self->marginsAdded(*margins);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

QRectF* QRectF_MarginsRemoved(const QRectF* self, QMarginsF* margins) {
	QRectF _ret = self->marginsRemoved(*margins);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRectF*>(new QRectF(_ret));
}

QRectF* QRectF_OperatorPlusAssign(QRectF* self, QMarginsF* margins) {
	QRectF& _ret = self->operator+=(*margins);
	// Cast returned reference into pointer
	return &_ret;
}

QRectF* QRectF_OperatorMinusAssign(QRectF* self, QMarginsF* margins) {
	QRectF& _ret = self->operator-=(*margins);
	// Cast returned reference into pointer
	return &_ret;
}

QRect* QRectF_ToRect(const QRectF* self) {
	QRect _ret = self->toRect();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

QRect* QRectF_ToAlignedRect(const QRectF* self) {
	QRect _ret = self->toAlignedRect();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(_ret));
}

void QRectF_Delete(QRectF* self) {
	delete self;
}

