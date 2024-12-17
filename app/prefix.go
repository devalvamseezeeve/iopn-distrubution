package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cmdcfg "code.zeeve.net/client-projects/iopn/v2/cmd/iopnd/config"
	ethcfg "github.com/evmos/ethermint/cmd/config"
)

func SetConfig() {
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	ethcfg.SetBip44CoinType(config)
	// Make sure address is compatible with ethereum
	config.SetAddressVerifier(VerifyAddressFormat)
	config.Seal()
}
