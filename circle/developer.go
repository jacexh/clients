package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	ResponseGetDeveloperPublicKey struct {
		PartialResponse
		Data *PublicKey `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) GetDeveloperPublicKey(ctx context.Context) (*PublicKey, error) {
	ret := new(ResponseGetDeveloperPublicKey)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/developer/publicKey", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetDeveloperConfiguration struct {
		PartialResponse
		Data *DeveloperConfiguration `json:"data,omitempty"`
	}

	DeveloperConfiguration struct {
		WebhookURL string `json:"webhookUrl"`
	}
)

func (client *CircleW3SClient) GetDeveloperConfiguration(ctx context.Context) (*DeveloperConfiguration, error) {
	ret := new(ResponseGetDeveloperConfiguration)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/developer/configuration", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
