package blst

import (
	blst "github.com/supranational/blst/bindings/go"
)

// bls12SecretKey used in the BLS signature scheme.
type bls12SecretKey struct {
	p *blst.SecretKey
}
