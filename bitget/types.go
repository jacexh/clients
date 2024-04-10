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
