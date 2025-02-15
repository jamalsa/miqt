#include <QBasicTimer>
#include <QObject>
#include <qbasictimer.h>
#include "gen_qbasictimer.h"
#include "_cgo_export.h"

void QBasicTimer_new(QBasicTimer** outptr_QBasicTimer) {
	QBasicTimer* ret = new QBasicTimer();
	*outptr_QBasicTimer = ret;
}

void QBasicTimer_Swap(QBasicTimer* self, QBasicTimer* other) {
	self->swap(*other);
}

bool QBasicTimer_IsActive(const QBasicTimer* self) {
	return self->isActive();
}

int QBasicTimer_TimerId(const QBasicTimer* self) {
	return self->timerId();
}

void QBasicTimer_Start(QBasicTimer* self, int msec, QObject* obj) {
	self->start(static_cast<int>(msec), obj);
}

void QBasicTimer_Start2(QBasicTimer* self, int msec, int timerType, QObject* obj) {
	self->start(static_cast<int>(msec), static_cast<Qt::TimerType>(timerType), obj);
}

void QBasicTimer_Stop(QBasicTimer* self) {
	self->stop();
}

void QBasicTimer_Delete(QBasicTimer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QBasicTimer*>( self );
	} else {
		delete self;
	}
}

