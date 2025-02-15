#include <QReadLocker>
#include <QReadWriteLock>
#include <QWriteLocker>
#include <qreadwritelock.h>
#include "gen_qreadwritelock.h"
#include "_cgo_export.h"

void QReadWriteLock_new(QReadWriteLock** outptr_QReadWriteLock) {
	QReadWriteLock* ret = new QReadWriteLock();
	*outptr_QReadWriteLock = ret;
}

void QReadWriteLock_new2(int recursionMode, QReadWriteLock** outptr_QReadWriteLock) {
	QReadWriteLock* ret = new QReadWriteLock(static_cast<QReadWriteLock::RecursionMode>(recursionMode));
	*outptr_QReadWriteLock = ret;
}

void QReadWriteLock_LockForRead(QReadWriteLock* self) {
	self->lockForRead();
}

bool QReadWriteLock_TryLockForRead(QReadWriteLock* self) {
	return self->tryLockForRead();
}

bool QReadWriteLock_TryLockForReadWithTimeout(QReadWriteLock* self, int timeout) {
	return self->tryLockForRead(static_cast<int>(timeout));
}

void QReadWriteLock_LockForWrite(QReadWriteLock* self) {
	self->lockForWrite();
}

bool QReadWriteLock_TryLockForWrite(QReadWriteLock* self) {
	return self->tryLockForWrite();
}

bool QReadWriteLock_TryLockForWriteWithTimeout(QReadWriteLock* self, int timeout) {
	return self->tryLockForWrite(static_cast<int>(timeout));
}

void QReadWriteLock_Unlock(QReadWriteLock* self) {
	self->unlock();
}

void QReadWriteLock_Delete(QReadWriteLock* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QReadWriteLock*>( self );
	} else {
		delete self;
	}
}

void QReadLocker_new(QReadWriteLock* readWriteLock, QReadLocker** outptr_QReadLocker) {
	QReadLocker* ret = new QReadLocker(readWriteLock);
	*outptr_QReadLocker = ret;
}

void QReadLocker_Unlock(QReadLocker* self) {
	self->unlock();
}

void QReadLocker_Relock(QReadLocker* self) {
	self->relock();
}

QReadWriteLock* QReadLocker_ReadWriteLock(const QReadLocker* self) {
	return self->readWriteLock();
}

void QReadLocker_Delete(QReadLocker* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QReadLocker*>( self );
	} else {
		delete self;
	}
}

void QWriteLocker_new(QReadWriteLock* readWriteLock, QWriteLocker** outptr_QWriteLocker) {
	QWriteLocker* ret = new QWriteLocker(readWriteLock);
	*outptr_QWriteLocker = ret;
}

void QWriteLocker_Unlock(QWriteLocker* self) {
	self->unlock();
}

void QWriteLocker_Relock(QWriteLocker* self) {
	self->relock();
}

QReadWriteLock* QWriteLocker_ReadWriteLock(const QWriteLocker* self) {
	return self->readWriteLock();
}

void QWriteLocker_Delete(QWriteLocker* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWriteLocker*>( self );
	} else {
		delete self;
	}
}

