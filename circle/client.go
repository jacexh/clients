package circle

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jacexh/requests"
)

// CircleW3SClient is a client for interacting with the Circle W3S API.
// API documentation: https://developers.circle.com/api-reference/w3s/common/create-subscription
type CircleW3SClient struct {
	session *requests.Session
	opt     *Options
}

type Options struct {
	baseURL string
	apiKey  string
}

type Option func(*Options)

type CircleResponse interface {
	GetError() error
}

func CheckResponse(err error, res CircleResponse) error {
	if err != nil {
		return err
	}
	return res.GetError()
}

// NewCircleW3SClient creates a new CircleW3SClient with the provided options.
func NewCircleW3SClient(opts ...Option) *CircleW3SClient {
	session := requests.NewSession(requests.WithClient(http.DefaultClient))
	client := &CircleW3SClient{session: session, opt: &Options{baseURL: "https://api.circle.com"}}

	for _, opt := range opts {
		opt(client.opt)
	}
	session.Apply(requests.WithBeforeHooks(client.withAPIKey(), client.withXRequestID()))
	return client
}

func (client *CircleW3SClient) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(client.opt)
	}
}

func (client *CircleW3SClient) withAPIKey() requests.BeforeRequestHook {
	return func(r *http.Request, _ []byte) {
		if client.opt.apiKey != "" {
			r.Header.Set("Authorization", "Bearer "+client.opt.apiKey)
		}
	}
}

func (client *CircleW3SClient) withXRequestID() requests.BeforeRequestHook {
	return func(r *http.Request, _ []byte) {
		if r.Header.Get("X-Request-Id") == "" {
			r.Header.Set("X-Request-Id", uuid.New().String())
		}
	}
}

// WithAPIKey sets the API key for the CircleW3SClient.
func WithAPIKey(apiKey string) Option {
	return func(opt *Options) {
		opt.apiKey = apiKey
	}
}

// WithBaseURL sets the base URL for the CircleW3SClient.
func WithBaseURL(baseURL string) Option {
	return func(opt *Options) {
		opt.baseURL = baseURL
	}
}
