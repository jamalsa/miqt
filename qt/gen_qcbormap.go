package qt

/*

#include "gen_qcbormap.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QCborMap struct {
	h *C.QCborMap
}

func (this *QCborMap) cPointer() *C.QCborMap {
	if this == nil {
		return nil
	}
	return this.h
}

func newQCborMap(h *C.QCborMap) *QCborMap {
	if h == nil {
		return nil
	}
	return &QCborMap{h: h}
}

func newQCborMap_U(h unsafe.Pointer) *QCborMap {
	return newQCborMap((*C.QCborMap)(h))
}

// NewQCborMap constructs a new QCborMap object.
func NewQCborMap() *QCborMap {
	ret := C.QCborMap_new()
	return newQCborMap(ret)
}

// NewQCborMap2 constructs a new QCborMap object.
func NewQCborMap2(other *QCborMap) *QCborMap {
	ret := C.QCborMap_new2(other.cPointer())
	return newQCborMap(ret)
}

func (this *QCborMap) OperatorAssign(other *QCborMap) {
	C.QCborMap_OperatorAssign(this.h, other.cPointer())
}

func (this *QCborMap) Swap(other *QCborMap) {
	C.QCborMap_Swap(this.h, other.cPointer())
}

func (this *QCborMap) ToCborValue() *QCborValue {
	_ret := C.QCborMap_ToCborValue(this.h)
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Size() int64 {
	return (int64)(C.QCborMap_Size(this.h))
}

func (this *QCborMap) IsEmpty() bool {
	return (bool)(C.QCborMap_IsEmpty(this.h))
}

func (this *QCborMap) Clear() {
	C.QCborMap_Clear(this.h)
}

func (this *QCborMap) Keys() []QCborValue {
	var _ma *C.struct_miqt_array = C.QCborMap_Keys(this.h)
	_ret := make([]QCborValue, int(_ma.len))
	_outCast := (*[0xffff]*C.QCborValue)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_vv_ret := _outCast[i]
		_vv_goptr := newQCborValue(_vv_ret)
		_vv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_vv_goptr
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QCborMap) Value(key int64) *QCborValue {
	_ret := C.QCborMap_Value(this.h, (C.longlong)(key))
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Value2(key string) *QCborValue {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_Value2(this.h, (*C.struct_miqt_string)(key_ms))
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Value3(key *QCborValue) *QCborValue {
	_ret := C.QCborMap_Value3(this.h, key.cPointer())
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript(key int64) *QCborValue {
	_ret := C.QCborMap_OperatorSubscript(this.h, (C.longlong)(key))
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript2(key string) *QCborValue {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_OperatorSubscript2(this.h, (*C.struct_miqt_string)(key_ms))
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript3(key *QCborValue) *QCborValue {
	_ret := C.QCborMap_OperatorSubscript3(this.h, key.cPointer())
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript4(key int64) *QCborValueRef {
	_ret := C.QCborMap_OperatorSubscript4(this.h, (C.longlong)(key))
	_goptr := newQCborValueRef(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript6(key string) *QCborValueRef {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_OperatorSubscript6(this.h, (*C.struct_miqt_string)(key_ms))
	_goptr := newQCborValueRef(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript7(key *QCborValue) *QCborValueRef {
	_ret := C.QCborMap_OperatorSubscript7(this.h, key.cPointer())
	_goptr := newQCborValueRef(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Take(key int64) *QCborValue {
	_ret := C.QCborMap_Take(this.h, (C.longlong)(key))
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Take2(key string) *QCborValue {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_Take2(this.h, (*C.struct_miqt_string)(key_ms))
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Take3(key *QCborValue) *QCborValue {
	_ret := C.QCborMap_Take3(this.h, key.cPointer())
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Remove(key int64) {
	C.QCborMap_Remove(this.h, (C.longlong)(key))
}

func (this *QCborMap) Remove2(key string) {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	C.QCborMap_Remove2(this.h, (*C.struct_miqt_string)(key_ms))
}

func (this *QCborMap) Remove3(key *QCborValue) {
	C.QCborMap_Remove3(this.h, key.cPointer())
}

func (this *QCborMap) Contains(key int64) bool {
	return (bool)(C.QCborMap_Contains(this.h, (C.longlong)(key)))
}

func (this *QCborMap) Contains2(key string) bool {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	return (bool)(C.QCborMap_Contains2(this.h, (*C.struct_miqt_string)(key_ms)))
}

func (this *QCborMap) Contains3(key *QCborValue) bool {
	return (bool)(C.QCborMap_Contains3(this.h, key.cPointer()))
}

func (this *QCborMap) Compare(other *QCborMap) int {
	return (int)(C.QCborMap_Compare(this.h, other.cPointer()))
}

func (this *QCborMap) OperatorEqual(other *QCborMap) bool {
	return (bool)(C.QCborMap_OperatorEqual(this.h, other.cPointer()))
}

func (this *QCborMap) OperatorNotEqual(other *QCborMap) bool {
	return (bool)(C.QCborMap_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QCborMap) OperatorLesser(other *QCborMap) bool {
	return (bool)(C.QCborMap_OperatorLesser(this.h, other.cPointer()))
}

func (this *QCborMap) Begin() *QCborMap__Iterator {
	_ret := C.QCborMap_Begin(this.h)
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ConstBegin() *QCborMap__ConstIterator {
	_ret := C.QCborMap_ConstBegin(this.h)
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Begin2() *QCborMap__ConstIterator {
	_ret := C.QCborMap_Begin2(this.h)
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Cbegin() *QCborMap__ConstIterator {
	_ret := C.QCborMap_Cbegin(this.h)
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) End() *QCborMap__Iterator {
	_ret := C.QCborMap_End(this.h)
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ConstEnd() *QCborMap__ConstIterator {
	_ret := C.QCborMap_ConstEnd(this.h)
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) End2() *QCborMap__ConstIterator {
	_ret := C.QCborMap_End2(this.h)
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Cend() *QCborMap__ConstIterator {
	_ret := C.QCborMap_Cend(this.h)
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Erase(it QCborMap__Iterator) *QCborMap__Iterator {
	_ret := C.QCborMap_Erase(this.h, it.cPointer())
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) EraseWithIt(it QCborMap__ConstIterator) *QCborMap__Iterator {
	_ret := C.QCborMap_EraseWithIt(this.h, it.cPointer())
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Extract(it QCborMap__Iterator) *QCborValue {
	_ret := C.QCborMap_Extract(this.h, it.cPointer())
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ExtractWithIt(it QCborMap__ConstIterator) *QCborValue {
	_ret := C.QCborMap_ExtractWithIt(this.h, it.cPointer())
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Empty() bool {
	return (bool)(C.QCborMap_Empty(this.h))
}

func (this *QCborMap) Find(key int64) *QCborMap__Iterator {
	_ret := C.QCborMap_Find(this.h, (C.longlong)(key))
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Find2(key string) *QCborMap__Iterator {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_Find2(this.h, (*C.struct_miqt_string)(key_ms))
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Find3(key *QCborValue) *QCborMap__Iterator {
	_ret := C.QCborMap_Find3(this.h, key.cPointer())
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ConstFind(key int64) *QCborMap__ConstIterator {
	_ret := C.QCborMap_ConstFind(this.h, (C.longlong)(key))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ConstFind2(key string) *QCborMap__ConstIterator {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_ConstFind2(this.h, (*C.struct_miqt_string)(key_ms))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ConstFind3(key *QCborValue) *QCborMap__ConstIterator {
	_ret := C.QCborMap_ConstFind3(this.h, key.cPointer())
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Find4(key int64) *QCborMap__ConstIterator {
	_ret := C.QCborMap_Find4(this.h, (C.longlong)(key))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Find6(key string) *QCborMap__ConstIterator {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_Find6(this.h, (*C.struct_miqt_string)(key_ms))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Find7(key *QCborValue) *QCborMap__ConstIterator {
	_ret := C.QCborMap_Find7(this.h, key.cPointer())
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Insert(key int64, value_ *QCborValue) *QCborMap__Iterator {
	_ret := C.QCborMap_Insert(this.h, (C.longlong)(key), value_.cPointer())
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Insert3(key string, value_ *QCborValue) *QCborMap__Iterator {
	key_ms := miqt_strdupg(key)
	defer C.free(key_ms)
	_ret := C.QCborMap_Insert3(this.h, (*C.struct_miqt_string)(key_ms), value_.cPointer())
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Insert4(key *QCborValue, value_ *QCborValue) *QCborMap__Iterator {
	_ret := C.QCborMap_Insert4(this.h, key.cPointer(), value_.cPointer())
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborMap_FromJsonObject(o *QJsonObject) *QCborMap {
	_ret := C.QCborMap_FromJsonObject(o.cPointer())
	_goptr := newQCborMap(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ToJsonObject() *QJsonObject {
	_ret := C.QCborMap_ToJsonObject(this.h)
	_goptr := newQJsonObject(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QCborMap) Delete() {
	C.QCborMap_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCborMap) GoGC() {
	runtime.SetFinalizer(this, func(this *QCborMap) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QCborMap__Iterator struct {
	h *C.QCborMap__Iterator
}

func (this *QCborMap__Iterator) cPointer() *C.QCborMap__Iterator {
	if this == nil {
		return nil
	}
	return this.h
}

func newQCborMap__Iterator(h *C.QCborMap__Iterator) *QCborMap__Iterator {
	if h == nil {
		return nil
	}
	return &QCborMap__Iterator{h: h}
}

func newQCborMap__Iterator_U(h unsafe.Pointer) *QCborMap__Iterator {
	return newQCborMap__Iterator((*C.QCborMap__Iterator)(h))
}

// NewQCborMap__Iterator constructs a new QCborMap::Iterator object.
func NewQCborMap__Iterator() *QCborMap__Iterator {
	ret := C.QCborMap__Iterator_new()
	return newQCborMap__Iterator(ret)
}

// NewQCborMap__Iterator2 constructs a new QCborMap::Iterator object.
func NewQCborMap__Iterator2(param1 *QCborMap__Iterator) *QCborMap__Iterator {
	ret := C.QCborMap__Iterator_new2(param1.cPointer())
	return newQCborMap__Iterator(ret)
}

func (this *QCborMap__Iterator) OperatorAssign(other *QCborMap__Iterator) {
	C.QCborMap__Iterator_OperatorAssign(this.h, other.cPointer())
}

func (this *QCborMap__Iterator) OperatorMinusGreater() *QCborValueRef {
	return newQCborValueRef_U(unsafe.Pointer(C.QCborMap__Iterator_OperatorMinusGreater(this.h)))
}

func (this *QCborMap__Iterator) Key() *QCborValue {
	_ret := C.QCborMap__Iterator_Key(this.h)
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) Value() *QCborValueRef {
	_ret := C.QCborMap__Iterator_Value(this.h)
	_goptr := newQCborValueRef(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) OperatorEqual(o *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorEqual(this.h, o.cPointer()))
}

func (this *QCborMap__Iterator) OperatorNotEqual(o *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorNotEqual(this.h, o.cPointer()))
}

func (this *QCborMap__Iterator) OperatorLesser(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorLesser(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorLesserOrEqual(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorLesserOrEqual(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorGreater(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorGreater(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorGreaterOrEqual(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorGreaterOrEqual(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorEqualWithQCborMapConstIterator(o *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorEqualWithQCborMapConstIterator(this.h, o.cPointer()))
}

func (this *QCborMap__Iterator) OperatorNotEqualWithQCborMapConstIterator(o *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorNotEqualWithQCborMapConstIterator(this.h, o.cPointer()))
}

func (this *QCborMap__Iterator) OperatorLesserWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorLesserWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorLesserOrEqualWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorLesserOrEqualWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorGreaterWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorGreaterWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorGreaterOrEqualWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__Iterator_OperatorGreaterOrEqualWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__Iterator) OperatorPlusPlus() *QCborMap__Iterator {
	return newQCborMap__Iterator_U(unsafe.Pointer(C.QCborMap__Iterator_OperatorPlusPlus(this.h)))
}

func (this *QCborMap__Iterator) OperatorPlusPlusWithInt(param1 int) *QCborMap__Iterator {
	_ret := C.QCborMap__Iterator_OperatorPlusPlusWithInt(this.h, (C.int)(param1))
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) OperatorMinusMinus() *QCborMap__Iterator {
	return newQCborMap__Iterator_U(unsafe.Pointer(C.QCborMap__Iterator_OperatorMinusMinus(this.h)))
}

func (this *QCborMap__Iterator) OperatorMinusMinusWithInt(param1 int) *QCborMap__Iterator {
	_ret := C.QCborMap__Iterator_OperatorMinusMinusWithInt(this.h, (C.int)(param1))
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) OperatorPlusAssign(j int64) *QCborMap__Iterator {
	return newQCborMap__Iterator_U(unsafe.Pointer(C.QCborMap__Iterator_OperatorPlusAssign(this.h, (C.ptrdiff_t)(j))))
}

func (this *QCborMap__Iterator) OperatorMinusAssign(j int64) *QCborMap__Iterator {
	return newQCborMap__Iterator_U(unsafe.Pointer(C.QCborMap__Iterator_OperatorMinusAssign(this.h, (C.ptrdiff_t)(j))))
}

func (this *QCborMap__Iterator) OperatorPlus(j int64) *QCborMap__Iterator {
	_ret := C.QCborMap__Iterator_OperatorPlus(this.h, (C.ptrdiff_t)(j))
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) OperatorMinus(j int64) *QCborMap__Iterator {
	_ret := C.QCborMap__Iterator_OperatorMinus(this.h, (C.ptrdiff_t)(j))
	_goptr := newQCborMap__Iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) OperatorMinusWithQCborMapIterator(j QCborMap__Iterator) int64 {
	return (int64)(C.QCborMap__Iterator_OperatorMinusWithQCborMapIterator(this.h, j.cPointer()))
}

// Delete this object from C++ memory.
func (this *QCborMap__Iterator) Delete() {
	C.QCborMap__Iterator_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCborMap__Iterator) GoGC() {
	runtime.SetFinalizer(this, func(this *QCborMap__Iterator) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QCborMap__ConstIterator struct {
	h *C.QCborMap__ConstIterator
}

func (this *QCborMap__ConstIterator) cPointer() *C.QCborMap__ConstIterator {
	if this == nil {
		return nil
	}
	return this.h
}

func newQCborMap__ConstIterator(h *C.QCborMap__ConstIterator) *QCborMap__ConstIterator {
	if h == nil {
		return nil
	}
	return &QCborMap__ConstIterator{h: h}
}

func newQCborMap__ConstIterator_U(h unsafe.Pointer) *QCborMap__ConstIterator {
	return newQCborMap__ConstIterator((*C.QCborMap__ConstIterator)(h))
}

// NewQCborMap__ConstIterator constructs a new QCborMap::ConstIterator object.
func NewQCborMap__ConstIterator() *QCborMap__ConstIterator {
	ret := C.QCborMap__ConstIterator_new()
	return newQCborMap__ConstIterator(ret)
}

// NewQCborMap__ConstIterator2 constructs a new QCborMap::ConstIterator object.
func NewQCborMap__ConstIterator2(param1 *QCborMap__ConstIterator) *QCborMap__ConstIterator {
	ret := C.QCborMap__ConstIterator_new2(param1.cPointer())
	return newQCborMap__ConstIterator(ret)
}

func (this *QCborMap__ConstIterator) OperatorAssign(other *QCborMap__ConstIterator) {
	C.QCborMap__ConstIterator_OperatorAssign(this.h, other.cPointer())
}

func (this *QCborMap__ConstIterator) OperatorMinusGreater() *QCborValueRef {
	return newQCborValueRef_U(unsafe.Pointer(C.QCborMap__ConstIterator_OperatorMinusGreater(this.h)))
}

func (this *QCborMap__ConstIterator) Key() *QCborValue {
	_ret := C.QCborMap__ConstIterator_Key(this.h)
	_goptr := newQCborValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) Value() *QCborValueRef {
	_ret := C.QCborMap__ConstIterator_Value(this.h)
	_goptr := newQCborValueRef(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) OperatorEqual(o *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorEqual(this.h, o.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorNotEqual(o *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorNotEqual(this.h, o.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorLesser(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorLesser(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorLesserOrEqual(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorLesserOrEqual(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorGreater(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorGreater(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorGreaterOrEqual(other *QCborMap__Iterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorGreaterOrEqual(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorEqualWithQCborMapConstIterator(o *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorEqualWithQCborMapConstIterator(this.h, o.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorNotEqualWithQCborMapConstIterator(o *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorNotEqualWithQCborMapConstIterator(this.h, o.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorLesserWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorLesserWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorLesserOrEqualWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorLesserOrEqualWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorGreaterWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorGreaterWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorGreaterOrEqualWithOther(other *QCborMap__ConstIterator) bool {
	return (bool)(C.QCborMap__ConstIterator_OperatorGreaterOrEqualWithOther(this.h, other.cPointer()))
}

func (this *QCborMap__ConstIterator) OperatorPlusPlus() *QCborMap__ConstIterator {
	return newQCborMap__ConstIterator_U(unsafe.Pointer(C.QCborMap__ConstIterator_OperatorPlusPlus(this.h)))
}

func (this *QCborMap__ConstIterator) OperatorPlusPlusWithInt(param1 int) *QCborMap__ConstIterator {
	_ret := C.QCborMap__ConstIterator_OperatorPlusPlusWithInt(this.h, (C.int)(param1))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) OperatorMinusMinus() *QCborMap__ConstIterator {
	return newQCborMap__ConstIterator_U(unsafe.Pointer(C.QCborMap__ConstIterator_OperatorMinusMinus(this.h)))
}

func (this *QCborMap__ConstIterator) OperatorMinusMinusWithInt(param1 int) *QCborMap__ConstIterator {
	_ret := C.QCborMap__ConstIterator_OperatorMinusMinusWithInt(this.h, (C.int)(param1))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) OperatorPlusAssign(j int64) *QCborMap__ConstIterator {
	return newQCborMap__ConstIterator_U(unsafe.Pointer(C.QCborMap__ConstIterator_OperatorPlusAssign(this.h, (C.ptrdiff_t)(j))))
}

func (this *QCborMap__ConstIterator) OperatorMinusAssign(j int64) *QCborMap__ConstIterator {
	return newQCborMap__ConstIterator_U(unsafe.Pointer(C.QCborMap__ConstIterator_OperatorMinusAssign(this.h, (C.ptrdiff_t)(j))))
}

func (this *QCborMap__ConstIterator) OperatorPlus(j int64) *QCborMap__ConstIterator {
	_ret := C.QCborMap__ConstIterator_OperatorPlus(this.h, (C.ptrdiff_t)(j))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) OperatorMinus(j int64) *QCborMap__ConstIterator {
	_ret := C.QCborMap__ConstIterator_OperatorMinus(this.h, (C.ptrdiff_t)(j))
	_goptr := newQCborMap__ConstIterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) OperatorMinusWithQCborMapConstIterator(j QCborMap__ConstIterator) int64 {
	return (int64)(C.QCborMap__ConstIterator_OperatorMinusWithQCborMapConstIterator(this.h, j.cPointer()))
}

// Delete this object from C++ memory.
func (this *QCborMap__ConstIterator) Delete() {
	C.QCborMap__ConstIterator_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCborMap__ConstIterator) GoGC() {
	runtime.SetFinalizer(this, func(this *QCborMap__ConstIterator) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
