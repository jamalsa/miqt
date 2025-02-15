#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLPRESHAREDKEYAUTHENTICATOR_H
#define MIQT_QT_NETWORK_GEN_QSSLPRESHAREDKEYAUTHENTICATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QSslPreSharedKeyAuthenticator;
#else
typedef struct QByteArray QByteArray;
typedef struct QSslPreSharedKeyAuthenticator QSslPreSharedKeyAuthenticator;
#endif

void QSslPreSharedKeyAuthenticator_new(QSslPreSharedKeyAuthenticator** outptr_QSslPreSharedKeyAuthenticator);
void QSslPreSharedKeyAuthenticator_new2(QSslPreSharedKeyAuthenticator* authenticator, QSslPreSharedKeyAuthenticator** outptr_QSslPreSharedKeyAuthenticator);
void QSslPreSharedKeyAuthenticator_OperatorAssign(QSslPreSharedKeyAuthenticator* self, QSslPreSharedKeyAuthenticator* authenticator);
void QSslPreSharedKeyAuthenticator_Swap(QSslPreSharedKeyAuthenticator* self, QSslPreSharedKeyAuthenticator* other);
struct miqt_string QSslPreSharedKeyAuthenticator_IdentityHint(const QSslPreSharedKeyAuthenticator* self);
void QSslPreSharedKeyAuthenticator_SetIdentity(QSslPreSharedKeyAuthenticator* self, struct miqt_string identity);
struct miqt_string QSslPreSharedKeyAuthenticator_Identity(const QSslPreSharedKeyAuthenticator* self);
int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(const QSslPreSharedKeyAuthenticator* self);
void QSslPreSharedKeyAuthenticator_SetPreSharedKey(QSslPreSharedKeyAuthenticator* self, struct miqt_string preSharedKey);
struct miqt_string QSslPreSharedKeyAuthenticator_PreSharedKey(const QSslPreSharedKeyAuthenticator* self);
int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(const QSslPreSharedKeyAuthenticator* self);
void QSslPreSharedKeyAuthenticator_Delete(QSslPreSharedKeyAuthenticator* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
