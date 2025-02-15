#include <QAbstractTextDocumentLayout>
#define WORKAROUND_INNER_CLASS_DEFINITION_QAbstractTextDocumentLayout__PaintContext
#define WORKAROUND_INNER_CLASS_DEFINITION_QAbstractTextDocumentLayout__Selection
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPainter>
#include <QPointF>
#include <QRectF>
#include <QSizeF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTextBlock>
#include <QTextDocument>
#include <QTextFormat>
#include <QTextFrame>
#include <QTextInlineObject>
#include <QTextObjectInterface>
#include <qabstracttextdocumentlayout.h>
#include "gen_qabstracttextdocumentlayout.h"
#include "_cgo_export.h"

QMetaObject* QAbstractTextDocumentLayout_MetaObject(const QAbstractTextDocumentLayout* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAbstractTextDocumentLayout_Metacast(QAbstractTextDocumentLayout* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAbstractTextDocumentLayout_Tr(const char* s) {
	QString _ret = QAbstractTextDocumentLayout::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractTextDocumentLayout_TrUtf8(const char* s) {
	QString _ret = QAbstractTextDocumentLayout::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractTextDocumentLayout_Draw(QAbstractTextDocumentLayout* self, QPainter* painter, QAbstractTextDocumentLayout__PaintContext* context) {
	self->draw(painter, *context);
}

int QAbstractTextDocumentLayout_HitTest(const QAbstractTextDocumentLayout* self, QPointF* point, int accuracy) {
	return self->hitTest(*point, static_cast<Qt::HitTestAccuracy>(accuracy));
}

struct miqt_string QAbstractTextDocumentLayout_AnchorAt(const QAbstractTextDocumentLayout* self, QPointF* pos) {
	QString _ret = self->anchorAt(*pos);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractTextDocumentLayout_ImageAt(const QAbstractTextDocumentLayout* self, QPointF* pos) {
	QString _ret = self->imageAt(*pos);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QTextFormat* QAbstractTextDocumentLayout_FormatAt(const QAbstractTextDocumentLayout* self, QPointF* pos) {
	return new QTextFormat(self->formatAt(*pos));
}

QTextBlock* QAbstractTextDocumentLayout_BlockWithMarkerAt(const QAbstractTextDocumentLayout* self, QPointF* pos) {
	return new QTextBlock(self->blockWithMarkerAt(*pos));
}

int QAbstractTextDocumentLayout_PageCount(const QAbstractTextDocumentLayout* self) {
	return self->pageCount();
}

QSizeF* QAbstractTextDocumentLayout_DocumentSize(const QAbstractTextDocumentLayout* self) {
	return new QSizeF(self->documentSize());
}

QRectF* QAbstractTextDocumentLayout_FrameBoundingRect(const QAbstractTextDocumentLayout* self, QTextFrame* frame) {
	return new QRectF(self->frameBoundingRect(frame));
}

QRectF* QAbstractTextDocumentLayout_BlockBoundingRect(const QAbstractTextDocumentLayout* self, QTextBlock* block) {
	return new QRectF(self->blockBoundingRect(*block));
}

void QAbstractTextDocumentLayout_SetPaintDevice(QAbstractTextDocumentLayout* self, QPaintDevice* device) {
	self->setPaintDevice(device);
}

QPaintDevice* QAbstractTextDocumentLayout_PaintDevice(const QAbstractTextDocumentLayout* self) {
	return self->paintDevice();
}

QTextDocument* QAbstractTextDocumentLayout_Document(const QAbstractTextDocumentLayout* self) {
	return self->document();
}

void QAbstractTextDocumentLayout_RegisterHandler(QAbstractTextDocumentLayout* self, int objectType, QObject* component) {
	self->registerHandler(static_cast<int>(objectType), component);
}

void QAbstractTextDocumentLayout_UnregisterHandler(QAbstractTextDocumentLayout* self, int objectType) {
	self->unregisterHandler(static_cast<int>(objectType));
}

QTextObjectInterface* QAbstractTextDocumentLayout_HandlerForObject(const QAbstractTextDocumentLayout* self, int objectType) {
	return self->handlerForObject(static_cast<int>(objectType));
}

void QAbstractTextDocumentLayout_Update(QAbstractTextDocumentLayout* self) {
	self->update();
}

void QAbstractTextDocumentLayout_connect_Update(QAbstractTextDocumentLayout* self, intptr_t slot) {
	QAbstractTextDocumentLayout::connect(self, static_cast<void (QAbstractTextDocumentLayout::*)(const QRectF&)>(&QAbstractTextDocumentLayout::update), self, [=]() {
		miqt_exec_callback_QAbstractTextDocumentLayout_Update(slot);
	});
}

void QAbstractTextDocumentLayout_UpdateBlock(QAbstractTextDocumentLayout* self, QTextBlock* block) {
	self->updateBlock(*block);
}

void QAbstractTextDocumentLayout_connect_UpdateBlock(QAbstractTextDocumentLayout* self, intptr_t slot) {
	QAbstractTextDocumentLayout::connect(self, static_cast<void (QAbstractTextDocumentLayout::*)(const QTextBlock&)>(&QAbstractTextDocumentLayout::updateBlock), self, [=](const QTextBlock& block) {
		const QTextBlock& block_ret = block;
		// Cast returned reference into pointer
		QTextBlock* sigval1 = const_cast<QTextBlock*>(&block_ret);
		miqt_exec_callback_QAbstractTextDocumentLayout_UpdateBlock(slot, sigval1);
	});
}

void QAbstractTextDocumentLayout_DocumentSizeChanged(QAbstractTextDocumentLayout* self, QSizeF* newSize) {
	self->documentSizeChanged(*newSize);
}

void QAbstractTextDocumentLayout_connect_DocumentSizeChanged(QAbstractTextDocumentLayout* self, intptr_t slot) {
	QAbstractTextDocumentLayout::connect(self, static_cast<void (QAbstractTextDocumentLayout::*)(const QSizeF&)>(&QAbstractTextDocumentLayout::documentSizeChanged), self, [=](const QSizeF& newSize) {
		const QSizeF& newSize_ret = newSize;
		// Cast returned reference into pointer
		QSizeF* sigval1 = const_cast<QSizeF*>(&newSize_ret);
		miqt_exec_callback_QAbstractTextDocumentLayout_DocumentSizeChanged(slot, sigval1);
	});
}

void QAbstractTextDocumentLayout_PageCountChanged(QAbstractTextDocumentLayout* self, int newPages) {
	self->pageCountChanged(static_cast<int>(newPages));
}

void QAbstractTextDocumentLayout_connect_PageCountChanged(QAbstractTextDocumentLayout* self, intptr_t slot) {
	QAbstractTextDocumentLayout::connect(self, static_cast<void (QAbstractTextDocumentLayout::*)(int)>(&QAbstractTextDocumentLayout::pageCountChanged), self, [=](int newPages) {
		int sigval1 = newPages;
		miqt_exec_callback_QAbstractTextDocumentLayout_PageCountChanged(slot, sigval1);
	});
}

struct miqt_string QAbstractTextDocumentLayout_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractTextDocumentLayout::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractTextDocumentLayout_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractTextDocumentLayout::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractTextDocumentLayout_TrUtf82(const char* s, const char* c) {
	QString _ret = QAbstractTextDocumentLayout::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractTextDocumentLayout_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QAbstractTextDocumentLayout::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractTextDocumentLayout_UnregisterHandler2(QAbstractTextDocumentLayout* self, int objectType, QObject* component) {
	self->unregisterHandler(static_cast<int>(objectType), component);
}

void QAbstractTextDocumentLayout_Update1(QAbstractTextDocumentLayout* self, QRectF* param1) {
	self->update(*param1);
}

void QAbstractTextDocumentLayout_connect_Update1(QAbstractTextDocumentLayout* self, intptr_t slot) {
	QAbstractTextDocumentLayout::connect(self, static_cast<void (QAbstractTextDocumentLayout::*)(const QRectF&)>(&QAbstractTextDocumentLayout::update), self, [=](const QRectF& param1) {
		const QRectF& param1_ret = param1;
		// Cast returned reference into pointer
		QRectF* sigval1 = const_cast<QRectF*>(&param1_ret);
		miqt_exec_callback_QAbstractTextDocumentLayout_Update1(slot, sigval1);
	});
}

void QAbstractTextDocumentLayout_Delete(QAbstractTextDocumentLayout* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractTextDocumentLayout*>( self );
	} else {
		delete self;
	}
}

QSizeF* QTextObjectInterface_IntrinsicSize(QTextObjectInterface* self, QTextDocument* doc, int posInDocument, QTextFormat* format) {
	return new QSizeF(self->intrinsicSize(doc, static_cast<int>(posInDocument), *format));
}

void QTextObjectInterface_DrawObject(QTextObjectInterface* self, QPainter* painter, QRectF* rect, QTextDocument* doc, int posInDocument, QTextFormat* format) {
	self->drawObject(painter, *rect, doc, static_cast<int>(posInDocument), *format);
}

void QTextObjectInterface_OperatorAssign(QTextObjectInterface* self, QTextObjectInterface* param1) {
	self->operator=(*param1);
}

void QTextObjectInterface_Delete(QTextObjectInterface* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTextObjectInterface*>( self );
	} else {
		delete self;
	}
}

void QAbstractTextDocumentLayout__Selection_new(QAbstractTextDocumentLayout__Selection* param1, QAbstractTextDocumentLayout__Selection** outptr_QAbstractTextDocumentLayout__Selection) {
	QAbstractTextDocumentLayout::Selection* ret = new QAbstractTextDocumentLayout::Selection(*param1);
	*outptr_QAbstractTextDocumentLayout__Selection = ret;
}

void QAbstractTextDocumentLayout__Selection_OperatorAssign(QAbstractTextDocumentLayout__Selection* self, QAbstractTextDocumentLayout__Selection* param1) {
	self->operator=(*param1);
}

void QAbstractTextDocumentLayout__Selection_Delete(QAbstractTextDocumentLayout__Selection* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractTextDocumentLayout::Selection*>( self );
	} else {
		delete self;
	}
}

void QAbstractTextDocumentLayout__PaintContext_new(QAbstractTextDocumentLayout__PaintContext** outptr_QAbstractTextDocumentLayout__PaintContext) {
	QAbstractTextDocumentLayout::PaintContext* ret = new QAbstractTextDocumentLayout::PaintContext();
	*outptr_QAbstractTextDocumentLayout__PaintContext = ret;
}

void QAbstractTextDocumentLayout__PaintContext_new2(QAbstractTextDocumentLayout__PaintContext* param1, QAbstractTextDocumentLayout__PaintContext** outptr_QAbstractTextDocumentLayout__PaintContext) {
	QAbstractTextDocumentLayout::PaintContext* ret = new QAbstractTextDocumentLayout::PaintContext(*param1);
	*outptr_QAbstractTextDocumentLayout__PaintContext = ret;
}

void QAbstractTextDocumentLayout__PaintContext_OperatorAssign(QAbstractTextDocumentLayout__PaintContext* self, QAbstractTextDocumentLayout__PaintContext* param1) {
	self->operator=(*param1);
}

void QAbstractTextDocumentLayout__PaintContext_Delete(QAbstractTextDocumentLayout__PaintContext* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractTextDocumentLayout::PaintContext*>( self );
	} else {
		delete self;
	}
}

