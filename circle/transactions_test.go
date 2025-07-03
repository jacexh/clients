package circle

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleW3SClient_GetTransaction(t *testing.T) {
	client := NewCircleW3SClient(WithAPIKey("test-api-key"))
	// This is a placeholder test.
	// In a real-world scenario, you would mock the HTTP response.
	// For now, we just check that the method doesn't panic and returns an error
	// because we are not providing a valid transaction ID.
	_, err := client.GetTransaction(context.Background(), "some-id")
	assert.Error(t, err)
}

func TestCircleW3SClient_CancelTransaction(t *testing.T) {
	client := NewCircleW3SClient(WithAPIKey("test-api-key"))
	// This is a placeholder test.
	// In a real-world scenario, you would mock the HTTP response.
	// For now, we just check that the method doesn't panic and returns an error
	// because we are not providing a valid transaction ID.
	_, err := client.CancelTransaction(context.Background(), "some-id")
	assert.Error(t, err)
}

func TestCircleW3SClient_CreateContractExecutionTransaction(t *testing.T) {
	client := NewCircleW3SClient(WithAPIKey("test-api-key"))
	// This is a placeholder test.
	// In a real-world scenario, you would mock the HTTP response.
	// For now, we just check that the method doesn't panic and returns an error
	// because we are not providing valid request parameters.
	_, err := client.CreateContractExecutionTransaction(context.Background(), &RequestCreateContractExecutionTransaction{
		WalletID:   "some-wallet-id",
		ContractID: "some-contract-id",
		MethodName: "some-method",
		Args:       []string{"arg1", "arg2"},
	})
	assert.Error(t, err)
}

func TestCircleW3SClient_ValidateAddress(t *testing.T) {
	client := NewCircleW3SClient(WithAPIKey("test-api-key"))
	// This is a placeholder test.
	// In a real-world scenario, you would mock the HTTP response.
	// For now, we just check that the method doesn't panic and returns an error
	// because we are not providing a valid address.
	_, err := client.ValidateAddress(context.Background(), &RequestValidateAddress{
		Address: "not-a-valid-address",
	})
	assert.Error(t, err)
}
