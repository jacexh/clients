package bitget

import (
	"context"
	"net/http"

	"github.com/jacexh/requests"
)

type (
	SpotClientV2 struct {
		client  *requests.Session
		baseURL string
		opt     Option
	}

	Ticker struct {
		Symbol      string  `json:"symbol"`
		High24H     float64 `json:"high24h,string"`
		Open        float64 `json:"open,string"`
		LastPrice   float64 `json:"lastPr,string"`
		Low24H      float64 `json:"low24h,string"`
		QuoteVolume float64 `json:"quoteVolume,string"`
		BaseVolume  float64 `json:"baseVolume,string"`
		USDTVolume  float64 `json:"usdtVolume,string"`
		BidPrice    float64 `json:"bidPr,string"`
		AskPrice    float64 `json:"askPr,string"`
		BidSize     float64 `json:"bidSZ,string"`
		AskSize     float64 `json:"askSz,string"`
		OpenUTC     float64 `json:"openUtc,string"`
		Timestamp   int64   `json:"ts,string"`
		ChantUTC24H float64 `json:"changeUtc24h,string"`
		Change24H   float64 `json:"change24h,string"`
	}

	ResponseTicker struct {
		Response
		Data []Ticker `json:"data"`
	}
)

func NewSpotClient(opt Option) *SpotClientV2 {
	client := requests.NewSession(
		requests.WithClient(http.DefaultClient),
		requests.WithBeforeHooks(Sign(opt), requests.LogRequest(nil)),
		requests.WithAfterHooks(requests.LogResponse(nil)),
	)
	return &SpotClientV2{client: client, baseURL: BaseURL}
}

func (sc *SpotClientV2) GetTickers(ctx context.Context, symbol string) (*ResponseTicker, error) {
	res := new(ResponseTicker)
	_, _, err := sc.client.GetWithContext(
		ctx, sc.baseURL+"/api/v2/spot/market/tickers",
		requests.Params{Query: map[string]string{"symbol": symbol}}, requests.UnmarshalJSON(res),
	)
	return res, err
}
