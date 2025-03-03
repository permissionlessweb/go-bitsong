package app

// to run: bitsongd export --for-zero-height --output-document v0.21.5-export.json
import (
	"encoding/json"
	"fmt"
	"log"

	storetypes "cosmossdk.io/store/types"
	v020 "github.com/bitsongofficial/go-bitsong/app/upgrades/v020"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// ExportAppStateAndValidators exports the state of the application for a genesis
// file.
func (app *BitsongApp) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string,
) (servertypes.ExportedApp, error) {
	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true)

	// We export at last height + 1, because that's the height at which
	// Tendermint will start InitChain.
	height := app.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
		app.prepForZeroHeightGenesis(ctx, jailAllowedAddrs)
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
func (app *BitsongApp) prepForZeroHeightGenesis(ctx sdk.Context, jailAllowedAddrs []string) {
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

	/* Just to be safe, assert the invariants on current state. */
	// app.AppKeepers.CrisisKeeper.AssertInvariants(ctx)

	/* Handle fee distribution state. */
	// app.AppKeepers.StakingKeeper.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
	// 	valAddr := sdk.ValAddress(val.GetOperator())
	// 	dels, _ := app.AppKeepers.StakingKeeper.GetValidatorDelegations(ctx, valAddr)
	// 	for _, del := range dels {
	// 		fmt.Println(fmt.Printf("del_info: %q %v", val.GetOperator(), del.GetDelegatorAddr()))

	// 		// if del.Shares.LTE(math.LegacyZeroDec()) {
	// 		// 	ctx.Logger().Info(fmt.Sprintf("removing negative delegations: %q %v", val.GetOperator(), del.GetDelegatorAddr()))
	// 		// 	// remove reward information from distribution store
	// 		// 	has, _ := app.AppKeepers.DistrKeeper.HasDelegatorStartingInfo(ctx, valAddr, sdk.AccAddress(del.DelegatorAddress))
	// 		// 	if has {
	// 		// 		app.AppKeepers.DistrKeeper.DeleteDelegatorStartingInfo(ctx, valAddr, sdk.AccAddress(del.DelegatorAddress))
	// 		// 	}
	// 		// 	// remove delegation from staking store
	// 		// 	if err := app.AppKeepers.StakingKeeper.RemoveDelegation(ctx, del); err != nil {
	// 		// 		panic(err)
	// 		// 	}
	// 		// } else {
	// 		notHas, _ := app.AppKeepers.DistrKeeper.HasDelegatorStartingInfo(ctx, valAddr, sdk.AccAddress(del.GetDelegatorAddr()))
	// 		if !notHas {
	// 			panic(distrtypes.ErrEmptyDelegationDistInfo)
	// 		}
	// 		endingPeriod, _ := app.AppKeepers.DistrKeeper.IncrementValidatorPeriod(ctx, val)
	// 		rewardsRaw, patched := v020.CustomCalculateDelegationRewards(ctx, &app.AppKeepers, val, del, endingPeriod)
	// 		outstanding, _ := app.AppKeepers.DistrKeeper.GetValidatorOutstandingRewardsCoins(ctx, valAddr)

	// 		if patched {
	// 			ctx.Logger().Info("~=~=~=~=~~=~=~=~=~~=~=~=~=~~=~=~=~=~~=~=~=~=~~=~=~=~=~~=~=~=~=~~=~=~=~=~~=~=~=~=~~=~=~=~=~")
	// 			ctx.Logger().Info(fmt.Sprintf("PATCHED: %q %v", val.GetOperator(), del.GetDelegatorAddr()))
	// 			err := v020.V018ManualDelegationRewardsPatch(ctx, rewardsRaw, outstanding, &app.AppKeepers, val, del, endingPeriod)
	// 			if err != nil {
	// 				panic(err)
	// 			}
	// 		}
	// 	}
	// 	// }

	// 	return false
	// })

	/*  ensure no delegations exist without starting info*/
	dels, _ := app.AppKeepers.StakingKeeper.GetAllDelegations(ctx)
	for _, del := range dels {
		valAddr := sdk.ValAddress(del.GetValidatorAddr())
		// delAddr := sdk.AccAddress(del.GetDelegatorAddr())
		has, _ := app.AppKeepers.DistrKeeper.HasDelegatorStartingInfo(ctx, valAddr, sdk.AccAddress(del.GetDelegatorAddr()))
		// if valAddr.String() == "bitsongvaloper1jsaud5d8weze74a5e5w9ercxamtglg9wapc43m" && delAddr.String() == "bitsong1qqqj207k07465gu4s7e4d3lr39as9dxxtlwzum" {
		// 	// calculate rewards
		// 	val, _ := app.AppKeepers.StakingKeeper.Validator(ctx, valAddr)
		// 	endingPeriod, err := app.AppKeepers.DistrKeeper.IncrementValidatorPeriod(ctx, val)
		// 	rewardsRaw, _ := v020.CustomCalculateDelegationRewards(ctx, &app.AppKeepers, val, del, endingPeriod)
		// 	outstanding, err := app.AppKeepers.DistrKeeper.GetValidatorOutstandingRewardsCoins(ctx, sdk.ValAddress(del.GetValidatorAddr()))

		// 	err = v020.V018ManualDelegationRewardsPatch(ctx, rewardsRaw, outstanding, &app.AppKeepers, val, del, endingPeriod)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }
		if !has {
			// val, _ := app.AppKeepers.StakingKeeper.GetValidator(ctx, valAddr)
			// if val.Jailed {
			// 	panic("val is not jailed")
			// } else {
			// 	app.AppKeepers.DistrKeeper.SetDelegatorStartingInfo(ctx, valAddr, delAddr, distrtypes.NewDelegatorStartingInfo(0, math.LegacyZeroDec(), uint64(ctx.BlockHeight())))
			// }
			// if delegation starting info does not exist, we assume & assert:
			// - validator is jailed
			// if this is the case, then we must:
			// - inject dummy delegator starting info
			// - manually claim rewards for delegator, using the patch
			// continue
			// panic(distrtypes.ErrEmptyDelegationDistInfo)

		}

		/* ensure all rewards are patched */
		val, _ := app.AppKeepers.StakingKeeper.Validator(ctx, valAddr)
		endingPeriod, _ := app.AppKeepers.DistrKeeper.IncrementValidatorPeriod(ctx, val)
		if val.GetTokens().IsNil() {
			/* will error if still broken */
			fmt.Println(fmt.Printf("delegator: %q", del.DelegatorAddress))
			fmt.Println(fmt.Printf("validator: %q", del.ValidatorAddress))
			fmt.Println(fmt.Printf(" del.Shares: %q", del.Shares))
			fmt.Println(fmt.Printf("  val.GetStatus(): %q", val.GetStatus()))
			fmt.Println(fmt.Print(val.IsBonded()))
			fmt.Println(fmt.Printf(" val.GetTokens(): %q", val.GetTokens()))

		} else {
			app.AppKeepers.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)
		}
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

	// withdraw all delegator rewards
	dels, _ = app.AppKeepers.StakingKeeper.GetAllDelegations(ctx)
	for _, delegation := range dels {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		delAddr := sdk.AccAddress(delegation.DelegatorAddress)
		if err != nil {
			panic(err)
		}
		val, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, valAddr)
		if err != nil {
			panic(err)
		} else if val.GetTokens().IsNil() {
			ctx.Logger().Info(fmt.Sprintf("val tokens for %q: %v", val.GetOperator(), val.GetTokens()))
		} else {
			ctx.Logger().Info(fmt.Sprintf("val tokens for %q: %v", val.GetOperator(), val.GetTokens()))
			ctx.Logger().Info(fmt.Sprintf("withdrawing %q: %v", val.GetOperator(), delAddr.String()))
			endingPeriod, err := app.AppKeepers.DistrKeeper.IncrementValidatorPeriod(ctx, val)
			if err != nil {
				panic(err)
			}
			rewardsRaw, patched := v020.CustomCalculateDelegationRewards(ctx, &app.AppKeepers, val, delegation, endingPeriod)
			outstanding, err := app.AppKeepers.DistrKeeper.GetValidatorOutstandingRewardsCoins(ctx, sdk.ValAddress(delegation.GetValidatorAddr()))
			if patched {
				err = v020.V018ManualDelegationRewardsPatch(ctx, rewardsRaw, outstanding, &app.AppKeepers, val, delegation, endingPeriod)
				if err != nil {
					panic(err)
				}
			}
			// _, err = app.AppKeepers.DistrKeeper.WithdrawDelegationRewards(ctx, delAddr, valAddr)
			if err != nil {
				panic(err)
			}
		}
	}

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
		// fmt.Println(fmt.Printf("delegator: %q", del.DelegatorAddress))
		// fmt.Println(fmt.Printf("validator: %q", del.ValidatorAddress))
		// fmt.Println(fmt.Printf("refCount: %q", refCount))

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

	// Iterate through validators by power descending, reset bond heights, and
	// update bond intra-tx counters.
	store := ctx.KVStore(app.keys[stakingtypes.StoreKey])
	iter := storetypes.KVStoreReversePrefixIterator(store, stakingtypes.ValidatorsKey)
	deletedCounter := int16(0)
	notDeletedCounter := int16(0)

	for ; iter.Valid(); iter.Next() {
		key := iter.Key()[1:]
		addr := sdk.ValAddress(key)
		validator, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, addr)
		if err != nil {
			ctx.Logger().Info(fmt.Sprintf("expected validator, not found: %q. removing key from store...", addr.String()))
			store.Delete(key)
			deletedCounter++
			continue
		}

		validator.UnbondingHeight = 0
		if applyAllowedAddrs && !allowedAddrsMap[addr.String()] {
			validator.Jailed = true
		}

		app.AppKeepers.StakingKeeper.SetValidator(ctx, validator)
		notDeletedCounter++
	}

	iter.Close()
	/* Handle slashing state. */

	fmt.Println(fmt.Printf("notdeleted validator store key count: %q", notDeletedCounter))
	fmt.Println(fmt.Printf("deleted validator store key count: %q", deletedCounter))

	// reset start height on signing infos
	app.AppKeepers.SlashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			app.AppKeepers.SlashingKeeper.SetValidatorSigningInfo(ctx, addr, info)
			return false
		},
	)
}

// /* remove any remaining validator keys from store.This runs after we retrieve all current validators from staking keeper store,
//  preventing us from deleting active validators store. */
// store := sdkCtx.KVStore(k.GetKey(stakingtypes.StoreKey))
// iter := storetypes.KVStoreReversePrefixIterator(store, stakingtypes.ValidatorsKey)
// counter := int16(0)

// for ; iter.Valid(); iter.Next() {
// 	key := iter.Key()[1:]
// 	addr := sdk.ValAddress(key)
// 	validator, err := k.StakingKeeper.GetValidator(sdkCtx, addr)
// 	if err != nil {
// 		sdkCtx.Logger().Info(fmt.Sprintf("expected validator, not found: %q", addr.String()))
// 		store.Delete(key)
// 		counter++
// 		continue
// 	} else {
// 		sdkCtx.Logger().Info("-==-=-=-==---=-=-=-=-=--=-=-=-")
// 		sdkCtx.Logger().Info(fmt.Sprintf("found: %q", validator.OperatorAddress))
// 	}
// 	counter++
// }
