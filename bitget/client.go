package bitget

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/jacexh/requests"
)

type (
	Client struct {
		client  *http.Client
		option  Option
		baseURL string
	}

	Option struct {
		AccessKey  string `json:"access_key" yaml:"access_key" toml:"access_key"`
		SecretKey  string `json:"secret_key" yaml:"secret_key" toml:"secret_key"`
		Passphrase string `json:"passphrase" yaml:"passphrase" toml:"passphrase"`
	}

	Response struct {
		Code    string `json:"code"`
		Message string `json:"msg"`
	}
)

const (
	HeaderAccessKey   = "ACCESS-KEY"
	HeaderTimestamp   = "ACCESS-TIMESTAMP"
	HeaderSign        = "ACCESS-SIGN"
	HeaderPassphrase  = "ACCESS-PASSPHRASE"
	HeaderContentType = "Content-Type"
	HeaderLocale      = "locale"
)

var (
	BaseURL = "https://api.bitget.com"
)

func Sign(opt Option) requests.BeforeRequestHook {
	return func(request *http.Request, body []byte) {
		ts := time.Now().UnixNano() / 1000000
		timestamp := strconv.FormatInt(ts, 10)

		signer := bytes.NewBuffer([]byte(timestamp))
		signer.WriteString(request.Method)
		signer.WriteString(request.RequestURI)
		if body != nil {
			signer.Write(body)
		}
		// make sign
		h := hmac.New(sha256.New, []byte(opt.SecretKey))
		if _, err := h.Write([]byte(signer.Bytes())); err != nil {
			panic(err)
		}
		sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

		// set header
		request.Header.Set(HeaderSign, sign)
		request.Header.Set(HeaderTimestamp, timestamp)
		request.Header.Set(HeaderAccessKey, opt.AccessKey)
		request.Header.Set(HeaderPassphrase, opt.Passphrase)
		request.Header.Set(HeaderContentType, "application/json")
		request.Header.Set(HeaderLocale, "en-US")
	}
}
