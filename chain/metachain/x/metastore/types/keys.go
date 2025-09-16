package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "metastore"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"

	// Version defines the current version the IBC module supports
	Version = "ibc-proto-1"

	// PortID is the default port id that module binds to
	PortID = "metastore"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = collections.NewPrefix("metastore-port-")
)

// ParamsKey is the prefix to retrieve all Params
var ParamsKey = collections.NewPrefix("p_metastore")
