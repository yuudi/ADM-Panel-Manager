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
	LoginRequestTokenString string `json:"login_token_string"`
	LoginHmac               string `json:"login_hmac"`
}

type UserToken struct {
	Username string `json:"username"`
}

type LoginResponse struct {
	Sueecss         bool   `json:"success"`
	Message         string `json:"message"`
	UserTokenString string `json:"user_token_string"`
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

func (l *LoginRequest) Verify() (bool, error) {
	loginRequestToken := LoginRequestToken{}
	err := loginRequestToken.ParseJwtString(l.LoginRequestTokenString)
	if err != nil {
		return false, err
	}
	// if loginRequestToken.RemoteAddress != remoteAddress {
	// 	return false, errors.New("remote address not match")
	// }
	if loginRequestToken.Timestamp+60 < time.Now().Unix() {
		return false, errors.New("token expired")
	}
	mac := HmacSha256String(panelInstance.config.Auth.PasswordHmac, l.LoginRequestTokenString)
	if mac != l.LoginHmac {
		return false, errors.New("hmac not match")
	}
	return true, nil
}
