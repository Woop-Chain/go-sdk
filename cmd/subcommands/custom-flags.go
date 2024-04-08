package cmd

import (
	"github.com/pkg/errors"
	"github.com/woop-chain/go-sdk/pkg/address"
	"github.com/woop-chain/go-sdk/pkg/common"
	"github.com/woop-chain/go-sdk/pkg/validation"
)

type wocAddress struct {
	address string
}

func (wocAddress wocAddress) String() string {
	return wocAddress.address
}

func (wocAddress *wocAddress) Set(s string) error {
	err := validation.ValidateAddress(s)
	if err != nil {
		return err
	}
	_, err = address.Bech32ToAddress(s)
	if err != nil {
		return errors.Wrap(err, "not a valid woc address")
	}
	wocAddress.address = s
	return nil
}

func (wocAddress wocAddress) Type() string {
	return "woc-address"
}

type chainIDWrapper struct {
	chainID *common.ChainID
}

func (chainIDWrapper chainIDWrapper) String() string {
	return chainIDWrapper.chainID.Name
}

func (chainIDWrapper *chainIDWrapper) Set(s string) error {
	chain, err := common.StringToChainID(s)
	chainIDWrapper.chainID = chain
	if err != nil {
		return err
	}
	return nil
}

func (chainIDWrapper chainIDWrapper) Type() string {
	return "chain-id"
}
