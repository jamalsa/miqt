package qt

/*

#include "gen_qurlquery.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QUrlQuery struct {
	h *C.QUrlQuery
}

func (this *QUrlQuery) cPointer() *C.QUrlQuery {
	if this == nil {
		return nil
	}
	return this.h
}

func newQUrlQuery(h *C.QUrlQuery) *QUrlQuery {
	if h == nil {
		return nil
	}
	return &QUrlQuery{h: h}
}

func newQUrlQuery_U(h unsafe.Pointer) *QUrlQuery {
	return newQUrlQuery((*C.QUrlQuery)(h))
}

// NewQUrlQuery constructs a new QUrlQuery object.
func NewQUrlQuery() *QUrlQuery {
	ret := C.QUrlQuery_new()
	return newQUrlQuery(ret)
}

// NewQUrlQuery2 constructs a new QUrlQuery object.
func NewQUrlQuery2(url *QUrl) *QUrlQuery {
	ret := C.QUrlQuery_new2(url.cPointer())
	return newQUrlQuery(ret)
}

// NewQUrlQuery3 constructs a new QUrlQuery object.
func NewQUrlQuery3(queryString string) *QUrlQuery {
	queryString_ms := miqt_strdupg(queryString)
	defer C.free(queryString_ms)
	ret := C.QUrlQuery_new3((*C.struct_miqt_string)(queryString_ms))
	return newQUrlQuery(ret)
}

// NewQUrlQuery4 constructs a new QUrlQuery object.
func NewQUrlQuery4(other *QUrlQuery) *QUrlQuery {
	ret := C.QUrlQuery_new4(other.cPointer())
	return newQUrlQuery(ret)
}

func (this *QUrlQuery) OperatorAssign(other *QUrlQuery) {
	C.QUrlQuery_OperatorAssign(this.h, other.cPointer())
}

func (this *QUrlQuery) OperatorEqual(other *QUrlQuery) bool {
	return (bool)(C.QUrlQuery_OperatorEqual(this.h, other.cPointer()))
}

func (this *QUrlQuery) OperatorNotEqual(other *QUrlQuery) bool {
	return (bool)(C.QUrlQuery_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QUrlQuery) Swap(other *QUrlQuery) {
	C.QUrlQuery_Swap(this.h, other.cPointer())
}

func (this *QUrlQuery) IsEmpty() bool {
	return (bool)(C.QUrlQuery_IsEmpty(this.h))
}

func (this *QUrlQuery) IsDetached() bool {
	return (bool)(C.QUrlQuery_IsDetached(this.h))
}

func (this *QUrlQuery) Clear() {
	C.QUrlQuery_Clear(this.h)
}

func (this *QUrlQuery) Query() string {
	var _ms *C.struct_miqt_string = C.QUrlQuery_Query(this.h)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QUrlQuery) SetQuery(queryString string) {
	queryString_ms := miqt_strdupg(queryString)
	defer C.free(queryString_ms)
	C.QUrlQuery_SetQuery(this.h, (*C.struct_miqt_string)(queryString_ms))
}

func (this *QUrlQuery) ToString() string {
	var _ms *C.struct_miqt_string = C.QUrlQuery_ToString(this.h)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QUrlQuery) SetQueryDelimiters(valueDelimiter QChar, pairDelimiter QChar) {
	C.QUrlQuery_SetQueryDelimiters(this.h, valueDelimiter.cPointer(), pairDelimiter.cPointer())
}

func (this *QUrlQuery) QueryValueDelimiter() *QChar {
	_ret := C.QUrlQuery_QueryValueDelimiter(this.h)
	_goptr := newQChar(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrlQuery) QueryPairDelimiter() *QChar {
	_ret := C.QUrlQuery_QueryPairDelimiter(this.h)
	_goptr := newQChar(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrlQuery) HasQueryItem(key string) bool {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	return (bool)(C.QUrlQuery_HasQueryItem(this.h, (*C.struct_miqt_string)(key_ms)))
}

func (this *QUrlQuery) AddQueryItem(key string, value string) {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	value_ms := miqt_strdupg(value)
	defer C.free(value_ms)
	C.QUrlQuery_AddQueryItem(this.h, (*C.struct_miqt_string)(key_ms), (*C.struct_miqt_string)(value_ms))
}

func (this *QUrlQuery) RemoveQueryItem(key string) {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	C.QUrlQuery_RemoveQueryItem(this.h, (*C.struct_miqt_string)(key_ms))
}

func (this *QUrlQuery) QueryItemValue(key string) string {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	var _ms *C.struct_miqt_string = C.QUrlQuery_QueryItemValue(this.h, (*C.struct_miqt_string)(key_ms))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QUrlQuery) AllQueryItemValues(key string) []string {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	var _ma *C.struct_miqt_array = C.QUrlQuery_AllQueryItemValues(this.h, (*C.struct_miqt_string)(key_ms))
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]*C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms *C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(&_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms))
		_ret[i] = _lv_ret
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QUrlQuery) RemoveAllQueryItems(key string) {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	C.QUrlQuery_RemoveAllQueryItems(this.h, (*C.struct_miqt_string)(key_ms))
}

func QUrlQuery_DefaultQueryValueDelimiter() *QChar {
	_ret := C.QUrlQuery_DefaultQueryValueDelimiter()
	_goptr := newQChar(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QUrlQuery_DefaultQueryPairDelimiter() *QChar {
	_ret := C.QUrlQuery_DefaultQueryPairDelimiter()
	_goptr := newQChar(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrlQuery) Query1(encoding int) string {
	var _ms *C.struct_miqt_string = C.QUrlQuery_Query1(this.h, encoding)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QUrlQuery) ToString1(encoding int) string {
	var _ms *C.struct_miqt_string = C.QUrlQuery_ToString1(this.h, encoding)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QUrlQuery) QueryItemValue2(key string, encoding int) string {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	var _ms *C.struct_miqt_string = C.QUrlQuery_QueryItemValue2(this.h, (*C.struct_miqt_string)(key_ms), encoding)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QUrlQuery) AllQueryItemValues2(key string, encoding int) []string {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	var _ma *C.struct_miqt_array = C.QUrlQuery_AllQueryItemValues2(this.h, (*C.struct_miqt_string)(key_ms), encoding)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]*C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms *C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(&_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms))
		_ret[i] = _lv_ret
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

// Delete this object from C++ memory.
func (this *QUrlQuery) Delete() {
	C.QUrlQuery_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QUrlQuery) GoGC() {
	runtime.SetFinalizer(this, func(this *QUrlQuery) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
