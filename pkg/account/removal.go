package account

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/woop-chain/go-sdk/pkg/common"
	"github.com/woop-chain/go-sdk/pkg/store"
)

// RemoveAccount - removes an account from the keystore
func RemoveAccount(name string) error {
	accountExists := store.DoesNamedAccountExist(name)

	if !accountExists {
		return fmt.Errorf("account %s doesn't exist", name)
	}

	uDir, _ := homedir.Dir()
	wikiCLIDir := path.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName)
	accountDir := fmt.Sprintf("%s/%s", wikiCLIDir, name)
	os.RemoveAll(accountDir)

	return nil
}
