#ifndef GEN_QLISTVIEW_H
#define GEN_QLISTVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QListView;
class QMetaObject;
class QModelIndex;
class QPoint;
class QRect;
class QSize;
class QWidget;
#else
typedef struct QListView QListView;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
#endif

QListView* QListView_new();
QListView* QListView_new2(QWidget* parent);
QMetaObject* QListView_MetaObject(QListView* self);
void QListView_Tr(const char* s, char** _out, int* _out_Strlen);
void QListView_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QListView_SetMovement(QListView* self, uintptr_t movement);
uintptr_t QListView_Movement(QListView* self);
void QListView_SetFlow(QListView* self, uintptr_t flow);
uintptr_t QListView_Flow(QListView* self);
void QListView_SetWrapping(QListView* self, bool enable);
bool QListView_IsWrapping(QListView* self);
void QListView_SetResizeMode(QListView* self, uintptr_t mode);
uintptr_t QListView_ResizeMode(QListView* self);
void QListView_SetLayoutMode(QListView* self, uintptr_t mode);
uintptr_t QListView_LayoutMode(QListView* self);
void QListView_SetSpacing(QListView* self, int space);
int QListView_Spacing(QListView* self);
void QListView_SetBatchSize(QListView* self, int batchSize);
int QListView_BatchSize(QListView* self);
void QListView_SetGridSize(QListView* self, QSize* size);
QSize* QListView_GridSize(QListView* self);
void QListView_SetViewMode(QListView* self, uintptr_t mode);
uintptr_t QListView_ViewMode(QListView* self);
void QListView_ClearPropertyFlags(QListView* self);
bool QListView_IsRowHidden(QListView* self, int row);
void QListView_SetRowHidden(QListView* self, int row, bool hide);
void QListView_SetModelColumn(QListView* self, int column);
int QListView_ModelColumn(QListView* self);
void QListView_SetUniformItemSizes(QListView* self, bool enable);
bool QListView_UniformItemSizes(QListView* self);
void QListView_SetWordWrap(QListView* self, bool on);
bool QListView_WordWrap(QListView* self);
void QListView_SetSelectionRectVisible(QListView* self, bool show);
bool QListView_IsSelectionRectVisible(QListView* self);
void QListView_SetItemAlignment(QListView* self, int alignment);
int QListView_ItemAlignment(QListView* self);
QRect* QListView_VisualRect(QListView* self, QModelIndex* index);
void QListView_ScrollTo(QListView* self, QModelIndex* index);
QModelIndex* QListView_IndexAt(QListView* self, QPoint* p);
void QListView_DoItemsLayout(QListView* self);
void QListView_Reset(QListView* self);
void QListView_SetRootIndex(QListView* self, QModelIndex* index);
void QListView_IndexesMoved(QListView* self, QModelIndex** indexes, size_t indexes_len);
void QListView_connect_IndexesMoved(QListView* self, void* slot);
void QListView_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QListView_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QListView_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QListView_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QListView_ScrollTo2(QListView* self, QModelIndex* index, uintptr_t hint);
void QListView_Delete(QListView* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
