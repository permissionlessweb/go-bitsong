package authenticator_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txsigning "cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/bitsongofficial/go-bitsong/app"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/blst"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/common"

	"github.com/bitsongofficial/go-bitsong/x/smart-account/authenticator"
	"github.com/bitsongofficial/go-bitsong/x/smart-account/types"
	smartaccounttypes "github.com/bitsongofficial/go-bitsong/x/smart-account/types"
)

type SigVerifyAuthenticationSuite struct {
	BaseAuthenticatorSuite

	SigVerificationAuthenticator authenticator.SignatureVerification
}

func TestSigVerifyAuthenticationSuite(t *testing.T) {
	suite.Run(t, new(SigVerifyAuthenticationSuite))
}

func (s *SigVerifyAuthenticationSuite) SetupTest() {
	s.SetupKeys()

	s.EncodingConfig = app.MakeEncodingConfig()
	ak := s.BitsongApp.AccountKeeper

	// Create a new Secp256k1SignatureAuthenticator for testing
	s.SigVerificationAuthenticator = authenticator.NewSignatureVerification(
		ak,
	)
}

func (s *SigVerifyAuthenticationSuite) TearDownTest() {
	os.RemoveAll(s.HomeDir)
}

type SignatureVerificationTestData struct {
	Msgs                               []sdk.Msg
	AccNums                            []uint64
	AccSeqs                            []uint64
	Signers                            []cryptotypes.PrivKey
	Signatures                         []cryptotypes.PrivKey
	NumberOfExpectedSigners            int
	NumberOfExpectedSignatures         int
	ShouldSucceedGettingData           bool
	ShouldSucceedSignatureVerification bool
}

type SignatureVerificationTest struct {
	Description string
	TestData    SignatureVerificationTestData
}

// TestSignatureAuthenticator test a non-smart account signature verification
func (s *SigVerifyAuthenticationSuite) TestSignatureAuthenticator() {
	bitsongToken := "bitsong"
	coins := sdk.Coins{sdk.NewInt64Coin(bitsongToken, 2500)}

	// Create a test messages for signing
	testMsg1 := &banktypes.MsgSend{
		FromAddress: sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[0]),
		ToAddress:   sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[1]),
		Amount:      coins,
	}
	testMsg2 := &banktypes.MsgSend{
		FromAddress: sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[1]),
		ToAddress:   sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[1]),
		Amount:      coins,
	}
	testMsg3 := &banktypes.MsgSend{
		FromAddress: sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[2]),
		ToAddress:   sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[1]),
		Amount:      coins,
	}
	//testMsg4 := &banktypes.MsgSend{
	//	FromAddress: sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[0]),
	//	ToAddress:   sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[1]),
	//	Amount:      coins,
	//}
	feeCoins := sdk.Coins{sdk.NewInt64Coin(bitsongToken, 2500)}

	tests := []SignatureVerificationTest{
		{
			Description: "Test: successfully verified authenticator with one signer: base case: PASS",
			TestData: SignatureVerificationTestData{
				[]sdk.Msg{
					testMsg1,
				},
				[]uint64{0},
				[]uint64{0},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
				},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
				},
				1,
				1,
				true,
				true,
			},
		},

		{
			Description: "Test: successfully verified authenticator: multiple signers: PASS",
			TestData: SignatureVerificationTestData{
				[]sdk.Msg{
					testMsg1,
					testMsg2,
					testMsg3,
				},
				[]uint64{0, 0, 0, 0},
				[]uint64{0, 0, 0, 0},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
					s.TestPrivKeys[1],
					s.TestPrivKeys[2],
				},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
					s.TestPrivKeys[1],
					s.TestPrivKeys[2],
				},
				3,
				3,
				true,
				true,
			},
		},

		{
			// This test case tests if there are two messages with the same signer
			// with two successful signatures.
			Description: "Test: verified authenticator with 2 messages signed correctly with the same address: PASS",
			TestData: SignatureVerificationTestData{
				[]sdk.Msg{
					testMsg1,
					testMsg2,
					testMsg2,
				},
				[]uint64{0, 0},
				[]uint64{0, 0},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
					s.TestPrivKeys[1],
				},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
					s.TestPrivKeys[1],
				},
				2,
				2,
				true,
				true,
			},
		},

		{
			// This test case tests if there are two messages with the same signer
			// with two successful signatures.
			Description: "Test: verified authenticator with 2 messages but only first signed signed correctly: Fail",
			TestData: SignatureVerificationTestData{
				[]sdk.Msg{
					testMsg1,
					testMsg2,
					testMsg2,
				},
				[]uint64{0, 0},
				[]uint64{0, 0},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
					s.TestPrivKeys[1],
				},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
					s.TestPrivKeys[0],
				},
				2,
				2,
				true,
				false,
			},
		},

		{
			// This test case tests if there are two messages with the same signer
			// with two successful signatures.
			Description: "Test: verified authenticator with 2 messages but only second signed signed correctly: Fail",
			TestData: SignatureVerificationTestData{
				[]sdk.Msg{
					testMsg1,
					testMsg2,
					testMsg2,
				},
				[]uint64{0, 0},
				[]uint64{0, 0},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[0],
					s.TestPrivKeys[1],
				},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[1],
					s.TestPrivKeys[1],
				},
				2,
				2,
				true,
				false,
			},
		},

		{
			Description: "Test: unsuccessful signature authentication invalid signatures: FAIL",
			TestData: SignatureVerificationTestData{
				[]sdk.Msg{
					testMsg1,
					testMsg2,
				},
				[]uint64{0, 0},
				[]uint64{0, 0},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[1],
					s.TestPrivKeys[0],
				},
				[]cryptotypes.PrivKey{
					s.TestPrivKeys[2],
					s.TestPrivKeys[0],
				},
				2,
				2,
				false,
				false,
			},
		},
	}

	for _, tc := range tests {
		s.Run(tc.Description, func() {
			// Generate a transaction based on the test cases
			tx, _ := GenTx(
				s.Ctx,
				s.EncodingConfig.TxConfig,
				tc.TestData.Msgs,
				feeCoins,
				300000,
				"",
				tc.TestData.AccNums,
				tc.TestData.AccSeqs,
				tc.TestData.Signers,
				tc.TestData.Signatures,
			)
			ak := s.BitsongApp.AccountKeeper
			sigModeHandler := s.EncodingConfig.TxConfig.SignModeHandler()

			// Only the first message is tested for authenticate
			addr := sdk.AccAddress(tc.TestData.Signers[0].PubKey().Address())

			if tc.TestData.ShouldSucceedGettingData {
				// request for the first message
				request, err := authenticator.GenerateAuthenticationRequest(s.Ctx, s.BitsongApp.AppCodec(), ak, sigModeHandler, addr, addr, nil, sdk.NewCoins(), tc.TestData.Msgs[0], tx, 0, false, authenticator.SequenceMatch, &types.AgAuthData{})
				s.Require().NoError(err)

				// Test Authenticate method
				if tc.TestData.ShouldSucceedSignatureVerification {
					initialized, err := s.SigVerificationAuthenticator.Initialize(tc.TestData.Signers[0].PubKey().Bytes())
					s.Require().NoError(err)
					err = initialized.Authenticate(s.Ctx, request)
					s.Require().NoError(err)
				} else {
					err = s.SigVerificationAuthenticator.Authenticate(s.Ctx, request)
					s.Require().Error(err)
				}
			} else {
				_, err := authenticator.GenerateAuthenticationRequest(s.Ctx, s.BitsongApp.AppCodec(), ak, sigModeHandler, addr, addr, nil, sdk.NewCoins(), tc.TestData.Msgs[0], tx, 0, false, authenticator.SequenceMatch, &types.AgAuthData{})
				s.Require().Error(err)
			}
		})
	}
}

// TODO: revisit multisignature
//func (s *SigVerifyAuthenticationSuite) TestMultiSignatureAuthenticator() {
//	bitsongToken := "bitsong"
//	priv := []cryptotypes.PrivKey{
//		s.TestPrivKeys[0],
//		s.TestPrivKeys[1],
//	}
//
//	feeCoins := sdk.Coins{sdk.NewInt64Coin(bitsongToken, 2500)}
//
//	sigs := make([]signing.SignatureV2, 1)
//	gen := s.EncodingConfig.TxConfig
//
//	// create a random length memo
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	memo := simulation.RandStringOfLength(r, simulation.RandIntBetween(r, 0, 100))
//	signMode := gen.SignModeHandler().Modes()
//
//	pkSet1 := generatePubKeysForMultiSig(priv...)
//	multisigKey1 := kmultisig.NewLegacyAminoPubKey(2, pkSet1)
//	multisignature1 := multisig.NewMultisig(len(pkSet1))
//
//	accAddress := sdk.AccAddress(multisigKey1.Address())
//	account := authtypes.NewBaseAccount(accAddress, multisigKey1, 0, 0)
//	s.BitsongApp.AccountKeeper.SetAccount(s.Ctx, account)
//
//	coins := sdk.Coins{sdk.NewInt64Coin(bitsongToken, 2500)}
//	msg := &banktypes.MsgSend{
//		FromAddress: accAddress.String(),
//		ToAddress:   sdk.MustBech32ifyAddressBytes(bitsongToken, s.TestAccAddress[1]),
//		Amount:      coins,
//	}
//
//	tx := gen.NewTxBuilder()
//	err := tx.SetMsgs(msg)
//	s.Require().NoError(err)
//
//	tx.SetMemo(memo)
//	sigs[0] = signing.SignatureV2{
//		PubKey:   multisigKey1,
//		Data:     &signing.MultiSignatureData{},
//		Sequence: 0,
//	}
//	err = tx.SetSignatures(sigs...)
//	s.Require().NoError(err)
//
//	tx.SetFeeAmount(feeCoins)
//	tx.SetGasLimit(300000)
//	s.Require().NoError(err)
//
//	signerData := authsigning.SignerData{
//		ChainID:       "",
//		AccountNumber: 0,
//		Sequence:      0,
//	}
//
//	// Get the signer bytes, use signModeLegacyAminoJSONHandler
//	signBytes, err := gen.SignModeHandler().GetSignBytes(signMode[1], signerData, tx.GetTx())
//
//	// Generate multisig signatures
//	sigSet1 := generateSignaturesForMultiSig(signBytes, priv...)
//
//	// Add signatures to signaturesv2 struct
//	for i := 0; i < len(pkSet1); i++ {
//		stdSig := legacytx.StdSignature{PubKey: pkSet1[i], Signature: sigSet1[i]}
//		sigV2, err := legacytx.StdSignatureToSignatureV2(s.EncodingConfig.Amino, stdSig)
//		s.Require().NoError(err)
//		err = multisig.AddSignatureV2(multisignature1, sigV2, pkSet1)
//		s.Require().NoError(err)
//	}
//
//	// 1st round: set SignatureV2 with empty signatures, to set correct
//	// signer infos.
//	sigs[0].Data = multisignature1
//
//	err = tx.SetSignatures(sigs...)
//	s.Require().NoError(err)
//
//	// Test GetAuthenticationData
//	authData, err := s.SigVerificationAuthenticator.GetAuthenticationData(s.Ctx, tx.GetTx(), -1, false)
//	s.Require().NoError(err)
//
//	// the signer data should contain 2 signers
//	sigData := authData.(authenticator.SignatureData)
//	s.Require().Equal(1, len(sigData.Signers))
//
//	// the signature data should contain 2 signatures
//	s.Require().Equal(1, len(sigData.Signatures))
//
//	// Test Authenticate method
//	authentication := s.SigVerificationAuthenticator.Authenticate(s.Ctx, nil, nil, authData)
//	s.Require().True(authentication.IsAuthenticated())
//}

func MakeTxBuilderBls381(
	ctx sdk.Context,
	gen client.TxConfig,
	msgs []sdk.Msg,
	feeAmt sdk.Coins,
	gas uint64,
	chainID string,
	accNums, accSeqs []uint64,
	signers, signatures []common.SecretKey,
) (client.TxBuilder, error) {
	// Validate inputs
	if len(msgs) == 0 {
		return nil, fmt.Errorf("no messages provided")
	}
	if len(signers) == 0 || len(signatures) == 0 {
		return nil, fmt.Errorf("no signers or signatures provided")
	}
	if len(signers) != len(signatures) {
		return nil, fmt.Errorf("mismatched lengths: signatures=%d, signers=%d", len(signatures), len(signers))
	}
	if len(accNums) != len(signers) || len(accSeqs) != len(signers) {
		return nil, fmt.Errorf("mismatched lengths: accNums=%d, accSeqs=%d, signers=%d", len(accNums), len(accSeqs), len(signers))
	}
	if feeAmt.IsZero() || !feeAmt.IsValid() {
		return nil, fmt.Errorf("invalid fee amount: %v", feeAmt)
	}
	if gas == 0 {
		return nil, fmt.Errorf("gas limit cannot be zero")
	}

	cosmosSigs := make([]signing.SignatureV2, len(signatures))
	pk := make([]common.PublicKey, 0)
	sigsInside := make([]common.Signature, 0)

	// Create a random length memo
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	memo := simulation.RandStringOfLength(r, simulation.RandIntBetween(r, 0, 100))
	signMode, err := authsigning.APISignModeToInternal(gen.SignModeHandler().DefaultMode())
	if err != nil {
		return nil, fmt.Errorf("failed to convert sign mode: %v", err)
	}

	// supportedModes := gen.SignModeHandler().SupportedModes()
	// fmt.Printf("supportedModes: %v\n", supportedModes)
	// if ok := supportedModes[signingv1beta1.SignMode(signMode)]; ok == nil {
	// 	return nil, fmt.Errorf("sign mode %v not supported by SignModeHandler", signMode)
	// }

	// 1st round: set SignatureV2 with derived public keys and empty signatures

	for i, privKey := range signers {
		pubKey, err := blst.GetCosmosBlsPubkey(privKey)
		if err != nil {
			return nil, fmt.Errorf("failed to derive BLS public key for signer %d: %v", i, err)
		}
		if pubKey == nil {
			return nil, fmt.Errorf("derived BLS public key is nil for signer %d", i)
		}

		cosmosSigs[i] = signing.SignatureV2{
			PubKey: pubKey,
			Data: &signing.SingleSignatureData{
				SignMode:  signMode,
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}
		pk = append(pk, privKey.PublicKey())
	}

	tx := gen.NewTxBuilder()
	extx, ok := tx.(client.ExtendedTxBuilder)
	if !ok {
		return nil, fmt.Errorf("failed to use ExtendTxBuilder interface: %v", err)
	}

	err = tx.SetMsgs(msgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to set messages: %v", err)
	}

	tx.SetMemo(memo)
	tx.SetFeeAmount(feeAmt)
	tx.SetGasLimit(gas)

	// Log transaction details
	txObj := tx.GetTx()
	if txObj == nil {
		return nil, fmt.Errorf("tx.GetTx() returned nil")
	}
	// fmt.Printf("tx.GetTx(): %+v\n", txObj)

	// Verify V2AdaptableTx implementation
	adaptableTx, ok := txObj.(authsigning.V2AdaptableTx)
	if !ok {
		return nil, fmt.Errorf("tx does not implement V2AdaptableTx, got %T", txObj)
	}
	txData := adaptableTx.GetSigningTxData()
	// fmt.Printf("txData: %+v\n", txData)

	// 2nd round: sign the transaction
	for i, p := range signatures {
		signerData := authsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}

		var pubKey *anypb.Any
		if cosmosSigs[i].PubKey != nil {
			cosmosPubkey, err := blst.GetCosmosBlsPubkey(p)
			anyPk, err := codectypes.NewAnyWithValue(cosmosPubkey)
			if err != nil {
				return nil, fmt.Errorf("failed to GetCosmosBlsPubkey %d: %v", i, err)
			}
			if err != nil {
				return nil, fmt.Errorf("failed to encode public key for signer %d: %v", i, err)
			}
			pubKey = &anypb.Any{
				TypeUrl: anyPk.TypeUrl,
				Value:   anyPk.Value,
			}
		}

		txSignerData := txsigning.SignerData{
			ChainID:       signerData.ChainID,
			AccountNumber: signerData.AccountNumber,
			Sequence:      signerData.Sequence,
			PubKey:        pubKey,
		}

		// gets the value to sign
		signBytes, err := gen.SignModeHandler().GetSignBytes(ctx, signingv1beta1.SignMode(signMode), txSignerData, txData)

		if err != nil {
			return nil, fmt.Errorf("failed to get sign bytes for signer %d: %v", i, err)
		}
		if signBytes == nil {
			return nil, fmt.Errorf("GetSignBytes returned nil for signer %d", i)
		}

		sig := p.Sign(signBytes)
		cosmosSigs[i].Data.(*signing.SingleSignatureData).Signature = sig.Marshal()
		fmt.Printf("cosmosSigs[i].Data: %v\n", cosmosSigs[i].Data)
		sigsInside = append(sigsInside, sig)
	}

	// Set the signature details into the TxExtension
	authExtSignature, err := gen.MarshalSignatureJSON(cosmosSigs)
	if err != nil {
		return nil, fmt.Errorf("failed to set messages: %v", err)
	}

	txExtensionData := &smartaccounttypes.TxExtension{
		SelectedAuthenticators: []uint64{1},
		SmartAccount: &smartaccounttypes.AgAuthData{
			Data: authExtSignature,
		},
	}
	extOpts, err := codectypes.NewAnyWithValue(txExtensionData)
	if err != nil {
		return nil, fmt.Errorf("failed to set messages: %v", err)
	}

	// aggregate signatures & pubkey, set into signature options
	fmt.Printf("pk: %v\n", pk[0])
	aggPubkeys := blst.AggregateMultiplePubkeys(pk)
	fmt.Printf("aggPubkeys: %v\n", aggPubkeys)

	aggPubkey, err := blst.GetCosmosBlsPubkeyFromPubkey(aggPubkeys)
	if err != nil {
		return nil, fmt.Errorf("failed to blst.GetCosmosBlsPubkeyFromPubkey: %v", err)
	}
	fmt.Printf("aggPubkey: %v\n", aggPubkey)
	aggSig := blst.AggregateSignatures(sigsInside)
	extx.SetExtensionOptions(extOpts)
	err = tx.SetSignatures(signing.SignatureV2{
		PubKey: aggPubkey,
		Data: &signing.SingleSignatureData{
			SignMode:  signMode,
			Signature: aggSig.Marshal(),
		},
		Sequence: accSeqs[0],
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set final signatures: %v", err)
	}

	sigTx, ok := tx.(authsigning.Tx)
	pubkeys, _ := sigTx.GetPubKeys()
	fmt.Printf("pubkeys: %v\n", pubkeys)

	tx.GetTx()

	return tx, nil
}

func MakeTxBuilder(ctx sdk.Context,
	gen client.TxConfig,
	msgs []sdk.Msg,
	feeAmt sdk.Coins,
	gas uint64,
	chainID string,
	accNums,
	accSeqs []uint64,
	signers []cryptotypes.PrivKey,
	signatures []cryptotypes.PrivKey,
) (client.TxBuilder, error) {
	sigs := make([]signing.SignatureV2, len(signatures))

	// create a random length memo
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	memo := simulation.RandStringOfLength(r, simulation.RandIntBetween(r, 0, 100))
	signMode, err := authsigning.APISignModeToInternal(gen.SignModeHandler().DefaultMode())
	if err != nil {
		return nil, err
	}

	// 1st round: set SignatureV2 with empty signatures, to set correct
	// signer infos.
	for i, p := range signers {
		sigs[i] = signing.SignatureV2{
			PubKey: p.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode: signMode,
			},
			Sequence: accSeqs[i],
		}
	}

	tx := gen.NewTxBuilder()
	err = tx.SetMsgs(msgs...)
	if err != nil {
		return nil, err
	}
	err = tx.SetSignatures(sigs...)
	if err != nil {
		return nil, err
	}
	tx.SetMemo(memo)
	tx.SetFeeAmount(feeAmt)
	tx.SetGasLimit(gas)

	// 2nd round: once all signer infos are set, every signer can sign.
	signers = signers[0:len(signatures)]
	for i, p := range signatures {
		signerData := authsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		signBytes, err := authsigning.GetSignBytesAdapter(
			ctx, gen.SignModeHandler(), signMode, signerData, tx.GetTx())
		if err != nil {
			panic(err)
		}
		sig, err := p.Sign(signBytes)
		if err != nil {
			panic(err)
		}
		sigs[i].Data.(*signing.SingleSignatureData).Signature = sig
	}

	err = tx.SetSignatures(sigs...)
	if err != nil {
		panic(err)
	}
	return tx, nil
}

// GenTx generates a signed mock transaction.
func GenTx(
	ctx sdk.Context,
	gen client.TxConfig,
	msgs []sdk.Msg,
	feeAmt sdk.Coins,
	gas uint64,
	chainID string,
	accNums,
	accSeqs []uint64,
	signers []cryptotypes.PrivKey,
	signatures []cryptotypes.PrivKey,
) (sdk.Tx, error) {
	tx, err := MakeTxBuilder(ctx, gen, msgs, feeAmt, gas, chainID, accNums, accSeqs, signers, signatures)
	if err != nil {
		return nil, err
	}
	return tx.GetTx(), nil
}

// GenTx generates a signed mock transaction.
func GenTxBls12381(
	ctx sdk.Context,
	gen client.TxConfig,
	msgs []sdk.Msg,
	feeAmt sdk.Coins,
	gas uint64,
	chainID string,
	accNums, accSeqs []uint64,
	signers, signatures []common.SecretKey,
) (sdk.Tx, error) {
	tx, err := MakeTxBuilderBls381(ctx, gen, msgs, feeAmt, gas, chainID, accNums, accSeqs, signers, signatures)
	if err != nil {
		return nil, err
	}
	return tx.GetTx(), nil
}

func generatePubKeysForMultiSig(
	priv ...cryptotypes.PrivKey,
) (pubkeys []cryptotypes.PubKey) {
	pubkeys = make([]cryptotypes.PubKey, len(priv))
	for i, p := range priv {
		pubkeys[i] = p.PubKey()
	}
	return
}

func generateSignaturesForMultiSig(
	msg []byte,
	priv ...cryptotypes.PrivKey,
) (signatures [][]byte) {
	signatures = make([][]byte, len(priv))
	for i, p := range priv {
		var err error
		signatures[i], err = p.Sign(msg)
		if err != nil {
			panic(err)
		}
	}
	return
}
