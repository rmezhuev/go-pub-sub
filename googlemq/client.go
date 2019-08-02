package googlemq

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func Client() *pubsub.Client {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, config.ProjectId)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}

	return client
}
