package qt

/*

#include "gen_qabstractitemview.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QAbstractItemView__SelectionMode int

const (
	QAbstractItemView__NoSelection         QAbstractItemView__SelectionMode = 0
	QAbstractItemView__SingleSelection     QAbstractItemView__SelectionMode = 1
	QAbstractItemView__MultiSelection      QAbstractItemView__SelectionMode = 2
	QAbstractItemView__ExtendedSelection   QAbstractItemView__SelectionMode = 3
	QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = 4
)

type QAbstractItemView__SelectionBehavior int

const (
	QAbstractItemView__SelectItems   QAbstractItemView__SelectionBehavior = 0
	QAbstractItemView__SelectRows    QAbstractItemView__SelectionBehavior = 1
	QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = 2
)

type QAbstractItemView__ScrollHint int

const (
	QAbstractItemView__EnsureVisible    QAbstractItemView__ScrollHint = 0
	QAbstractItemView__PositionAtTop    QAbstractItemView__ScrollHint = 1
	QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = 2
	QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = 3
)

type QAbstractItemView__EditTrigger int

const (
	QAbstractItemView__NoEditTriggers  QAbstractItemView__EditTrigger = 0
	QAbstractItemView__CurrentChanged  QAbstractItemView__EditTrigger = 1
	QAbstractItemView__DoubleClicked   QAbstractItemView__EditTrigger = 2
	QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = 4
	QAbstractItemView__EditKeyPressed  QAbstractItemView__EditTrigger = 8
	QAbstractItemView__AnyKeyPressed   QAbstractItemView__EditTrigger = 16
	QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = 31
)

type QAbstractItemView__ScrollMode int

const (
	QAbstractItemView__ScrollPerItem  QAbstractItemView__ScrollMode = 0
	QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = 1
)

type QAbstractItemView__DragDropMode int

const (
	QAbstractItemView__NoDragDrop   QAbstractItemView__DragDropMode = 0
	QAbstractItemView__DragOnly     QAbstractItemView__DragDropMode = 1
	QAbstractItemView__DropOnly     QAbstractItemView__DragDropMode = 2
	QAbstractItemView__DragDrop     QAbstractItemView__DragDropMode = 3
	QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = 4
)

type QAbstractItemView__CursorAction int

const (
	QAbstractItemView__MoveUp       QAbstractItemView__CursorAction = 0
	QAbstractItemView__MoveDown     QAbstractItemView__CursorAction = 1
	QAbstractItemView__MoveLeft     QAbstractItemView__CursorAction = 2
	QAbstractItemView__MoveRight    QAbstractItemView__CursorAction = 3
	QAbstractItemView__MoveHome     QAbstractItemView__CursorAction = 4
	QAbstractItemView__MoveEnd      QAbstractItemView__CursorAction = 5
	QAbstractItemView__MovePageUp   QAbstractItemView__CursorAction = 6
	QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = 7
	QAbstractItemView__MoveNext     QAbstractItemView__CursorAction = 8
	QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = 9
)

type QAbstractItemView__State int

const (
	QAbstractItemView__NoState            QAbstractItemView__State = 0
	QAbstractItemView__DraggingState      QAbstractItemView__State = 1
	QAbstractItemView__DragSelectingState QAbstractItemView__State = 2
	QAbstractItemView__EditingState       QAbstractItemView__State = 3
	QAbstractItemView__ExpandingState     QAbstractItemView__State = 4
	QAbstractItemView__CollapsingState    QAbstractItemView__State = 5
	QAbstractItemView__AnimatingState     QAbstractItemView__State = 6
)

type QAbstractItemView__DropIndicatorPosition int

const (
	QAbstractItemView__OnItem     QAbstractItemView__DropIndicatorPosition = 0
	QAbstractItemView__AboveItem  QAbstractItemView__DropIndicatorPosition = 1
	QAbstractItemView__BelowItem  QAbstractItemView__DropIndicatorPosition = 2
	QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = 3
)

type QAbstractItemView struct {
	h          *C.QAbstractItemView
	isSubclass bool
	*QAbstractScrollArea
}

func (this *QAbstractItemView) cPointer() *C.QAbstractItemView {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QAbstractItemView) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQAbstractItemView constructs the type using only CGO pointers.
func newQAbstractItemView(h *C.QAbstractItemView, h_QAbstractScrollArea *C.QAbstractScrollArea, h_QFrame *C.QFrame, h_QWidget *C.QWidget, h_QObject *C.QObject, h_QPaintDevice *C.QPaintDevice) *QAbstractItemView {
	if h == nil {
		return nil
	}
	return &QAbstractItemView{h: h,
		QAbstractScrollArea: newQAbstractScrollArea(h_QAbstractScrollArea, h_QFrame, h_QWidget, h_QObject, h_QPaintDevice)}
}

// UnsafeNewQAbstractItemView constructs the type using only unsafe pointers.
func UnsafeNewQAbstractItemView(h unsafe.Pointer, h_QAbstractScrollArea unsafe.Pointer, h_QFrame unsafe.Pointer, h_QWidget unsafe.Pointer, h_QObject unsafe.Pointer, h_QPaintDevice unsafe.Pointer) *QAbstractItemView {
	if h == nil {
		return nil
	}

	return &QAbstractItemView{h: (*C.QAbstractItemView)(h),
		QAbstractScrollArea: UnsafeNewQAbstractScrollArea(h_QAbstractScrollArea, h_QFrame, h_QWidget, h_QObject, h_QPaintDevice)}
}

func (this *QAbstractItemView) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QAbstractItemView_MetaObject(this.h)))
}

func (this *QAbstractItemView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QAbstractItemView_Metacast(this.h, param1_Cstring))
}

func QAbstractItemView_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractItemView_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractItemView_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractItemView_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractItemView) SetModel(model *QAbstractItemModel) {
	C.QAbstractItemView_SetModel(this.h, model.cPointer())
}

func (this *QAbstractItemView) Model() *QAbstractItemModel {
	return UnsafeNewQAbstractItemModel(unsafe.Pointer(C.QAbstractItemView_Model(this.h)), nil)
}

func (this *QAbstractItemView) SetSelectionModel(selectionModel *QItemSelectionModel) {
	C.QAbstractItemView_SetSelectionModel(this.h, selectionModel.cPointer())
}

func (this *QAbstractItemView) SelectionModel() *QItemSelectionModel {
	return UnsafeNewQItemSelectionModel(unsafe.Pointer(C.QAbstractItemView_SelectionModel(this.h)), nil)
}

func (this *QAbstractItemView) SetItemDelegate(delegate *QAbstractItemDelegate) {
	C.QAbstractItemView_SetItemDelegate(this.h, delegate.cPointer())
}

func (this *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate {
	return UnsafeNewQAbstractItemDelegate(unsafe.Pointer(C.QAbstractItemView_ItemDelegate(this.h)), nil)
}

func (this *QAbstractItemView) SetSelectionMode(mode QAbstractItemView__SelectionMode) {
	C.QAbstractItemView_SetSelectionMode(this.h, (C.int)(mode))
}

func (this *QAbstractItemView) SelectionMode() QAbstractItemView__SelectionMode {
	return (QAbstractItemView__SelectionMode)(C.QAbstractItemView_SelectionMode(this.h))
}

func (this *QAbstractItemView) SetSelectionBehavior(behavior QAbstractItemView__SelectionBehavior) {
	C.QAbstractItemView_SetSelectionBehavior(this.h, (C.int)(behavior))
}

func (this *QAbstractItemView) SelectionBehavior() QAbstractItemView__SelectionBehavior {
	return (QAbstractItemView__SelectionBehavior)(C.QAbstractItemView_SelectionBehavior(this.h))
}

func (this *QAbstractItemView) CurrentIndex() *QModelIndex {
	_ret := C.QAbstractItemView_CurrentIndex(this.h)
	_goptr := newQModelIndex(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) RootIndex() *QModelIndex {
	_ret := C.QAbstractItemView_RootIndex(this.h)
	_goptr := newQModelIndex(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SetEditTriggers(triggers QAbstractItemView__EditTrigger) {
	C.QAbstractItemView_SetEditTriggers(this.h, (C.int)(triggers))
}

func (this *QAbstractItemView) EditTriggers() QAbstractItemView__EditTrigger {
	return (QAbstractItemView__EditTrigger)(C.QAbstractItemView_EditTriggers(this.h))
}

func (this *QAbstractItemView) SetVerticalScrollMode(mode QAbstractItemView__ScrollMode) {
	C.QAbstractItemView_SetVerticalScrollMode(this.h, (C.int)(mode))
}

func (this *QAbstractItemView) VerticalScrollMode() QAbstractItemView__ScrollMode {
	return (QAbstractItemView__ScrollMode)(C.QAbstractItemView_VerticalScrollMode(this.h))
}

func (this *QAbstractItemView) ResetVerticalScrollMode() {
	C.QAbstractItemView_ResetVerticalScrollMode(this.h)
}

func (this *QAbstractItemView) SetHorizontalScrollMode(mode QAbstractItemView__ScrollMode) {
	C.QAbstractItemView_SetHorizontalScrollMode(this.h, (C.int)(mode))
}

func (this *QAbstractItemView) HorizontalScrollMode() QAbstractItemView__ScrollMode {
	return (QAbstractItemView__ScrollMode)(C.QAbstractItemView_HorizontalScrollMode(this.h))
}

func (this *QAbstractItemView) ResetHorizontalScrollMode() {
	C.QAbstractItemView_ResetHorizontalScrollMode(this.h)
}

func (this *QAbstractItemView) SetAutoScroll(enable bool) {
	C.QAbstractItemView_SetAutoScroll(this.h, (C.bool)(enable))
}

func (this *QAbstractItemView) HasAutoScroll() bool {
	return (bool)(C.QAbstractItemView_HasAutoScroll(this.h))
}

func (this *QAbstractItemView) SetAutoScrollMargin(margin int) {
	C.QAbstractItemView_SetAutoScrollMargin(this.h, (C.int)(margin))
}

func (this *QAbstractItemView) AutoScrollMargin() int {
	return (int)(C.QAbstractItemView_AutoScrollMargin(this.h))
}

func (this *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	C.QAbstractItemView_SetTabKeyNavigation(this.h, (C.bool)(enable))
}

func (this *QAbstractItemView) TabKeyNavigation() bool {
	return (bool)(C.QAbstractItemView_TabKeyNavigation(this.h))
}

func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	C.QAbstractItemView_SetDropIndicatorShown(this.h, (C.bool)(enable))
}

func (this *QAbstractItemView) ShowDropIndicator() bool {
	return (bool)(C.QAbstractItemView_ShowDropIndicator(this.h))
}

func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	C.QAbstractItemView_SetDragEnabled(this.h, (C.bool)(enable))
}

func (this *QAbstractItemView) DragEnabled() bool {
	return (bool)(C.QAbstractItemView_DragEnabled(this.h))
}

func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	C.QAbstractItemView_SetDragDropOverwriteMode(this.h, (C.bool)(overwrite))
}

func (this *QAbstractItemView) DragDropOverwriteMode() bool {
	return (bool)(C.QAbstractItemView_DragDropOverwriteMode(this.h))
}

func (this *QAbstractItemView) SetDragDropMode(behavior QAbstractItemView__DragDropMode) {
	C.QAbstractItemView_SetDragDropMode(this.h, (C.int)(behavior))
}

func (this *QAbstractItemView) DragDropMode() QAbstractItemView__DragDropMode {
	return (QAbstractItemView__DragDropMode)(C.QAbstractItemView_DragDropMode(this.h))
}

func (this *QAbstractItemView) SetDefaultDropAction(dropAction DropAction) {
	C.QAbstractItemView_SetDefaultDropAction(this.h, (C.int)(dropAction))
}

func (this *QAbstractItemView) DefaultDropAction() DropAction {
	return (DropAction)(C.QAbstractItemView_DefaultDropAction(this.h))
}

func (this *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	C.QAbstractItemView_SetAlternatingRowColors(this.h, (C.bool)(enable))
}

func (this *QAbstractItemView) AlternatingRowColors() bool {
	return (bool)(C.QAbstractItemView_AlternatingRowColors(this.h))
}

func (this *QAbstractItemView) SetIconSize(size *QSize) {
	C.QAbstractItemView_SetIconSize(this.h, size.cPointer())
}

func (this *QAbstractItemView) IconSize() *QSize {
	_ret := C.QAbstractItemView_IconSize(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SetTextElideMode(mode TextElideMode) {
	C.QAbstractItemView_SetTextElideMode(this.h, (C.int)(mode))
}

func (this *QAbstractItemView) TextElideMode() TextElideMode {
	return (TextElideMode)(C.QAbstractItemView_TextElideMode(this.h))
}

func (this *QAbstractItemView) KeyboardSearch(search string) {
	search_ms := C.struct_miqt_string{}
	search_ms.data = C.CString(search)
	search_ms.len = C.size_t(len(search))
	defer C.free(unsafe.Pointer(search_ms.data))
	C.QAbstractItemView_KeyboardSearch(this.h, search_ms)
}

func (this *QAbstractItemView) VisualRect(index *QModelIndex) *QRect {
	_ret := C.QAbstractItemView_VisualRect(this.h, index.cPointer())
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) ScrollTo(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	C.QAbstractItemView_ScrollTo(this.h, index.cPointer(), (C.int)(hint))
}

func (this *QAbstractItemView) IndexAt(point *QPoint) *QModelIndex {
	_ret := C.QAbstractItemView_IndexAt(this.h, point.cPointer())
	_goptr := newQModelIndex(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SizeHintForIndex(index *QModelIndex) *QSize {
	_ret := C.QAbstractItemView_SizeHintForIndex(this.h, index.cPointer())
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SizeHintForRow(row int) int {
	return (int)(C.QAbstractItemView_SizeHintForRow(this.h, (C.int)(row)))
}

func (this *QAbstractItemView) SizeHintForColumn(column int) int {
	return (int)(C.QAbstractItemView_SizeHintForColumn(this.h, (C.int)(column)))
}

func (this *QAbstractItemView) OpenPersistentEditor(index *QModelIndex) {
	C.QAbstractItemView_OpenPersistentEditor(this.h, index.cPointer())
}

func (this *QAbstractItemView) ClosePersistentEditor(index *QModelIndex) {
	C.QAbstractItemView_ClosePersistentEditor(this.h, index.cPointer())
}

func (this *QAbstractItemView) IsPersistentEditorOpen(index *QModelIndex) bool {
	return (bool)(C.QAbstractItemView_IsPersistentEditorOpen(this.h, index.cPointer()))
}

func (this *QAbstractItemView) SetIndexWidget(index *QModelIndex, widget *QWidget) {
	C.QAbstractItemView_SetIndexWidget(this.h, index.cPointer(), widget.cPointer())
}

func (this *QAbstractItemView) IndexWidget(index *QModelIndex) *QWidget {
	return UnsafeNewQWidget(unsafe.Pointer(C.QAbstractItemView_IndexWidget(this.h, index.cPointer())), nil, nil)
}

func (this *QAbstractItemView) SetItemDelegateForRow(row int, delegate *QAbstractItemDelegate) {
	C.QAbstractItemView_SetItemDelegateForRow(this.h, (C.int)(row), delegate.cPointer())
}

func (this *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate {
	return UnsafeNewQAbstractItemDelegate(unsafe.Pointer(C.QAbstractItemView_ItemDelegateForRow(this.h, (C.int)(row))), nil)
}

func (this *QAbstractItemView) SetItemDelegateForColumn(column int, delegate *QAbstractItemDelegate) {
	C.QAbstractItemView_SetItemDelegateForColumn(this.h, (C.int)(column), delegate.cPointer())
}

func (this *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate {
	return UnsafeNewQAbstractItemDelegate(unsafe.Pointer(C.QAbstractItemView_ItemDelegateForColumn(this.h, (C.int)(column))), nil)
}

func (this *QAbstractItemView) ItemDelegateWithIndex(index *QModelIndex) *QAbstractItemDelegate {
	return UnsafeNewQAbstractItemDelegate(unsafe.Pointer(C.QAbstractItemView_ItemDelegateWithIndex(this.h, index.cPointer())), nil)
}

func (this *QAbstractItemView) InputMethodQuery(query InputMethodQuery) *QVariant {
	_ret := C.QAbstractItemView_InputMethodQuery(this.h, (C.int)(query))
	_goptr := newQVariant(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) Reset() {
	C.QAbstractItemView_Reset(this.h)
}

func (this *QAbstractItemView) SetRootIndex(index *QModelIndex) {
	C.QAbstractItemView_SetRootIndex(this.h, index.cPointer())
}

func (this *QAbstractItemView) DoItemsLayout() {
	C.QAbstractItemView_DoItemsLayout(this.h)
}

func (this *QAbstractItemView) SelectAll() {
	C.QAbstractItemView_SelectAll(this.h)
}

func (this *QAbstractItemView) Edit(index *QModelIndex) {
	C.QAbstractItemView_Edit(this.h, index.cPointer())
}

func (this *QAbstractItemView) ClearSelection() {
	C.QAbstractItemView_ClearSelection(this.h)
}

func (this *QAbstractItemView) SetCurrentIndex(index *QModelIndex) {
	C.QAbstractItemView_SetCurrentIndex(this.h, index.cPointer())
}

func (this *QAbstractItemView) ScrollToTop() {
	C.QAbstractItemView_ScrollToTop(this.h)
}

func (this *QAbstractItemView) ScrollToBottom() {
	C.QAbstractItemView_ScrollToBottom(this.h)
}

func (this *QAbstractItemView) Update(index *QModelIndex) {
	C.QAbstractItemView_Update(this.h, index.cPointer())
}

func (this *QAbstractItemView) Pressed(index *QModelIndex) {
	C.QAbstractItemView_Pressed(this.h, index.cPointer())
}
func (this *QAbstractItemView) OnPressed(slot func(index *QModelIndex)) {
	C.QAbstractItemView_connect_Pressed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Pressed
func miqt_exec_callback_QAbstractItemView_Pressed(cb C.intptr_t, index *C.QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQModelIndex(unsafe.Pointer(index))

	gofunc(slotval1)
}

func (this *QAbstractItemView) Clicked(index *QModelIndex) {
	C.QAbstractItemView_Clicked(this.h, index.cPointer())
}
func (this *QAbstractItemView) OnClicked(slot func(index *QModelIndex)) {
	C.QAbstractItemView_connect_Clicked(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Clicked
func miqt_exec_callback_QAbstractItemView_Clicked(cb C.intptr_t, index *C.QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQModelIndex(unsafe.Pointer(index))

	gofunc(slotval1)
}

func (this *QAbstractItemView) DoubleClicked(index *QModelIndex) {
	C.QAbstractItemView_DoubleClicked(this.h, index.cPointer())
}
func (this *QAbstractItemView) OnDoubleClicked(slot func(index *QModelIndex)) {
	C.QAbstractItemView_connect_DoubleClicked(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DoubleClicked
func miqt_exec_callback_QAbstractItemView_DoubleClicked(cb C.intptr_t, index *C.QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQModelIndex(unsafe.Pointer(index))

	gofunc(slotval1)
}

func (this *QAbstractItemView) Activated(index *QModelIndex) {
	C.QAbstractItemView_Activated(this.h, index.cPointer())
}
func (this *QAbstractItemView) OnActivated(slot func(index *QModelIndex)) {
	C.QAbstractItemView_connect_Activated(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Activated
func miqt_exec_callback_QAbstractItemView_Activated(cb C.intptr_t, index *C.QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQModelIndex(unsafe.Pointer(index))

	gofunc(slotval1)
}

func (this *QAbstractItemView) Entered(index *QModelIndex) {
	C.QAbstractItemView_Entered(this.h, index.cPointer())
}
func (this *QAbstractItemView) OnEntered(slot func(index *QModelIndex)) {
	C.QAbstractItemView_connect_Entered(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Entered
func miqt_exec_callback_QAbstractItemView_Entered(cb C.intptr_t, index *C.QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQModelIndex(unsafe.Pointer(index))

	gofunc(slotval1)
}

func (this *QAbstractItemView) ViewportEntered() {
	C.QAbstractItemView_ViewportEntered(this.h)
}
func (this *QAbstractItemView) OnViewportEntered(slot func()) {
	C.QAbstractItemView_connect_ViewportEntered(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ViewportEntered
func miqt_exec_callback_QAbstractItemView_ViewportEntered(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractItemView) IconSizeChanged(size *QSize) {
	C.QAbstractItemView_IconSizeChanged(this.h, size.cPointer())
}
func (this *QAbstractItemView) OnIconSizeChanged(slot func(size *QSize)) {
	C.QAbstractItemView_connect_IconSizeChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_IconSizeChanged
func miqt_exec_callback_QAbstractItemView_IconSizeChanged(cb C.intptr_t, size *C.QSize) {
	gofunc, ok := cgo.Handle(cb).Value().(func(size *QSize))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQSize(unsafe.Pointer(size))

	gofunc(slotval1)
}

func QAbstractItemView_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractItemView_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractItemView_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractItemView_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractItemView_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractItemView_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractItemView_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractItemView_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QAbstractItemView) Delete() {
	C.QAbstractItemView_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAbstractItemView) GoGC() {
	runtime.SetFinalizer(this, func(this *QAbstractItemView) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
