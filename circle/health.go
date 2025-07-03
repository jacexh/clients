package circle

import (
	"context"

	"github.com/jacexh/requests"
)

// PingResponse represents the response from the Circle W3S ping endpoint.
// It embeds PartialResponse to include common response fields.
type PingResponse struct {
	PartialResponse
}

func (client *CircleW3SClient) Ping(ctx context.Context) error {
	ret := new(PingResponse)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/ping", requests.Params{}, requests.UnmarshalJSON(ret))
	return CheckResponse(err, ret)
}
