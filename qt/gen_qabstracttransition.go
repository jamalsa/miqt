package qt

/*

#include "gen_qabstracttransition.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QAbstractTransition__TransitionType int

const (
	QAbstractTransition__ExternalTransition QAbstractTransition__TransitionType = 0
	QAbstractTransition__InternalTransition QAbstractTransition__TransitionType = 1
)

type QAbstractTransition struct {
	h          *C.QAbstractTransition
	isSubclass bool
	*QObject
}

func (this *QAbstractTransition) cPointer() *C.QAbstractTransition {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QAbstractTransition) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQAbstractTransition constructs the type using only CGO pointers.
func newQAbstractTransition(h *C.QAbstractTransition, h_QObject *C.QObject) *QAbstractTransition {
	if h == nil {
		return nil
	}
	return &QAbstractTransition{h: h,
		QObject: newQObject(h_QObject)}
}

// UnsafeNewQAbstractTransition constructs the type using only unsafe pointers.
func UnsafeNewQAbstractTransition(h unsafe.Pointer, h_QObject unsafe.Pointer) *QAbstractTransition {
	if h == nil {
		return nil
	}

	return &QAbstractTransition{h: (*C.QAbstractTransition)(h),
		QObject: UnsafeNewQObject(h_QObject)}
}

func (this *QAbstractTransition) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QAbstractTransition_MetaObject(this.h)))
}

func (this *QAbstractTransition) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QAbstractTransition_Metacast(this.h, param1_Cstring))
}

func QAbstractTransition_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractTransition_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractTransition_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractTransition_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractTransition) SourceState() *QState {
	return UnsafeNewQState(unsafe.Pointer(C.QAbstractTransition_SourceState(this.h)), nil, nil)
}

func (this *QAbstractTransition) TargetState() *QAbstractState {
	return UnsafeNewQAbstractState(unsafe.Pointer(C.QAbstractTransition_TargetState(this.h)), nil)
}

func (this *QAbstractTransition) SetTargetState(target *QAbstractState) {
	C.QAbstractTransition_SetTargetState(this.h, target.cPointer())
}

func (this *QAbstractTransition) TargetStates() []*QAbstractState {
	var _ma C.struct_miqt_array = C.QAbstractTransition_TargetStates(this.h)
	_ret := make([]*QAbstractState, int(_ma.len))
	_outCast := (*[0xffff]*C.QAbstractState)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = UnsafeNewQAbstractState(unsafe.Pointer(_outCast[i]), nil)
	}
	return _ret
}

func (this *QAbstractTransition) SetTargetStates(targets []*QAbstractState) {
	targets_CArray := (*[0xffff]*C.QAbstractState)(C.malloc(C.size_t(8 * len(targets))))
	defer C.free(unsafe.Pointer(targets_CArray))
	for i := range targets {
		targets_CArray[i] = targets[i].cPointer()
	}
	targets_ma := C.struct_miqt_array{len: C.size_t(len(targets)), data: unsafe.Pointer(targets_CArray)}
	C.QAbstractTransition_SetTargetStates(this.h, targets_ma)
}

func (this *QAbstractTransition) TransitionType() QAbstractTransition__TransitionType {
	return (QAbstractTransition__TransitionType)(C.QAbstractTransition_TransitionType(this.h))
}

func (this *QAbstractTransition) SetTransitionType(typeVal QAbstractTransition__TransitionType) {
	C.QAbstractTransition_SetTransitionType(this.h, (C.int)(typeVal))
}

func (this *QAbstractTransition) Machine() *QStateMachine {
	return UnsafeNewQStateMachine(unsafe.Pointer(C.QAbstractTransition_Machine(this.h)), nil, nil, nil)
}

func (this *QAbstractTransition) AddAnimation(animation *QAbstractAnimation) {
	C.QAbstractTransition_AddAnimation(this.h, animation.cPointer())
}

func (this *QAbstractTransition) RemoveAnimation(animation *QAbstractAnimation) {
	C.QAbstractTransition_RemoveAnimation(this.h, animation.cPointer())
}

func (this *QAbstractTransition) Animations() []*QAbstractAnimation {
	var _ma C.struct_miqt_array = C.QAbstractTransition_Animations(this.h)
	_ret := make([]*QAbstractAnimation, int(_ma.len))
	_outCast := (*[0xffff]*C.QAbstractAnimation)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = UnsafeNewQAbstractAnimation(unsafe.Pointer(_outCast[i]), nil)
	}
	return _ret
}

func QAbstractTransition_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractTransition_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractTransition_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractTransition_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractTransition_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractTransition_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractTransition_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractTransition_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QAbstractTransition) Delete() {
	C.QAbstractTransition_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAbstractTransition) GoGC() {
	runtime.SetFinalizer(this, func(this *QAbstractTransition) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
