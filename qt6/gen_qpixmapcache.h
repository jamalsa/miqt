#pragma once
#ifndef MIQT_QT6_GEN_QPIXMAPCACHE_H
#define MIQT_QT6_GEN_QPIXMAPCACHE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPixmap;
class QPixmapCache;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPixmapCache__Key)
typedef QPixmapCache::Key QPixmapCache__Key;
#else
class QPixmapCache__Key;
#endif
#else
typedef struct QPixmap QPixmap;
typedef struct QPixmapCache QPixmapCache;
typedef struct QPixmapCache__Key QPixmapCache__Key;
#endif

int QPixmapCache_CacheLimit();
void QPixmapCache_SetCacheLimit(int cacheLimit);
bool QPixmapCache_Find(struct miqt_string key, QPixmap* pixmap);
bool QPixmapCache_Find2(QPixmapCache__Key* key, QPixmap* pixmap);
bool QPixmapCache_Insert(struct miqt_string key, QPixmap* pixmap);
QPixmapCache__Key* QPixmapCache_InsertWithPixmap(QPixmap* pixmap);
bool QPixmapCache_Replace(QPixmapCache__Key* key, QPixmap* pixmap);
void QPixmapCache_Remove(struct miqt_string key);
void QPixmapCache_RemoveWithKey(QPixmapCache__Key* key);
void QPixmapCache_Clear();
void QPixmapCache_Delete(QPixmapCache* self, bool isSubclass);

void QPixmapCache__Key_new(QPixmapCache__Key** outptr_QPixmapCache__Key);
void QPixmapCache__Key_new2(QPixmapCache__Key* other, QPixmapCache__Key** outptr_QPixmapCache__Key);
bool QPixmapCache__Key_OperatorEqual(const QPixmapCache__Key* self, QPixmapCache__Key* key);
bool QPixmapCache__Key_OperatorNotEqual(const QPixmapCache__Key* self, QPixmapCache__Key* key);
void QPixmapCache__Key_OperatorAssign(QPixmapCache__Key* self, QPixmapCache__Key* other);
void QPixmapCache__Key_Swap(QPixmapCache__Key* self, QPixmapCache__Key* other);
bool QPixmapCache__Key_IsValid(const QPixmapCache__Key* self);
void QPixmapCache__Key_Delete(QPixmapCache__Key* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
