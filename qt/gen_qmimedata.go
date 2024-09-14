package qt

/*

#include "gen_qmimedata.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QMimeData struct {
	h *C.QMimeData
	*QObject
}

func (this *QMimeData) cPointer() *C.QMimeData {
	if this == nil {
		return nil
	}
	return this.h
}

func newQMimeData(h *C.QMimeData) *QMimeData {
	if h == nil {
		return nil
	}
	return &QMimeData{h: h, QObject: newQObject_U(unsafe.Pointer(h))}
}

func newQMimeData_U(h unsafe.Pointer) *QMimeData {
	return newQMimeData((*C.QMimeData)(h))
}

// NewQMimeData constructs a new QMimeData object.
func NewQMimeData() *QMimeData {
	ret := C.QMimeData_new()
	return newQMimeData(ret)
}

func (this *QMimeData) MetaObject() *QMetaObject {
	_ret := C.QMimeData_MetaObject(this.h)
	return newQMetaObject_U(unsafe.Pointer(_ret))
}

func QMimeData_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QMimeData_Tr(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QMimeData_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QMimeData_TrUtf8(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QMimeData) Urls() []QUrl {
	var _ma *C.struct_miqt_array = C.QMimeData_Urls(this.h)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*C.QUrl)(unsafe.Pointer(_ma.data)) // mrs jackson
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = *newQUrl(_outCast[i])
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QMimeData) SetUrls(urls []QUrl) {
	// For the C ABI, malloc a C array of raw pointers
	urls_CArray := (*[0xffff]*C.QUrl)(C.malloc(C.size_t(8 * len(urls))))
	defer C.free(unsafe.Pointer(urls_CArray))
	for i := range urls {
		urls_CArray[i] = urls[i].cPointer()
	}
	urls_ma := &C.struct_miqt_array{len: C.size_t(len(urls)), data: unsafe.Pointer(urls_CArray)}
	defer runtime.KeepAlive(unsafe.Pointer(urls_ma))
	C.QMimeData_SetUrls(this.h, urls_ma)
}

func (this *QMimeData) HasUrls() bool {
	_ret := C.QMimeData_HasUrls(this.h)
	return (bool)(_ret)
}

func (this *QMimeData) Text() string {
	var _ms *C.struct_miqt_string = C.QMimeData_Text(this.h)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QMimeData) SetText(text string) {
	text_ms := miqt_strdupg(text)
	defer C.free(text_ms)
	C.QMimeData_SetText(this.h, (*C.struct_miqt_string)(text_ms))
}

func (this *QMimeData) HasText() bool {
	_ret := C.QMimeData_HasText(this.h)
	return (bool)(_ret)
}

func (this *QMimeData) Html() string {
	var _ms *C.struct_miqt_string = C.QMimeData_Html(this.h)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QMimeData) SetHtml(html string) {
	html_ms := miqt_strdupg(html)
	defer C.free(html_ms)
	C.QMimeData_SetHtml(this.h, (*C.struct_miqt_string)(html_ms))
}

func (this *QMimeData) HasHtml() bool {
	_ret := C.QMimeData_HasHtml(this.h)
	return (bool)(_ret)
}

func (this *QMimeData) ImageData() *QVariant {
	_ret := C.QMimeData_ImageData(this.h)
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeData) SetImageData(image *QVariant) {
	C.QMimeData_SetImageData(this.h, image.cPointer())
}

func (this *QMimeData) HasImage() bool {
	_ret := C.QMimeData_HasImage(this.h)
	return (bool)(_ret)
}

func (this *QMimeData) ColorData() *QVariant {
	_ret := C.QMimeData_ColorData(this.h)
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeData) SetColorData(color *QVariant) {
	C.QMimeData_SetColorData(this.h, color.cPointer())
}

func (this *QMimeData) HasColor() bool {
	_ret := C.QMimeData_HasColor(this.h)
	return (bool)(_ret)
}

func (this *QMimeData) Data(mimetype string) *QByteArray {
	mimetype_ms := miqt_strdupg(mimetype)
	defer C.free(mimetype_ms)
	_ret := C.QMimeData_Data(this.h, (*C.struct_miqt_string)(mimetype_ms))
	_goptr := newQByteArray(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeData) SetData(mimetype string, data *QByteArray) {
	mimetype_ms := miqt_strdupg(mimetype)
	defer C.free(mimetype_ms)
	C.QMimeData_SetData(this.h, (*C.struct_miqt_string)(mimetype_ms), data.cPointer())
}

func (this *QMimeData) RemoveFormat(mimetype string) {
	mimetype_ms := miqt_strdupg(mimetype)
	defer C.free(mimetype_ms)
	C.QMimeData_RemoveFormat(this.h, (*C.struct_miqt_string)(mimetype_ms))
}

func (this *QMimeData) HasFormat(mimetype string) bool {
	mimetype_ms := miqt_strdupg(mimetype)
	defer C.free(mimetype_ms)
	_ret := C.QMimeData_HasFormat(this.h, (*C.struct_miqt_string)(mimetype_ms))
	return (bool)(_ret)
}

func (this *QMimeData) Formats() []string {
	var _ma *C.struct_miqt_array = C.QMimeData_Formats(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]*C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = C.GoStringN(&_outCast[i].data, C.int(int64(_outCast[i].len)))
		C.free(unsafe.Pointer(_outCast[i])) // free the inner miqt_string*
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QMimeData) Clear() {
	C.QMimeData_Clear(this.h)
}

func QMimeData_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QMimeData_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QMimeData_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QMimeData_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QMimeData_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QMimeData_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QMimeData_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QMimeData_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

// Delete this object from C++ memory.
func (this *QMimeData) Delete() {
	C.QMimeData_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMimeData) GoGC() {
	runtime.SetFinalizer(this, func(this *QMimeData) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
