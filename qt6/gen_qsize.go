package qt6

/*

#include "gen_qsize.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QSize struct {
	h          *C.QSize
	isSubclass bool
}

func (this *QSize) cPointer() *C.QSize {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QSize) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQSize constructs the type using only CGO pointers.
func newQSize(h *C.QSize) *QSize {
	if h == nil {
		return nil
	}
	return &QSize{h: h}
}

// UnsafeNewQSize constructs the type using only unsafe pointers.
func UnsafeNewQSize(h unsafe.Pointer) *QSize {
	if h == nil {
		return nil
	}

	return &QSize{h: (*C.QSize)(h)}
}

// NewQSize constructs a new QSize object.
func NewQSize() *QSize {
	var outptr_QSize *C.QSize = nil

	C.QSize_new(&outptr_QSize)
	ret := newQSize(outptr_QSize)
	ret.isSubclass = true
	return ret
}

// NewQSize2 constructs a new QSize object.
func NewQSize2(w int, h int) *QSize {
	var outptr_QSize *C.QSize = nil

	C.QSize_new2((C.int)(w), (C.int)(h), &outptr_QSize)
	ret := newQSize(outptr_QSize)
	ret.isSubclass = true
	return ret
}

// NewQSize3 constructs a new QSize object.
func NewQSize3(param1 *QSize) *QSize {
	var outptr_QSize *C.QSize = nil

	C.QSize_new3(param1.cPointer(), &outptr_QSize)
	ret := newQSize(outptr_QSize)
	ret.isSubclass = true
	return ret
}

func (this *QSize) IsNull() bool {
	return (bool)(C.QSize_IsNull(this.h))
}

func (this *QSize) IsEmpty() bool {
	return (bool)(C.QSize_IsEmpty(this.h))
}

func (this *QSize) IsValid() bool {
	return (bool)(C.QSize_IsValid(this.h))
}

func (this *QSize) Width() int {
	return (int)(C.QSize_Width(this.h))
}

func (this *QSize) Height() int {
	return (int)(C.QSize_Height(this.h))
}

func (this *QSize) SetWidth(w int) {
	C.QSize_SetWidth(this.h, (C.int)(w))
}

func (this *QSize) SetHeight(h int) {
	C.QSize_SetHeight(this.h, (C.int)(h))
}

func (this *QSize) Transpose() {
	C.QSize_Transpose(this.h)
}

func (this *QSize) Transposed() *QSize {
	_ret := C.QSize_Transposed(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) Scale(w int, h int, mode AspectRatioMode) {
	C.QSize_Scale(this.h, (C.int)(w), (C.int)(h), (C.int)(mode))
}

func (this *QSize) Scale2(s *QSize, mode AspectRatioMode) {
	C.QSize_Scale2(this.h, s.cPointer(), (C.int)(mode))
}

func (this *QSize) Scaled(w int, h int, mode AspectRatioMode) *QSize {
	_ret := C.QSize_Scaled(this.h, (C.int)(w), (C.int)(h), (C.int)(mode))
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) Scaled2(s *QSize, mode AspectRatioMode) *QSize {
	_ret := C.QSize_Scaled2(this.h, s.cPointer(), (C.int)(mode))
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) ExpandedTo(param1 *QSize) *QSize {
	_ret := C.QSize_ExpandedTo(this.h, param1.cPointer())
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) BoundedTo(param1 *QSize) *QSize {
	_ret := C.QSize_BoundedTo(this.h, param1.cPointer())
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) GrownBy(m QMargins) *QSize {
	_ret := C.QSize_GrownBy(this.h, m.cPointer())
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) ShrunkBy(m QMargins) *QSize {
	_ret := C.QSize_ShrunkBy(this.h, m.cPointer())
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) OperatorPlusAssign(param1 *QSize) *QSize {
	return UnsafeNewQSize(unsafe.Pointer(C.QSize_OperatorPlusAssign(this.h, param1.cPointer())))
}

func (this *QSize) OperatorMinusAssign(param1 *QSize) *QSize {
	return UnsafeNewQSize(unsafe.Pointer(C.QSize_OperatorMinusAssign(this.h, param1.cPointer())))
}

func (this *QSize) OperatorMultiplyAssign(c float64) *QSize {
	return UnsafeNewQSize(unsafe.Pointer(C.QSize_OperatorMultiplyAssign(this.h, (C.double)(c))))
}

func (this *QSize) OperatorDivideAssign(c float64) *QSize {
	return UnsafeNewQSize(unsafe.Pointer(C.QSize_OperatorDivideAssign(this.h, (C.double)(c))))
}

func (this *QSize) ToSizeF() *QSizeF {
	_ret := C.QSize_ToSizeF(this.h)
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QSize) Delete() {
	C.QSize_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSize) GoGC() {
	runtime.SetFinalizer(this, func(this *QSize) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QSizeF struct {
	h          *C.QSizeF
	isSubclass bool
}

func (this *QSizeF) cPointer() *C.QSizeF {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QSizeF) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQSizeF constructs the type using only CGO pointers.
func newQSizeF(h *C.QSizeF) *QSizeF {
	if h == nil {
		return nil
	}
	return &QSizeF{h: h}
}

// UnsafeNewQSizeF constructs the type using only unsafe pointers.
func UnsafeNewQSizeF(h unsafe.Pointer) *QSizeF {
	if h == nil {
		return nil
	}

	return &QSizeF{h: (*C.QSizeF)(h)}
}

// NewQSizeF constructs a new QSizeF object.
func NewQSizeF() *QSizeF {
	var outptr_QSizeF *C.QSizeF = nil

	C.QSizeF_new(&outptr_QSizeF)
	ret := newQSizeF(outptr_QSizeF)
	ret.isSubclass = true
	return ret
}

// NewQSizeF2 constructs a new QSizeF object.
func NewQSizeF2(sz *QSize) *QSizeF {
	var outptr_QSizeF *C.QSizeF = nil

	C.QSizeF_new2(sz.cPointer(), &outptr_QSizeF)
	ret := newQSizeF(outptr_QSizeF)
	ret.isSubclass = true
	return ret
}

// NewQSizeF3 constructs a new QSizeF object.
func NewQSizeF3(w float64, h float64) *QSizeF {
	var outptr_QSizeF *C.QSizeF = nil

	C.QSizeF_new3((C.double)(w), (C.double)(h), &outptr_QSizeF)
	ret := newQSizeF(outptr_QSizeF)
	ret.isSubclass = true
	return ret
}

// NewQSizeF4 constructs a new QSizeF object.
func NewQSizeF4(param1 *QSizeF) *QSizeF {
	var outptr_QSizeF *C.QSizeF = nil

	C.QSizeF_new4(param1.cPointer(), &outptr_QSizeF)
	ret := newQSizeF(outptr_QSizeF)
	ret.isSubclass = true
	return ret
}

func (this *QSizeF) IsNull() bool {
	return (bool)(C.QSizeF_IsNull(this.h))
}

func (this *QSizeF) IsEmpty() bool {
	return (bool)(C.QSizeF_IsEmpty(this.h))
}

func (this *QSizeF) IsValid() bool {
	return (bool)(C.QSizeF_IsValid(this.h))
}

func (this *QSizeF) Width() float64 {
	return (float64)(C.QSizeF_Width(this.h))
}

func (this *QSizeF) Height() float64 {
	return (float64)(C.QSizeF_Height(this.h))
}

func (this *QSizeF) SetWidth(w float64) {
	C.QSizeF_SetWidth(this.h, (C.double)(w))
}

func (this *QSizeF) SetHeight(h float64) {
	C.QSizeF_SetHeight(this.h, (C.double)(h))
}

func (this *QSizeF) Transpose() {
	C.QSizeF_Transpose(this.h)
}

func (this *QSizeF) Transposed() *QSizeF {
	_ret := C.QSizeF_Transposed(this.h)
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) Scale(w float64, h float64, mode AspectRatioMode) {
	C.QSizeF_Scale(this.h, (C.double)(w), (C.double)(h), (C.int)(mode))
}

func (this *QSizeF) Scale2(s *QSizeF, mode AspectRatioMode) {
	C.QSizeF_Scale2(this.h, s.cPointer(), (C.int)(mode))
}

func (this *QSizeF) Scaled(w float64, h float64, mode AspectRatioMode) *QSizeF {
	_ret := C.QSizeF_Scaled(this.h, (C.double)(w), (C.double)(h), (C.int)(mode))
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) Scaled2(s *QSizeF, mode AspectRatioMode) *QSizeF {
	_ret := C.QSizeF_Scaled2(this.h, s.cPointer(), (C.int)(mode))
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) ExpandedTo(param1 *QSizeF) *QSizeF {
	_ret := C.QSizeF_ExpandedTo(this.h, param1.cPointer())
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) BoundedTo(param1 *QSizeF) *QSizeF {
	_ret := C.QSizeF_BoundedTo(this.h, param1.cPointer())
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) GrownBy(m QMarginsF) *QSizeF {
	_ret := C.QSizeF_GrownBy(this.h, m.cPointer())
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) ShrunkBy(m QMarginsF) *QSizeF {
	_ret := C.QSizeF_ShrunkBy(this.h, m.cPointer())
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) OperatorPlusAssign(param1 *QSizeF) *QSizeF {
	return UnsafeNewQSizeF(unsafe.Pointer(C.QSizeF_OperatorPlusAssign(this.h, param1.cPointer())))
}

func (this *QSizeF) OperatorMinusAssign(param1 *QSizeF) *QSizeF {
	return UnsafeNewQSizeF(unsafe.Pointer(C.QSizeF_OperatorMinusAssign(this.h, param1.cPointer())))
}

func (this *QSizeF) OperatorMultiplyAssign(c float64) *QSizeF {
	return UnsafeNewQSizeF(unsafe.Pointer(C.QSizeF_OperatorMultiplyAssign(this.h, (C.double)(c))))
}

func (this *QSizeF) OperatorDivideAssign(c float64) *QSizeF {
	return UnsafeNewQSizeF(unsafe.Pointer(C.QSizeF_OperatorDivideAssign(this.h, (C.double)(c))))
}

func (this *QSizeF) ToSize() *QSize {
	_ret := C.QSizeF_ToSize(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QSizeF) Delete() {
	C.QSizeF_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSizeF) GoGC() {
	runtime.SetFinalizer(this, func(this *QSizeF) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
