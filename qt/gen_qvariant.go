package qt

/*

#include "gen_qvariant.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QVariant__Type int

const (
	QVariant__Type__Invalid              QVariant__Type = 0
	QVariant__Type__Bool                 QVariant__Type = 1
	QVariant__Type__Int                  QVariant__Type = 2
	QVariant__Type__UInt                 QVariant__Type = 3
	QVariant__Type__LongLong             QVariant__Type = 4
	QVariant__Type__ULongLong            QVariant__Type = 5
	QVariant__Type__Double               QVariant__Type = 6
	QVariant__Type__Char                 QVariant__Type = 7
	QVariant__Type__Map                  QVariant__Type = 8
	QVariant__Type__List                 QVariant__Type = 9
	QVariant__Type__String               QVariant__Type = 10
	QVariant__Type__StringList           QVariant__Type = 11
	QVariant__Type__ByteArray            QVariant__Type = 12
	QVariant__Type__BitArray             QVariant__Type = 13
	QVariant__Type__Date                 QVariant__Type = 14
	QVariant__Type__Time                 QVariant__Type = 15
	QVariant__Type__DateTime             QVariant__Type = 16
	QVariant__Type__Url                  QVariant__Type = 17
	QVariant__Type__Locale               QVariant__Type = 18
	QVariant__Type__Rect                 QVariant__Type = 19
	QVariant__Type__RectF                QVariant__Type = 20
	QVariant__Type__Size                 QVariant__Type = 21
	QVariant__Type__SizeF                QVariant__Type = 22
	QVariant__Type__Line                 QVariant__Type = 23
	QVariant__Type__LineF                QVariant__Type = 24
	QVariant__Type__Point                QVariant__Type = 25
	QVariant__Type__PointF               QVariant__Type = 26
	QVariant__Type__RegExp               QVariant__Type = 27
	QVariant__Type__RegularExpression    QVariant__Type = 44
	QVariant__Type__Hash                 QVariant__Type = 28
	QVariant__Type__EasingCurve          QVariant__Type = 29
	QVariant__Type__Uuid                 QVariant__Type = 30
	QVariant__Type__ModelIndex           QVariant__Type = 42
	QVariant__Type__PersistentModelIndex QVariant__Type = 50
	QVariant__Type__LastCoreType         QVariant__Type = 55
	QVariant__Type__Font                 QVariant__Type = 64
	QVariant__Type__Pixmap               QVariant__Type = 65
	QVariant__Type__Brush                QVariant__Type = 66
	QVariant__Type__Color                QVariant__Type = 67
	QVariant__Type__Palette              QVariant__Type = 68
	QVariant__Type__Image                QVariant__Type = 70
	QVariant__Type__Polygon              QVariant__Type = 71
	QVariant__Type__Region               QVariant__Type = 72
	QVariant__Type__Bitmap               QVariant__Type = 73
	QVariant__Type__Cursor               QVariant__Type = 74
	QVariant__Type__KeySequence          QVariant__Type = 75
	QVariant__Type__Pen                  QVariant__Type = 76
	QVariant__Type__TextLength           QVariant__Type = 77
	QVariant__Type__TextFormat           QVariant__Type = 78
	QVariant__Type__Matrix               QVariant__Type = 79
	QVariant__Type__Transform            QVariant__Type = 80
	QVariant__Type__Matrix4x4            QVariant__Type = 81
	QVariant__Type__Vector2D             QVariant__Type = 82
	QVariant__Type__Vector3D             QVariant__Type = 83
	QVariant__Type__Vector4D             QVariant__Type = 84
	QVariant__Type__Quaternion           QVariant__Type = 85
	QVariant__Type__PolygonF             QVariant__Type = 86
	QVariant__Type__Icon                 QVariant__Type = 69
	QVariant__Type__LastGuiType          QVariant__Type = 87
	QVariant__Type__SizePolicy           QVariant__Type = 121
	QVariant__Type__UserType             QVariant__Type = 1024
	QVariant__Type__LastType             QVariant__Type = 4294967295
)

type QVariant struct {
	h *C.QVariant
}

func (this *QVariant) cPointer() *C.QVariant {
	if this == nil {
		return nil
	}
	return this.h
}

func newQVariant(h *C.QVariant) *QVariant {
	if h == nil {
		return nil
	}
	return &QVariant{h: h}
}

func newQVariant_U(h unsafe.Pointer) *QVariant {
	return newQVariant((*C.QVariant)(h))
}

// NewQVariant constructs a new QVariant object.
func NewQVariant() *QVariant {
	ret := C.QVariant_new()
	return newQVariant(ret)
}

// NewQVariant2 constructs a new QVariant object.
func NewQVariant2(typeVal QVariant__Type) *QVariant {
	ret := C.QVariant_new2((C.int)(typeVal))
	return newQVariant(ret)
}

// NewQVariant3 constructs a new QVariant object.
func NewQVariant3(other *QVariant) *QVariant {
	ret := C.QVariant_new3(other.cPointer())
	return newQVariant(ret)
}

// NewQVariant4 constructs a new QVariant object.
func NewQVariant4(s *QDataStream) *QVariant {
	ret := C.QVariant_new4(s.cPointer())
	return newQVariant(ret)
}

// NewQVariant5 constructs a new QVariant object.
func NewQVariant5(i int) *QVariant {
	ret := C.QVariant_new5((C.int)(i))
	return newQVariant(ret)
}

// NewQVariant6 constructs a new QVariant object.
func NewQVariant6(ui uint) *QVariant {
	ret := C.QVariant_new6((C.uint)(ui))
	return newQVariant(ret)
}

// NewQVariant7 constructs a new QVariant object.
func NewQVariant7(ll int64) *QVariant {
	ret := C.QVariant_new7((C.longlong)(ll))
	return newQVariant(ret)
}

// NewQVariant8 constructs a new QVariant object.
func NewQVariant8(ull uint64) *QVariant {
	ret := C.QVariant_new8((C.ulonglong)(ull))
	return newQVariant(ret)
}

// NewQVariant9 constructs a new QVariant object.
func NewQVariant9(b bool) *QVariant {
	ret := C.QVariant_new9((C.bool)(b))
	return newQVariant(ret)
}

// NewQVariant10 constructs a new QVariant object.
func NewQVariant10(d float64) *QVariant {
	ret := C.QVariant_new10((C.double)(d))
	return newQVariant(ret)
}

// NewQVariant11 constructs a new QVariant object.
func NewQVariant11(f float32) *QVariant {
	ret := C.QVariant_new11((C.float)(f))
	return newQVariant(ret)
}

// NewQVariant12 constructs a new QVariant object.
func NewQVariant12(str string) *QVariant {
	str_Cstring := C.CString(str)
	defer C.free(unsafe.Pointer(str_Cstring))
	ret := C.QVariant_new12(str_Cstring)
	return newQVariant(ret)
}

// NewQVariant13 constructs a new QVariant object.
func NewQVariant13(bytearray *QByteArray) *QVariant {
	ret := C.QVariant_new13(bytearray.cPointer())
	return newQVariant(ret)
}

// NewQVariant14 constructs a new QVariant object.
func NewQVariant14(bitarray *QBitArray) *QVariant {
	ret := C.QVariant_new14(bitarray.cPointer())
	return newQVariant(ret)
}

// NewQVariant15 constructs a new QVariant object.
func NewQVariant15(stringVal string) *QVariant {
	stringVal_ms := miqt_strdupg(stringVal)
	defer C.free(stringVal_ms)
	ret := C.QVariant_new15((*C.struct_miqt_string)(stringVal_ms))
	return newQVariant(ret)
}

// NewQVariant16 constructs a new QVariant object.
func NewQVariant16(stringlist []string) *QVariant {
	// For the C ABI, malloc a C array of raw pointers
	stringlist_CArray := (*[0xffff]*C.struct_miqt_string)(C.malloc(C.size_t(8 * len(stringlist))))
	defer C.free(unsafe.Pointer(stringlist_CArray))
	for i := range stringlist {
		stringlist_i_ms := miqt_strdupg(stringlist[i])
		defer C.free(stringlist_i_ms)
		stringlist_CArray[i] = (*C.struct_miqt_string)(stringlist_i_ms)
	}
	stringlist_ma := &C.struct_miqt_array{len: C.size_t(len(stringlist)), data: unsafe.Pointer(stringlist_CArray)}
	defer runtime.KeepAlive(unsafe.Pointer(stringlist_ma))
	ret := C.QVariant_new16(stringlist_ma)
	return newQVariant(ret)
}

// NewQVariant17 constructs a new QVariant object.
func NewQVariant17(qchar QChar) *QVariant {
	ret := C.QVariant_new17(qchar.cPointer())
	return newQVariant(ret)
}

// NewQVariant18 constructs a new QVariant object.
func NewQVariant18(date *QDate) *QVariant {
	ret := C.QVariant_new18(date.cPointer())
	return newQVariant(ret)
}

// NewQVariant19 constructs a new QVariant object.
func NewQVariant19(time *QTime) *QVariant {
	ret := C.QVariant_new19(time.cPointer())
	return newQVariant(ret)
}

// NewQVariant20 constructs a new QVariant object.
func NewQVariant20(datetime *QDateTime) *QVariant {
	ret := C.QVariant_new20(datetime.cPointer())
	return newQVariant(ret)
}

// NewQVariant21 constructs a new QVariant object.
func NewQVariant21(size *QSize) *QVariant {
	ret := C.QVariant_new21(size.cPointer())
	return newQVariant(ret)
}

// NewQVariant22 constructs a new QVariant object.
func NewQVariant22(size *QSizeF) *QVariant {
	ret := C.QVariant_new22(size.cPointer())
	return newQVariant(ret)
}

// NewQVariant23 constructs a new QVariant object.
func NewQVariant23(pt *QPoint) *QVariant {
	ret := C.QVariant_new23(pt.cPointer())
	return newQVariant(ret)
}

// NewQVariant24 constructs a new QVariant object.
func NewQVariant24(pt *QPointF) *QVariant {
	ret := C.QVariant_new24(pt.cPointer())
	return newQVariant(ret)
}

// NewQVariant25 constructs a new QVariant object.
func NewQVariant25(line *QLine) *QVariant {
	ret := C.QVariant_new25(line.cPointer())
	return newQVariant(ret)
}

// NewQVariant26 constructs a new QVariant object.
func NewQVariant26(line *QLineF) *QVariant {
	ret := C.QVariant_new26(line.cPointer())
	return newQVariant(ret)
}

// NewQVariant27 constructs a new QVariant object.
func NewQVariant27(rect *QRect) *QVariant {
	ret := C.QVariant_new27(rect.cPointer())
	return newQVariant(ret)
}

// NewQVariant28 constructs a new QVariant object.
func NewQVariant28(rect *QRectF) *QVariant {
	ret := C.QVariant_new28(rect.cPointer())
	return newQVariant(ret)
}

// NewQVariant29 constructs a new QVariant object.
func NewQVariant29(locale *QLocale) *QVariant {
	ret := C.QVariant_new29(locale.cPointer())
	return newQVariant(ret)
}

// NewQVariant30 constructs a new QVariant object.
func NewQVariant30(regExp *QRegExp) *QVariant {
	ret := C.QVariant_new30(regExp.cPointer())
	return newQVariant(ret)
}

// NewQVariant31 constructs a new QVariant object.
func NewQVariant31(re *QRegularExpression) *QVariant {
	ret := C.QVariant_new31(re.cPointer())
	return newQVariant(ret)
}

// NewQVariant32 constructs a new QVariant object.
func NewQVariant32(easing *QEasingCurve) *QVariant {
	ret := C.QVariant_new32(easing.cPointer())
	return newQVariant(ret)
}

// NewQVariant33 constructs a new QVariant object.
func NewQVariant33(uuid *QUuid) *QVariant {
	ret := C.QVariant_new33(uuid.cPointer())
	return newQVariant(ret)
}

// NewQVariant34 constructs a new QVariant object.
func NewQVariant34(url *QUrl) *QVariant {
	ret := C.QVariant_new34(url.cPointer())
	return newQVariant(ret)
}

// NewQVariant35 constructs a new QVariant object.
func NewQVariant35(jsonValue *QJsonValue) *QVariant {
	ret := C.QVariant_new35(jsonValue.cPointer())
	return newQVariant(ret)
}

// NewQVariant36 constructs a new QVariant object.
func NewQVariant36(jsonObject *QJsonObject) *QVariant {
	ret := C.QVariant_new36(jsonObject.cPointer())
	return newQVariant(ret)
}

// NewQVariant37 constructs a new QVariant object.
func NewQVariant37(jsonArray *QJsonArray) *QVariant {
	ret := C.QVariant_new37(jsonArray.cPointer())
	return newQVariant(ret)
}

// NewQVariant38 constructs a new QVariant object.
func NewQVariant38(jsonDocument *QJsonDocument) *QVariant {
	ret := C.QVariant_new38(jsonDocument.cPointer())
	return newQVariant(ret)
}

// NewQVariant39 constructs a new QVariant object.
func NewQVariant39(modelIndex *QModelIndex) *QVariant {
	ret := C.QVariant_new39(modelIndex.cPointer())
	return newQVariant(ret)
}

// NewQVariant40 constructs a new QVariant object.
func NewQVariant40(modelIndex *QPersistentModelIndex) *QVariant {
	ret := C.QVariant_new40(modelIndex.cPointer())
	return newQVariant(ret)
}

func (this *QVariant) OperatorAssign(other *QVariant) {
	C.QVariant_OperatorAssign(this.h, other.cPointer())
}

func (this *QVariant) Swap(other *QVariant) {
	C.QVariant_Swap(this.h, other.cPointer())
}

func (this *QVariant) Type() QVariant__Type {
	return (QVariant__Type)(C.QVariant_Type(this.h))
}

func (this *QVariant) UserType() int {
	return (int)(C.QVariant_UserType(this.h))
}

func (this *QVariant) TypeName() unsafe.Pointer {
	_ret := C.QVariant_TypeName(this.h)
	return (unsafe.Pointer)(_ret)
}

func (this *QVariant) CanConvert(targetTypeId int) bool {
	return (bool)(C.QVariant_CanConvert(this.h, (C.int)(targetTypeId)))
}

func (this *QVariant) Convert(targetTypeId int) bool {
	return (bool)(C.QVariant_Convert(this.h, (C.int)(targetTypeId)))
}

func (this *QVariant) IsValid() bool {
	return (bool)(C.QVariant_IsValid(this.h))
}

func (this *QVariant) IsNull() bool {
	return (bool)(C.QVariant_IsNull(this.h))
}

func (this *QVariant) Clear() {
	C.QVariant_Clear(this.h)
}

func (this *QVariant) Detach() {
	C.QVariant_Detach(this.h)
}

func (this *QVariant) IsDetached() bool {
	return (bool)(C.QVariant_IsDetached(this.h))
}

func (this *QVariant) ToInt() int {
	return (int)(C.QVariant_ToInt(this.h))
}

func (this *QVariant) ToUInt() uint {
	return (uint)(C.QVariant_ToUInt(this.h))
}

func (this *QVariant) ToLongLong() int64 {
	return (int64)(C.QVariant_ToLongLong(this.h))
}

func (this *QVariant) ToULongLong() uint64 {
	return (uint64)(C.QVariant_ToULongLong(this.h))
}

func (this *QVariant) ToBool() bool {
	return (bool)(C.QVariant_ToBool(this.h))
}

func (this *QVariant) ToDouble() float64 {
	return (float64)(C.QVariant_ToDouble(this.h))
}

func (this *QVariant) ToFloat() float32 {
	return (float32)(C.QVariant_ToFloat(this.h))
}

func (this *QVariant) ToReal() float64 {
	return (float64)(C.QVariant_ToReal(this.h))
}

func (this *QVariant) ToByteArray() *QByteArray {
	_ret := C.QVariant_ToByteArray(this.h)
	_goptr := newQByteArray(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToBitArray() *QBitArray {
	_ret := C.QVariant_ToBitArray(this.h)
	_goptr := newQBitArray(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToString() string {
	var _ms *C.struct_miqt_string = C.QVariant_ToString(this.h)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QVariant) ToStringList() []string {
	var _ma *C.struct_miqt_array = C.QVariant_ToStringList(this.h)
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

func (this *QVariant) ToChar() *QChar {
	_ret := C.QVariant_ToChar(this.h)
	_goptr := newQChar(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToDate() *QDate {
	_ret := C.QVariant_ToDate(this.h)
	_goptr := newQDate(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToTime() *QTime {
	_ret := C.QVariant_ToTime(this.h)
	_goptr := newQTime(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToDateTime() *QDateTime {
	_ret := C.QVariant_ToDateTime(this.h)
	_goptr := newQDateTime(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToPoint() *QPoint {
	_ret := C.QVariant_ToPoint(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToPointF() *QPointF {
	_ret := C.QVariant_ToPointF(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToRect() *QRect {
	_ret := C.QVariant_ToRect(this.h)
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToSize() *QSize {
	_ret := C.QVariant_ToSize(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToSizeF() *QSizeF {
	_ret := C.QVariant_ToSizeF(this.h)
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToLine() *QLine {
	_ret := C.QVariant_ToLine(this.h)
	_goptr := newQLine(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToLineF() *QLineF {
	_ret := C.QVariant_ToLineF(this.h)
	_goptr := newQLineF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToRectF() *QRectF {
	_ret := C.QVariant_ToRectF(this.h)
	_goptr := newQRectF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToLocale() *QLocale {
	_ret := C.QVariant_ToLocale(this.h)
	_goptr := newQLocale(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToRegExp() *QRegExp {
	_ret := C.QVariant_ToRegExp(this.h)
	_goptr := newQRegExp(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToRegularExpression() *QRegularExpression {
	_ret := C.QVariant_ToRegularExpression(this.h)
	_goptr := newQRegularExpression(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToEasingCurve() *QEasingCurve {
	_ret := C.QVariant_ToEasingCurve(this.h)
	_goptr := newQEasingCurve(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToUuid() *QUuid {
	_ret := C.QVariant_ToUuid(this.h)
	_goptr := newQUuid(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToUrl() *QUrl {
	_ret := C.QVariant_ToUrl(this.h)
	_goptr := newQUrl(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonValue() *QJsonValue {
	_ret := C.QVariant_ToJsonValue(this.h)
	_goptr := newQJsonValue(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonObject() *QJsonObject {
	_ret := C.QVariant_ToJsonObject(this.h)
	_goptr := newQJsonObject(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonArray() *QJsonArray {
	_ret := C.QVariant_ToJsonArray(this.h)
	_goptr := newQJsonArray(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonDocument() *QJsonDocument {
	_ret := C.QVariant_ToJsonDocument(this.h)
	_goptr := newQJsonDocument(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToModelIndex() *QModelIndex {
	_ret := C.QVariant_ToModelIndex(this.h)
	_goptr := newQModelIndex(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToPersistentModelIndex() *QPersistentModelIndex {
	_ret := C.QVariant_ToPersistentModelIndex(this.h)
	_goptr := newQPersistentModelIndex(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) Load(ds *QDataStream) {
	C.QVariant_Load(this.h, ds.cPointer())
}

func (this *QVariant) Save(ds *QDataStream) {
	C.QVariant_Save(this.h, ds.cPointer())
}

func QVariant_TypeToName(typeId int) unsafe.Pointer {
	_ret := C.QVariant_TypeToName((C.int)(typeId))
	return (unsafe.Pointer)(_ret)
}

func QVariant_NameToType(name string) QVariant__Type {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	return (QVariant__Type)(C.QVariant_NameToType(name_Cstring))
}

func (this *QVariant) OperatorEqual(v *QVariant) bool {
	return (bool)(C.QVariant_OperatorEqual(this.h, v.cPointer()))
}

func (this *QVariant) OperatorNotEqual(v *QVariant) bool {
	return (bool)(C.QVariant_OperatorNotEqual(this.h, v.cPointer()))
}

func (this *QVariant) OperatorLesser(v *QVariant) bool {
	return (bool)(C.QVariant_OperatorLesser(this.h, v.cPointer()))
}

func (this *QVariant) OperatorLesserOrEqual(v *QVariant) bool {
	return (bool)(C.QVariant_OperatorLesserOrEqual(this.h, v.cPointer()))
}

func (this *QVariant) OperatorGreater(v *QVariant) bool {
	return (bool)(C.QVariant_OperatorGreater(this.h, v.cPointer()))
}

func (this *QVariant) OperatorGreaterOrEqual(v *QVariant) bool {
	return (bool)(C.QVariant_OperatorGreaterOrEqual(this.h, v.cPointer()))
}

func (this *QVariant) ToInt1(ok *bool) int {
	return (int)(C.QVariant_ToInt1(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToUInt1(ok *bool) uint {
	return (uint)(C.QVariant_ToUInt1(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToLongLong1(ok *bool) int64 {
	return (int64)(C.QVariant_ToLongLong1(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToULongLong1(ok *bool) uint64 {
	return (uint64)(C.QVariant_ToULongLong1(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToDouble1(ok *bool) float64 {
	return (float64)(C.QVariant_ToDouble1(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToFloat1(ok *bool) float32 {
	return (float32)(C.QVariant_ToFloat1(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToReal1(ok *bool) float64 {
	return (float64)(C.QVariant_ToReal1(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

// Delete this object from C++ memory.
func (this *QVariant) Delete() {
	C.QVariant_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QVariant) GoGC() {
	runtime.SetFinalizer(this, func(this *QVariant) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QVariantComparisonHelper struct {
	h *C.QVariantComparisonHelper
}

func (this *QVariantComparisonHelper) cPointer() *C.QVariantComparisonHelper {
	if this == nil {
		return nil
	}
	return this.h
}

func newQVariantComparisonHelper(h *C.QVariantComparisonHelper) *QVariantComparisonHelper {
	if h == nil {
		return nil
	}
	return &QVariantComparisonHelper{h: h}
}

func newQVariantComparisonHelper_U(h unsafe.Pointer) *QVariantComparisonHelper {
	return newQVariantComparisonHelper((*C.QVariantComparisonHelper)(h))
}

// NewQVariantComparisonHelper constructs a new QVariantComparisonHelper object.
func NewQVariantComparisonHelper(varVal *QVariant) *QVariantComparisonHelper {
	ret := C.QVariantComparisonHelper_new(varVal.cPointer())
	return newQVariantComparisonHelper(ret)
}

// NewQVariantComparisonHelper2 constructs a new QVariantComparisonHelper object.
func NewQVariantComparisonHelper2(param1 *QVariantComparisonHelper) *QVariantComparisonHelper {
	ret := C.QVariantComparisonHelper_new2(param1.cPointer())
	return newQVariantComparisonHelper(ret)
}

// Delete this object from C++ memory.
func (this *QVariantComparisonHelper) Delete() {
	C.QVariantComparisonHelper_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QVariantComparisonHelper) GoGC() {
	runtime.SetFinalizer(this, func(this *QVariantComparisonHelper) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QSequentialIterable struct {
	h *C.QSequentialIterable
}

func (this *QSequentialIterable) cPointer() *C.QSequentialIterable {
	if this == nil {
		return nil
	}
	return this.h
}

func newQSequentialIterable(h *C.QSequentialIterable) *QSequentialIterable {
	if h == nil {
		return nil
	}
	return &QSequentialIterable{h: h}
}

func newQSequentialIterable_U(h unsafe.Pointer) *QSequentialIterable {
	return newQSequentialIterable((*C.QSequentialIterable)(h))
}

// NewQSequentialIterable constructs a new QSequentialIterable object.
func NewQSequentialIterable(impl QtMetaTypePrivate__QSequentialIterableImpl) *QSequentialIterable {
	ret := C.QSequentialIterable_new(impl.cPointer())
	return newQSequentialIterable(ret)
}

// NewQSequentialIterable2 constructs a new QSequentialIterable object.
func NewQSequentialIterable2(param1 *QSequentialIterable) *QSequentialIterable {
	ret := C.QSequentialIterable_new2(param1.cPointer())
	return newQSequentialIterable(ret)
}

func (this *QSequentialIterable) Begin() *QSequentialIterable__const_iterator {
	_ret := C.QSequentialIterable_Begin(this.h)
	_goptr := newQSequentialIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSequentialIterable) End() *QSequentialIterable__const_iterator {
	_ret := C.QSequentialIterable_End(this.h)
	_goptr := newQSequentialIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSequentialIterable) At(idx int) *QVariant {
	_ret := C.QSequentialIterable_At(this.h, (C.int)(idx))
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSequentialIterable) Size() int {
	return (int)(C.QSequentialIterable_Size(this.h))
}

func (this *QSequentialIterable) CanReverseIterate() bool {
	return (bool)(C.QSequentialIterable_CanReverseIterate(this.h))
}

// Delete this object from C++ memory.
func (this *QSequentialIterable) Delete() {
	C.QSequentialIterable_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSequentialIterable) GoGC() {
	runtime.SetFinalizer(this, func(this *QSequentialIterable) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QAssociativeIterable struct {
	h *C.QAssociativeIterable
}

func (this *QAssociativeIterable) cPointer() *C.QAssociativeIterable {
	if this == nil {
		return nil
	}
	return this.h
}

func newQAssociativeIterable(h *C.QAssociativeIterable) *QAssociativeIterable {
	if h == nil {
		return nil
	}
	return &QAssociativeIterable{h: h}
}

func newQAssociativeIterable_U(h unsafe.Pointer) *QAssociativeIterable {
	return newQAssociativeIterable((*C.QAssociativeIterable)(h))
}

// NewQAssociativeIterable constructs a new QAssociativeIterable object.
func NewQAssociativeIterable(impl QtMetaTypePrivate__QAssociativeIterableImpl) *QAssociativeIterable {
	ret := C.QAssociativeIterable_new(impl.cPointer())
	return newQAssociativeIterable(ret)
}

// NewQAssociativeIterable2 constructs a new QAssociativeIterable object.
func NewQAssociativeIterable2(param1 *QAssociativeIterable) *QAssociativeIterable {
	ret := C.QAssociativeIterable_new2(param1.cPointer())
	return newQAssociativeIterable(ret)
}

func (this *QAssociativeIterable) Begin() *QAssociativeIterable__const_iterator {
	_ret := C.QAssociativeIterable_Begin(this.h)
	_goptr := newQAssociativeIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable) End() *QAssociativeIterable__const_iterator {
	_ret := C.QAssociativeIterable_End(this.h)
	_goptr := newQAssociativeIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable) Find(key *QVariant) *QAssociativeIterable__const_iterator {
	_ret := C.QAssociativeIterable_Find(this.h, key.cPointer())
	_goptr := newQAssociativeIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable) Value(key *QVariant) *QVariant {
	_ret := C.QAssociativeIterable_Value(this.h, key.cPointer())
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable) Size() int {
	return (int)(C.QAssociativeIterable_Size(this.h))
}

// Delete this object from C++ memory.
func (this *QAssociativeIterable) Delete() {
	C.QAssociativeIterable_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAssociativeIterable) GoGC() {
	runtime.SetFinalizer(this, func(this *QAssociativeIterable) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QVariant__PrivateShared struct {
	h *C.QVariant__PrivateShared
}

func (this *QVariant__PrivateShared) cPointer() *C.QVariant__PrivateShared {
	if this == nil {
		return nil
	}
	return this.h
}

func newQVariant__PrivateShared(h *C.QVariant__PrivateShared) *QVariant__PrivateShared {
	if h == nil {
		return nil
	}
	return &QVariant__PrivateShared{h: h}
}

func newQVariant__PrivateShared_U(h unsafe.Pointer) *QVariant__PrivateShared {
	return newQVariant__PrivateShared((*C.QVariant__PrivateShared)(h))
}

// Delete this object from C++ memory.
func (this *QVariant__PrivateShared) Delete() {
	C.QVariant__PrivateShared_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QVariant__PrivateShared) GoGC() {
	runtime.SetFinalizer(this, func(this *QVariant__PrivateShared) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QVariant__Handler struct {
	h *C.QVariant__Handler
}

func (this *QVariant__Handler) cPointer() *C.QVariant__Handler {
	if this == nil {
		return nil
	}
	return this.h
}

func newQVariant__Handler(h *C.QVariant__Handler) *QVariant__Handler {
	if h == nil {
		return nil
	}
	return &QVariant__Handler{h: h}
}

func newQVariant__Handler_U(h unsafe.Pointer) *QVariant__Handler {
	return newQVariant__Handler((*C.QVariant__Handler)(h))
}

// Delete this object from C++ memory.
func (this *QVariant__Handler) Delete() {
	C.QVariant__Handler_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QVariant__Handler) GoGC() {
	runtime.SetFinalizer(this, func(this *QVariant__Handler) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QSequentialIterable__const_iterator struct {
	h *C.QSequentialIterable__const_iterator
}

func (this *QSequentialIterable__const_iterator) cPointer() *C.QSequentialIterable__const_iterator {
	if this == nil {
		return nil
	}
	return this.h
}

func newQSequentialIterable__const_iterator(h *C.QSequentialIterable__const_iterator) *QSequentialIterable__const_iterator {
	if h == nil {
		return nil
	}
	return &QSequentialIterable__const_iterator{h: h}
}

func newQSequentialIterable__const_iterator_U(h unsafe.Pointer) *QSequentialIterable__const_iterator {
	return newQSequentialIterable__const_iterator((*C.QSequentialIterable__const_iterator)(h))
}

// NewQSequentialIterable__const_iterator constructs a new QSequentialIterable::const_iterator object.
func NewQSequentialIterable__const_iterator(other *QSequentialIterable__const_iterator) *QSequentialIterable__const_iterator {
	ret := C.QSequentialIterable__const_iterator_new(other.cPointer())
	return newQSequentialIterable__const_iterator(ret)
}

func (this *QSequentialIterable__const_iterator) OperatorAssign(other *QSequentialIterable__const_iterator) {
	C.QSequentialIterable__const_iterator_OperatorAssign(this.h, other.cPointer())
}

func (this *QSequentialIterable__const_iterator) OperatorMultiply() *QVariant {
	_ret := C.QSequentialIterable__const_iterator_OperatorMultiply(this.h)
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSequentialIterable__const_iterator) OperatorEqual(o *QSequentialIterable__const_iterator) bool {
	return (bool)(C.QSequentialIterable__const_iterator_OperatorEqual(this.h, o.cPointer()))
}

func (this *QSequentialIterable__const_iterator) OperatorNotEqual(o *QSequentialIterable__const_iterator) bool {
	return (bool)(C.QSequentialIterable__const_iterator_OperatorNotEqual(this.h, o.cPointer()))
}

func (this *QSequentialIterable__const_iterator) OperatorPlusPlus() *QSequentialIterable__const_iterator {
	return newQSequentialIterable__const_iterator_U(unsafe.Pointer(C.QSequentialIterable__const_iterator_OperatorPlusPlus(this.h)))
}

func (this *QSequentialIterable__const_iterator) OperatorPlusPlusWithInt(param1 int) *QSequentialIterable__const_iterator {
	_ret := C.QSequentialIterable__const_iterator_OperatorPlusPlusWithInt(this.h, (C.int)(param1))
	_goptr := newQSequentialIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSequentialIterable__const_iterator) OperatorMinusMinus() *QSequentialIterable__const_iterator {
	return newQSequentialIterable__const_iterator_U(unsafe.Pointer(C.QSequentialIterable__const_iterator_OperatorMinusMinus(this.h)))
}

func (this *QSequentialIterable__const_iterator) OperatorMinusMinusWithInt(param1 int) *QSequentialIterable__const_iterator {
	_ret := C.QSequentialIterable__const_iterator_OperatorMinusMinusWithInt(this.h, (C.int)(param1))
	_goptr := newQSequentialIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSequentialIterable__const_iterator) OperatorPlusAssign(j int) *QSequentialIterable__const_iterator {
	return newQSequentialIterable__const_iterator_U(unsafe.Pointer(C.QSequentialIterable__const_iterator_OperatorPlusAssign(this.h, (C.int)(j))))
}

func (this *QSequentialIterable__const_iterator) OperatorMinusAssign(j int) *QSequentialIterable__const_iterator {
	return newQSequentialIterable__const_iterator_U(unsafe.Pointer(C.QSequentialIterable__const_iterator_OperatorMinusAssign(this.h, (C.int)(j))))
}

func (this *QSequentialIterable__const_iterator) OperatorPlus(j int) *QSequentialIterable__const_iterator {
	_ret := C.QSequentialIterable__const_iterator_OperatorPlus(this.h, (C.int)(j))
	_goptr := newQSequentialIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSequentialIterable__const_iterator) OperatorMinus(j int) *QSequentialIterable__const_iterator {
	_ret := C.QSequentialIterable__const_iterator_OperatorMinus(this.h, (C.int)(j))
	_goptr := newQSequentialIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QSequentialIterable__const_iterator) Delete() {
	C.QSequentialIterable__const_iterator_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSequentialIterable__const_iterator) GoGC() {
	runtime.SetFinalizer(this, func(this *QSequentialIterable__const_iterator) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QAssociativeIterable__const_iterator struct {
	h *C.QAssociativeIterable__const_iterator
}

func (this *QAssociativeIterable__const_iterator) cPointer() *C.QAssociativeIterable__const_iterator {
	if this == nil {
		return nil
	}
	return this.h
}

func newQAssociativeIterable__const_iterator(h *C.QAssociativeIterable__const_iterator) *QAssociativeIterable__const_iterator {
	if h == nil {
		return nil
	}
	return &QAssociativeIterable__const_iterator{h: h}
}

func newQAssociativeIterable__const_iterator_U(h unsafe.Pointer) *QAssociativeIterable__const_iterator {
	return newQAssociativeIterable__const_iterator((*C.QAssociativeIterable__const_iterator)(h))
}

// NewQAssociativeIterable__const_iterator constructs a new QAssociativeIterable::const_iterator object.
func NewQAssociativeIterable__const_iterator(other *QAssociativeIterable__const_iterator) *QAssociativeIterable__const_iterator {
	ret := C.QAssociativeIterable__const_iterator_new(other.cPointer())
	return newQAssociativeIterable__const_iterator(ret)
}

func (this *QAssociativeIterable__const_iterator) OperatorAssign(other *QAssociativeIterable__const_iterator) {
	C.QAssociativeIterable__const_iterator_OperatorAssign(this.h, other.cPointer())
}

func (this *QAssociativeIterable__const_iterator) Key() *QVariant {
	_ret := C.QAssociativeIterable__const_iterator_Key(this.h)
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable__const_iterator) Value() *QVariant {
	_ret := C.QAssociativeIterable__const_iterator_Value(this.h)
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable__const_iterator) OperatorMultiply() *QVariant {
	_ret := C.QAssociativeIterable__const_iterator_OperatorMultiply(this.h)
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable__const_iterator) OperatorEqual(o *QAssociativeIterable__const_iterator) bool {
	return (bool)(C.QAssociativeIterable__const_iterator_OperatorEqual(this.h, o.cPointer()))
}

func (this *QAssociativeIterable__const_iterator) OperatorNotEqual(o *QAssociativeIterable__const_iterator) bool {
	return (bool)(C.QAssociativeIterable__const_iterator_OperatorNotEqual(this.h, o.cPointer()))
}

func (this *QAssociativeIterable__const_iterator) OperatorPlusPlus() *QAssociativeIterable__const_iterator {
	return newQAssociativeIterable__const_iterator_U(unsafe.Pointer(C.QAssociativeIterable__const_iterator_OperatorPlusPlus(this.h)))
}

func (this *QAssociativeIterable__const_iterator) OperatorPlusPlusWithInt(param1 int) *QAssociativeIterable__const_iterator {
	_ret := C.QAssociativeIterable__const_iterator_OperatorPlusPlusWithInt(this.h, (C.int)(param1))
	_goptr := newQAssociativeIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable__const_iterator) OperatorMinusMinus() *QAssociativeIterable__const_iterator {
	return newQAssociativeIterable__const_iterator_U(unsafe.Pointer(C.QAssociativeIterable__const_iterator_OperatorMinusMinus(this.h)))
}

func (this *QAssociativeIterable__const_iterator) OperatorMinusMinusWithInt(param1 int) *QAssociativeIterable__const_iterator {
	_ret := C.QAssociativeIterable__const_iterator_OperatorMinusMinusWithInt(this.h, (C.int)(param1))
	_goptr := newQAssociativeIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable__const_iterator) OperatorPlusAssign(j int) *QAssociativeIterable__const_iterator {
	return newQAssociativeIterable__const_iterator_U(unsafe.Pointer(C.QAssociativeIterable__const_iterator_OperatorPlusAssign(this.h, (C.int)(j))))
}

func (this *QAssociativeIterable__const_iterator) OperatorMinusAssign(j int) *QAssociativeIterable__const_iterator {
	return newQAssociativeIterable__const_iterator_U(unsafe.Pointer(C.QAssociativeIterable__const_iterator_OperatorMinusAssign(this.h, (C.int)(j))))
}

func (this *QAssociativeIterable__const_iterator) OperatorPlus(j int) *QAssociativeIterable__const_iterator {
	_ret := C.QAssociativeIterable__const_iterator_OperatorPlus(this.h, (C.int)(j))
	_goptr := newQAssociativeIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAssociativeIterable__const_iterator) OperatorMinus(j int) *QAssociativeIterable__const_iterator {
	_ret := C.QAssociativeIterable__const_iterator_OperatorMinus(this.h, (C.int)(j))
	_goptr := newQAssociativeIterable__const_iterator(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QAssociativeIterable__const_iterator) Delete() {
	C.QAssociativeIterable__const_iterator_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAssociativeIterable__const_iterator) GoGC() {
	runtime.SetFinalizer(this, func(this *QAssociativeIterable__const_iterator) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
