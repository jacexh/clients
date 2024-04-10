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

func NewSpotClient(opt bitget.Option) *SpotClientV2 {
	client := requests.NewSession(
		requests.WithClient(http.DefaultClient),
		requests.WithBeforeHooks(bitget.Sign(opt), requests.LogRequest(nil)),
		requests.WithAfterHooks(requests.LogResponse(nil)), // TODO: remove
	)
	return &SpotClientV2{client: client, baseURL: bitget.BaseURL, opt: opt}
}

