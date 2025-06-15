package blst

import (
	"fmt"

	"github.com/bitsongofficial/go-bitsong/btsgutils/cache/nonblocking"
	"github.com/bitsongofficial/go-bitsong/crypto/bls/common"
	blst "github.com/supranational/blst/bindings/go"
)

func init() {
	// Limit blst operations to a single core
	blst.SetMaxProcs(1)
	onEvict := func(_ [48]byte, _ common.PublicKey) {}
	keysCache, err := nonblocking.NewLRU(maxKeys, onEvict)
	if err != nil {
		panic(fmt.Sprintf("Could not initiate public keys cache: %v", err))
	}
	pubkeyCache = keysCache
}
