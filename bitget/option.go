package bitget

type Option struct {
	AccessKey  string `json:"access_key" yaml:"access_key" toml:"access_key"`
	SecretKey  string `json:"secret_key" yaml:"secret_key" toml:"secret_key"`
	Passphrase string `json:"passphrase" yaml:"passphrase" toml:"passphrase"`
	PrivateKey string `json:"private_key" yaml:"private_key" toml:"private_key"`
}

var BaseURL = "https://api.bitget.com"
