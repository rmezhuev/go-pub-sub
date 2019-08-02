package features

import (
	"cloud.google.com/go/pubsub"
	"fmt"
)

func Feature1(msg *pubsub.Message) error {
	fmt.Println("Executing feature 1")
	return nil
}

func Feature2(msg *pubsub.Message) error{
	fmt.Println("Executing feature 2")
	return nil
}

func Feature3(msg *pubsub.Message) error  {
	fmt.Println("Executing feature 3")
	return nil
}

func Feature4(msg *pubsub.Message) error  {
	fmt.Println("Executing feature 4")
	return nil
}

func Feature5(msg *pubsub.Message) error  {
	fmt.Println("Executing feature 5")
	return nil
}
