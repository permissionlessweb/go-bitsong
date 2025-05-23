package cadence

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/go-bitsong/x/cadence/keeper"
	"github.com/bitsongofficial/go-bitsong/x/cadence/types"
)

// NewGenesisState - Create a new genesis state
func NewGenesisState(params types.Params) *types.GenesisState {
	return &types.GenesisState{
		Params: params,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *types.GenesisState {
	return NewGenesisState(types.DefaultParams())
}

// GetGenesisStateFromAppState returns x/auth GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) *types.GenesisState {
	var genesisState types.GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

func ValidateGenesis(data types.GenesisState) error {
	err := data.Params.Validate()
	if err != nil {
		return err
	}

	return nil
}

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) {
	// Validate init contents
	if err := ValidateGenesis(data); err != nil {
		panic(err)
	}

	// Set params
	if err := k.SetParams(ctx, data.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis export module state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)

	return &types.GenesisState{
		Params: params,
	}
}
