package circle_test

import (
	"context"
	"testing"

	"github.com/jacexh/clients/circle"
	"github.com/stretchr/testify/assert"
)

const StandardAPIKey = "TEST_API_KEY:c800debdd1e39985b13b57455568396f:4d809c459f1117f79e5a27612ce021d6"

func TestListAllNotificationSubscriptions(t *testing.T) {
	client := circle.NewCircleW3SClient(circle.WithAPIKey(StandardAPIKey))
	subs, err := client.ListAllNotificationSubscriptions(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, subs)
}

func TestCreateSubscription(t *testing.T) {
	client := circle.NewCircleW3SClient(circle.WithAPIKey(StandardAPIKey))

	sub, err := client.CreateNotificationSubscription(context.Background(), &circle.CreateSubscriptionRequest{
		Endpoint: "https://example.com/webhook",
	})
	assert.NoError(t, err)
	assert.NotNil(t, sub)
}
