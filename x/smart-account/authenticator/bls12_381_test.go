package authenticator_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"

	"github.com/bitsongofficial/go-bitsong/crypto/bls"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/blst"
	btsgblst "github.com/bitsongofficial/go-bitsong/crypto/bls/blst"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/common"
	"github.com/bitsongofficial/go-bitsong/x/smart-account/authenticator"
	"github.com/bitsongofficial/go-bitsong/x/smart-account/types"
)

type Bls12381AuthenticatorTest struct {
	BaseAuthenticatorSuite
	Bls12381Auth authenticator.Bls12381
}

func TestBls12381Auth(t *testing.T) {
	suite.Run(t, new(Bls12381AuthenticatorTest))
}

func (s *Bls12381AuthenticatorTest) SetupTest() {
	s.SetupKeys()
	am := authenticator.NewAuthenticatorManager()

	// Define authenticators
	s.Bls12381Auth = authenticator.NewBls12381(am, s.BitsongApp.GetKVStoreKey()[types.ModuleName])
	am.RegisterAuthenticator(s.Bls12381Auth)
}

func (s *Bls12381AuthenticatorTest) TearDownTest() {
	os.RemoveAll(s.HomeDir)
}

func (s *Bls12381AuthenticatorTest) TestBls12381() {
	// Define test cases
	type testCase struct {
		name                  string
		includeTxExt          bool
		includeAggPkSig       bool
		expectSuccessfulAdded bool   // expect registering an auth to be valid
		expectSuccessful      bool   // expect authentication to be valid
		expectConfirm         bool   // expect confirm execution to be valid
		numKeys               uint64 // # off agg keys in set
		threshold             uint64 // signing threshold
		signers               uint64 // # of keys to include in msg
	}

	testCases := []testCase{

		{
			name:                  "with txExt",
			numKeys:               1,
			threshold:             1,
			signers:               1,
			includeTxExt:          true,
			includeAggPkSig:       true,
			expectSuccessfulAdded: true,
			expectSuccessful:      true,
			expectConfirm:         true,
		},
		{
			name:                  "without txExt",
			numKeys:               1,
			threshold:             1,
			signers:               1,
			includeTxExt:          false,
			includeAggPkSig:       true,
			expectSuccessfulAdded: true,
			expectSuccessful:      false,
			expectConfirm:         true,
		},
		{
			name:                  "multiple keys, threshold met",
			numKeys:               2,
			threshold:             2,
			signers:               2,
			includeTxExt:          true,
			includeAggPkSig:       true,
			expectSuccessfulAdded: true,
			expectSuccessful:      true,
			expectConfirm:         true,
		},
		{
			name:                  "multiple keys, threshold not met",
			numKeys:               1,
			threshold:             2,
			signers:               1,
			includeTxExt:          true,
			includeAggPkSig:       true,
			expectSuccessfulAdded: false,
			expectSuccessful:      false,
			expectConfirm:         true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			cdc := s.BitsongApp.AppCodec()

			secretKeys, blsConfig, err := GenerateBLSPrivateKeysReturnBlsConfig(int(tc.numKeys), tc.threshold, 369)
			s.Require().NoError(err)
			fmt.Printf("len(secretKeys): %v\n", len(secretKeys))

			txSender := s.TestPrivKeys[0].PubKey().Address()
			// fmt.Printf("txSender: %v\n", txSender)
			// fmt.Printf("secretKeys[0].PublicKey().Marshal(): %v\n", secretKeys[0].PublicKey().Marshal())
			// fmt.Printf("len(secretKeys[0].PublicKey().Marshal()): %v\n", len(secretKeys[0].PublicKey().Marshal()))
			initializedAuth, err := s.Bls12381Auth.Initialize([]byte{})
			s.Require().NoError(err)

			s.Require().NoError(err)

			// Generate authentication request
			ak := s.BitsongApp.AccountKeeper

			// sample msg
			msg := &bank.MsgSend{FromAddress: s.TestAccAddress[0].String(), ToAddress: "to", Amount: sdk.NewCoins(sdk.NewInt64Coin("foo", 1))}

			// digest msg into hash being signed
			msgsToHash := []sdk.Msg{msg}
			var anyMsgs []authenticator.LocalAny
			for _, in := range msgsToHash {
				anyMsg, _ := codectypes.NewAnyWithValue(in)
				anyMsgs = append(anyMsgs, authenticator.LocalAny{
					TypeURL: anyMsg.TypeUrl,
					Value:   anyMsg.Value,
				})
			}

			msgDigestHash := authenticator.Sha256Msgs(anyMsgs)
			// fmt.Printf("msgDigestHash: %v\n", msgDigestHash)
			signerKeys := secretKeys
			if tc.signers < uint64(len(secretKeys)) {
				signerKeys = secretKeys[:tc.signers]
			}
			// Sign the message with the keys
			agAuthData, err := SignMessageWithTestBls12Keys(s.EncodingConfig.TxConfig, msgDigestHash[:], signerKeys)

			s.Require().NoError(err)

			// omit inclusion
			var smartAccAuth *types.AgAuthData
			if tc.includeTxExt {
				smartAccAuth = agAuthData
			}
			// // fmt.Printf("smartAccAuth: %v\n", smartAccAuth)
			// //  todo: aggregate pubkey
			// if tc.includeAggPkSig {}

			// sample tx
			tx, err := s.GenSimpleTxBls12381(msgsToHash, secretKeys, txSender)
			s.Require().NoError(err)

			// pass msgs based on test instance
			request, err := authenticator.GenerateAuthenticationRequest(
				s.Ctx, cdc, ak,
				s.EncodingConfig.TxConfig.SignModeHandler(),
				s.TestAccAddress[0], s.TestAccAddress[0],
				nil, sdk.NewCoins(),
				msg, tx,
				0, false,
				authenticator.SequenceMatch,
				smartAccAuth,
			)
			s.Require().NoError(err)

			// sign, err := request.SignatureData.Signers[0].Marshal()
			// fmt.Printf("request.SignatureData: %v\n", sign)
			request.AuthenticatorId = "1"

			// fmt.Printf("request.Account.String(): %v\n", request.Account.String())
			// fmt.Printf("len(request.Account): %v\n", len(request.Account))
			// Attempt to authenticate using initialized authenticator
			bzBlsConfig, err := blsConfig.Marshal()
			s.Require().NoError(err)
			// fmt.Printf("blsConfig: %v\n", blsConfig)
			err = initializedAuth.OnAuthenticatorAdded(s.Ctx, request.Account, bzBlsConfig, request.AuthenticatorId)
			s.Require().Equal(tc.expectSuccessfulAdded, err == nil)
			err = initializedAuth.Authenticate(s.Ctx, request)
			fmt.Printf("err: %v\n", err)
			s.Require().Equal(tc.expectSuccessful, err == nil)
			err = initializedAuth.Track(s.Ctx, request)
			s.Require().NoError(err)
			err = initializedAuth.ConfirmExecution(s.Ctx, request)
			s.Require().Equal(tc.expectConfirm, err == nil)
		})
	}
}

// GenerateBLSPrivateKeys creates n BLS12-381 private keys.
func GenerateBLSPrivateKeysReturnBlsConfig(n int, threshold uint64, seed int64) ([]common.SecretKey, types.BlsConfig, error) {
	if n < 0 {
		return nil, types.BlsConfig{}, fmt.Errorf("number of keys must be non-negative, got %d", n)
	}
	if n == 0 {
		return []common.SecretKey{}, types.BlsConfig{}, nil
	}

	secretKeys := make([]common.SecretKey, n)
	pubkeys := make([][]byte, n)

	for i := 0; i < n; i++ {
		key, err := blst.RandKey()
		if err != nil {
			return nil, types.BlsConfig{}, fmt.Errorf("RandKey() returned an error: %v", err)
		}
		// Convert to BLS secret key
		secretKey, err := bls.SecretKeyFromBytes(key.Marshal())
		if err != nil {
			return nil, types.BlsConfig{}, fmt.Errorf("failed to create secret key %d: %v", i, err)
		}

		secretKeys[i] = secretKey
		pubkeys[i] = secretKey.PublicKey().Marshal()

	}
	// fmt.Printf("pubkeys: %v\n", pubkeys)
	return secretKeys, types.BlsConfig{
		Pubkeys:   pubkeys,
		Threshold: threshold,
	}, nil
}

// SignMessageWithKeys signs a 32-byte message hash with a list of BLS private keys
// and returns a SmartAccountAuthData object with public keys and signatures.
func SignMessageWithTestBls12Keys(gen client.TxConfig, msgHash []byte, secretKeys []common.SecretKey) (*types.AgAuthData, error) {
	// Validate message hash
	if len(msgHash) != 32 {
		return nil, fmt.Errorf("message hash must be 32 bytes, got %d", len(msgHash))
	}

	// Validate secret keys
	if len(secretKeys) == 0 {
		return nil, fmt.Errorf("no secret keys provided")
	}

	// Initialize SmartAccountAuthData
	auth := &types.AgAuthData{
		Data: []byte{},
	}

	// Sign with each key
	sigs := make([]signing.SignatureV2, 0, len(secretKeys))
	for i, sk := range secretKeys {
		sig := sk.Sign(msgHash)
		if sig == nil {
			return nil, fmt.Errorf("failed to sk.Sign %d", i)
		}
		pubkey, err := btsgblst.GetCosmosBlsPubkeyFromPubkey(sk.PublicKey())
		if err != nil { // Fix: check err != nil
			return nil, fmt.Errorf("failed to NewPublicKeyFromBytes %d: %w", i, err)
		}
		sigV2 := signing.SignatureV2{
			PubKey: pubkey,
			Data: &signing.SingleSignatureData{
				SignMode:  0,
				Signature: sig.Marshal(),
			},
			Sequence: 0,
		}
		// define signature with correct Interface
		sigs = append(sigs, sigV2)
	}

	// Marshal the signatures array into bytes
	fmt.Printf("sigs: %v\n", sigs)
	signBz, err := authenticator.MarshalSignatureJSON(sigs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signatures: %w", err)
	}
	auth.Data = signBz
	return auth, nil
}
