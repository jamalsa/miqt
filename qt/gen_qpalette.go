package qt

/*

#include "gen_qpalette.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QPalette__ColorGroup int

const (
	QPalette__ColorGroup__Active       QPalette__ColorGroup = 0
	QPalette__ColorGroup__Disabled     QPalette__ColorGroup = 1
	QPalette__ColorGroup__Inactive     QPalette__ColorGroup = 2
	QPalette__ColorGroup__NColorGroups QPalette__ColorGroup = 3
	QPalette__ColorGroup__Current      QPalette__ColorGroup = 4
	QPalette__ColorGroup__All          QPalette__ColorGroup = 5
	QPalette__ColorGroup__Normal       QPalette__ColorGroup = 0
)

type QPalette__ColorRole int

const (
	QPalette__ColorRole__WindowText      QPalette__ColorRole = 0
	QPalette__ColorRole__Button          QPalette__ColorRole = 1
	QPalette__ColorRole__Light           QPalette__ColorRole = 2
	QPalette__ColorRole__Midlight        QPalette__ColorRole = 3
	QPalette__ColorRole__Dark            QPalette__ColorRole = 4
	QPalette__ColorRole__Mid             QPalette__ColorRole = 5
	QPalette__ColorRole__Text            QPalette__ColorRole = 6
	QPalette__ColorRole__BrightText      QPalette__ColorRole = 7
	QPalette__ColorRole__ButtonText      QPalette__ColorRole = 8
	QPalette__ColorRole__Base            QPalette__ColorRole = 9
	QPalette__ColorRole__Window          QPalette__ColorRole = 10
	QPalette__ColorRole__Shadow          QPalette__ColorRole = 11
	QPalette__ColorRole__Highlight       QPalette__ColorRole = 12
	QPalette__ColorRole__HighlightedText QPalette__ColorRole = 13
	QPalette__ColorRole__Link            QPalette__ColorRole = 14
	QPalette__ColorRole__LinkVisited     QPalette__ColorRole = 15
	QPalette__ColorRole__AlternateBase   QPalette__ColorRole = 16
	QPalette__ColorRole__NoRole          QPalette__ColorRole = 17
	QPalette__ColorRole__ToolTipBase     QPalette__ColorRole = 18
	QPalette__ColorRole__ToolTipText     QPalette__ColorRole = 19
	QPalette__ColorRole__PlaceholderText QPalette__ColorRole = 20
	QPalette__ColorRole__NColorRoles     QPalette__ColorRole = 21
	QPalette__ColorRole__Foreground      QPalette__ColorRole = 0
	QPalette__ColorRole__Background      QPalette__ColorRole = 10
)

type QPalette struct {
	h *C.QPalette
}

func (this *QPalette) cPointer() *C.QPalette {
	if this == nil {
		return nil
	}
	return this.h
}

func newQPalette(h *C.QPalette) *QPalette {
	if h == nil {
		return nil
	}
	return &QPalette{h: h}
}

func newQPalette_U(h unsafe.Pointer) *QPalette {
	return newQPalette((*C.QPalette)(h))
}

// NewQPalette constructs a new QPalette object.
func NewQPalette() *QPalette {
	ret := C.QPalette_new()
	return newQPalette(ret)
}

// NewQPalette2 constructs a new QPalette object.
func NewQPalette2(button *QColor) *QPalette {
	ret := C.QPalette_new2(button.cPointer())
	return newQPalette(ret)
}

// NewQPalette3 constructs a new QPalette object.
func NewQPalette3(button GlobalColor) *QPalette {
	ret := C.QPalette_new3((C.int)(button))
	return newQPalette(ret)
}

// NewQPalette4 constructs a new QPalette object.
func NewQPalette4(button *QColor, window *QColor) *QPalette {
	ret := C.QPalette_new4(button.cPointer(), window.cPointer())
	return newQPalette(ret)
}

// NewQPalette5 constructs a new QPalette object.
func NewQPalette5(windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) *QPalette {
	ret := C.QPalette_new5(windowText.cPointer(), button.cPointer(), light.cPointer(), dark.cPointer(), mid.cPointer(), text.cPointer(), bright_text.cPointer(), base.cPointer(), window.cPointer())
	return newQPalette(ret)
}

// NewQPalette6 constructs a new QPalette object.
func NewQPalette6(windowText *QColor, window *QColor, light *QColor, dark *QColor, mid *QColor, text *QColor, base *QColor) *QPalette {
	ret := C.QPalette_new6(windowText.cPointer(), window.cPointer(), light.cPointer(), dark.cPointer(), mid.cPointer(), text.cPointer(), base.cPointer())
	return newQPalette(ret)
}

// NewQPalette7 constructs a new QPalette object.
func NewQPalette7(palette *QPalette) *QPalette {
	ret := C.QPalette_new7(palette.cPointer())
	return newQPalette(ret)
}

func (this *QPalette) OperatorAssign(palette *QPalette) {
	C.QPalette_OperatorAssign(this.h, palette.cPointer())
}

func (this *QPalette) Swap(other *QPalette) {
	C.QPalette_Swap(this.h, other.cPointer())
}

func (this *QPalette) CurrentColorGroup() QPalette__ColorGroup {
	return (QPalette__ColorGroup)(C.QPalette_CurrentColorGroup(this.h))
}

func (this *QPalette) SetCurrentColorGroup(cg QPalette__ColorGroup) {
	C.QPalette_SetCurrentColorGroup(this.h, (C.int)(cg))
}

func (this *QPalette) Color(cg QPalette__ColorGroup, cr QPalette__ColorRole) *QColor {
	return newQColor_U(unsafe.Pointer(C.QPalette_Color(this.h, (C.int)(cg), (C.int)(cr))))
}

func (this *QPalette) Brush(cg QPalette__ColorGroup, cr QPalette__ColorRole) *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Brush(this.h, (C.int)(cg), (C.int)(cr))))
}

func (this *QPalette) SetColor(cg QPalette__ColorGroup, cr QPalette__ColorRole, color *QColor) {
	C.QPalette_SetColor(this.h, (C.int)(cg), (C.int)(cr), color.cPointer())
}

func (this *QPalette) SetColor2(cr QPalette__ColorRole, color *QColor) {
	C.QPalette_SetColor2(this.h, (C.int)(cr), color.cPointer())
}

func (this *QPalette) SetBrush(cr QPalette__ColorRole, brush *QBrush) {
	C.QPalette_SetBrush(this.h, (C.int)(cr), brush.cPointer())
}

func (this *QPalette) IsBrushSet(cg QPalette__ColorGroup, cr QPalette__ColorRole) bool {
	return (bool)(C.QPalette_IsBrushSet(this.h, (C.int)(cg), (C.int)(cr)))
}

func (this *QPalette) SetBrush2(cg QPalette__ColorGroup, cr QPalette__ColorRole, brush *QBrush) {
	C.QPalette_SetBrush2(this.h, (C.int)(cg), (C.int)(cr), brush.cPointer())
}

func (this *QPalette) SetColorGroup(cr QPalette__ColorGroup, windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) {
	C.QPalette_SetColorGroup(this.h, (C.int)(cr), windowText.cPointer(), button.cPointer(), light.cPointer(), dark.cPointer(), mid.cPointer(), text.cPointer(), bright_text.cPointer(), base.cPointer(), window.cPointer())
}

func (this *QPalette) IsEqual(cr1 QPalette__ColorGroup, cr2 QPalette__ColorGroup) bool {
	return (bool)(C.QPalette_IsEqual(this.h, (C.int)(cr1), (C.int)(cr2)))
}

func (this *QPalette) ColorWithCr(cr QPalette__ColorRole) *QColor {
	return newQColor_U(unsafe.Pointer(C.QPalette_ColorWithCr(this.h, (C.int)(cr))))
}

func (this *QPalette) BrushWithCr(cr QPalette__ColorRole) *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_BrushWithCr(this.h, (C.int)(cr))))
}

func (this *QPalette) WindowText() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_WindowText(this.h)))
}

func (this *QPalette) Button() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Button(this.h)))
}

func (this *QPalette) Light() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Light(this.h)))
}

func (this *QPalette) Dark() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Dark(this.h)))
}

func (this *QPalette) Mid() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Mid(this.h)))
}

func (this *QPalette) Text() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Text(this.h)))
}

func (this *QPalette) Base() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Base(this.h)))
}

func (this *QPalette) AlternateBase() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_AlternateBase(this.h)))
}

func (this *QPalette) ToolTipBase() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_ToolTipBase(this.h)))
}

func (this *QPalette) ToolTipText() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_ToolTipText(this.h)))
}

func (this *QPalette) Window() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Window(this.h)))
}

func (this *QPalette) Midlight() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Midlight(this.h)))
}

func (this *QPalette) BrightText() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_BrightText(this.h)))
}

func (this *QPalette) ButtonText() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_ButtonText(this.h)))
}

func (this *QPalette) Shadow() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Shadow(this.h)))
}

func (this *QPalette) Highlight() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Highlight(this.h)))
}

func (this *QPalette) HighlightedText() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_HighlightedText(this.h)))
}

func (this *QPalette) Link() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Link(this.h)))
}

func (this *QPalette) LinkVisited() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_LinkVisited(this.h)))
}

func (this *QPalette) PlaceholderText() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_PlaceholderText(this.h)))
}

func (this *QPalette) Foreground() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Foreground(this.h)))
}

func (this *QPalette) Background() *QBrush {
	return newQBrush_U(unsafe.Pointer(C.QPalette_Background(this.h)))
}

func (this *QPalette) OperatorEqual(p *QPalette) bool {
	return (bool)(C.QPalette_OperatorEqual(this.h, p.cPointer()))
}

func (this *QPalette) OperatorNotEqual(p *QPalette) bool {
	return (bool)(C.QPalette_OperatorNotEqual(this.h, p.cPointer()))
}

func (this *QPalette) IsCopyOf(p *QPalette) bool {
	return (bool)(C.QPalette_IsCopyOf(this.h, p.cPointer()))
}

func (this *QPalette) CacheKey() int64 {
	return (int64)(C.QPalette_CacheKey(this.h))
}

func (this *QPalette) Resolve(param1 *QPalette) *QPalette {
	_ret := C.QPalette_Resolve(this.h, param1.cPointer())
	_goptr := newQPalette(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPalette) Resolve2() uint {
	return (uint)(C.QPalette_Resolve2(this.h))
}

func (this *QPalette) ResolveWithMask(mask uint) {
	C.QPalette_ResolveWithMask(this.h, (C.uint)(mask))
}

// Delete this object from C++ memory.
func (this *QPalette) Delete() {
	C.QPalette_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPalette) GoGC() {
	runtime.SetFinalizer(this, func(this *QPalette) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
