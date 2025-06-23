package cctp

type (
	// PublicKey represents a public key structure with its version
	PublicKey struct {
		PublicKey string `json:"publicKey"`
		Version   int    `json:"cctpVersion"`
	}

	// ResponseGetPublicKeys represents the response structure for getting public keys.
	ResponseGetPublicKeys struct {
		PublicKeys []*PublicKey `json:"publicKeys"`
	}
)
