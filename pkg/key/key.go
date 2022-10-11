package key

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tendermint/tendermint/libs/bytes"
)

const (
	TruncatedSize = 20
	DIDScheme     = "did"
	DIDMethod     = "ctn"
	DIDFormat     = "%s:%s:%s:%s"
	DIDPKID       = "%s#keys-%d"
	DIDServiceID  = "%s#endpoint-%d"
)
type Address = bytes.HexBytes

func GetDIDFromPubKey(pubKey ed25519.PublicKey) string {
	address := GetAddressFromPubKey(pubKey)
	id := base58.Encode(address)
	did := fmt.Sprintf(DIDFormat, DIDScheme, DIDMethod, "01", id)
	return did
}

func GetPubKeyBytes(privKey ed25519.PrivateKey) ed25519.PublicKey {
	pubKey := make([]byte, ed25519.PublicKeySize)
	copy(pubKey, privKey[32:])

	return pubKey
}

func GetAddressFromPubKey(pubKey ed25519.PublicKey) Address {
	hash := sha256.Sum256(pubKey[:])

	return hash[:TruncatedSize]
}
