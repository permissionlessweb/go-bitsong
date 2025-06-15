package authenticator

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/bls12381"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	sat "github.com/bitsongofficial/go-bitsong/x/smart-account/types"

	txsigning "cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//
// These structs define the data structure for authentication, used with AuthenticationRequest struct.
//

// SignModeData represents the signing modes with direct bytes and textual representation.
type SignModeData struct {
	Direct  []byte `json:"sign_mode_direct"`
	Textual string `json:"sign_mode_textual"`
}

// LocalAny holds a message with its type URL and byte value. This is necessary because the type Any fails
// to serialize and deserialize properly in nested contexts.
type LocalAny struct {
	TypeURL string `json:"type_url"`
	Value   []byte `json:"value"`
}

// SimplifiedSignatureData contains lists of signers and their corresponding signatures.
type SimplifiedSignatureData struct {
	Signers    []sdk.AccAddress `json:"signers"`
	Signatures [][]byte         `json:"signatures"`
}

// ExplicitTxData encapsulates key transaction data like chain ID, account info, and messages.
type ExplicitTxData struct {
	ChainID         string     `json:"chain_id"`
	AccountNumber   uint64     `json:"account_number"`
	AccountSequence uint64     `json:"sequence"`
	TimeoutHeight   uint64     `json:"timeout_height"`
	Msgs            []LocalAny `json:"msgs"`
	Memo            string     `json:"memo"`
}

// GetSignerAndSignatures gets an array of signer and an array of signatures from the transaction
// checks they're the same length and returns both.
//
// A signer can only have one signature, so if it appears in multiple messages, the signatures must be
// the same, and it will only be returned once by this function. This is to mimic the way the classic
// sdk authentication works, and we will probably want to change this in the future
func GetSignerAndSignatures(tx sdk.Tx, aggSig *sat.AgAuthData) (signers []sdk.AccAddress, signatures []signing.SignatureV2, err error) {
	// Attempt to cast the provided transaction to an authsigning.Tx.
	sigTx, ok := tx.(authsigning.Tx)
	if !ok {
		return nil, nil,
			errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// Retrieve signatures from the transaction.
	signatures, err = sigTx.GetSignaturesV2()
	if err != nil {
		return nil, nil, err
	}

	// Retrieve messages from the transaction.
	signerBytes, err := sigTx.GetSigners()
	if err != nil {
		return nil, nil, err
	}

	if aggSig != nil {
		// static accAddress for account keys that registered agg authenticator.
		// Used by all keys in agg key group, we always identify agg keys by their pubkeys .
		signers = append(signers, sdk.AccAddress(signerBytes[0]))
		// we expect one signature to have been validated already
		return signers, signatures, nil
	}

	for _, signer := range signerBytes {
		signers = append(signers, sdk.AccAddress(signer))
	}
	// check that signer length and signature length are the same
	if len(signatures) != len(signers) {
		return nil, nil,
			errorsmod.Wrap(sdkerrors.ErrTxDecode, fmt.Sprintf("invalid number of signer;  expected: %d, got %d", len(signers), len(signatures)))
	}

	return signers, signatures, nil
}

// getSignerData returns the signer data for a given account. This is part of the data that needs to be signed.
func getSignerData(ctx sdk.Context, ak authante.AccountKeeper, account sdk.AccAddress) authsigning.SignerData {
	// Retrieve and build the signer data struct
	baseAccount := ak.GetAccount(ctx, account)
	genesis := ctx.BlockHeight() == 0
	chainID := ctx.ChainID()
	var accNum uint64
	if !genesis {
		accNum = baseAccount.GetAccountNumber()
	}
	var sequence uint64
	if baseAccount != nil {
		sequence = baseAccount.GetSequence()
	}

	return authsigning.SignerData{
		ChainID:       chainID,
		AccountNumber: accNum,
		Sequence:      sequence,
	}
}

// extractExplicitTxData makes the transaction data concrete for the authentication request. This is necessary to
// pass the parsed data to the cosmwasm authenticator.
func extractExplicitTxData(tx sdk.Tx, signerData authsigning.SignerData) (ExplicitTxData, error) {
	timeoutTx, ok := tx.(sdk.TxWithTimeoutHeight)
	if !ok {
		return ExplicitTxData{}, errorsmod.Wrap(sdkerrors.ErrInvalidType, "failed to cast tx to TxWithTimeoutHeight")
	}
	memoTx, ok := tx.(sdk.TxWithMemo)
	if !ok {
		return ExplicitTxData{}, errorsmod.Wrap(sdkerrors.ErrInvalidType, "failed to cast tx to TxWithMemo")
	}

	// Encode messages as Anys and manually convert them to a struct we can serialize to json for cosmwasm.
	txMsgs := tx.GetMsgs()
	msgs := make([]LocalAny, len(txMsgs))
	for i, txMsg := range txMsgs {
		encodedMsg, err := codectypes.NewAnyWithValue(txMsg)
		if err != nil {
			return ExplicitTxData{}, errorsmod.Wrap(err, "failed to encode msg")
		}
		msgs[i] = LocalAny{
			TypeURL: encodedMsg.TypeUrl,
			Value:   encodedMsg.Value,
		}
	}

	return ExplicitTxData{
		ChainID:         signerData.ChainID,
		AccountNumber:   signerData.AccountNumber,
		AccountSequence: signerData.Sequence,
		TimeoutHeight:   timeoutTx.GetTimeoutHeight(),
		Msgs:            msgs,
		Memo:            memoTx.GetMemo(),
	}, nil
}

// extractSignatures returns the signature data for each signature in the transaction and the one for the current signer.
//
// This function also checks for replay attacks. The replay protection needs to be able to match the signature to the
// corresponding signer, which involves iterating over the signatures. To avoid iterating over the signatures twice,
// we do replay protection here instead of in a separate replay protection function.
//
// If this tx is making use of Aggregated Signatures,we optionally expect a single aggregated pk & sig, or else we return nothing.
func extractSignatures(cdc codec.Codec, txSigners []sdk.AccAddress, txSignatures []signing.SignatureV2, txData ExplicitTxData, account sdk.AccAddress, replayProtection ReplayProtection, agAuthData *sat.AgAuthData) (signatures [][]byte, msgSignature []byte, err error) {
	if agAuthData != nil {
		// check if an aggregated signature was generated and provided offline
		if len(txSignatures) > 0 {
			// set agg key & sig first
			aggregatedSig := txSignatures[0]
			single, ok := aggregatedSig.Data.(*signing.SingleSignatureData)
			if !ok {
				return nil, nil, errorsmod.Wrap(sdkerrors.ErrInvalidType, "failed to cast aggregated signature to SingleSignatureData")
			}
			fmt.Printf("single: %v\n", single)
			if replayProtection(&txData, &aggregatedSig); err != nil {
				return nil, nil, err
			}
			signatures = append(signatures, single.Signature)
		} else {
			return nil, nil, errorsmod.Wrap(sdkerrors.ErrInvalidType, "no tx signatures provided")
		}

		rawSigs, err := UnmarshalSignatureJSON(cdc, agAuthData.Data)
		if err != nil {
			return nil, nil, errorsmod.Wrap(sdkerrors.ErrInvalidType, "failed to UnmarshalSignatureJSON")
		}

		for i, extSig := range rawSigs {
			single, ok := extSig.Data.(*signing.SingleSignatureData)
			if !ok {
				return nil, nil, errorsmod.Wrap(sdkerrors.ErrInvalidType, "failed to cast extTx signature to SingleSignatureData")
			}
			// fmt.Printf("single: %v\n", single)
			signatures = append(signatures, single.Signature)
			if txSigners[i].Equals(account) {
				err = replayProtection(&txData, &extSig)
				if err != nil {
					return nil, nil, err
				}
				msgSignature = single.Signature
			}
		}
		return signatures, msgSignature, nil
	}

	for i, signature := range txSignatures {
		single, ok := signature.Data.(*signing.SingleSignatureData)
		if !ok {
			return nil, nil, errorsmod.Wrap(sdkerrors.ErrInvalidType, "failed to cast signature to SingleSignatureData")
		}

		// fmt.Printf("single: %v\n", single)
		signatures = append(signatures, single.Signature)

		if txSigners[i].Equals(account) {
			err = replayProtection(&txData, &signature)
			if err != nil {
				return nil, nil, err
			}
			msgSignature = single.Signature
		}
	}
	return signatures, msgSignature, nil
}

// GenerateAuthenticationRequest creates an AuthenticationRequest for the transaction.
func GenerateAuthenticationRequest(
	ctx sdk.Context,
	cdc codec.Codec,
	ak authante.AccountKeeper,
	sigModeHandler *txsigning.HandlerMap,
	account sdk.AccAddress,
	feePayer sdk.AccAddress,
	feeGranter sdk.AccAddress,
	fee sdk.Coins,
	msg sdk.Msg,
	tx sdk.Tx,
	msgIndex int,
	simulate bool,
	replayProtection ReplayProtection,
	agAuthData *sat.AgAuthData,
) (AuthenticationRequest, error) {
	var aggEnabled bool
	var simpleSignatureData = SimplifiedSignatureData{
		Signers:    make([]sdk.AccAddress, 0),
		Signatures: make([][]byte, 0),
	}
	// Only supporting one signer per message in default signer data
	signers, _, err := cdc.GetMsgV1Signers(msg)
	if err != nil {
		return AuthenticationRequest{}, err
	}
	fmt.Printf("signers: %v\n", signers)

	// either actual signer, or aggregated pubkey & address of account agg pubkeys control.
	signer := sdk.AccAddress(signers[0])
	if !signer.Equals(account) {
		return AuthenticationRequest{}, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid signer")
	}
	// Check if this request is using aggregated consensus authentication
	if agAuthData != nil {
		if len(agAuthData.Data) > 0 {
			aggEnabled = true
		}
	}
	fmt.Printf("aggEnabled: %v\n", aggEnabled)

	// Get the signers and signatures from the transaction.
	txSigners, txSignatures, err := GetSignerAndSignatures(tx, agAuthData)
	if err != nil {
		return AuthenticationRequest{}, errorsmod.Wrap(err, "failed to get signers and signatures")
	}
	// Get the signer data for the account. This is needed in the SignDoc
	signerData := getSignerData(ctx, ak, account)

	fmt.Printf("txSigners: %v\n", txSigners)
	fmt.Printf("txSignatures: %v\n", txSignatures)
	fmt.Printf("signerData: %v\n", signerData)

	// Get the concrete transaction data to be passed to the authenticators
	txData, err := extractExplicitTxData(tx, signerData)
	if err != nil {
		return AuthenticationRequest{}, errorsmod.Wrap(err, "failed to get explicit tx data")
	}
	// fmt.Printf("txData: %v\n", txData)

	// Get the signatures for the transaction and execute replay protection.
	// If aggregate keys are in use, set agg key & sig as first value, followed by all key/sig pairs in extension
	signatures, msgSignature, err := extractSignatures(cdc, txSigners, txSignatures, txData, account, replayProtection, agAuthData)
	if err != nil {
		return AuthenticationRequest{}, errorsmod.Wrap(err, "failed to get signatures")
	}

	simpleSignatureData.Signatures = append(simpleSignatureData.Signatures, signatures...)
	simpleSignatureData.Signers = append(simpleSignatureData.Signers, signer)

	// Build the authentication request
	authRequest := AuthenticationRequest{
		Account:    account,
		FeePayer:   feePayer,
		FeeGranter: feeGranter,
		Fee:        fee,
		Msg:        txData.Msgs[msgIndex],
		MsgIndex:   uint64(msgIndex),
		Signature:  msgSignature,
		TxData:     txData,
		SignModeTxData: SignModeData{
			Direct: []byte("signBytes"),
		},
		SignatureData:       simpleSignatureData,
		Simulate:            simulate,
		AuthenticatorParams: nil,
	}

	// We do not generate the sign bytes if simulate is true or isCheckTx is true
	if simulate && ctx.IsCheckTx() {
		return authRequest, nil
	}

	// Get the sign bytes for the transaction
	signBytes, err := authsigning.GetSignBytesAdapter(ctx, sigModeHandler, signing.SignMode_SIGN_MODE_DIRECT, signerData, tx)
	if err != nil {
		return AuthenticationRequest{}, errorsmod.Wrap(err, "failed to get signBytes")
	}

	// TODO: Add other sign modes. Specifically json when it becomes available
	authRequest.SignModeTxData = SignModeData{
		Direct: signBytes,
	}
	return authRequest, nil
}

// Generates the SHA256SUM for an array of cosmos-sdk messages
func Sha256Msgs(msgs []LocalAny) [32]byte {
	jsonBytes, _ := json.Marshal(msgs)
	return sha256.Sum256(jsonBytes)
}

func MarshalSignatureJSON(sigs []signing.SignatureV2) ([]byte, error) {
	descs := make([]*signing.SignatureDescriptor, len(sigs))

	for i, sig := range sigs {
		descData := signing.SignatureDataToProto(sig.Data)
		// assert public key interface works
		pubKey, ok := sig.PubKey.(*bls12381.PubKey)
		if !ok {
			return nil, fmt.Errorf("failed to get bls12381.PubKey")
		}
		any, err := codectypes.NewAnyWithValue(pubKey)
		if err != nil {
			return nil, err
		}
		fmt.Printf("any: %v\n", any)
		descs[i] = &signing.SignatureDescriptor{
			PublicKey: any,
			Data:      descData,
			Sequence:  sig.Sequence,
		}
	}

	toJSON := &signing.SignatureDescriptors{Signatures: descs}

	return codec.ProtoMarshalJSON(toJSON, nil)
}

func UnmarshalSignatureJSON(cdc codec.Codec, bz []byte) ([]signing.SignatureV2, error) {
	var sigDescs signing.SignatureDescriptors
	err := cdc.UnmarshalJSON(bz, &sigDescs)
	if err != nil {
		return nil, err
	}

	sigs := make([]signing.SignatureV2, len(sigDescs.Signatures))
	for i, desc := range sigDescs.Signatures {
		pubKey, _ := desc.PublicKey.GetCachedValue().(cryptotypes.PubKey)

		data := signing.SignatureDataFromProto(desc.Data)

		sigs[i] = signing.SignatureV2{
			PubKey:   pubKey,
			Data:     data,
			Sequence: desc.Sequence,
		}
	}

	return sigs, nil
}
