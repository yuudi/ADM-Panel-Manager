package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Auth struct {
		PasswordHmac string `json:"password_hmac"`
		PasswordSalt string `json:"password_salt"`
	} `json:"auth"`
	SecondFactorAuth struct {
		TOTP struct {
			Enabled bool   `json:"enabled"`
			Key     string `json:"key"`
		} `json:"totp"`
		WebAuthn struct {
			Enabled bool `json:"enabled"`
		} `json:"webauthn"`
	} `json:"second_factor_auth"`
}

func (c *Configuration) Load(path string) error {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileContent, &c)
}
