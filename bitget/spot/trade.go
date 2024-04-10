package spot

import (
	"context"

	"github.com/jacexh/clients/bitget"
	"github.com/jacexh/requests"
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

func (sc *SpotClientV2) PlaceOrder(ctx context.Context, req *RequestPlaceOrder) (*ResponsePlaceOrder, error) {
	res := new(ResponsePlaceOrder)
	_, _, err := sc.client.PostWithContext(ctx, bitget.BaseURL+"/api/v2/spot/trade/place-order",
		requests.Params{Json: req}, bitget.DecodeAndCheckResponse(res))
	return res, err
}
