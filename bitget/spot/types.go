package spot

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jacexh/clients/bitget"
)

type (
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

	ResponseGetTickers struct {
		bitget.PartialHTTPResponse
		Data []*Ticker `json:"data"`
	}

	Coin struct {
		CoinID   string   `json:"coinId"`
		Coin     string   `json:"coin"`
		Transfer bool     `json:"transfer,string"`
		Chains   []*Chain `json:"chains"`
	}
	Chain struct {
		Chain             string  `json:"chain"`
		NeedTag           bool    `json:"needTag,string"`
		Withdrawable      bool    `json:"withdrawable,string"`
		Rechargeable      bool    `json:"rechargeable,string"`
		WithdrawFee       float64 `json:"withdrawFee,string"`
		ExtraWithdrawFee  float64 `json:"extraWithdrawFee,string"`
		DepositConfirm    float64 `json:"depositConfirm,string"`
		WithdrawConfirm   float64 `json:"withdrawConfirm,string"`
		MinDepositAmount  float64 `json:"minDepositAmount,string"`
		MinWithdrawAmount float64 `json:"minWithdrawAmount,string"`
		BrowserURL        string  `json:"browserUrl"`
		ContractAddress   string  `json:"contractAddress"`
		WithdrawStep      int     `json:"withdrawStep,string"`
	}

	ResponseGetCoins struct {
		bitget.PartialHTTPResponse
		Data []*Coin `json:"data"`
	}

	ResponseGetSymbols struct {
		bitget.PartialHTTPResponse
		Data []*Symbol `json:"data"`
	}

	Symbol struct {
		Symbol               string  `json:"symbol"`
		BaseCoin             string  `json:"baseCoin"`
		QuoteCoin            string  `json:"quoteCoin"`
		MinTradeAmount       float64 `json:"minTradeAmount,string"`
		MaxTradeAmount       float64 `json:"maxTradeAmount,string"`
		TakerFeeRate         float64 `json:"takerFeeRate,string"`
		MakerFeeRate         float64 `json:"makerFeeRate,string"`
		PricePrecision       int     `json:"pricePrecision,string"`
		QuantityPrecision    int     `json:"quantityPrecision,string"`
		MinTradeUSDT         float64 `json:"minTradeUSDT,string"`
		Status               string  `json:"status"`
		BuyLimitPriceRation  float64 `json:"buyLimitPriceRation"`
		SellLimitPriceRation float64 `json:"sellLimitPriceRation"`
	}

	ResponseGetVIPFeeRate struct {
		bitget.PartialHTTPResponse
		Data []*VIPFeeRate `json:"data"`
	}

	VIPFeeRate struct {
		Level              int     `json:"level"`
		DealAmount         float64 `json:"dealAmount,string"`
		AssetAmount        float64 `json:"assetAmount,string"`
		TakerFeeRate       float64 `json:"takerFeeRate,string"`
		MakerFeeRate       float64 `json:"makerFeeRate,string"`
		BTCWithdrawAmount  float64 `json:"btcWithdrawAmount,string"`
		USDTWithdrawAmount float64 `json:"usdtWithdrawAmount,string"`
	}

	ResponseMergeDepth struct {
		bitget.PartialHTTPResponse
		Data *MergeDepth `json:"data"`
	}

	MergeDepth struct {
		Asks           []string `json:"asks"`
		Bids           []string `json:"bids"`
		Percision      string   `json:"percision"`
		Scale          string   `json:"scale"`
		IsMaxPrecision string   `json:"isMaxPercision"`
		Timestamp      int64    `json:"ts,string"`
	}

	ResponseGetOrderBook struct {
		bitget.PartialHTTPResponse
		Data *OrderBook `json:"data"`
	}

	OrderBook struct {
		Asks      []string `json:"asks"`
		Bids      []string `json:"bids"`
		Timestamp int64    `json:"ts,string"`
	}

	ResponseGetCandles struct {
		bitget.PartialHTTPResponse
		Data []*Candle `json:"data"`
	}

	Candle struct {
		Time         time.Time
		OpeningPrice float64
		HighestPrice float64
		LowestPrice  float64
		ClosingPrice float64
		BaseVolume   float64
		QuoteVolume  float64
		USDTVolume   float64
	}

	Trade struct {
		Symbol    string  `json:"symbol"`
		TradeID   int64   `json:"tradeId,string"`
		Side      string  `json:"side"`
		Price     float64 `json:"price,string"`
		Size      float64 `json:"size,string"`
		Timestamp int64   `json:"ts,string"`
	}

	ResponseGetRecentTrades struct {
		bitget.PartialHTTPResponse
		Data []*Trade `json:"data"`
	}
)

type (
	RequestPlaceOrder struct {
		Symbol        string  `json:"symbol"`
		Side          string  `json:"side"`
		OrdeType      string  `json:"orderType"`
		Force         string  `json:"force"`
		Price         float64 `json:"price,string,omitempty"`
		Size          float64 `json:"size,string"`
		ClientOID     string  `json:"clientOid,omitempty"`
		TPSLType      string  `json:"tpsl,omitempty"`
		TriggerPrice  float64 `json:"triggerPrice,string,omitempty"`
		RequestTime   int64   `json:"requestTime,string,omitempty"`
		ReceiveWindow string  `json:"receiveWindow,omitempty"`
	}

	ResponsePlaceOrder struct {
		bitget.PartialHTTPResponse
		Data *struct {
			OrderID   string `json:"orderId"`
			ClientOID string `json:"clientOid"`
		} `json:"data"`
	}
)

func (candle *Candle) UnmarshalJSON(data []byte) error {
	strings := make([]string, 8)
	if err := json.Unmarshal(data, &strings); err != nil {
		return err
	}
	ts, err := strconv.ParseInt(strings[0], 10, 64)
	if err != nil {
		return err
	}
	candle.Time = time.UnixMilli(ts)

	candle.OpeningPrice, _ = strconv.ParseFloat(strings[1], 64)
	candle.HighestPrice, _ = strconv.ParseFloat(strings[2], 64)
	candle.LowestPrice, _ = strconv.ParseFloat(strings[3], 64)
	candle.ClosingPrice, _ = strconv.ParseFloat(strings[4], 64)
	candle.BaseVolume, _ = strconv.ParseFloat(strings[5], 64)
	candle.QuoteVolume, _ = strconv.ParseFloat(strings[6], 64)
	candle.USDTVolume, _ = strconv.ParseFloat(strings[7], 64)
	return nil
}
