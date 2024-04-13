package spot

import (
	"context"

	"github.com/jacexh/clients/bitget"
	"github.com/jacexh/requests"
)

// GetAccountInfo https://www.bitget.com/zh-CN/api-doc/spot/account/Get-Account-Info
func (sc *SpotClientV2) GetAccountInfo(ctx context.Context) (*ResponseAccountInfo, error) {
	res := new(ResponseAccountInfo)
	_, _, err := sc.client.GetWithContext(ctx, bitget.BaseURL+"/api/v2/spot/account/info", requests.Params{}, bitget.DecodeAndCheckResponse(res))
	return res, err
}

// GetAssets https://www.bitget.com/zh-CN/api-doc/spot/account/Get-Account-Assets
func (sc *SpotClientV2) GetAssets(ctx context.Context, coin, assetType string) (*ResponseGetAssets, error) {
	q := requests.Any{"coin": coin, "assetType": assetType}
	res := new(ResponseGetAssets)
	_, _, err := sc.client.GetWithContext(ctx, bitget.BaseURL+"/api/v2/spot/account/assets", requests.Params{Query: q}, bitget.DecodeAndCheckResponse(res))
	return res, err
}
