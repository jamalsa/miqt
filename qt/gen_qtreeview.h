#ifndef GEN_QTREEVIEW_H
#define GEN_QTREEVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractItemModel;
class QHeaderView;
class QItemSelectionModel;
class QMetaObject;
class QModelIndex;
class QPoint;
class QRect;
class QTreeView;
class QWidget;
#else
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QHeaderView QHeaderView;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QTreeView QTreeView;
typedef struct QWidget QWidget;
#endif

QTreeView* QTreeView_new();
QTreeView* QTreeView_new2(QWidget* parent);
QMetaObject* QTreeView_MetaObject(QTreeView* self);
void QTreeView_Tr(const char* s, char** _out, int* _out_Strlen);
void QTreeView_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QTreeView_SetModel(QTreeView* self, QAbstractItemModel* model);
void QTreeView_SetRootIndex(QTreeView* self, QModelIndex* index);
void QTreeView_SetSelectionModel(QTreeView* self, QItemSelectionModel* selectionModel);
QHeaderView* QTreeView_Header(QTreeView* self);
void QTreeView_SetHeader(QTreeView* self, QHeaderView* header);
int QTreeView_AutoExpandDelay(QTreeView* self);
void QTreeView_SetAutoExpandDelay(QTreeView* self, int delay);
int QTreeView_Indentation(QTreeView* self);
void QTreeView_SetIndentation(QTreeView* self, int i);
void QTreeView_ResetIndentation(QTreeView* self);
bool QTreeView_RootIsDecorated(QTreeView* self);
void QTreeView_SetRootIsDecorated(QTreeView* self, bool show);
bool QTreeView_UniformRowHeights(QTreeView* self);
void QTreeView_SetUniformRowHeights(QTreeView* self, bool uniform);
bool QTreeView_ItemsExpandable(QTreeView* self);
void QTreeView_SetItemsExpandable(QTreeView* self, bool enable);
bool QTreeView_ExpandsOnDoubleClick(QTreeView* self);
void QTreeView_SetExpandsOnDoubleClick(QTreeView* self, bool enable);
int QTreeView_ColumnViewportPosition(QTreeView* self, int column);
int QTreeView_ColumnWidth(QTreeView* self, int column);
void QTreeView_SetColumnWidth(QTreeView* self, int column, int width);
int QTreeView_ColumnAt(QTreeView* self, int x);
bool QTreeView_IsColumnHidden(QTreeView* self, int column);
void QTreeView_SetColumnHidden(QTreeView* self, int column, bool hide);
bool QTreeView_IsHeaderHidden(QTreeView* self);
void QTreeView_SetHeaderHidden(QTreeView* self, bool hide);
bool QTreeView_IsRowHidden(QTreeView* self, int row, QModelIndex* parent);
void QTreeView_SetRowHidden(QTreeView* self, int row, QModelIndex* parent, bool hide);
bool QTreeView_IsFirstColumnSpanned(QTreeView* self, int row, QModelIndex* parent);
void QTreeView_SetFirstColumnSpanned(QTreeView* self, int row, QModelIndex* parent, bool span);
bool QTreeView_IsExpanded(QTreeView* self, QModelIndex* index);
void QTreeView_SetExpanded(QTreeView* self, QModelIndex* index, bool expand);
void QTreeView_SetSortingEnabled(QTreeView* self, bool enable);
bool QTreeView_IsSortingEnabled(QTreeView* self);
void QTreeView_SetAnimated(QTreeView* self, bool enable);
bool QTreeView_IsAnimated(QTreeView* self);
void QTreeView_SetAllColumnsShowFocus(QTreeView* self, bool enable);
bool QTreeView_AllColumnsShowFocus(QTreeView* self);
void QTreeView_SetWordWrap(QTreeView* self, bool on);
bool QTreeView_WordWrap(QTreeView* self);
void QTreeView_SetTreePosition(QTreeView* self, int logicalIndex);
int QTreeView_TreePosition(QTreeView* self);
void QTreeView_KeyboardSearch(QTreeView* self, const char* search, size_t search_Strlen);
QRect* QTreeView_VisualRect(QTreeView* self, QModelIndex* index);
void QTreeView_ScrollTo(QTreeView* self, QModelIndex* index);
QModelIndex* QTreeView_IndexAt(QTreeView* self, QPoint* p);
QModelIndex* QTreeView_IndexAbove(QTreeView* self, QModelIndex* index);
QModelIndex* QTreeView_IndexBelow(QTreeView* self, QModelIndex* index);
void QTreeView_DoItemsLayout(QTreeView* self);
void QTreeView_Reset(QTreeView* self);
void QTreeView_DataChanged(QTreeView* self, QModelIndex* topLeft, QModelIndex* bottomRight);
void QTreeView_SelectAll(QTreeView* self);
void QTreeView_Expanded(QTreeView* self, QModelIndex* index);
void QTreeView_connect_Expanded(QTreeView* self, void* slot);
void QTreeView_Collapsed(QTreeView* self, QModelIndex* index);
void QTreeView_connect_Collapsed(QTreeView* self, void* slot);
void QTreeView_HideColumn(QTreeView* self, int column);
void QTreeView_ShowColumn(QTreeView* self, int column);
void QTreeView_Expand(QTreeView* self, QModelIndex* index);
void QTreeView_Collapse(QTreeView* self, QModelIndex* index);
void QTreeView_ResizeColumnToContents(QTreeView* self, int column);
void QTreeView_SortByColumn(QTreeView* self, int column);
void QTreeView_SortByColumn2(QTreeView* self, int column, uintptr_t order);
void QTreeView_ExpandAll(QTreeView* self);
void QTreeView_ExpandRecursively(QTreeView* self, QModelIndex* index);
void QTreeView_CollapseAll(QTreeView* self);
void QTreeView_ExpandToDepth(QTreeView* self, int depth);
void QTreeView_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QTreeView_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QTreeView_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QTreeView_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QTreeView_ScrollTo2(QTreeView* self, QModelIndex* index, uintptr_t hint);
void QTreeView_DataChanged3(QTreeView* self, QModelIndex* topLeft, QModelIndex* bottomRight, int* roles, size_t roles_len);
void QTreeView_ExpandRecursively2(QTreeView* self, QModelIndex* index, int depth);
void QTreeView_Delete(QTreeView* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
