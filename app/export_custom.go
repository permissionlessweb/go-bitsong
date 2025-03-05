package app

// to run: bitsongd export --for-zero-height --output-document v0.21.5-export.json
import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	storetypes "cosmossdk.io/store/types"
	v020 "github.com/bitsongofficial/go-bitsong/app/upgrades/v020"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// ConditionalJSON represents a JSON object for logging conditionals

type ConditionalJSON struct {
	NoStartingInfo           []NoStartingInfo     `json:"no_starting_info"`
	NoStartingInfoCount      int                  `json:"no_starting_info_count"`
	NoSigningInfo            []NoSigningInfo      `json:"no_signing_info"`
	ZeroRewards              []ZeroRewards        `json:"zero_rewards"`
	ZeroRewardsCount         int                  `json:"zero_rewards_count"`
	ZeroTokenValidators      []ZeroTokenValidator `json:"zero_token_validators"` // todo: add total count
	ZeroTokenValidatorsCount int                  `json:"zero_token_validators_count"`
}

type NoStartingInfo struct {
	ValidatorAddress string `json:"validator_address"`
	DelegatorAddress string `json:"delegator_address"`
	Power            string `json:"power"`
	KVStoreKey       string `json:"kv_store_key"`
}
type NoSigningInfo struct {
	ValidatorAddress string `json:"validator_address"`
	Tokens           string `json:"tokens"`
}

type ZeroRewards struct {
	ValidatorAddress string `json:"validator_address"`
	DelegatorAddress string `json:"delegator_address"`
	Power            string `json:"power"`
	KVStoreKey       string `json:"kv_store_key"`
}

type ZeroTokenValidator struct {
	OperatorAddress string `json:"operator_address"`
	Power           string `json:"power"`
	KVStoreKey      string `json:"kv_store_key"`
}

// Custom Export Debugging Our Current KvStores:
// - x/distribution:
//   - delegator starting info:
//   - validator slash event:
//   - historical reference:
//
// x/slashing:
//   - validator signing info:
//   - historical reference:
func (app *BitsongApp) CustomExportAppStateAndValidators(

	forZeroHeight bool, jailAllowedAddrs []string,

) (servertypes.ExportedApp, error) {
	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true)

	// We export at last height + 1, because that's the height at which
	// Tendermint will start InitChain.
	height := app.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
		app.customPrepForZeroHeightGenesis(ctx, jailAllowedAddrs)
	}

	genState, _ := app.mm.ExportGenesis(ctx, app.appCodec)
	appState, err := json.MarshalIndent(genState, "", "  ")
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	validators, err := staking.WriteValidators(ctx, app.AppKeepers.StakingKeeper)
	if err != nil {
		return servertypes.ExportedApp{}, err
	}
	return servertypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: app.BaseApp.GetConsensusParams(ctx),
	}, nil
}

// prepare for fresh start at zero height
// NOTE zero height genesis is a temporary feature which will be deprecated
//
//	in favour of export at a block height
func (app *BitsongApp) customPrepForZeroHeightGenesis(ctx sdk.Context, jailAllowedAddrs []string) {
	condJSON := ConditionalJSON{
		NoStartingInfo:      make([]NoStartingInfo, 0),
		NoStartingInfoCount: 0,
		NoSigningInfo:       make([]NoSigningInfo, 0),
		ZeroRewards:         make([]ZeroRewards, 0),
		ZeroRewardsCount:    0,
		ZeroTokenValidators: make([]ZeroTokenValidator, 0),
	}
	applyAllowedAddrs := false

	// check if there is a allowed address list
	if len(jailAllowedAddrs) > 0 {
		applyAllowedAddrs = true
	}

	allowedAddrsMap := make(map[string]bool)

	for _, addr := range jailAllowedAddrs {
		_, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			log.Fatal(err)
		}
		allowedAddrsMap[addr] = true
	}

	// withdraw all validator commission
	app.AppKeepers.StakingKeeper.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		_, err := app.AppKeepers.DistrKeeper.WithdrawValidatorCommission(ctx, sdk.ValAddress(val.GetOperator()))
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("attempted to withdraw commission from validator with none, skipping: %q", val.GetOperator()))
			return false
		}
		return false
	})

	app.AppKeepers.StakingKeeper.IterateAllDelegations(ctx, func(del stakingtypes.Delegation) (stop bool) {
		valAddr, err := sdk.ValAddressFromBech32(del.ValidatorAddress)
		delAddr := sdk.AccAddress(del.DelegatorAddress)
		if err != nil {
			panic(err)
		}
		has, _ := app.AppKeepers.DistrKeeper.HasDelegatorStartingInfo(ctx, valAddr, sdk.AccAddress(del.GetDelegatorAddr()))

		// add count to the log file printed for nostarting infos
		if !has {
			// Append no starting info conditional to the ConditionalJSON object
			condJSON.NoStartingInfo = append(condJSON.NoStartingInfo, NoStartingInfo{
				ValidatorAddress: valAddr.String(),
				DelegatorAddress: delAddr.String(),
				Power:            del.Shares.String(),
				KVStoreKey:       "", // Add the KV store key here
			})
			condJSON.NoStartingInfoCount = len(condJSON.NoStartingInfo)
			// todo: continue to the next del in the iteration of dels
		}

		val, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, valAddr)
		if err != nil {
			panic(err)
		} else if val.GetTokens().IsZero() {
			condJSON.ZeroTokenValidators = append(condJSON.ZeroTokenValidators, ZeroTokenValidator{
				OperatorAddress: val.GetOperator(),
				Power:           del.Shares.String(),
				KVStoreKey:      "", // Add the KV store key here
			})
			condJSON.ZeroTokenValidatorsCount = len(condJSON.ZeroTokenValidators)
		} else {
			valBz, err := app.AppKeepers.StakingKeeper.ValidatorAddressCodec().StringToBytes(val.GetOperator())
			if err != nil {
				panic(err)
			}
			rewards, err := app.AppKeepers.DistrKeeper.GetValidatorCurrentRewards(ctx, valBz)
			if err != nil {
				panic(err)
			}

			endingPeriod := rewards.Period
			// endingPeriod, err := app.AppKeepers.DistrKeeper.IncrementValidatorPeriod(ctx, val)
			rewardsRaw, patched := v020.CustomCalculateDelegationRewards(ctx, &app.AppKeepers, val, del, endingPeriod)
			outstanding, err := app.AppKeepers.DistrKeeper.GetValidatorOutstandingRewardsCoins(ctx, sdk.ValAddress(del.GetValidatorAddr()))
			if err != nil {
				panic(err)
			}
			if rewardsRaw.IsZero() {
				// append to log json
				condJSON.ZeroRewards = append(condJSON.ZeroRewards, ZeroRewards{
					ValidatorAddress: valAddr.String(),
					DelegatorAddress: delAddr.String(),
					Power:            del.Shares.String(),
					KVStoreKey:       "", // Add the KV store key here
				})
				condJSON.ZeroRewardsCount = len(condJSON.ZeroRewards)

			} else if patched {
				//  claim rewards with logic to patch
				fmt.Printf("patched: %v\n", del.ValidatorAddress)
				err = v020.V018ManualDelegationRewardsPatch(ctx, rewardsRaw, outstanding, &app.AppKeepers, val, del, endingPeriod)
				if err != nil {
					panic(err)
				}
			} else {
				// claim rewards normally
				_, err := app.AppKeepers.DistrKeeper.WithdrawDelegationRewards(ctx, delAddr, valAddr)
				if err != nil {
					// todo: if err is panic: no delegation for (address, validator) tuple, we remove from the kvstore
					panic(err)
				}
			}
		}
		return false
	})

	// VALIDATION
	xStake := ctx.KVStore(app.keys[stakingtypes.StoreKey])
	stakeIter := storetypes.KVStoreReversePrefixIterator(xStake, stakingtypes.ValidatorsKey)
	deletedCounter := int16(0)
	notDeletedCounter := int16(0)

	for ; stakeIter.Valid(); stakeIter.Next() {
		key := stakeIter.Key()[1:]
		addr := sdk.ValAddress(key)
		// confirm by sdk.ValAddr
		validator, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, addr)

		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("expected validator, not found: %q. removing key from store...", addr.String()))
			xStake.Delete(key)
			deletedCounter++
			continue
		}
		validator.UnbondingHeight = 0
		if applyAllowedAddrs && !allowedAddrsMap[addr.String()] {
			validator.Jailed = true
		}

		// 	// assert validator signing info exists in x/slashing
		_, err = app.AppKeepers.SlashingKeeper.GetValidatorSigningInfo(ctx, validator.ConsensusPubkey.Value)
		if err != nil {
			condJSON.NoSigningInfo = append(condJSON.NoSigningInfo, NoSigningInfo{
				ValidatorAddress: validator.OperatorAddress,
				Tokens:           validator.Tokens.String(),
			})
			panic(err)
		}

		// 	// iterate & assert delegations for this validators key from smallest delegations first
		delpre := stakingtypes.GetDelegationsByValPrefixKey(addr)
		delIter := storetypes.KVStoreReversePrefixIterator(xStake, storetypes.PrefixEndBytes(delpre))
		for ; delIter.Valid(); delIter.Next() {
			// delKey
			valAddr, delAddr, err := stakingtypes.ParseDelegationsByValKey(delIter.Key())
			if err != nil {
				panic(err)
			}

			// 		// get from keeper
			del, err := app.AppKeepers.StakingKeeper.GetDelegation(ctx, delAddr, valAddr)
			if err != nil {
				// TODO: improve err
				panic(err)
			}
			fmt.Printf("del: %v\n", del)

		}
		app.AppKeepers.StakingKeeper.SetValidator(ctx, validator)
		notDeletedCounter++
	}

	stakeIter.Close()

	// Marshal the ConditionalJSON object to JSON
	jsonBytes, err := json.MarshalIndent(condJSON, "", "  ")
	if err != nil {
		panic(err)
	}
	// Write the JSON to a file
	fileName := "conditionals.json"
	err = os.WriteFile(fileName, jsonBytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote conditionals to %s\n", fileName)

	// clear validator slash events
	app.AppKeepers.DistrKeeper.DeleteAllValidatorSlashEvents(ctx)

	// clear validator historical rewards
	app.AppKeepers.DistrKeeper.DeleteAllValidatorHistoricalRewards(ctx)

	// set context height to zero
	height := ctx.BlockHeight()
	ctx = ctx.WithBlockHeight(0)

	// reinitialize all validators
	app.AppKeepers.StakingKeeper.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		valAddr := sdk.ValAddress(val.GetOperator())
		// donate any unwithdrawn outstanding reward fraction tokens to the community pool
		scraps, _ := app.AppKeepers.DistrKeeper.GetValidatorOutstandingRewardsCoins(ctx, valAddr)
		feePool, _ := app.AppKeepers.DistrKeeper.FeePool.Get(ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(scraps...)
		app.AppKeepers.DistrKeeper.FeePool.Set(ctx, feePool)

		app.AppKeepers.DistrKeeper.Hooks().AfterValidatorCreated(ctx, valAddr)
		return false
	})
	// withdraw all delegator rewards
	dels, err := app.AppKeepers.StakingKeeper.GetAllDelegations(ctx)
	if err != nil {
		panic(err)
	}

	// reinitialize all delegations
	for _, del := range dels {
		valAddr, err := sdk.ValAddressFromBech32(del.ValidatorAddress)
		if err != nil {
			panic(err)
		}
		delAddr, err := sdk.AccAddressFromBech32(del.DelegatorAddress)
		if err != nil {
			panic(err)
		}
		refCount := app.AppKeepers.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx)

		// omit specific jailed vals
		if refCount != uint64(0) {
			app.AppKeepers.DistrKeeper.Hooks().BeforeDelegationCreated(ctx, delAddr, valAddr)
			app.AppKeepers.DistrKeeper.Hooks().AfterDelegationModified(ctx, delAddr, valAddr)
		}
	}

	// reset context height
	ctx = ctx.WithBlockHeight(height)

	/* Handle staking state. */

	// iterate through redelegations, reset creation height
	app.AppKeepers.StakingKeeper.IterateRedelegations(ctx, func(_ int64, red stakingtypes.Redelegation) (stop bool) {
		for i := range red.Entries {
			red.Entries[i].CreationHeight = 0
		}
		app.AppKeepers.StakingKeeper.SetRedelegation(ctx, red)
		return false
	})

	// iterate through unbonding delegations, reset creation height
	app.AppKeepers.StakingKeeper.IterateUnbondingDelegations(ctx, func(_ int64, ubd stakingtypes.UnbondingDelegation) (stop bool) {
		for i := range ubd.Entries {
			ubd.Entries[i].CreationHeight = 0
		}
		app.AppKeepers.StakingKeeper.SetUnbondingDelegation(ctx, ubd)
		return false
	})

	// x/slashing store assertion: pubkey relation
	xSlash := ctx.KVStore(app.keys[slashingtypes.StoreKey])
	slashiter := storetypes.KVStoreReversePrefixIterator(xSlash, slashingtypes.AddrPubkeyRelationKeyPrefix)
	for ; slashiter.Valid(); slashiter.Next() {

	}
	slashiter.Close()

	// x/distr store assertion - delegator starting info
	xDistr := ctx.KVStore(app.keys[distrtypes.StoreKey])
	distrIter := storetypes.KVStoreReversePrefixIterator(xDistr, distrtypes.DelegatorStartingInfoPrefix)
	for ; distrIter.Valid(); distrIter.Next() {

	}
	distrIter.Close()

	/* Handle slashing state. */

	// reset start height on signing infos
	app.AppKeepers.SlashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			app.AppKeepers.SlashingKeeper.SetValidatorSigningInfo(ctx, addr, info)
			return false
		},
	)

	// /* Just to be safe, assert the invariants on current state. */
	// app.AppKeepers.CrisisKeeper.AssertInvariants(ctx)

}
