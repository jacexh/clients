package bitget

import "fmt"

type (
	PartialHTTPResponse struct {
		Code    string `json:"code"`
		Message string `json:"msg"`
	}

	CheckableResponse interface {
		Error() error
	}

	Option struct {
		AccessKey  string `json:"access_key" yaml:"access_key" toml:"access_key"`
		SecretKey  string `json:"secret_key" yaml:"secret_key" toml:"secret_key"`
		Passphrase string `json:"passphrase" yaml:"passphrase" toml:"passphrase"`
	}
)

var (
	BaseURL = "https://api.bitget.com"
)

const (
	// CodeSucc success
	CodeSucc = "00000"

	HeaderAccessKey   = "ACCESS-KEY"
	HeaderTimestamp   = "ACCESS-TIMESTAMP"
	HeaderSign        = "ACCESS-SIGN"
	HeaderPassphrase  = "ACCESS-PASSPHRASE"
	HeaderContentType = "Content-Type"
	HeaderLocale      = "locale"
)

func (rp PartialHTTPResponse) Error() error {
	if rp.Code == CodeSucc {
		return nil
	}
	return fmt.Errorf("[%s] %s", rp.Code, rp.Message)
}
