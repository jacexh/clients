package circle

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleW3SClient_GetWallet(t *testing.T) {
	client := NewCircleW3SClient(WithAPIKey("test-api-key"))
	// This is a placeholder test.
	// In a real-world scenario, you would mock the HTTP response.
	// For now, we just check that the method doesn't panic and returns an error
	// because we are not providing a valid wallet ID.
	_, err := client.GetWallet(context.Background(), "some-id")
	assert.Error(t, err)
}

func TestCircleW3SClient_GetWalletTokenBalance(t *testing.T) {
	client := NewCircleW3SClient(WithAPIKey("test-api-key"))
	// This is a placeholder test.
	// In a real-world scenario, you would mock the HTTP response.
	// For now, we just check that the method doesn't panic and returns an error
	// because we are not providing a valid wallet ID.
	_, err := client.GetWalletTokenBalance(context.Background(), "some-id")
	assert.Error(t, err)
}
