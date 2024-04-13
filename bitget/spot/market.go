package spot

import (
	"context"
	"strconv"
	"time"

	"github.com/jacexh/clients/bitget"
	"github.com/jacexh/requests"
)

// GetTickers https://www.bitget.com/zh-CN/api-doc/spot/market/Get-Tickers
func (sc *SpotClientV2) GetTickers(ctx context.Context, symbol string) (*ResponseGetTickers, error) {
	res := new(ResponseGetTickers)
	_, _, err := sc.client.GetWithContext(
		ctx, sc.baseURL+"/api/v2/spot/market/tickers",
		requests.Params{Query: requests.Any{"symbol": symbol}},
		bitget.DecodeAndCheckResponse(res),
	)
	return res, err
}

// GetCoins https://www.bitget.com/zh-CN/api-doc/spot/market/Get-Coin-List
func (sc *SpotClientV2) GetCoins(ctx context.Context, coin string) (*ResponseGetCoins, error) {
	res := new(ResponseGetCoins)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/public/coins",
		requests.Params{Query: requests.Any{"coin": coin}},
		bitget.DecodeAndCheckResponse(res),
	)
	return res, err
}

// GetSymbols https://www.bitget.com/zh-CN/api-doc/spot/market/Get-Symbols
func (sc *SpotClientV2) GetSymbols(ctx context.Context, symbol string) (*ResponseGetSymbols, error) {
	res := new(ResponseGetSymbols)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/public/symbols",
		requests.Params{Query: requests.Any{"symbol": symbol}},
		bitget.DecodeAndCheckResponse(res))
	return res, err
}

// GetVIPFeeRate https://www.bitget.com/zh-CN/api-doc/spot/market/Get-VIP-Fee-Rate
func (sc *SpotClientV2) GetVIPFeeRate(ctx context.Context) (*ResponseGetVIPFeeRate, error) {
	res := new(ResponseGetVIPFeeRate)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/market/vip-fee-rate",
		requests.Params{},
		bitget.DecodeAndCheckResponse(res))
	return res, err
}

// MergeDepth https://www.bitget.com/zh-CN/api-doc/spot/market/Merge-Orderbook
func (sc *SpotClientV2) MergeDepth(ctx context.Context, symbol string, precision string, limit string) (*ResponseMergeDepth, error) {
	res := new(ResponseMergeDepth)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/market/merge-depth",
		requests.Params{Query: requests.Any{"symbol": symbol, "precision": precision, "limit": limit}},
		bitget.DecodeAndCheckResponse(res),
	)
	return res, err
}

// GetOrderBook https://www.bitget.com/zh-CN/api-doc/spot/market/Get-Orderbook
func (sc *SpotClientV2) GetOrderBook(ctx context.Context, symbol, tp string, limit int) (*ResponseGetOrderBook, error) {
	q := requests.Any{"symbol": symbol, "type": tp}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	res := new(ResponseGetOrderBook)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/market/orderbook",
		requests.Params{Query: q}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// GetCandles https://www.bitget.com/zh-CN/api-doc/spot/market/Get-Candle-Data
func (sc *SpotClientV2) GetCandles(ctx context.Context, symbol, granularity string, start, end time.Time, limit int) (*ResponseGetCandles, error) {
	q := requests.Any{"symbol": symbol, "granularity": granularity}
	if !start.IsZero() {
		q["startTime"] = start.UnixMilli()
	}
	if !end.IsZero() {
		q["endTime"] = end.UnixMilli()
	}
	if limit > 0 {
		q["limit"] = limit
	}

	ret := new(ResponseGetCandles)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/market/candles", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(ret))
	return ret, err
}

// GetHistoryCandles https://www.bitget.com/zh-CN/api-doc/spot/market/Get-History-Candle-Data
func (sc *SpotClientV2) GetHistoryCandles(ctx context.Context, symbol, granularity string, end time.Time, limit int) (*ResponseGetCandles, error) {
	q := requests.Any{"symbol": symbol, "granularity": granularity}
	if !end.IsZero() {
		q["endTime"] = end.UnixMilli()
	}
	if limit > 0 {
		q["limit"] = limit
	}

	ret := new(ResponseGetCandles)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/market/candles", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(ret))
	return ret, err
}

// GetRecentTrades https://www.bitget.com/zh-CN/api-doc/spot/market/Get-Recent-Trades
func (sc *SpotClientV2) GetRecentTrades(ctx context.Context, symbol string, limit int) (*ResponseGetRecentTrades, error) {
	q := requests.Any{"symbol": symbol}
	if limit > 0 {
		q["limit"] = limit
	}

	res := new(ResponseGetRecentTrades)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/market/fills", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// GetMarketTrades https://www.bitget.com/zh-CN/api-doc/spot/market/Get-Market-Trades
func (sc *SpotClientV2) GetMarketTrades(ctx context.Context, symbol string, idLessThan int, start, end time.Time, limit int) (*ResponseGetRecentTrades, error) {
	q := requests.Any{"symbol": symbol}
	if idLessThan > 0 {
		q["idLessThan"] = idLessThan
	}
	if !start.IsZero() {
		q["startTime"] = start.UnixMilli()
	}
	if !end.IsZero() {
		q["endTime"] = end.UnixMilli()
	}
	if limit > 0 {
		q["limit"] = limit
	}

	res := new(ResponseGetRecentTrades)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/market/fills-history", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(res))
	return res, err
}
