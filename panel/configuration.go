package panel

import (
	"encoding/json"
	"os"
)

func EnvironmentVariable(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var configPath = EnvironmentVariable("ADM_CONFIG_PATH", "config.json")

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
	// test if config file is readable
	if _, err := os.Stat(configPath); err != nil {
		c.Save()
		return nil
	}

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
