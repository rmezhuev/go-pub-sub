package googlemq

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

func Pull(topicName string) error {
	ctx := context.Background()
	sub := initSubscription(topicName)

	fmt.Println("[] waiting for messages")
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		attr := extractMessageAttributes(msg.Attributes)
		Handle(topicName, msg, &attr)
	})
	if err != nil {
		LogError(err, "Error on receiving message")
	}

	return nil
}

func initSubscription(topicName string) *pubsub.Subscription {
	ctx := context.Background()
	client := Client()
	subName := composeSubscriptionName(topicName)
	sub := client.Subscription(subName)

	if exists, _ := sub.Exists(ctx); !exists {
		sub, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
			Topic: client.Topic(topicName),
		})
		HandleError(err, "Unable to create subscription")
		fmt.Printf("Created subscription: %v\n", sub)
		return sub
	}

	return sub
}

func Handle(topicName string, msg *pubsub.Message, attr *MessageAttributes) {

	handler := GetHandler(topicName, attr.RoutingKey)

	//message for different subscription no matching handler or version we just auto ack it here
	if handler == nil || attr.Version != config.Version {
		msg.Ack()
		fmt.Printf("Auto acked not relevant message %v", msg.ID)
		return
	}

	//process message in handler
	err := handler(msg)
	if err != nil {
		msg.Nack()
		LogError(err, "Message consuming failed")
		return
	}

	msg.Ack()
	fmt.Printf("Consumed message: %q\n", string(msg.ID))
}
