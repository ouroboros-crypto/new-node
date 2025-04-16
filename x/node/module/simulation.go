package node

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ouroboros-crypto/node/testutil/sample"
	nodesimulation "github.com/ouroboros-crypto/node/x/node/simulation"
	"github.com/ouroboros-crypto/node/x/node/types"
)

// avoid unused import issue
var (
	_ = nodesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgReinvest = "op_weight_msg_reinvest"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReinvest int = 100

	opWeightMsgBurnExtraCoins = "op_weight_msg_burn_extra_coins"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnExtraCoins int = 100

	opWeightMsgUpdateStakingParams = "op_weight_msg_update_staking_params"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStakingParams int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nodeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nodeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgReinvest int
	simState.AppParams.GetOrGenerate(opWeightMsgReinvest, &weightMsgReinvest, nil,
		func(_ *rand.Rand) {
			weightMsgReinvest = defaultWeightMsgReinvest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReinvest,
		nodesimulation.SimulateMsgReinvest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurnExtraCoins int
	simState.AppParams.GetOrGenerate(opWeightMsgBurnExtraCoins, &weightMsgBurnExtraCoins, nil,
		func(_ *rand.Rand) {
			weightMsgBurnExtraCoins = defaultWeightMsgBurnExtraCoins
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnExtraCoins,
		nodesimulation.SimulateMsgBurnExtraCoins(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStakingParams int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStakingParams, &weightMsgUpdateStakingParams, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStakingParams = defaultWeightMsgUpdateStakingParams
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStakingParams,
		nodesimulation.SimulateMsgUpdateStakingParams(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgReinvest,
			defaultWeightMsgReinvest,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nodesimulation.SimulateMsgReinvest(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
	opWeightMsgBurnExtraCoins,
	defaultWeightMsgBurnExtraCoins,
	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		nodesimulation.SimulateMsgBurnExtraCoins(am.accountKeeper, am.bankKeeper, am.keeper)
		return nil
	},
),
simulation.NewWeightedProposalMsg(
	opWeightMsgUpdateStakingParams,
	defaultWeightMsgUpdateStakingParams,
	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		nodesimulation.SimulateMsgUpdateStakingParams(am.accountKeeper, am.bankKeeper, am.keeper)
		return nil
	},
),
// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
