package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	ResponseListTransactions struct {
		PartialResponse
		Data *Transactions `json:"data,omitempty"`
	}

	Transactions struct {
		Transactions []Transaction `json:"transactions"`
	}

	Transaction struct {
		ID                  string   `json:"id"`
		WalletID            string   `json:"walletId"`
		State               string   `json:"state"`
		TransactionType     string   `json:"transactionType"`
		Blockchain          string   `json:"blockchain"`
		CustodyType         string   `json:"custodyType"`
		Operation           string   `json:"operation"`
		TransactionHash     string   `json:"transactionHash"`
		ExplorerURL         string   `json:"explorerUrl"`
		CreateDate          string   `json:"createDate"`
		UpdateDate          string   `json:"updateDate"`
		SourceAddress       string   `json:"sourceAddress"`
		DestinationAddress  string   `json:"destinationAddress"`
		TokenID             string   `json:"tokenId"`
		Amount              []string `json:"amount"`
		FeeLevel            string   `json:"feeLevel"`
		GasLimit            string   `json:"gasLimit"`
		GasPrice            string   `json:"gasPrice"`
		PriorityFee         string   `json:"priorityFee"`
		MaxFee              string   `json:"maxFee"`
		NetworkFee          string   `json:"networkFee"`
		FirstConfirmDate    string   `json:"firstConfirmDate"`
		EstimatedCompletion string   `json:"estimatedCompletion"`
		BlockHeight         string   `json:"blockHeight"`
		BlockHash           string   `json:"blockHash"`
		ErrorCode           string   `json:"errorCode"`
		RefID               string   `json:"refId"`
	}
)

func (client *CircleW3SClient) ListTransactions(ctx context.Context) (*Transactions, error) {
	ret := new(ResponseListTransactions)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/transactions", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
		return ret.Data, nil
}

type (
	ResponseCancelTransaction struct {
		PartialResponse
		Data *Transaction `json:"data,omitempty"`
	}
)

type (
	RequestAccelerateTransaction struct {
		FeeLevel string `json:"feeLevel,omitempty"`
		GasLimit string `json:"gasLimit,omitempty"`
		GasPrice string `json:"gasPrice,omitempty"`
	}

	ResponseAccelerateTransaction struct {
		PartialResponse
		Data *Transaction `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) AccelerateTransaction(ctx context.Context, id string, req *RequestAccelerateTransaction) (*Transaction, error) {
	ret := new(ResponseAccelerateTransaction)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/transactions/"+id+"/accelerate", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetTransaction struct {
		PartialResponse
		Data *Transaction `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) GetTransaction(ctx context.Context, id string) (*Transaction, error) {
	ret := new(ResponseGetTransaction)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/transactions/"+id, requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

func (client *CircleW3SClient) CancelTransaction(ctx context.Context, id string) (*Transaction, error) {
	ret := new(ResponseCancelTransaction)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/transactions/"+id+"/cancel", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	RequestCreateContractExecutionTransaction struct {
		WalletID      string   `json:"walletId"`
		ContractID    string   `json:"contractId"`
		MethodName    string   `json:"methodName"`
		Args          []string `json:"args"`
		FeeLevel      string   `json:"feeLevel,omitempty"`
	}

	ResponseCreateContractExecutionTransaction struct {
		PartialResponse
		Data *Transaction `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) CreateContractExecutionTransaction(ctx context.Context, req *RequestCreateContractExecutionTransaction) (*Transaction, error) {
	ret := new(ResponseCreateContractExecutionTransaction)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/transactions/contract-execution", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	RequestValidateAddress struct {
		Address string `json:"address"`
	}

	ResponseValidateAddress struct {
		PartialResponse
		Data *ValidatedAddress `json:"data,omitempty"`
	}

	ValidatedAddress struct {
		IsValid bool `json:"isValid"`
	}
)

func (client *CircleW3SClient) ValidateAddress(ctx context.Context, req *RequestValidateAddress) (*ValidatedAddress, error) {
	ret := new(ResponseValidateAddress)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/transactions/validate-address", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

