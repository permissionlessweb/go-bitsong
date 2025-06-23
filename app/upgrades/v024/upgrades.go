package v024

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/go-bitsong/app/keepers"
	"github.com/bitsongofficial/go-bitsong/app/upgrades"

	"github.com/cosmos/cosmos-sdk/types/module"
)

// iterates through all existing bs721 contracts,
// expects bs721 minter to be a bs721-curve,
// retrives params for bs721-curve, initialized new curve factory w/ temporary params for reminting nfts to existing holders
// iterates through bs721-curve `address_tokens` state, mints new nfts to addresses with > 0 nfts
// all mint fees are routed to gov module thanks to temporary params
// once all token holders for collection own nfts in new bs721-base state, update params back to orignal ones
// - no need to instantiate royalty contracts, existing ones can be reused
// - actual bs721 nft collection is created on bs721-curve initialization
// - no need to create new bs721-curve through factory
// x/protocolpool is re-enabled as externalCommunityPool, which breaks old unmigratable contracts, preventing any new mints of the old contracts
// sets ubtsg
func CreateV024UpgradeHandler(mm *module.Manager, configurator module.Configurator, bpm upgrades.BaseAppParamManager, k *keepers.AppKeepers) upgradetypes.UpgradeHandler {
	return func(context context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		sdkCtx := sdk.UnwrapSDKContext(context)
		logger := sdkCtx.Logger().With("upgrade", UpgradeName)

		// Run migrations first
		logger.Info(fmt.Sprintf("pre migrate version map: %v", vm))
		versionMap, err := mm.RunMigrations(sdkCtx, configurator, vm)
		if err != nil {
			return nil, err
		}

		// apply logic patch
		err = CustomV024PatchLogic(sdkCtx, k)
		if err != nil {
			return nil, err
		}

		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))
		return versionMap, err
	}
}

// transfer all funds from protocolpool module account to distirbution module account
func CustomV024PatchLogic(ctx sdk.Context, k *keepers.AppKeepers) error {
	NewFactoryAddress, err := InstantiateNewFactory(ctx, k)
	if err != nil {
		return err
	}
	if err := MigrateExistingContracts(ctx, k, NewFactoryAddress); err != nil {
		return err
	}
	return nil
}

// // transfer all funds from protocolpool module account to distirbution module account
func InstantiateNewFactory(ctx sdk.Context, k *keepers.AppKeepers) (sdk.AccAddress, error) {
	NewFactoryAddress, _, err := k.ContractKeeper.Instantiate(ctx, NewFactoryCodeId, k.GovKeeper.ModuleAccountAddress(), sdk.MustAccAddressFromBech32(ContractAdmin), []byte{}, "", nil)
	if err != nil {
		return nil, err

	}
	return NewFactoryAddress, nil
}

// iterate through all bs721's, retrieve necessary info to
func MigrateExistingContracts(ctx sdk.Context, k *keepers.AppKeepers, NewFactoryAddress sdk.AccAddress) error {
	var gracefulErr error
	// for each collection, we need to:
	k.WasmKeeper.IterateContractsByCode(ctx, OldBs721CodeId, func(bs721 sdk.AccAddress) bool {
		govModAddress := k.GovKeeper.ModuleAccountAddress()
		// old variables
		var bs721Minter sdk.AccAddress
		var oldBs721Ownership ContractOwnership

		// new variables
		var newBs721CurveInitMsg Bs721CurveInitMsg

		// retrive the minter, nft_info, & ownership from contract state
		k.WasmKeeper.IterateContractState(ctx, bs721, func(key, value []byte) bool {
			switch {
			case bytes.Equal(key, []byte("minter")):
				bs721Minter = sdk.MustAccAddressFromBech32(string(value))
			case bytes.Equal(key, []byte("ownership")):
				json.Unmarshal(value, &oldBs721Ownership)
			// case bytes.Equal(key, []byte("nft_info")):
			// case bytes.Equal(key, []byte("contract_info")):
			default:
				fmt.Println("key not critical, skipping:", key)
			}
			return false
		})

		// we expect the minter to be bs721 curve contract, so retrive _ from its contract state.
		oldCurveCInfo := k.WasmKeeper.QueryRaw(ctx, bs721Minter, []byte("contract_info"))
		if oldCurveCInfo != nil {
			var actualContractInfo RawContractInfo
			err := json.Unmarshal(oldCurveCInfo, &actualContractInfo)
			if err != nil {
				gracefulErr = err
				return true
			}
			if actualContractInfo.Contract != "crates.io:bs721-curve" {
				gracefulErr = fmt.Errorf("bs721-contract minter is not a bs721-curve contract")
				return true
			}
			// we know this contract is bs721-curve contract. retrive the original init-msg from contract-history in preparation to create new contract
			oldCurveCHistory := k.WasmKeeper.GetContractHistory(ctx, bs721Minter)[0]
			oldCurveCInfo := k.WasmKeeper.GetContractInfo(ctx, bs721Minter)

			if oldCurveCHistory.Operation == wasmtypes.ContractCodeHistoryOperationTypeInit {
				var sellerBps int
				var referralBps int
				var protocolBps int
				var paymentDenom string
				var paymentAddress string
				json.Unmarshal(oldCurveCHistory.Msg, &newBs721CurveInitMsg)
				// update to make use of new code-ids
				newBs721CurveInitMsg.Bs721Admin = oldBs721Ownership.Owner
				newBs721CurveInitMsg.Bs721CodeID = NewBs721CodeId

				sellerBps = newBs721CurveInitMsg.SellerFeeBps
				referralBps = newBs721CurveInitMsg.ReferralFeeBps
				protocolBps = newBs721CurveInitMsg.ProtocolFeeBps
				paymentAddress = newBs721CurveInitMsg.PaymentAddress

				// set to values used only during migration. These will be updated after existing tokens are reminted
				newBs721CurveInitMsg.SellerFeeBps = 0
				newBs721CurveInitMsg.ReferralFeeBps = 0
				newBs721CurveInitMsg.ProtocolFeeBps = 0
				newBs721CurveInitMsg.PaymentAddress = govModAddress.String()
				newBs721CurveInitMsg.PaymentDenom = "ubtsg"
				newBs721CurveInitMsg.Ratio = 1

				// instantiate new bs721curve
				bs721CurveInitBz, _ := json.Marshal(newBs721CurveInitMsg)
				newBs721CurveAddr, _, err := k.ContractKeeper.Instantiate(ctx, NewBs721CurveCodeId, govModAddress, oldCurveCInfo.AdminAddr(), bs721CurveInitBz, oldCurveCInfo.Label, nil)
				if err != nil {
					gracefulErr = err
					return true
				}

				// iterate through old bs721curve state for all ``.
				k.WasmKeeper.IterateContractState(ctx, bs721Minter, func(key, value []byte) bool {
					var price QueryTotalMintPriceResponse
					var mintMsg Mint
					// Define the prefix you're interested in
					prefix := []byte("address_tokens")
					// Check if the key starts with the prefix
					if bytes.HasPrefix(key, prefix) {
						address := string(key[len(prefix):])
						// check how many nfts we should mint
						count := binary.LittleEndian.Uint32(value)

						// we wont save addresses that have 0 minted in store
						if count != 0 {
							queryBz, _ := json.Marshal(QueryTotalMintPrice{
								Amount: int(count),
							})
							// get the total purchase price
							qRes, _ := k.WasmKeeper.QuerySmart(ctx, newBs721CurveAddr, queryBz)
							json.Unmarshal(qRes, price)

							// form mintmsg
							mintMsg.Amount = count
							mintMsg.MintTo = address
							mintMsg.Referral = govModAddress.String()

							mintMsgBz, _ := json.Marshal(mintMsg)

							// mint the number of tokens to this address. All minted funds are expected to go right back to govModAddress
							_, err := k.ContractKeeper.Execute(ctx, newBs721CurveAddr, govModAddress, mintMsgBz, sdk.NewCoins(sdk.NewCoin("ubtsg", math.NewIntFromUint64(uint64(price.PubTotalPrice)))))
							if err != nil {
								gracefulErr = err
								return true
							}
						}
					}

					// if no errors during remint, continue to next nft owner in store
					if gracefulErr != nil {
						return true
					}
					return false
				})
				// if err exist, dont continue
				if gracefulErr != nil {
					return true
				}
				// once all holders have been minted new nft, lets update params back to normal
				var updateConfig UpdateConfig
				updateConfig.Cfg.PaymentAddress = paymentAddress
				updateConfig.Cfg.PaymentDenom = paymentDenom
				updateConfig.Cfg.ProtocolFeeBps = uint32(protocolBps)
				updateConfig.Cfg.ReferralFeeBps = uint32(referralBps)
				updateConfig.Cfg.SellerFeeBps = uint32(sellerBps)
				updateConfigBz, _ := json.Marshal(updateConfig)
				// execute as gov mod
				_, err = k.ContractKeeper.Execute(ctx, newBs721CurveAddr, govModAddress, updateConfigBz, nil)
				if err != nil {
					gracefulErr = err
					return true
				}
			}

		}
		return false
	})

	if gracefulErr != nil {
		return gracefulErr
	}

	return nil
}
