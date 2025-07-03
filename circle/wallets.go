package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	RequestCreateWallet struct {
		WalletSetID string   `json:"walletSetId"`
		Blockchains []string `json:"blockchains"`
		Count       int      `json:"count,omitempty"`
	}

	ResponseCreateWallet struct {
		PartialResponse
		Data *Wallets `json:"data,omitempty"`
	}

	Wallets struct {
		Wallets []Wallet `json:"wallets"`
	}

	Wallet struct {
		ID           string `json:"id"`
		WalletSetID  string `json:"walletSetId"`
		CustodyType  string `json:"custodyType"`
		UserID       string `json:"userId"`
		Address      string `json:"address"`
		Blockchain   string `json:"blockchain"`
		AccountType  string `json:"accountType"`
		UpdateDate   string `json:"updateDate"`
		CreateDate   string `json:"createDate"`
		State        string `json:"state"`
		StateDetails string `json:"stateDetails"`
	}
)

func (client *CircleW3SClient) CreateWallet(ctx context.Context, req *RequestCreateWallet) (*Wallets, error) {
	ret := new(ResponseCreateWallet)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/wallets", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetWalletNFTs struct {
		PartialResponse
		Data *NFTs `json:"data,omitempty"`
	}

	NFTs struct {
		NFTs []NFT `json:"nfts"`
	}

	NFT struct {
		ID           string `json:"id"`
		WalletID     string `json:"walletId"`
		TokenID      string `json:"tokenId"`
		Collection   Collection `json:"collection"`
		Standard     string `json:"standard"`
		Blockchain   string `json:"blockchain"`
		OwnerAddress string `json:"ownerAddress"`
		UpdateDate   string `json:"updateDate"`
		CreateDate   string `json:"createDate"`
	}

	Collection struct {
		Name         string `json:"name"`
		TokenAddress string `json:"tokenAddress"`
		ImageURL     string `json:"imageUrl"`
	}
)

func (client *CircleW3SClient) GetWalletNFTs(ctx context.Context, id string) (*NFTs, error) {
	ret := new(ResponseGetWalletNFTs)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/wallets/"+id+"/nfts", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetWallet struct {
		PartialResponse
		Data *Wallet `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) GetWallet(ctx context.Context, id string) (*Wallet, error) {
	ret := new(ResponseGetWallet)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/wallets/"+id, requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetWalletTokenBalance struct {
		PartialResponse
		Data *TokenBalances `json:"data,omitempty"`
	}

	TokenBalances struct {
		TokenBalances []TokenBalance `json:"tokenBalances"`
	}

	TokenBalance struct {
		Token       Token  `json:"token"`
		Amount      string `json:"amount"`
		UpdateDate  string `json:"updateDate"`
	}
)

func (client *CircleW3SClient) GetWalletTokenBalance(ctx context.Context, id string) (*TokenBalances, error) {
	ret := new(ResponseGetWalletTokenBalance)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/wallets/"+id+"/balances", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

