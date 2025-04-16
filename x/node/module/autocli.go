package node

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ouroboros-crypto/node/api/node/node"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{

				{
					RpcMethod:      "GetPosmining",
					Use:            "get-posmining [address] [coin]",
					Short:          "Query get-posmining",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "coin"}},
				},

				{
					RpcMethod:      "GetPosmining",
					Use:            "get-posmining [address] [coin]",
					Short:          "Query get-posmining",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "coin"}},
				},

				{
					RpcMethod:      "GetProfile",
					Use:            "get-profile [address]",
					Short:          "Query get-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "Reinvest",
					Use:            "reinvest [coin]",
					Short:          "Send a reinvest tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "coin"}},
				},
				{
					RpcMethod:      "BurnExtraCoins",
					Use:            "burn-extra-coins",
					Short:          "Send a burn-extra-coins tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "UpdateStakingParams",
					Use:            "update-staking-params [unbonding_time] [max_validators]",
					Short:          "Send a update-staking-params tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "unbonding_time"}, {ProtoField: "max_validators"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
