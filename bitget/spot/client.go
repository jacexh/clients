package spot

import (
	"net/http"

	"github.com/jacexh/clients/bitget"
	"github.com/jacexh/requests"
)

type SpotClientV2 struct {
	client  *requests.Session
	baseURL string
	opt     bitget.Option
}

func NewSpotClientV2(opt bitget.Option) *SpotClientV2 {
	client := requests.NewSession(
		requests.WithClient(http.DefaultClient),
		requests.WithBeforeHooks(bitget.Sign(opt)),
	)
	return &SpotClientV2{client: client, baseURL: bitget.BaseURL, opt: opt}
}
