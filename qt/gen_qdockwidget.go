package qt

/*

#include "gen_qdockwidget.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QDockWidget__DockWidgetFeature int

const (
	QDockWidget__DockWidgetFeature__DockWidgetClosable         QDockWidget__DockWidgetFeature = 1
	QDockWidget__DockWidgetFeature__DockWidgetMovable          QDockWidget__DockWidgetFeature = 2
	QDockWidget__DockWidgetFeature__DockWidgetFloatable        QDockWidget__DockWidgetFeature = 4
	QDockWidget__DockWidgetFeature__DockWidgetVerticalTitleBar QDockWidget__DockWidgetFeature = 8
	QDockWidget__DockWidgetFeature__DockWidgetFeatureMask      QDockWidget__DockWidgetFeature = 15
	QDockWidget__DockWidgetFeature__AllDockWidgetFeatures      QDockWidget__DockWidgetFeature = 7
	QDockWidget__DockWidgetFeature__NoDockWidgetFeatures       QDockWidget__DockWidgetFeature = 0
	QDockWidget__DockWidgetFeature__Reserved                   QDockWidget__DockWidgetFeature = 255
)

type QDockWidget struct {
	h *C.QDockWidget
	*QWidget
}

func (this *QDockWidget) cPointer() *C.QDockWidget {
	if this == nil {
		return nil
	}
	return this.h
}

func newQDockWidget(h *C.QDockWidget) *QDockWidget {
	if h == nil {
		return nil
	}
	return &QDockWidget{h: h, QWidget: newQWidget_U(unsafe.Pointer(h))}
}

func newQDockWidget_U(h unsafe.Pointer) *QDockWidget {
	return newQDockWidget((*C.QDockWidget)(h))
}

// NewQDockWidget constructs a new QDockWidget object.
func NewQDockWidget(title string) *QDockWidget {
	title_ms := miqt_strdupg(title)
	defer C.free(title_ms)
	ret := C.QDockWidget_new((*C.struct_miqt_string)(title_ms))
	return newQDockWidget(ret)
}

// NewQDockWidget2 constructs a new QDockWidget object.
func NewQDockWidget2() *QDockWidget {
	ret := C.QDockWidget_new2()
	return newQDockWidget(ret)
}

// NewQDockWidget3 constructs a new QDockWidget object.
func NewQDockWidget3(title string, parent *QWidget) *QDockWidget {
	title_ms := miqt_strdupg(title)
	defer C.free(title_ms)
	ret := C.QDockWidget_new3((*C.struct_miqt_string)(title_ms), parent.cPointer())
	return newQDockWidget(ret)
}

// NewQDockWidget4 constructs a new QDockWidget object.
func NewQDockWidget4(title string, parent *QWidget, flags int) *QDockWidget {
	title_ms := miqt_strdupg(title)
	defer C.free(title_ms)
	ret := C.QDockWidget_new4((*C.struct_miqt_string)(title_ms), parent.cPointer(), (C.int)(flags))
	return newQDockWidget(ret)
}

// NewQDockWidget5 constructs a new QDockWidget object.
func NewQDockWidget5(parent *QWidget) *QDockWidget {
	ret := C.QDockWidget_new5(parent.cPointer())
	return newQDockWidget(ret)
}

// NewQDockWidget6 constructs a new QDockWidget object.
func NewQDockWidget6(parent *QWidget, flags int) *QDockWidget {
	ret := C.QDockWidget_new6(parent.cPointer(), (C.int)(flags))
	return newQDockWidget(ret)
}

func (this *QDockWidget) MetaObject() *QMetaObject {
	_ret := C.QDockWidget_MetaObject(this.h)
	return newQMetaObject_U(unsafe.Pointer(_ret))
}

func QDockWidget_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_Tr(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_TrUtf8(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QDockWidget) Widget() *QWidget {
	_ret := C.QDockWidget_Widget(this.h)
	return newQWidget_U(unsafe.Pointer(_ret))
}

func (this *QDockWidget) SetWidget(widget *QWidget) {
	C.QDockWidget_SetWidget(this.h, widget.cPointer())
}

func (this *QDockWidget) SetFeatures(features int) {
	C.QDockWidget_SetFeatures(this.h, (C.int)(features))
}

func (this *QDockWidget) Features() int {
	_ret := C.QDockWidget_Features(this.h)
	return (int)(_ret)
}

func (this *QDockWidget) SetFloating(floating bool) {
	C.QDockWidget_SetFloating(this.h, (C.bool)(floating))
}

func (this *QDockWidget) IsFloating() bool {
	_ret := C.QDockWidget_IsFloating(this.h)
	return (bool)(_ret)
}

func (this *QDockWidget) SetAllowedAreas(areas int) {
	C.QDockWidget_SetAllowedAreas(this.h, (C.int)(areas))
}

func (this *QDockWidget) AllowedAreas() int {
	_ret := C.QDockWidget_AllowedAreas(this.h)
	return (int)(_ret)
}

func (this *QDockWidget) SetTitleBarWidget(widget *QWidget) {
	C.QDockWidget_SetTitleBarWidget(this.h, widget.cPointer())
}

func (this *QDockWidget) TitleBarWidget() *QWidget {
	_ret := C.QDockWidget_TitleBarWidget(this.h)
	return newQWidget_U(unsafe.Pointer(_ret))
}

func (this *QDockWidget) IsAreaAllowed(area DockWidgetArea) bool {
	_ret := C.QDockWidget_IsAreaAllowed(this.h, (C.uintptr_t)(area))
	return (bool)(_ret)
}

func (this *QDockWidget) ToggleViewAction() *QAction {
	_ret := C.QDockWidget_ToggleViewAction(this.h)
	return newQAction_U(unsafe.Pointer(_ret))
}

func (this *QDockWidget) FeaturesChanged(features int) {
	C.QDockWidget_FeaturesChanged(this.h, (C.int)(features))
}
func (this *QDockWidget) OnFeaturesChanged(slot func(features int)) {
	C.QDockWidget_connect_FeaturesChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_FeaturesChanged
func miqt_exec_callback_QDockWidget_FeaturesChanged(cb *C.void, features C.int) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(features int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	features_ret := features
	slotval1 := (int)(features_ret)

	gofunc(slotval1)
}

func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	C.QDockWidget_TopLevelChanged(this.h, (C.bool)(topLevel))
}
func (this *QDockWidget) OnTopLevelChanged(slot func(topLevel bool)) {
	C.QDockWidget_connect_TopLevelChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_TopLevelChanged
func miqt_exec_callback_QDockWidget_TopLevelChanged(cb *C.void, topLevel C.bool) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(topLevel bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	topLevel_ret := topLevel
	slotval1 := (bool)(topLevel_ret)

	gofunc(slotval1)
}

func (this *QDockWidget) AllowedAreasChanged(allowedAreas int) {
	C.QDockWidget_AllowedAreasChanged(this.h, (C.int)(allowedAreas))
}
func (this *QDockWidget) OnAllowedAreasChanged(slot func(allowedAreas int)) {
	C.QDockWidget_connect_AllowedAreasChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_AllowedAreasChanged
func miqt_exec_callback_QDockWidget_AllowedAreasChanged(cb *C.void, allowedAreas C.int) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(allowedAreas int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	allowedAreas_ret := allowedAreas
	slotval1 := (int)(allowedAreas_ret)

	gofunc(slotval1)
}

func (this *QDockWidget) VisibilityChanged(visible bool) {
	C.QDockWidget_VisibilityChanged(this.h, (C.bool)(visible))
}
func (this *QDockWidget) OnVisibilityChanged(slot func(visible bool)) {
	C.QDockWidget_connect_VisibilityChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_VisibilityChanged
func miqt_exec_callback_QDockWidget_VisibilityChanged(cb *C.void, visible C.bool) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	visible_ret := visible
	slotval1 := (bool)(visible_ret)

	gofunc(slotval1)
}

func (this *QDockWidget) DockLocationChanged(area DockWidgetArea) {
	C.QDockWidget_DockLocationChanged(this.h, (C.uintptr_t)(area))
}
func (this *QDockWidget) OnDockLocationChanged(slot func(area DockWidgetArea)) {
	C.QDockWidget_connect_DockLocationChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_DockLocationChanged
func miqt_exec_callback_QDockWidget_DockLocationChanged(cb *C.void, area C.uintptr_t) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(area DockWidgetArea))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	area_ret := area
	slotval1 := (DockWidgetArea)(area_ret)

	gofunc(slotval1)
}

func QDockWidget_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

// Delete this object from C++ memory.
func (this *QDockWidget) Delete() {
	C.QDockWidget_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QDockWidget) GoGC() {
	runtime.SetFinalizer(this, func(this *QDockWidget) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
