package common

import (
	"github.com/woop-chain/woop/accounts/keystore"
)

func KeyStoreForPath(p string) *keystore.KeyStore {
	return keystore.NewKeyStore(p, ScryptN, ScryptP)
}
