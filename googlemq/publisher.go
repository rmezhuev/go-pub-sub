package googlemq

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

func Publish(topic string, msg string, attr *MessageAttributes) error {
	ctx := context.Background()
	t := Client().Topic(topic)

	attrMap := compileMessageAttributes(attr)
	fmt.Printf("%+v\n", attrMap)

	result := t.Publish(ctx, &pubsub.Message{
		Data:       []byte(msg),
		Attributes: attrMap,
	})

	id, err := result.Get(ctx)

	if err != nil {
		return err
	} else {
		fmt.Printf("Published a message; msg ID: %v\n", id)
	}
	return nil
}


