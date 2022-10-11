package random

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

const (
	keyLen = 30
)

func GetRandomKey(prefix string, nonce string, txTime int64) (string, error) {
	newTime := time.Unix(0, txTime)
	newTimeBytes, err := newTime.MarshalBinary()
	if err != nil {
		return "", err
	}

	nonceBytes := base58.Decode(nonce)

	spsBytes := append(newTimeBytes[:], nonceBytes[:]...)
	//sps := base58.Encode(spsBytes[:keyLen])
	sps := base58.Encode(spsBytes)

	key := fmt.Sprintf("%s:%s", prefix, sps)

	return key, nil
}
