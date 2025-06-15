package authenticator_test

import (
	"crypto/sha256"
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
		name             string
		includeTxExt     bool
		includeAggPkSig  bool
		expectInit       bool
		expectSuccessful bool
		expectConfirm    bool
		numKeys          uint64
		threshold        uint64
	}

	testCases := []testCase{

		{
			name:             "with txExt, w/out agg",
			numKeys:          1,
			threshold:        1,
			includeTxExt:     true,
			includeAggPkSig:  true,
			expectInit:       true,
			expectSuccessful: true,
			expectConfirm:    true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {

			secretKeys, blsConfig, err := GenerateBLSPrivateKeysReturnBlsConfig(int(tc.numKeys), tc.threshold, 369)
			s.Require().NoError(err)

			bzBlsConfig, err := blsConfig.Marshal()
			s.Require().NoError(err)
			fmt.Printf("bzBlsConfig: %v\n", bzBlsConfig)

			initializedAuth, err := s.Bls12381Auth.Initialize(bzBlsConfig)
			s.Require().NoError(err)

			if !tc.expectInit {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var smartAccAuth *types.AgAuthData

				// Generate authentication request
				ak := s.BitsongApp.AccountKeeper
				sigModeHandler := s.EncodingConfig.TxConfig.SignModeHandler()

				// sample msg
				msg := &bank.MsgSend{FromAddress: s.TestAccAddress[0].String(), ToAddress: "to", Amount: sdk.NewCoins(sdk.NewInt64Coin("foo", 1))}

				// digest msg into hash being signed
				msgsToHash := []sdk.Msg{msg}
				var anyMsgs []string
				for _, in := range msgsToHash {
					anyMsg, _ := codectypes.NewAnyWithValue(in)
					anyMsgs = append(anyMsgs, anyMsg.String())
				}
				valueToHash, _ := types.Amino.Marshal(anyMsgs)
				hash := sha256.Sum256(valueToHash)

				// Sign the message with the keys
				smartAccTxExtension, err := SignMessageWithTestBls12Keys(s.EncodingConfig.TxConfig, hash[:], secretKeys)
				s.Require().NoError(err)
				// omit inclusion
				if tc.includeTxExt {
					smartAccAuth = smartAccTxExtension
				}

				//  todo: aggregate pubkey
				if tc.includeAggPkSig {
				}

				// sample tx
				tx, err := s.GenSimpleTxBls12381(msgsToHash, secretKeys, s.TestPrivKeys[0].PubKey().Address())
				s.Require().NoError(err)

				cdc := s.BitsongApp.AppCodec()
				fmt.Printf("smartAccAuth: %v\n", smartAccAuth)
				fmt.Printf("hash: %v\n", hash)
				//  pass msgs based on test instance
				request, err := authenticator.GenerateAuthenticationRequest(s.Ctx, cdc, ak, sigModeHandler, s.TestAccAddress[0], s.TestAccAddress[0], nil, sdk.NewCoins(), msg, tx, 0, false, authenticator.SequenceMatch, smartAccAuth)
				s.Require().NoError(err)

				// Attempt to authenticate using initialized authenticator
				err = initializedAuth.Authenticate(s.Ctx, request)
				s.Require().Equal(tc.expectSuccessful, err == nil)

				err = initializedAuth.ConfirmExecution(s.Ctx, request)
				s.Require().Equal(tc.expectConfirm, err == nil)

			}
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
		// Convert to BLS secret key
		privKeyBytes := []byte{0x25, 0x29, 0x5f, 0x0d, 0x1d, 0x59, 0x2a, 0x90, 0xb3, 0x33, 0xe2, 0x6e, 0x85, 0x14, 0x97, 0x08, 0x20, 0x8e, 0x9f, 0x8e, 0x8b, 0xc1, 0x8f, 0x6c, 0x77, 0xbd, 0x62, 0xf8, 0xad, 0x7a, 0x68, 0x66}
		secretKey, err := bls.SecretKeyFromBytes(privKeyBytes)
		if err != nil {
			return nil, types.BlsConfig{}, fmt.Errorf("failed to create secret key %d: %v", i, err)
		}

		secretKeys[i] = secretKey

	}
	return secretKeys, types.BlsConfig{
		Pubkeys:   pubkeys,
		Threshold: threshold,
	}, nil
}

// SignMessageWithKeys signs a 32-byte message hash with a list of BLS private keys
// and returns a AgAuthData object with public keys and signatures.
func SignMessageWithTestBls12Keys(gen client.TxConfig, msgHash []byte, secretKeys []common.SecretKey) (*types.AgAuthData, error) {
	// Validate message hash
	if len(msgHash) != 32 {
		return nil, fmt.Errorf("message hash must be 32 bytes, got %d", len(msgHash))
	}

	// Validate secret keys
	if len(secretKeys) == 0 {
		return nil, fmt.Errorf("no secret keys provided")
	}

	// Initialize AgAuthData
	auth := &types.AgAuthData{
		Signatures: []byte{},
	}

	// Sign with each key
	sigs := make([]signing.SignatureV2, 0, len(secretKeys))
	for i, sk := range secretKeys {
		sig := sk.Sign(msgHash)
		if sig == nil {
			return nil, fmt.Errorf("failed to sk.Sign %d", i)
		}

		fmt.Printf("sk.PublicKey().Marshal(): %v\n", sk.PublicKey().Marshal())
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
	fmt.Printf("signBz: %v\n", signBz)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signatures: %w", err)
	}
	auth.Signatures = signBz

	return auth, nil
}
