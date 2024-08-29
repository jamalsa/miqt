#include <QMetaObject>
#include <QScrollArea>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include "qscrollarea.h"

#include "gen_qscrollarea.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QScrollArea* QScrollArea_new() {
	return new QScrollArea();
}

QScrollArea* QScrollArea_new2(QWidget* parent) {
	return new QScrollArea(parent);
}

QMetaObject* QScrollArea_MetaObject(QScrollArea* self) {
	return (QMetaObject*) const_cast<const QScrollArea*>(self)->metaObject();
}

void QScrollArea_Tr(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QScrollArea::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QScrollArea_TrUtf8(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QScrollArea::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

QWidget* QScrollArea_Widget(QScrollArea* self) {
	return const_cast<const QScrollArea*>(self)->widget();
}

void QScrollArea_SetWidget(QScrollArea* self, QWidget* widget) {
	self->setWidget(widget);
}

QWidget* QScrollArea_TakeWidget(QScrollArea* self) {
	return self->takeWidget();
}

bool QScrollArea_WidgetResizable(QScrollArea* self) {
	return const_cast<const QScrollArea*>(self)->widgetResizable();
}

void QScrollArea_SetWidgetResizable(QScrollArea* self, bool resizable) {
	self->setWidgetResizable(resizable);
}

QSize* QScrollArea_SizeHint(QScrollArea* self) {
	QSize ret = const_cast<const QScrollArea*>(self)->sizeHint();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
}

bool QScrollArea_FocusNextPrevChild(QScrollArea* self, bool next) {
	return self->focusNextPrevChild(next);
}

int QScrollArea_Alignment(QScrollArea* self) {
	Qt::Alignment ret = const_cast<const QScrollArea*>(self)->alignment();
	return static_cast<int>(ret);
}

void QScrollArea_SetAlignment(QScrollArea* self, int alignment) {
	self->setAlignment(static_cast<Qt::Alignment>(alignment));
}

void QScrollArea_EnsureVisible(QScrollArea* self, int x, int y) {
	self->ensureVisible(static_cast<int>(x), static_cast<int>(y));
}

void QScrollArea_EnsureWidgetVisible(QScrollArea* self, QWidget* childWidget) {
	self->ensureWidgetVisible(childWidget);
}

void QScrollArea_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QScrollArea::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QScrollArea_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QScrollArea::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QScrollArea_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QScrollArea::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QScrollArea_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QScrollArea::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QScrollArea_EnsureVisible3(QScrollArea* self, int x, int y, int xmargin) {
	self->ensureVisible(static_cast<int>(x), static_cast<int>(y), static_cast<int>(xmargin));
}

void QScrollArea_EnsureVisible4(QScrollArea* self, int x, int y, int xmargin, int ymargin) {
	self->ensureVisible(static_cast<int>(x), static_cast<int>(y), static_cast<int>(xmargin), static_cast<int>(ymargin));
}

void QScrollArea_EnsureWidgetVisible2(QScrollArea* self, QWidget* childWidget, int xmargin) {
	self->ensureWidgetVisible(childWidget, static_cast<int>(xmargin));
}

void QScrollArea_EnsureWidgetVisible3(QScrollArea* self, QWidget* childWidget, int xmargin, int ymargin) {
	self->ensureWidgetVisible(childWidget, static_cast<int>(xmargin), static_cast<int>(ymargin));
}

void QScrollArea_Delete(QScrollArea* self) {
	delete self;
}

