package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	RequestSignMessage struct {
		WalletID string `json:"walletId"`
		Message  string `json:"message"`
	}

	ResponseSignMessage struct {
		PartialResponse
		Data *Signature `json:"data,omitempty"`
	}

	Signature struct {
		Signature string `json:"signature"`
	}
)

func (client *CircleW3SClient) SignMessage(ctx context.Context, req *RequestSignMessage) (*Signature, error) {
	ret := new(ResponseSignMessage)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/sign/message", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
		return ret.Data, nil
}

type (
	RequestSignDelegateAction struct {
		WalletID      string `json:"walletId"`
		DelegateAction any    `json:"delegateAction"`
	}

	ResponseSignDelegateAction struct {
		PartialResponse
		Data *Signature `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) SignDelegateAction(ctx context.Context, req *RequestSignDelegateAction) (*Signature, error) {
	ret := new(ResponseSignDelegateAction)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/sign/delegateAction", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

