package qt6

/*

#include "gen_qevent.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QWheelEvent__ int

const (
	QWheelEvent__DefaultDeltasPerStep QWheelEvent__ = 120
)

type QPlatformSurfaceEvent__SurfaceEventType int

const (
	QPlatformSurfaceEvent__SurfaceCreated            QPlatformSurfaceEvent__SurfaceEventType = 0
	QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed QPlatformSurfaceEvent__SurfaceEventType = 1
)

type QContextMenuEvent__Reason int

const (
	QContextMenuEvent__Mouse    QContextMenuEvent__Reason = 0
	QContextMenuEvent__Keyboard QContextMenuEvent__Reason = 1
	QContextMenuEvent__Other    QContextMenuEvent__Reason = 2
)

type QInputMethodEvent__AttributeType int

const (
	QInputMethodEvent__TextFormat QInputMethodEvent__AttributeType = 0
	QInputMethodEvent__Cursor     QInputMethodEvent__AttributeType = 1
	QInputMethodEvent__Language   QInputMethodEvent__AttributeType = 2
	QInputMethodEvent__Ruby       QInputMethodEvent__AttributeType = 3
	QInputMethodEvent__Selection  QInputMethodEvent__AttributeType = 4
)

type QScrollEvent__ScrollState int

const (
	QScrollEvent__ScrollStarted  QScrollEvent__ScrollState = 0
	QScrollEvent__ScrollUpdated  QScrollEvent__ScrollState = 1
	QScrollEvent__ScrollFinished QScrollEvent__ScrollState = 2
)

type QInputEvent struct {
	h          *C.QInputEvent
	isSubclass bool
	*QEvent
}

func (this *QInputEvent) cPointer() *C.QInputEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QInputEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQInputEvent constructs the type using only CGO pointers.
func newQInputEvent(h *C.QInputEvent, h_QEvent *C.QEvent) *QInputEvent {
	if h == nil {
		return nil
	}
	return &QInputEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQInputEvent constructs the type using only unsafe pointers.
func UnsafeNewQInputEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QInputEvent {
	if h == nil {
		return nil
	}

	return &QInputEvent{h: (*C.QInputEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQInputEvent constructs a new QInputEvent object.
func NewQInputEvent(typeVal QEvent__Type, m_dev *QInputDevice) *QInputEvent {
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QInputEvent_new((C.int)(typeVal), m_dev.cPointer(), &outptr_QInputEvent, &outptr_QEvent)
	ret := newQInputEvent(outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQInputEvent2 constructs a new QInputEvent object.
func NewQInputEvent2(typeVal QEvent__Type, m_dev *QInputDevice, modifiers KeyboardModifier) *QInputEvent {
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QInputEvent_new2((C.int)(typeVal), m_dev.cPointer(), (C.int)(modifiers), &outptr_QInputEvent, &outptr_QEvent)
	ret := newQInputEvent(outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QInputEvent) Clone() *QInputEvent {
	return UnsafeNewQInputEvent(unsafe.Pointer(C.QInputEvent_Clone(this.h)), nil)
}

func (this *QInputEvent) Device() *QInputDevice {
	return UnsafeNewQInputDevice(unsafe.Pointer(C.QInputEvent_Device(this.h)), nil)
}

func (this *QInputEvent) DeviceType() QInputDevice__DeviceType {
	return (QInputDevice__DeviceType)(C.QInputEvent_DeviceType(this.h))
}

func (this *QInputEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(C.QInputEvent_Modifiers(this.h))
}

func (this *QInputEvent) SetModifiers(modifiers KeyboardModifier) {
	C.QInputEvent_SetModifiers(this.h, (C.int)(modifiers))
}

func (this *QInputEvent) Timestamp() uint64 {
	return (uint64)(C.QInputEvent_Timestamp(this.h))
}

func (this *QInputEvent) SetTimestamp(timestamp uint64) {
	C.QInputEvent_SetTimestamp(this.h, (C.ulonglong)(timestamp))
}

func (this *QInputEvent) callVirtualBase_Clone() *QInputEvent {

	return UnsafeNewQInputEvent(unsafe.Pointer(C.QInputEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QInputEvent) OnClone(slot func(super func() *QInputEvent) *QInputEvent) {
	C.QInputEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputEvent_Clone
func miqt_exec_callback_QInputEvent_Clone(self *C.QInputEvent, cb C.intptr_t) *C.QInputEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QInputEvent) *QInputEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QInputEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	C.QInputEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (C.ulonglong)(timestamp))

}
func (this *QInputEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	C.QInputEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputEvent_SetTimestamp
func miqt_exec_callback_QInputEvent_SetTimestamp(self *C.QInputEvent, cb C.intptr_t, timestamp C.ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QInputEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

func (this *QInputEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QInputEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QInputEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QInputEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputEvent_SetAccepted
func miqt_exec_callback_QInputEvent_SetAccepted(self *C.QInputEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QInputEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QInputEvent) Delete() {
	C.QInputEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QInputEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QInputEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QPointerEvent struct {
	h          *C.QPointerEvent
	isSubclass bool
	*QInputEvent
}

func (this *QPointerEvent) cPointer() *C.QPointerEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPointerEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPointerEvent constructs the type using only CGO pointers.
func newQPointerEvent(h *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QPointerEvent {
	if h == nil {
		return nil
	}
	return &QPointerEvent{h: h,
		QInputEvent: newQInputEvent(h_QInputEvent, h_QEvent)}
}

// UnsafeNewQPointerEvent constructs the type using only unsafe pointers.
func UnsafeNewQPointerEvent(h unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QPointerEvent {
	if h == nil {
		return nil
	}

	return &QPointerEvent{h: (*C.QPointerEvent)(h),
		QInputEvent: UnsafeNewQInputEvent(h_QInputEvent, h_QEvent)}
}

// NewQPointerEvent constructs a new QPointerEvent object.
func NewQPointerEvent(typeVal QEvent__Type, dev *QPointingDevice) *QPointerEvent {
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QPointerEvent_new((C.int)(typeVal), dev.cPointer(), &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQPointerEvent(outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQPointerEvent2 constructs a new QPointerEvent object.
func NewQPointerEvent2(typeVal QEvent__Type, dev *QPointingDevice, modifiers KeyboardModifier) *QPointerEvent {
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QPointerEvent_new2((C.int)(typeVal), dev.cPointer(), (C.int)(modifiers), &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQPointerEvent(outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQPointerEvent3 constructs a new QPointerEvent object.
func NewQPointerEvent3(typeVal QEvent__Type, dev *QPointingDevice, modifiers KeyboardModifier, points []QEventPoint) *QPointerEvent {
	points_CArray := (*[0xffff]*C.QEventPoint)(C.malloc(C.size_t(8 * len(points))))
	defer C.free(unsafe.Pointer(points_CArray))
	for i := range points {
		points_CArray[i] = points[i].cPointer()
	}
	points_ma := C.struct_miqt_array{len: C.size_t(len(points)), data: unsafe.Pointer(points_CArray)}
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QPointerEvent_new3((C.int)(typeVal), dev.cPointer(), (C.int)(modifiers), points_ma, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQPointerEvent(outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QPointerEvent) Clone() *QPointerEvent {
	return UnsafeNewQPointerEvent(unsafe.Pointer(C.QPointerEvent_Clone(this.h)), nil, nil)
}

func (this *QPointerEvent) PointingDevice() *QPointingDevice {
	return UnsafeNewQPointingDevice(unsafe.Pointer(C.QPointerEvent_PointingDevice(this.h)), nil, nil)
}

func (this *QPointerEvent) PointerType() QPointingDevice__PointerType {
	return (QPointingDevice__PointerType)(C.QPointerEvent_PointerType(this.h))
}

func (this *QPointerEvent) SetTimestamp(timestamp uint64) {
	C.QPointerEvent_SetTimestamp(this.h, (C.ulonglong)(timestamp))
}

func (this *QPointerEvent) PointCount() int64 {
	return (int64)(C.QPointerEvent_PointCount(this.h))
}

func (this *QPointerEvent) Point(i int64) *QEventPoint {
	return UnsafeNewQEventPoint(unsafe.Pointer(C.QPointerEvent_Point(this.h, (C.ptrdiff_t)(i))))
}

func (this *QPointerEvent) Points() []QEventPoint {
	var _ma C.struct_miqt_array = C.QPointerEvent_Points(this.h)
	_ret := make([]QEventPoint, int(_ma.len))
	_outCast := (*[0xffff]*C.QEventPoint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_ret := _outCast[i]
		_lv_goptr := newQEventPoint(_lv_ret)
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPointerEvent) PointById(id int) *QEventPoint {
	return UnsafeNewQEventPoint(unsafe.Pointer(C.QPointerEvent_PointById(this.h, (C.int)(id))))
}

func (this *QPointerEvent) AllPointsGrabbed() bool {
	return (bool)(C.QPointerEvent_AllPointsGrabbed(this.h))
}

func (this *QPointerEvent) IsBeginEvent() bool {
	return (bool)(C.QPointerEvent_IsBeginEvent(this.h))
}

func (this *QPointerEvent) IsUpdateEvent() bool {
	return (bool)(C.QPointerEvent_IsUpdateEvent(this.h))
}

func (this *QPointerEvent) IsEndEvent() bool {
	return (bool)(C.QPointerEvent_IsEndEvent(this.h))
}

func (this *QPointerEvent) AllPointsAccepted() bool {
	return (bool)(C.QPointerEvent_AllPointsAccepted(this.h))
}

func (this *QPointerEvent) SetAccepted(accepted bool) {
	C.QPointerEvent_SetAccepted(this.h, (C.bool)(accepted))
}

func (this *QPointerEvent) ExclusiveGrabber(point *QEventPoint) *QObject {
	return UnsafeNewQObject(unsafe.Pointer(C.QPointerEvent_ExclusiveGrabber(this.h, point.cPointer())))
}

func (this *QPointerEvent) SetExclusiveGrabber(point *QEventPoint, exclusiveGrabber *QObject) {
	C.QPointerEvent_SetExclusiveGrabber(this.h, point.cPointer(), exclusiveGrabber.cPointer())
}

func (this *QPointerEvent) ClearPassiveGrabbers(point *QEventPoint) {
	C.QPointerEvent_ClearPassiveGrabbers(this.h, point.cPointer())
}

func (this *QPointerEvent) AddPassiveGrabber(point *QEventPoint, grabber *QObject) bool {
	return (bool)(C.QPointerEvent_AddPassiveGrabber(this.h, point.cPointer(), grabber.cPointer()))
}

func (this *QPointerEvent) RemovePassiveGrabber(point *QEventPoint, grabber *QObject) bool {
	return (bool)(C.QPointerEvent_RemovePassiveGrabber(this.h, point.cPointer(), grabber.cPointer()))
}

func (this *QPointerEvent) callVirtualBase_Clone() *QPointerEvent {

	return UnsafeNewQPointerEvent(unsafe.Pointer(C.QPointerEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil)
}
func (this *QPointerEvent) OnClone(slot func(super func() *QPointerEvent) *QPointerEvent) {
	C.QPointerEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_Clone
func miqt_exec_callback_QPointerEvent_Clone(self *C.QPointerEvent, cb C.intptr_t) *C.QPointerEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPointerEvent) *QPointerEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QPointerEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	C.QPointerEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (C.ulonglong)(timestamp))

}
func (this *QPointerEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	C.QPointerEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_SetTimestamp
func miqt_exec_callback_QPointerEvent_SetTimestamp(self *C.QPointerEvent, cb C.intptr_t, timestamp C.ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QPointerEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

func (this *QPointerEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QPointerEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QPointerEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QPointerEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_IsBeginEvent
func miqt_exec_callback_QPointerEvent_IsBeginEvent(self *C.QPointerEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QPointerEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QPointerEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QPointerEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QPointerEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_IsUpdateEvent
func miqt_exec_callback_QPointerEvent_IsUpdateEvent(self *C.QPointerEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QPointerEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QPointerEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QPointerEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QPointerEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_IsEndEvent
func miqt_exec_callback_QPointerEvent_IsEndEvent(self *C.QPointerEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

func (this *QPointerEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QPointerEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QPointerEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QPointerEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_SetAccepted
func miqt_exec_callback_QPointerEvent_SetAccepted(self *C.QPointerEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QPointerEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QPointerEvent) Delete() {
	C.QPointerEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPointerEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QPointerEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QSinglePointEvent struct {
	h          *C.QSinglePointEvent
	isSubclass bool
	*QPointerEvent
}

func (this *QSinglePointEvent) cPointer() *C.QSinglePointEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QSinglePointEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQSinglePointEvent constructs the type using only CGO pointers.
func newQSinglePointEvent(h *C.QSinglePointEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QSinglePointEvent {
	if h == nil {
		return nil
	}
	return &QSinglePointEvent{h: h,
		QPointerEvent: newQPointerEvent(h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQSinglePointEvent constructs the type using only unsafe pointers.
func UnsafeNewQSinglePointEvent(h unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QSinglePointEvent {
	if h == nil {
		return nil
	}

	return &QSinglePointEvent{h: (*C.QSinglePointEvent)(h),
		QPointerEvent: UnsafeNewQPointerEvent(h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

func (this *QSinglePointEvent) Clone() *QSinglePointEvent {
	return UnsafeNewQSinglePointEvent(unsafe.Pointer(C.QSinglePointEvent_Clone(this.h)), nil, nil, nil)
}

func (this *QSinglePointEvent) Button() MouseButton {
	return (MouseButton)(C.QSinglePointEvent_Button(this.h))
}

func (this *QSinglePointEvent) Buttons() MouseButton {
	return (MouseButton)(C.QSinglePointEvent_Buttons(this.h))
}

func (this *QSinglePointEvent) Position() *QPointF {
	_ret := C.QSinglePointEvent_Position(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSinglePointEvent) ScenePosition() *QPointF {
	_ret := C.QSinglePointEvent_ScenePosition(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSinglePointEvent) GlobalPosition() *QPointF {
	_ret := C.QSinglePointEvent_GlobalPosition(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSinglePointEvent) IsBeginEvent() bool {
	return (bool)(C.QSinglePointEvent_IsBeginEvent(this.h))
}

func (this *QSinglePointEvent) IsUpdateEvent() bool {
	return (bool)(C.QSinglePointEvent_IsUpdateEvent(this.h))
}

func (this *QSinglePointEvent) IsEndEvent() bool {
	return (bool)(C.QSinglePointEvent_IsEndEvent(this.h))
}

func (this *QSinglePointEvent) ExclusivePointGrabber() *QObject {
	return UnsafeNewQObject(unsafe.Pointer(C.QSinglePointEvent_ExclusivePointGrabber(this.h)))
}

func (this *QSinglePointEvent) SetExclusivePointGrabber(exclusiveGrabber *QObject) {
	C.QSinglePointEvent_SetExclusivePointGrabber(this.h, exclusiveGrabber.cPointer())
}

// Delete this object from C++ memory.
func (this *QSinglePointEvent) Delete() {
	C.QSinglePointEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSinglePointEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QSinglePointEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QEnterEvent struct {
	h          *C.QEnterEvent
	isSubclass bool
	*QSinglePointEvent
}

func (this *QEnterEvent) cPointer() *C.QEnterEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QEnterEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQEnterEvent constructs the type using only CGO pointers.
func newQEnterEvent(h *C.QEnterEvent, h_QSinglePointEvent *C.QSinglePointEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QEnterEvent {
	if h == nil {
		return nil
	}
	return &QEnterEvent{h: h,
		QSinglePointEvent: newQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQEnterEvent constructs the type using only unsafe pointers.
func UnsafeNewQEnterEvent(h unsafe.Pointer, h_QSinglePointEvent unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QEnterEvent {
	if h == nil {
		return nil
	}

	return &QEnterEvent{h: (*C.QEnterEvent)(h),
		QSinglePointEvent: UnsafeNewQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// NewQEnterEvent constructs a new QEnterEvent object.
func NewQEnterEvent(localPos *QPointF, scenePos *QPointF, globalPos *QPointF) *QEnterEvent {
	var outptr_QEnterEvent *C.QEnterEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QEnterEvent_new(localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), &outptr_QEnterEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQEnterEvent(outptr_QEnterEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQEnterEvent2 constructs a new QEnterEvent object.
func NewQEnterEvent2(localPos *QPointF, scenePos *QPointF, globalPos *QPointF, device *QPointingDevice) *QEnterEvent {
	var outptr_QEnterEvent *C.QEnterEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QEnterEvent_new2(localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), device.cPointer(), &outptr_QEnterEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQEnterEvent(outptr_QEnterEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QEnterEvent) Clone() *QEnterEvent {
	return UnsafeNewQEnterEvent(unsafe.Pointer(C.QEnterEvent_Clone(this.h)), nil, nil, nil, nil)
}

func (this *QEnterEvent) Pos() *QPoint {
	_ret := C.QEnterEvent_Pos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) GlobalPos() *QPoint {
	_ret := C.QEnterEvent_GlobalPos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) X() int {
	return (int)(C.QEnterEvent_X(this.h))
}

func (this *QEnterEvent) Y() int {
	return (int)(C.QEnterEvent_Y(this.h))
}

func (this *QEnterEvent) GlobalX() int {
	return (int)(C.QEnterEvent_GlobalX(this.h))
}

func (this *QEnterEvent) GlobalY() int {
	return (int)(C.QEnterEvent_GlobalY(this.h))
}

func (this *QEnterEvent) LocalPos() *QPointF {
	_ret := C.QEnterEvent_LocalPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) WindowPos() *QPointF {
	_ret := C.QEnterEvent_WindowPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) ScreenPos() *QPointF {
	_ret := C.QEnterEvent_ScreenPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) callVirtualBase_Clone() *QEnterEvent {

	return UnsafeNewQEnterEvent(unsafe.Pointer(C.QEnterEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil, nil)
}
func (this *QEnterEvent) OnClone(slot func(super func() *QEnterEvent) *QEnterEvent) {
	C.QEnterEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_Clone
func miqt_exec_callback_QEnterEvent_Clone(self *C.QEnterEvent, cb C.intptr_t) *C.QEnterEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QEnterEvent) *QEnterEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QEnterEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QEnterEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QEnterEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QEnterEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_IsBeginEvent
func miqt_exec_callback_QEnterEvent_IsBeginEvent(self *C.QEnterEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QEnterEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QEnterEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QEnterEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QEnterEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_IsUpdateEvent
func miqt_exec_callback_QEnterEvent_IsUpdateEvent(self *C.QEnterEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QEnterEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QEnterEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QEnterEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QEnterEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_IsEndEvent
func miqt_exec_callback_QEnterEvent_IsEndEvent(self *C.QEnterEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QEnterEvent) Delete() {
	C.QEnterEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QEnterEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QEnterEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QMouseEvent struct {
	h          *C.QMouseEvent
	isSubclass bool
	*QSinglePointEvent
}

func (this *QMouseEvent) cPointer() *C.QMouseEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMouseEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMouseEvent constructs the type using only CGO pointers.
func newQMouseEvent(h *C.QMouseEvent, h_QSinglePointEvent *C.QSinglePointEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QMouseEvent {
	if h == nil {
		return nil
	}
	return &QMouseEvent{h: h,
		QSinglePointEvent: newQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQMouseEvent constructs the type using only unsafe pointers.
func UnsafeNewQMouseEvent(h unsafe.Pointer, h_QSinglePointEvent unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QMouseEvent {
	if h == nil {
		return nil
	}

	return &QMouseEvent{h: (*C.QMouseEvent)(h),
		QSinglePointEvent: UnsafeNewQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// NewQMouseEvent constructs a new QMouseEvent object.
func NewQMouseEvent(typeVal QEvent__Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new((C.int)(typeVal), localPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent2 constructs a new QMouseEvent object.
func NewQMouseEvent2(typeVal QEvent__Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new2((C.int)(typeVal), localPos.cPointer(), globalPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent3 constructs a new QMouseEvent object.
func NewQMouseEvent3(typeVal QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new3((C.int)(typeVal), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent4 constructs a new QMouseEvent object.
func NewQMouseEvent4(typeVal QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new4((C.int)(typeVal), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), (C.int)(source), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent5 constructs a new QMouseEvent object.
func NewQMouseEvent5(typeVal QEvent__Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new5((C.int)(typeVal), localPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), device.cPointer(), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent6 constructs a new QMouseEvent object.
func NewQMouseEvent6(typeVal QEvent__Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new6((C.int)(typeVal), localPos.cPointer(), globalPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), device.cPointer(), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent7 constructs a new QMouseEvent object.
func NewQMouseEvent7(typeVal QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new7((C.int)(typeVal), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), device.cPointer(), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent8 constructs a new QMouseEvent object.
func NewQMouseEvent8(typeVal QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource, device *QPointingDevice) *QMouseEvent {
	var outptr_QMouseEvent *C.QMouseEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMouseEvent_new8((C.int)(typeVal), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (C.int)(button), (C.int)(buttons), (C.int)(modifiers), (C.int)(source), device.cPointer(), &outptr_QMouseEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQMouseEvent(outptr_QMouseEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QMouseEvent) Clone() *QMouseEvent {
	return UnsafeNewQMouseEvent(unsafe.Pointer(C.QMouseEvent_Clone(this.h)), nil, nil, nil, nil)
}

func (this *QMouseEvent) Pos() *QPoint {
	_ret := C.QMouseEvent_Pos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) GlobalPos() *QPoint {
	_ret := C.QMouseEvent_GlobalPos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) X() int {
	return (int)(C.QMouseEvent_X(this.h))
}

func (this *QMouseEvent) Y() int {
	return (int)(C.QMouseEvent_Y(this.h))
}

func (this *QMouseEvent) GlobalX() int {
	return (int)(C.QMouseEvent_GlobalX(this.h))
}

func (this *QMouseEvent) GlobalY() int {
	return (int)(C.QMouseEvent_GlobalY(this.h))
}

func (this *QMouseEvent) LocalPos() *QPointF {
	_ret := C.QMouseEvent_LocalPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) WindowPos() *QPointF {
	_ret := C.QMouseEvent_WindowPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) ScreenPos() *QPointF {
	_ret := C.QMouseEvent_ScreenPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) Source() MouseEventSource {
	return (MouseEventSource)(C.QMouseEvent_Source(this.h))
}

func (this *QMouseEvent) Flags() MouseEventFlag {
	return (MouseEventFlag)(C.QMouseEvent_Flags(this.h))
}

func (this *QMouseEvent) callVirtualBase_Clone() *QMouseEvent {

	return UnsafeNewQMouseEvent(unsafe.Pointer(C.QMouseEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil, nil)
}
func (this *QMouseEvent) OnClone(slot func(super func() *QMouseEvent) *QMouseEvent) {
	C.QMouseEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_Clone
func miqt_exec_callback_QMouseEvent_Clone(self *C.QMouseEvent, cb C.intptr_t) *C.QMouseEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMouseEvent) *QMouseEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QMouseEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QMouseEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QMouseEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QMouseEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_IsBeginEvent
func miqt_exec_callback_QMouseEvent_IsBeginEvent(self *C.QMouseEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QMouseEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QMouseEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QMouseEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QMouseEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_IsUpdateEvent
func miqt_exec_callback_QMouseEvent_IsUpdateEvent(self *C.QMouseEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QMouseEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QMouseEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QMouseEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QMouseEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_IsEndEvent
func miqt_exec_callback_QMouseEvent_IsEndEvent(self *C.QMouseEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QMouseEvent) Delete() {
	C.QMouseEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMouseEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QMouseEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QHoverEvent struct {
	h          *C.QHoverEvent
	isSubclass bool
	*QSinglePointEvent
}

func (this *QHoverEvent) cPointer() *C.QHoverEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QHoverEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQHoverEvent constructs the type using only CGO pointers.
func newQHoverEvent(h *C.QHoverEvent, h_QSinglePointEvent *C.QSinglePointEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QHoverEvent {
	if h == nil {
		return nil
	}
	return &QHoverEvent{h: h,
		QSinglePointEvent: newQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQHoverEvent constructs the type using only unsafe pointers.
func UnsafeNewQHoverEvent(h unsafe.Pointer, h_QSinglePointEvent unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QHoverEvent {
	if h == nil {
		return nil
	}

	return &QHoverEvent{h: (*C.QHoverEvent)(h),
		QSinglePointEvent: UnsafeNewQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// NewQHoverEvent constructs a new QHoverEvent object.
func NewQHoverEvent(typeVal QEvent__Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF) *QHoverEvent {
	var outptr_QHoverEvent *C.QHoverEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHoverEvent_new((C.int)(typeVal), scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer(), &outptr_QHoverEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQHoverEvent(outptr_QHoverEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent2 constructs a new QHoverEvent object.
func NewQHoverEvent2(typeVal QEvent__Type, pos *QPointF, oldPos *QPointF) *QHoverEvent {
	var outptr_QHoverEvent *C.QHoverEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHoverEvent_new2((C.int)(typeVal), pos.cPointer(), oldPos.cPointer(), &outptr_QHoverEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQHoverEvent(outptr_QHoverEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent3 constructs a new QHoverEvent object.
func NewQHoverEvent3(typeVal QEvent__Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {
	var outptr_QHoverEvent *C.QHoverEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHoverEvent_new3((C.int)(typeVal), scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer(), (C.int)(modifiers), &outptr_QHoverEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQHoverEvent(outptr_QHoverEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent4 constructs a new QHoverEvent object.
func NewQHoverEvent4(typeVal QEvent__Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {
	var outptr_QHoverEvent *C.QHoverEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHoverEvent_new4((C.int)(typeVal), scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer(), (C.int)(modifiers), device.cPointer(), &outptr_QHoverEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQHoverEvent(outptr_QHoverEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent5 constructs a new QHoverEvent object.
func NewQHoverEvent5(typeVal QEvent__Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {
	var outptr_QHoverEvent *C.QHoverEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHoverEvent_new5((C.int)(typeVal), pos.cPointer(), oldPos.cPointer(), (C.int)(modifiers), &outptr_QHoverEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQHoverEvent(outptr_QHoverEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent6 constructs a new QHoverEvent object.
func NewQHoverEvent6(typeVal QEvent__Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {
	var outptr_QHoverEvent *C.QHoverEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHoverEvent_new6((C.int)(typeVal), pos.cPointer(), oldPos.cPointer(), (C.int)(modifiers), device.cPointer(), &outptr_QHoverEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQHoverEvent(outptr_QHoverEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QHoverEvent) Clone() *QHoverEvent {
	return UnsafeNewQHoverEvent(unsafe.Pointer(C.QHoverEvent_Clone(this.h)), nil, nil, nil, nil)
}

func (this *QHoverEvent) Pos() *QPoint {
	_ret := C.QHoverEvent_Pos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHoverEvent) PosF() *QPointF {
	_ret := C.QHoverEvent_PosF(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHoverEvent) IsUpdateEvent() bool {
	return (bool)(C.QHoverEvent_IsUpdateEvent(this.h))
}

func (this *QHoverEvent) OldPos() *QPoint {
	_ret := C.QHoverEvent_OldPos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHoverEvent) OldPosF() *QPointF {
	_ret := C.QHoverEvent_OldPosF(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHoverEvent) callVirtualBase_Clone() *QHoverEvent {

	return UnsafeNewQHoverEvent(unsafe.Pointer(C.QHoverEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil, nil)
}
func (this *QHoverEvent) OnClone(slot func(super func() *QHoverEvent) *QHoverEvent) {
	C.QHoverEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_Clone
func miqt_exec_callback_QHoverEvent_Clone(self *C.QHoverEvent, cb C.intptr_t) *C.QHoverEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QHoverEvent) *QHoverEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QHoverEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QHoverEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QHoverEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QHoverEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_IsUpdateEvent
func miqt_exec_callback_QHoverEvent_IsUpdateEvent(self *C.QHoverEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QHoverEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QHoverEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QHoverEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QHoverEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_IsBeginEvent
func miqt_exec_callback_QHoverEvent_IsBeginEvent(self *C.QHoverEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QHoverEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QHoverEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QHoverEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QHoverEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_IsEndEvent
func miqt_exec_callback_QHoverEvent_IsEndEvent(self *C.QHoverEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QHoverEvent) Delete() {
	C.QHoverEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QHoverEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QHoverEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QWheelEvent struct {
	h          *C.QWheelEvent
	isSubclass bool
	*QSinglePointEvent
}

func (this *QWheelEvent) cPointer() *C.QWheelEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWheelEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWheelEvent constructs the type using only CGO pointers.
func newQWheelEvent(h *C.QWheelEvent, h_QSinglePointEvent *C.QSinglePointEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QWheelEvent {
	if h == nil {
		return nil
	}
	return &QWheelEvent{h: h,
		QSinglePointEvent: newQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQWheelEvent constructs the type using only unsafe pointers.
func UnsafeNewQWheelEvent(h unsafe.Pointer, h_QSinglePointEvent unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QWheelEvent {
	if h == nil {
		return nil
	}

	return &QWheelEvent{h: (*C.QWheelEvent)(h),
		QSinglePointEvent: UnsafeNewQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// NewQWheelEvent constructs a new QWheelEvent object.
func NewQWheelEvent(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool) *QWheelEvent {
	var outptr_QWheelEvent *C.QWheelEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QWheelEvent_new(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (C.int)(buttons), (C.int)(modifiers), (C.int)(phase), (C.bool)(inverted), &outptr_QWheelEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQWheelEvent(outptr_QWheelEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQWheelEvent2 constructs a new QWheelEvent object.
func NewQWheelEvent2(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource) *QWheelEvent {
	var outptr_QWheelEvent *C.QWheelEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QWheelEvent_new2(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (C.int)(buttons), (C.int)(modifiers), (C.int)(phase), (C.bool)(inverted), (C.int)(source), &outptr_QWheelEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQWheelEvent(outptr_QWheelEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQWheelEvent3 constructs a new QWheelEvent object.
func NewQWheelEvent3(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource, device *QPointingDevice) *QWheelEvent {
	var outptr_QWheelEvent *C.QWheelEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QWheelEvent_new3(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (C.int)(buttons), (C.int)(modifiers), (C.int)(phase), (C.bool)(inverted), (C.int)(source), device.cPointer(), &outptr_QWheelEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQWheelEvent(outptr_QWheelEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QWheelEvent) Clone() *QWheelEvent {
	return UnsafeNewQWheelEvent(unsafe.Pointer(C.QWheelEvent_Clone(this.h)), nil, nil, nil, nil)
}

func (this *QWheelEvent) PixelDelta() *QPoint {
	_ret := C.QWheelEvent_PixelDelta(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWheelEvent) AngleDelta() *QPoint {
	_ret := C.QWheelEvent_AngleDelta(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWheelEvent) Phase() ScrollPhase {
	return (ScrollPhase)(C.QWheelEvent_Phase(this.h))
}

func (this *QWheelEvent) Inverted() bool {
	return (bool)(C.QWheelEvent_Inverted(this.h))
}

func (this *QWheelEvent) IsInverted() bool {
	return (bool)(C.QWheelEvent_IsInverted(this.h))
}

func (this *QWheelEvent) HasPixelDelta() bool {
	return (bool)(C.QWheelEvent_HasPixelDelta(this.h))
}

func (this *QWheelEvent) IsBeginEvent() bool {
	return (bool)(C.QWheelEvent_IsBeginEvent(this.h))
}

func (this *QWheelEvent) IsUpdateEvent() bool {
	return (bool)(C.QWheelEvent_IsUpdateEvent(this.h))
}

func (this *QWheelEvent) IsEndEvent() bool {
	return (bool)(C.QWheelEvent_IsEndEvent(this.h))
}

func (this *QWheelEvent) Source() MouseEventSource {
	return (MouseEventSource)(C.QWheelEvent_Source(this.h))
}

func (this *QWheelEvent) callVirtualBase_Clone() *QWheelEvent {

	return UnsafeNewQWheelEvent(unsafe.Pointer(C.QWheelEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil, nil)
}
func (this *QWheelEvent) OnClone(slot func(super func() *QWheelEvent) *QWheelEvent) {
	C.QWheelEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_Clone
func miqt_exec_callback_QWheelEvent_Clone(self *C.QWheelEvent, cb C.intptr_t) *C.QWheelEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWheelEvent) *QWheelEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QWheelEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QWheelEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QWheelEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QWheelEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_IsBeginEvent
func miqt_exec_callback_QWheelEvent_IsBeginEvent(self *C.QWheelEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QWheelEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QWheelEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QWheelEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QWheelEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_IsUpdateEvent
func miqt_exec_callback_QWheelEvent_IsUpdateEvent(self *C.QWheelEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QWheelEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QWheelEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QWheelEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QWheelEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_IsEndEvent
func miqt_exec_callback_QWheelEvent_IsEndEvent(self *C.QWheelEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QWheelEvent) Delete() {
	C.QWheelEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWheelEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QWheelEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QTabletEvent struct {
	h          *C.QTabletEvent
	isSubclass bool
	*QSinglePointEvent
}

func (this *QTabletEvent) cPointer() *C.QTabletEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QTabletEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQTabletEvent constructs the type using only CGO pointers.
func newQTabletEvent(h *C.QTabletEvent, h_QSinglePointEvent *C.QSinglePointEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QTabletEvent {
	if h == nil {
		return nil
	}
	return &QTabletEvent{h: h,
		QSinglePointEvent: newQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQTabletEvent constructs the type using only unsafe pointers.
func UnsafeNewQTabletEvent(h unsafe.Pointer, h_QSinglePointEvent unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QTabletEvent {
	if h == nil {
		return nil
	}

	return &QTabletEvent{h: (*C.QTabletEvent)(h),
		QSinglePointEvent: UnsafeNewQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// NewQTabletEvent constructs a new QTabletEvent object.
func NewQTabletEvent(t QEvent__Type, device *QPointingDevice, pos *QPointF, globalPos *QPointF, pressure float64, xTilt float32, yTilt float32, tangentialPressure float32, rotation float64, z float32, keyState KeyboardModifier, button MouseButton, buttons MouseButton) *QTabletEvent {
	var outptr_QTabletEvent *C.QTabletEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QTabletEvent_new((C.int)(t), device.cPointer(), pos.cPointer(), globalPos.cPointer(), (C.double)(pressure), (C.float)(xTilt), (C.float)(yTilt), (C.float)(tangentialPressure), (C.double)(rotation), (C.float)(z), (C.int)(keyState), (C.int)(button), (C.int)(buttons), &outptr_QTabletEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQTabletEvent(outptr_QTabletEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QTabletEvent) Clone() *QTabletEvent {
	return UnsafeNewQTabletEvent(unsafe.Pointer(C.QTabletEvent_Clone(this.h)), nil, nil, nil, nil)
}

func (this *QTabletEvent) Pos() *QPoint {
	_ret := C.QTabletEvent_Pos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) GlobalPos() *QPoint {
	_ret := C.QTabletEvent_GlobalPos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) PosF() *QPointF {
	_ret := C.QTabletEvent_PosF(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) GlobalPosF() *QPointF {
	_ret := C.QTabletEvent_GlobalPosF(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) X() int {
	return (int)(C.QTabletEvent_X(this.h))
}

func (this *QTabletEvent) Y() int {
	return (int)(C.QTabletEvent_Y(this.h))
}

func (this *QTabletEvent) GlobalX() int {
	return (int)(C.QTabletEvent_GlobalX(this.h))
}

func (this *QTabletEvent) GlobalY() int {
	return (int)(C.QTabletEvent_GlobalY(this.h))
}

func (this *QTabletEvent) HiResGlobalX() float64 {
	return (float64)(C.QTabletEvent_HiResGlobalX(this.h))
}

func (this *QTabletEvent) HiResGlobalY() float64 {
	return (float64)(C.QTabletEvent_HiResGlobalY(this.h))
}

func (this *QTabletEvent) UniqueId() int64 {
	return (int64)(C.QTabletEvent_UniqueId(this.h))
}

func (this *QTabletEvent) Pressure() float64 {
	return (float64)(C.QTabletEvent_Pressure(this.h))
}

func (this *QTabletEvent) Rotation() float64 {
	return (float64)(C.QTabletEvent_Rotation(this.h))
}

func (this *QTabletEvent) Z() float64 {
	return (float64)(C.QTabletEvent_Z(this.h))
}

func (this *QTabletEvent) TangentialPressure() float64 {
	return (float64)(C.QTabletEvent_TangentialPressure(this.h))
}

func (this *QTabletEvent) XTilt() float64 {
	return (float64)(C.QTabletEvent_XTilt(this.h))
}

func (this *QTabletEvent) YTilt() float64 {
	return (float64)(C.QTabletEvent_YTilt(this.h))
}

func (this *QTabletEvent) callVirtualBase_Clone() *QTabletEvent {

	return UnsafeNewQTabletEvent(unsafe.Pointer(C.QTabletEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil, nil)
}
func (this *QTabletEvent) OnClone(slot func(super func() *QTabletEvent) *QTabletEvent) {
	C.QTabletEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_Clone
func miqt_exec_callback_QTabletEvent_Clone(self *C.QTabletEvent, cb C.intptr_t) *C.QTabletEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QTabletEvent) *QTabletEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QTabletEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QTabletEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QTabletEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QTabletEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_IsBeginEvent
func miqt_exec_callback_QTabletEvent_IsBeginEvent(self *C.QTabletEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QTabletEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QTabletEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QTabletEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QTabletEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_IsUpdateEvent
func miqt_exec_callback_QTabletEvent_IsUpdateEvent(self *C.QTabletEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QTabletEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QTabletEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QTabletEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QTabletEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_IsEndEvent
func miqt_exec_callback_QTabletEvent_IsEndEvent(self *C.QTabletEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QTabletEvent) Delete() {
	C.QTabletEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QTabletEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QTabletEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QNativeGestureEvent struct {
	h          *C.QNativeGestureEvent
	isSubclass bool
	*QSinglePointEvent
}

func (this *QNativeGestureEvent) cPointer() *C.QNativeGestureEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QNativeGestureEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQNativeGestureEvent constructs the type using only CGO pointers.
func newQNativeGestureEvent(h *C.QNativeGestureEvent, h_QSinglePointEvent *C.QSinglePointEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QNativeGestureEvent {
	if h == nil {
		return nil
	}
	return &QNativeGestureEvent{h: h,
		QSinglePointEvent: newQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQNativeGestureEvent constructs the type using only unsafe pointers.
func UnsafeNewQNativeGestureEvent(h unsafe.Pointer, h_QSinglePointEvent unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QNativeGestureEvent {
	if h == nil {
		return nil
	}

	return &QNativeGestureEvent{h: (*C.QNativeGestureEvent)(h),
		QSinglePointEvent: UnsafeNewQSinglePointEvent(h_QSinglePointEvent, h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// NewQNativeGestureEvent constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent(typeVal NativeGestureType, dev *QPointingDevice, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, sequenceId uint64, intArgument uint64) *QNativeGestureEvent {
	var outptr_QNativeGestureEvent *C.QNativeGestureEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QNativeGestureEvent_new((C.int)(typeVal), dev.cPointer(), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (C.double)(value), (C.ulonglong)(sequenceId), (C.ulonglong)(intArgument), &outptr_QNativeGestureEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQNativeGestureEvent(outptr_QNativeGestureEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQNativeGestureEvent2 constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent2(typeVal NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF) *QNativeGestureEvent {
	var outptr_QNativeGestureEvent *C.QNativeGestureEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QNativeGestureEvent_new2((C.int)(typeVal), dev.cPointer(), (C.int)(fingerCount), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (C.double)(value), delta.cPointer(), &outptr_QNativeGestureEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQNativeGestureEvent(outptr_QNativeGestureEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQNativeGestureEvent3 constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent3(typeVal NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF, sequenceId uint64) *QNativeGestureEvent {
	var outptr_QNativeGestureEvent *C.QNativeGestureEvent = nil
	var outptr_QSinglePointEvent *C.QSinglePointEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QNativeGestureEvent_new3((C.int)(typeVal), dev.cPointer(), (C.int)(fingerCount), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (C.double)(value), delta.cPointer(), (C.ulonglong)(sequenceId), &outptr_QNativeGestureEvent, &outptr_QSinglePointEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQNativeGestureEvent(outptr_QNativeGestureEvent, outptr_QSinglePointEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QNativeGestureEvent) Clone() *QNativeGestureEvent {
	return UnsafeNewQNativeGestureEvent(unsafe.Pointer(C.QNativeGestureEvent_Clone(this.h)), nil, nil, nil, nil)
}

func (this *QNativeGestureEvent) GestureType() NativeGestureType {
	return (NativeGestureType)(C.QNativeGestureEvent_GestureType(this.h))
}

func (this *QNativeGestureEvent) FingerCount() int {
	return (int)(C.QNativeGestureEvent_FingerCount(this.h))
}

func (this *QNativeGestureEvent) Value() float64 {
	return (float64)(C.QNativeGestureEvent_Value(this.h))
}

func (this *QNativeGestureEvent) Delta() *QPointF {
	_ret := C.QNativeGestureEvent_Delta(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) Pos() *QPoint {
	_ret := C.QNativeGestureEvent_Pos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) GlobalPos() *QPoint {
	_ret := C.QNativeGestureEvent_GlobalPos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) LocalPos() *QPointF {
	_ret := C.QNativeGestureEvent_LocalPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) WindowPos() *QPointF {
	_ret := C.QNativeGestureEvent_WindowPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) ScreenPos() *QPointF {
	_ret := C.QNativeGestureEvent_ScreenPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) callVirtualBase_Clone() *QNativeGestureEvent {

	return UnsafeNewQNativeGestureEvent(unsafe.Pointer(C.QNativeGestureEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil, nil)
}
func (this *QNativeGestureEvent) OnClone(slot func(super func() *QNativeGestureEvent) *QNativeGestureEvent) {
	C.QNativeGestureEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_Clone
func miqt_exec_callback_QNativeGestureEvent_Clone(self *C.QNativeGestureEvent, cb C.intptr_t) *C.QNativeGestureEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QNativeGestureEvent) *QNativeGestureEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QNativeGestureEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QNativeGestureEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QNativeGestureEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QNativeGestureEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_IsBeginEvent
func miqt_exec_callback_QNativeGestureEvent_IsBeginEvent(self *C.QNativeGestureEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QNativeGestureEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QNativeGestureEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QNativeGestureEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QNativeGestureEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_IsUpdateEvent
func miqt_exec_callback_QNativeGestureEvent_IsUpdateEvent(self *C.QNativeGestureEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QNativeGestureEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QNativeGestureEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QNativeGestureEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QNativeGestureEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_IsEndEvent
func miqt_exec_callback_QNativeGestureEvent_IsEndEvent(self *C.QNativeGestureEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QNativeGestureEvent) Delete() {
	C.QNativeGestureEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QNativeGestureEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QNativeGestureEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QKeyEvent struct {
	h          *C.QKeyEvent
	isSubclass bool
	*QInputEvent
}

func (this *QKeyEvent) cPointer() *C.QKeyEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QKeyEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQKeyEvent constructs the type using only CGO pointers.
func newQKeyEvent(h *C.QKeyEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QKeyEvent {
	if h == nil {
		return nil
	}
	return &QKeyEvent{h: h,
		QInputEvent: newQInputEvent(h_QInputEvent, h_QEvent)}
}

// UnsafeNewQKeyEvent constructs the type using only unsafe pointers.
func UnsafeNewQKeyEvent(h unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QKeyEvent {
	if h == nil {
		return nil
	}

	return &QKeyEvent{h: (*C.QKeyEvent)(h),
		QInputEvent: UnsafeNewQInputEvent(h_QInputEvent, h_QEvent)}
}

// NewQKeyEvent constructs a new QKeyEvent object.
func NewQKeyEvent(typeVal QEvent__Type, key int, modifiers KeyboardModifier) *QKeyEvent {
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent2 constructs a new QKeyEvent object.
func NewQKeyEvent2(typeVal QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint) *QKeyEvent {
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new2((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), (C.uint)(nativeScanCode), (C.uint)(nativeVirtualKey), (C.uint)(nativeModifiers), &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent3 constructs a new QKeyEvent object.
func NewQKeyEvent3(typeVal QEvent__Type, key int, modifiers KeyboardModifier, text string) *QKeyEvent {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new3((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), text_ms, &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent4 constructs a new QKeyEvent object.
func NewQKeyEvent4(typeVal QEvent__Type, key int, modifiers KeyboardModifier, text string, autorep bool) *QKeyEvent {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new4((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), text_ms, (C.bool)(autorep), &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent5 constructs a new QKeyEvent object.
func NewQKeyEvent5(typeVal QEvent__Type, key int, modifiers KeyboardModifier, text string, autorep bool, count uint16) *QKeyEvent {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new5((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), text_ms, (C.bool)(autorep), (C.uint16_t)(count), &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent6 constructs a new QKeyEvent object.
func NewQKeyEvent6(typeVal QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string) *QKeyEvent {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new6((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), (C.uint)(nativeScanCode), (C.uint)(nativeVirtualKey), (C.uint)(nativeModifiers), text_ms, &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent7 constructs a new QKeyEvent object.
func NewQKeyEvent7(typeVal QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool) *QKeyEvent {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new7((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), (C.uint)(nativeScanCode), (C.uint)(nativeVirtualKey), (C.uint)(nativeModifiers), text_ms, (C.bool)(autorep), &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent8 constructs a new QKeyEvent object.
func NewQKeyEvent8(typeVal QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16) *QKeyEvent {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new8((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), (C.uint)(nativeScanCode), (C.uint)(nativeVirtualKey), (C.uint)(nativeModifiers), text_ms, (C.bool)(autorep), (C.uint16_t)(count), &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent9 constructs a new QKeyEvent object.
func NewQKeyEvent9(typeVal QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16, device *QInputDevice) *QKeyEvent {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	var outptr_QKeyEvent *C.QKeyEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QKeyEvent_new9((C.int)(typeVal), (C.int)(key), (C.int)(modifiers), (C.uint)(nativeScanCode), (C.uint)(nativeVirtualKey), (C.uint)(nativeModifiers), text_ms, (C.bool)(autorep), (C.uint16_t)(count), device.cPointer(), &outptr_QKeyEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQKeyEvent(outptr_QKeyEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QKeyEvent) Clone() *QKeyEvent {
	return UnsafeNewQKeyEvent(unsafe.Pointer(C.QKeyEvent_Clone(this.h)), nil, nil)
}

func (this *QKeyEvent) Key() int {
	return (int)(C.QKeyEvent_Key(this.h))
}

func (this *QKeyEvent) Matches(key QKeySequence__StandardKey) bool {
	return (bool)(C.QKeyEvent_Matches(this.h, (C.int)(key)))
}

func (this *QKeyEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(C.QKeyEvent_Modifiers(this.h))
}

func (this *QKeyEvent) KeyCombination() *QKeyCombination {
	_ret := C.QKeyEvent_KeyCombination(this.h)
	_goptr := newQKeyCombination(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QKeyEvent) Text() string {
	var _ms C.struct_miqt_string = C.QKeyEvent_Text(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QKeyEvent) IsAutoRepeat() bool {
	return (bool)(C.QKeyEvent_IsAutoRepeat(this.h))
}

func (this *QKeyEvent) Count() int {
	return (int)(C.QKeyEvent_Count(this.h))
}

func (this *QKeyEvent) NativeScanCode() uint {
	return (uint)(C.QKeyEvent_NativeScanCode(this.h))
}

func (this *QKeyEvent) NativeVirtualKey() uint {
	return (uint)(C.QKeyEvent_NativeVirtualKey(this.h))
}

func (this *QKeyEvent) NativeModifiers() uint {
	return (uint)(C.QKeyEvent_NativeModifiers(this.h))
}

func (this *QKeyEvent) callVirtualBase_Clone() *QKeyEvent {

	return UnsafeNewQKeyEvent(unsafe.Pointer(C.QKeyEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil)
}
func (this *QKeyEvent) OnClone(slot func(super func() *QKeyEvent) *QKeyEvent) {
	C.QKeyEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeyEvent_Clone
func miqt_exec_callback_QKeyEvent_Clone(self *C.QKeyEvent, cb C.intptr_t) *C.QKeyEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QKeyEvent) *QKeyEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeyEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QKeyEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	C.QKeyEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (C.ulonglong)(timestamp))

}
func (this *QKeyEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	C.QKeyEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeyEvent_SetTimestamp
func miqt_exec_callback_QKeyEvent_SetTimestamp(self *C.QKeyEvent, cb C.intptr_t, timestamp C.ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QKeyEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

// Delete this object from C++ memory.
func (this *QKeyEvent) Delete() {
	C.QKeyEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QKeyEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QKeyEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QFocusEvent struct {
	h          *C.QFocusEvent
	isSubclass bool
	*QEvent
}

func (this *QFocusEvent) cPointer() *C.QFocusEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QFocusEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQFocusEvent constructs the type using only CGO pointers.
func newQFocusEvent(h *C.QFocusEvent, h_QEvent *C.QEvent) *QFocusEvent {
	if h == nil {
		return nil
	}
	return &QFocusEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQFocusEvent constructs the type using only unsafe pointers.
func UnsafeNewQFocusEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QFocusEvent {
	if h == nil {
		return nil
	}

	return &QFocusEvent{h: (*C.QFocusEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQFocusEvent constructs a new QFocusEvent object.
func NewQFocusEvent(typeVal QEvent__Type) *QFocusEvent {
	var outptr_QFocusEvent *C.QFocusEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QFocusEvent_new((C.int)(typeVal), &outptr_QFocusEvent, &outptr_QEvent)
	ret := newQFocusEvent(outptr_QFocusEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQFocusEvent2 constructs a new QFocusEvent object.
func NewQFocusEvent2(typeVal QEvent__Type, reason FocusReason) *QFocusEvent {
	var outptr_QFocusEvent *C.QFocusEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QFocusEvent_new2((C.int)(typeVal), (C.int)(reason), &outptr_QFocusEvent, &outptr_QEvent)
	ret := newQFocusEvent(outptr_QFocusEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QFocusEvent) Clone() *QFocusEvent {
	return UnsafeNewQFocusEvent(unsafe.Pointer(C.QFocusEvent_Clone(this.h)), nil)
}

func (this *QFocusEvent) GotFocus() bool {
	return (bool)(C.QFocusEvent_GotFocus(this.h))
}

func (this *QFocusEvent) LostFocus() bool {
	return (bool)(C.QFocusEvent_LostFocus(this.h))
}

func (this *QFocusEvent) Reason() FocusReason {
	return (FocusReason)(C.QFocusEvent_Reason(this.h))
}

func (this *QFocusEvent) callVirtualBase_Clone() *QFocusEvent {

	return UnsafeNewQFocusEvent(unsafe.Pointer(C.QFocusEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QFocusEvent) OnClone(slot func(super func() *QFocusEvent) *QFocusEvent) {
	C.QFocusEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFocusEvent_Clone
func miqt_exec_callback_QFocusEvent_Clone(self *C.QFocusEvent, cb C.intptr_t) *C.QFocusEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QFocusEvent) *QFocusEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFocusEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QFocusEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QFocusEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QFocusEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QFocusEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFocusEvent_SetAccepted
func miqt_exec_callback_QFocusEvent_SetAccepted(self *C.QFocusEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QFocusEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QFocusEvent) Delete() {
	C.QFocusEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QFocusEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QFocusEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QPaintEvent struct {
	h          *C.QPaintEvent
	isSubclass bool
	*QEvent
}

func (this *QPaintEvent) cPointer() *C.QPaintEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPaintEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPaintEvent constructs the type using only CGO pointers.
func newQPaintEvent(h *C.QPaintEvent, h_QEvent *C.QEvent) *QPaintEvent {
	if h == nil {
		return nil
	}
	return &QPaintEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQPaintEvent constructs the type using only unsafe pointers.
func UnsafeNewQPaintEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QPaintEvent {
	if h == nil {
		return nil
	}

	return &QPaintEvent{h: (*C.QPaintEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQPaintEvent constructs a new QPaintEvent object.
func NewQPaintEvent(paintRegion *QRegion) *QPaintEvent {
	var outptr_QPaintEvent *C.QPaintEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QPaintEvent_new(paintRegion.cPointer(), &outptr_QPaintEvent, &outptr_QEvent)
	ret := newQPaintEvent(outptr_QPaintEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQPaintEvent2 constructs a new QPaintEvent object.
func NewQPaintEvent2(paintRect *QRect) *QPaintEvent {
	var outptr_QPaintEvent *C.QPaintEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QPaintEvent_new2(paintRect.cPointer(), &outptr_QPaintEvent, &outptr_QEvent)
	ret := newQPaintEvent(outptr_QPaintEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QPaintEvent) Clone() *QPaintEvent {
	return UnsafeNewQPaintEvent(unsafe.Pointer(C.QPaintEvent_Clone(this.h)), nil)
}

func (this *QPaintEvent) Rect() *QRect {
	return UnsafeNewQRect(unsafe.Pointer(C.QPaintEvent_Rect(this.h)))
}

func (this *QPaintEvent) Region() *QRegion {
	return UnsafeNewQRegion(unsafe.Pointer(C.QPaintEvent_Region(this.h)))
}

func (this *QPaintEvent) callVirtualBase_Clone() *QPaintEvent {

	return UnsafeNewQPaintEvent(unsafe.Pointer(C.QPaintEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QPaintEvent) OnClone(slot func(super func() *QPaintEvent) *QPaintEvent) {
	C.QPaintEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEvent_Clone
func miqt_exec_callback_QPaintEvent_Clone(self *C.QPaintEvent, cb C.intptr_t) *C.QPaintEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEvent) *QPaintEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPaintEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QPaintEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QPaintEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QPaintEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QPaintEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEvent_SetAccepted
func miqt_exec_callback_QPaintEvent_SetAccepted(self *C.QPaintEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QPaintEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QPaintEvent) Delete() {
	C.QPaintEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPaintEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QPaintEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QMoveEvent struct {
	h          *C.QMoveEvent
	isSubclass bool
	*QEvent
}

func (this *QMoveEvent) cPointer() *C.QMoveEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMoveEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMoveEvent constructs the type using only CGO pointers.
func newQMoveEvent(h *C.QMoveEvent, h_QEvent *C.QEvent) *QMoveEvent {
	if h == nil {
		return nil
	}
	return &QMoveEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQMoveEvent constructs the type using only unsafe pointers.
func UnsafeNewQMoveEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QMoveEvent {
	if h == nil {
		return nil
	}

	return &QMoveEvent{h: (*C.QMoveEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQMoveEvent constructs a new QMoveEvent object.
func NewQMoveEvent(pos *QPoint, oldPos *QPoint) *QMoveEvent {
	var outptr_QMoveEvent *C.QMoveEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QMoveEvent_new(pos.cPointer(), oldPos.cPointer(), &outptr_QMoveEvent, &outptr_QEvent)
	ret := newQMoveEvent(outptr_QMoveEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QMoveEvent) Clone() *QMoveEvent {
	return UnsafeNewQMoveEvent(unsafe.Pointer(C.QMoveEvent_Clone(this.h)), nil)
}

func (this *QMoveEvent) Pos() *QPoint {
	return UnsafeNewQPoint(unsafe.Pointer(C.QMoveEvent_Pos(this.h)))
}

func (this *QMoveEvent) OldPos() *QPoint {
	return UnsafeNewQPoint(unsafe.Pointer(C.QMoveEvent_OldPos(this.h)))
}

func (this *QMoveEvent) callVirtualBase_Clone() *QMoveEvent {

	return UnsafeNewQMoveEvent(unsafe.Pointer(C.QMoveEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QMoveEvent) OnClone(slot func(super func() *QMoveEvent) *QMoveEvent) {
	C.QMoveEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMoveEvent_Clone
func miqt_exec_callback_QMoveEvent_Clone(self *C.QMoveEvent, cb C.intptr_t) *C.QMoveEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMoveEvent) *QMoveEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMoveEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QMoveEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QMoveEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QMoveEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QMoveEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMoveEvent_SetAccepted
func miqt_exec_callback_QMoveEvent_SetAccepted(self *C.QMoveEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QMoveEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QMoveEvent) Delete() {
	C.QMoveEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMoveEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QMoveEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QExposeEvent struct {
	h          *C.QExposeEvent
	isSubclass bool
	*QEvent
}

func (this *QExposeEvent) cPointer() *C.QExposeEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QExposeEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQExposeEvent constructs the type using only CGO pointers.
func newQExposeEvent(h *C.QExposeEvent, h_QEvent *C.QEvent) *QExposeEvent {
	if h == nil {
		return nil
	}
	return &QExposeEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQExposeEvent constructs the type using only unsafe pointers.
func UnsafeNewQExposeEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QExposeEvent {
	if h == nil {
		return nil
	}

	return &QExposeEvent{h: (*C.QExposeEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQExposeEvent constructs a new QExposeEvent object.
func NewQExposeEvent(m_region *QRegion) *QExposeEvent {
	var outptr_QExposeEvent *C.QExposeEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QExposeEvent_new(m_region.cPointer(), &outptr_QExposeEvent, &outptr_QEvent)
	ret := newQExposeEvent(outptr_QExposeEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QExposeEvent) Clone() *QExposeEvent {
	return UnsafeNewQExposeEvent(unsafe.Pointer(C.QExposeEvent_Clone(this.h)), nil)
}

func (this *QExposeEvent) Region() *QRegion {
	return UnsafeNewQRegion(unsafe.Pointer(C.QExposeEvent_Region(this.h)))
}

func (this *QExposeEvent) callVirtualBase_Clone() *QExposeEvent {

	return UnsafeNewQExposeEvent(unsafe.Pointer(C.QExposeEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QExposeEvent) OnClone(slot func(super func() *QExposeEvent) *QExposeEvent) {
	C.QExposeEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QExposeEvent_Clone
func miqt_exec_callback_QExposeEvent_Clone(self *C.QExposeEvent, cb C.intptr_t) *C.QExposeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QExposeEvent) *QExposeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QExposeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QExposeEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QExposeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QExposeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QExposeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QExposeEvent_SetAccepted
func miqt_exec_callback_QExposeEvent_SetAccepted(self *C.QExposeEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QExposeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QExposeEvent) Delete() {
	C.QExposeEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QExposeEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QExposeEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QPlatformSurfaceEvent struct {
	h          *C.QPlatformSurfaceEvent
	isSubclass bool
	*QEvent
}

func (this *QPlatformSurfaceEvent) cPointer() *C.QPlatformSurfaceEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPlatformSurfaceEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPlatformSurfaceEvent constructs the type using only CGO pointers.
func newQPlatformSurfaceEvent(h *C.QPlatformSurfaceEvent, h_QEvent *C.QEvent) *QPlatformSurfaceEvent {
	if h == nil {
		return nil
	}
	return &QPlatformSurfaceEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQPlatformSurfaceEvent constructs the type using only unsafe pointers.
func UnsafeNewQPlatformSurfaceEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QPlatformSurfaceEvent {
	if h == nil {
		return nil
	}

	return &QPlatformSurfaceEvent{h: (*C.QPlatformSurfaceEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQPlatformSurfaceEvent constructs a new QPlatformSurfaceEvent object.
func NewQPlatformSurfaceEvent(surfaceEventType QPlatformSurfaceEvent__SurfaceEventType) *QPlatformSurfaceEvent {
	var outptr_QPlatformSurfaceEvent *C.QPlatformSurfaceEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QPlatformSurfaceEvent_new((C.int)(surfaceEventType), &outptr_QPlatformSurfaceEvent, &outptr_QEvent)
	ret := newQPlatformSurfaceEvent(outptr_QPlatformSurfaceEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QPlatformSurfaceEvent) Clone() *QPlatformSurfaceEvent {
	return UnsafeNewQPlatformSurfaceEvent(unsafe.Pointer(C.QPlatformSurfaceEvent_Clone(this.h)), nil)
}

func (this *QPlatformSurfaceEvent) SurfaceEventType() QPlatformSurfaceEvent__SurfaceEventType {
	return (QPlatformSurfaceEvent__SurfaceEventType)(C.QPlatformSurfaceEvent_SurfaceEventType(this.h))
}

func (this *QPlatformSurfaceEvent) callVirtualBase_Clone() *QPlatformSurfaceEvent {

	return UnsafeNewQPlatformSurfaceEvent(unsafe.Pointer(C.QPlatformSurfaceEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QPlatformSurfaceEvent) OnClone(slot func(super func() *QPlatformSurfaceEvent) *QPlatformSurfaceEvent) {
	C.QPlatformSurfaceEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlatformSurfaceEvent_Clone
func miqt_exec_callback_QPlatformSurfaceEvent_Clone(self *C.QPlatformSurfaceEvent, cb C.intptr_t) *C.QPlatformSurfaceEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPlatformSurfaceEvent) *QPlatformSurfaceEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPlatformSurfaceEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QPlatformSurfaceEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QPlatformSurfaceEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QPlatformSurfaceEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QPlatformSurfaceEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlatformSurfaceEvent_SetAccepted
func miqt_exec_callback_QPlatformSurfaceEvent_SetAccepted(self *C.QPlatformSurfaceEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QPlatformSurfaceEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QPlatformSurfaceEvent) Delete() {
	C.QPlatformSurfaceEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPlatformSurfaceEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QPlatformSurfaceEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QResizeEvent struct {
	h          *C.QResizeEvent
	isSubclass bool
	*QEvent
}

func (this *QResizeEvent) cPointer() *C.QResizeEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QResizeEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQResizeEvent constructs the type using only CGO pointers.
func newQResizeEvent(h *C.QResizeEvent, h_QEvent *C.QEvent) *QResizeEvent {
	if h == nil {
		return nil
	}
	return &QResizeEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQResizeEvent constructs the type using only unsafe pointers.
func UnsafeNewQResizeEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QResizeEvent {
	if h == nil {
		return nil
	}

	return &QResizeEvent{h: (*C.QResizeEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQResizeEvent constructs a new QResizeEvent object.
func NewQResizeEvent(size *QSize, oldSize *QSize) *QResizeEvent {
	var outptr_QResizeEvent *C.QResizeEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QResizeEvent_new(size.cPointer(), oldSize.cPointer(), &outptr_QResizeEvent, &outptr_QEvent)
	ret := newQResizeEvent(outptr_QResizeEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QResizeEvent) Clone() *QResizeEvent {
	return UnsafeNewQResizeEvent(unsafe.Pointer(C.QResizeEvent_Clone(this.h)), nil)
}

func (this *QResizeEvent) Size() *QSize {
	return UnsafeNewQSize(unsafe.Pointer(C.QResizeEvent_Size(this.h)))
}

func (this *QResizeEvent) OldSize() *QSize {
	return UnsafeNewQSize(unsafe.Pointer(C.QResizeEvent_OldSize(this.h)))
}

func (this *QResizeEvent) callVirtualBase_Clone() *QResizeEvent {

	return UnsafeNewQResizeEvent(unsafe.Pointer(C.QResizeEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QResizeEvent) OnClone(slot func(super func() *QResizeEvent) *QResizeEvent) {
	C.QResizeEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QResizeEvent_Clone
func miqt_exec_callback_QResizeEvent_Clone(self *C.QResizeEvent, cb C.intptr_t) *C.QResizeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QResizeEvent) *QResizeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QResizeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QResizeEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QResizeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QResizeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QResizeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QResizeEvent_SetAccepted
func miqt_exec_callback_QResizeEvent_SetAccepted(self *C.QResizeEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QResizeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QResizeEvent) Delete() {
	C.QResizeEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QResizeEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QResizeEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QCloseEvent struct {
	h          *C.QCloseEvent
	isSubclass bool
	*QEvent
}

func (this *QCloseEvent) cPointer() *C.QCloseEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QCloseEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQCloseEvent constructs the type using only CGO pointers.
func newQCloseEvent(h *C.QCloseEvent, h_QEvent *C.QEvent) *QCloseEvent {
	if h == nil {
		return nil
	}
	return &QCloseEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQCloseEvent constructs the type using only unsafe pointers.
func UnsafeNewQCloseEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QCloseEvent {
	if h == nil {
		return nil
	}

	return &QCloseEvent{h: (*C.QCloseEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQCloseEvent constructs a new QCloseEvent object.
func NewQCloseEvent() *QCloseEvent {
	var outptr_QCloseEvent *C.QCloseEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QCloseEvent_new(&outptr_QCloseEvent, &outptr_QEvent)
	ret := newQCloseEvent(outptr_QCloseEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QCloseEvent) Clone() *QCloseEvent {
	return UnsafeNewQCloseEvent(unsafe.Pointer(C.QCloseEvent_Clone(this.h)), nil)
}

func (this *QCloseEvent) callVirtualBase_Clone() *QCloseEvent {

	return UnsafeNewQCloseEvent(unsafe.Pointer(C.QCloseEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QCloseEvent) OnClone(slot func(super func() *QCloseEvent) *QCloseEvent) {
	C.QCloseEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCloseEvent_Clone
func miqt_exec_callback_QCloseEvent_Clone(self *C.QCloseEvent, cb C.intptr_t) *C.QCloseEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QCloseEvent) *QCloseEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCloseEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QCloseEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QCloseEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QCloseEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QCloseEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCloseEvent_SetAccepted
func miqt_exec_callback_QCloseEvent_SetAccepted(self *C.QCloseEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QCloseEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QCloseEvent) Delete() {
	C.QCloseEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCloseEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QCloseEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QIconDragEvent struct {
	h          *C.QIconDragEvent
	isSubclass bool
	*QEvent
}

func (this *QIconDragEvent) cPointer() *C.QIconDragEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QIconDragEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQIconDragEvent constructs the type using only CGO pointers.
func newQIconDragEvent(h *C.QIconDragEvent, h_QEvent *C.QEvent) *QIconDragEvent {
	if h == nil {
		return nil
	}
	return &QIconDragEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQIconDragEvent constructs the type using only unsafe pointers.
func UnsafeNewQIconDragEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QIconDragEvent {
	if h == nil {
		return nil
	}

	return &QIconDragEvent{h: (*C.QIconDragEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQIconDragEvent constructs a new QIconDragEvent object.
func NewQIconDragEvent() *QIconDragEvent {
	var outptr_QIconDragEvent *C.QIconDragEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QIconDragEvent_new(&outptr_QIconDragEvent, &outptr_QEvent)
	ret := newQIconDragEvent(outptr_QIconDragEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QIconDragEvent) Clone() *QIconDragEvent {
	return UnsafeNewQIconDragEvent(unsafe.Pointer(C.QIconDragEvent_Clone(this.h)), nil)
}

func (this *QIconDragEvent) callVirtualBase_Clone() *QIconDragEvent {

	return UnsafeNewQIconDragEvent(unsafe.Pointer(C.QIconDragEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QIconDragEvent) OnClone(slot func(super func() *QIconDragEvent) *QIconDragEvent) {
	C.QIconDragEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconDragEvent_Clone
func miqt_exec_callback_QIconDragEvent_Clone(self *C.QIconDragEvent, cb C.intptr_t) *C.QIconDragEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QIconDragEvent) *QIconDragEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconDragEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QIconDragEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QIconDragEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QIconDragEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QIconDragEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconDragEvent_SetAccepted
func miqt_exec_callback_QIconDragEvent_SetAccepted(self *C.QIconDragEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QIconDragEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QIconDragEvent) Delete() {
	C.QIconDragEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QIconDragEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QIconDragEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QShowEvent struct {
	h          *C.QShowEvent
	isSubclass bool
	*QEvent
}

func (this *QShowEvent) cPointer() *C.QShowEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QShowEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQShowEvent constructs the type using only CGO pointers.
func newQShowEvent(h *C.QShowEvent, h_QEvent *C.QEvent) *QShowEvent {
	if h == nil {
		return nil
	}
	return &QShowEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQShowEvent constructs the type using only unsafe pointers.
func UnsafeNewQShowEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QShowEvent {
	if h == nil {
		return nil
	}

	return &QShowEvent{h: (*C.QShowEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQShowEvent constructs a new QShowEvent object.
func NewQShowEvent() *QShowEvent {
	var outptr_QShowEvent *C.QShowEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QShowEvent_new(&outptr_QShowEvent, &outptr_QEvent)
	ret := newQShowEvent(outptr_QShowEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QShowEvent) Clone() *QShowEvent {
	return UnsafeNewQShowEvent(unsafe.Pointer(C.QShowEvent_Clone(this.h)), nil)
}

func (this *QShowEvent) callVirtualBase_Clone() *QShowEvent {

	return UnsafeNewQShowEvent(unsafe.Pointer(C.QShowEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QShowEvent) OnClone(slot func(super func() *QShowEvent) *QShowEvent) {
	C.QShowEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShowEvent_Clone
func miqt_exec_callback_QShowEvent_Clone(self *C.QShowEvent, cb C.intptr_t) *C.QShowEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QShowEvent) *QShowEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QShowEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QShowEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QShowEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QShowEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QShowEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShowEvent_SetAccepted
func miqt_exec_callback_QShowEvent_SetAccepted(self *C.QShowEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QShowEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QShowEvent) Delete() {
	C.QShowEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QShowEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QShowEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QHideEvent struct {
	h          *C.QHideEvent
	isSubclass bool
	*QEvent
}

func (this *QHideEvent) cPointer() *C.QHideEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QHideEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQHideEvent constructs the type using only CGO pointers.
func newQHideEvent(h *C.QHideEvent, h_QEvent *C.QEvent) *QHideEvent {
	if h == nil {
		return nil
	}
	return &QHideEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQHideEvent constructs the type using only unsafe pointers.
func UnsafeNewQHideEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QHideEvent {
	if h == nil {
		return nil
	}

	return &QHideEvent{h: (*C.QHideEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQHideEvent constructs a new QHideEvent object.
func NewQHideEvent() *QHideEvent {
	var outptr_QHideEvent *C.QHideEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHideEvent_new(&outptr_QHideEvent, &outptr_QEvent)
	ret := newQHideEvent(outptr_QHideEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QHideEvent) Clone() *QHideEvent {
	return UnsafeNewQHideEvent(unsafe.Pointer(C.QHideEvent_Clone(this.h)), nil)
}

func (this *QHideEvent) callVirtualBase_Clone() *QHideEvent {

	return UnsafeNewQHideEvent(unsafe.Pointer(C.QHideEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QHideEvent) OnClone(slot func(super func() *QHideEvent) *QHideEvent) {
	C.QHideEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHideEvent_Clone
func miqt_exec_callback_QHideEvent_Clone(self *C.QHideEvent, cb C.intptr_t) *C.QHideEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QHideEvent) *QHideEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHideEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QHideEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QHideEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QHideEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QHideEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHideEvent_SetAccepted
func miqt_exec_callback_QHideEvent_SetAccepted(self *C.QHideEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QHideEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QHideEvent) Delete() {
	C.QHideEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QHideEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QHideEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QContextMenuEvent struct {
	h          *C.QContextMenuEvent
	isSubclass bool
	*QInputEvent
}

func (this *QContextMenuEvent) cPointer() *C.QContextMenuEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QContextMenuEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQContextMenuEvent constructs the type using only CGO pointers.
func newQContextMenuEvent(h *C.QContextMenuEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QContextMenuEvent {
	if h == nil {
		return nil
	}
	return &QContextMenuEvent{h: h,
		QInputEvent: newQInputEvent(h_QInputEvent, h_QEvent)}
}

// UnsafeNewQContextMenuEvent constructs the type using only unsafe pointers.
func UnsafeNewQContextMenuEvent(h unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QContextMenuEvent {
	if h == nil {
		return nil
	}

	return &QContextMenuEvent{h: (*C.QContextMenuEvent)(h),
		QInputEvent: UnsafeNewQInputEvent(h_QInputEvent, h_QEvent)}
}

// NewQContextMenuEvent constructs a new QContextMenuEvent object.
func NewQContextMenuEvent(reason QContextMenuEvent__Reason, pos *QPoint, globalPos *QPoint) *QContextMenuEvent {
	var outptr_QContextMenuEvent *C.QContextMenuEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QContextMenuEvent_new((C.int)(reason), pos.cPointer(), globalPos.cPointer(), &outptr_QContextMenuEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQContextMenuEvent(outptr_QContextMenuEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQContextMenuEvent2 constructs a new QContextMenuEvent object.
func NewQContextMenuEvent2(reason QContextMenuEvent__Reason, pos *QPoint) *QContextMenuEvent {
	var outptr_QContextMenuEvent *C.QContextMenuEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QContextMenuEvent_new2((C.int)(reason), pos.cPointer(), &outptr_QContextMenuEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQContextMenuEvent(outptr_QContextMenuEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQContextMenuEvent3 constructs a new QContextMenuEvent object.
func NewQContextMenuEvent3(reason QContextMenuEvent__Reason, pos *QPoint, globalPos *QPoint, modifiers KeyboardModifier) *QContextMenuEvent {
	var outptr_QContextMenuEvent *C.QContextMenuEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QContextMenuEvent_new3((C.int)(reason), pos.cPointer(), globalPos.cPointer(), (C.int)(modifiers), &outptr_QContextMenuEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQContextMenuEvent(outptr_QContextMenuEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QContextMenuEvent) Clone() *QContextMenuEvent {
	return UnsafeNewQContextMenuEvent(unsafe.Pointer(C.QContextMenuEvent_Clone(this.h)), nil, nil)
}

func (this *QContextMenuEvent) X() int {
	return (int)(C.QContextMenuEvent_X(this.h))
}

func (this *QContextMenuEvent) Y() int {
	return (int)(C.QContextMenuEvent_Y(this.h))
}

func (this *QContextMenuEvent) GlobalX() int {
	return (int)(C.QContextMenuEvent_GlobalX(this.h))
}

func (this *QContextMenuEvent) GlobalY() int {
	return (int)(C.QContextMenuEvent_GlobalY(this.h))
}

func (this *QContextMenuEvent) Pos() *QPoint {
	return UnsafeNewQPoint(unsafe.Pointer(C.QContextMenuEvent_Pos(this.h)))
}

func (this *QContextMenuEvent) GlobalPos() *QPoint {
	return UnsafeNewQPoint(unsafe.Pointer(C.QContextMenuEvent_GlobalPos(this.h)))
}

func (this *QContextMenuEvent) Reason() QContextMenuEvent__Reason {
	return (QContextMenuEvent__Reason)(C.QContextMenuEvent_Reason(this.h))
}

func (this *QContextMenuEvent) callVirtualBase_Clone() *QContextMenuEvent {

	return UnsafeNewQContextMenuEvent(unsafe.Pointer(C.QContextMenuEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil)
}
func (this *QContextMenuEvent) OnClone(slot func(super func() *QContextMenuEvent) *QContextMenuEvent) {
	C.QContextMenuEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QContextMenuEvent_Clone
func miqt_exec_callback_QContextMenuEvent_Clone(self *C.QContextMenuEvent, cb C.intptr_t) *C.QContextMenuEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QContextMenuEvent) *QContextMenuEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QContextMenuEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QContextMenuEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	C.QContextMenuEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (C.ulonglong)(timestamp))

}
func (this *QContextMenuEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	C.QContextMenuEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QContextMenuEvent_SetTimestamp
func miqt_exec_callback_QContextMenuEvent_SetTimestamp(self *C.QContextMenuEvent, cb C.intptr_t, timestamp C.ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QContextMenuEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

// Delete this object from C++ memory.
func (this *QContextMenuEvent) Delete() {
	C.QContextMenuEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QContextMenuEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QContextMenuEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QInputMethodEvent struct {
	h          *C.QInputMethodEvent
	isSubclass bool
	*QEvent
}

func (this *QInputMethodEvent) cPointer() *C.QInputMethodEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QInputMethodEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQInputMethodEvent constructs the type using only CGO pointers.
func newQInputMethodEvent(h *C.QInputMethodEvent, h_QEvent *C.QEvent) *QInputMethodEvent {
	if h == nil {
		return nil
	}
	return &QInputMethodEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQInputMethodEvent constructs the type using only unsafe pointers.
func UnsafeNewQInputMethodEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QInputMethodEvent {
	if h == nil {
		return nil
	}

	return &QInputMethodEvent{h: (*C.QInputMethodEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQInputMethodEvent constructs a new QInputMethodEvent object.
func NewQInputMethodEvent() *QInputMethodEvent {
	var outptr_QInputMethodEvent *C.QInputMethodEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QInputMethodEvent_new(&outptr_QInputMethodEvent, &outptr_QEvent)
	ret := newQInputMethodEvent(outptr_QInputMethodEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQInputMethodEvent2 constructs a new QInputMethodEvent object.
func NewQInputMethodEvent2(preeditText string, attributes []QInputMethodEvent__Attribute) *QInputMethodEvent {
	preeditText_ms := C.struct_miqt_string{}
	preeditText_ms.data = C.CString(preeditText)
	preeditText_ms.len = C.size_t(len(preeditText))
	defer C.free(unsafe.Pointer(preeditText_ms.data))
	attributes_CArray := (*[0xffff]*C.QInputMethodEvent__Attribute)(C.malloc(C.size_t(8 * len(attributes))))
	defer C.free(unsafe.Pointer(attributes_CArray))
	for i := range attributes {
		attributes_CArray[i] = attributes[i].cPointer()
	}
	attributes_ma := C.struct_miqt_array{len: C.size_t(len(attributes)), data: unsafe.Pointer(attributes_CArray)}
	var outptr_QInputMethodEvent *C.QInputMethodEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QInputMethodEvent_new2(preeditText_ms, attributes_ma, &outptr_QInputMethodEvent, &outptr_QEvent)
	ret := newQInputMethodEvent(outptr_QInputMethodEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QInputMethodEvent) Clone() *QInputMethodEvent {
	return UnsafeNewQInputMethodEvent(unsafe.Pointer(C.QInputMethodEvent_Clone(this.h)), nil)
}

func (this *QInputMethodEvent) SetCommitString(commitString string) {
	commitString_ms := C.struct_miqt_string{}
	commitString_ms.data = C.CString(commitString)
	commitString_ms.len = C.size_t(len(commitString))
	defer C.free(unsafe.Pointer(commitString_ms.data))
	C.QInputMethodEvent_SetCommitString(this.h, commitString_ms)
}

func (this *QInputMethodEvent) Attributes() []QInputMethodEvent__Attribute {
	var _ma C.struct_miqt_array = C.QInputMethodEvent_Attributes(this.h)
	_ret := make([]QInputMethodEvent__Attribute, int(_ma.len))
	_outCast := (*[0xffff]*C.QInputMethodEvent__Attribute)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_ret := _outCast[i]
		_lv_goptr := newQInputMethodEvent__Attribute(_lv_ret)
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QInputMethodEvent) PreeditString() string {
	var _ms C.struct_miqt_string = C.QInputMethodEvent_PreeditString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputMethodEvent) CommitString() string {
	var _ms C.struct_miqt_string = C.QInputMethodEvent_CommitString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputMethodEvent) ReplacementStart() int {
	return (int)(C.QInputMethodEvent_ReplacementStart(this.h))
}

func (this *QInputMethodEvent) ReplacementLength() int {
	return (int)(C.QInputMethodEvent_ReplacementLength(this.h))
}

func (this *QInputMethodEvent) SetCommitString2(commitString string, replaceFrom int) {
	commitString_ms := C.struct_miqt_string{}
	commitString_ms.data = C.CString(commitString)
	commitString_ms.len = C.size_t(len(commitString))
	defer C.free(unsafe.Pointer(commitString_ms.data))
	C.QInputMethodEvent_SetCommitString2(this.h, commitString_ms, (C.int)(replaceFrom))
}

func (this *QInputMethodEvent) SetCommitString3(commitString string, replaceFrom int, replaceLength int) {
	commitString_ms := C.struct_miqt_string{}
	commitString_ms.data = C.CString(commitString)
	commitString_ms.len = C.size_t(len(commitString))
	defer C.free(unsafe.Pointer(commitString_ms.data))
	C.QInputMethodEvent_SetCommitString3(this.h, commitString_ms, (C.int)(replaceFrom), (C.int)(replaceLength))
}

func (this *QInputMethodEvent) callVirtualBase_Clone() *QInputMethodEvent {

	return UnsafeNewQInputMethodEvent(unsafe.Pointer(C.QInputMethodEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QInputMethodEvent) OnClone(slot func(super func() *QInputMethodEvent) *QInputMethodEvent) {
	C.QInputMethodEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodEvent_Clone
func miqt_exec_callback_QInputMethodEvent_Clone(self *C.QInputMethodEvent, cb C.intptr_t) *C.QInputMethodEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QInputMethodEvent) *QInputMethodEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputMethodEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QInputMethodEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QInputMethodEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QInputMethodEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QInputMethodEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodEvent_SetAccepted
func miqt_exec_callback_QInputMethodEvent_SetAccepted(self *C.QInputMethodEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QInputMethodEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QInputMethodEvent) Delete() {
	C.QInputMethodEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QInputMethodEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QInputMethodEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QInputMethodQueryEvent struct {
	h          *C.QInputMethodQueryEvent
	isSubclass bool
	*QEvent
}

func (this *QInputMethodQueryEvent) cPointer() *C.QInputMethodQueryEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QInputMethodQueryEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQInputMethodQueryEvent constructs the type using only CGO pointers.
func newQInputMethodQueryEvent(h *C.QInputMethodQueryEvent, h_QEvent *C.QEvent) *QInputMethodQueryEvent {
	if h == nil {
		return nil
	}
	return &QInputMethodQueryEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQInputMethodQueryEvent constructs the type using only unsafe pointers.
func UnsafeNewQInputMethodQueryEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QInputMethodQueryEvent {
	if h == nil {
		return nil
	}

	return &QInputMethodQueryEvent{h: (*C.QInputMethodQueryEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQInputMethodQueryEvent constructs a new QInputMethodQueryEvent object.
func NewQInputMethodQueryEvent(queries InputMethodQuery) *QInputMethodQueryEvent {
	var outptr_QInputMethodQueryEvent *C.QInputMethodQueryEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QInputMethodQueryEvent_new((C.int)(queries), &outptr_QInputMethodQueryEvent, &outptr_QEvent)
	ret := newQInputMethodQueryEvent(outptr_QInputMethodQueryEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QInputMethodQueryEvent) Clone() *QInputMethodQueryEvent {
	return UnsafeNewQInputMethodQueryEvent(unsafe.Pointer(C.QInputMethodQueryEvent_Clone(this.h)), nil)
}

func (this *QInputMethodQueryEvent) Queries() InputMethodQuery {
	return (InputMethodQuery)(C.QInputMethodQueryEvent_Queries(this.h))
}

func (this *QInputMethodQueryEvent) SetValue(query InputMethodQuery, value *QVariant) {
	C.QInputMethodQueryEvent_SetValue(this.h, (C.int)(query), value.cPointer())
}

func (this *QInputMethodQueryEvent) Value(query InputMethodQuery) *QVariant {
	_ret := C.QInputMethodQueryEvent_Value(this.h, (C.int)(query))
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethodQueryEvent) callVirtualBase_Clone() *QInputMethodQueryEvent {

	return UnsafeNewQInputMethodQueryEvent(unsafe.Pointer(C.QInputMethodQueryEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QInputMethodQueryEvent) OnClone(slot func(super func() *QInputMethodQueryEvent) *QInputMethodQueryEvent) {
	C.QInputMethodQueryEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodQueryEvent_Clone
func miqt_exec_callback_QInputMethodQueryEvent_Clone(self *C.QInputMethodQueryEvent, cb C.intptr_t) *C.QInputMethodQueryEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QInputMethodQueryEvent) *QInputMethodQueryEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputMethodQueryEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QInputMethodQueryEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QInputMethodQueryEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QInputMethodQueryEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QInputMethodQueryEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodQueryEvent_SetAccepted
func miqt_exec_callback_QInputMethodQueryEvent_SetAccepted(self *C.QInputMethodQueryEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QInputMethodQueryEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QInputMethodQueryEvent) Delete() {
	C.QInputMethodQueryEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QInputMethodQueryEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QInputMethodQueryEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QDropEvent struct {
	h          *C.QDropEvent
	isSubclass bool
	*QEvent
}

func (this *QDropEvent) cPointer() *C.QDropEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QDropEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQDropEvent constructs the type using only CGO pointers.
func newQDropEvent(h *C.QDropEvent, h_QEvent *C.QEvent) *QDropEvent {
	if h == nil {
		return nil
	}
	return &QDropEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQDropEvent constructs the type using only unsafe pointers.
func UnsafeNewQDropEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QDropEvent {
	if h == nil {
		return nil
	}

	return &QDropEvent{h: (*C.QDropEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQDropEvent constructs a new QDropEvent object.
func NewQDropEvent(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDropEvent {
	var outptr_QDropEvent *C.QDropEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QDropEvent_new(pos.cPointer(), (C.int)(actions), data.cPointer(), (C.int)(buttons), (C.int)(modifiers), &outptr_QDropEvent, &outptr_QEvent)
	ret := newQDropEvent(outptr_QDropEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQDropEvent2 constructs a new QDropEvent object.
func NewQDropEvent2(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, typeVal QEvent__Type) *QDropEvent {
	var outptr_QDropEvent *C.QDropEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QDropEvent_new2(pos.cPointer(), (C.int)(actions), data.cPointer(), (C.int)(buttons), (C.int)(modifiers), (C.int)(typeVal), &outptr_QDropEvent, &outptr_QEvent)
	ret := newQDropEvent(outptr_QDropEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QDropEvent) Clone() *QDropEvent {
	return UnsafeNewQDropEvent(unsafe.Pointer(C.QDropEvent_Clone(this.h)), nil)
}

func (this *QDropEvent) Pos() *QPoint {
	_ret := C.QDropEvent_Pos(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDropEvent) PosF() *QPointF {
	_ret := C.QDropEvent_PosF(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDropEvent) MouseButtons() MouseButton {
	return (MouseButton)(C.QDropEvent_MouseButtons(this.h))
}

func (this *QDropEvent) KeyboardModifiers() KeyboardModifier {
	return (KeyboardModifier)(C.QDropEvent_KeyboardModifiers(this.h))
}

func (this *QDropEvent) Position() *QPointF {
	_ret := C.QDropEvent_Position(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDropEvent) Buttons() MouseButton {
	return (MouseButton)(C.QDropEvent_Buttons(this.h))
}

func (this *QDropEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(C.QDropEvent_Modifiers(this.h))
}

func (this *QDropEvent) PossibleActions() DropAction {
	return (DropAction)(C.QDropEvent_PossibleActions(this.h))
}

func (this *QDropEvent) ProposedAction() DropAction {
	return (DropAction)(C.QDropEvent_ProposedAction(this.h))
}

func (this *QDropEvent) AcceptProposedAction() {
	C.QDropEvent_AcceptProposedAction(this.h)
}

func (this *QDropEvent) DropAction() DropAction {
	return (DropAction)(C.QDropEvent_DropAction(this.h))
}

func (this *QDropEvent) SetDropAction(action DropAction) {
	C.QDropEvent_SetDropAction(this.h, (C.int)(action))
}

func (this *QDropEvent) Source() *QObject {
	return UnsafeNewQObject(unsafe.Pointer(C.QDropEvent_Source(this.h)))
}

func (this *QDropEvent) MimeData() *QMimeData {
	return UnsafeNewQMimeData(unsafe.Pointer(C.QDropEvent_MimeData(this.h)), nil)
}

func (this *QDropEvent) callVirtualBase_Clone() *QDropEvent {

	return UnsafeNewQDropEvent(unsafe.Pointer(C.QDropEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QDropEvent) OnClone(slot func(super func() *QDropEvent) *QDropEvent) {
	C.QDropEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDropEvent_Clone
func miqt_exec_callback_QDropEvent_Clone(self *C.QDropEvent, cb C.intptr_t) *C.QDropEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDropEvent) *QDropEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDropEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QDropEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QDropEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QDropEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QDropEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDropEvent_SetAccepted
func miqt_exec_callback_QDropEvent_SetAccepted(self *C.QDropEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QDropEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QDropEvent) Delete() {
	C.QDropEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QDropEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QDropEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QDragMoveEvent struct {
	h          *C.QDragMoveEvent
	isSubclass bool
	*QDropEvent
}

func (this *QDragMoveEvent) cPointer() *C.QDragMoveEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QDragMoveEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQDragMoveEvent constructs the type using only CGO pointers.
func newQDragMoveEvent(h *C.QDragMoveEvent, h_QDropEvent *C.QDropEvent, h_QEvent *C.QEvent) *QDragMoveEvent {
	if h == nil {
		return nil
	}
	return &QDragMoveEvent{h: h,
		QDropEvent: newQDropEvent(h_QDropEvent, h_QEvent)}
}

// UnsafeNewQDragMoveEvent constructs the type using only unsafe pointers.
func UnsafeNewQDragMoveEvent(h unsafe.Pointer, h_QDropEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QDragMoveEvent {
	if h == nil {
		return nil
	}

	return &QDragMoveEvent{h: (*C.QDragMoveEvent)(h),
		QDropEvent: UnsafeNewQDropEvent(h_QDropEvent, h_QEvent)}
}

// NewQDragMoveEvent constructs a new QDragMoveEvent object.
func NewQDragMoveEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragMoveEvent {
	var outptr_QDragMoveEvent *C.QDragMoveEvent = nil
	var outptr_QDropEvent *C.QDropEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QDragMoveEvent_new(pos.cPointer(), (C.int)(actions), data.cPointer(), (C.int)(buttons), (C.int)(modifiers), &outptr_QDragMoveEvent, &outptr_QDropEvent, &outptr_QEvent)
	ret := newQDragMoveEvent(outptr_QDragMoveEvent, outptr_QDropEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQDragMoveEvent2 constructs a new QDragMoveEvent object.
func NewQDragMoveEvent2(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, typeVal QEvent__Type) *QDragMoveEvent {
	var outptr_QDragMoveEvent *C.QDragMoveEvent = nil
	var outptr_QDropEvent *C.QDropEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QDragMoveEvent_new2(pos.cPointer(), (C.int)(actions), data.cPointer(), (C.int)(buttons), (C.int)(modifiers), (C.int)(typeVal), &outptr_QDragMoveEvent, &outptr_QDropEvent, &outptr_QEvent)
	ret := newQDragMoveEvent(outptr_QDragMoveEvent, outptr_QDropEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QDragMoveEvent) Clone() *QDragMoveEvent {
	return UnsafeNewQDragMoveEvent(unsafe.Pointer(C.QDragMoveEvent_Clone(this.h)), nil, nil)
}

func (this *QDragMoveEvent) AnswerRect() *QRect {
	_ret := C.QDragMoveEvent_AnswerRect(this.h)
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDragMoveEvent) Accept() {
	C.QDragMoveEvent_Accept(this.h)
}

func (this *QDragMoveEvent) Ignore() {
	C.QDragMoveEvent_Ignore(this.h)
}

func (this *QDragMoveEvent) AcceptWithQRect(r *QRect) {
	C.QDragMoveEvent_AcceptWithQRect(this.h, r.cPointer())
}

func (this *QDragMoveEvent) IgnoreWithQRect(r *QRect) {
	C.QDragMoveEvent_IgnoreWithQRect(this.h, r.cPointer())
}

func (this *QDragMoveEvent) callVirtualBase_Clone() *QDragMoveEvent {

	return UnsafeNewQDragMoveEvent(unsafe.Pointer(C.QDragMoveEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil)
}
func (this *QDragMoveEvent) OnClone(slot func(super func() *QDragMoveEvent) *QDragMoveEvent) {
	C.QDragMoveEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragMoveEvent_Clone
func miqt_exec_callback_QDragMoveEvent_Clone(self *C.QDragMoveEvent, cb C.intptr_t) *C.QDragMoveEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDragMoveEvent) *QDragMoveEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDragMoveEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

// Delete this object from C++ memory.
func (this *QDragMoveEvent) Delete() {
	C.QDragMoveEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QDragMoveEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QDragMoveEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QDragEnterEvent struct {
	h          *C.QDragEnterEvent
	isSubclass bool
	*QDragMoveEvent
}

func (this *QDragEnterEvent) cPointer() *C.QDragEnterEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QDragEnterEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQDragEnterEvent constructs the type using only CGO pointers.
func newQDragEnterEvent(h *C.QDragEnterEvent, h_QDragMoveEvent *C.QDragMoveEvent, h_QDropEvent *C.QDropEvent, h_QEvent *C.QEvent) *QDragEnterEvent {
	if h == nil {
		return nil
	}
	return &QDragEnterEvent{h: h,
		QDragMoveEvent: newQDragMoveEvent(h_QDragMoveEvent, h_QDropEvent, h_QEvent)}
}

// UnsafeNewQDragEnterEvent constructs the type using only unsafe pointers.
func UnsafeNewQDragEnterEvent(h unsafe.Pointer, h_QDragMoveEvent unsafe.Pointer, h_QDropEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QDragEnterEvent {
	if h == nil {
		return nil
	}

	return &QDragEnterEvent{h: (*C.QDragEnterEvent)(h),
		QDragMoveEvent: UnsafeNewQDragMoveEvent(h_QDragMoveEvent, h_QDropEvent, h_QEvent)}
}

// NewQDragEnterEvent constructs a new QDragEnterEvent object.
func NewQDragEnterEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragEnterEvent {
	var outptr_QDragEnterEvent *C.QDragEnterEvent = nil
	var outptr_QDragMoveEvent *C.QDragMoveEvent = nil
	var outptr_QDropEvent *C.QDropEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QDragEnterEvent_new(pos.cPointer(), (C.int)(actions), data.cPointer(), (C.int)(buttons), (C.int)(modifiers), &outptr_QDragEnterEvent, &outptr_QDragMoveEvent, &outptr_QDropEvent, &outptr_QEvent)
	ret := newQDragEnterEvent(outptr_QDragEnterEvent, outptr_QDragMoveEvent, outptr_QDropEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QDragEnterEvent) Clone() *QDragEnterEvent {
	return UnsafeNewQDragEnterEvent(unsafe.Pointer(C.QDragEnterEvent_Clone(this.h)), nil, nil, nil)
}

func (this *QDragEnterEvent) callVirtualBase_Clone() *QDragEnterEvent {

	return UnsafeNewQDragEnterEvent(unsafe.Pointer(C.QDragEnterEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil)
}
func (this *QDragEnterEvent) OnClone(slot func(super func() *QDragEnterEvent) *QDragEnterEvent) {
	C.QDragEnterEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragEnterEvent_Clone
func miqt_exec_callback_QDragEnterEvent_Clone(self *C.QDragEnterEvent, cb C.intptr_t) *C.QDragEnterEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDragEnterEvent) *QDragEnterEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDragEnterEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

// Delete this object from C++ memory.
func (this *QDragEnterEvent) Delete() {
	C.QDragEnterEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QDragEnterEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QDragEnterEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QDragLeaveEvent struct {
	h          *C.QDragLeaveEvent
	isSubclass bool
	*QEvent
}

func (this *QDragLeaveEvent) cPointer() *C.QDragLeaveEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QDragLeaveEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQDragLeaveEvent constructs the type using only CGO pointers.
func newQDragLeaveEvent(h *C.QDragLeaveEvent, h_QEvent *C.QEvent) *QDragLeaveEvent {
	if h == nil {
		return nil
	}
	return &QDragLeaveEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQDragLeaveEvent constructs the type using only unsafe pointers.
func UnsafeNewQDragLeaveEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QDragLeaveEvent {
	if h == nil {
		return nil
	}

	return &QDragLeaveEvent{h: (*C.QDragLeaveEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQDragLeaveEvent constructs a new QDragLeaveEvent object.
func NewQDragLeaveEvent() *QDragLeaveEvent {
	var outptr_QDragLeaveEvent *C.QDragLeaveEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QDragLeaveEvent_new(&outptr_QDragLeaveEvent, &outptr_QEvent)
	ret := newQDragLeaveEvent(outptr_QDragLeaveEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QDragLeaveEvent) Clone() *QDragLeaveEvent {
	return UnsafeNewQDragLeaveEvent(unsafe.Pointer(C.QDragLeaveEvent_Clone(this.h)), nil)
}

func (this *QDragLeaveEvent) callVirtualBase_Clone() *QDragLeaveEvent {

	return UnsafeNewQDragLeaveEvent(unsafe.Pointer(C.QDragLeaveEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QDragLeaveEvent) OnClone(slot func(super func() *QDragLeaveEvent) *QDragLeaveEvent) {
	C.QDragLeaveEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragLeaveEvent_Clone
func miqt_exec_callback_QDragLeaveEvent_Clone(self *C.QDragLeaveEvent, cb C.intptr_t) *C.QDragLeaveEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDragLeaveEvent) *QDragLeaveEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDragLeaveEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QDragLeaveEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QDragLeaveEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QDragLeaveEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QDragLeaveEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragLeaveEvent_SetAccepted
func miqt_exec_callback_QDragLeaveEvent_SetAccepted(self *C.QDragLeaveEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QDragLeaveEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QDragLeaveEvent) Delete() {
	C.QDragLeaveEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QDragLeaveEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QDragLeaveEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QHelpEvent struct {
	h          *C.QHelpEvent
	isSubclass bool
	*QEvent
}

func (this *QHelpEvent) cPointer() *C.QHelpEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QHelpEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQHelpEvent constructs the type using only CGO pointers.
func newQHelpEvent(h *C.QHelpEvent, h_QEvent *C.QEvent) *QHelpEvent {
	if h == nil {
		return nil
	}
	return &QHelpEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQHelpEvent constructs the type using only unsafe pointers.
func UnsafeNewQHelpEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QHelpEvent {
	if h == nil {
		return nil
	}

	return &QHelpEvent{h: (*C.QHelpEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQHelpEvent constructs a new QHelpEvent object.
func NewQHelpEvent(typeVal QEvent__Type, pos *QPoint, globalPos *QPoint) *QHelpEvent {
	var outptr_QHelpEvent *C.QHelpEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QHelpEvent_new((C.int)(typeVal), pos.cPointer(), globalPos.cPointer(), &outptr_QHelpEvent, &outptr_QEvent)
	ret := newQHelpEvent(outptr_QHelpEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QHelpEvent) Clone() *QHelpEvent {
	return UnsafeNewQHelpEvent(unsafe.Pointer(C.QHelpEvent_Clone(this.h)), nil)
}

func (this *QHelpEvent) X() int {
	return (int)(C.QHelpEvent_X(this.h))
}

func (this *QHelpEvent) Y() int {
	return (int)(C.QHelpEvent_Y(this.h))
}

func (this *QHelpEvent) GlobalX() int {
	return (int)(C.QHelpEvent_GlobalX(this.h))
}

func (this *QHelpEvent) GlobalY() int {
	return (int)(C.QHelpEvent_GlobalY(this.h))
}

func (this *QHelpEvent) Pos() *QPoint {
	return UnsafeNewQPoint(unsafe.Pointer(C.QHelpEvent_Pos(this.h)))
}

func (this *QHelpEvent) GlobalPos() *QPoint {
	return UnsafeNewQPoint(unsafe.Pointer(C.QHelpEvent_GlobalPos(this.h)))
}

func (this *QHelpEvent) callVirtualBase_Clone() *QHelpEvent {

	return UnsafeNewQHelpEvent(unsafe.Pointer(C.QHelpEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QHelpEvent) OnClone(slot func(super func() *QHelpEvent) *QHelpEvent) {
	C.QHelpEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHelpEvent_Clone
func miqt_exec_callback_QHelpEvent_Clone(self *C.QHelpEvent, cb C.intptr_t) *C.QHelpEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QHelpEvent) *QHelpEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHelpEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QHelpEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QHelpEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QHelpEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QHelpEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHelpEvent_SetAccepted
func miqt_exec_callback_QHelpEvent_SetAccepted(self *C.QHelpEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QHelpEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QHelpEvent) Delete() {
	C.QHelpEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QHelpEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QHelpEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QStatusTipEvent struct {
	h          *C.QStatusTipEvent
	isSubclass bool
	*QEvent
}

func (this *QStatusTipEvent) cPointer() *C.QStatusTipEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QStatusTipEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQStatusTipEvent constructs the type using only CGO pointers.
func newQStatusTipEvent(h *C.QStatusTipEvent, h_QEvent *C.QEvent) *QStatusTipEvent {
	if h == nil {
		return nil
	}
	return &QStatusTipEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQStatusTipEvent constructs the type using only unsafe pointers.
func UnsafeNewQStatusTipEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QStatusTipEvent {
	if h == nil {
		return nil
	}

	return &QStatusTipEvent{h: (*C.QStatusTipEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQStatusTipEvent constructs a new QStatusTipEvent object.
func NewQStatusTipEvent(tip string) *QStatusTipEvent {
	tip_ms := C.struct_miqt_string{}
	tip_ms.data = C.CString(tip)
	tip_ms.len = C.size_t(len(tip))
	defer C.free(unsafe.Pointer(tip_ms.data))
	var outptr_QStatusTipEvent *C.QStatusTipEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QStatusTipEvent_new(tip_ms, &outptr_QStatusTipEvent, &outptr_QEvent)
	ret := newQStatusTipEvent(outptr_QStatusTipEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QStatusTipEvent) Clone() *QStatusTipEvent {
	return UnsafeNewQStatusTipEvent(unsafe.Pointer(C.QStatusTipEvent_Clone(this.h)), nil)
}

func (this *QStatusTipEvent) Tip() string {
	var _ms C.struct_miqt_string = C.QStatusTipEvent_Tip(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStatusTipEvent) callVirtualBase_Clone() *QStatusTipEvent {

	return UnsafeNewQStatusTipEvent(unsafe.Pointer(C.QStatusTipEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QStatusTipEvent) OnClone(slot func(super func() *QStatusTipEvent) *QStatusTipEvent) {
	C.QStatusTipEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusTipEvent_Clone
func miqt_exec_callback_QStatusTipEvent_Clone(self *C.QStatusTipEvent, cb C.intptr_t) *C.QStatusTipEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QStatusTipEvent) *QStatusTipEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusTipEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QStatusTipEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QStatusTipEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QStatusTipEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QStatusTipEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusTipEvent_SetAccepted
func miqt_exec_callback_QStatusTipEvent_SetAccepted(self *C.QStatusTipEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QStatusTipEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QStatusTipEvent) Delete() {
	C.QStatusTipEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QStatusTipEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QStatusTipEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QWhatsThisClickedEvent struct {
	h          *C.QWhatsThisClickedEvent
	isSubclass bool
	*QEvent
}

func (this *QWhatsThisClickedEvent) cPointer() *C.QWhatsThisClickedEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWhatsThisClickedEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWhatsThisClickedEvent constructs the type using only CGO pointers.
func newQWhatsThisClickedEvent(h *C.QWhatsThisClickedEvent, h_QEvent *C.QEvent) *QWhatsThisClickedEvent {
	if h == nil {
		return nil
	}
	return &QWhatsThisClickedEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQWhatsThisClickedEvent constructs the type using only unsafe pointers.
func UnsafeNewQWhatsThisClickedEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QWhatsThisClickedEvent {
	if h == nil {
		return nil
	}

	return &QWhatsThisClickedEvent{h: (*C.QWhatsThisClickedEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQWhatsThisClickedEvent constructs a new QWhatsThisClickedEvent object.
func NewQWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	href_ms := C.struct_miqt_string{}
	href_ms.data = C.CString(href)
	href_ms.len = C.size_t(len(href))
	defer C.free(unsafe.Pointer(href_ms.data))
	var outptr_QWhatsThisClickedEvent *C.QWhatsThisClickedEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QWhatsThisClickedEvent_new(href_ms, &outptr_QWhatsThisClickedEvent, &outptr_QEvent)
	ret := newQWhatsThisClickedEvent(outptr_QWhatsThisClickedEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QWhatsThisClickedEvent) Clone() *QWhatsThisClickedEvent {
	return UnsafeNewQWhatsThisClickedEvent(unsafe.Pointer(C.QWhatsThisClickedEvent_Clone(this.h)), nil)
}

func (this *QWhatsThisClickedEvent) Href() string {
	var _ms C.struct_miqt_string = C.QWhatsThisClickedEvent_Href(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWhatsThisClickedEvent) callVirtualBase_Clone() *QWhatsThisClickedEvent {

	return UnsafeNewQWhatsThisClickedEvent(unsafe.Pointer(C.QWhatsThisClickedEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QWhatsThisClickedEvent) OnClone(slot func(super func() *QWhatsThisClickedEvent) *QWhatsThisClickedEvent) {
	C.QWhatsThisClickedEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWhatsThisClickedEvent_Clone
func miqt_exec_callback_QWhatsThisClickedEvent_Clone(self *C.QWhatsThisClickedEvent, cb C.intptr_t) *C.QWhatsThisClickedEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWhatsThisClickedEvent) *QWhatsThisClickedEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWhatsThisClickedEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QWhatsThisClickedEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QWhatsThisClickedEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QWhatsThisClickedEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QWhatsThisClickedEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWhatsThisClickedEvent_SetAccepted
func miqt_exec_callback_QWhatsThisClickedEvent_SetAccepted(self *C.QWhatsThisClickedEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QWhatsThisClickedEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QWhatsThisClickedEvent) Delete() {
	C.QWhatsThisClickedEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWhatsThisClickedEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QWhatsThisClickedEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QActionEvent struct {
	h          *C.QActionEvent
	isSubclass bool
	*QEvent
}

func (this *QActionEvent) cPointer() *C.QActionEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QActionEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQActionEvent constructs the type using only CGO pointers.
func newQActionEvent(h *C.QActionEvent, h_QEvent *C.QEvent) *QActionEvent {
	if h == nil {
		return nil
	}
	return &QActionEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQActionEvent constructs the type using only unsafe pointers.
func UnsafeNewQActionEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QActionEvent {
	if h == nil {
		return nil
	}

	return &QActionEvent{h: (*C.QActionEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQActionEvent constructs a new QActionEvent object.
func NewQActionEvent(typeVal int, action *QAction) *QActionEvent {
	var outptr_QActionEvent *C.QActionEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QActionEvent_new((C.int)(typeVal), action.cPointer(), &outptr_QActionEvent, &outptr_QEvent)
	ret := newQActionEvent(outptr_QActionEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQActionEvent2 constructs a new QActionEvent object.
func NewQActionEvent2(typeVal int, action *QAction, before *QAction) *QActionEvent {
	var outptr_QActionEvent *C.QActionEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QActionEvent_new2((C.int)(typeVal), action.cPointer(), before.cPointer(), &outptr_QActionEvent, &outptr_QEvent)
	ret := newQActionEvent(outptr_QActionEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QActionEvent) Clone() *QActionEvent {
	return UnsafeNewQActionEvent(unsafe.Pointer(C.QActionEvent_Clone(this.h)), nil)
}

func (this *QActionEvent) Action() *QAction {
	return UnsafeNewQAction(unsafe.Pointer(C.QActionEvent_Action(this.h)), nil)
}

func (this *QActionEvent) Before() *QAction {
	return UnsafeNewQAction(unsafe.Pointer(C.QActionEvent_Before(this.h)), nil)
}

func (this *QActionEvent) callVirtualBase_Clone() *QActionEvent {

	return UnsafeNewQActionEvent(unsafe.Pointer(C.QActionEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QActionEvent) OnClone(slot func(super func() *QActionEvent) *QActionEvent) {
	C.QActionEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionEvent_Clone
func miqt_exec_callback_QActionEvent_Clone(self *C.QActionEvent, cb C.intptr_t) *C.QActionEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QActionEvent) *QActionEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QActionEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QActionEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QActionEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QActionEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QActionEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionEvent_SetAccepted
func miqt_exec_callback_QActionEvent_SetAccepted(self *C.QActionEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QActionEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QActionEvent) Delete() {
	C.QActionEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QActionEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QActionEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QFileOpenEvent struct {
	h          *C.QFileOpenEvent
	isSubclass bool
	*QEvent
}

func (this *QFileOpenEvent) cPointer() *C.QFileOpenEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QFileOpenEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQFileOpenEvent constructs the type using only CGO pointers.
func newQFileOpenEvent(h *C.QFileOpenEvent, h_QEvent *C.QEvent) *QFileOpenEvent {
	if h == nil {
		return nil
	}
	return &QFileOpenEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQFileOpenEvent constructs the type using only unsafe pointers.
func UnsafeNewQFileOpenEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QFileOpenEvent {
	if h == nil {
		return nil
	}

	return &QFileOpenEvent{h: (*C.QFileOpenEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQFileOpenEvent constructs a new QFileOpenEvent object.
func NewQFileOpenEvent(file string) *QFileOpenEvent {
	file_ms := C.struct_miqt_string{}
	file_ms.data = C.CString(file)
	file_ms.len = C.size_t(len(file))
	defer C.free(unsafe.Pointer(file_ms.data))
	var outptr_QFileOpenEvent *C.QFileOpenEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QFileOpenEvent_new(file_ms, &outptr_QFileOpenEvent, &outptr_QEvent)
	ret := newQFileOpenEvent(outptr_QFileOpenEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQFileOpenEvent2 constructs a new QFileOpenEvent object.
func NewQFileOpenEvent2(url *QUrl) *QFileOpenEvent {
	var outptr_QFileOpenEvent *C.QFileOpenEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QFileOpenEvent_new2(url.cPointer(), &outptr_QFileOpenEvent, &outptr_QEvent)
	ret := newQFileOpenEvent(outptr_QFileOpenEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QFileOpenEvent) Clone() *QFileOpenEvent {
	return UnsafeNewQFileOpenEvent(unsafe.Pointer(C.QFileOpenEvent_Clone(this.h)), nil)
}

func (this *QFileOpenEvent) File() string {
	var _ms C.struct_miqt_string = C.QFileOpenEvent_File(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileOpenEvent) Url() *QUrl {
	_ret := C.QFileOpenEvent_Url(this.h)
	_goptr := newQUrl(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileOpenEvent) OpenFile(file *QFile, flags QIODeviceBase__OpenModeFlag) bool {
	return (bool)(C.QFileOpenEvent_OpenFile(this.h, file.cPointer(), (C.int)(flags)))
}

func (this *QFileOpenEvent) callVirtualBase_Clone() *QFileOpenEvent {

	return UnsafeNewQFileOpenEvent(unsafe.Pointer(C.QFileOpenEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QFileOpenEvent) OnClone(slot func(super func() *QFileOpenEvent) *QFileOpenEvent) {
	C.QFileOpenEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileOpenEvent_Clone
func miqt_exec_callback_QFileOpenEvent_Clone(self *C.QFileOpenEvent, cb C.intptr_t) *C.QFileOpenEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QFileOpenEvent) *QFileOpenEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileOpenEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QFileOpenEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QFileOpenEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QFileOpenEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QFileOpenEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileOpenEvent_SetAccepted
func miqt_exec_callback_QFileOpenEvent_SetAccepted(self *C.QFileOpenEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QFileOpenEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QFileOpenEvent) Delete() {
	C.QFileOpenEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QFileOpenEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QFileOpenEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QToolBarChangeEvent struct {
	h          *C.QToolBarChangeEvent
	isSubclass bool
	*QEvent
}

func (this *QToolBarChangeEvent) cPointer() *C.QToolBarChangeEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QToolBarChangeEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQToolBarChangeEvent constructs the type using only CGO pointers.
func newQToolBarChangeEvent(h *C.QToolBarChangeEvent, h_QEvent *C.QEvent) *QToolBarChangeEvent {
	if h == nil {
		return nil
	}
	return &QToolBarChangeEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQToolBarChangeEvent constructs the type using only unsafe pointers.
func UnsafeNewQToolBarChangeEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QToolBarChangeEvent {
	if h == nil {
		return nil
	}

	return &QToolBarChangeEvent{h: (*C.QToolBarChangeEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQToolBarChangeEvent constructs a new QToolBarChangeEvent object.
func NewQToolBarChangeEvent(t bool) *QToolBarChangeEvent {
	var outptr_QToolBarChangeEvent *C.QToolBarChangeEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QToolBarChangeEvent_new((C.bool)(t), &outptr_QToolBarChangeEvent, &outptr_QEvent)
	ret := newQToolBarChangeEvent(outptr_QToolBarChangeEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QToolBarChangeEvent) Clone() *QToolBarChangeEvent {
	return UnsafeNewQToolBarChangeEvent(unsafe.Pointer(C.QToolBarChangeEvent_Clone(this.h)), nil)
}

func (this *QToolBarChangeEvent) Toggle() bool {
	return (bool)(C.QToolBarChangeEvent_Toggle(this.h))
}

func (this *QToolBarChangeEvent) callVirtualBase_Clone() *QToolBarChangeEvent {

	return UnsafeNewQToolBarChangeEvent(unsafe.Pointer(C.QToolBarChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QToolBarChangeEvent) OnClone(slot func(super func() *QToolBarChangeEvent) *QToolBarChangeEvent) {
	C.QToolBarChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBarChangeEvent_Clone
func miqt_exec_callback_QToolBarChangeEvent_Clone(self *C.QToolBarChangeEvent, cb C.intptr_t) *C.QToolBarChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QToolBarChangeEvent) *QToolBarChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBarChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QToolBarChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QToolBarChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QToolBarChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QToolBarChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBarChangeEvent_SetAccepted
func miqt_exec_callback_QToolBarChangeEvent_SetAccepted(self *C.QToolBarChangeEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QToolBarChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QToolBarChangeEvent) Delete() {
	C.QToolBarChangeEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QToolBarChangeEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QToolBarChangeEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QShortcutEvent struct {
	h          *C.QShortcutEvent
	isSubclass bool
	*QEvent
}

func (this *QShortcutEvent) cPointer() *C.QShortcutEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QShortcutEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQShortcutEvent constructs the type using only CGO pointers.
func newQShortcutEvent(h *C.QShortcutEvent, h_QEvent *C.QEvent) *QShortcutEvent {
	if h == nil {
		return nil
	}
	return &QShortcutEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQShortcutEvent constructs the type using only unsafe pointers.
func UnsafeNewQShortcutEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QShortcutEvent {
	if h == nil {
		return nil
	}

	return &QShortcutEvent{h: (*C.QShortcutEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQShortcutEvent constructs a new QShortcutEvent object.
func NewQShortcutEvent(key *QKeySequence, id int) *QShortcutEvent {
	var outptr_QShortcutEvent *C.QShortcutEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QShortcutEvent_new(key.cPointer(), (C.int)(id), &outptr_QShortcutEvent, &outptr_QEvent)
	ret := newQShortcutEvent(outptr_QShortcutEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQShortcutEvent2 constructs a new QShortcutEvent object.
func NewQShortcutEvent2(key *QKeySequence, id int, ambiguous bool) *QShortcutEvent {
	var outptr_QShortcutEvent *C.QShortcutEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QShortcutEvent_new2(key.cPointer(), (C.int)(id), (C.bool)(ambiguous), &outptr_QShortcutEvent, &outptr_QEvent)
	ret := newQShortcutEvent(outptr_QShortcutEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QShortcutEvent) Clone() *QShortcutEvent {
	return UnsafeNewQShortcutEvent(unsafe.Pointer(C.QShortcutEvent_Clone(this.h)), nil)
}

func (this *QShortcutEvent) Key() *QKeySequence {
	return UnsafeNewQKeySequence(unsafe.Pointer(C.QShortcutEvent_Key(this.h)))
}

func (this *QShortcutEvent) ShortcutId() int {
	return (int)(C.QShortcutEvent_ShortcutId(this.h))
}

func (this *QShortcutEvent) IsAmbiguous() bool {
	return (bool)(C.QShortcutEvent_IsAmbiguous(this.h))
}

func (this *QShortcutEvent) callVirtualBase_Clone() *QShortcutEvent {

	return UnsafeNewQShortcutEvent(unsafe.Pointer(C.QShortcutEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QShortcutEvent) OnClone(slot func(super func() *QShortcutEvent) *QShortcutEvent) {
	C.QShortcutEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcutEvent_Clone
func miqt_exec_callback_QShortcutEvent_Clone(self *C.QShortcutEvent, cb C.intptr_t) *C.QShortcutEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QShortcutEvent) *QShortcutEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QShortcutEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QShortcutEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QShortcutEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QShortcutEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QShortcutEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcutEvent_SetAccepted
func miqt_exec_callback_QShortcutEvent_SetAccepted(self *C.QShortcutEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QShortcutEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QShortcutEvent) Delete() {
	C.QShortcutEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QShortcutEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QShortcutEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QWindowStateChangeEvent struct {
	h          *C.QWindowStateChangeEvent
	isSubclass bool
	*QEvent
}

func (this *QWindowStateChangeEvent) cPointer() *C.QWindowStateChangeEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWindowStateChangeEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWindowStateChangeEvent constructs the type using only CGO pointers.
func newQWindowStateChangeEvent(h *C.QWindowStateChangeEvent, h_QEvent *C.QEvent) *QWindowStateChangeEvent {
	if h == nil {
		return nil
	}
	return &QWindowStateChangeEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQWindowStateChangeEvent constructs the type using only unsafe pointers.
func UnsafeNewQWindowStateChangeEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QWindowStateChangeEvent {
	if h == nil {
		return nil
	}

	return &QWindowStateChangeEvent{h: (*C.QWindowStateChangeEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQWindowStateChangeEvent constructs a new QWindowStateChangeEvent object.
func NewQWindowStateChangeEvent(oldState WindowState) *QWindowStateChangeEvent {
	var outptr_QWindowStateChangeEvent *C.QWindowStateChangeEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QWindowStateChangeEvent_new((C.int)(oldState), &outptr_QWindowStateChangeEvent, &outptr_QEvent)
	ret := newQWindowStateChangeEvent(outptr_QWindowStateChangeEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQWindowStateChangeEvent2 constructs a new QWindowStateChangeEvent object.
func NewQWindowStateChangeEvent2(oldState WindowState, isOverride bool) *QWindowStateChangeEvent {
	var outptr_QWindowStateChangeEvent *C.QWindowStateChangeEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QWindowStateChangeEvent_new2((C.int)(oldState), (C.bool)(isOverride), &outptr_QWindowStateChangeEvent, &outptr_QEvent)
	ret := newQWindowStateChangeEvent(outptr_QWindowStateChangeEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QWindowStateChangeEvent) Clone() *QWindowStateChangeEvent {
	return UnsafeNewQWindowStateChangeEvent(unsafe.Pointer(C.QWindowStateChangeEvent_Clone(this.h)), nil)
}

func (this *QWindowStateChangeEvent) OldState() WindowState {
	return (WindowState)(C.QWindowStateChangeEvent_OldState(this.h))
}

func (this *QWindowStateChangeEvent) IsOverride() bool {
	return (bool)(C.QWindowStateChangeEvent_IsOverride(this.h))
}

func (this *QWindowStateChangeEvent) callVirtualBase_Clone() *QWindowStateChangeEvent {

	return UnsafeNewQWindowStateChangeEvent(unsafe.Pointer(C.QWindowStateChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QWindowStateChangeEvent) OnClone(slot func(super func() *QWindowStateChangeEvent) *QWindowStateChangeEvent) {
	C.QWindowStateChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowStateChangeEvent_Clone
func miqt_exec_callback_QWindowStateChangeEvent_Clone(self *C.QWindowStateChangeEvent, cb C.intptr_t) *C.QWindowStateChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWindowStateChangeEvent) *QWindowStateChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindowStateChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QWindowStateChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QWindowStateChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QWindowStateChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QWindowStateChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowStateChangeEvent_SetAccepted
func miqt_exec_callback_QWindowStateChangeEvent_SetAccepted(self *C.QWindowStateChangeEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QWindowStateChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QWindowStateChangeEvent) Delete() {
	C.QWindowStateChangeEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWindowStateChangeEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QWindowStateChangeEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QTouchEvent struct {
	h          *C.QTouchEvent
	isSubclass bool
	*QPointerEvent
}

func (this *QTouchEvent) cPointer() *C.QTouchEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QTouchEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQTouchEvent constructs the type using only CGO pointers.
func newQTouchEvent(h *C.QTouchEvent, h_QPointerEvent *C.QPointerEvent, h_QInputEvent *C.QInputEvent, h_QEvent *C.QEvent) *QTouchEvent {
	if h == nil {
		return nil
	}
	return &QTouchEvent{h: h,
		QPointerEvent: newQPointerEvent(h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// UnsafeNewQTouchEvent constructs the type using only unsafe pointers.
func UnsafeNewQTouchEvent(h unsafe.Pointer, h_QPointerEvent unsafe.Pointer, h_QInputEvent unsafe.Pointer, h_QEvent unsafe.Pointer) *QTouchEvent {
	if h == nil {
		return nil
	}

	return &QTouchEvent{h: (*C.QTouchEvent)(h),
		QPointerEvent: UnsafeNewQPointerEvent(h_QPointerEvent, h_QInputEvent, h_QEvent)}
}

// NewQTouchEvent constructs a new QTouchEvent object.
func NewQTouchEvent(eventType QEvent__Type) *QTouchEvent {
	var outptr_QTouchEvent *C.QTouchEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QTouchEvent_new((C.int)(eventType), &outptr_QTouchEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQTouchEvent(outptr_QTouchEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent2 constructs a new QTouchEvent object.
func NewQTouchEvent2(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates QEventPoint__State) *QTouchEvent {
	var outptr_QTouchEvent *C.QTouchEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QTouchEvent_new2((C.int)(eventType), device.cPointer(), (C.int)(modifiers), (C.uint8_t)(touchPointStates), &outptr_QTouchEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQTouchEvent(outptr_QTouchEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent3 constructs a new QTouchEvent object.
func NewQTouchEvent3(eventType QEvent__Type, device *QPointingDevice) *QTouchEvent {
	var outptr_QTouchEvent *C.QTouchEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QTouchEvent_new3((C.int)(eventType), device.cPointer(), &outptr_QTouchEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQTouchEvent(outptr_QTouchEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent4 constructs a new QTouchEvent object.
func NewQTouchEvent4(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier) *QTouchEvent {
	var outptr_QTouchEvent *C.QTouchEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QTouchEvent_new4((C.int)(eventType), device.cPointer(), (C.int)(modifiers), &outptr_QTouchEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQTouchEvent(outptr_QTouchEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent5 constructs a new QTouchEvent object.
func NewQTouchEvent5(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPoints []QEventPoint) *QTouchEvent {
	touchPoints_CArray := (*[0xffff]*C.QEventPoint)(C.malloc(C.size_t(8 * len(touchPoints))))
	defer C.free(unsafe.Pointer(touchPoints_CArray))
	for i := range touchPoints {
		touchPoints_CArray[i] = touchPoints[i].cPointer()
	}
	touchPoints_ma := C.struct_miqt_array{len: C.size_t(len(touchPoints)), data: unsafe.Pointer(touchPoints_CArray)}
	var outptr_QTouchEvent *C.QTouchEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QTouchEvent_new5((C.int)(eventType), device.cPointer(), (C.int)(modifiers), touchPoints_ma, &outptr_QTouchEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQTouchEvent(outptr_QTouchEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent6 constructs a new QTouchEvent object.
func NewQTouchEvent6(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates QEventPoint__State, touchPoints []QEventPoint) *QTouchEvent {
	touchPoints_CArray := (*[0xffff]*C.QEventPoint)(C.malloc(C.size_t(8 * len(touchPoints))))
	defer C.free(unsafe.Pointer(touchPoints_CArray))
	for i := range touchPoints {
		touchPoints_CArray[i] = touchPoints[i].cPointer()
	}
	touchPoints_ma := C.struct_miqt_array{len: C.size_t(len(touchPoints)), data: unsafe.Pointer(touchPoints_CArray)}
	var outptr_QTouchEvent *C.QTouchEvent = nil
	var outptr_QPointerEvent *C.QPointerEvent = nil
	var outptr_QInputEvent *C.QInputEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QTouchEvent_new6((C.int)(eventType), device.cPointer(), (C.int)(modifiers), (C.uint8_t)(touchPointStates), touchPoints_ma, &outptr_QTouchEvent, &outptr_QPointerEvent, &outptr_QInputEvent, &outptr_QEvent)
	ret := newQTouchEvent(outptr_QTouchEvent, outptr_QPointerEvent, outptr_QInputEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QTouchEvent) Clone() *QTouchEvent {
	return UnsafeNewQTouchEvent(unsafe.Pointer(C.QTouchEvent_Clone(this.h)), nil, nil, nil)
}

func (this *QTouchEvent) Target() *QObject {
	return UnsafeNewQObject(unsafe.Pointer(C.QTouchEvent_Target(this.h)))
}

func (this *QTouchEvent) TouchPointStates() QEventPoint__State {
	return (QEventPoint__State)(C.QTouchEvent_TouchPointStates(this.h))
}

func (this *QTouchEvent) TouchPoints() []QEventPoint {
	var _ma C.struct_miqt_array = C.QTouchEvent_TouchPoints(this.h)
	_ret := make([]QEventPoint, int(_ma.len))
	_outCast := (*[0xffff]*C.QEventPoint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_ret := _outCast[i]
		_lv_goptr := newQEventPoint(_lv_ret)
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTouchEvent) IsBeginEvent() bool {
	return (bool)(C.QTouchEvent_IsBeginEvent(this.h))
}

func (this *QTouchEvent) IsUpdateEvent() bool {
	return (bool)(C.QTouchEvent_IsUpdateEvent(this.h))
}

func (this *QTouchEvent) IsEndEvent() bool {
	return (bool)(C.QTouchEvent_IsEndEvent(this.h))
}

func (this *QTouchEvent) callVirtualBase_Clone() *QTouchEvent {

	return UnsafeNewQTouchEvent(unsafe.Pointer(C.QTouchEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil, nil, nil)
}
func (this *QTouchEvent) OnClone(slot func(super func() *QTouchEvent) *QTouchEvent) {
	C.QTouchEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_Clone
func miqt_exec_callback_QTouchEvent_Clone(self *C.QTouchEvent, cb C.intptr_t) *C.QTouchEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QTouchEvent) *QTouchEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QTouchEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(C.QTouchEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QTouchEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	C.QTouchEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_IsBeginEvent
func miqt_exec_callback_QTouchEvent_IsBeginEvent(self *C.QTouchEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (C.bool)(virtualReturn)

}

func (this *QTouchEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(C.QTouchEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QTouchEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	C.QTouchEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_IsUpdateEvent
func miqt_exec_callback_QTouchEvent_IsUpdateEvent(self *C.QTouchEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (C.bool)(virtualReturn)

}

func (this *QTouchEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(C.QTouchEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QTouchEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	C.QTouchEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_IsEndEvent
func miqt_exec_callback_QTouchEvent_IsEndEvent(self *C.QTouchEvent, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_IsEndEvent)

	return (C.bool)(virtualReturn)

}

func (this *QTouchEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	C.QTouchEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (C.ulonglong)(timestamp))

}
func (this *QTouchEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	C.QTouchEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_SetTimestamp
func miqt_exec_callback_QTouchEvent_SetTimestamp(self *C.QTouchEvent, cb C.intptr_t, timestamp C.ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QTouchEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

func (this *QTouchEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QTouchEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QTouchEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QTouchEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_SetAccepted
func miqt_exec_callback_QTouchEvent_SetAccepted(self *C.QTouchEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QTouchEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QTouchEvent) Delete() {
	C.QTouchEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QTouchEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QTouchEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QScrollPrepareEvent struct {
	h          *C.QScrollPrepareEvent
	isSubclass bool
	*QEvent
}

func (this *QScrollPrepareEvent) cPointer() *C.QScrollPrepareEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QScrollPrepareEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQScrollPrepareEvent constructs the type using only CGO pointers.
func newQScrollPrepareEvent(h *C.QScrollPrepareEvent, h_QEvent *C.QEvent) *QScrollPrepareEvent {
	if h == nil {
		return nil
	}
	return &QScrollPrepareEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQScrollPrepareEvent constructs the type using only unsafe pointers.
func UnsafeNewQScrollPrepareEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QScrollPrepareEvent {
	if h == nil {
		return nil
	}

	return &QScrollPrepareEvent{h: (*C.QScrollPrepareEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQScrollPrepareEvent constructs a new QScrollPrepareEvent object.
func NewQScrollPrepareEvent(startPos *QPointF) *QScrollPrepareEvent {
	var outptr_QScrollPrepareEvent *C.QScrollPrepareEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QScrollPrepareEvent_new(startPos.cPointer(), &outptr_QScrollPrepareEvent, &outptr_QEvent)
	ret := newQScrollPrepareEvent(outptr_QScrollPrepareEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QScrollPrepareEvent) Clone() *QScrollPrepareEvent {
	return UnsafeNewQScrollPrepareEvent(unsafe.Pointer(C.QScrollPrepareEvent_Clone(this.h)), nil)
}

func (this *QScrollPrepareEvent) StartPos() *QPointF {
	_ret := C.QScrollPrepareEvent_StartPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) ViewportSize() *QSizeF {
	_ret := C.QScrollPrepareEvent_ViewportSize(this.h)
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) ContentPosRange() *QRectF {
	_ret := C.QScrollPrepareEvent_ContentPosRange(this.h)
	_goptr := newQRectF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) ContentPos() *QPointF {
	_ret := C.QScrollPrepareEvent_ContentPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) SetViewportSize(size *QSizeF) {
	C.QScrollPrepareEvent_SetViewportSize(this.h, size.cPointer())
}

func (this *QScrollPrepareEvent) SetContentPosRange(rect *QRectF) {
	C.QScrollPrepareEvent_SetContentPosRange(this.h, rect.cPointer())
}

func (this *QScrollPrepareEvent) SetContentPos(pos *QPointF) {
	C.QScrollPrepareEvent_SetContentPos(this.h, pos.cPointer())
}

func (this *QScrollPrepareEvent) callVirtualBase_Clone() *QScrollPrepareEvent {

	return UnsafeNewQScrollPrepareEvent(unsafe.Pointer(C.QScrollPrepareEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QScrollPrepareEvent) OnClone(slot func(super func() *QScrollPrepareEvent) *QScrollPrepareEvent) {
	C.QScrollPrepareEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollPrepareEvent_Clone
func miqt_exec_callback_QScrollPrepareEvent_Clone(self *C.QScrollPrepareEvent, cb C.intptr_t) *C.QScrollPrepareEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QScrollPrepareEvent) *QScrollPrepareEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollPrepareEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QScrollPrepareEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QScrollPrepareEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QScrollPrepareEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QScrollPrepareEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollPrepareEvent_SetAccepted
func miqt_exec_callback_QScrollPrepareEvent_SetAccepted(self *C.QScrollPrepareEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QScrollPrepareEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QScrollPrepareEvent) Delete() {
	C.QScrollPrepareEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QScrollPrepareEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QScrollPrepareEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QScrollEvent struct {
	h          *C.QScrollEvent
	isSubclass bool
	*QEvent
}

func (this *QScrollEvent) cPointer() *C.QScrollEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QScrollEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQScrollEvent constructs the type using only CGO pointers.
func newQScrollEvent(h *C.QScrollEvent, h_QEvent *C.QEvent) *QScrollEvent {
	if h == nil {
		return nil
	}
	return &QScrollEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQScrollEvent constructs the type using only unsafe pointers.
func UnsafeNewQScrollEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QScrollEvent {
	if h == nil {
		return nil
	}

	return &QScrollEvent{h: (*C.QScrollEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQScrollEvent constructs a new QScrollEvent object.
func NewQScrollEvent(contentPos *QPointF, overshoot *QPointF, scrollState QScrollEvent__ScrollState) *QScrollEvent {
	var outptr_QScrollEvent *C.QScrollEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QScrollEvent_new(contentPos.cPointer(), overshoot.cPointer(), (C.int)(scrollState), &outptr_QScrollEvent, &outptr_QEvent)
	ret := newQScrollEvent(outptr_QScrollEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QScrollEvent) Clone() *QScrollEvent {
	return UnsafeNewQScrollEvent(unsafe.Pointer(C.QScrollEvent_Clone(this.h)), nil)
}

func (this *QScrollEvent) ContentPos() *QPointF {
	_ret := C.QScrollEvent_ContentPos(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollEvent) OvershootDistance() *QPointF {
	_ret := C.QScrollEvent_OvershootDistance(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollEvent) ScrollState() QScrollEvent__ScrollState {
	return (QScrollEvent__ScrollState)(C.QScrollEvent_ScrollState(this.h))
}

func (this *QScrollEvent) callVirtualBase_Clone() *QScrollEvent {

	return UnsafeNewQScrollEvent(unsafe.Pointer(C.QScrollEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QScrollEvent) OnClone(slot func(super func() *QScrollEvent) *QScrollEvent) {
	C.QScrollEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollEvent_Clone
func miqt_exec_callback_QScrollEvent_Clone(self *C.QScrollEvent, cb C.intptr_t) *C.QScrollEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QScrollEvent) *QScrollEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QScrollEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QScrollEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QScrollEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QScrollEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollEvent_SetAccepted
func miqt_exec_callback_QScrollEvent_SetAccepted(self *C.QScrollEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QScrollEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QScrollEvent) Delete() {
	C.QScrollEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QScrollEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QScrollEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QScreenOrientationChangeEvent struct {
	h          *C.QScreenOrientationChangeEvent
	isSubclass bool
	*QEvent
}

func (this *QScreenOrientationChangeEvent) cPointer() *C.QScreenOrientationChangeEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QScreenOrientationChangeEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQScreenOrientationChangeEvent constructs the type using only CGO pointers.
func newQScreenOrientationChangeEvent(h *C.QScreenOrientationChangeEvent, h_QEvent *C.QEvent) *QScreenOrientationChangeEvent {
	if h == nil {
		return nil
	}
	return &QScreenOrientationChangeEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQScreenOrientationChangeEvent constructs the type using only unsafe pointers.
func UnsafeNewQScreenOrientationChangeEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QScreenOrientationChangeEvent {
	if h == nil {
		return nil
	}

	return &QScreenOrientationChangeEvent{h: (*C.QScreenOrientationChangeEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQScreenOrientationChangeEvent constructs a new QScreenOrientationChangeEvent object.
func NewQScreenOrientationChangeEvent(screen *QScreen, orientation ScreenOrientation) *QScreenOrientationChangeEvent {
	var outptr_QScreenOrientationChangeEvent *C.QScreenOrientationChangeEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QScreenOrientationChangeEvent_new(screen.cPointer(), (C.int)(orientation), &outptr_QScreenOrientationChangeEvent, &outptr_QEvent)
	ret := newQScreenOrientationChangeEvent(outptr_QScreenOrientationChangeEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QScreenOrientationChangeEvent) Clone() *QScreenOrientationChangeEvent {
	return UnsafeNewQScreenOrientationChangeEvent(unsafe.Pointer(C.QScreenOrientationChangeEvent_Clone(this.h)), nil)
}

func (this *QScreenOrientationChangeEvent) Screen() *QScreen {
	return UnsafeNewQScreen(unsafe.Pointer(C.QScreenOrientationChangeEvent_Screen(this.h)), nil)
}

func (this *QScreenOrientationChangeEvent) Orientation() ScreenOrientation {
	return (ScreenOrientation)(C.QScreenOrientationChangeEvent_Orientation(this.h))
}

func (this *QScreenOrientationChangeEvent) callVirtualBase_Clone() *QScreenOrientationChangeEvent {

	return UnsafeNewQScreenOrientationChangeEvent(unsafe.Pointer(C.QScreenOrientationChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QScreenOrientationChangeEvent) OnClone(slot func(super func() *QScreenOrientationChangeEvent) *QScreenOrientationChangeEvent) {
	C.QScreenOrientationChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenOrientationChangeEvent_Clone
func miqt_exec_callback_QScreenOrientationChangeEvent_Clone(self *C.QScreenOrientationChangeEvent, cb C.intptr_t) *C.QScreenOrientationChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QScreenOrientationChangeEvent) *QScreenOrientationChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScreenOrientationChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QScreenOrientationChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QScreenOrientationChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QScreenOrientationChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QScreenOrientationChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenOrientationChangeEvent_SetAccepted
func miqt_exec_callback_QScreenOrientationChangeEvent_SetAccepted(self *C.QScreenOrientationChangeEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QScreenOrientationChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QScreenOrientationChangeEvent) Delete() {
	C.QScreenOrientationChangeEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QScreenOrientationChangeEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QScreenOrientationChangeEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QApplicationStateChangeEvent struct {
	h          *C.QApplicationStateChangeEvent
	isSubclass bool
	*QEvent
}

func (this *QApplicationStateChangeEvent) cPointer() *C.QApplicationStateChangeEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QApplicationStateChangeEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQApplicationStateChangeEvent constructs the type using only CGO pointers.
func newQApplicationStateChangeEvent(h *C.QApplicationStateChangeEvent, h_QEvent *C.QEvent) *QApplicationStateChangeEvent {
	if h == nil {
		return nil
	}
	return &QApplicationStateChangeEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQApplicationStateChangeEvent constructs the type using only unsafe pointers.
func UnsafeNewQApplicationStateChangeEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QApplicationStateChangeEvent {
	if h == nil {
		return nil
	}

	return &QApplicationStateChangeEvent{h: (*C.QApplicationStateChangeEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQApplicationStateChangeEvent constructs a new QApplicationStateChangeEvent object.
func NewQApplicationStateChangeEvent(state ApplicationState) *QApplicationStateChangeEvent {
	var outptr_QApplicationStateChangeEvent *C.QApplicationStateChangeEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QApplicationStateChangeEvent_new((C.int)(state), &outptr_QApplicationStateChangeEvent, &outptr_QEvent)
	ret := newQApplicationStateChangeEvent(outptr_QApplicationStateChangeEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QApplicationStateChangeEvent) Clone() *QApplicationStateChangeEvent {
	return UnsafeNewQApplicationStateChangeEvent(unsafe.Pointer(C.QApplicationStateChangeEvent_Clone(this.h)), nil)
}

func (this *QApplicationStateChangeEvent) ApplicationState() ApplicationState {
	return (ApplicationState)(C.QApplicationStateChangeEvent_ApplicationState(this.h))
}

func (this *QApplicationStateChangeEvent) callVirtualBase_Clone() *QApplicationStateChangeEvent {

	return UnsafeNewQApplicationStateChangeEvent(unsafe.Pointer(C.QApplicationStateChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h))), nil)
}
func (this *QApplicationStateChangeEvent) OnClone(slot func(super func() *QApplicationStateChangeEvent) *QApplicationStateChangeEvent) {
	C.QApplicationStateChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QApplicationStateChangeEvent_Clone
func miqt_exec_callback_QApplicationStateChangeEvent_Clone(self *C.QApplicationStateChangeEvent, cb C.intptr_t) *C.QApplicationStateChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QApplicationStateChangeEvent) *QApplicationStateChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QApplicationStateChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QApplicationStateChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QApplicationStateChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QApplicationStateChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	C.QApplicationStateChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QApplicationStateChangeEvent_SetAccepted
func miqt_exec_callback_QApplicationStateChangeEvent_SetAccepted(self *C.QApplicationStateChangeEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QApplicationStateChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

// Delete this object from C++ memory.
func (this *QApplicationStateChangeEvent) Delete() {
	C.QApplicationStateChangeEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QApplicationStateChangeEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QApplicationStateChangeEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QInputMethodEvent__Attribute struct {
	h          *C.QInputMethodEvent__Attribute
	isSubclass bool
}

func (this *QInputMethodEvent__Attribute) cPointer() *C.QInputMethodEvent__Attribute {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QInputMethodEvent__Attribute) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQInputMethodEvent__Attribute constructs the type using only CGO pointers.
func newQInputMethodEvent__Attribute(h *C.QInputMethodEvent__Attribute) *QInputMethodEvent__Attribute {
	if h == nil {
		return nil
	}
	return &QInputMethodEvent__Attribute{h: h}
}

// UnsafeNewQInputMethodEvent__Attribute constructs the type using only unsafe pointers.
func UnsafeNewQInputMethodEvent__Attribute(h unsafe.Pointer) *QInputMethodEvent__Attribute {
	if h == nil {
		return nil
	}

	return &QInputMethodEvent__Attribute{h: (*C.QInputMethodEvent__Attribute)(h)}
}

// NewQInputMethodEvent__Attribute constructs a new QInputMethodEvent::Attribute object.
func NewQInputMethodEvent__Attribute(typ QInputMethodEvent__AttributeType, s int, l int, val QVariant) *QInputMethodEvent__Attribute {
	var outptr_QInputMethodEvent__Attribute *C.QInputMethodEvent__Attribute = nil

	C.QInputMethodEvent__Attribute_new((C.int)(typ), (C.int)(s), (C.int)(l), val.cPointer(), &outptr_QInputMethodEvent__Attribute)
	ret := newQInputMethodEvent__Attribute(outptr_QInputMethodEvent__Attribute)
	ret.isSubclass = true
	return ret
}

// NewQInputMethodEvent__Attribute2 constructs a new QInputMethodEvent::Attribute object.
func NewQInputMethodEvent__Attribute2(typ QInputMethodEvent__AttributeType, s int, l int) *QInputMethodEvent__Attribute {
	var outptr_QInputMethodEvent__Attribute *C.QInputMethodEvent__Attribute = nil

	C.QInputMethodEvent__Attribute_new2((C.int)(typ), (C.int)(s), (C.int)(l), &outptr_QInputMethodEvent__Attribute)
	ret := newQInputMethodEvent__Attribute(outptr_QInputMethodEvent__Attribute)
	ret.isSubclass = true
	return ret
}

// NewQInputMethodEvent__Attribute3 constructs a new QInputMethodEvent::Attribute object.
func NewQInputMethodEvent__Attribute3(param1 *QInputMethodEvent__Attribute) *QInputMethodEvent__Attribute {
	var outptr_QInputMethodEvent__Attribute *C.QInputMethodEvent__Attribute = nil

	C.QInputMethodEvent__Attribute_new3(param1.cPointer(), &outptr_QInputMethodEvent__Attribute)
	ret := newQInputMethodEvent__Attribute(outptr_QInputMethodEvent__Attribute)
	ret.isSubclass = true
	return ret
}

func (this *QInputMethodEvent__Attribute) OperatorAssign(param1 *QInputMethodEvent__Attribute) {
	C.QInputMethodEvent__Attribute_OperatorAssign(this.h, param1.cPointer())
}

// Delete this object from C++ memory.
func (this *QInputMethodEvent__Attribute) Delete() {
	C.QInputMethodEvent__Attribute_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QInputMethodEvent__Attribute) GoGC() {
	runtime.SetFinalizer(this, func(this *QInputMethodEvent__Attribute) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
