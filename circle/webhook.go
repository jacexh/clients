package circle

import (
	"context"
	"fmt"
	"time"

	"github.com/jacexh/requests"
)

type (
	ListAllNotificationSubscriptionsResponse struct {
		PartialResponse
		Data []*Subscription `json:"data,omitempty"`
	}

	Subscription struct {
		ID                string    `json:"id"`
		Name              string    `json:"name"`
		Endpoint          string    `json:"endpoint"`
		Enabled           bool      `json:"enabled"`
		Restricted        bool      `json:"restricted"`
		NotificationTypes []string  `json:"notificationTypes"`
		CreateDate        time.Time `json:"createDate"`
		UpdateDate        time.Time `json:"updateDate"`
	}

	CreateSubscriptionRequest struct {
		Endpoint          string   `json:"endpoint"`
		NotificationTypes []string `json:"notificationTypes,omitempty"`
	}

	CreateSubscriptionResponse struct {
		PartialResponse
		Data *Subscription `json:"data,omitempty"`
	}

	RetriveNotificationSubscriptionResponse struct {
		PartialResponse
		Data *Subscription `json:"data,omitempty"`
	}

	UpdateNotificationSubscriptionResponse struct {
		PartialResponse
		Data *Subscription `json:"data,omitempty"`
	}

	UpdateNotificationSubscriptionRequest struct {
		ID      string `json:"_"`
		Enabled bool   `json:"enabled"`
		Name    string `json:"name"`
	}

	DeleteNotificationSubscriptionResponse struct {
		PartialResponse
		Data any `json:"data,omitempty"`
	}

	PublicKey struct {
		ID         string    `json:"id"`
		Algorithm  string    `json:"algorithm"`
		PublicKey  string    `json:"publicKey"`
		CreateDate time.Time `json:"createDate"`
	}

	GetNotificationSignaturePublicKey struct {
		PartialResponse
		Data *PublicKey `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) ListAllNotificationSubscriptions(ctx context.Context) ([]*Subscription, error) {
	ret := new(ListAllNotificationSubscriptionsResponse)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/notifications/subscriptions", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

// CreateNotificationSubscription creates a new notification subscription.
// API documentation: https://developers.circle.com/api-reference/w3s/common/create-subscription
func (client *CircleW3SClient) CreateNotificationSubscription(ctx context.Context, req *CreateSubscriptionRequest) (*Subscription, error) {
	ret := new(CreateSubscriptionResponse)
	_, _, err := client.session.PostWithContext(
		ctx,
		client.opt.baseURL+"/v2/notifications/subscriptions",
		requests.Params{Json: req},
		requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

// RetriveNotificationSubscription retrieves a notification subscription by its ID.
func (client *CircleW3SClient) RetriveNotificationSubscription(ctx context.Context, id string) (*Subscription, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: subscription ID cannot be empty", ErrCircleAPI)
	}
	ret := new(RetriveNotificationSubscriptionResponse)
	_, _, err := client.session.GetWithContext(
		ctx,
		client.opt.baseURL+"/v2/notifications/subscriptions/"+id,
		requests.Params{},
		requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

func (client *CircleW3SClient) UpdateNotificationSubscription(ctx context.Context, req *UpdateNotificationSubscriptionRequest) (*Subscription, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("%w: subscription ID cannot be empty", ErrCircleAPI)
	}
	ret := new(UpdateNotificationSubscriptionResponse)
	_, _, err := client.session.PatchWithContext(ctx, client.opt.baseURL+"/v2/notifications/subscriptions/"+req.ID, requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

func (client *CircleW3SClient) DeleteNotionficationSubscription(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("%w: subscription ID cannot be empty", ErrCircleAPI)
	}
	ret := new(DeleteNotificationSubscriptionResponse)
	_, _, err := client.session.DeleteWithContext(ctx, client.opt.baseURL+"/v2/notifications/subscriptions/"+id, requests.Params{}, requests.UnmarshalJSON(ret))
	return CheckResponse(err, ret)
}

func (client *CircleW3SClient) GetNotificationSignaturePublicKey(ctx context.Context, id string) (*PublicKey, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: subscription ID cannot be empty", ErrCircleAPI)
	}
	ret := new(GetNotificationSignaturePublicKey)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/notifications/publicKey/"+id, requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
