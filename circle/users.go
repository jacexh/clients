package circle

import (
	"context"

	"github.com/jacexh/requests"
)

type (
	RequestCreateUser struct {
		UserID string `json:"userId"`
	}

	ResponseCreateUser struct {
		PartialResponse
		Data *User `json:"data,omitempty"`
	}

	User struct {
		ID         string `json:"id"`
		UserID     string `json:"userId"`
		CreateDate string `json:"createDate"`
		UpdateDate string `json:"updateDate"`
	}
)

func (client *CircleW3SClient) CreateUser(ctx context.Context, req *RequestCreateUser) (*User, error) {
	ret := new(ResponseCreateUser)
	_, _, err := client.session.PostWithContext(ctx, client.opt.baseURL+"/v2/users", requests.Params{Json: req}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseListUsers struct {
		PartialResponse
		Data *Users `json:"data,omitempty"`
	}

	Users struct {
		Users []User `json:"users"`
	}
)

func (client *CircleW3SClient) ListUsers(ctx context.Context) (*Users, error) {
	ret := new(ResponseListUsers)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/users", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetUser struct {
		PartialResponse
		Data *User `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) GetUser(ctx context.Context, id string) (*User, error) {
	ret := new(ResponseGetUser)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/users/"+id, requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type (
	ResponseGetMe struct {
		PartialResponse
		Data *User `json:"data,omitempty"`
	}
)

func (client *CircleW3SClient) GetMe(ctx context.Context) (*User, error) {
	ret := new(ResponseGetMe)
	_, _, err := client.session.GetWithContext(ctx, client.opt.baseURL+"/v2/user", requests.Params{}, requests.UnmarshalJSON(ret))
	if err := CheckResponse(err, ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
