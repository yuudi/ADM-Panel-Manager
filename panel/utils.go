package panel

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256(key, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func HmacSha256Hex(keyHex, data string) string {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		panic(err)
	}
	macBytes := (HmacSha256([]byte(key), []byte(data)))
	macHex := hex.EncodeToString(macBytes)
	return macHex
}
