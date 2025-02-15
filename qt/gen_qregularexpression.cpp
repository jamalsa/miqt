#include <QList>
#include <QRegularExpression>
#include <QRegularExpressionMatch>
#include <QRegularExpressionMatchIterator>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qregularexpression.h>
#include "gen_qregularexpression.h"
#include "_cgo_export.h"

void QRegularExpression_new(QRegularExpression** outptr_QRegularExpression) {
	QRegularExpression* ret = new QRegularExpression();
	*outptr_QRegularExpression = ret;
}

void QRegularExpression_new2(struct miqt_string pattern, QRegularExpression** outptr_QRegularExpression) {
	QString pattern_QString = QString::fromUtf8(pattern.data, pattern.len);
	QRegularExpression* ret = new QRegularExpression(pattern_QString);
	*outptr_QRegularExpression = ret;
}

void QRegularExpression_new3(QRegularExpression* re, QRegularExpression** outptr_QRegularExpression) {
	QRegularExpression* ret = new QRegularExpression(*re);
	*outptr_QRegularExpression = ret;
}

void QRegularExpression_new4(struct miqt_string pattern, int options, QRegularExpression** outptr_QRegularExpression) {
	QString pattern_QString = QString::fromUtf8(pattern.data, pattern.len);
	QRegularExpression* ret = new QRegularExpression(pattern_QString, static_cast<QRegularExpression::PatternOptions>(options));
	*outptr_QRegularExpression = ret;
}

int QRegularExpression_PatternOptions(const QRegularExpression* self) {
	QRegularExpression::PatternOptions _ret = self->patternOptions();
	return static_cast<int>(_ret);
}

void QRegularExpression_SetPatternOptions(QRegularExpression* self, int options) {
	self->setPatternOptions(static_cast<QRegularExpression::PatternOptions>(options));
}

void QRegularExpression_OperatorAssign(QRegularExpression* self, QRegularExpression* re) {
	self->operator=(*re);
}

void QRegularExpression_Swap(QRegularExpression* self, QRegularExpression* other) {
	self->swap(*other);
}

struct miqt_string QRegularExpression_Pattern(const QRegularExpression* self) {
	QString _ret = self->pattern();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QRegularExpression_SetPattern(QRegularExpression* self, struct miqt_string pattern) {
	QString pattern_QString = QString::fromUtf8(pattern.data, pattern.len);
	self->setPattern(pattern_QString);
}

bool QRegularExpression_IsValid(const QRegularExpression* self) {
	return self->isValid();
}

int QRegularExpression_PatternErrorOffset(const QRegularExpression* self) {
	return self->patternErrorOffset();
}

struct miqt_string QRegularExpression_ErrorString(const QRegularExpression* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QRegularExpression_CaptureCount(const QRegularExpression* self) {
	return self->captureCount();
}

struct miqt_array /* of struct miqt_string */  QRegularExpression_NamedCaptureGroups(const QRegularExpression* self) {
	QStringList _ret = self->namedCaptureGroups();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QRegularExpressionMatch* QRegularExpression_Match(const QRegularExpression* self, struct miqt_string subject) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatch(self->match(subject_QString));
}

QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch(const QRegularExpression* self, struct miqt_string subject) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatchIterator(self->globalMatch(subject_QString));
}

void QRegularExpression_Optimize(const QRegularExpression* self) {
	self->optimize();
}

struct miqt_string QRegularExpression_Escape(struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QString _ret = QRegularExpression::escape(str_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegularExpression_WildcardToRegularExpression(struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QString _ret = QRegularExpression::wildcardToRegularExpression(str_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegularExpression_AnchoredPattern(struct miqt_string expression) {
	QString expression_QString = QString::fromUtf8(expression.data, expression.len);
	QString _ret = QRegularExpression::anchoredPattern(expression_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QRegularExpression_OperatorEqual(const QRegularExpression* self, QRegularExpression* re) {
	return self->operator==(*re);
}

bool QRegularExpression_OperatorNotEqual(const QRegularExpression* self, QRegularExpression* re) {
	return self->operator!=(*re);
}

QRegularExpressionMatch* QRegularExpression_Match2(const QRegularExpression* self, struct miqt_string subject, int offset) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatch(self->match(subject_QString, static_cast<int>(offset)));
}

QRegularExpressionMatch* QRegularExpression_Match3(const QRegularExpression* self, struct miqt_string subject, int offset, int matchType) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatch(self->match(subject_QString, static_cast<int>(offset), static_cast<QRegularExpression::MatchType>(matchType)));
}

QRegularExpressionMatch* QRegularExpression_Match4(const QRegularExpression* self, struct miqt_string subject, int offset, int matchType, int matchOptions) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatch(self->match(subject_QString, static_cast<int>(offset), static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOptions>(matchOptions)));
}

QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch2(const QRegularExpression* self, struct miqt_string subject, int offset) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatchIterator(self->globalMatch(subject_QString, static_cast<int>(offset)));
}

QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch3(const QRegularExpression* self, struct miqt_string subject, int offset, int matchType) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatchIterator(self->globalMatch(subject_QString, static_cast<int>(offset), static_cast<QRegularExpression::MatchType>(matchType)));
}

QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch4(const QRegularExpression* self, struct miqt_string subject, int offset, int matchType, int matchOptions) {
	QString subject_QString = QString::fromUtf8(subject.data, subject.len);
	return new QRegularExpressionMatchIterator(self->globalMatch(subject_QString, static_cast<int>(offset), static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOptions>(matchOptions)));
}

void QRegularExpression_Delete(QRegularExpression* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QRegularExpression*>( self );
	} else {
		delete self;
	}
}

void QRegularExpressionMatch_new(QRegularExpressionMatch** outptr_QRegularExpressionMatch) {
	QRegularExpressionMatch* ret = new QRegularExpressionMatch();
	*outptr_QRegularExpressionMatch = ret;
}

void QRegularExpressionMatch_new2(QRegularExpressionMatch* match, QRegularExpressionMatch** outptr_QRegularExpressionMatch) {
	QRegularExpressionMatch* ret = new QRegularExpressionMatch(*match);
	*outptr_QRegularExpressionMatch = ret;
}

void QRegularExpressionMatch_OperatorAssign(QRegularExpressionMatch* self, QRegularExpressionMatch* match) {
	self->operator=(*match);
}

void QRegularExpressionMatch_Swap(QRegularExpressionMatch* self, QRegularExpressionMatch* other) {
	self->swap(*other);
}

QRegularExpression* QRegularExpressionMatch_RegularExpression(const QRegularExpressionMatch* self) {
	return new QRegularExpression(self->regularExpression());
}

int QRegularExpressionMatch_MatchType(const QRegularExpressionMatch* self) {
	QRegularExpression::MatchType _ret = self->matchType();
	return static_cast<int>(_ret);
}

int QRegularExpressionMatch_MatchOptions(const QRegularExpressionMatch* self) {
	QRegularExpression::MatchOptions _ret = self->matchOptions();
	return static_cast<int>(_ret);
}

bool QRegularExpressionMatch_HasMatch(const QRegularExpressionMatch* self) {
	return self->hasMatch();
}

bool QRegularExpressionMatch_HasPartialMatch(const QRegularExpressionMatch* self) {
	return self->hasPartialMatch();
}

bool QRegularExpressionMatch_IsValid(const QRegularExpressionMatch* self) {
	return self->isValid();
}

int QRegularExpressionMatch_LastCapturedIndex(const QRegularExpressionMatch* self) {
	return self->lastCapturedIndex();
}

struct miqt_string QRegularExpressionMatch_Captured(const QRegularExpressionMatch* self) {
	QString _ret = self->captured();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRegularExpressionMatch_CapturedWithName(const QRegularExpressionMatch* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString _ret = self->captured(name_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of struct miqt_string */  QRegularExpressionMatch_CapturedTexts(const QRegularExpressionMatch* self) {
	QStringList _ret = self->capturedTexts();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

int QRegularExpressionMatch_CapturedStart(const QRegularExpressionMatch* self) {
	return self->capturedStart();
}

int QRegularExpressionMatch_CapturedLength(const QRegularExpressionMatch* self) {
	return self->capturedLength();
}

int QRegularExpressionMatch_CapturedEnd(const QRegularExpressionMatch* self) {
	return self->capturedEnd();
}

int QRegularExpressionMatch_CapturedStartWithName(const QRegularExpressionMatch* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return self->capturedStart(name_QString);
}

int QRegularExpressionMatch_CapturedLengthWithName(const QRegularExpressionMatch* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return self->capturedLength(name_QString);
}

int QRegularExpressionMatch_CapturedEndWithName(const QRegularExpressionMatch* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return self->capturedEnd(name_QString);
}

struct miqt_string QRegularExpressionMatch_Captured1(const QRegularExpressionMatch* self, int nth) {
	QString _ret = self->captured(static_cast<int>(nth));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QRegularExpressionMatch_CapturedStart1(const QRegularExpressionMatch* self, int nth) {
	return self->capturedStart(static_cast<int>(nth));
}

int QRegularExpressionMatch_CapturedLength1(const QRegularExpressionMatch* self, int nth) {
	return self->capturedLength(static_cast<int>(nth));
}

int QRegularExpressionMatch_CapturedEnd1(const QRegularExpressionMatch* self, int nth) {
	return self->capturedEnd(static_cast<int>(nth));
}

void QRegularExpressionMatch_Delete(QRegularExpressionMatch* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QRegularExpressionMatch*>( self );
	} else {
		delete self;
	}
}

void QRegularExpressionMatchIterator_new(QRegularExpressionMatchIterator** outptr_QRegularExpressionMatchIterator) {
	QRegularExpressionMatchIterator* ret = new QRegularExpressionMatchIterator();
	*outptr_QRegularExpressionMatchIterator = ret;
}

void QRegularExpressionMatchIterator_new2(QRegularExpressionMatchIterator* iterator, QRegularExpressionMatchIterator** outptr_QRegularExpressionMatchIterator) {
	QRegularExpressionMatchIterator* ret = new QRegularExpressionMatchIterator(*iterator);
	*outptr_QRegularExpressionMatchIterator = ret;
}

void QRegularExpressionMatchIterator_OperatorAssign(QRegularExpressionMatchIterator* self, QRegularExpressionMatchIterator* iterator) {
	self->operator=(*iterator);
}

void QRegularExpressionMatchIterator_Swap(QRegularExpressionMatchIterator* self, QRegularExpressionMatchIterator* other) {
	self->swap(*other);
}

bool QRegularExpressionMatchIterator_IsValid(const QRegularExpressionMatchIterator* self) {
	return self->isValid();
}

bool QRegularExpressionMatchIterator_HasNext(const QRegularExpressionMatchIterator* self) {
	return self->hasNext();
}

QRegularExpressionMatch* QRegularExpressionMatchIterator_Next(QRegularExpressionMatchIterator* self) {
	return new QRegularExpressionMatch(self->next());
}

QRegularExpressionMatch* QRegularExpressionMatchIterator_PeekNext(const QRegularExpressionMatchIterator* self) {
	return new QRegularExpressionMatch(self->peekNext());
}

QRegularExpression* QRegularExpressionMatchIterator_RegularExpression(const QRegularExpressionMatchIterator* self) {
	return new QRegularExpression(self->regularExpression());
}

int QRegularExpressionMatchIterator_MatchType(const QRegularExpressionMatchIterator* self) {
	QRegularExpression::MatchType _ret = self->matchType();
	return static_cast<int>(_ret);
}

int QRegularExpressionMatchIterator_MatchOptions(const QRegularExpressionMatchIterator* self) {
	QRegularExpression::MatchOptions _ret = self->matchOptions();
	return static_cast<int>(_ret);
}

void QRegularExpressionMatchIterator_Delete(QRegularExpressionMatchIterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QRegularExpressionMatchIterator*>( self );
	} else {
		delete self;
	}
}

