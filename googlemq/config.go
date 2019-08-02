package googlemq

import (
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load()

var defaultAttributes = MessageAttributes{
	Sender:  os.Getenv("PUBSUB_SERVICE_NAME"),
	Version: "master",
}

var config = Config{
	ProjectId:  os.Getenv("GOOGLE_CLOUD_PROJECT"),
	Version: os.Getenv("PUBSUB_SUBSCRIBER_VERSION"),
	Service:  os.Getenv("PUBSUB_SERVICE_NAME"),
}


