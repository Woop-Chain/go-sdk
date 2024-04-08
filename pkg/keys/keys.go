package keys

import (
	"fmt"
	"os"
	"path"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/woop-chain/go-sdk/pkg/address"
	"github.com/woop-chain/woop/accounts/keystore"

	// "github.com/ethereum/go-ethereum/crypto"

	homedir "github.com/mitchellh/go-homedir"
)

func checkAndMakeKeyDirIfNeeded() string {
	userDir, _ := homedir.Dir()
	wikiCLIDir := path.Join(userDir, ".wiki_cli", "keystore")
	if _, err := os.Stat(wikiCLIDir); os.IsNotExist(err) {
		// Double check with Leo what is right file persmission
		os.Mkdir(wikiCLIDir, 0700)
	}

	return wikiCLIDir
}

func ListKeys(keystoreDir string) {
	wikiCLIDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(wikiCLIDir, scryptN, scryptP)
	// keystore.KeyStore
	allAccounts := ks.Accounts()
	fmt.Printf("Woop Address:%s File URL:\n", strings.Repeat(" ", ethCommon.AddressLength*2))
	for _, account := range allAccounts {
		fmt.Printf("%s\t\t %s\n", address.ToBech32(account.Address), account.URL)
	}
}

func AddNewKey(password string) {
	wikiCLIDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(wikiCLIDir, scryptN, scryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Printf("new account error: %v\n", err)
	}
	fmt.Printf("account: %s\n", address.ToBech32(account.Address))
	fmt.Printf("URL: %s\n", account.URL)
}
