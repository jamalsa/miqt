#pragma once
#ifndef MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA_GEN_QSCILEXERBASH_H
#define MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA_GEN_QSCILEXERBASH_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QColor;
class QFont;
class QMetaObject;
class QObject;
class QSettings;
class QsciLexer;
class QsciLexerBash;
class QsciScintilla;
#else
typedef struct QColor QColor;
typedef struct QFont QFont;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSettings QSettings;
typedef struct QsciLexer QsciLexer;
typedef struct QsciLexerBash QsciLexerBash;
typedef struct QsciScintilla QsciScintilla;
#endif

void QsciLexerBash_new(QsciLexerBash** outptr_QsciLexerBash, QsciLexer** outptr_QsciLexer, QObject** outptr_QObject);
void QsciLexerBash_new2(QObject* parent, QsciLexerBash** outptr_QsciLexerBash, QsciLexer** outptr_QsciLexer, QObject** outptr_QObject);
QMetaObject* QsciLexerBash_MetaObject(const QsciLexerBash* self);
void* QsciLexerBash_Metacast(QsciLexerBash* self, const char* param1);
struct miqt_string QsciLexerBash_Tr(const char* s);
struct miqt_string QsciLexerBash_TrUtf8(const char* s);
const char* QsciLexerBash_Language(const QsciLexerBash* self);
const char* QsciLexerBash_Lexer(const QsciLexerBash* self);
int QsciLexerBash_BraceStyle(const QsciLexerBash* self);
const char* QsciLexerBash_WordCharacters(const QsciLexerBash* self);
QColor* QsciLexerBash_DefaultColor(const QsciLexerBash* self, int style);
bool QsciLexerBash_DefaultEolFill(const QsciLexerBash* self, int style);
QFont* QsciLexerBash_DefaultFont(const QsciLexerBash* self, int style);
QColor* QsciLexerBash_DefaultPaper(const QsciLexerBash* self, int style);
const char* QsciLexerBash_Keywords(const QsciLexerBash* self, int set);
struct miqt_string QsciLexerBash_Description(const QsciLexerBash* self, int style);
void QsciLexerBash_RefreshProperties(QsciLexerBash* self);
bool QsciLexerBash_FoldComments(const QsciLexerBash* self);
bool QsciLexerBash_FoldCompact(const QsciLexerBash* self);
void QsciLexerBash_SetFoldComments(QsciLexerBash* self, bool fold);
void QsciLexerBash_SetFoldCompact(QsciLexerBash* self, bool fold);
struct miqt_string QsciLexerBash_Tr2(const char* s, const char* c);
struct miqt_string QsciLexerBash_Tr3(const char* s, const char* c, int n);
struct miqt_string QsciLexerBash_TrUtf82(const char* s, const char* c);
struct miqt_string QsciLexerBash_TrUtf83(const char* s, const char* c, int n);
void QsciLexerBash_override_virtual_SetFoldComments(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetFoldComments(void* self, bool fold);
void QsciLexerBash_override_virtual_SetFoldCompact(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetFoldCompact(void* self, bool fold);
void QsciLexerBash_override_virtual_Language(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_Language(const void* self);
void QsciLexerBash_override_virtual_Lexer(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_Lexer(const void* self);
void QsciLexerBash_override_virtual_LexerId(void* self, intptr_t slot);
int QsciLexerBash_virtualbase_LexerId(const void* self);
void QsciLexerBash_override_virtual_AutoCompletionFillups(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_AutoCompletionFillups(const void* self);
void QsciLexerBash_override_virtual_AutoCompletionWordSeparators(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QsciLexerBash_virtualbase_AutoCompletionWordSeparators(const void* self);
void QsciLexerBash_override_virtual_BlockEnd(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_BlockEnd(const void* self, int* style);
void QsciLexerBash_override_virtual_BlockLookback(void* self, intptr_t slot);
int QsciLexerBash_virtualbase_BlockLookback(const void* self);
void QsciLexerBash_override_virtual_BlockStart(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_BlockStart(const void* self, int* style);
void QsciLexerBash_override_virtual_BlockStartKeyword(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_BlockStartKeyword(const void* self, int* style);
void QsciLexerBash_override_virtual_BraceStyle(void* self, intptr_t slot);
int QsciLexerBash_virtualbase_BraceStyle(const void* self);
void QsciLexerBash_override_virtual_CaseSensitive(void* self, intptr_t slot);
bool QsciLexerBash_virtualbase_CaseSensitive(const void* self);
void QsciLexerBash_override_virtual_Color(void* self, intptr_t slot);
QColor* QsciLexerBash_virtualbase_Color(const void* self, int style);
void QsciLexerBash_override_virtual_EolFill(void* self, intptr_t slot);
bool QsciLexerBash_virtualbase_EolFill(const void* self, int style);
void QsciLexerBash_override_virtual_Font(void* self, intptr_t slot);
QFont* QsciLexerBash_virtualbase_Font(const void* self, int style);
void QsciLexerBash_override_virtual_IndentationGuideView(void* self, intptr_t slot);
int QsciLexerBash_virtualbase_IndentationGuideView(const void* self);
void QsciLexerBash_override_virtual_Keywords(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_Keywords(const void* self, int set);
void QsciLexerBash_override_virtual_DefaultStyle(void* self, intptr_t slot);
int QsciLexerBash_virtualbase_DefaultStyle(const void* self);
void QsciLexerBash_override_virtual_Description(void* self, intptr_t slot);
struct miqt_string QsciLexerBash_virtualbase_Description(const void* self, int style);
void QsciLexerBash_override_virtual_Paper(void* self, intptr_t slot);
QColor* QsciLexerBash_virtualbase_Paper(const void* self, int style);
void QsciLexerBash_override_virtual_DefaultColorWithStyle(void* self, intptr_t slot);
QColor* QsciLexerBash_virtualbase_DefaultColorWithStyle(const void* self, int style);
void QsciLexerBash_override_virtual_DefaultEolFill(void* self, intptr_t slot);
bool QsciLexerBash_virtualbase_DefaultEolFill(const void* self, int style);
void QsciLexerBash_override_virtual_DefaultFontWithStyle(void* self, intptr_t slot);
QFont* QsciLexerBash_virtualbase_DefaultFontWithStyle(const void* self, int style);
void QsciLexerBash_override_virtual_DefaultPaperWithStyle(void* self, intptr_t slot);
QColor* QsciLexerBash_virtualbase_DefaultPaperWithStyle(const void* self, int style);
void QsciLexerBash_override_virtual_SetEditor(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetEditor(void* self, QsciScintilla* editor);
void QsciLexerBash_override_virtual_RefreshProperties(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_RefreshProperties(void* self);
void QsciLexerBash_override_virtual_StyleBitsNeeded(void* self, intptr_t slot);
int QsciLexerBash_virtualbase_StyleBitsNeeded(const void* self);
void QsciLexerBash_override_virtual_WordCharacters(void* self, intptr_t slot);
const char* QsciLexerBash_virtualbase_WordCharacters(const void* self);
void QsciLexerBash_override_virtual_SetAutoIndentStyle(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetAutoIndentStyle(void* self, int autoindentstyle);
void QsciLexerBash_override_virtual_SetColor(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetColor(void* self, QColor* c, int style);
void QsciLexerBash_override_virtual_SetEolFill(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetEolFill(void* self, bool eoffill, int style);
void QsciLexerBash_override_virtual_SetFont(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetFont(void* self, QFont* f, int style);
void QsciLexerBash_override_virtual_SetPaper(void* self, intptr_t slot);
void QsciLexerBash_virtualbase_SetPaper(void* self, QColor* c, int style);
void QsciLexerBash_override_virtual_ReadProperties(void* self, intptr_t slot);
bool QsciLexerBash_virtualbase_ReadProperties(void* self, QSettings* qs, struct miqt_string prefix);
void QsciLexerBash_override_virtual_WriteProperties(void* self, intptr_t slot);
bool QsciLexerBash_virtualbase_WriteProperties(const void* self, QSettings* qs, struct miqt_string prefix);
void QsciLexerBash_Delete(QsciLexerBash* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
