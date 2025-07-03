package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	RequestRequestTestnetTokens struct {
		Address string `json:"address"`
		Chain   string `json:"chain"`
	}

	ResponseRequestTestnetTokens struct {
		PartialResponse
		Data *Transaction `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) RequestTestnetTokens(ctx context.Context, req *RequestRequestTestnetTokens) (*Transaction, error) {
	ret := new(ResponseRequestTestnetTokens)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/faucet/tokens", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
