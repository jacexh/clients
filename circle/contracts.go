package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	ResponseListContracts struct {
		PartialResponse
		Data *Contracts `json:"data,omitempty"`
	}

	Contracts struct {
		Contracts []Contract `json:"contracts"`
	}

	Contract struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Description     string `json:"description"`
		Blockchain      string `json:"blockchain"`
		ContractAddress string `json:"contractAddress"`
		ABI             string `json:"abi"`
		CreateDate      string `json:"createDate"`
		UpdateDate      string `json:"updateDate"`
	}
)

func (client *CircleW3SClient) ListContracts(ctx context.Context) (*Contracts, error) {
	ret := new(ResponseListContracts)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/contracts", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	RequestUpdateContract struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	ResponseUpdateContract struct {
		PartialResponse
		Data *Contract `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) UpdateContract(ctx context.Context, id string, req *RequestUpdateContract) (*Contract, error) {
	ret := new(ResponseUpdateContract)
	_, _, err := client.session.PatchWithContext(ctx, client.opt.baseURL+"/v2/contracts/"+id, requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

