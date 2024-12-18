package simulation

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/devalvamseezeeve/iopn-distrubution/v2/x/iopn/types"
)

const (
	ibcCroDenomKey          = "ibc_cro_denom"
	ibcTimeoutKey           = "ibc_timeout"
	iopnAdminKey          = "iopn_admin"
	enableAutoDeploymentKey = "enable_auto_deployment"
	maxCallbackGasKey       = "max_callback_gas"
)

func GenIbcCroDenom(r *rand.Rand) string {
	randDenom := make([]byte, 32)
	r.Read(randDenom)
	return fmt.Sprintf("ibc/%s", hex.EncodeToString(randDenom))
}

func GenIbcTimeout(r *rand.Rand) uint64 {
	timeout := r.Uint64()
	return timeout
}

func GenCronosAdmin(r *rand.Rand, simState *module.SimulationState) string {
	adminAccount, _ := simtypes.RandomAcc(r, simState.Accounts)
	return adminAccount.Address.String()
}

func GenEnableAutoDeployment(r *rand.Rand) bool {
	return r.Intn(2) > 0
}

func GenMaxCallbackGas(r *rand.Rand) uint64 {
	maxCallbackGas := r.Uint64()
	return maxCallbackGas
}

// RandomizedGenState generates a random GenesisState for the iopn module
func RandomizedGenState(simState *module.SimulationState) {
	// iopn params
	var (
		ibcCroDenom          string
		ibcTimeout           uint64
		iopnAdmin          string
		enableAutoDeployment bool
		maxCallbackGas       uint64
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, ibcCroDenomKey, &ibcCroDenom, simState.Rand,
		func(r *rand.Rand) { ibcCroDenom = GenIbcCroDenom(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, ibcTimeoutKey, &ibcTimeout, simState.Rand,
		func(r *rand.Rand) { ibcTimeout = GenIbcTimeout(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, iopnAdminKey, &iopnAdmin, simState.Rand,
		func(r *rand.Rand) { iopnAdmin = GenCronosAdmin(r, simState) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, enableAutoDeploymentKey, &enableAutoDeployment, simState.Rand,
		func(r *rand.Rand) { enableAutoDeployment = GenEnableAutoDeployment(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, maxCallbackGasKey, &ibcTimeout, simState.Rand,
		func(r *rand.Rand) { maxCallbackGas = GenIbcTimeout(r) },
	)

	params := types.NewParams(ibcCroDenom, ibcTimeout, iopnAdmin, enableAutoDeployment, maxCallbackGas)
	iopnGenesis := &types.GenesisState{
		Params:            params,
		ExternalContracts: nil,
		AutoContracts:     nil,
	}

	bz, err := json.MarshalIndent(iopnGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(iopnGenesis)
}
