package qt

/*

#include "gen_qelapsedtimer.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QElapsedTimer__ClockType int

const (
	QElapsedTimer__ClockType__SystemTime         QElapsedTimer__ClockType = 0
	QElapsedTimer__ClockType__MonotonicClock     QElapsedTimer__ClockType = 1
	QElapsedTimer__ClockType__TickCounter        QElapsedTimer__ClockType = 2
	QElapsedTimer__ClockType__MachAbsoluteTime   QElapsedTimer__ClockType = 3
	QElapsedTimer__ClockType__PerformanceCounter QElapsedTimer__ClockType = 4
)

type QElapsedTimer struct {
	h *C.QElapsedTimer
}

func (this *QElapsedTimer) cPointer() *C.QElapsedTimer {
	if this == nil {
		return nil
	}
	return this.h
}

func newQElapsedTimer(h *C.QElapsedTimer) *QElapsedTimer {
	if h == nil {
		return nil
	}
	return &QElapsedTimer{h: h}
}

func newQElapsedTimer_U(h unsafe.Pointer) *QElapsedTimer {
	return newQElapsedTimer((*C.QElapsedTimer)(h))
}

// NewQElapsedTimer constructs a new QElapsedTimer object.
func NewQElapsedTimer() *QElapsedTimer {
	ret := C.QElapsedTimer_new()
	return newQElapsedTimer(ret)
}

func QElapsedTimer_ClockType() QElapsedTimer__ClockType {
	_ret := C.QElapsedTimer_ClockType()
	return (QElapsedTimer__ClockType)(_ret)
}

func QElapsedTimer_IsMonotonic() bool {
	_ret := C.QElapsedTimer_IsMonotonic()
	return (bool)(_ret)
}

func (this *QElapsedTimer) Start() {
	C.QElapsedTimer_Start(this.h)
}

func (this *QElapsedTimer) Restart() int64 {
	_ret := C.QElapsedTimer_Restart(this.h)
	return (int64)(_ret)
}

func (this *QElapsedTimer) Invalidate() {
	C.QElapsedTimer_Invalidate(this.h)
}

func (this *QElapsedTimer) IsValid() bool {
	_ret := C.QElapsedTimer_IsValid(this.h)
	return (bool)(_ret)
}

func (this *QElapsedTimer) NsecsElapsed() int64 {
	_ret := C.QElapsedTimer_NsecsElapsed(this.h)
	return (int64)(_ret)
}

func (this *QElapsedTimer) Elapsed() int64 {
	_ret := C.QElapsedTimer_Elapsed(this.h)
	return (int64)(_ret)
}

func (this *QElapsedTimer) HasExpired(timeout int64) bool {
	_ret := C.QElapsedTimer_HasExpired(this.h, (C.longlong)(timeout))
	return (bool)(_ret)
}

func (this *QElapsedTimer) MsecsSinceReference() int64 {
	_ret := C.QElapsedTimer_MsecsSinceReference(this.h)
	return (int64)(_ret)
}

func (this *QElapsedTimer) MsecsTo(other *QElapsedTimer) int64 {
	_ret := C.QElapsedTimer_MsecsTo(this.h, other.cPointer())
	return (int64)(_ret)
}

func (this *QElapsedTimer) SecsTo(other *QElapsedTimer) int64 {
	_ret := C.QElapsedTimer_SecsTo(this.h, other.cPointer())
	return (int64)(_ret)
}

func (this *QElapsedTimer) OperatorEqual(other *QElapsedTimer) bool {
	_ret := C.QElapsedTimer_OperatorEqual(this.h, other.cPointer())
	return (bool)(_ret)
}

func (this *QElapsedTimer) OperatorNotEqual(other *QElapsedTimer) bool {
	_ret := C.QElapsedTimer_OperatorNotEqual(this.h, other.cPointer())
	return (bool)(_ret)
}

// Delete this object from C++ memory.
func (this *QElapsedTimer) Delete() {
	C.QElapsedTimer_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QElapsedTimer) GoGC() {
	runtime.SetFinalizer(this, func(this *QElapsedTimer) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
