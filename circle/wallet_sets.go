package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	ResponseListWalletSets struct {
		PartialResponse
		Data *WalletSets `json:"data,omitempty"`
	}

	WalletSets struct {
		WalletSets []WalletSet `json:"walletSets"`
	}

	WalletSet struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		CustodyType string `json:"custodyType"`
		UpdateDate  string `json:"updateDate"`
		CreateDate  string `json:"createDate"`
	}
)

func (client *CircleW3SClient) ListWalletSets(ctx context.Context) (*WalletSets, error) {
	ret := new(ResponseListWalletSets)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/walletSets", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	RequestUpdateWalletSet struct {
		Name string `json:"name"`
	}

	ResponseUpdateWalletSet struct {
		PartialResponse
		Data *WalletSet `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) UpdateWalletSet(ctx context.Context, id string, req *RequestUpdateWalletSet) (*WalletSet, error) {
	ret := new(ResponseUpdateWalletSet)
	_, _, err := client.session.PutWithContext(ctx, client.opt.baseURL+"/v2/walletSets/"+id, requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetWalletSet struct {
		PartialResponse
		Data *WalletSet `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) GetWalletSet(ctx context.Context, id string) (*WalletSet, error) {
	ret := new(ResponseGetWalletSet)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/walletSets/"+id, requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

