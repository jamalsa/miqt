package qt

/*

#include "gen_qcheckbox.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QCheckBox struct {
	h          *C.QCheckBox
	isSubclass bool
	*QAbstractButton
}

func (this *QCheckBox) cPointer() *C.QCheckBox {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QCheckBox) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQCheckBox constructs the type using only CGO pointers.
func newQCheckBox(h *C.QCheckBox, h_QAbstractButton *C.QAbstractButton, h_QWidget *C.QWidget, h_QObject *C.QObject, h_QPaintDevice *C.QPaintDevice) *QCheckBox {
	if h == nil {
		return nil
	}
	return &QCheckBox{h: h,
		QAbstractButton: newQAbstractButton(h_QAbstractButton, h_QWidget, h_QObject, h_QPaintDevice)}
}

// UnsafeNewQCheckBox constructs the type using only unsafe pointers.
func UnsafeNewQCheckBox(h unsafe.Pointer, h_QAbstractButton unsafe.Pointer, h_QWidget unsafe.Pointer, h_QObject unsafe.Pointer, h_QPaintDevice unsafe.Pointer) *QCheckBox {
	if h == nil {
		return nil
	}

	return &QCheckBox{h: (*C.QCheckBox)(h),
		QAbstractButton: UnsafeNewQAbstractButton(h_QAbstractButton, h_QWidget, h_QObject, h_QPaintDevice)}
}

// NewQCheckBox constructs a new QCheckBox object.
func NewQCheckBox(parent *QWidget) *QCheckBox {
	var outptr_QCheckBox *C.QCheckBox = nil
	var outptr_QAbstractButton *C.QAbstractButton = nil
	var outptr_QWidget *C.QWidget = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QCheckBox_new(parent.cPointer(), &outptr_QCheckBox, &outptr_QAbstractButton, &outptr_QWidget, &outptr_QObject, &outptr_QPaintDevice)
	ret := newQCheckBox(outptr_QCheckBox, outptr_QAbstractButton, outptr_QWidget, outptr_QObject, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

// NewQCheckBox2 constructs a new QCheckBox object.
func NewQCheckBox2() *QCheckBox {
	var outptr_QCheckBox *C.QCheckBox = nil
	var outptr_QAbstractButton *C.QAbstractButton = nil
	var outptr_QWidget *C.QWidget = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QCheckBox_new2(&outptr_QCheckBox, &outptr_QAbstractButton, &outptr_QWidget, &outptr_QObject, &outptr_QPaintDevice)
	ret := newQCheckBox(outptr_QCheckBox, outptr_QAbstractButton, outptr_QWidget, outptr_QObject, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

// NewQCheckBox3 constructs a new QCheckBox object.
func NewQCheckBox3(text string) *QCheckBox {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QCheckBox *C.QCheckBox = nil
	var outptr_QAbstractButton *C.QAbstractButton = nil
	var outptr_QWidget *C.QWidget = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QCheckBox_new3(text_ms, &outptr_QCheckBox, &outptr_QAbstractButton, &outptr_QWidget, &outptr_QObject, &outptr_QPaintDevice)
	ret := newQCheckBox(outptr_QCheckBox, outptr_QAbstractButton, outptr_QWidget, outptr_QObject, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

// NewQCheckBox4 constructs a new QCheckBox object.
func NewQCheckBox4(text string, parent *QWidget) *QCheckBox {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QCheckBox *C.QCheckBox = nil
	var outptr_QAbstractButton *C.QAbstractButton = nil
	var outptr_QWidget *C.QWidget = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QCheckBox_new4(text_ms, parent.cPointer(), &outptr_QCheckBox, &outptr_QAbstractButton, &outptr_QWidget, &outptr_QObject, &outptr_QPaintDevice)
	ret := newQCheckBox(outptr_QCheckBox, outptr_QAbstractButton, outptr_QWidget, outptr_QObject, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

func (this *QCheckBox) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QCheckBox_MetaObject(this.h)))
}

func (this *QCheckBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QCheckBox_Metacast(this.h, param1_Cstring))
}

func QCheckBox_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QCheckBox_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCheckBox_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QCheckBox_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCheckBox) SizeHint() *QSize {
	_ret := C.QCheckBox_SizeHint(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCheckBox) MinimumSizeHint() *QSize {
	_ret := C.QCheckBox_MinimumSizeHint(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCheckBox) SetTristate() {
	C.QCheckBox_SetTristate(this.h)
}

func (this *QCheckBox) IsTristate() bool {
	return (bool)(C.QCheckBox_IsTristate(this.h))
}

func (this *QCheckBox) CheckState() CheckState {
	return (CheckState)(C.QCheckBox_CheckState(this.h))
}

func (this *QCheckBox) SetCheckState(state CheckState) {
	C.QCheckBox_SetCheckState(this.h, (C.int)(state))
}

func (this *QCheckBox) StateChanged(param1 int) {
	C.QCheckBox_StateChanged(this.h, (C.int)(param1))
}
func (this *QCheckBox) OnStateChanged(slot func(param1 int)) {
	C.QCheckBox_connect_StateChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_StateChanged
func miqt_exec_callback_QCheckBox_StateChanged(cb C.intptr_t, param1 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func QCheckBox_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCheckBox_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCheckBox_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCheckBox_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCheckBox_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCheckBox_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCheckBox_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCheckBox_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCheckBox) SetTristate1(y bool) {
	C.QCheckBox_SetTristate1(this.h, (C.bool)(y))
}

func (this *QCheckBox) callVirtualBase_SizeHint() *QSize {

	_ret := C.QCheckBox_virtualbase_SizeHint(unsafe.Pointer(this.h))
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCheckBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	C.QCheckBox_override_virtual_SizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_SizeHint
func miqt_exec_callback_QCheckBox_SizeHint(self *C.QCheckBox, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QCheckBox) callVirtualBase_MinimumSizeHint() *QSize {

	_ret := C.QCheckBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h))
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCheckBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	C.QCheckBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MinimumSizeHint
func miqt_exec_callback_QCheckBox_MinimumSizeHint(self *C.QCheckBox, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QCheckBox) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(C.QCheckBox_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QCheckBox) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	C.QCheckBox_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_Event
func miqt_exec_callback_QCheckBox_Event(self *C.QCheckBox, cb C.intptr_t, e *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(e))

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QCheckBox) callVirtualBase_HitButton(pos *QPoint) bool {

	return (bool)(C.QCheckBox_virtualbase_HitButton(unsafe.Pointer(this.h), pos.cPointer()))

}
func (this *QCheckBox) OnHitButton(slot func(super func(pos *QPoint) bool, pos *QPoint) bool) {
	C.QCheckBox_override_virtual_HitButton(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_HitButton
func miqt_exec_callback_QCheckBox_HitButton(self *C.QCheckBox, cb C.intptr_t, pos *C.QPoint) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos *QPoint) bool, pos *QPoint) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPoint(unsafe.Pointer(pos))

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_HitButton, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QCheckBox) callVirtualBase_CheckStateSet() {

	C.QCheckBox_virtualbase_CheckStateSet(unsafe.Pointer(this.h))

}
func (this *QCheckBox) OnCheckStateSet(slot func(super func())) {
	C.QCheckBox_override_virtual_CheckStateSet(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_CheckStateSet
func miqt_exec_callback_QCheckBox_CheckStateSet(self *C.QCheckBox, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QCheckBox{h: self}).callVirtualBase_CheckStateSet)

}

func (this *QCheckBox) callVirtualBase_NextCheckState() {

	C.QCheckBox_virtualbase_NextCheckState(unsafe.Pointer(this.h))

}
func (this *QCheckBox) OnNextCheckState(slot func(super func())) {
	C.QCheckBox_override_virtual_NextCheckState(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_NextCheckState
func miqt_exec_callback_QCheckBox_NextCheckState(self *C.QCheckBox, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QCheckBox{h: self}).callVirtualBase_NextCheckState)

}

func (this *QCheckBox) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	C.QCheckBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCheckBox) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	C.QCheckBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_PaintEvent
func miqt_exec_callback_QCheckBox_PaintEvent(self *C.QCheckBox, cb C.intptr_t, param1 *C.QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPaintEvent(unsafe.Pointer(param1), nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	C.QCheckBox_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCheckBox) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	C.QCheckBox_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MouseMoveEvent
func miqt_exec_callback_QCheckBox_MouseMoveEvent(self *C.QCheckBox, cb C.intptr_t, param1 *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQMouseEvent(unsafe.Pointer(param1), nil, nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_KeyPressEvent(e *QKeyEvent) {

	C.QCheckBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnKeyPressEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	C.QCheckBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_KeyPressEvent
func miqt_exec_callback_QCheckBox_KeyPressEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQKeyEvent(unsafe.Pointer(e), nil, nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_KeyReleaseEvent(e *QKeyEvent) {

	C.QCheckBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnKeyReleaseEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	C.QCheckBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_KeyReleaseEvent
func miqt_exec_callback_QCheckBox_KeyReleaseEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQKeyEvent(unsafe.Pointer(e), nil, nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_MousePressEvent(e *QMouseEvent) {

	C.QCheckBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnMousePressEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	C.QCheckBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MousePressEvent
func miqt_exec_callback_QCheckBox_MousePressEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQMouseEvent(unsafe.Pointer(e), nil, nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	C.QCheckBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	C.QCheckBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MouseReleaseEvent
func miqt_exec_callback_QCheckBox_MouseReleaseEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQMouseEvent(unsafe.Pointer(e), nil, nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_FocusInEvent(e *QFocusEvent) {

	C.QCheckBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnFocusInEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	C.QCheckBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_FocusInEvent
func miqt_exec_callback_QCheckBox_FocusInEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQFocusEvent(unsafe.Pointer(e), nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_FocusOutEvent(e *QFocusEvent) {

	C.QCheckBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnFocusOutEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	C.QCheckBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_FocusOutEvent
func miqt_exec_callback_QCheckBox_FocusOutEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQFocusEvent(unsafe.Pointer(e), nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_ChangeEvent(e *QEvent) {

	C.QCheckBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	C.QCheckBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_ChangeEvent
func miqt_exec_callback_QCheckBox_ChangeEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(e))

	gofunc((&QCheckBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_TimerEvent(e *QTimerEvent) {

	C.QCheckBox_virtualbase_TimerEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnTimerEvent(slot func(super func(e *QTimerEvent), e *QTimerEvent)) {
	C.QCheckBox_override_virtual_TimerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_TimerEvent
func miqt_exec_callback_QCheckBox_TimerEvent(self *C.QCheckBox, cb C.intptr_t, e *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QTimerEvent), e *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQTimerEvent(unsafe.Pointer(e), nil)

	gofunc((&QCheckBox{h: self}).callVirtualBase_TimerEvent, slotval1)

}

// Delete this object from C++ memory.
func (this *QCheckBox) Delete() {
	C.QCheckBox_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCheckBox) GoGC() {
	runtime.SetFinalizer(this, func(this *QCheckBox) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
