package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	ResponseGetToken struct {
		PartialResponse
		Data *Token `json:"data,omitempty"`
	}

	Token struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Decimals     int    `json:"decimals"`
		Blockchain   string `json:"blockchain"`
		TokenAddress string `json:"tokenAddress"`
		Standard     string `json:"standard"`
	}
)

func (client *CircleW3SClient) GetToken(ctx context.Context, id string) (*Token, error) {
	ret := new(ResponseGetToken)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/tokens/"+id, requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
