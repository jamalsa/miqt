package qscintilla

/*

#include "gen_qsciprinter.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"github.com/mappu/miqt/qt/printsupport"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QsciPrinter struct {
	h          *C.QsciPrinter
	isSubclass bool
	*printsupport.QPrinter
}

func (this *QsciPrinter) cPointer() *C.QsciPrinter {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QsciPrinter) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQsciPrinter constructs the type using only CGO pointers.
func newQsciPrinter(h *C.QsciPrinter, h_QPrinter *C.QPrinter, h_QPagedPaintDevice *C.QPagedPaintDevice, h_QPaintDevice *C.QPaintDevice) *QsciPrinter {
	if h == nil {
		return nil
	}
	return &QsciPrinter{h: h,
		QPrinter: printsupport.UnsafeNewQPrinter(unsafe.Pointer(h_QPrinter), unsafe.Pointer(h_QPagedPaintDevice), unsafe.Pointer(h_QPaintDevice))}
}

// UnsafeNewQsciPrinter constructs the type using only unsafe pointers.
func UnsafeNewQsciPrinter(h unsafe.Pointer, h_QPrinter unsafe.Pointer, h_QPagedPaintDevice unsafe.Pointer, h_QPaintDevice unsafe.Pointer) *QsciPrinter {
	if h == nil {
		return nil
	}

	return &QsciPrinter{h: (*C.QsciPrinter)(h),
		QPrinter: printsupport.UnsafeNewQPrinter(h_QPrinter, h_QPagedPaintDevice, h_QPaintDevice)}
}

// NewQsciPrinter constructs a new QsciPrinter object.
func NewQsciPrinter() *QsciPrinter {
	var outptr_QsciPrinter *C.QsciPrinter = nil
	var outptr_QPrinter *C.QPrinter = nil
	var outptr_QPagedPaintDevice *C.QPagedPaintDevice = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QsciPrinter_new(&outptr_QsciPrinter, &outptr_QPrinter, &outptr_QPagedPaintDevice, &outptr_QPaintDevice)
	ret := newQsciPrinter(outptr_QsciPrinter, outptr_QPrinter, outptr_QPagedPaintDevice, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

// NewQsciPrinter2 constructs a new QsciPrinter object.
func NewQsciPrinter2(mode printsupport.QPrinter__PrinterMode) *QsciPrinter {
	var outptr_QsciPrinter *C.QsciPrinter = nil
	var outptr_QPrinter *C.QPrinter = nil
	var outptr_QPagedPaintDevice *C.QPagedPaintDevice = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QsciPrinter_new2((C.int)(mode), &outptr_QsciPrinter, &outptr_QPrinter, &outptr_QPagedPaintDevice, &outptr_QPaintDevice)
	ret := newQsciPrinter(outptr_QsciPrinter, outptr_QPrinter, outptr_QPagedPaintDevice, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

func (this *QsciPrinter) FormatPage(painter *qt.QPainter, drawing bool, area *qt.QRect, pagenr int) {
	C.QsciPrinter_FormatPage(this.h, (*C.QPainter)(painter.UnsafePointer()), (C.bool)(drawing), (*C.QRect)(area.UnsafePointer()), (C.int)(pagenr))
}

func (this *QsciPrinter) Magnification() int {
	return (int)(C.QsciPrinter_Magnification(this.h))
}

func (this *QsciPrinter) SetMagnification(magnification int) {
	C.QsciPrinter_SetMagnification(this.h, (C.int)(magnification))
}

func (this *QsciPrinter) PrintRange(qsb *QsciScintillaBase, painter *qt.QPainter, from int, to int) int {
	return (int)(C.QsciPrinter_PrintRange(this.h, qsb.cPointer(), (*C.QPainter)(painter.UnsafePointer()), (C.int)(from), (C.int)(to)))
}

func (this *QsciPrinter) PrintRange2(qsb *QsciScintillaBase, from int, to int) int {
	return (int)(C.QsciPrinter_PrintRange2(this.h, qsb.cPointer(), (C.int)(from), (C.int)(to)))
}

func (this *QsciPrinter) WrapMode() QsciScintilla__WrapMode {
	return (QsciScintilla__WrapMode)(C.QsciPrinter_WrapMode(this.h))
}

func (this *QsciPrinter) SetWrapMode(wmode QsciScintilla__WrapMode) {
	C.QsciPrinter_SetWrapMode(this.h, (C.int)(wmode))
}

func (this *QsciPrinter) callVirtualBase_FormatPage(painter *qt.QPainter, drawing bool, area *qt.QRect, pagenr int) {

	C.QsciPrinter_virtualbase_FormatPage(unsafe.Pointer(this.h), (*C.QPainter)(painter.UnsafePointer()), (C.bool)(drawing), (*C.QRect)(area.UnsafePointer()), (C.int)(pagenr))

}
func (this *QsciPrinter) OnFormatPage(slot func(super func(painter *qt.QPainter, drawing bool, area *qt.QRect, pagenr int), painter *qt.QPainter, drawing bool, area *qt.QRect, pagenr int)) {
	C.QsciPrinter_override_virtual_FormatPage(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_FormatPage
func miqt_exec_callback_QsciPrinter_FormatPage(self *C.QsciPrinter, cb C.intptr_t, painter *C.QPainter, drawing C.bool, area *C.QRect, pagenr C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter, drawing bool, area *qt.QRect, pagenr int), painter *qt.QPainter, drawing bool, area *qt.QRect, pagenr int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval2 := (bool)(drawing)

	slotval3 := qt.UnsafeNewQRect(unsafe.Pointer(area))
	slotval4 := (int)(pagenr)

	gofunc((&QsciPrinter{h: self}).callVirtualBase_FormatPage, slotval1, slotval2, slotval3, slotval4)

}

func (this *QsciPrinter) callVirtualBase_SetMagnification(magnification int) {

	C.QsciPrinter_virtualbase_SetMagnification(unsafe.Pointer(this.h), (C.int)(magnification))

}
func (this *QsciPrinter) OnSetMagnification(slot func(super func(magnification int), magnification int)) {
	C.QsciPrinter_override_virtual_SetMagnification(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_SetMagnification
func miqt_exec_callback_QsciPrinter_SetMagnification(self *C.QsciPrinter, cb C.intptr_t, magnification C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(magnification int), magnification int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(magnification)

	gofunc((&QsciPrinter{h: self}).callVirtualBase_SetMagnification, slotval1)

}

func (this *QsciPrinter) callVirtualBase_PrintRange(qsb *QsciScintillaBase, painter *qt.QPainter, from int, to int) int {

	return (int)(C.QsciPrinter_virtualbase_PrintRange(unsafe.Pointer(this.h), qsb.cPointer(), (*C.QPainter)(painter.UnsafePointer()), (C.int)(from), (C.int)(to)))

}
func (this *QsciPrinter) OnPrintRange(slot func(super func(qsb *QsciScintillaBase, painter *qt.QPainter, from int, to int) int, qsb *QsciScintillaBase, painter *qt.QPainter, from int, to int) int) {
	C.QsciPrinter_override_virtual_PrintRange(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_PrintRange
func miqt_exec_callback_QsciPrinter_PrintRange(self *C.QsciPrinter, cb C.intptr_t, qsb *C.QsciScintillaBase, painter *C.QPainter, from C.int, to C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(qsb *QsciScintillaBase, painter *qt.QPainter, from int, to int) int, qsb *QsciScintillaBase, painter *qt.QPainter, from int, to int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQsciScintillaBase(unsafe.Pointer(qsb), nil, nil, nil, nil, nil)
	slotval2 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval3 := (int)(from)

	slotval4 := (int)(to)

	virtualReturn := gofunc((&QsciPrinter{h: self}).callVirtualBase_PrintRange, slotval1, slotval2, slotval3, slotval4)

	return (C.int)(virtualReturn)

}

func (this *QsciPrinter) callVirtualBase_PrintRange2(qsb *QsciScintillaBase, from int, to int) int {

	return (int)(C.QsciPrinter_virtualbase_PrintRange2(unsafe.Pointer(this.h), qsb.cPointer(), (C.int)(from), (C.int)(to)))

}
func (this *QsciPrinter) OnPrintRange2(slot func(super func(qsb *QsciScintillaBase, from int, to int) int, qsb *QsciScintillaBase, from int, to int) int) {
	C.QsciPrinter_override_virtual_PrintRange2(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_PrintRange2
func miqt_exec_callback_QsciPrinter_PrintRange2(self *C.QsciPrinter, cb C.intptr_t, qsb *C.QsciScintillaBase, from C.int, to C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(qsb *QsciScintillaBase, from int, to int) int, qsb *QsciScintillaBase, from int, to int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQsciScintillaBase(unsafe.Pointer(qsb), nil, nil, nil, nil, nil)
	slotval2 := (int)(from)

	slotval3 := (int)(to)

	virtualReturn := gofunc((&QsciPrinter{h: self}).callVirtualBase_PrintRange2, slotval1, slotval2, slotval3)

	return (C.int)(virtualReturn)

}

func (this *QsciPrinter) callVirtualBase_SetWrapMode(wmode QsciScintilla__WrapMode) {

	C.QsciPrinter_virtualbase_SetWrapMode(unsafe.Pointer(this.h), (C.int)(wmode))

}
func (this *QsciPrinter) OnSetWrapMode(slot func(super func(wmode QsciScintilla__WrapMode), wmode QsciScintilla__WrapMode)) {
	C.QsciPrinter_override_virtual_SetWrapMode(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_SetWrapMode
func miqt_exec_callback_QsciPrinter_SetWrapMode(self *C.QsciPrinter, cb C.intptr_t, wmode C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(wmode QsciScintilla__WrapMode), wmode QsciScintilla__WrapMode))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QsciScintilla__WrapMode)(wmode)

	gofunc((&QsciPrinter{h: self}).callVirtualBase_SetWrapMode, slotval1)

}

func (this *QsciPrinter) callVirtualBase_DevType() int {

	return (int)(C.QsciPrinter_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QsciPrinter) OnDevType(slot func(super func() int) int) {
	C.QsciPrinter_override_virtual_DevType(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_DevType
func miqt_exec_callback_QsciPrinter_DevType(self *C.QsciPrinter, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciPrinter{h: self}).callVirtualBase_DevType)

	return (C.int)(virtualReturn)

}

func (this *QsciPrinter) callVirtualBase_SetPageSize(pageSize qt.QPagedPaintDevice__PageSize) {

	C.QsciPrinter_virtualbase_SetPageSize(unsafe.Pointer(this.h), (C.int)(pageSize))

}
func (this *QsciPrinter) OnSetPageSize(slot func(super func(pageSize qt.QPagedPaintDevice__PageSize), pageSize qt.QPagedPaintDevice__PageSize)) {
	C.QsciPrinter_override_virtual_SetPageSize(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_SetPageSize
func miqt_exec_callback_QsciPrinter_SetPageSize(self *C.QsciPrinter, cb C.intptr_t, pageSize C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pageSize qt.QPagedPaintDevice__PageSize), pageSize qt.QPagedPaintDevice__PageSize))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QPagedPaintDevice__PageSize)(pageSize)

	gofunc((&QsciPrinter{h: self}).callVirtualBase_SetPageSize, slotval1)

}

func (this *QsciPrinter) callVirtualBase_SetPageSizeMM(size *qt.QSizeF) {

	C.QsciPrinter_virtualbase_SetPageSizeMM(unsafe.Pointer(this.h), (*C.QSizeF)(size.UnsafePointer()))

}
func (this *QsciPrinter) OnSetPageSizeMM(slot func(super func(size *qt.QSizeF), size *qt.QSizeF)) {
	C.QsciPrinter_override_virtual_SetPageSizeMM(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_SetPageSizeMM
func miqt_exec_callback_QsciPrinter_SetPageSizeMM(self *C.QsciPrinter, cb C.intptr_t, size *C.QSizeF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size *qt.QSizeF), size *qt.QSizeF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQSizeF(unsafe.Pointer(size))

	gofunc((&QsciPrinter{h: self}).callVirtualBase_SetPageSizeMM, slotval1)

}

func (this *QsciPrinter) callVirtualBase_NewPage() bool {

	return (bool)(C.QsciPrinter_virtualbase_NewPage(unsafe.Pointer(this.h)))

}
func (this *QsciPrinter) OnNewPage(slot func(super func() bool) bool) {
	C.QsciPrinter_override_virtual_NewPage(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_NewPage
func miqt_exec_callback_QsciPrinter_NewPage(self *C.QsciPrinter, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciPrinter{h: self}).callVirtualBase_NewPage)

	return (C.bool)(virtualReturn)

}

func (this *QsciPrinter) callVirtualBase_PaintEngine() *qt.QPaintEngine {

	return qt.UnsafeNewQPaintEngine(unsafe.Pointer(C.QsciPrinter_virtualbase_PaintEngine(unsafe.Pointer(this.h))))
}
func (this *QsciPrinter) OnPaintEngine(slot func(super func() *qt.QPaintEngine) *qt.QPaintEngine) {
	C.QsciPrinter_override_virtual_PaintEngine(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_PaintEngine
func miqt_exec_callback_QsciPrinter_PaintEngine(self *C.QsciPrinter, cb C.intptr_t) *C.QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPaintEngine) *qt.QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciPrinter{h: self}).callVirtualBase_PaintEngine)

	return (*C.QPaintEngine)(virtualReturn.UnsafePointer())

}

func (this *QsciPrinter) callVirtualBase_SetMargins(m *qt.QPagedPaintDevice__Margins) {

	C.QsciPrinter_virtualbase_SetMargins(unsafe.Pointer(this.h), (*C.QPagedPaintDevice__Margins)(m.UnsafePointer()))

}
func (this *QsciPrinter) OnSetMargins(slot func(super func(m *qt.QPagedPaintDevice__Margins), m *qt.QPagedPaintDevice__Margins)) {
	C.QsciPrinter_override_virtual_SetMargins(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_SetMargins
func miqt_exec_callback_QsciPrinter_SetMargins(self *C.QsciPrinter, cb C.intptr_t, m *C.QPagedPaintDevice__Margins) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(m *qt.QPagedPaintDevice__Margins), m *qt.QPagedPaintDevice__Margins))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPagedPaintDevice__Margins(unsafe.Pointer(m))

	gofunc((&QsciPrinter{h: self}).callVirtualBase_SetMargins, slotval1)

}

func (this *QsciPrinter) callVirtualBase_Metric(param1 qt.QPaintDevice__PaintDeviceMetric) int {

	return (int)(C.QsciPrinter_virtualbase_Metric(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QsciPrinter) OnMetric(slot func(super func(param1 qt.QPaintDevice__PaintDeviceMetric) int, param1 qt.QPaintDevice__PaintDeviceMetric) int) {
	C.QsciPrinter_override_virtual_Metric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciPrinter_Metric
func miqt_exec_callback_QsciPrinter_Metric(self *C.QsciPrinter, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt.QPaintDevice__PaintDeviceMetric) int, param1 qt.QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QPaintDevice__PaintDeviceMetric)(param1)

	virtualReturn := gofunc((&QsciPrinter{h: self}).callVirtualBase_Metric, slotval1)

	return (C.int)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QsciPrinter) Delete() {
	C.QsciPrinter_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QsciPrinter) GoGC() {
	runtime.SetFinalizer(this, func(this *QsciPrinter) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
