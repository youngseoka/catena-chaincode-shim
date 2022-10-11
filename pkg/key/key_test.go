package key

import (
	"crypto/ed25519"
	"github.com/btcsuite/btcd/btcec"
	"github.com/stretchr/testify/assert"
	"github.com/youngseoka/catena-chaincode-shim/pkg/test"
	testData "github.com/youngseoka/catena-chaincode-shim/pkg/test/data"
	"testing"
)

func TestBtcecKeyPair(t *testing.T) {
	t.Run("PrivKey", func(t *testing.T) {
		privKey, err := test.NewPrivKey()
		if err != nil {
			t.Fatal(err)
		}
		assert.IsType(t, &btcec.PrivateKey{}, privKey)

		privKeyBytes := privKey.Serialize()
		assert.Len(t, privKeyBytes, btcec.PrivKeyBytesLen)
	})

	t.Run("PubKey", func(t *testing.T) {
		privKey, err := test.NewPrivKey()
		if err != nil {
			t.Fatal(err)
		}
		pubKey := privKey.PubKey()
		assert.IsType(t, &btcec.PublicKey{}, pubKey)

		pubKeyBytes := pubKey.SerializeCompressed()
		assert.Len(t, pubKeyBytes, btcec.PubKeyBytesLenCompressed)
	})
}

func TestGetAddressFromPubKey(t *testing.T) {
	// Test btcec key get address
	t.Run("BtcecKey", func(t *testing.T) {
		privKey := test.NewPrivKeyFromBytes(testData.TestingBtcecPrivKey1)
		pubKey := privKey.PubKey()
		addr := GetAddressFromPubKey(pubKey)
		assert.Equal(t, testData.ExpectedBtcecPrivKey1Addr, addr)
	})
}

func TestGetDIDFromPubKey(t *testing.T) {
	// Test btcec key get did(id)
	t.Run("BtcecKey", func(t *testing.T) {
		privKey := test.NewPrivKeyFromBytes(testData.TestingBtcecPrivKey1)
		pubKey := privKey.PubKey()
		did := GetDIDFromPubKey(pubKey)
		assert.Equal(t, testData.ExpectedBtcecPrivKey1DID, did)
	})
}

func Testdd(t *testing.T) {
	// Test btcec key get did(id)
	t.Run("BtcecKey", func(t *testing.T) {
		_, privKey, _ := ed25519.GenerateKey()

	})
}
