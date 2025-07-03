package circle

import (
	"context"
	"time"

	"github.com/jacexh/requests"
)

type (
	ResponseRetrieveMonitoredTokens struct {
		PartialResponse
		Data *MonitoredTokens `json:"data,omitempty"`
	}

	MonitoredTokens struct {
		Tokens     []Token   `json:"tokens"`
		UpdateDate time.Time `json:"updateDate"`
	}
)

func (client *CircleW3SClient) RetrieveMonitoredTokens(ctx context.Context) (*MonitoredTokens, error) {
	ret := new(ResponseRetrieveMonitoredTokens)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/tokens/monitored", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
		return ret.Data, nil
}

type (
	RequestDeleteMonitoredTokens struct {
		Tokens []Token `json:"tokens"`
	}

	ResponseDeleteMonitoredTokens struct {
		PartialResponse
	}
)

func (client *CircleW3SClient) DeleteMonitoredTokens(ctx context.Context, req *RequestDeleteMonitoredTokens) error {
	ret := new(ResponseDeleteMonitoredTokens)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/tokens/monitored/delete", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	return CheckResponse(err, ret)
}

