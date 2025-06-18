package authenticator

import (
	"bytes"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	btsgblst "github.com/bitsongofficial/go-bitsong/crypto/bls/blst"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/common"
	"github.com/bitsongofficial/go-bitsong/x/smart-account/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Bls12381 authenticates aggregate signatures from an set of public keys registered.
// It allows for complex pattern matching to support advanced authentication flows.
type Bls12381 struct {
	am       *AuthenticatorManager
	cdc      codec.Codec
	storeKey storetypes.StoreKey
}

var _ Authenticator = &Bls12381{}

func NewBls12381(am *AuthenticatorManager, storeKey storetypes.StoreKey) Bls12381 {
	return Bls12381{
		am:       am,
		storeKey: storeKey,
	}
}

func (bls Bls12381) Type() string {
	return "Bls12381"
}

func (bls Bls12381) StaticGas() uint64 {
	return 0
}

func (bls Bls12381) Initialize(cfg []byte) (Authenticator, error) {
	return bls, nil
}

func (bls Bls12381) Authenticate(ctx sdk.Context, req AuthenticationRequest) error {
	// Validate input
	// fmt.Printf("req.AuthenticatorId: %v\n", req.AuthenticatorId)
	// fmt.Printf("len(req.SignatureData.Signatures): %v\n", len(req.SignatureData.Signatures))
	// fmt.Printf("len(req.SignatureData.Signers): %v\n", len(req.SignatureData.Signers))
	// ensure threshold is met & keys provided are expected for this authenticator
	var blsConfig types.BlsConfig
	store := ctx.KVStore(bls.storeKey)
	key := types.KeyAccountBlsKeySet(req.Account, req.AuthenticatorId)
	found, err := types.Get(store, key, &blsConfig)
	if err != nil || !found {
		return errorsmod.Wrap(err, "failed to get authenticator")
	}
	// ensure threshold has been met
	if blsConfig.Threshold+1 < uint64(len(req.SignatureData.Signatures)) {
		return errorsmod.Wrap(err, "aggregate signature threshold not satisfied")

	}

	msgDigestHash := Sha256Msgs(req.TxData.Msgs)
	// fmt.Printf("msgDigestHash: %v\n", msgDigestHash)

	// Aggregate public keys
	var g1 [][]byte
	// first sig details is ALWAYS aggregated key, so we skip
	for i, signer := range req.SignatureData.Signers[1:] {
		validPoint := checkPubkeyExistence(&blsConfig, signer)
		if !validPoint {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "aggregate key is not valid point %d: %v", i, err)
		}
		// pk, _ := blst.PublicKeyFromBytes(signer)
		// good, err := btsgblst.VerifySignature(req.SignatureData.Signatures[i+1], msgDigestHash, pk)
		// fmt.Printf("good: %v\n", good)
		// if err != nil || !good {
		// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "AHHHHHHHH: %v", err)
		// }
		// Aggregate public keys (add them in G1)
		g1 = append(g1, signer)
	}

	var g2 []common.Signature
	for i, sigBytes := range req.SignatureData.Signatures[1:] {
		if len(sigBytes) == 0 {
			continue
		}

		sig, err := btsgblst.SignatureFromBytes(sigBytes)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to deserialize wavs operator signature at index %d: %v", i, err)
		}
		g2 = append(g2, sig)
	}

	// Aggregate Signature
	// aggregate signature is in default signature location
	aggregatedSignature := btsgblst.AggregateSignatures(g2)
	aggregatedPubkey, err := btsgblst.AggregatePublicKeys(g1)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Aggregated Signature Failed: %v", err)
	}

	// providedAggregateSignature, err := btsgblst.SignatureFromBytesNoValidation(req.Signature)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed btsgblst.SignatureFromBytesNoValidatio: %v", err)
	// }
	// // fmt.Printf("aggregatedSignature: %v\n", aggregatedSignature.Marshal())
	// // fmt.Printf("providedAggregateSignature.Marshal(): %v\n", providedAggregateSignature.Marshal())
	// if providedAggregateSignature != aggregatedSignature {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Aggregated Signature Failed: %v", "computed aggregate signature does not match provided aggregate signature")
	// }

	valid := aggregatedSignature.Verify(aggregatedPubkey, msgDigestHash[:])
	if !valid {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "btsgblst.VerifySignature Failed: %v", "aggregate signature verification failed")
	}

	return nil
}

func (bls Bls12381) Track(ctx sdk.Context, request AuthenticationRequest) error {
	return nil
	// return subTrack(ctx, request, bls.SubAuthenticators)
}

func (bls Bls12381) ConfirmExecution(ctx sdk.Context, request AuthenticationRequest) error {
	return nil
}

func (bls Bls12381) OnAuthenticatorAdded(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	return onBls12381Added(ctx, bls.storeKey, account, config, authenticatorId, bls.am)
}

func (bls Bls12381) OnAuthenticatorRemoved(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	return onBls12381Removed(ctx, bls.storeKey, account, authenticatorId)
}

func onBls12381Added(ctx sdk.Context, storekey storetypes.StoreKey, account sdk.AccAddress, data []byte, authenticatorId string, am *AuthenticatorManager) error {
	var config types.BlsConfig
	if err := config.Unmarshal(data); err != nil {
		return errorsmod.Wrapf(err, "failed to unmarshal BlsConfig init data")
	}

	if len(config.Pubkeys) < int(config.Threshold) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "must set threshold to atleast to the number of keys: %d", config.Threshold)
	}

	key := types.KeyAccountBlsKeySet(account, authenticatorId)
	types.MustSet(ctx.KVStore(storekey), key, &config)
	return nil

}
func onBls12381Removed(ctx sdk.Context, storekey storetypes.StoreKey, account sdk.AccAddress, authenticatorId string) error {
	key := types.KeyAccountBlsKeySet(account, authenticatorId)
	ctx.KVStore(storekey).Delete(key)
	return nil
}

func checkPubkeyExistence(blsConfig *types.BlsConfig, pubkey []byte) bool {
	for _, pk := range blsConfig.Pubkeys {
		if bytes.Equal(pk, pubkey) {
			return true
		}
	}
	return false
}
