package types

const (
	// ModuleName defines the module name
	ModuleName = "node"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_node"

	// PosminingKey defines the primary module store key
	PosminingKey = "Posmining/value/"
)

var ParamsKey = []byte("p_node")
var LastValidatorRewardTimeKey = []byte("last_validator_reward_time")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
