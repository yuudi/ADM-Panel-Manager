package panel

import (
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	l := LoginRequestToken{
		Username:      "admin",
		Timestamp:     time.Now().Unix(),
		RemoteAddress: "1.2.3.4",
		Nonce:         "test",
	}
	tokenString, err := l.ToJwtString()
	if err != nil {
		t.Error(err)
	}

	l2 := LoginRequestToken{}

	err = l2.ParseJwtString(tokenString)
	if err != nil {
		t.Error(err)
	}

	secretKeyBak := secretKey
	defer func() {
		secretKey = secretKeyBak
	}()
	secretKey = []byte("different")
	err = l2.ParseJwtString(tokenString)
	if err == nil {
		t.Error("should fail")
	}
}
