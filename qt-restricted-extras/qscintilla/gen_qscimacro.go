package qscintilla

/*

#include "gen_qscimacro.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QsciMacro struct {
	h          *C.QsciMacro
	isSubclass bool
	*qt.QObject
}

func (this *QsciMacro) cPointer() *C.QsciMacro {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QsciMacro) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQsciMacro constructs the type using only CGO pointers.
func newQsciMacro(h *C.QsciMacro, h_QObject *C.QObject) *QsciMacro {
	if h == nil {
		return nil
	}
	return &QsciMacro{h: h,
		QObject: qt.UnsafeNewQObject(unsafe.Pointer(h_QObject))}
}

// UnsafeNewQsciMacro constructs the type using only unsafe pointers.
func UnsafeNewQsciMacro(h unsafe.Pointer, h_QObject unsafe.Pointer) *QsciMacro {
	if h == nil {
		return nil
	}

	return &QsciMacro{h: (*C.QsciMacro)(h),
		QObject: qt.UnsafeNewQObject(h_QObject)}
}

// NewQsciMacro constructs a new QsciMacro object.
func NewQsciMacro(parent *QsciScintilla) *QsciMacro {
	var outptr_QsciMacro *C.QsciMacro = nil
	var outptr_QObject *C.QObject = nil

	C.QsciMacro_new(parent.cPointer(), &outptr_QsciMacro, &outptr_QObject)
	ret := newQsciMacro(outptr_QsciMacro, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQsciMacro2 constructs a new QsciMacro object.
func NewQsciMacro2(asc string, parent *QsciScintilla) *QsciMacro {
	asc_ms := C.struct_miqt_string{}
	asc_ms.data = C.CString(asc)
	asc_ms.len = C.size_t(len(asc))
	defer C.free(unsafe.Pointer(asc_ms.data))
	var outptr_QsciMacro *C.QsciMacro = nil
	var outptr_QObject *C.QObject = nil

	C.QsciMacro_new2(asc_ms, parent.cPointer(), &outptr_QsciMacro, &outptr_QObject)
	ret := newQsciMacro(outptr_QsciMacro, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QsciMacro) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QsciMacro_MetaObject(this.h)))
}

func (this *QsciMacro) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QsciMacro_Metacast(this.h, param1_Cstring))
}

func QsciMacro_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciMacro_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciMacro_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciMacro_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciMacro) Clear() {
	C.QsciMacro_Clear(this.h)
}

func (this *QsciMacro) Load(asc string) bool {
	asc_ms := C.struct_miqt_string{}
	asc_ms.data = C.CString(asc)
	asc_ms.len = C.size_t(len(asc))
	defer C.free(unsafe.Pointer(asc_ms.data))
	return (bool)(C.QsciMacro_Load(this.h, asc_ms))
}

func (this *QsciMacro) Save() string {
	var _ms C.struct_miqt_string = C.QsciMacro_Save(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciMacro) Play() {
	C.QsciMacro_Play(this.h)
}

func (this *QsciMacro) StartRecording() {
	C.QsciMacro_StartRecording(this.h)
}

func (this *QsciMacro) EndRecording() {
	C.QsciMacro_EndRecording(this.h)
}

func QsciMacro_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciMacro_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciMacro_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciMacro_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciMacro_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciMacro_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciMacro_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciMacro_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciMacro) callVirtualBase_Play() {

	C.QsciMacro_virtualbase_Play(unsafe.Pointer(this.h))

}
func (this *QsciMacro) OnPlay(slot func(super func())) {
	C.QsciMacro_override_virtual_Play(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_Play
func miqt_exec_callback_QsciMacro_Play(self *C.QsciMacro, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QsciMacro{h: self}).callVirtualBase_Play)

}

func (this *QsciMacro) callVirtualBase_StartRecording() {

	C.QsciMacro_virtualbase_StartRecording(unsafe.Pointer(this.h))

}
func (this *QsciMacro) OnStartRecording(slot func(super func())) {
	C.QsciMacro_override_virtual_StartRecording(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_StartRecording
func miqt_exec_callback_QsciMacro_StartRecording(self *C.QsciMacro, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QsciMacro{h: self}).callVirtualBase_StartRecording)

}

func (this *QsciMacro) callVirtualBase_EndRecording() {

	C.QsciMacro_virtualbase_EndRecording(unsafe.Pointer(this.h))

}
func (this *QsciMacro) OnEndRecording(slot func(super func())) {
	C.QsciMacro_override_virtual_EndRecording(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_EndRecording
func miqt_exec_callback_QsciMacro_EndRecording(self *C.QsciMacro, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QsciMacro{h: self}).callVirtualBase_EndRecording)

}

func (this *QsciMacro) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(C.QsciMacro_virtualbase_Event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QsciMacro) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	C.QsciMacro_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_Event
func miqt_exec_callback_QsciMacro_Event(self *C.QsciMacro, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QsciMacro{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QsciMacro) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(C.QsciMacro_virtualbase_EventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QsciMacro) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	C.QsciMacro_override_virtual_EventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_EventFilter
func miqt_exec_callback_QsciMacro_EventFilter(self *C.QsciMacro, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))
	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QsciMacro{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QsciMacro) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	C.QsciMacro_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QsciMacro) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	C.QsciMacro_override_virtual_TimerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_TimerEvent
func miqt_exec_callback_QsciMacro_TimerEvent(self *C.QsciMacro, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event), nil)

	gofunc((&QsciMacro{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QsciMacro) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	C.QsciMacro_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QsciMacro) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	C.QsciMacro_override_virtual_ChildEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_ChildEvent
func miqt_exec_callback_QsciMacro_ChildEvent(self *C.QsciMacro, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event), nil)

	gofunc((&QsciMacro{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QsciMacro) callVirtualBase_CustomEvent(event *qt.QEvent) {

	C.QsciMacro_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QsciMacro) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	C.QsciMacro_override_virtual_CustomEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_CustomEvent
func miqt_exec_callback_QsciMacro_CustomEvent(self *C.QsciMacro, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QsciMacro{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QsciMacro) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	C.QsciMacro_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QsciMacro) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	C.QsciMacro_override_virtual_ConnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_ConnectNotify
func miqt_exec_callback_QsciMacro_ConnectNotify(self *C.QsciMacro, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QsciMacro{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QsciMacro) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	C.QsciMacro_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QsciMacro) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	C.QsciMacro_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciMacro_DisconnectNotify
func miqt_exec_callback_QsciMacro_DisconnectNotify(self *C.QsciMacro, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QsciMacro{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QsciMacro) Delete() {
	C.QsciMacro_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QsciMacro) GoGC() {
	runtime.SetFinalizer(this, func(this *QsciMacro) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
