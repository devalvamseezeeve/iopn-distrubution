package icaauth

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"code.zeeve.net/client-projects/iopn/v2/x/icaauth/keeper"
	"code.zeeve.net/client-projects/iopn/v2/x/icaauth/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(fmt.Sprintf("Invalid icaauth module params: %v\n", genState.Params))
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	return genesis
}
