package authenticator

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	blst "github.com/Layr-Labs/eigensdk-go/crypto/bls"
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
		// signatureAssignment: Partitioned,
	}
}

func (bls Bls12) Type() string {
	// if bls.signatureAssignment == Single {
	// 	return "Bls12"
	// }
	// return "PartitionedBls12"
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

	// if len(initDatas) <= 1 {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "allOf must have at least 2 sub-authenticators")
	// }

	// for _, initData := range initDatas {
	// 	authenticatorCode := bls.am.GetAuthenticatorByType(initData.Type)
	// 	instance, err := authenticatorCode.Initialize(initData.Config)
	// 	if err != nil {
	// 		return nil, errorsmod.Wrapf(err, "failed to initialize sub-authenticator (type = %s)", initData.Type)
	// 	}
	// 	bls.SubAuthenticators = append(bls.SubAuthenticators, instance)
	// }

	// // If not all sub-authenticators are registered, return an error
	// if len(bls.SubAuthenticators) != len(initDatas) {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to initialize all sub-authenticators")
	// }

	return bls, nil
}

func (bls Bls12) Authenticate(ctx sdk.Context, request AuthenticationRequest) error {
	// Validate input
	if len(request.SignatureData.Signers) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no public keys provided")
	}
	if len(request.SignModeTxData.Direct) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no message hash provided")
	}
	if len(request.Signature) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no aggregated signature provided")
	}

	b1Points := request.SignatureData.Signers
	b2Points := request.SignatureData.Signatures
	// sha256MsgHash := request.SignModeTxData.Direct

	aggregateSignature := request.Signature

	// Initialize aggregated public key (G1 point)
	var aggPubKey blst.G1Point
	isFirstKey := true

	for i, pubKeyBytes := range b1Points {
		if len(pubKeyBytes) == 0 {
			continue
		}
		var pubKey blst.G1Point
		if err := pubKey.Deserialize(pubKeyBytes); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to deserialize public key at index %d: %v", i, err)
		}

		// Aggregate public keys (add them in G1)
		if isFirstKey {
			aggPubKey = pubKey
			isFirstKey = false
		} else {
			aggPubKey.Add(&pubKey)
		}
	}

	// Deserialize the aggregated signature (G2 point)
	var aggSig blst.G2Point
	isFirstSig := true
	for i, signatureBytes := range b2Points {
		if len(signatureBytes) == 0 {
			continue
		}
		var wavsOpsSig blst.G2Point
		if err := wavsOpsSig.Deserialize(signatureBytes); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to deserialize wavs operator signature at index %d: %v", i, err)
		}

		// Aggregate public keys (add them in G1)
		if isFirstSig {
			aggSig = wavsOpsSig
			isFirstSig = false
		} else {
			aggSig.Add(&wavsOpsSig)
		}
	}

	if err := aggSig.Deserialize(aggregateSignature); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to deserialize aggregated signature: "+err.String())
	}

	// TODO: Verify the aggregated signature against the aggregated public key and message hash

	return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "aggregated BLS signature verification failed")

	// return nil
}

func (bls Bls12) Track(ctx sdk.Context, request AuthenticationRequest) error {
	return subTrack(ctx, request, bls.SubAuthenticators)
}

func (bls Bls12) ConfirmExecution(ctx sdk.Context, request AuthenticationRequest) error {
	// var signatures [][]byte
	// var err error
	// if bls.signatureAssignment == Partitioned {
	// 	// Partitioned signatures are decoded and passed one by one as the signature of the sub-authenticator
	// 	signatures, err = splitSignatures(request.Signature, len(bls.SubAuthenticators))
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// baseId := request.AuthenticatorId
	// for i, auth := range bls.SubAuthenticators {
	// 	// update the authenticator id to include the sub-authenticator id
	// 	request.AuthenticatorId = compositeId(baseId, i)
	// 	// update the request to include the sub-authenticator signature
	// 	// if bls.signatureAssignment == Partitioned {
	// 	// 	request.Signature = signatures[i]
	// 	// }

	// 	if err := auth.ConfirmExecution(ctx, request); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (bls Bls12) OnAuthenticatorAdded(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	return onSubAuthenticatorsAdded(ctx, account, config, authenticatorId, bls.am)
}

func (bls Bls12) OnAuthenticatorRemoved(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	return onSubAuthenticatorsRemoved(ctx, account, config, authenticatorId, bls.am)
}
