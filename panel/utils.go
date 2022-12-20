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

func HmacSha256String(key, data string) string {
	macBytes := (HmacSha256([]byte(key), []byte(data)))
	macHex := hex.EncodeToString(macBytes)
	return macHex
}
