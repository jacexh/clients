package spot

import "github.com/jacexh/clients/bitget"

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
		Data []Ticker `json:"data"`
	}

	Coin struct {
		CoinID   string  `json:"coinId"`
		Coin     string  `json:"coin"`
		Transfer bool    `json:"transfer,string"`
		Chains   []Chain `json:"chains"`
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

	ResponseGetCoin struct {
		bitget.PartialHTTPResponse
		Data []Coin `json:"data"`
	}
)
