package panel

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// LoginPreRequest fetch the LoginRequesttoken
type LoginPreRequest struct {
	Username string `json:"username"`
}

type LoginRequestToken struct {
	Username      string `json:"username"`
	Timestamp     int64  `json:"timestamp"`
	RemoteAddress string `json:"remote_address"`
	Nonce         string `json:"nonce"`
}

type LoginRequest struct {
	LoginRequestTokenString string `json:"login_request_token"`
	LoginHmacHex            string `json:"login_hmac_hex"`
}

type UserToken struct {
	Username   string `json:"username"`
	ExpireTime int64  `json:"expire_time"`
}

type LoginResponse struct {
	Sueecss         bool   `json:"success"`
	Message         string `json:"message"`
	UserTokenString string `json:"user_token_string"`
}

func NewLoginRequestToken(username string, remoteAddress string) *LoginRequestToken {
	nonce := randString(32)
	return &LoginRequestToken{
		Username:      username,
		Timestamp:     time.Now().Unix(),
		RemoteAddress: remoteAddress,
		Nonce:         nonce,
	}
}

func (l *LoginRequestToken) ToJwtString() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = l.Username
	claims["timestamp"] = l.Timestamp
	claims["remote_address"] = l.RemoteAddress
	claims["nonce"] = l.Nonce
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (l *LoginRequestToken) ParseJwtString(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token")
	}
	l.Username = claims["username"].(string)
	l.Timestamp = int64(claims["timestamp"].(float64))
	l.RemoteAddress = claims["remote_address"].(string)
	l.Nonce = claims["nonce"].(string)
	return nil
}

func (l *LoginRequest) VerifyPassword(passwordHmacHex string) (bool, error) {
	mac := HmacSha256Hex(passwordHmacHex, l.LoginRequestTokenString)
	if mac != l.LoginHmacHex {
		return false, errors.New("hmac not match")
	}
	return true, nil
}

func (u *UserToken) ToJwtString() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["expire_time"] = u.ExpireTime
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *UserToken) ParseJwtString(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token")
	}
	u.Username = claims["username"].(string)
	u.ExpireTime = int64(claims["expire_time"].(float64))
	return nil
}
