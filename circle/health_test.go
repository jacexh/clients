package circle_test

import (
	"context"
	"testing"

	"github.com/jacexh/clients/circle"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	client := circle.NewCircleW3SClient()
	err := client.Ping(context.Background())
	assert.NoError(t, err)
}
