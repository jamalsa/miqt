package qt

/*

#include "gen_qfloat16.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type qfloat16 struct {
	h *C.qfloat16
}

func (this *qfloat16) cPointer() *C.qfloat16 {
	if this == nil {
		return nil
	}
	return this.h
}

func newqfloat16(h *C.qfloat16) *qfloat16 {
	if h == nil {
		return nil
	}
	return &qfloat16{h: h}
}

func newqfloat16_U(h unsafe.Pointer) *qfloat16 {
	return newqfloat16((*C.qfloat16)(h))
}

// Newqfloat16 constructs a new qfloat16 object.
func Newqfloat16() *qfloat16 {
	ret := C.qfloat16_new()
	return newqfloat16(ret)
}

// Newqfloat162 constructs a new qfloat16 object.
func Newqfloat162(f float32) *qfloat16 {
	ret := C.qfloat16_new2((C.float)(f))
	return newqfloat16(ret)
}

func (this *qfloat16) IsInf() bool {
	_ret := C.qfloat16_IsInf(this.h)
	return (bool)(_ret)
}

func (this *qfloat16) IsNaN() bool {
	_ret := C.qfloat16_IsNaN(this.h)
	return (bool)(_ret)
}

func (this *qfloat16) IsFinite() bool {
	_ret := C.qfloat16_IsFinite(this.h)
	return (bool)(_ret)
}

func (this *qfloat16) FpClassify() int {
	_ret := C.qfloat16_FpClassify(this.h)
	return (int)(_ret)
}

func (this *qfloat16) IsNormal() bool {
	_ret := C.qfloat16_IsNormal(this.h)
	return (bool)(_ret)
}

// Delete this object from C++ memory.
func (this *qfloat16) Delete() {
	C.qfloat16_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *qfloat16) GoGC() {
	runtime.SetFinalizer(this, func(this *qfloat16) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
