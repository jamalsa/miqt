#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QUDPSOCKET_H
#define MIQT_QT_NETWORK_GEN_QUDPSOCKET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractSocket;
class QByteArray;
class QHostAddress;
class QIODevice;
class QMetaObject;
class QNetworkDatagram;
class QNetworkInterface;
class QObject;
class QUdpSocket;
class QVariant;
#else
typedef struct QAbstractSocket QAbstractSocket;
typedef struct QByteArray QByteArray;
typedef struct QHostAddress QHostAddress;
typedef struct QIODevice QIODevice;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkDatagram QNetworkDatagram;
typedef struct QNetworkInterface QNetworkInterface;
typedef struct QObject QObject;
typedef struct QUdpSocket QUdpSocket;
typedef struct QVariant QVariant;
#endif

void QUdpSocket_new(QUdpSocket** outptr_QUdpSocket, QAbstractSocket** outptr_QAbstractSocket, QIODevice** outptr_QIODevice, QObject** outptr_QObject);
void QUdpSocket_new2(QObject* parent, QUdpSocket** outptr_QUdpSocket, QAbstractSocket** outptr_QAbstractSocket, QIODevice** outptr_QIODevice, QObject** outptr_QObject);
QMetaObject* QUdpSocket_MetaObject(const QUdpSocket* self);
void* QUdpSocket_Metacast(QUdpSocket* self, const char* param1);
struct miqt_string QUdpSocket_Tr(const char* s);
struct miqt_string QUdpSocket_TrUtf8(const char* s);
bool QUdpSocket_JoinMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress);
bool QUdpSocket_JoinMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface);
bool QUdpSocket_LeaveMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress);
bool QUdpSocket_LeaveMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface);
QNetworkInterface* QUdpSocket_MulticastInterface(const QUdpSocket* self);
void QUdpSocket_SetMulticastInterface(QUdpSocket* self, QNetworkInterface* iface);
bool QUdpSocket_HasPendingDatagrams(const QUdpSocket* self);
long long QUdpSocket_PendingDatagramSize(const QUdpSocket* self);
QNetworkDatagram* QUdpSocket_ReceiveDatagram(QUdpSocket* self);
long long QUdpSocket_ReadDatagram(QUdpSocket* self, char* data, long long maxlen);
long long QUdpSocket_WriteDatagram(QUdpSocket* self, QNetworkDatagram* datagram);
long long QUdpSocket_WriteDatagram2(QUdpSocket* self, const char* data, long long lenVal, QHostAddress* host, uint16_t port);
long long QUdpSocket_WriteDatagram3(QUdpSocket* self, struct miqt_string datagram, QHostAddress* host, uint16_t port);
struct miqt_string QUdpSocket_Tr2(const char* s, const char* c);
struct miqt_string QUdpSocket_Tr3(const char* s, const char* c, int n);
struct miqt_string QUdpSocket_TrUtf82(const char* s, const char* c);
struct miqt_string QUdpSocket_TrUtf83(const char* s, const char* c, int n);
QNetworkDatagram* QUdpSocket_ReceiveDatagram1(QUdpSocket* self, long long maxSize);
long long QUdpSocket_ReadDatagram3(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host);
long long QUdpSocket_ReadDatagram4(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host, uint16_t* port);
void QUdpSocket_override_virtual_Resume(void* self, intptr_t slot);
void QUdpSocket_virtualbase_Resume(void* self);
void QUdpSocket_override_virtual_ConnectToHost(void* self, intptr_t slot);
void QUdpSocket_virtualbase_ConnectToHost(void* self, struct miqt_string hostName, uint16_t port, int mode, int protocol);
void QUdpSocket_override_virtual_DisconnectFromHost(void* self, intptr_t slot);
void QUdpSocket_virtualbase_DisconnectFromHost(void* self);
void QUdpSocket_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_BytesAvailable(const void* self);
void QUdpSocket_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_BytesToWrite(const void* self);
void QUdpSocket_override_virtual_CanReadLine(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_CanReadLine(const void* self);
void QUdpSocket_override_virtual_SetReadBufferSize(void* self, intptr_t slot);
void QUdpSocket_virtualbase_SetReadBufferSize(void* self, long long size);
void QUdpSocket_override_virtual_SocketDescriptor(void* self, intptr_t slot);
intptr_t QUdpSocket_virtualbase_SocketDescriptor(const void* self);
void QUdpSocket_override_virtual_SetSocketDescriptor(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_SetSocketDescriptor(void* self, intptr_t socketDescriptor, int state, int openMode);
void QUdpSocket_override_virtual_SetSocketOption(void* self, intptr_t slot);
void QUdpSocket_virtualbase_SetSocketOption(void* self, int option, QVariant* value);
void QUdpSocket_override_virtual_SocketOption(void* self, intptr_t slot);
QVariant* QUdpSocket_virtualbase_SocketOption(void* self, int option);
void QUdpSocket_override_virtual_Close(void* self, intptr_t slot);
void QUdpSocket_virtualbase_Close(void* self);
void QUdpSocket_override_virtual_IsSequential(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_IsSequential(const void* self);
void QUdpSocket_override_virtual_AtEnd(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_AtEnd(const void* self);
void QUdpSocket_override_virtual_WaitForConnected(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForConnected(void* self, int msecs);
void QUdpSocket_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForReadyRead(void* self, int msecs);
void QUdpSocket_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForBytesWritten(void* self, int msecs);
void QUdpSocket_override_virtual_WaitForDisconnected(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForDisconnected(void* self, int msecs);
void QUdpSocket_override_virtual_ReadData(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_ReadData(void* self, char* data, long long maxlen);
void QUdpSocket_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
void QUdpSocket_override_virtual_WriteData(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_WriteData(void* self, const char* data, long long lenVal);
void QUdpSocket_Delete(QUdpSocket* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
