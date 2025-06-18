package blst_test

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"testing"

	"github.com/bitsongofficial/go-bitsong/crypto/bls/blst"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/common"
	"github.com/stretchr/testify/require"
)

func TestMarshalUnmarshal(t *testing.T) {
	priv, err := blst.RandKey()
	require.NoError(t, err)
	b := priv.Marshal()
	// b32 := bytesutil.ToBytes32(b)
	pk, err := blst.SecretKeyFromBytes(b[:])
	require.NoError(t, err)
	pk2, err := blst.SecretKeyFromBytes(b[:])
	require.NoError(t, err)
	require.Equal(t, pk.Marshal(), pk2.Marshal(), "Keys not equal")
}

func TestSecretKeyFromBytes(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   error
	}{
		{
			name: "Nil",
			err:  errors.New("secret key must be 32 bytes"),
		},
		{
			name:  "Empty",
			input: []byte{},
			err:   errors.New("secret key must be 32 bytes"),
		},
		{
			name:  "Short",
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			err:   errors.New("secret key must be 32 bytes"),
		},
		{
			name:  "Long",
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			err:   errors.New("secret key must be 32 bytes"),
		},
		{
			name:  "Bad",
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			err:   common.ErrSecretUnmarshal,
		},
		{
			name:  "Good",
			input: []byte{0x25, 0x29, 0x5f, 0x0d, 0x1d, 0x59, 0x2a, 0x90, 0xb3, 0x33, 0xe2, 0x6e, 0x85, 0x14, 0x97, 0x08, 0x20, 0x8e, 0x9f, 0x8e, 0x8b, 0xc1, 0x8f, 0x6c, 0x77, 0xbd, 0x62, 0xf8, 0xad, 0x7a, 0x68, 0x66},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := blst.SecretKeyFromBytes(test.input)
			if test.err != nil {
				require.NotEqual(t, nil, err, "No error returned")
				require.ErrorContains(t, test.err, err.Error(), "Unexpected error returned")
			} else {
				require.NoError(t, err)
				require.Equal(t, 0, bytes.Compare(res.Marshal(), test.input))
			}
		})
	}
}

func TestSerialize(t *testing.T) {
	rk, err := blst.RandKey()
	require.NoError(t, err)
	b := rk.Marshal()

	_, err = blst.SecretKeyFromBytes(b)
	require.NoError(t, err)
}

func TestZeroKey(t *testing.T) {
	// Is Zero
	var zKey [32]byte
	require.Equal(t, true, blst.IsZero(zKey[:]))

	// Is Not Zero
	_, err := rand.Read(zKey[:])
	require.NoError(t, err)
	require.Equal(t, false, blst.IsZero(zKey[:]))
}

func TestRandKeyUniqueness(t *testing.T) {
	numKeys := 10
	keys := make([]common.SecretKey, numKeys)

	for i := 0; i < numKeys; i++ {
		key, err := blst.RandKey()
		if err != nil {
			t.Errorf("RandKey() returned an error: %v", err)
		}
		keys[i] = key
		fmt.Printf("key.Marshal(): %v\n", key.Marshal())
	}

	// Compare each key with every other key
	for i := 0; i < numKeys; i++ {
		for j := i + 1; j < numKeys; j++ {
			if bytes.Equal(keys[i].Marshal(), keys[j].Marshal()) {
				t.Errorf("RandKey() generated duplicate keys at indices %d and %d", i, j)
			}
		}
	}
}
