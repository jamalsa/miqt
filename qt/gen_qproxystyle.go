package qt

/*

#include "gen_qproxystyle.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QProxyStyle struct {
	h          *C.QProxyStyle
	isSubclass bool
	*QCommonStyle
}

func (this *QProxyStyle) cPointer() *C.QProxyStyle {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QProxyStyle) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQProxyStyle constructs the type using only CGO pointers.
func newQProxyStyle(h *C.QProxyStyle, h_QCommonStyle *C.QCommonStyle, h_QStyle *C.QStyle, h_QObject *C.QObject) *QProxyStyle {
	if h == nil {
		return nil
	}
	return &QProxyStyle{h: h,
		QCommonStyle: newQCommonStyle(h_QCommonStyle, h_QStyle, h_QObject)}
}

// UnsafeNewQProxyStyle constructs the type using only unsafe pointers.
func UnsafeNewQProxyStyle(h unsafe.Pointer, h_QCommonStyle unsafe.Pointer, h_QStyle unsafe.Pointer, h_QObject unsafe.Pointer) *QProxyStyle {
	if h == nil {
		return nil
	}

	return &QProxyStyle{h: (*C.QProxyStyle)(h),
		QCommonStyle: UnsafeNewQCommonStyle(h_QCommonStyle, h_QStyle, h_QObject)}
}

// NewQProxyStyle constructs a new QProxyStyle object.
func NewQProxyStyle() *QProxyStyle {
	var outptr_QProxyStyle *C.QProxyStyle = nil
	var outptr_QCommonStyle *C.QCommonStyle = nil
	var outptr_QStyle *C.QStyle = nil
	var outptr_QObject *C.QObject = nil

	C.QProxyStyle_new(&outptr_QProxyStyle, &outptr_QCommonStyle, &outptr_QStyle, &outptr_QObject)
	ret := newQProxyStyle(outptr_QProxyStyle, outptr_QCommonStyle, outptr_QStyle, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQProxyStyle2 constructs a new QProxyStyle object.
func NewQProxyStyle2(key string) *QProxyStyle {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	var outptr_QProxyStyle *C.QProxyStyle = nil
	var outptr_QCommonStyle *C.QCommonStyle = nil
	var outptr_QStyle *C.QStyle = nil
	var outptr_QObject *C.QObject = nil

	C.QProxyStyle_new2(key_ms, &outptr_QProxyStyle, &outptr_QCommonStyle, &outptr_QStyle, &outptr_QObject)
	ret := newQProxyStyle(outptr_QProxyStyle, outptr_QCommonStyle, outptr_QStyle, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQProxyStyle3 constructs a new QProxyStyle object.
func NewQProxyStyle3(style *QStyle) *QProxyStyle {
	var outptr_QProxyStyle *C.QProxyStyle = nil
	var outptr_QCommonStyle *C.QCommonStyle = nil
	var outptr_QStyle *C.QStyle = nil
	var outptr_QObject *C.QObject = nil

	C.QProxyStyle_new3(style.cPointer(), &outptr_QProxyStyle, &outptr_QCommonStyle, &outptr_QStyle, &outptr_QObject)
	ret := newQProxyStyle(outptr_QProxyStyle, outptr_QCommonStyle, outptr_QStyle, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QProxyStyle) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QProxyStyle_MetaObject(this.h)))
}

func (this *QProxyStyle) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QProxyStyle_Metacast(this.h, param1_Cstring))
}

func QProxyStyle_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QProxyStyle_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProxyStyle_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QProxyStyle_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProxyStyle) BaseStyle() *QStyle {
	return UnsafeNewQStyle(unsafe.Pointer(C.QProxyStyle_BaseStyle(this.h)), nil)
}

func (this *QProxyStyle) SetBaseStyle(style *QStyle) {
	C.QProxyStyle_SetBaseStyle(this.h, style.cPointer())
}

func (this *QProxyStyle) DrawPrimitive(element QStyle__PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	C.QProxyStyle_DrawPrimitive(this.h, (C.int)(element), option.cPointer(), painter.cPointer(), widget.cPointer())
}

func (this *QProxyStyle) DrawControl(element QStyle__ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	C.QProxyStyle_DrawControl(this.h, (C.int)(element), option.cPointer(), painter.cPointer(), widget.cPointer())
}

func (this *QProxyStyle) DrawComplexControl(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget) {
	C.QProxyStyle_DrawComplexControl(this.h, (C.int)(control), option.cPointer(), painter.cPointer(), widget.cPointer())
}

func (this *QProxyStyle) DrawItemText(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QProxyStyle_DrawItemText(this.h, painter.cPointer(), rect.cPointer(), (C.int)(flags), pal.cPointer(), (C.bool)(enabled), text_ms, (C.int)(textRole))
}

func (this *QProxyStyle) DrawItemPixmap(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	C.QProxyStyle_DrawItemPixmap(this.h, painter.cPointer(), rect.cPointer(), (C.int)(alignment), pixmap.cPointer())
}

func (this *QProxyStyle) SizeFromContents(typeVal QStyle__ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize {
	_ret := C.QProxyStyle_SizeFromContents(this.h, (C.int)(typeVal), option.cPointer(), size.cPointer(), widget.cPointer())
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) SubElementRect(element QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect {
	_ret := C.QProxyStyle_SubElementRect(this.h, (C.int)(element), option.cPointer(), widget.cPointer())
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) SubControlRect(cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect {
	_ret := C.QProxyStyle_SubControlRect(this.h, (C.int)(cc), opt.cPointer(), (C.int)(sc), widget.cPointer())
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) ItemTextRect(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	_ret := C.QProxyStyle_ItemTextRect(this.h, fm.cPointer(), r.cPointer(), (C.int)(flags), (C.bool)(enabled), text_ms)
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) ItemPixmapRect(r *QRect, flags int, pixmap *QPixmap) *QRect {
	_ret := C.QProxyStyle_ItemPixmapRect(this.h, r.cPointer(), (C.int)(flags), pixmap.cPointer())
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) HitTestComplexControl(control QStyle__ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) QStyle__SubControl {
	return (QStyle__SubControl)(C.QProxyStyle_HitTestComplexControl(this.h, (C.int)(control), option.cPointer(), pos.cPointer(), widget.cPointer()))
}

func (this *QProxyStyle) StyleHint(hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {
	return (int)(C.QProxyStyle_StyleHint(this.h, (C.int)(hint), option.cPointer(), widget.cPointer(), returnData.cPointer()))
}

func (this *QProxyStyle) PixelMetric(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int {
	return (int)(C.QProxyStyle_PixelMetric(this.h, (C.int)(metric), option.cPointer(), widget.cPointer()))
}

func (this *QProxyStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {
	return (int)(C.QProxyStyle_LayoutSpacing(this.h, (C.int)(control1), (C.int)(control2), (C.int)(orientation), option.cPointer(), widget.cPointer()))
}

func (this *QProxyStyle) StandardIcon(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon {
	_ret := C.QProxyStyle_StandardIcon(this.h, (C.int)(standardIcon), option.cPointer(), widget.cPointer())
	_goptr := newQIcon(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) StandardPixmap(standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	_ret := C.QProxyStyle_StandardPixmap(this.h, (C.int)(standardPixmap), opt.cPointer(), widget.cPointer())
	_goptr := newQPixmap(_ret, nil)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) GeneratedIconPixmap(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	_ret := C.QProxyStyle_GeneratedIconPixmap(this.h, (C.int)(iconMode), pixmap.cPointer(), opt.cPointer())
	_goptr := newQPixmap(_ret, nil)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) StandardPalette() *QPalette {
	_ret := C.QProxyStyle_StandardPalette(this.h)
	_goptr := newQPalette(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) Polish(widget *QWidget) {
	C.QProxyStyle_Polish(this.h, widget.cPointer())
}

func (this *QProxyStyle) PolishWithPal(pal *QPalette) {
	C.QProxyStyle_PolishWithPal(this.h, pal.cPointer())
}

func (this *QProxyStyle) PolishWithApp(app *QApplication) {
	C.QProxyStyle_PolishWithApp(this.h, app.cPointer())
}

func (this *QProxyStyle) Unpolish(widget *QWidget) {
	C.QProxyStyle_Unpolish(this.h, widget.cPointer())
}

func (this *QProxyStyle) UnpolishWithApp(app *QApplication) {
	C.QProxyStyle_UnpolishWithApp(this.h, app.cPointer())
}

func QProxyStyle_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QProxyStyle_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProxyStyle_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QProxyStyle_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProxyStyle_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QProxyStyle_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProxyStyle_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QProxyStyle_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProxyStyle) callVirtualBase_DrawPrimitive(element QStyle__PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget) {

	C.QProxyStyle_virtualbase_DrawPrimitive(unsafe.Pointer(this.h), (C.int)(element), option.cPointer(), painter.cPointer(), widget.cPointer())

}
func (this *QProxyStyle) OnDrawPrimitive(slot func(super func(element QStyle__PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget), element QStyle__PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget)) {
	C.QProxyStyle_override_virtual_DrawPrimitive(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawPrimitive
func miqt_exec_callback_QProxyStyle_DrawPrimitive(self *C.QProxyStyle, cb C.intptr_t, element C.int, option *C.QStyleOption, painter *C.QPainter, widget *C.QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(element QStyle__PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget), element QStyle__PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__PrimitiveElement)(element)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval3 := UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval4 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawPrimitive, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_DrawControl(element QStyle__ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget) {

	C.QProxyStyle_virtualbase_DrawControl(unsafe.Pointer(this.h), (C.int)(element), option.cPointer(), painter.cPointer(), widget.cPointer())

}
func (this *QProxyStyle) OnDrawControl(slot func(super func(element QStyle__ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget), element QStyle__ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget)) {
	C.QProxyStyle_override_virtual_DrawControl(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawControl
func miqt_exec_callback_QProxyStyle_DrawControl(self *C.QProxyStyle, cb C.intptr_t, element C.int, option *C.QStyleOption, painter *C.QPainter, widget *C.QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(element QStyle__ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget), element QStyle__ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__ControlElement)(element)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval3 := UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval4 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawControl, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_DrawComplexControl(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget) {

	C.QProxyStyle_virtualbase_DrawComplexControl(unsafe.Pointer(this.h), (C.int)(control), option.cPointer(), painter.cPointer(), widget.cPointer())

}
func (this *QProxyStyle) OnDrawComplexControl(slot func(super func(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget), control QStyle__ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget)) {
	C.QProxyStyle_override_virtual_DrawComplexControl(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawComplexControl
func miqt_exec_callback_QProxyStyle_DrawComplexControl(self *C.QProxyStyle, cb C.intptr_t, control C.int, option *C.QStyleOptionComplex, painter *C.QPainter, widget *C.QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget), control QStyle__ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__ComplexControl)(control)

	slotval2 := UnsafeNewQStyleOptionComplex(unsafe.Pointer(option), nil)
	slotval3 := UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval4 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawComplexControl, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_DrawItemText(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	C.QProxyStyle_virtualbase_DrawItemText(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), (C.int)(flags), pal.cPointer(), (C.bool)(enabled), text_ms, (C.int)(textRole))

}
func (this *QProxyStyle) OnDrawItemText(slot func(super func(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole), painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole)) {
	C.QProxyStyle_override_virtual_DrawItemText(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawItemText
func miqt_exec_callback_QProxyStyle_DrawItemText(self *C.QProxyStyle, cb C.intptr_t, painter *C.QPainter, rect *C.QRect, flags C.int, pal *C.QPalette, enabled C.bool, text C.struct_miqt_string, textRole C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole), painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval2 := UnsafeNewQRect(unsafe.Pointer(rect))
	slotval3 := (int)(flags)

	slotval4 := UnsafeNewQPalette(unsafe.Pointer(pal))
	slotval5 := (bool)(enabled)

	var text_ms C.struct_miqt_string = text
	text_ret := C.GoStringN(text_ms.data, C.int(int64(text_ms.len)))
	C.free(unsafe.Pointer(text_ms.data))
	slotval6 := text_ret
	slotval7 := (QPalette__ColorRole)(textRole)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawItemText, slotval1, slotval2, slotval3, slotval4, slotval5, slotval6, slotval7)

}

func (this *QProxyStyle) callVirtualBase_DrawItemPixmap(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {

	C.QProxyStyle_virtualbase_DrawItemPixmap(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), (C.int)(alignment), pixmap.cPointer())

}
func (this *QProxyStyle) OnDrawItemPixmap(slot func(super func(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap), painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap)) {
	C.QProxyStyle_override_virtual_DrawItemPixmap(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawItemPixmap
func miqt_exec_callback_QProxyStyle_DrawItemPixmap(self *C.QProxyStyle, cb C.intptr_t, painter *C.QPainter, rect *C.QRect, alignment C.int, pixmap *C.QPixmap) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap), painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval2 := UnsafeNewQRect(unsafe.Pointer(rect))
	slotval3 := (int)(alignment)

	slotval4 := UnsafeNewQPixmap(unsafe.Pointer(pixmap), nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawItemPixmap, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_SizeFromContents(typeVal QStyle__ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize {

	_ret := C.QProxyStyle_virtualbase_SizeFromContents(unsafe.Pointer(this.h), (C.int)(typeVal), option.cPointer(), size.cPointer(), widget.cPointer())
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnSizeFromContents(slot func(super func(typeVal QStyle__ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize, typeVal QStyle__ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize) {
	C.QProxyStyle_override_virtual_SizeFromContents(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_SizeFromContents
func miqt_exec_callback_QProxyStyle_SizeFromContents(self *C.QProxyStyle, cb C.intptr_t, typeVal C.int, option *C.QStyleOption, size *C.QSize, widget *C.QWidget) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal QStyle__ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize, typeVal QStyle__ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__ContentsType)(typeVal)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval3 := UnsafeNewQSize(unsafe.Pointer(size))
	slotval4 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_SizeFromContents, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_SubElementRect(element QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect {

	_ret := C.QProxyStyle_virtualbase_SubElementRect(unsafe.Pointer(this.h), (C.int)(element), option.cPointer(), widget.cPointer())
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnSubElementRect(slot func(super func(element QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect, element QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect) {
	C.QProxyStyle_override_virtual_SubElementRect(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_SubElementRect
func miqt_exec_callback_QProxyStyle_SubElementRect(self *C.QProxyStyle, cb C.intptr_t, element C.int, option *C.QStyleOption, widget *C.QWidget) *C.QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(element QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect, element QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__SubElement)(element)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval3 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_SubElementRect, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_SubControlRect(cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect {

	_ret := C.QProxyStyle_virtualbase_SubControlRect(unsafe.Pointer(this.h), (C.int)(cc), opt.cPointer(), (C.int)(sc), widget.cPointer())
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnSubControlRect(slot func(super func(cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect, cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect) {
	C.QProxyStyle_override_virtual_SubControlRect(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_SubControlRect
func miqt_exec_callback_QProxyStyle_SubControlRect(self *C.QProxyStyle, cb C.intptr_t, cc C.int, opt *C.QStyleOptionComplex, sc C.int, widget *C.QWidget) *C.QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect, cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__ComplexControl)(cc)

	slotval2 := UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt), nil)
	slotval3 := (QStyle__SubControl)(sc)

	slotval4 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_SubControlRect, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_ItemTextRect(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	_ret := C.QProxyStyle_virtualbase_ItemTextRect(unsafe.Pointer(this.h), fm.cPointer(), r.cPointer(), (C.int)(flags), (C.bool)(enabled), text_ms)
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnItemTextRect(slot func(super func(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect, fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect) {
	C.QProxyStyle_override_virtual_ItemTextRect(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_ItemTextRect
func miqt_exec_callback_QProxyStyle_ItemTextRect(self *C.QProxyStyle, cb C.intptr_t, fm *C.QFontMetrics, r *C.QRect, flags C.int, enabled C.bool, text C.struct_miqt_string) *C.QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect, fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQFontMetrics(unsafe.Pointer(fm))
	slotval2 := UnsafeNewQRect(unsafe.Pointer(r))
	slotval3 := (int)(flags)

	slotval4 := (bool)(enabled)

	var text_ms C.struct_miqt_string = text
	text_ret := C.GoStringN(text_ms.data, C.int(int64(text_ms.len)))
	C.free(unsafe.Pointer(text_ms.data))
	slotval5 := text_ret

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_ItemTextRect, slotval1, slotval2, slotval3, slotval4, slotval5)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_ItemPixmapRect(r *QRect, flags int, pixmap *QPixmap) *QRect {

	_ret := C.QProxyStyle_virtualbase_ItemPixmapRect(unsafe.Pointer(this.h), r.cPointer(), (C.int)(flags), pixmap.cPointer())
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnItemPixmapRect(slot func(super func(r *QRect, flags int, pixmap *QPixmap) *QRect, r *QRect, flags int, pixmap *QPixmap) *QRect) {
	C.QProxyStyle_override_virtual_ItemPixmapRect(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_ItemPixmapRect
func miqt_exec_callback_QProxyStyle_ItemPixmapRect(self *C.QProxyStyle, cb C.intptr_t, r *C.QRect, flags C.int, pixmap *C.QPixmap) *C.QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r *QRect, flags int, pixmap *QPixmap) *QRect, r *QRect, flags int, pixmap *QPixmap) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQRect(unsafe.Pointer(r))
	slotval2 := (int)(flags)

	slotval3 := UnsafeNewQPixmap(unsafe.Pointer(pixmap), nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_ItemPixmapRect, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_HitTestComplexControl(control QStyle__ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) QStyle__SubControl {

	return (QStyle__SubControl)(C.QProxyStyle_virtualbase_HitTestComplexControl(unsafe.Pointer(this.h), (C.int)(control), option.cPointer(), pos.cPointer(), widget.cPointer()))

}
func (this *QProxyStyle) OnHitTestComplexControl(slot func(super func(control QStyle__ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) QStyle__SubControl, control QStyle__ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) QStyle__SubControl) {
	C.QProxyStyle_override_virtual_HitTestComplexControl(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_HitTestComplexControl
func miqt_exec_callback_QProxyStyle_HitTestComplexControl(self *C.QProxyStyle, cb C.intptr_t, control C.int, option *C.QStyleOptionComplex, pos *C.QPoint, widget *C.QWidget) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(control QStyle__ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) QStyle__SubControl, control QStyle__ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) QStyle__SubControl)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__ComplexControl)(control)

	slotval2 := UnsafeNewQStyleOptionComplex(unsafe.Pointer(option), nil)
	slotval3 := UnsafeNewQPoint(unsafe.Pointer(pos))
	slotval4 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_HitTestComplexControl, slotval1, slotval2, slotval3, slotval4)

	return (C.int)(virtualReturn)

}

func (this *QProxyStyle) callVirtualBase_StyleHint(hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {

	return (int)(C.QProxyStyle_virtualbase_StyleHint(unsafe.Pointer(this.h), (C.int)(hint), option.cPointer(), widget.cPointer(), returnData.cPointer()))

}
func (this *QProxyStyle) OnStyleHint(slot func(super func(hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int, hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int) {
	C.QProxyStyle_override_virtual_StyleHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StyleHint
func miqt_exec_callback_QProxyStyle_StyleHint(self *C.QProxyStyle, cb C.intptr_t, hint C.int, option *C.QStyleOption, widget *C.QWidget, returnData *C.QStyleHintReturn) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int, hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__StyleHint)(hint)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval3 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)
	slotval4 := UnsafeNewQStyleHintReturn(unsafe.Pointer(returnData))

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StyleHint, slotval1, slotval2, slotval3, slotval4)

	return (C.int)(virtualReturn)

}

func (this *QProxyStyle) callVirtualBase_PixelMetric(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int {

	return (int)(C.QProxyStyle_virtualbase_PixelMetric(unsafe.Pointer(this.h), (C.int)(metric), option.cPointer(), widget.cPointer()))

}
func (this *QProxyStyle) OnPixelMetric(slot func(super func(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int, metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int) {
	C.QProxyStyle_override_virtual_PixelMetric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_PixelMetric
func miqt_exec_callback_QProxyStyle_PixelMetric(self *C.QProxyStyle, cb C.intptr_t, metric C.int, option *C.QStyleOption, widget *C.QWidget) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int, metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__PixelMetric)(metric)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval3 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_PixelMetric, slotval1, slotval2, slotval3)

	return (C.int)(virtualReturn)

}

func (this *QProxyStyle) callVirtualBase_LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {

	return (int)(C.QProxyStyle_virtualbase_LayoutSpacing(unsafe.Pointer(this.h), (C.int)(control1), (C.int)(control2), (C.int)(orientation), option.cPointer(), widget.cPointer()))

}
func (this *QProxyStyle) OnLayoutSpacing(slot func(super func(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int, control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int) {
	C.QProxyStyle_override_virtual_LayoutSpacing(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_LayoutSpacing
func miqt_exec_callback_QProxyStyle_LayoutSpacing(self *C.QProxyStyle, cb C.intptr_t, control1 C.int, control2 C.int, orientation C.int, option *C.QStyleOption, widget *C.QWidget) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int, control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QSizePolicy__ControlType)(control1)

	slotval2 := (QSizePolicy__ControlType)(control2)

	slotval3 := (Orientation)(orientation)

	slotval4 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval5 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_LayoutSpacing, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (C.int)(virtualReturn)

}

func (this *QProxyStyle) callVirtualBase_StandardIcon(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon {

	_ret := C.QProxyStyle_virtualbase_StandardIcon(unsafe.Pointer(this.h), (C.int)(standardIcon), option.cPointer(), widget.cPointer())
	_goptr := newQIcon(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnStandardIcon(slot func(super func(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon, standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon) {
	C.QProxyStyle_override_virtual_StandardIcon(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StandardIcon
func miqt_exec_callback_QProxyStyle_StandardIcon(self *C.QProxyStyle, cb C.intptr_t, standardIcon C.int, option *C.QStyleOption, widget *C.QWidget) *C.QIcon {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon, standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__StandardPixmap)(standardIcon)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(option))
	slotval3 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StandardIcon, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_StandardPixmap(standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {

	_ret := C.QProxyStyle_virtualbase_StandardPixmap(unsafe.Pointer(this.h), (C.int)(standardPixmap), opt.cPointer(), widget.cPointer())
	_goptr := newQPixmap(_ret, nil)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnStandardPixmap(slot func(super func(standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap, standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap) {
	C.QProxyStyle_override_virtual_StandardPixmap(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StandardPixmap
func miqt_exec_callback_QProxyStyle_StandardPixmap(self *C.QProxyStyle, cb C.intptr_t, standardPixmap C.int, opt *C.QStyleOption, widget *C.QWidget) *C.QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap, standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QStyle__StandardPixmap)(standardPixmap)

	slotval2 := UnsafeNewQStyleOption(unsafe.Pointer(opt))
	slotval3 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StandardPixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_GeneratedIconPixmap(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {

	_ret := C.QProxyStyle_virtualbase_GeneratedIconPixmap(unsafe.Pointer(this.h), (C.int)(iconMode), pixmap.cPointer(), opt.cPointer())
	_goptr := newQPixmap(_ret, nil)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnGeneratedIconPixmap(slot func(super func(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap, iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap) {
	C.QProxyStyle_override_virtual_GeneratedIconPixmap(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_GeneratedIconPixmap
func miqt_exec_callback_QProxyStyle_GeneratedIconPixmap(self *C.QProxyStyle, cb C.intptr_t, iconMode C.int, pixmap *C.QPixmap, opt *C.QStyleOption) *C.QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap, iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QIcon__Mode)(iconMode)

	slotval2 := UnsafeNewQPixmap(unsafe.Pointer(pixmap), nil)
	slotval3 := UnsafeNewQStyleOption(unsafe.Pointer(opt))

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_GeneratedIconPixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_StandardPalette() *QPalette {

	_ret := C.QProxyStyle_virtualbase_StandardPalette(unsafe.Pointer(this.h))
	_goptr := newQPalette(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnStandardPalette(slot func(super func() *QPalette) *QPalette) {
	C.QProxyStyle_override_virtual_StandardPalette(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StandardPalette
func miqt_exec_callback_QProxyStyle_StandardPalette(self *C.QProxyStyle, cb C.intptr_t) *C.QPalette {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPalette) *QPalette)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StandardPalette)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_Polish(widget *QWidget) {

	C.QProxyStyle_virtualbase_Polish(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QProxyStyle) OnPolish(slot func(super func(widget *QWidget), widget *QWidget)) {
	C.QProxyStyle_override_virtual_Polish(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_Polish
func miqt_exec_callback_QProxyStyle_Polish(self *C.QProxyStyle, cb C.intptr_t, widget *C.QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_Polish, slotval1)

}

func (this *QProxyStyle) callVirtualBase_PolishWithPal(pal *QPalette) {

	C.QProxyStyle_virtualbase_PolishWithPal(unsafe.Pointer(this.h), pal.cPointer())

}
func (this *QProxyStyle) OnPolishWithPal(slot func(super func(pal *QPalette), pal *QPalette)) {
	C.QProxyStyle_override_virtual_PolishWithPal(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_PolishWithPal
func miqt_exec_callback_QProxyStyle_PolishWithPal(self *C.QProxyStyle, cb C.intptr_t, pal *C.QPalette) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pal *QPalette), pal *QPalette))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPalette(unsafe.Pointer(pal))

	gofunc((&QProxyStyle{h: self}).callVirtualBase_PolishWithPal, slotval1)

}

func (this *QProxyStyle) callVirtualBase_PolishWithApp(app *QApplication) {

	C.QProxyStyle_virtualbase_PolishWithApp(unsafe.Pointer(this.h), app.cPointer())

}
func (this *QProxyStyle) OnPolishWithApp(slot func(super func(app *QApplication), app *QApplication)) {
	C.QProxyStyle_override_virtual_PolishWithApp(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_PolishWithApp
func miqt_exec_callback_QProxyStyle_PolishWithApp(self *C.QProxyStyle, cb C.intptr_t, app *C.QApplication) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(app *QApplication), app *QApplication))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQApplication(unsafe.Pointer(app), nil, nil, nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_PolishWithApp, slotval1)

}

func (this *QProxyStyle) callVirtualBase_Unpolish(widget *QWidget) {

	C.QProxyStyle_virtualbase_Unpolish(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QProxyStyle) OnUnpolish(slot func(super func(widget *QWidget), widget *QWidget)) {
	C.QProxyStyle_override_virtual_Unpolish(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_Unpolish
func miqt_exec_callback_QProxyStyle_Unpolish(self *C.QProxyStyle, cb C.intptr_t, widget *C.QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQWidget(unsafe.Pointer(widget), nil, nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_Unpolish, slotval1)

}

func (this *QProxyStyle) callVirtualBase_UnpolishWithApp(app *QApplication) {

	C.QProxyStyle_virtualbase_UnpolishWithApp(unsafe.Pointer(this.h), app.cPointer())

}
func (this *QProxyStyle) OnUnpolishWithApp(slot func(super func(app *QApplication), app *QApplication)) {
	C.QProxyStyle_override_virtual_UnpolishWithApp(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_UnpolishWithApp
func miqt_exec_callback_QProxyStyle_UnpolishWithApp(self *C.QProxyStyle, cb C.intptr_t, app *C.QApplication) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(app *QApplication), app *QApplication))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQApplication(unsafe.Pointer(app), nil, nil, nil)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_UnpolishWithApp, slotval1)

}

func (this *QProxyStyle) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(C.QProxyStyle_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QProxyStyle) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	C.QProxyStyle_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_Event
func miqt_exec_callback_QProxyStyle_Event(self *C.QProxyStyle, cb C.intptr_t, e *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(e))

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QProxyStyle) Delete() {
	C.QProxyStyle_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QProxyStyle) GoGC() {
	runtime.SetFinalizer(this, func(this *QProxyStyle) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
