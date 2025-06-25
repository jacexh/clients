package circle

import (
	"errors"
	"fmt"
	"net/http"
)

// PartialResponse represents a partial HTTP response with a code and message.
type (
	PartialResponse struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

var ErrCircleAPI = errors.New("Circle API error")

// GetError checks the response code and returns an error if the code indicates a failure.
func (pr PartialResponse) GetError() error {
	if pr.Code == 0 || pr.Code == http.StatusOK {
		return nil
	}
	return fmt.Errorf("%w: %s (code: %d)", ErrCircleAPI, pr.Message, pr.Code)
}
