package bitget

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jacexh/requests"
)

func DecodeAndCheckResponse(v CheckableResponse) requests.Unmarshaller {
	return func(b []byte) error {
		if err := requests.UnmarshalJSON(v)(b); err != nil {
			return err
		}
		return v.Error()
	}
}

func hmacSign(plainText []byte, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	_, err := h.Write(plainText)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func rsaSign(plainText []byte, privateKey string) string {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		panic(errors.New("bad public key"))
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hash := sha256.New()
	if _, err = hash.Write(plainText); err != nil {
		panic(err)
	}

	sign, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(sign)
}

func Sign(opt Option) requests.BeforeRequestHook {
	return func(request *http.Request, body []byte) {
		ts := time.Now().UnixNano() / 1000000
		timestamp := strconv.FormatInt(ts, 10)

		signer := requests.GetBuffer()
		defer requests.PutBuffer(signer)

		signer.WriteString(timestamp)
		signer.WriteString(request.Method)
		signer.WriteString(request.URL.RequestURI())
		if body != nil {
			signer.Write(body)
		}

		// make sign
		var sign string
		if opt.PrivateKey == "" {
			sign = hmacSign(signer.Bytes(), opt.SecretKey)
		} else {
			sign = rsaSign(signer.Bytes(), opt.PrivateKey)
		}

		// set header
		request.Header.Set(HeaderSign, sign)
		request.Header.Set(HeaderTimestamp, timestamp)
		request.Header.Set(HeaderAccessKey, opt.AccessKey)
		request.Header.Set(HeaderPassphrase, opt.Passphrase)
		request.Header.Set(HeaderContentType, "application/json")
		request.Header.Set(HeaderLocale, "en-US")
	}
}
