package panel

import (
	"crypto/rand"
)

// secretKey changes every time the program starts
var secretKey []byte

func init() {
	secretKey := make([]byte, 8)
	_, _ = rand.Read(secretKey)
}
