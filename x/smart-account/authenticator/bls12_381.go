package authenticator

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	blst "github.com/bitsongofficial/go-bitsong/crypto/bls/blst"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type Bls12 struct {
	SubAuthenticators []Authenticator
	am                *AuthenticatorManager
	// signatureAssignment SignatureAssignment
}

var _ Authenticator = &Bls12{}

func NewBls12(am *AuthenticatorManager) Bls12 {
	return Bls12{
		am:                am,
		SubAuthenticators: []Authenticator{},
		// signatureAssignment: Single,
	}
}

func NewPartitionedBls12(am *AuthenticatorManager) Bls12 {
	return Bls12{
		am:                am,
		SubAuthenticators: []Authenticator{},
	}
}

func (bls Bls12) Type() string {

	return "Bls12"
}

func (bls Bls12) StaticGas() uint64 {
	var totalGas uint64
	// for _, auth := range bls.SubAuthenticators {
	// 	totalGas += auth.StaticGas()
	// }
	return totalGas
}

func (bls Bls12) Initialize(config []byte) (Authenticator, error) {
	var initDatas []SubAuthenticatorInitData
	if err := json.Unmarshal(config, &initDatas); err != nil {
		return nil, errorsmod.Wrap(err, "failed to parse sub-authenticators initialization data")
	}

	return bls, nil
}

func (bls Bls12) Authenticate(ctx sdk.Context, request AuthenticationRequest) error {
	// Validate input
	if len(request.SignatureData.Signers) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no public keys provided")
	}
	// expect sha256sum signed to be provided here.
	// We dont perform hash client side to save costs, but can be implemented.
	if len(request.SignModeTxData.Direct) != 32 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "sha256sum hash signed by AVS set not provided")
	}
	if len(request.Signature) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no aggregated signature provided")
	}

	b1Points := request.SignatureData.Signers
	b2Points := request.SignatureData.Signatures
	msgDigestHash := request.SignModeTxData.Direct
	providedAggregateSignature, err := blst.SignatureFromBytesNoValidation(request.Signature)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed blst.SignatureFromBytesNoValidatio: %v", err)
	}
	// Aggregate public keys
	var aggPubKeys [][]byte
	for i, pubKeyBytes := range b1Points {
		if len(pubKeyBytes) == 0 {
			continue
		}
		pubKey, err := blst.PublicKeyFromBytes(pubKeyBytes)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to deserialize public key at index %d: %v", i, err)
		}

		// Aggregate public keys (add them in G1)
		aggPubKeys = append(aggPubKeys, pubKey.Marshal())
	}

	var aggSigArray []common.Signature
	for i, signatureBytes := range b2Points {
		if len(signatureBytes) == 0 {
			continue
		}

		sig, err := blst.SignatureFromBytes(signatureBytes)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to deserialize wavs operator signature at index %d: %v", i, err)
		}
		aggSigArray = append(aggSigArray, sig)
	}
	// Aggregate Signature
	aggregatedSignature := blst.AggregateSignatures(aggSigArray)
	aggregatedPubkey, err := blst.AggregatePublicKeys(aggPubKeys)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Aggregated Signature Failed: %v", err)
	}
	if providedAggregateSignature != aggregatedSignature {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Aggregated Signature Failed: %v", err)
	}

	// digest msg hash that was signed
	var digest [32]byte
	copy(digest[:], msgDigestHash)

	verified, err := blst.VerifySignature(providedAggregateSignature.Marshal(), digest, aggregatedPubkey)
	if err != nil || !verified {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "blst.VerifySignature Failed: %v", err)
	}

	return nil
}

func (bls Bls12) Track(ctx sdk.Context, request AuthenticationRequest) error {
	return subTrack(ctx, request, bls.SubAuthenticators)
}

func (bls Bls12) ConfirmExecution(ctx sdk.Context, request AuthenticationRequest) error {

	return nil
}

func (bls Bls12) OnAuthenticatorAdded(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	return onSubAuthenticatorsAdded(ctx, account, config, authenticatorId, bls.am)
}

func (bls Bls12) OnAuthenticatorRemoved(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	return onSubAuthenticatorsRemoved(ctx, account, config, authenticatorId, bls.am)
}
