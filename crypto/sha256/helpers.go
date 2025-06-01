package helpers

import (
	"crypto/sha256"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Generates the SHA256SUM for an array of cosmos-sdk messages
func Sha256Msgs(msgs []sdk.Msg) [32]byte {
	var msgStrings []string
	for _, msg := range msgs {
		msgStr := msg.String() // ensure this is linted
		msgStrings = append(msgStrings, msgStr)
	}
	jsonBytes, _ := json.Marshal(msgStrings)
	return sha256.Sum256(jsonBytes)
}
