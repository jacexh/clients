package spot

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"strconv"
	"time"

	"github.com/jacexh/clients/bitget"
	"github.com/mitchellh/mapstructure"
)

type (
	Ticker struct {
		Symbol      string     `json:"symbol"`
		High24H     float64    `json:"high24h,string"`
		Open        float64    `json:"open,string"`
		LastPrice   float64    `json:"lastPr,string"`
		Low24H      float64    `json:"low24h,string"`
		QuoteVolume float64    `json:"quoteVolume,string"`
		BaseVolume  float64    `json:"baseVolume,string"`
		USDTVolume  float64    `json:"usdtVolume,string"`
		BidPrice    float64    `json:"bidPr,string"`
		AskPrice    float64    `json:"askPr,string"`
		BidSize     float64    `json:"bidSZ,string"`
		AskSize     float64    `json:"askSz,string"`
		OpenUTC     float64    `json:"openUtc,string"`
		Timestamp   *Timestamp `json:"ts,string"`
		ChantUTC24H float64    `json:"changeUtc24h,string"`
		Change24H   float64    `json:"change24h,string"`
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
		Asks           []string   `json:"asks"`
		Bids           []string   `json:"bids"`
		Percision      string     `json:"percision"`
		Scale          string     `json:"scale"`
		IsMaxPrecision string     `json:"isMaxPercision"`
		Timestamp      *Timestamp `json:"ts,string"`
	}

	ResponseGetOrderBook struct {
		bitget.PartialHTTPResponse
		Data *OrderBook `json:"data"`
	}

	OrderBook struct {
		Asks      []string   `json:"asks"`
		Bids      []string   `json:"bids"`
		Timestamp *Timestamp `json:"ts,string"`
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
		Symbol    string     `json:"symbol"`
		TradeID   int        `json:"tradeId,string"`
		Side      string     `json:"side"`
		Price     float64    `json:"price,string"`
		Size      float64    `json:"size,string"`
		Timestamp *Timestamp `json:"ts,string"`
	}

	ResponseGetRecentTrades struct {
		bitget.PartialHTTPResponse
		Data []*Trade `json:"data"`
	}
)

type (
	RequestPlaceOrder struct {
		Symbol        string     `json:"symbol"`
		Side          string     `json:"side"`
		OrdeType      string     `json:"orderType"`
		Force         string     `json:"force"`
		Price         float64    `json:"price,string,omitempty"`
		Size          float64    `json:"size,string"`
		ClientOID     string     `json:"clientOid,omitempty"`
		TPSLType      string     `json:"tpsl,omitempty"`
		TriggerPrice  float64    `json:"triggerPrice,string,omitempty"`
		RequestTime   *Timestamp `json:"requestTime,string,omitempty"`
		ReceiveWindow string     `json:"receiveWindow,omitempty"`
	}

	ResponsePlaceOrder struct {
		bitget.PartialHTTPResponse
		Data *struct {
			OrderID   string `json:"orderId"`
			ClientOID string `json:"clientOid"`
		} `json:"data"`
	}

	RequestGetFills struct {
		Symbol     string
		OrderID    int
		Start      time.Time
		End        time.Time
		Limit      int
		IDLessThan int
	}

	ResponseGetFills struct {
		bitget.PartialHTTPResponse
		Data []*Fill `json:"data"`
	}

	Fill struct {
		UserID     string     `json:"userId"`
		Symbol     string     `json:"symbol"`
		OrderID    int        `json:"orderId,string"`
		TradeID    int        `json:"tradeId,string"`
		OrderType  string     `json:"orderType"`
		Side       string     `json:"side"`
		PriceAvg   float64    `json:"priceAvg,string"`
		Size       float64    `json:"size,string"`
		Amount     float64    `json:"amount,string"`
		CTime      *Timestamp `json:"cTime,string"`
		UTime      *Timestamp `json:"uTime,string"`
		TradeScope string     `json:"tradeScope"`
		FeeDetail  *FeeDetail `json:"feeDetail,omitempty"`
	}

	FeeDetail struct {
		Deduction         string  `json:"deduction"`
		FeeCoin           string  `json:"feeCoin"`
		TotalDeductionFee float64 `json:"totalDeductionFee,string"`
		TotalFee          float64 `json:"totalFee,string"`
	}

	RequestCancelOrder struct {
		Symbol        string `json:"symbol"`
		OrderID       int    `json:"orderId,string"`
		ClientOrderID string `json:"clientOid"`
	}

	ResponseCancelOrder struct {
		bitget.PartialHTTPResponse
		Data *OrderPair `json:"data"`
	}

	RequestBatchOrders struct {
		Symbol string `json:"symbol"`
		Orders []*struct {
			Side          string  `json:"side"`
			OrderType     string  `json:"orderType"`
			Force         string  `json:"force"`
			Price         float64 `json:"price,string"`
			Size          float64 `json:"size,string"`
			ClientOrderID string  `json:"clientOid,omitempty"`
		} `json:"orderList"`
	}

	OrderPair struct {
		OrderID       int    `json:"orderId,string"`
		ClientOrderID string `json:"clientOid"`
		ErrorMessage  string `json:"errorMsg,omitempty"`
		ErrorCode     string `json:"errorCode,omitempty"`
	}

	ResponseBatchOrder struct {
		bitget.PartialHTTPResponse
		Data *BatchOrders `json:"data"`
	}

	BatchOrders struct {
		SuccessList []*OrderPair `json:"successList"`
		FailureList []*OrderPair `json:"failureList"`
	}

	RequestBatchCancelOrders struct {
		Symbol string       `json:"symbol"`
		Orders []*OrderPair `json:"orderList"`
	}

	ResponseBatchCancelOrders = ResponseBatchOrder

	ResponseCancelSymbolOrders struct {
		bitget.PartialHTTPResponse
		Symbol string `json:"symbol"`
	}

	RequestGetOrder struct {
		OrderID       int
		ClientOrderID string
		RequestTime   time.Time
		ReceiveWindow time.Duration
	}

	Order struct {
		UserID          string          `json:"userId"`
		Symbol          string          `json:"symbol"`
		OrderID         int             `json:"orderId,string"`
		ClientOrderID   string          `json:"clientOid"`
		Price           float64         `json:"price,string"`
		Size            float64         `json:"size,string"`
		Status          string          `json:"status"`
		PriceAvg        float64         `json:"priceAvg,string"`
		BaseVolume      float64         `json:"baseVolume,string"`
		QuoteVolume     float64         `json:"quoteVolume,string"`
		EnterPointSouce string          `json:"enterPointSource"`
		CTime           *Timestamp      `json:"cTime,string"`
		UTime           *Timestamp      `json:"uTime,string"`
		OrderSource     string          `json:"orderSource"`
		FeeDetail       *OrderFeeDetail `json:"feeDetail"`
	}

	OrderFeeDetail struct {
		NewFees Fees `json:"newFees"`
		BGB     *BGB `json:"BGB"`
	}

	Fees struct {
		C                 float64 `json:"c"`
		D                 float64 `json:"d"`
		R                 float64 `json:"r"`
		T                 float64 `json:"t"`
		Detuction         bool    `json:"detution"`
		TotalDeductionFee float64 `json:"totalDetuctionFee"`
	}

	BGB struct {
		Deduction         bool    `json:"deduction"`
		FeeCoinCode       string  `json:"feeCoinCode"`
		TotalDeductionFee float64 `json:"totalDeductionFee"`
		TotalFee          float64 `json:"totalFee"`
	}

	ResponseGetOrder struct {
		bitget.PartialHTTPResponse
		Data []*Order `json:"data,omitempty"`
	}

	RequestGetUnfilledOrders struct {
		Symbol        string
		StartTime     time.Time
		EndTime       time.Time
		IDLessThan    int
		Limit         int
		OrderID       int
		TPSLType      string
		RequestTime   time.Time
		ReceiveWindow time.Duration
	}

	ResponseGetUnfilledOrders struct {
		bitget.PartialHTTPResponse
		Data []*UnfilledOrder `json:"data,omitempty"`
	}

	UnfilledOrder struct {
		UserID          string     `json:"userId"`
		Symbol          string     `json:"symbol"`
		OrderID         int        `json:"orderId,string"`
		ClientOrderID   string     `json:"clientOid"`
		Size            float64    `json:"size,string"`
		OrderType       string     `json:"orderType"`
		Side            string     `json:"side"`
		Status          string     `json:"status"`
		BasePrice       float64    `json:"basePrice,string"`
		PriceAvg        float64    `json:"priceAvg,string"`
		TriggerPrice    float64    `json:"triggerPrice"`
		TPSLType        string     `json:"tpslType"`
		BaseVolume      float64    `json:"baseVolume,string"`
		QuoteVolume     float64    `json:"quoteVolume,string"`
		EnterPointSouce string     `json:"enterPointSource"`
		CTime           *Timestamp `json:"cTime,string"`
		UTime           *Timestamp `json:"uTime,string"`
		OrderSource     string     `json:"orderSource"`
	}

	ResponseGetHistoryOrders struct {
		bitget.PartialHTTPResponse
		Data []*HistoryOrder `json:"data,omitempty"`
	}

	HistoryOrder struct {
		UserID          string          `json:"userId"`
		Symbol          string          `json:"symbol"`
		OrderID         int             `json:"orderId,string"`
		ClientOrderID   string          `json:"clientOid"`
		Size            float64         `json:"size,string"`
		OrderType       string          `json:"orderType"`
		Side            string          `json:"side"`
		Status          string          `json:"status"`
		PriceAvg        float64         `json:"priceAvg,string"`
		TriggerPrice    float64         `json:"triggerPrice"`
		TPSLType        string          `json:"tpslType"`
		BaseVolume      float64         `json:"baseVolume,string"`
		QuoteVolume     float64         `json:"quoteVolume,string"`
		EnterPointSouce string          `json:"enterPointSource"`
		CTime           *Timestamp      `json:"cTime,string"`
		UTime           *Timestamp      `json:"uTime,string"`
		OrderSource     string          `json:"orderSource"`
		FeeDetail       *OrderFeeDetail `json:"feeDetail"`
	}
)

type (
	ResponseAccountInfo struct {
		bitget.PartialHTTPResponse
		Data *Account `json:"data"`
	}

	Account struct {
		UserID      int        `json:"userId,string"`
		InviterID   int        `json:"inviterId,string"`
		ChannelCode string     `json:"channelCode"`
		Channel     string     `json:"channel"`
		IPs         string     `json:"ips"`
		Authorities []string   `json:"authorities"`
		ParentID    int        `json:"parentId"`
		TraderType  string     `json:"traderType"`
		RegisTime   *Timestamp `json:"regisTime,string"`
	}

	Asset struct {
		Coin           string     `json:"coin"`
		Available      float64    `json:"available,string"`
		Frozen         float64    `json:"frozen,string"`
		Locked         float64    `json:"locked,string"`
		LimitAvailable float64    `json:"limitAvailable,string"`
		UTime          *Timestamp `json:"uTime"`
	}

	Timestamp struct {
		T time.Time
	}

	ResponseGetAssets struct {
		bitget.PartialHTTPResponse
		Data []*Asset `json:"data"`
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

func (fee *OrderFeeDetail) UnmarshalJSON(data []byte) error {
	t := make(map[string]interface{})
	data = data[1 : len(data)-1]
	data = bytes.ReplaceAll(data, []byte("\\\""), []byte("\""))
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	return mapstructure.Decode(t, fee)
}

func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	slog.Info(string(data))
	d, err := strconv.ParseInt(string(data[1:len(data)-1]), 10, 64)
	if err != nil {
		return err
	}
	ts.T = time.UnixMilli(d)
	return nil
}

func (ts *Timestamp) MarshalJSON() ([]byte, error) {
	t := ts.T.UnixMilli()
	return []byte(strconv.FormatInt(t, 10)), nil
}
