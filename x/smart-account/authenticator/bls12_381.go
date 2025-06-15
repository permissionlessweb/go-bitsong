package authenticator

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	btsgblst "github.com/bitsongofficial/go-bitsong/crypto/bls/blst"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/common"
	"github.com/bitsongofficial/go-bitsong/x/smart-account/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Bls12381 authenticates aggregate signatures from an set of public keys registered.
// It allows for complex pattern matching to support advanced authentication flows.
type Bls12381 struct {
	am       *AuthenticatorManager
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
	fmt.Printf("req.SignatureData: %v\n", req.SignatureData)
	fmt.Printf("req.AuthenticatorId: %v\n", req.AuthenticatorId)
	if len(req.SignatureData.Signatures) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no public keys provided")
	}
	fmt.Printf("req.TxData.Msgs: %v\n", req.TxData.Msgs)
	store := ctx.KVStore(bls.storeKey)
	var blsConfig types.BlsConfig
	key := types.KeyAccountBlsKeySet(req.Account, req.AuthenticatorId)
	fmt.Printf("key: %v\n", key)
	found, err := types.Get(store, key, &blsConfig)
	if err != nil {
		return errorsmod.Wrap(err, "failed to get authenticator")
	}
	if !found {
		return fmt.Errorf("could not get key accunt by id & account")
	}

	// TODO: ensure keys provided are in expected keys array
	// blsConfig.Pubkeys

	msgDigestHash := Sha256Msgs(req.TxData.Msgs)
	fmt.Printf("msgDigestHash: %v\n", msgDigestHash)

	providedAggregateSignature, err := btsgblst.SignatureFromBytesNoValidation(req.Signature)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed btsgblst.SignatureFromBytesNoValidatio: %v", err)
	}
	// Aggregate public keys
	var aggb1 [][]byte
	for i, pubKeyBytes := range req.SignatureData.Signers {
		if len(pubKeyBytes) == 0 {
			continue
		}

		fmt.Printf("len(pubKeyBytes): %v\n", len(pubKeyBytes))
		pubKey, err := btsgblst.PublicKeyFromBytes(pubKeyBytes.Bytes())
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to deserialize public key at index %d: %v", i, err)
		}

		// Aggregate public keys (add them in G1)
		aggb1 = append(aggb1, pubKey.Marshal())
	}

	var aggb2 []common.Signature
	for i, signatureBytes := range req.SignatureData.Signatures {
		if len(signatureBytes) == 0 {
			continue
		}

		sig, err := btsgblst.SignatureFromBytes(signatureBytes)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to deserialize wavs operator signature at index %d: %v", i, err)
		}
		aggb2 = append(aggb2, sig)
	}
	// Aggregate Signature
	aggregatedSignature := btsgblst.AggregateSignatures(aggb2)
	aggregatedPubkey, err := btsgblst.AggregatePublicKeys(aggb1)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Aggregated Signature Failed: %v", err)
	}
	if providedAggregateSignature != aggregatedSignature {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Aggregated Signature Failed: %v", err)
	}

	verified, err := btsgblst.VerifySignature(providedAggregateSignature.Marshal(), msgDigestHash, aggregatedPubkey)
	if err != nil || !verified {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "btsgblst.VerifySignature Failed: %v", err)
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
	return onSubAuthenticatorsRemoved(ctx, account, config, authenticatorId, bls.am)
}

func onBls12381Added(ctx sdk.Context, storekey storetypes.StoreKey, account sdk.AccAddress, data []byte, authenticatorId string, am *AuthenticatorManager) error {
	var config types.BlsConfig
	if err := config.Unmarshal(data); err != nil {
		return errorsmod.Wrapf(err, "failed to unmarshal BlsConfig init data")
	}

	if len(config.Pubkeys) < int(config.Threshold) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "must set threshold to atleast to the number of keys, but got %d", config.Threshold)
	}

	fmt.Printf("config: %v\n", config)
	key := types.KeyAccountBlsKeySet(account, authenticatorId)
	fmt.Printf("key: %v\n", key)
	types.MustSet(ctx.KVStore(storekey), key, &config)

	// If not all sub-authenticators are registered, return an error

	return nil
}
