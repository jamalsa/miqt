package qt

/*

#include "gen_qrasterwindow.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QRasterWindow struct {
	h          *C.QRasterWindow
	isSubclass bool
	*QPaintDeviceWindow
}

func (this *QRasterWindow) cPointer() *C.QRasterWindow {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QRasterWindow) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQRasterWindow constructs the type using only CGO pointers.
func newQRasterWindow(h *C.QRasterWindow, h_QPaintDeviceWindow *C.QPaintDeviceWindow, h_QWindow *C.QWindow, h_QObject *C.QObject, h_QSurface *C.QSurface, h_QPaintDevice *C.QPaintDevice) *QRasterWindow {
	if h == nil {
		return nil
	}
	return &QRasterWindow{h: h,
		QPaintDeviceWindow: newQPaintDeviceWindow(h_QPaintDeviceWindow, h_QWindow, h_QObject, h_QSurface, h_QPaintDevice)}
}

// UnsafeNewQRasterWindow constructs the type using only unsafe pointers.
func UnsafeNewQRasterWindow(h unsafe.Pointer, h_QPaintDeviceWindow unsafe.Pointer, h_QWindow unsafe.Pointer, h_QObject unsafe.Pointer, h_QSurface unsafe.Pointer, h_QPaintDevice unsafe.Pointer) *QRasterWindow {
	if h == nil {
		return nil
	}

	return &QRasterWindow{h: (*C.QRasterWindow)(h),
		QPaintDeviceWindow: UnsafeNewQPaintDeviceWindow(h_QPaintDeviceWindow, h_QWindow, h_QObject, h_QSurface, h_QPaintDevice)}
}

// NewQRasterWindow constructs a new QRasterWindow object.
func NewQRasterWindow() *QRasterWindow {
	var outptr_QRasterWindow *C.QRasterWindow = nil
	var outptr_QPaintDeviceWindow *C.QPaintDeviceWindow = nil
	var outptr_QWindow *C.QWindow = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QSurface *C.QSurface = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QRasterWindow_new(&outptr_QRasterWindow, &outptr_QPaintDeviceWindow, &outptr_QWindow, &outptr_QObject, &outptr_QSurface, &outptr_QPaintDevice)
	ret := newQRasterWindow(outptr_QRasterWindow, outptr_QPaintDeviceWindow, outptr_QWindow, outptr_QObject, outptr_QSurface, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

// NewQRasterWindow2 constructs a new QRasterWindow object.
func NewQRasterWindow2(parent *QWindow) *QRasterWindow {
	var outptr_QRasterWindow *C.QRasterWindow = nil
	var outptr_QPaintDeviceWindow *C.QPaintDeviceWindow = nil
	var outptr_QWindow *C.QWindow = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QSurface *C.QSurface = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QRasterWindow_new2(parent.cPointer(), &outptr_QRasterWindow, &outptr_QPaintDeviceWindow, &outptr_QWindow, &outptr_QObject, &outptr_QSurface, &outptr_QPaintDevice)
	ret := newQRasterWindow(outptr_QRasterWindow, outptr_QPaintDeviceWindow, outptr_QWindow, outptr_QObject, outptr_QSurface, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

func (this *QRasterWindow) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QRasterWindow_MetaObject(this.h)))
}

func (this *QRasterWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QRasterWindow_Metacast(this.h, param1_Cstring))
}

func QRasterWindow_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QRasterWindow_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QRasterWindow_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QRasterWindow_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QRasterWindow_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QRasterWindow_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QRasterWindow_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRasterWindow) callVirtualBase_Metric(metric QPaintDevice__PaintDeviceMetric) int {

	return (int)(C.QRasterWindow_virtualbase_Metric(unsafe.Pointer(this.h), (C.int)(metric)))

}
func (this *QRasterWindow) OnMetric(slot func(super func(metric QPaintDevice__PaintDeviceMetric) int, metric QPaintDevice__PaintDeviceMetric) int) {
	C.QRasterWindow_override_virtual_Metric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_Metric
func miqt_exec_callback_QRasterWindow_Metric(self *C.QRasterWindow, cb C.intptr_t, metric C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(metric QPaintDevice__PaintDeviceMetric) int, metric QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QPaintDevice__PaintDeviceMetric)(metric)

	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_Metric, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QRasterWindow) callVirtualBase_Redirected(param1 *QPoint) *QPaintDevice {

	return UnsafeNewQPaintDevice(unsafe.Pointer(C.QRasterWindow_virtualbase_Redirected(unsafe.Pointer(this.h), param1.cPointer())))
}
func (this *QRasterWindow) OnRedirected(slot func(super func(param1 *QPoint) *QPaintDevice, param1 *QPoint) *QPaintDevice) {
	C.QRasterWindow_override_virtual_Redirected(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_Redirected
func miqt_exec_callback_QRasterWindow_Redirected(self *C.QRasterWindow, cb C.intptr_t, param1 *C.QPoint) *C.QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPoint) *QPaintDevice, param1 *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPoint(unsafe.Pointer(param1))

	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QRasterWindow) callVirtualBase_PaintEvent(event *QPaintEvent) {

	C.QRasterWindow_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRasterWindow) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	C.QRasterWindow_override_virtual_PaintEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_PaintEvent
func miqt_exec_callback_QRasterWindow_PaintEvent(self *C.QRasterWindow, cb C.intptr_t, event *C.QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPaintEvent(unsafe.Pointer(event), nil)

	gofunc((&QRasterWindow{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QRasterWindow) callVirtualBase_ExposeEvent(param1 *QExposeEvent) {

	C.QRasterWindow_virtualbase_ExposeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRasterWindow) OnExposeEvent(slot func(super func(param1 *QExposeEvent), param1 *QExposeEvent)) {
	C.QRasterWindow_override_virtual_ExposeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_ExposeEvent
func miqt_exec_callback_QRasterWindow_ExposeEvent(self *C.QRasterWindow, cb C.intptr_t, param1 *C.QExposeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QExposeEvent), param1 *QExposeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQExposeEvent(unsafe.Pointer(param1), nil)

	gofunc((&QRasterWindow{h: self}).callVirtualBase_ExposeEvent, slotval1)

}

func (this *QRasterWindow) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(C.QRasterWindow_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QRasterWindow) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	C.QRasterWindow_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_Event
func miqt_exec_callback_QRasterWindow_Event(self *C.QRasterWindow, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QRasterWindow) Delete() {
	C.QRasterWindow_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QRasterWindow) GoGC() {
	runtime.SetFinalizer(this, func(this *QRasterWindow) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
