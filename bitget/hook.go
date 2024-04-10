package bitget

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log/slog"
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
		h := hmac.New(sha256.New, []byte(opt.SecretKey))
		if _, err := h.Write([]byte(signer.Bytes())); err != nil {
			slog.Error("failed to sign the request", slog.String("error", err.Error()))
			return
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
