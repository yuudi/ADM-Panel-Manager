package main

type Configuration struct {
	MasterPasswordHmac string `json:"master_password_hmac"`
	SecondFactorAuth   struct {
		Enabled bool `json:"enabled"`
	} `json:"second_factor_auth"`
}

type ContainerManager struct {
}

type SiteManager struct{}

func (c *Configuration) Load(path string) error {
	return nil
}
