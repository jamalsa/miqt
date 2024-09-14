package qt

/*

#include "gen_qtemporaryfile.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QTemporaryFile struct {
	h *C.QTemporaryFile
	*QFile
}

func (this *QTemporaryFile) cPointer() *C.QTemporaryFile {
	if this == nil {
		return nil
	}
	return this.h
}

func newQTemporaryFile(h *C.QTemporaryFile) *QTemporaryFile {
	if h == nil {
		return nil
	}
	return &QTemporaryFile{h: h, QFile: newQFile_U(unsafe.Pointer(h))}
}

func newQTemporaryFile_U(h unsafe.Pointer) *QTemporaryFile {
	return newQTemporaryFile((*C.QTemporaryFile)(h))
}

// NewQTemporaryFile constructs a new QTemporaryFile object.
func NewQTemporaryFile() *QTemporaryFile {
	ret := C.QTemporaryFile_new()
	return newQTemporaryFile(ret)
}

// NewQTemporaryFile2 constructs a new QTemporaryFile object.
func NewQTemporaryFile2(templateName string) *QTemporaryFile {
	templateName_ms := miqt_strdupg(templateName)
	defer C.free(templateName_ms)
	ret := C.QTemporaryFile_new2((*C.struct_miqt_string)(templateName_ms))
	return newQTemporaryFile(ret)
}

// NewQTemporaryFile3 constructs a new QTemporaryFile object.
func NewQTemporaryFile3(parent *QObject) *QTemporaryFile {
	ret := C.QTemporaryFile_new3(parent.cPointer())
	return newQTemporaryFile(ret)
}

// NewQTemporaryFile4 constructs a new QTemporaryFile object.
func NewQTemporaryFile4(templateName string, parent *QObject) *QTemporaryFile {
	templateName_ms := miqt_strdupg(templateName)
	defer C.free(templateName_ms)
	ret := C.QTemporaryFile_new4((*C.struct_miqt_string)(templateName_ms), parent.cPointer())
	return newQTemporaryFile(ret)
}

func (this *QTemporaryFile) MetaObject() *QMetaObject {
	_ret := C.QTemporaryFile_MetaObject(this.h)
	return newQMetaObject_U(unsafe.Pointer(_ret))
}

func QTemporaryFile_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QTemporaryFile_Tr(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QTemporaryFile_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QTemporaryFile_TrUtf8(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QTemporaryFile) AutoRemove() bool {
	_ret := C.QTemporaryFile_AutoRemove(this.h)
	return (bool)(_ret)
}

func (this *QTemporaryFile) SetAutoRemove(b bool) {
	C.QTemporaryFile_SetAutoRemove(this.h, (C.bool)(b))
}

func (this *QTemporaryFile) Open() bool {
	_ret := C.QTemporaryFile_Open(this.h)
	return (bool)(_ret)
}

func (this *QTemporaryFile) FileName() string {
	var _ms *C.struct_miqt_string = C.QTemporaryFile_FileName(this.h)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QTemporaryFile) FileTemplate() string {
	var _ms *C.struct_miqt_string = C.QTemporaryFile_FileTemplate(this.h)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QTemporaryFile) SetFileTemplate(name string) {
	name_ms := miqt_strdupg(name)
	defer C.free(name_ms)
	C.QTemporaryFile_SetFileTemplate(this.h, (*C.struct_miqt_string)(name_ms))
}

func (this *QTemporaryFile) Rename(newName string) bool {
	newName_ms := miqt_strdupg(newName)
	defer C.free(newName_ms)
	_ret := C.QTemporaryFile_Rename(this.h, (*C.struct_miqt_string)(newName_ms))
	return (bool)(_ret)
}

func QTemporaryFile_CreateLocalFile(fileName string) *QTemporaryFile {
	fileName_ms := miqt_strdupg(fileName)
	defer C.free(fileName_ms)
	_ret := C.QTemporaryFile_CreateLocalFile((*C.struct_miqt_string)(fileName_ms))
	return newQTemporaryFile_U(unsafe.Pointer(_ret))
}

func QTemporaryFile_CreateLocalFileWithFile(file *QFile) *QTemporaryFile {
	_ret := C.QTemporaryFile_CreateLocalFileWithFile(file.cPointer())
	return newQTemporaryFile_U(unsafe.Pointer(_ret))
}

func QTemporaryFile_CreateNativeFile(fileName string) *QTemporaryFile {
	fileName_ms := miqt_strdupg(fileName)
	defer C.free(fileName_ms)
	_ret := C.QTemporaryFile_CreateNativeFile((*C.struct_miqt_string)(fileName_ms))
	return newQTemporaryFile_U(unsafe.Pointer(_ret))
}

func QTemporaryFile_CreateNativeFileWithFile(file *QFile) *QTemporaryFile {
	_ret := C.QTemporaryFile_CreateNativeFileWithFile(file.cPointer())
	return newQTemporaryFile_U(unsafe.Pointer(_ret))
}

func QTemporaryFile_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QTemporaryFile_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QTemporaryFile_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QTemporaryFile_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QTemporaryFile_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QTemporaryFile_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QTemporaryFile_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QTemporaryFile_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

// Delete this object from C++ memory.
func (this *QTemporaryFile) Delete() {
	C.QTemporaryFile_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QTemporaryFile) GoGC() {
	runtime.SetFinalizer(this, func(this *QTemporaryFile) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
