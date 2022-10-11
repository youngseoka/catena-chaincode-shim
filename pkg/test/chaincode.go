package test

//
//import (
//	"github.com/btcsuite/btcd/btcec"
//	"github.com/whiteblockco/catena-camp/chaincode/auth/did/types"
//	"github.com/whiteblockco/catena-camp/internal/chaincode"
//	ccTypes "github.com/whiteblockco/catena-camp/internal/chaincode/types"
//	"github.com/whiteblockco/catena-camp/pkg/state"
//	"github.com/whiteblockco/pkg/jwt"
//	"github.com/whiteblockco/pkg/jwt/signing/ecdsa"
//)
//
//func init() {
//
//}
//
//func NewState() (*state.KvState, error) {
//	dbName := "campTestDB"
//	dbType := "goleveldb"
//	dbDir := "/tmp/testCamp/db"
//
//	st := state.NewState(dbName, dbType, dbDir)
//
//	return st, nil
//}
//
//func NewPrivKey() (*btcec.PrivateKey, error) {
//	privKey, err := btcec.NewPrivateKey(btcec.S256())
//	if err != nil {
//		return nil, err
//	}
//	return privKey, nil
//}
//
//func NewPrivKeyFromBytes(privKeyBytes []byte) *btcec.PrivateKey {
//	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
//	return privKey
//}
//
//func GetCreateDIDMsg(privKey *btcec.PrivateKey) *ccTypes.JwtReq {
//	// Create publicKey object
//	encodedPubKey := chaincode.EncodedPubKey(privKey)
//	publicKey := types.PublicKey{
//		ID:              types.DefaultDIDKey,
//		Type:            types.PubKeyType,
//		PublicKeyBase64: encodedPubKey,
//	}
//
//	// Nonce
//	nonce, err := chaincode.GetNonce()
//	if err != nil {
//		panic(err)
//	}
//
//	// Create object
//	createPayload := types.MsgCreatePayload{
//		Nonce:     nonce,
//		PublicKey: publicKey,
//	}
//
//	// Create signed token
//	kid := chaincode.GetDefaultKidFromPubKey(privKey.PubKey())
//	token, err := jwt.NewToken(createPayload, kid)
//	if err != nil {
//		panic(err)
//	}
//
//	tokenString, err := token.Sign(ecdsa.ES256k, privKey)
//	if err != nil {
//		panic(err)
//	}
//
//	// Create msg object
//	signedMsg := &ccTypes.JwtReq{
//		Token: tokenString,
//	}
//
//	return signedMsg
//}
//
//func GetExtraWithCID(cid string) []ccTypes.KeyValue {
//	var extra []ccTypes.KeyValue
//	kv := ccTypes.KeyValue{
//		Key:   []byte("CID"),
//		Value: []byte(cid),
//	}
//	extra = append(extra, kv)
//
//	return extra
//}
//
//func MustGetNonce() []byte {
//	// Nonce
//	nonce, err := chaincode.GetNonce()
//	if err != nil {
//		panic(err)
//	}
//	return nonce
//}
