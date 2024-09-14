#ifndef GEN_QCALENDARWIDGET_H
#define GEN_QCALENDARWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "binding.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCalendar;
class QCalendarWidget;
class QDate;
class QMetaObject;
class QSize;
class QTextCharFormat;
class QWidget;
#else
typedef struct QCalendar QCalendar;
typedef struct QCalendarWidget QCalendarWidget;
typedef struct QDate QDate;
typedef struct QMetaObject QMetaObject;
typedef struct QSize QSize;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QWidget QWidget;
#endif

QCalendarWidget* QCalendarWidget_new();
QCalendarWidget* QCalendarWidget_new2(QWidget* parent);
QMetaObject* QCalendarWidget_MetaObject(const QCalendarWidget* self);
struct miqt_string* QCalendarWidget_Tr(const char* s);
struct miqt_string* QCalendarWidget_TrUtf8(const char* s);
QSize* QCalendarWidget_SizeHint(const QCalendarWidget* self);
QSize* QCalendarWidget_MinimumSizeHint(const QCalendarWidget* self);
QDate* QCalendarWidget_SelectedDate(const QCalendarWidget* self);
int QCalendarWidget_YearShown(const QCalendarWidget* self);
int QCalendarWidget_MonthShown(const QCalendarWidget* self);
QDate* QCalendarWidget_MinimumDate(const QCalendarWidget* self);
void QCalendarWidget_SetMinimumDate(QCalendarWidget* self, QDate* date);
QDate* QCalendarWidget_MaximumDate(const QCalendarWidget* self);
void QCalendarWidget_SetMaximumDate(QCalendarWidget* self, QDate* date);
uintptr_t QCalendarWidget_FirstDayOfWeek(const QCalendarWidget* self);
void QCalendarWidget_SetFirstDayOfWeek(QCalendarWidget* self, uintptr_t dayOfWeek);
bool QCalendarWidget_IsNavigationBarVisible(const QCalendarWidget* self);
bool QCalendarWidget_IsGridVisible(const QCalendarWidget* self);
QCalendar* QCalendarWidget_Calendar(const QCalendarWidget* self);
void QCalendarWidget_SetCalendar(QCalendarWidget* self, QCalendar* calendar);
uintptr_t QCalendarWidget_SelectionMode(const QCalendarWidget* self);
void QCalendarWidget_SetSelectionMode(QCalendarWidget* self, uintptr_t mode);
uintptr_t QCalendarWidget_HorizontalHeaderFormat(const QCalendarWidget* self);
void QCalendarWidget_SetHorizontalHeaderFormat(QCalendarWidget* self, uintptr_t format);
uintptr_t QCalendarWidget_VerticalHeaderFormat(const QCalendarWidget* self);
void QCalendarWidget_SetVerticalHeaderFormat(QCalendarWidget* self, uintptr_t format);
QTextCharFormat* QCalendarWidget_HeaderTextFormat(const QCalendarWidget* self);
void QCalendarWidget_SetHeaderTextFormat(QCalendarWidget* self, QTextCharFormat* format);
QTextCharFormat* QCalendarWidget_WeekdayTextFormat(const QCalendarWidget* self, uintptr_t dayOfWeek);
void QCalendarWidget_SetWeekdayTextFormat(QCalendarWidget* self, uintptr_t dayOfWeek, QTextCharFormat* format);
QTextCharFormat* QCalendarWidget_DateTextFormatWithDate(const QCalendarWidget* self, QDate* date);
void QCalendarWidget_SetDateTextFormat(QCalendarWidget* self, QDate* date, QTextCharFormat* format);
bool QCalendarWidget_IsDateEditEnabled(const QCalendarWidget* self);
void QCalendarWidget_SetDateEditEnabled(QCalendarWidget* self, bool enable);
int QCalendarWidget_DateEditAcceptDelay(const QCalendarWidget* self);
void QCalendarWidget_SetDateEditAcceptDelay(QCalendarWidget* self, int delay);
void QCalendarWidget_SetSelectedDate(QCalendarWidget* self, QDate* date);
void QCalendarWidget_SetDateRange(QCalendarWidget* self, QDate* min, QDate* max);
void QCalendarWidget_SetCurrentPage(QCalendarWidget* self, int year, int month);
void QCalendarWidget_SetGridVisible(QCalendarWidget* self, bool show);
void QCalendarWidget_SetNavigationBarVisible(QCalendarWidget* self, bool visible);
void QCalendarWidget_ShowNextMonth(QCalendarWidget* self);
void QCalendarWidget_ShowPreviousMonth(QCalendarWidget* self);
void QCalendarWidget_ShowNextYear(QCalendarWidget* self);
void QCalendarWidget_ShowPreviousYear(QCalendarWidget* self);
void QCalendarWidget_ShowSelectedDate(QCalendarWidget* self);
void QCalendarWidget_ShowToday(QCalendarWidget* self);
void QCalendarWidget_SelectionChanged(QCalendarWidget* self);
void QCalendarWidget_connect_SelectionChanged(QCalendarWidget* self, void* slot);
void QCalendarWidget_Clicked(QCalendarWidget* self, QDate* date);
void QCalendarWidget_connect_Clicked(QCalendarWidget* self, void* slot);
void QCalendarWidget_Activated(QCalendarWidget* self, QDate* date);
void QCalendarWidget_connect_Activated(QCalendarWidget* self, void* slot);
void QCalendarWidget_CurrentPageChanged(QCalendarWidget* self, int year, int month);
void QCalendarWidget_connect_CurrentPageChanged(QCalendarWidget* self, void* slot);
struct miqt_string* QCalendarWidget_Tr2(const char* s, const char* c);
struct miqt_string* QCalendarWidget_Tr3(const char* s, const char* c, int n);
struct miqt_string* QCalendarWidget_TrUtf82(const char* s, const char* c);
struct miqt_string* QCalendarWidget_TrUtf83(const char* s, const char* c, int n);
void QCalendarWidget_Delete(QCalendarWidget* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
