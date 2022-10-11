package nonce

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/tendermint/tendermint/crypto"
)

const (
	nonceLen = 12
)

func GetNonceBytes() []byte {
	return crypto.CRandBytes(nonceLen)
}

func GetNonceString() string {
	nonceBytes := GetNonceBytes()
	return base58.Encode(nonceBytes)
}
