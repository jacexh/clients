package cctp_test

import (
	"context"
	"testing"

	"github.com/jacexh/clients/circle/cctp"
)

func TestGetPublicKeys(t *testing.T) {
	client := cctp.NewCCTPV2Client()

	// Test getting public keys
	publicKeys, err := client.GetPublicKeys(context.Background())
	if err != nil {
		t.Fatalf("Failed to get public keys: %v", err)
	}

	if len(publicKeys.PublicKeys) == 0 {
		t.Fatal("Expected at least one public key, got none")
	}

	for _, pk := range publicKeys.PublicKeys {
		if pk.PublicKey == "" || pk.Version <= 0 {
			t.Errorf("Invalid public key or version: %+v", pk)
		}
	}
}
