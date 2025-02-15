package qt

/*

#include "gen_qpagedpaintdevice.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QPagedPaintDevice__PageSize int

const (
	QPagedPaintDevice__A4                 QPagedPaintDevice__PageSize = 0
	QPagedPaintDevice__B5                 QPagedPaintDevice__PageSize = 1
	QPagedPaintDevice__Letter             QPagedPaintDevice__PageSize = 2
	QPagedPaintDevice__Legal              QPagedPaintDevice__PageSize = 3
	QPagedPaintDevice__Executive          QPagedPaintDevice__PageSize = 4
	QPagedPaintDevice__A0                 QPagedPaintDevice__PageSize = 5
	QPagedPaintDevice__A1                 QPagedPaintDevice__PageSize = 6
	QPagedPaintDevice__A2                 QPagedPaintDevice__PageSize = 7
	QPagedPaintDevice__A3                 QPagedPaintDevice__PageSize = 8
	QPagedPaintDevice__A5                 QPagedPaintDevice__PageSize = 9
	QPagedPaintDevice__A6                 QPagedPaintDevice__PageSize = 10
	QPagedPaintDevice__A7                 QPagedPaintDevice__PageSize = 11
	QPagedPaintDevice__A8                 QPagedPaintDevice__PageSize = 12
	QPagedPaintDevice__A9                 QPagedPaintDevice__PageSize = 13
	QPagedPaintDevice__B0                 QPagedPaintDevice__PageSize = 14
	QPagedPaintDevice__B1                 QPagedPaintDevice__PageSize = 15
	QPagedPaintDevice__B10                QPagedPaintDevice__PageSize = 16
	QPagedPaintDevice__B2                 QPagedPaintDevice__PageSize = 17
	QPagedPaintDevice__B3                 QPagedPaintDevice__PageSize = 18
	QPagedPaintDevice__B4                 QPagedPaintDevice__PageSize = 19
	QPagedPaintDevice__B6                 QPagedPaintDevice__PageSize = 20
	QPagedPaintDevice__B7                 QPagedPaintDevice__PageSize = 21
	QPagedPaintDevice__B8                 QPagedPaintDevice__PageSize = 22
	QPagedPaintDevice__B9                 QPagedPaintDevice__PageSize = 23
	QPagedPaintDevice__C5E                QPagedPaintDevice__PageSize = 24
	QPagedPaintDevice__Comm10E            QPagedPaintDevice__PageSize = 25
	QPagedPaintDevice__DLE                QPagedPaintDevice__PageSize = 26
	QPagedPaintDevice__Folio              QPagedPaintDevice__PageSize = 27
	QPagedPaintDevice__Ledger             QPagedPaintDevice__PageSize = 28
	QPagedPaintDevice__Tabloid            QPagedPaintDevice__PageSize = 29
	QPagedPaintDevice__Custom             QPagedPaintDevice__PageSize = 30
	QPagedPaintDevice__A10                QPagedPaintDevice__PageSize = 31
	QPagedPaintDevice__A3Extra            QPagedPaintDevice__PageSize = 32
	QPagedPaintDevice__A4Extra            QPagedPaintDevice__PageSize = 33
	QPagedPaintDevice__A4Plus             QPagedPaintDevice__PageSize = 34
	QPagedPaintDevice__A4Small            QPagedPaintDevice__PageSize = 35
	QPagedPaintDevice__A5Extra            QPagedPaintDevice__PageSize = 36
	QPagedPaintDevice__B5Extra            QPagedPaintDevice__PageSize = 37
	QPagedPaintDevice__JisB0              QPagedPaintDevice__PageSize = 38
	QPagedPaintDevice__JisB1              QPagedPaintDevice__PageSize = 39
	QPagedPaintDevice__JisB2              QPagedPaintDevice__PageSize = 40
	QPagedPaintDevice__JisB3              QPagedPaintDevice__PageSize = 41
	QPagedPaintDevice__JisB4              QPagedPaintDevice__PageSize = 42
	QPagedPaintDevice__JisB5              QPagedPaintDevice__PageSize = 43
	QPagedPaintDevice__JisB6              QPagedPaintDevice__PageSize = 44
	QPagedPaintDevice__JisB7              QPagedPaintDevice__PageSize = 45
	QPagedPaintDevice__JisB8              QPagedPaintDevice__PageSize = 46
	QPagedPaintDevice__JisB9              QPagedPaintDevice__PageSize = 47
	QPagedPaintDevice__JisB10             QPagedPaintDevice__PageSize = 48
	QPagedPaintDevice__AnsiC              QPagedPaintDevice__PageSize = 49
	QPagedPaintDevice__AnsiD              QPagedPaintDevice__PageSize = 50
	QPagedPaintDevice__AnsiE              QPagedPaintDevice__PageSize = 51
	QPagedPaintDevice__LegalExtra         QPagedPaintDevice__PageSize = 52
	QPagedPaintDevice__LetterExtra        QPagedPaintDevice__PageSize = 53
	QPagedPaintDevice__LetterPlus         QPagedPaintDevice__PageSize = 54
	QPagedPaintDevice__LetterSmall        QPagedPaintDevice__PageSize = 55
	QPagedPaintDevice__TabloidExtra       QPagedPaintDevice__PageSize = 56
	QPagedPaintDevice__ArchA              QPagedPaintDevice__PageSize = 57
	QPagedPaintDevice__ArchB              QPagedPaintDevice__PageSize = 58
	QPagedPaintDevice__ArchC              QPagedPaintDevice__PageSize = 59
	QPagedPaintDevice__ArchD              QPagedPaintDevice__PageSize = 60
	QPagedPaintDevice__ArchE              QPagedPaintDevice__PageSize = 61
	QPagedPaintDevice__Imperial7x9        QPagedPaintDevice__PageSize = 62
	QPagedPaintDevice__Imperial8x10       QPagedPaintDevice__PageSize = 63
	QPagedPaintDevice__Imperial9x11       QPagedPaintDevice__PageSize = 64
	QPagedPaintDevice__Imperial9x12       QPagedPaintDevice__PageSize = 65
	QPagedPaintDevice__Imperial10x11      QPagedPaintDevice__PageSize = 66
	QPagedPaintDevice__Imperial10x13      QPagedPaintDevice__PageSize = 67
	QPagedPaintDevice__Imperial10x14      QPagedPaintDevice__PageSize = 68
	QPagedPaintDevice__Imperial12x11      QPagedPaintDevice__PageSize = 69
	QPagedPaintDevice__Imperial15x11      QPagedPaintDevice__PageSize = 70
	QPagedPaintDevice__ExecutiveStandard  QPagedPaintDevice__PageSize = 71
	QPagedPaintDevice__Note               QPagedPaintDevice__PageSize = 72
	QPagedPaintDevice__Quarto             QPagedPaintDevice__PageSize = 73
	QPagedPaintDevice__Statement          QPagedPaintDevice__PageSize = 74
	QPagedPaintDevice__SuperA             QPagedPaintDevice__PageSize = 75
	QPagedPaintDevice__SuperB             QPagedPaintDevice__PageSize = 76
	QPagedPaintDevice__Postcard           QPagedPaintDevice__PageSize = 77
	QPagedPaintDevice__DoublePostcard     QPagedPaintDevice__PageSize = 78
	QPagedPaintDevice__Prc16K             QPagedPaintDevice__PageSize = 79
	QPagedPaintDevice__Prc32K             QPagedPaintDevice__PageSize = 80
	QPagedPaintDevice__Prc32KBig          QPagedPaintDevice__PageSize = 81
	QPagedPaintDevice__FanFoldUS          QPagedPaintDevice__PageSize = 82
	QPagedPaintDevice__FanFoldGerman      QPagedPaintDevice__PageSize = 83
	QPagedPaintDevice__FanFoldGermanLegal QPagedPaintDevice__PageSize = 84
	QPagedPaintDevice__EnvelopeB4         QPagedPaintDevice__PageSize = 85
	QPagedPaintDevice__EnvelopeB5         QPagedPaintDevice__PageSize = 86
	QPagedPaintDevice__EnvelopeB6         QPagedPaintDevice__PageSize = 87
	QPagedPaintDevice__EnvelopeC0         QPagedPaintDevice__PageSize = 88
	QPagedPaintDevice__EnvelopeC1         QPagedPaintDevice__PageSize = 89
	QPagedPaintDevice__EnvelopeC2         QPagedPaintDevice__PageSize = 90
	QPagedPaintDevice__EnvelopeC3         QPagedPaintDevice__PageSize = 91
	QPagedPaintDevice__EnvelopeC4         QPagedPaintDevice__PageSize = 92
	QPagedPaintDevice__EnvelopeC6         QPagedPaintDevice__PageSize = 93
	QPagedPaintDevice__EnvelopeC65        QPagedPaintDevice__PageSize = 94
	QPagedPaintDevice__EnvelopeC7         QPagedPaintDevice__PageSize = 95
	QPagedPaintDevice__Envelope9          QPagedPaintDevice__PageSize = 96
	QPagedPaintDevice__Envelope11         QPagedPaintDevice__PageSize = 97
	QPagedPaintDevice__Envelope12         QPagedPaintDevice__PageSize = 98
	QPagedPaintDevice__Envelope14         QPagedPaintDevice__PageSize = 99
	QPagedPaintDevice__EnvelopeMonarch    QPagedPaintDevice__PageSize = 100
	QPagedPaintDevice__EnvelopePersonal   QPagedPaintDevice__PageSize = 101
	QPagedPaintDevice__EnvelopeChou3      QPagedPaintDevice__PageSize = 102
	QPagedPaintDevice__EnvelopeChou4      QPagedPaintDevice__PageSize = 103
	QPagedPaintDevice__EnvelopeInvite     QPagedPaintDevice__PageSize = 104
	QPagedPaintDevice__EnvelopeItalian    QPagedPaintDevice__PageSize = 105
	QPagedPaintDevice__EnvelopeKaku2      QPagedPaintDevice__PageSize = 106
	QPagedPaintDevice__EnvelopeKaku3      QPagedPaintDevice__PageSize = 107
	QPagedPaintDevice__EnvelopePrc1       QPagedPaintDevice__PageSize = 108
	QPagedPaintDevice__EnvelopePrc2       QPagedPaintDevice__PageSize = 109
	QPagedPaintDevice__EnvelopePrc3       QPagedPaintDevice__PageSize = 110
	QPagedPaintDevice__EnvelopePrc4       QPagedPaintDevice__PageSize = 111
	QPagedPaintDevice__EnvelopePrc5       QPagedPaintDevice__PageSize = 112
	QPagedPaintDevice__EnvelopePrc6       QPagedPaintDevice__PageSize = 113
	QPagedPaintDevice__EnvelopePrc7       QPagedPaintDevice__PageSize = 114
	QPagedPaintDevice__EnvelopePrc8       QPagedPaintDevice__PageSize = 115
	QPagedPaintDevice__EnvelopePrc9       QPagedPaintDevice__PageSize = 116
	QPagedPaintDevice__EnvelopePrc10      QPagedPaintDevice__PageSize = 117
	QPagedPaintDevice__EnvelopeYou4       QPagedPaintDevice__PageSize = 118
	QPagedPaintDevice__LastPageSize       QPagedPaintDevice__PageSize = 118
	QPagedPaintDevice__NPageSize          QPagedPaintDevice__PageSize = 118
	QPagedPaintDevice__NPaperSize         QPagedPaintDevice__PageSize = 118
	QPagedPaintDevice__AnsiA              QPagedPaintDevice__PageSize = 2
	QPagedPaintDevice__AnsiB              QPagedPaintDevice__PageSize = 28
	QPagedPaintDevice__EnvelopeC5         QPagedPaintDevice__PageSize = 24
	QPagedPaintDevice__EnvelopeDL         QPagedPaintDevice__PageSize = 26
	QPagedPaintDevice__Envelope10         QPagedPaintDevice__PageSize = 25
)

type QPagedPaintDevice__PdfVersion int

const (
	QPagedPaintDevice__PdfVersion_1_4 QPagedPaintDevice__PdfVersion = 0
	QPagedPaintDevice__PdfVersion_A1b QPagedPaintDevice__PdfVersion = 1
	QPagedPaintDevice__PdfVersion_1_6 QPagedPaintDevice__PdfVersion = 2
)

type QPagedPaintDevice struct {
	h          *C.QPagedPaintDevice
	isSubclass bool
	*QPaintDevice
}

func (this *QPagedPaintDevice) cPointer() *C.QPagedPaintDevice {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPagedPaintDevice) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPagedPaintDevice constructs the type using only CGO pointers.
func newQPagedPaintDevice(h *C.QPagedPaintDevice, h_QPaintDevice *C.QPaintDevice) *QPagedPaintDevice {
	if h == nil {
		return nil
	}
	return &QPagedPaintDevice{h: h,
		QPaintDevice: newQPaintDevice(h_QPaintDevice)}
}

// UnsafeNewQPagedPaintDevice constructs the type using only unsafe pointers.
func UnsafeNewQPagedPaintDevice(h unsafe.Pointer, h_QPaintDevice unsafe.Pointer) *QPagedPaintDevice {
	if h == nil {
		return nil
	}

	return &QPagedPaintDevice{h: (*C.QPagedPaintDevice)(h),
		QPaintDevice: UnsafeNewQPaintDevice(h_QPaintDevice)}
}

func (this *QPagedPaintDevice) NewPage() bool {
	return (bool)(C.QPagedPaintDevice_NewPage(this.h))
}

func (this *QPagedPaintDevice) SetPageLayout(pageLayout *QPageLayout) bool {
	return (bool)(C.QPagedPaintDevice_SetPageLayout(this.h, pageLayout.cPointer()))
}

func (this *QPagedPaintDevice) SetPageSize(pageSize *QPageSize) bool {
	return (bool)(C.QPagedPaintDevice_SetPageSize(this.h, pageSize.cPointer()))
}

func (this *QPagedPaintDevice) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	return (bool)(C.QPagedPaintDevice_SetPageOrientation(this.h, (C.int)(orientation)))
}

func (this *QPagedPaintDevice) SetPageMargins(margins *QMarginsF) bool {
	return (bool)(C.QPagedPaintDevice_SetPageMargins(this.h, margins.cPointer()))
}

func (this *QPagedPaintDevice) SetPageMargins2(margins *QMarginsF, units QPageLayout__Unit) bool {
	return (bool)(C.QPagedPaintDevice_SetPageMargins2(this.h, margins.cPointer(), (C.int)(units)))
}

func (this *QPagedPaintDevice) PageLayout() *QPageLayout {
	_ret := C.QPagedPaintDevice_PageLayout(this.h)
	_goptr := newQPageLayout(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPagedPaintDevice) SetPageSizeWithSize(size QPagedPaintDevice__PageSize) {
	C.QPagedPaintDevice_SetPageSizeWithSize(this.h, (C.int)(size))
}

func (this *QPagedPaintDevice) PageSize() QPagedPaintDevice__PageSize {
	return (QPagedPaintDevice__PageSize)(C.QPagedPaintDevice_PageSize(this.h))
}

func (this *QPagedPaintDevice) SetPageSizeMM(size *QSizeF) {
	C.QPagedPaintDevice_SetPageSizeMM(this.h, size.cPointer())
}

func (this *QPagedPaintDevice) PageSizeMM() *QSizeF {
	_ret := C.QPagedPaintDevice_PageSizeMM(this.h)
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPagedPaintDevice) SetMargins(margins *QPagedPaintDevice__Margins) {
	C.QPagedPaintDevice_SetMargins(this.h, margins.cPointer())
}

func (this *QPagedPaintDevice) Margins() *QPagedPaintDevice__Margins {
	_ret := C.QPagedPaintDevice_Margins(this.h)
	_goptr := newQPagedPaintDevice__Margins(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QPagedPaintDevice) Delete() {
	C.QPagedPaintDevice_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPagedPaintDevice) GoGC() {
	runtime.SetFinalizer(this, func(this *QPagedPaintDevice) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QPagedPaintDevice__Margins struct {
	h          *C.QPagedPaintDevice__Margins
	isSubclass bool
}

func (this *QPagedPaintDevice__Margins) cPointer() *C.QPagedPaintDevice__Margins {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPagedPaintDevice__Margins) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPagedPaintDevice__Margins constructs the type using only CGO pointers.
func newQPagedPaintDevice__Margins(h *C.QPagedPaintDevice__Margins) *QPagedPaintDevice__Margins {
	if h == nil {
		return nil
	}
	return &QPagedPaintDevice__Margins{h: h}
}

// UnsafeNewQPagedPaintDevice__Margins constructs the type using only unsafe pointers.
func UnsafeNewQPagedPaintDevice__Margins(h unsafe.Pointer) *QPagedPaintDevice__Margins {
	if h == nil {
		return nil
	}

	return &QPagedPaintDevice__Margins{h: (*C.QPagedPaintDevice__Margins)(h)}
}

// Delete this object from C++ memory.
func (this *QPagedPaintDevice__Margins) Delete() {
	C.QPagedPaintDevice__Margins_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPagedPaintDevice__Margins) GoGC() {
	runtime.SetFinalizer(this, func(this *QPagedPaintDevice__Margins) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
