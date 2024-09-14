package qt

/*

#include "gen_qfontdatabase.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QFontDatabase__WritingSystem int

const (
	QFontDatabase__WritingSystem__Any                 QFontDatabase__WritingSystem = 0
	QFontDatabase__WritingSystem__Latin               QFontDatabase__WritingSystem = 1
	QFontDatabase__WritingSystem__Greek               QFontDatabase__WritingSystem = 2
	QFontDatabase__WritingSystem__Cyrillic            QFontDatabase__WritingSystem = 3
	QFontDatabase__WritingSystem__Armenian            QFontDatabase__WritingSystem = 4
	QFontDatabase__WritingSystem__Hebrew              QFontDatabase__WritingSystem = 5
	QFontDatabase__WritingSystem__Arabic              QFontDatabase__WritingSystem = 6
	QFontDatabase__WritingSystem__Syriac              QFontDatabase__WritingSystem = 7
	QFontDatabase__WritingSystem__Thaana              QFontDatabase__WritingSystem = 8
	QFontDatabase__WritingSystem__Devanagari          QFontDatabase__WritingSystem = 9
	QFontDatabase__WritingSystem__Bengali             QFontDatabase__WritingSystem = 10
	QFontDatabase__WritingSystem__Gurmukhi            QFontDatabase__WritingSystem = 11
	QFontDatabase__WritingSystem__Gujarati            QFontDatabase__WritingSystem = 12
	QFontDatabase__WritingSystem__Oriya               QFontDatabase__WritingSystem = 13
	QFontDatabase__WritingSystem__Tamil               QFontDatabase__WritingSystem = 14
	QFontDatabase__WritingSystem__Telugu              QFontDatabase__WritingSystem = 15
	QFontDatabase__WritingSystem__Kannada             QFontDatabase__WritingSystem = 16
	QFontDatabase__WritingSystem__Malayalam           QFontDatabase__WritingSystem = 17
	QFontDatabase__WritingSystem__Sinhala             QFontDatabase__WritingSystem = 18
	QFontDatabase__WritingSystem__Thai                QFontDatabase__WritingSystem = 19
	QFontDatabase__WritingSystem__Lao                 QFontDatabase__WritingSystem = 20
	QFontDatabase__WritingSystem__Tibetan             QFontDatabase__WritingSystem = 21
	QFontDatabase__WritingSystem__Myanmar             QFontDatabase__WritingSystem = 22
	QFontDatabase__WritingSystem__Georgian            QFontDatabase__WritingSystem = 23
	QFontDatabase__WritingSystem__Khmer               QFontDatabase__WritingSystem = 24
	QFontDatabase__WritingSystem__SimplifiedChinese   QFontDatabase__WritingSystem = 25
	QFontDatabase__WritingSystem__TraditionalChinese  QFontDatabase__WritingSystem = 26
	QFontDatabase__WritingSystem__Japanese            QFontDatabase__WritingSystem = 27
	QFontDatabase__WritingSystem__Korean              QFontDatabase__WritingSystem = 28
	QFontDatabase__WritingSystem__Vietnamese          QFontDatabase__WritingSystem = 29
	QFontDatabase__WritingSystem__Symbol              QFontDatabase__WritingSystem = 30
	QFontDatabase__WritingSystem__Other               QFontDatabase__WritingSystem = 30
	QFontDatabase__WritingSystem__Ogham               QFontDatabase__WritingSystem = 31
	QFontDatabase__WritingSystem__Runic               QFontDatabase__WritingSystem = 32
	QFontDatabase__WritingSystem__Nko                 QFontDatabase__WritingSystem = 33
	QFontDatabase__WritingSystem__WritingSystemsCount QFontDatabase__WritingSystem = 34
)

type QFontDatabase__SystemFont int

const (
	QFontDatabase__SystemFont__GeneralFont          QFontDatabase__SystemFont = 0
	QFontDatabase__SystemFont__FixedFont            QFontDatabase__SystemFont = 1
	QFontDatabase__SystemFont__TitleFont            QFontDatabase__SystemFont = 2
	QFontDatabase__SystemFont__SmallestReadableFont QFontDatabase__SystemFont = 3
)

type QFontDatabase struct {
	h *C.QFontDatabase
}

func (this *QFontDatabase) cPointer() *C.QFontDatabase {
	if this == nil {
		return nil
	}
	return this.h
}

func newQFontDatabase(h *C.QFontDatabase) *QFontDatabase {
	if h == nil {
		return nil
	}
	return &QFontDatabase{h: h}
}

func newQFontDatabase_U(h unsafe.Pointer) *QFontDatabase {
	return newQFontDatabase((*C.QFontDatabase)(h))
}

// NewQFontDatabase constructs a new QFontDatabase object.
func NewQFontDatabase() *QFontDatabase {
	ret := C.QFontDatabase_new()
	return newQFontDatabase(ret)
}

func QFontDatabase_StandardSizes() []int {
	var _ma *C.struct_miqt_array = C.QFontDatabase_StandardSizes()
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // mrs jackson
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) WritingSystems() []QFontDatabase__WritingSystem {
	var _ma *C.struct_miqt_array = C.QFontDatabase_WritingSystems(this.h)
	_ret := make([]QFontDatabase__WritingSystem, int(_ma.len))
	_outCast := (*[0xffff]C.uintptr_t)(unsafe.Pointer(_ma.data)) // mrs jackson
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QFontDatabase__WritingSystem)(_outCast[i])
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) WritingSystemsWithFamily(family string) []QFontDatabase__WritingSystem {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	var _ma *C.struct_miqt_array = C.QFontDatabase_WritingSystemsWithFamily(this.h, (*C.struct_miqt_string)(family_ms))
	_ret := make([]QFontDatabase__WritingSystem, int(_ma.len))
	_outCast := (*[0xffff]C.uintptr_t)(unsafe.Pointer(_ma.data)) // mrs jackson
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QFontDatabase__WritingSystem)(_outCast[i])
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) Families() []string {
	var _ma *C.struct_miqt_array = C.QFontDatabase_Families(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]*C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = C.GoStringN(&_outCast[i].data, C.int(int64(_outCast[i].len)))
		C.free(unsafe.Pointer(_outCast[i])) // free the inner miqt_string*
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) Styles(family string) []string {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	var _ma *C.struct_miqt_array = C.QFontDatabase_Styles(this.h, (*C.struct_miqt_string)(family_ms))
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]*C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = C.GoStringN(&_outCast[i].data, C.int(int64(_outCast[i].len)))
		C.free(unsafe.Pointer(_outCast[i])) // free the inner miqt_string*
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) PointSizes(family string) []int {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	var _ma *C.struct_miqt_array = C.QFontDatabase_PointSizes(this.h, (*C.struct_miqt_string)(family_ms))
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // mrs jackson
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) SmoothSizes(family string, style string) []int {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	var _ma *C.struct_miqt_array = C.QFontDatabase_SmoothSizes(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // mrs jackson
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) StyleString(font *QFont) string {
	var _ms *C.struct_miqt_string = C.QFontDatabase_StyleString(this.h, font.cPointer())
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QFontDatabase) StyleStringWithFontInfo(fontInfo *QFontInfo) string {
	var _ms *C.struct_miqt_string = C.QFontDatabase_StyleStringWithFontInfo(this.h, fontInfo.cPointer())
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QFontDatabase) Font(family string, style string, pointSize int) *QFont {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_Font(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms), (C.int)(pointSize))
	_goptr := newQFont(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDatabase) IsBitmapScalable(family string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	_ret := C.QFontDatabase_IsBitmapScalable(this.h, (*C.struct_miqt_string)(family_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) IsSmoothlyScalable(family string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	_ret := C.QFontDatabase_IsSmoothlyScalable(this.h, (*C.struct_miqt_string)(family_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) IsScalable(family string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	_ret := C.QFontDatabase_IsScalable(this.h, (*C.struct_miqt_string)(family_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) IsFixedPitch(family string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	_ret := C.QFontDatabase_IsFixedPitch(this.h, (*C.struct_miqt_string)(family_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) Italic(family string, style string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_Italic(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) Bold(family string, style string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_Bold(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) Weight(family string, style string) int {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_Weight(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	return (int)(_ret)
}

func (this *QFontDatabase) HasFamily(family string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	_ret := C.QFontDatabase_HasFamily(this.h, (*C.struct_miqt_string)(family_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) IsPrivateFamily(family string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	_ret := C.QFontDatabase_IsPrivateFamily(this.h, (*C.struct_miqt_string)(family_ms))
	return (bool)(_ret)
}

func QFontDatabase_WritingSystemName(writingSystem QFontDatabase__WritingSystem) string {
	var _ms *C.struct_miqt_string = C.QFontDatabase_WritingSystemName((C.uintptr_t)(writingSystem))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QFontDatabase_WritingSystemSample(writingSystem QFontDatabase__WritingSystem) string {
	var _ms *C.struct_miqt_string = C.QFontDatabase_WritingSystemSample((C.uintptr_t)(writingSystem))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QFontDatabase_AddApplicationFont(fileName string) int {
	fileName_ms := miqt_strdupg(fileName)
	defer C.free(fileName_ms)
	_ret := C.QFontDatabase_AddApplicationFont((*C.struct_miqt_string)(fileName_ms))
	return (int)(_ret)
}

func QFontDatabase_AddApplicationFontFromData(fontData *QByteArray) int {
	_ret := C.QFontDatabase_AddApplicationFontFromData(fontData.cPointer())
	return (int)(_ret)
}

func QFontDatabase_ApplicationFontFamilies(id int) []string {
	var _ma *C.struct_miqt_array = C.QFontDatabase_ApplicationFontFamilies((C.int)(id))
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]*C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = C.GoStringN(&_outCast[i].data, C.int(int64(_outCast[i].len)))
		C.free(unsafe.Pointer(_outCast[i])) // free the inner miqt_string*
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func QFontDatabase_RemoveApplicationFont(id int) bool {
	_ret := C.QFontDatabase_RemoveApplicationFont((C.int)(id))
	return (bool)(_ret)
}

func QFontDatabase_RemoveAllApplicationFonts() bool {
	_ret := C.QFontDatabase_RemoveAllApplicationFonts()
	return (bool)(_ret)
}

func QFontDatabase_SupportsThreadedFontRendering() bool {
	_ret := C.QFontDatabase_SupportsThreadedFontRendering()
	return (bool)(_ret)
}

func QFontDatabase_SystemFont(typeVal QFontDatabase__SystemFont) *QFont {
	_ret := C.QFontDatabase_SystemFont((C.uintptr_t)(typeVal))
	_goptr := newQFont(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDatabase) Families1(writingSystem QFontDatabase__WritingSystem) []string {
	var _ma *C.struct_miqt_array = C.QFontDatabase_Families1(this.h, (C.uintptr_t)(writingSystem))
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]*C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = C.GoStringN(&_outCast[i].data, C.int(int64(_outCast[i].len)))
		C.free(unsafe.Pointer(_outCast[i])) // free the inner miqt_string*
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) PointSizes2(family string, style string) []int {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	var _ma *C.struct_miqt_array = C.QFontDatabase_PointSizes2(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // mrs jackson
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QFontDatabase) IsBitmapScalable2(family string, style string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_IsBitmapScalable2(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) IsSmoothlyScalable2(family string, style string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_IsSmoothlyScalable2(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) IsScalable2(family string, style string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_IsScalable2(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	return (bool)(_ret)
}

func (this *QFontDatabase) IsFixedPitch2(family string, style string) bool {
	family_ms := miqt_strdupg(family)
	defer C.free(family_ms)
	style_ms := miqt_strdupg(style)
	defer C.free(style_ms)
	_ret := C.QFontDatabase_IsFixedPitch2(this.h, (*C.struct_miqt_string)(family_ms), (*C.struct_miqt_string)(style_ms))
	return (bool)(_ret)
}

// Delete this object from C++ memory.
func (this *QFontDatabase) Delete() {
	C.QFontDatabase_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QFontDatabase) GoGC() {
	runtime.SetFinalizer(this, func(this *QFontDatabase) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
