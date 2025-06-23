package cctp

import (
	"context"
	"net/http"

	"github.com/jacexh/requests"
)

type (
	CCTPV2Client struct {
		session *requests.Session
		option  *Options
	}

	Options struct {
		baseURL string
	}

	Option func(*Options)
)

const (
	Mainnet = "https://iris-api.circle.com"
	Testnet = "https://iris-api-sandbox.circle.com"
)

var (
	defaultOpts = &Options{
		baseURL: Mainnet,
	}
)

func NewCCTPV2Client(opts ...Option) *CCTPV2Client {
	options := defaultOpts
	for _, opt := range opts {
		opt(options)
	}

	session := requests.NewSession(requests.WithClient(http.DefaultClient))
	return &CCTPV2Client{
		session: session,
		option:  options,
	}
}

func (client *CCTPV2Client) GetPublicKeys(ctx context.Context) (*ResponseGetPublicKeys, error) {
	ret := new(ResponseGetPublicKeys)
	res, data, err := client.session.GetWithContext(ctx, client.option.baseURL+"/v2/publicKeys", requests.Params{}, requests.UnmarshalJSON(ret))
	// 调用正常
	if res.StatusCode == http.StatusOK {
		return ret, err
	}

	// 调用异常
	if res.StatusCode >= http.StatusBadRequest && res.StatusCode < http.StatusInternalServerError {
		return nil, &BadRequestResponse{Code: res.StatusCode, ErrorMessage: string(data)}
	}
	return nil, &UndefinedError{Code: res.StatusCode, ErrorMessage: string(data)}
}
