#include <QColor>
#include <QFont>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QSettings>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qscilexerpov.h>
#include "gen_qscilexerpov.h"
#include "_cgo_export.h"

class MiqtVirtualQsciLexerPOV : public virtual QsciLexerPOV {
public:

	MiqtVirtualQsciLexerPOV(): QsciLexerPOV() {};
	MiqtVirtualQsciLexerPOV(QObject* parent): QsciLexerPOV(parent) {};

	virtual ~MiqtVirtualQsciLexerPOV() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetFoldComments = 0;

	// Subclass to allow providing a Go implementation
	virtual void setFoldComments(bool fold) override {
		if (handle__SetFoldComments == 0) {
			QsciLexerPOV::setFoldComments(fold);
			return;
		}
		
		bool sigval1 = fold;

		miqt_exec_callback_QsciLexerPOV_SetFoldComments(this, handle__SetFoldComments, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetFoldComments(bool fold) {

		QsciLexerPOV::setFoldComments(fold);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetFoldCompact = 0;

	// Subclass to allow providing a Go implementation
	virtual void setFoldCompact(bool fold) override {
		if (handle__SetFoldCompact == 0) {
			QsciLexerPOV::setFoldCompact(fold);
			return;
		}
		
		bool sigval1 = fold;

		miqt_exec_callback_QsciLexerPOV_SetFoldCompact(this, handle__SetFoldCompact, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetFoldCompact(bool fold) {

		QsciLexerPOV::setFoldCompact(fold);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetFoldDirectives = 0;

	// Subclass to allow providing a Go implementation
	virtual void setFoldDirectives(bool fold) override {
		if (handle__SetFoldDirectives == 0) {
			QsciLexerPOV::setFoldDirectives(fold);
			return;
		}
		
		bool sigval1 = fold;

		miqt_exec_callback_QsciLexerPOV_SetFoldDirectives(this, handle__SetFoldDirectives, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetFoldDirectives(bool fold) {

		QsciLexerPOV::setFoldDirectives(fold);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Language = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* language() const override {
		if (handle__Language == 0) {
			return nullptr; // Pure virtual, there is no base we can call
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_Language(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__Language);

		return callback_return_value;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Lexer = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* lexer() const override {
		if (handle__Lexer == 0) {
			return QsciLexerPOV::lexer();
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_Lexer(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__Lexer);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_Lexer() const {

		return (const char*) QsciLexerPOV::lexer();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__LexerId = 0;

	// Subclass to allow providing a Go implementation
	virtual int lexerId() const override {
		if (handle__LexerId == 0) {
			return QsciLexerPOV::lexerId();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerPOV_LexerId(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__LexerId);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_LexerId() const {

		return QsciLexerPOV::lexerId();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__AutoCompletionFillups = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* autoCompletionFillups() const override {
		if (handle__AutoCompletionFillups == 0) {
			return QsciLexerPOV::autoCompletionFillups();
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_AutoCompletionFillups(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__AutoCompletionFillups);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_AutoCompletionFillups() const {

		return (const char*) QsciLexerPOV::autoCompletionFillups();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__AutoCompletionWordSeparators = 0;

	// Subclass to allow providing a Go implementation
	virtual QStringList autoCompletionWordSeparators() const override {
		if (handle__AutoCompletionWordSeparators == 0) {
			return QsciLexerPOV::autoCompletionWordSeparators();
		}
		

		struct miqt_array /* of struct miqt_string */  callback_return_value = miqt_exec_callback_QsciLexerPOV_AutoCompletionWordSeparators(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__AutoCompletionWordSeparators);
		QStringList callback_return_value_QList;
		callback_return_value_QList.reserve(callback_return_value.len);
		struct miqt_string* callback_return_value_arr = static_cast<struct miqt_string*>(callback_return_value.data);
		for(size_t i = 0; i < callback_return_value.len; ++i) {
			QString callback_return_value_arr_i_QString = QString::fromUtf8(callback_return_value_arr[i].data, callback_return_value_arr[i].len);
			callback_return_value_QList.push_back(callback_return_value_arr_i_QString);
		}

		return callback_return_value_QList;
	}

	// Wrapper to allow calling protected method
	struct miqt_array /* of struct miqt_string */  virtualbase_AutoCompletionWordSeparators() const {

		QStringList _ret = QsciLexerPOV::autoCompletionWordSeparators();
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

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockEnd = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* blockEnd(int* style) const override {
		if (handle__BlockEnd == 0) {
			return QsciLexerPOV::blockEnd(style);
		}
		
		int* sigval1 = style;

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_BlockEnd(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__BlockEnd, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_BlockEnd(int* style) const {

		return (const char*) QsciLexerPOV::blockEnd(static_cast<int*>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockLookback = 0;

	// Subclass to allow providing a Go implementation
	virtual int blockLookback() const override {
		if (handle__BlockLookback == 0) {
			return QsciLexerPOV::blockLookback();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerPOV_BlockLookback(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__BlockLookback);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_BlockLookback() const {

		return QsciLexerPOV::blockLookback();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockStart = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* blockStart(int* style) const override {
		if (handle__BlockStart == 0) {
			return QsciLexerPOV::blockStart(style);
		}
		
		int* sigval1 = style;

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_BlockStart(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__BlockStart, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_BlockStart(int* style) const {

		return (const char*) QsciLexerPOV::blockStart(static_cast<int*>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockStartKeyword = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* blockStartKeyword(int* style) const override {
		if (handle__BlockStartKeyword == 0) {
			return QsciLexerPOV::blockStartKeyword(style);
		}
		
		int* sigval1 = style;

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_BlockStartKeyword(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__BlockStartKeyword, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_BlockStartKeyword(int* style) const {

		return (const char*) QsciLexerPOV::blockStartKeyword(static_cast<int*>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BraceStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual int braceStyle() const override {
		if (handle__BraceStyle == 0) {
			return QsciLexerPOV::braceStyle();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerPOV_BraceStyle(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__BraceStyle);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_BraceStyle() const {

		return QsciLexerPOV::braceStyle();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CaseSensitive = 0;

	// Subclass to allow providing a Go implementation
	virtual bool caseSensitive() const override {
		if (handle__CaseSensitive == 0) {
			return QsciLexerPOV::caseSensitive();
		}
		

		bool callback_return_value = miqt_exec_callback_QsciLexerPOV_CaseSensitive(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__CaseSensitive);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_CaseSensitive() const {

		return QsciLexerPOV::caseSensitive();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Color = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor color(int style) const override {
		if (handle__Color == 0) {
			return QsciLexerPOV::color(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerPOV_Color(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__Color, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_Color(int style) const {

		return new QColor(QsciLexerPOV::color(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EolFill = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eolFill(int style) const override {
		if (handle__EolFill == 0) {
			return QsciLexerPOV::eolFill(style);
		}
		
		int sigval1 = style;

		bool callback_return_value = miqt_exec_callback_QsciLexerPOV_EolFill(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__EolFill, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EolFill(int style) const {

		return QsciLexerPOV::eolFill(static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Font = 0;

	// Subclass to allow providing a Go implementation
	virtual QFont font(int style) const override {
		if (handle__Font == 0) {
			return QsciLexerPOV::font(style);
		}
		
		int sigval1 = style;

		QFont* callback_return_value = miqt_exec_callback_QsciLexerPOV_Font(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__Font, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QFont* virtualbase_Font(int style) const {

		return new QFont(QsciLexerPOV::font(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IndentationGuideView = 0;

	// Subclass to allow providing a Go implementation
	virtual int indentationGuideView() const override {
		if (handle__IndentationGuideView == 0) {
			return QsciLexerPOV::indentationGuideView();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerPOV_IndentationGuideView(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__IndentationGuideView);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_IndentationGuideView() const {

		return QsciLexerPOV::indentationGuideView();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Keywords = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* keywords(int set) const override {
		if (handle__Keywords == 0) {
			return QsciLexerPOV::keywords(set);
		}
		
		int sigval1 = set;

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_Keywords(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__Keywords, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_Keywords(int set) const {

		return (const char*) QsciLexerPOV::keywords(static_cast<int>(set));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual int defaultStyle() const override {
		if (handle__DefaultStyle == 0) {
			return QsciLexerPOV::defaultStyle();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerPOV_DefaultStyle(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__DefaultStyle);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_DefaultStyle() const {

		return QsciLexerPOV::defaultStyle();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Description = 0;

	// Subclass to allow providing a Go implementation
	virtual QString description(int style) const override {
		if (handle__Description == 0) {
			return QString(); // Pure virtual, there is no base we can call
		}
		
		int sigval1 = style;

		struct miqt_string callback_return_value = miqt_exec_callback_QsciLexerPOV_Description(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__Description, sigval1);
		QString callback_return_value_QString = QString::fromUtf8(callback_return_value.data, callback_return_value.len);

		return callback_return_value_QString;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Paper = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor paper(int style) const override {
		if (handle__Paper == 0) {
			return QsciLexerPOV::paper(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerPOV_Paper(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__Paper, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_Paper(int style) const {

		return new QColor(QsciLexerPOV::paper(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultColorWithStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor defaultColor(int style) const override {
		if (handle__DefaultColorWithStyle == 0) {
			return QsciLexerPOV::defaultColor(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerPOV_DefaultColorWithStyle(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__DefaultColorWithStyle, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_DefaultColorWithStyle(int style) const {

		return new QColor(QsciLexerPOV::defaultColor(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultEolFill = 0;

	// Subclass to allow providing a Go implementation
	virtual bool defaultEolFill(int style) const override {
		if (handle__DefaultEolFill == 0) {
			return QsciLexerPOV::defaultEolFill(style);
		}
		
		int sigval1 = style;

		bool callback_return_value = miqt_exec_callback_QsciLexerPOV_DefaultEolFill(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__DefaultEolFill, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_DefaultEolFill(int style) const {

		return QsciLexerPOV::defaultEolFill(static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultFontWithStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual QFont defaultFont(int style) const override {
		if (handle__DefaultFontWithStyle == 0) {
			return QsciLexerPOV::defaultFont(style);
		}
		
		int sigval1 = style;

		QFont* callback_return_value = miqt_exec_callback_QsciLexerPOV_DefaultFontWithStyle(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__DefaultFontWithStyle, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QFont* virtualbase_DefaultFontWithStyle(int style) const {

		return new QFont(QsciLexerPOV::defaultFont(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultPaperWithStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor defaultPaper(int style) const override {
		if (handle__DefaultPaperWithStyle == 0) {
			return QsciLexerPOV::defaultPaper(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerPOV_DefaultPaperWithStyle(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__DefaultPaperWithStyle, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_DefaultPaperWithStyle(int style) const {

		return new QColor(QsciLexerPOV::defaultPaper(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetEditor = 0;

	// Subclass to allow providing a Go implementation
	virtual void setEditor(QsciScintilla* editor) override {
		if (handle__SetEditor == 0) {
			QsciLexerPOV::setEditor(editor);
			return;
		}
		
		QsciScintilla* sigval1 = editor;

		miqt_exec_callback_QsciLexerPOV_SetEditor(this, handle__SetEditor, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetEditor(QsciScintilla* editor) {

		QsciLexerPOV::setEditor(editor);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__RefreshProperties = 0;

	// Subclass to allow providing a Go implementation
	virtual void refreshProperties() override {
		if (handle__RefreshProperties == 0) {
			QsciLexerPOV::refreshProperties();
			return;
		}
		

		miqt_exec_callback_QsciLexerPOV_RefreshProperties(this, handle__RefreshProperties);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_RefreshProperties() {

		QsciLexerPOV::refreshProperties();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__StyleBitsNeeded = 0;

	// Subclass to allow providing a Go implementation
	virtual int styleBitsNeeded() const override {
		if (handle__StyleBitsNeeded == 0) {
			return QsciLexerPOV::styleBitsNeeded();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerPOV_StyleBitsNeeded(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__StyleBitsNeeded);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_StyleBitsNeeded() const {

		return QsciLexerPOV::styleBitsNeeded();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WordCharacters = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* wordCharacters() const override {
		if (handle__WordCharacters == 0) {
			return QsciLexerPOV::wordCharacters();
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerPOV_WordCharacters(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__WordCharacters);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_WordCharacters() const {

		return (const char*) QsciLexerPOV::wordCharacters();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetAutoIndentStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual void setAutoIndentStyle(int autoindentstyle) override {
		if (handle__SetAutoIndentStyle == 0) {
			QsciLexerPOV::setAutoIndentStyle(autoindentstyle);
			return;
		}
		
		int sigval1 = autoindentstyle;

		miqt_exec_callback_QsciLexerPOV_SetAutoIndentStyle(this, handle__SetAutoIndentStyle, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetAutoIndentStyle(int autoindentstyle) {

		QsciLexerPOV::setAutoIndentStyle(static_cast<int>(autoindentstyle));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetColor = 0;

	// Subclass to allow providing a Go implementation
	virtual void setColor(const QColor& c, int style) override {
		if (handle__SetColor == 0) {
			QsciLexerPOV::setColor(c, style);
			return;
		}
		
		const QColor& c_ret = c;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&c_ret);
		int sigval2 = style;

		miqt_exec_callback_QsciLexerPOV_SetColor(this, handle__SetColor, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetColor(QColor* c, int style) {

		QsciLexerPOV::setColor(*c, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetEolFill = 0;

	// Subclass to allow providing a Go implementation
	virtual void setEolFill(bool eoffill, int style) override {
		if (handle__SetEolFill == 0) {
			QsciLexerPOV::setEolFill(eoffill, style);
			return;
		}
		
		bool sigval1 = eoffill;
		int sigval2 = style;

		miqt_exec_callback_QsciLexerPOV_SetEolFill(this, handle__SetEolFill, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetEolFill(bool eoffill, int style) {

		QsciLexerPOV::setEolFill(eoffill, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetFont = 0;

	// Subclass to allow providing a Go implementation
	virtual void setFont(const QFont& f, int style) override {
		if (handle__SetFont == 0) {
			QsciLexerPOV::setFont(f, style);
			return;
		}
		
		const QFont& f_ret = f;
		// Cast returned reference into pointer
		QFont* sigval1 = const_cast<QFont*>(&f_ret);
		int sigval2 = style;

		miqt_exec_callback_QsciLexerPOV_SetFont(this, handle__SetFont, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetFont(QFont* f, int style) {

		QsciLexerPOV::setFont(*f, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPaper = 0;

	// Subclass to allow providing a Go implementation
	virtual void setPaper(const QColor& c, int style) override {
		if (handle__SetPaper == 0) {
			QsciLexerPOV::setPaper(c, style);
			return;
		}
		
		const QColor& c_ret = c;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&c_ret);
		int sigval2 = style;

		miqt_exec_callback_QsciLexerPOV_SetPaper(this, handle__SetPaper, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetPaper(QColor* c, int style) {

		QsciLexerPOV::setPaper(*c, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ReadProperties = 0;

	// Subclass to allow providing a Go implementation
	virtual bool readProperties(QSettings& qs, const QString& prefix) override {
		if (handle__ReadProperties == 0) {
			return QsciLexerPOV::readProperties(qs, prefix);
		}
		
		QSettings& qs_ret = qs;
		// Cast returned reference into pointer
		QSettings* sigval1 = &qs_ret;
		const QString prefix_ret = prefix;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray prefix_b = prefix_ret.toUtf8();
		struct miqt_string prefix_ms;
		prefix_ms.len = prefix_b.length();
		prefix_ms.data = static_cast<char*>(malloc(prefix_ms.len));
		memcpy(prefix_ms.data, prefix_b.data(), prefix_ms.len);
		struct miqt_string sigval2 = prefix_ms;

		bool callback_return_value = miqt_exec_callback_QsciLexerPOV_ReadProperties(this, handle__ReadProperties, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_ReadProperties(QSettings* qs, struct miqt_string prefix) {
		QString prefix_QString = QString::fromUtf8(prefix.data, prefix.len);

		return QsciLexerPOV::readProperties(*qs, prefix_QString);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WriteProperties = 0;

	// Subclass to allow providing a Go implementation
	virtual bool writeProperties(QSettings& qs, const QString& prefix) const override {
		if (handle__WriteProperties == 0) {
			return QsciLexerPOV::writeProperties(qs, prefix);
		}
		
		QSettings& qs_ret = qs;
		// Cast returned reference into pointer
		QSettings* sigval1 = &qs_ret;
		const QString prefix_ret = prefix;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray prefix_b = prefix_ret.toUtf8();
		struct miqt_string prefix_ms;
		prefix_ms.len = prefix_b.length();
		prefix_ms.data = static_cast<char*>(malloc(prefix_ms.len));
		memcpy(prefix_ms.data, prefix_b.data(), prefix_ms.len);
		struct miqt_string sigval2 = prefix_ms;

		bool callback_return_value = miqt_exec_callback_QsciLexerPOV_WriteProperties(const_cast<MiqtVirtualQsciLexerPOV*>(this), handle__WriteProperties, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_WriteProperties(QSettings* qs, struct miqt_string prefix) const {
		QString prefix_QString = QString::fromUtf8(prefix.data, prefix.len);

		return QsciLexerPOV::writeProperties(*qs, prefix_QString);

	}

};

void QsciLexerPOV_new(QsciLexerPOV** outptr_QsciLexerPOV, QsciLexer** outptr_QsciLexer, QObject** outptr_QObject) {
	MiqtVirtualQsciLexerPOV* ret = new MiqtVirtualQsciLexerPOV();
	*outptr_QsciLexerPOV = ret;
	*outptr_QsciLexer = static_cast<QsciLexer*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QsciLexerPOV_new2(QObject* parent, QsciLexerPOV** outptr_QsciLexerPOV, QsciLexer** outptr_QsciLexer, QObject** outptr_QObject) {
	MiqtVirtualQsciLexerPOV* ret = new MiqtVirtualQsciLexerPOV(parent);
	*outptr_QsciLexerPOV = ret;
	*outptr_QsciLexer = static_cast<QsciLexer*>(ret);
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QsciLexerPOV_MetaObject(const QsciLexerPOV* self) {
	return (QMetaObject*) self->metaObject();
}

void* QsciLexerPOV_Metacast(QsciLexerPOV* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QsciLexerPOV_Tr(const char* s) {
	QString _ret = QsciLexerPOV::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

const char* QsciLexerPOV_Language(const QsciLexerPOV* self) {
	return (const char*) self->language();
}

const char* QsciLexerPOV_Lexer(const QsciLexerPOV* self) {
	return (const char*) self->lexer();
}

int QsciLexerPOV_BraceStyle(const QsciLexerPOV* self) {
	return self->braceStyle();
}

const char* QsciLexerPOV_WordCharacters(const QsciLexerPOV* self) {
	return (const char*) self->wordCharacters();
}

QColor* QsciLexerPOV_DefaultColor(const QsciLexerPOV* self, int style) {
	return new QColor(self->defaultColor(static_cast<int>(style)));
}

bool QsciLexerPOV_DefaultEolFill(const QsciLexerPOV* self, int style) {
	return self->defaultEolFill(static_cast<int>(style));
}

QFont* QsciLexerPOV_DefaultFont(const QsciLexerPOV* self, int style) {
	return new QFont(self->defaultFont(static_cast<int>(style)));
}

QColor* QsciLexerPOV_DefaultPaper(const QsciLexerPOV* self, int style) {
	return new QColor(self->defaultPaper(static_cast<int>(style)));
}

const char* QsciLexerPOV_Keywords(const QsciLexerPOV* self, int set) {
	return (const char*) self->keywords(static_cast<int>(set));
}

struct miqt_string QsciLexerPOV_Description(const QsciLexerPOV* self, int style) {
	QString _ret = self->description(static_cast<int>(style));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QsciLexerPOV_RefreshProperties(QsciLexerPOV* self) {
	self->refreshProperties();
}

bool QsciLexerPOV_FoldComments(const QsciLexerPOV* self) {
	return self->foldComments();
}

bool QsciLexerPOV_FoldCompact(const QsciLexerPOV* self) {
	return self->foldCompact();
}

bool QsciLexerPOV_FoldDirectives(const QsciLexerPOV* self) {
	return self->foldDirectives();
}

void QsciLexerPOV_SetFoldComments(QsciLexerPOV* self, bool fold) {
	self->setFoldComments(fold);
}

void QsciLexerPOV_SetFoldCompact(QsciLexerPOV* self, bool fold) {
	self->setFoldCompact(fold);
}

void QsciLexerPOV_SetFoldDirectives(QsciLexerPOV* self, bool fold) {
	self->setFoldDirectives(fold);
}

struct miqt_string QsciLexerPOV_Tr2(const char* s, const char* c) {
	QString _ret = QsciLexerPOV::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QsciLexerPOV_Tr3(const char* s, const char* c, int n) {
	QString _ret = QsciLexerPOV::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QsciLexerPOV_override_virtual_SetFoldComments(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetFoldComments = slot;
}

void QsciLexerPOV_virtualbase_SetFoldComments(void* self, bool fold) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetFoldComments(fold);
}

void QsciLexerPOV_override_virtual_SetFoldCompact(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetFoldCompact = slot;
}

void QsciLexerPOV_virtualbase_SetFoldCompact(void* self, bool fold) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetFoldCompact(fold);
}

void QsciLexerPOV_override_virtual_SetFoldDirectives(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetFoldDirectives = slot;
}

void QsciLexerPOV_virtualbase_SetFoldDirectives(void* self, bool fold) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetFoldDirectives(fold);
}

void QsciLexerPOV_override_virtual_Language(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__Language = slot;
}

void QsciLexerPOV_override_virtual_Lexer(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__Lexer = slot;
}

const char* QsciLexerPOV_virtualbase_Lexer(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_Lexer();
}

void QsciLexerPOV_override_virtual_LexerId(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__LexerId = slot;
}

int QsciLexerPOV_virtualbase_LexerId(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_LexerId();
}

void QsciLexerPOV_override_virtual_AutoCompletionFillups(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__AutoCompletionFillups = slot;
}

const char* QsciLexerPOV_virtualbase_AutoCompletionFillups(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_AutoCompletionFillups();
}

void QsciLexerPOV_override_virtual_AutoCompletionWordSeparators(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__AutoCompletionWordSeparators = slot;
}

struct miqt_array /* of struct miqt_string */  QsciLexerPOV_virtualbase_AutoCompletionWordSeparators(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_AutoCompletionWordSeparators();
}

void QsciLexerPOV_override_virtual_BlockEnd(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__BlockEnd = slot;
}

const char* QsciLexerPOV_virtualbase_BlockEnd(const void* self, int* style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_BlockEnd(style);
}

void QsciLexerPOV_override_virtual_BlockLookback(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__BlockLookback = slot;
}

int QsciLexerPOV_virtualbase_BlockLookback(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_BlockLookback();
}

void QsciLexerPOV_override_virtual_BlockStart(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__BlockStart = slot;
}

const char* QsciLexerPOV_virtualbase_BlockStart(const void* self, int* style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_BlockStart(style);
}

void QsciLexerPOV_override_virtual_BlockStartKeyword(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__BlockStartKeyword = slot;
}

const char* QsciLexerPOV_virtualbase_BlockStartKeyword(const void* self, int* style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_BlockStartKeyword(style);
}

void QsciLexerPOV_override_virtual_BraceStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__BraceStyle = slot;
}

int QsciLexerPOV_virtualbase_BraceStyle(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_BraceStyle();
}

void QsciLexerPOV_override_virtual_CaseSensitive(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__CaseSensitive = slot;
}

bool QsciLexerPOV_virtualbase_CaseSensitive(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_CaseSensitive();
}

void QsciLexerPOV_override_virtual_Color(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__Color = slot;
}

QColor* QsciLexerPOV_virtualbase_Color(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_Color(style);
}

void QsciLexerPOV_override_virtual_EolFill(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__EolFill = slot;
}

bool QsciLexerPOV_virtualbase_EolFill(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_EolFill(style);
}

void QsciLexerPOV_override_virtual_Font(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__Font = slot;
}

QFont* QsciLexerPOV_virtualbase_Font(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_Font(style);
}

void QsciLexerPOV_override_virtual_IndentationGuideView(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__IndentationGuideView = slot;
}

int QsciLexerPOV_virtualbase_IndentationGuideView(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_IndentationGuideView();
}

void QsciLexerPOV_override_virtual_Keywords(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__Keywords = slot;
}

const char* QsciLexerPOV_virtualbase_Keywords(const void* self, int set) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_Keywords(set);
}

void QsciLexerPOV_override_virtual_DefaultStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__DefaultStyle = slot;
}

int QsciLexerPOV_virtualbase_DefaultStyle(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_DefaultStyle();
}

void QsciLexerPOV_override_virtual_Description(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__Description = slot;
}

void QsciLexerPOV_override_virtual_Paper(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__Paper = slot;
}

QColor* QsciLexerPOV_virtualbase_Paper(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_Paper(style);
}

void QsciLexerPOV_override_virtual_DefaultColorWithStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__DefaultColorWithStyle = slot;
}

QColor* QsciLexerPOV_virtualbase_DefaultColorWithStyle(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_DefaultColorWithStyle(style);
}

void QsciLexerPOV_override_virtual_DefaultEolFill(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__DefaultEolFill = slot;
}

bool QsciLexerPOV_virtualbase_DefaultEolFill(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_DefaultEolFill(style);
}

void QsciLexerPOV_override_virtual_DefaultFontWithStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__DefaultFontWithStyle = slot;
}

QFont* QsciLexerPOV_virtualbase_DefaultFontWithStyle(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_DefaultFontWithStyle(style);
}

void QsciLexerPOV_override_virtual_DefaultPaperWithStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__DefaultPaperWithStyle = slot;
}

QColor* QsciLexerPOV_virtualbase_DefaultPaperWithStyle(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_DefaultPaperWithStyle(style);
}

void QsciLexerPOV_override_virtual_SetEditor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetEditor = slot;
}

void QsciLexerPOV_virtualbase_SetEditor(void* self, QsciScintilla* editor) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetEditor(editor);
}

void QsciLexerPOV_override_virtual_RefreshProperties(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__RefreshProperties = slot;
}

void QsciLexerPOV_virtualbase_RefreshProperties(void* self) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_RefreshProperties();
}

void QsciLexerPOV_override_virtual_StyleBitsNeeded(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__StyleBitsNeeded = slot;
}

int QsciLexerPOV_virtualbase_StyleBitsNeeded(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_StyleBitsNeeded();
}

void QsciLexerPOV_override_virtual_WordCharacters(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__WordCharacters = slot;
}

const char* QsciLexerPOV_virtualbase_WordCharacters(const void* self) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_WordCharacters();
}

void QsciLexerPOV_override_virtual_SetAutoIndentStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetAutoIndentStyle = slot;
}

void QsciLexerPOV_virtualbase_SetAutoIndentStyle(void* self, int autoindentstyle) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetAutoIndentStyle(autoindentstyle);
}

void QsciLexerPOV_override_virtual_SetColor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetColor = slot;
}

void QsciLexerPOV_virtualbase_SetColor(void* self, QColor* c, int style) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetColor(c, style);
}

void QsciLexerPOV_override_virtual_SetEolFill(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetEolFill = slot;
}

void QsciLexerPOV_virtualbase_SetEolFill(void* self, bool eoffill, int style) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetEolFill(eoffill, style);
}

void QsciLexerPOV_override_virtual_SetFont(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetFont = slot;
}

void QsciLexerPOV_virtualbase_SetFont(void* self, QFont* f, int style) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetFont(f, style);
}

void QsciLexerPOV_override_virtual_SetPaper(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__SetPaper = slot;
}

void QsciLexerPOV_virtualbase_SetPaper(void* self, QColor* c, int style) {
	( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_SetPaper(c, style);
}

void QsciLexerPOV_override_virtual_ReadProperties(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__ReadProperties = slot;
}

bool QsciLexerPOV_virtualbase_ReadProperties(void* self, QSettings* qs, struct miqt_string prefix) {
	return ( (MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_ReadProperties(qs, prefix);
}

void QsciLexerPOV_override_virtual_WriteProperties(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerPOV*>( (QsciLexerPOV*)(self) )->handle__WriteProperties = slot;
}

bool QsciLexerPOV_virtualbase_WriteProperties(const void* self, QSettings* qs, struct miqt_string prefix) {
	return ( (const MiqtVirtualQsciLexerPOV*)(self) )->virtualbase_WriteProperties(qs, prefix);
}

void QsciLexerPOV_Delete(QsciLexerPOV* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQsciLexerPOV*>( self );
	} else {
		delete self;
	}
}

