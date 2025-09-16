package keeper_test

import (
	"testing"

	"metachain/x/metastore/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:        types.DefaultParams(),
		PortId:        types.PortID,
		StoredMetaMap: []types.StoredMeta{{Index: "0"}, {Index: "1"}}}

	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.Equal(t, genesisState.PortId, got.PortId)
	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.StoredMetaMap, got.StoredMetaMap)

}
