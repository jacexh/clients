package spot

import (
	"context"

	"github.com/jacexh/clients/bitget"
	"github.com/jacexh/requests"
)

func (sc *SpotClientV2) GetTickers(ctx context.Context, symbol string) (*ResponseGetTickers, error) {
	res := new(ResponseGetTickers)
	_, _, err := sc.client.GetWithContext(
		ctx, sc.baseURL+"/api/v2/spot/market/tickers",
		requests.Params{Query: map[string]string{"symbol": symbol}},
		bitget.DecodeAndCheckResponse(res),
	)
	return res, err
}

func (sc *SpotClientV2) GetCoin(ctx context.Context, coin string) (*ResponseGetCoin, error) {
	res := new(ResponseGetCoin)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/public/coins",
		requests.Params{Query: map[string]string{"coin": coin}},
		bitget.DecodeAndCheckResponse(res),
	)
	return res, err
}
