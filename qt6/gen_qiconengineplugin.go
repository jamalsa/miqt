package qt6

/*

#include "gen_qiconengineplugin.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QIconEnginePlugin struct {
	h          *C.QIconEnginePlugin
	isSubclass bool
	*QObject
}

func (this *QIconEnginePlugin) cPointer() *C.QIconEnginePlugin {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QIconEnginePlugin) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQIconEnginePlugin constructs the type using only CGO pointers.
func newQIconEnginePlugin(h *C.QIconEnginePlugin, h_QObject *C.QObject) *QIconEnginePlugin {
	if h == nil {
		return nil
	}
	return &QIconEnginePlugin{h: h,
		QObject: newQObject(h_QObject)}
}

// UnsafeNewQIconEnginePlugin constructs the type using only unsafe pointers.
func UnsafeNewQIconEnginePlugin(h unsafe.Pointer, h_QObject unsafe.Pointer) *QIconEnginePlugin {
	if h == nil {
		return nil
	}

	return &QIconEnginePlugin{h: (*C.QIconEnginePlugin)(h),
		QObject: UnsafeNewQObject(h_QObject)}
}

func (this *QIconEnginePlugin) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QIconEnginePlugin_MetaObject(this.h)))
}

func (this *QIconEnginePlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QIconEnginePlugin_Metacast(this.h, param1_Cstring))
}

func QIconEnginePlugin_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QIconEnginePlugin_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEnginePlugin) Create(filename string) *QIconEngine {
	filename_ms := C.struct_miqt_string{}
	filename_ms.data = C.CString(filename)
	filename_ms.len = C.size_t(len(filename))
	defer C.free(unsafe.Pointer(filename_ms.data))
	return UnsafeNewQIconEngine(unsafe.Pointer(C.QIconEnginePlugin_Create(this.h, filename_ms)))
}

func QIconEnginePlugin_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QIconEnginePlugin_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIconEnginePlugin_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QIconEnginePlugin_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QIconEnginePlugin) Delete() {
	C.QIconEnginePlugin_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QIconEnginePlugin) GoGC() {
	runtime.SetFinalizer(this, func(this *QIconEnginePlugin) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
