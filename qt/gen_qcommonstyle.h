#ifndef GEN_QCOMMONSTYLE_H
#define GEN_QCOMMONSTYLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "binding.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QApplication;
class QCommonStyle;
class QIcon;
class QMetaObject;
class QPainter;
class QPalette;
class QPixmap;
class QPoint;
class QRect;
class QSize;
class QStyleHintReturn;
class QStyleOption;
class QStyleOptionComplex;
class QWidget;
#else
typedef struct QApplication QApplication;
typedef struct QCommonStyle QCommonStyle;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QPainter QPainter;
typedef struct QPalette QPalette;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStyleHintReturn QStyleHintReturn;
typedef struct QStyleOption QStyleOption;
typedef struct QStyleOptionComplex QStyleOptionComplex;
typedef struct QWidget QWidget;
#endif

QCommonStyle* QCommonStyle_new();
QMetaObject* QCommonStyle_MetaObject(const QCommonStyle* self);
struct miqt_string* QCommonStyle_Tr(const char* s);
struct miqt_string* QCommonStyle_TrUtf8(const char* s);
void QCommonStyle_DrawPrimitive(const QCommonStyle* self, uintptr_t pe, QStyleOption* opt, QPainter* p);
void QCommonStyle_DrawControl(const QCommonStyle* self, uintptr_t element, QStyleOption* opt, QPainter* p);
QRect* QCommonStyle_SubElementRect(const QCommonStyle* self, uintptr_t r, QStyleOption* opt);
void QCommonStyle_DrawComplexControl(const QCommonStyle* self, uintptr_t cc, QStyleOptionComplex* opt, QPainter* p);
uintptr_t QCommonStyle_HitTestComplexControl(const QCommonStyle* self, uintptr_t cc, QStyleOptionComplex* opt, QPoint* pt);
QRect* QCommonStyle_SubControlRect(const QCommonStyle* self, uintptr_t cc, QStyleOptionComplex* opt, uintptr_t sc);
QSize* QCommonStyle_SizeFromContents(const QCommonStyle* self, uintptr_t ct, QStyleOption* opt, QSize* contentsSize);
int QCommonStyle_PixelMetric(const QCommonStyle* self, uintptr_t m);
int QCommonStyle_StyleHint(const QCommonStyle* self, uintptr_t sh);
QIcon* QCommonStyle_StandardIcon(const QCommonStyle* self, uintptr_t standardIcon);
QPixmap* QCommonStyle_StandardPixmap(const QCommonStyle* self, uintptr_t sp);
QPixmap* QCommonStyle_GeneratedIconPixmap(const QCommonStyle* self, uintptr_t iconMode, QPixmap* pixmap, QStyleOption* opt);
int QCommonStyle_LayoutSpacing(const QCommonStyle* self, uintptr_t control1, uintptr_t control2, uintptr_t orientation);
void QCommonStyle_Polish(QCommonStyle* self, QPalette* param1);
void QCommonStyle_PolishWithApp(QCommonStyle* self, QApplication* app);
void QCommonStyle_PolishWithWidget(QCommonStyle* self, QWidget* widget);
void QCommonStyle_Unpolish(QCommonStyle* self, QWidget* widget);
void QCommonStyle_UnpolishWithApplication(QCommonStyle* self, QApplication* application);
struct miqt_string* QCommonStyle_Tr2(const char* s, const char* c);
struct miqt_string* QCommonStyle_Tr3(const char* s, const char* c, int n);
struct miqt_string* QCommonStyle_TrUtf82(const char* s, const char* c);
struct miqt_string* QCommonStyle_TrUtf83(const char* s, const char* c, int n);
void QCommonStyle_DrawPrimitive4(const QCommonStyle* self, uintptr_t pe, QStyleOption* opt, QPainter* p, QWidget* w);
void QCommonStyle_DrawControl4(const QCommonStyle* self, uintptr_t element, QStyleOption* opt, QPainter* p, QWidget* w);
QRect* QCommonStyle_SubElementRect3(const QCommonStyle* self, uintptr_t r, QStyleOption* opt, QWidget* widget);
void QCommonStyle_DrawComplexControl4(const QCommonStyle* self, uintptr_t cc, QStyleOptionComplex* opt, QPainter* p, QWidget* w);
uintptr_t QCommonStyle_HitTestComplexControl4(const QCommonStyle* self, uintptr_t cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* w);
QRect* QCommonStyle_SubControlRect4(const QCommonStyle* self, uintptr_t cc, QStyleOptionComplex* opt, uintptr_t sc, QWidget* w);
QSize* QCommonStyle_SizeFromContents4(const QCommonStyle* self, uintptr_t ct, QStyleOption* opt, QSize* contentsSize, QWidget* widget);
int QCommonStyle_PixelMetric2(const QCommonStyle* self, uintptr_t m, QStyleOption* opt);
int QCommonStyle_PixelMetric3(const QCommonStyle* self, uintptr_t m, QStyleOption* opt, QWidget* widget);
int QCommonStyle_StyleHint2(const QCommonStyle* self, uintptr_t sh, QStyleOption* opt);
int QCommonStyle_StyleHint3(const QCommonStyle* self, uintptr_t sh, QStyleOption* opt, QWidget* w);
int QCommonStyle_StyleHint4(const QCommonStyle* self, uintptr_t sh, QStyleOption* opt, QWidget* w, QStyleHintReturn* shret);
QIcon* QCommonStyle_StandardIcon2(const QCommonStyle* self, uintptr_t standardIcon, QStyleOption* opt);
QIcon* QCommonStyle_StandardIcon3(const QCommonStyle* self, uintptr_t standardIcon, QStyleOption* opt, QWidget* widget);
QPixmap* QCommonStyle_StandardPixmap2(const QCommonStyle* self, uintptr_t sp, QStyleOption* opt);
QPixmap* QCommonStyle_StandardPixmap3(const QCommonStyle* self, uintptr_t sp, QStyleOption* opt, QWidget* widget);
int QCommonStyle_LayoutSpacing4(const QCommonStyle* self, uintptr_t control1, uintptr_t control2, uintptr_t orientation, QStyleOption* option);
int QCommonStyle_LayoutSpacing5(const QCommonStyle* self, uintptr_t control1, uintptr_t control2, uintptr_t orientation, QStyleOption* option, QWidget* widget);
void QCommonStyle_Delete(QCommonStyle* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
