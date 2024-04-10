package spot

import (
	"context"

	"github.com/jacexh/clients/bitget"
	"github.com/jacexh/requests"
)

func (sc *SpotClientV2) PlaceOrder(ctx context.Context, req *RequestPlaceOrder) (*ResponsePlaceOrder, error) {
	res := new(ResponsePlaceOrder)
	_, _, err := sc.client.PostWithContext(ctx, bitget.BaseURL+"/api/v2/spot/trade/place-order",
		requests.Params{Json: req}, bitget.DecodeAndCheckResponse(res))
	return res, err
}
