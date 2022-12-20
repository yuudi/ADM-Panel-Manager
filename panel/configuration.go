package panel

import (
	"encoding/json"
	"os"
)

var configPath = "/var/aid/config.json"

type Configuration struct {
	Auth struct {
		PasswordHmacHex string `json:"password_hmac_hex"`
		PasswordSaltHex string `json:"password_salt_hex"`
	} `json:"auth"`
	SecondFactorAuth struct {
		TOTP struct {
			Enabled bool   `json:"enabled"`
			KeyHex  string `json:"key_hex"`
		} `json:"totp"`
		WebAuthn struct {
			Enabled bool `json:"enabled"`
		} `json:"webauthn"`
	} `json:"second_factor_auth"`
}

func (c *Configuration) Load() error {
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileContent, &c)
}

func (c *Configuration) Save() error {
	fileContent, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, fileContent, 0644)
}
