package spot

import (
	"context"
	"strconv"

	"github.com/jacexh/clients/bitget"
	"github.com/jacexh/requests"
)

// PlaceOrder https://www.bitget.com/zh-CN/api-doc/spot/trade/Place-Order
func (sc *SpotClientV2) PlaceOrder(ctx context.Context, req *RequestPlaceOrder) (*ResponsePlaceOrder, error) {
	res := new(ResponsePlaceOrder)
	_, _, err := sc.client.PostWithContext(ctx, bitget.BaseURL+"/api/v2/spot/trade/place-order",
		requests.Params{Json: req}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// CancelOrder https://www.bitget.com/zh-CN/api-doc/spot/trade/Cancel-Order
func (sc *SpotClientV2) CancelOrder(ctx context.Context, req *RequestCancelOrder) (*ResponseCancelOrder, error) {
	res := new(ResponseCancelOrder)
	_, _, err := sc.client.PostWithContext(ctx, bitget.BaseURL+"/api/v2/spot/trade/cancel-order", requests.Params{Json: req}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// BatchOrders https://www.bitget.com/zh-CN/api-doc/spot/trade/Batch-Place-Orders
func (sc *SpotClientV2) BatchOrders(ctx context.Context, req *RequestBatchOrders) (*ResponseBatchOrder, error) {
	res := new(ResponseBatchOrder)
	_, _, err := sc.client.PostWithContext(ctx, sc.baseURL+"/api/v2/spot/trade/batch-orders", requests.Params{Json: req}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// BatchCancelOrders https://www.bitget.com/zh-CN/api-doc/spot/trade/Batch-Cancel-Orders
func (sc *SpotClientV2) BatchCancelOrders(ctx context.Context, req *RequestBatchCancelOrders) (*ResponseBatchCancelOrders, error) {
	res := new(ResponseBatchCancelOrders)
	_, _, err := sc.client.PostWithContext(ctx, sc.baseURL+"/api/v2/spot/trade/batch-cancel-order", requests.Params{Json: req}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// CancelSymbolOrder https://www.bitget.com/zh-CN/api-doc/spot/trade/Cancel-Symbol-Orders
func (sc *SpotClientV2) CancelSymbolOrders(ctx context.Context, symbol string) (*ResponseCancelSymbolOrders, error) {
	res := new(ResponseCancelSymbolOrders)
	_, _, err := sc.client.PostWithContext(ctx, sc.baseURL+"/api/v2/spot/trade/cancel-symbol-order", requests.Params{Json: map[string]string{"symbol": symbol}}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// GetFills https://www.bitget.com/zh-CN/api-doc/spot/trade/Get-Fills
func (sc *SpotClientV2) GetFills(ctx context.Context, req *RequestGetFills) (*ResponseGetFills, error) {
	q := map[string]string{"symbol": req.Symbol}
	if req.OrderID > 0 {
		q["orderId"] = strconv.Itoa(req.OrderID)
	}
	if !req.Start.IsZero() {
		q["startTime"] = strconv.Itoa(int(req.Start.UnixMilli()))
	}
	if !req.End.IsZero() {
		q["endTime"] = strconv.Itoa(int(req.End.UnixMilli()))
	}
	if req.Limit > 0 {
		q["limit"] = strconv.Itoa(req.Limit)
	}
	if req.IDLessThan > 0 {
		q["idLessThan"] = strconv.Itoa(req.IDLessThan)
	}

	res := new(ResponseGetFills)
	_, _, err := sc.client.GetWithContext(ctx, bitget.BaseURL+"/api/v2/spot/trade/fills",
		requests.Params{Query: q}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

func (sc *SpotClientV2) GetOrder(ctx context.Context, req *RequestGetOrder) (*ResponseGetOrder, error) {
	q := make(map[string]string)
	if req.OrderID > 0 {
		q["orderId"] = strconv.Itoa(req.OrderID)
	}
	if req.ClientOrderID != "" {
		q["clientOid"] = req.ClientOrderID
	}
	if !req.RequestTime.IsZero() {
		q["requestTime"] = strconv.FormatInt(req.RequestTime.UnixMilli(), 10)
	}
	if req.ReceiveWindow > 0 {
		q["receiveWindow"] = strconv.FormatInt(req.ReceiveWindow.Milliseconds(), 10)
	}

	ret := new(ResponseGetOrder)
	_, _, err := sc.client.GetWithContext(ctx, sc.baseURL+"/api/v2/spot/trade/orderInfo", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(ret))
	return ret, err
}

// GetUnfilledOrders https://www.bitget.com/zh-CN/api-doc/spot/trade/Get-Unfilled-Orders
func (sc *SpotClientV2) GetUnfilledOrders(ctx context.Context, req *RequestGetUnfilledOrders) (*ResponseGetUnfilledOrders, error) {
	q := make(map[string]string)
	if req.Symbol != "" {
		q["symbol"] = req.Symbol
	}
	if !req.StartTime.IsZero() {
		q["startTime"] = strconv.FormatInt(req.StartTime.UnixMilli(), 10)
	}
	if !req.EndTime.IsZero() {
		q["endTime"] = strconv.FormatInt(req.EndTime.UnixMilli(), 10)
	}
	if req.IDLessThan > 0 {
		q["idLessThan"] = strconv.Itoa(req.IDLessThan)
	}
	if req.Limit > 0 {
		q["limit"] = strconv.Itoa(req.Limit)
	}
	if req.OrderID > 0 {
		q["orderId"] = strconv.Itoa(req.OrderID)
	}
	if req.TPSLType != "" {
		q["tpslType"] = req.TPSLType
	}
	if !req.RequestTime.IsZero() {
		q["requestTime"] = strconv.FormatInt(req.RequestTime.UnixMilli(), 10)
	}
	if req.ReceiveWindow > 0 {
		q["receiveWindow"] = strconv.FormatInt(req.ReceiveWindow.Milliseconds(), 10)
	}

	res := new(ResponseGetUnfilledOrders)
	_, _, err := sc.client.GetWithContext(ctx, bitget.BaseURL+"/api/v2/spot/trade/unfilled-orders", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// GetUnfilledOrders https://www.bitget.com/zh-CN/api-doc/spot/trade/Get-History-Orders
func (sc *SpotClientV2) GetHistoryOrders(ctx context.Context, req *RequestGetUnfilledOrders) (*ResponseGetHistoryOrders, error) {
	q := make(map[string]string)
	if req.Symbol != "" {
		q["symbol"] = req.Symbol
	}
	if !req.StartTime.IsZero() {
		q["startTime"] = strconv.FormatInt(req.StartTime.UnixMilli(), 10)
	}
	if !req.EndTime.IsZero() {
		q["endTime"] = strconv.FormatInt(req.EndTime.UnixMilli(), 10)
	}
	if req.IDLessThan > 0 {
		q["idLessThan"] = strconv.Itoa(req.IDLessThan)
	}
	if req.Limit > 0 {
		q["limit"] = strconv.Itoa(req.Limit)
	}
	if req.OrderID > 0 {
		q["orderId"] = strconv.Itoa(req.OrderID)
	}
	if req.TPSLType != "" {
		q["tpslType"] = req.TPSLType
	}
	if !req.RequestTime.IsZero() {
		q["requestTime"] = strconv.FormatInt(req.RequestTime.UnixMilli(), 10)
	}
	if req.ReceiveWindow > 0 {
		q["receiveWindow"] = strconv.FormatInt(req.ReceiveWindow.Milliseconds(), 10)
	}

	res := new(ResponseGetHistoryOrders)
	_, _, err := sc.client.GetWithContext(ctx, bitget.BaseURL+"/api/v2/spot/trade/history-orders", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(res))
	return res, err
}
