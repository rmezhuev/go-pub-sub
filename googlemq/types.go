package googlemq

import "cloud.google.com/go/pubsub"

type MessageAttributes struct {
	TraceId    string `json:"trace_id"`
	SpanId     string `json:"span_id"`
	RoutingKey string `json:"routing_key"`
	Sender     string `json:"sender"`
	Version    string `json:"version"`
}

type Config struct {
	Service string
	ProjectId string
	Version   string
}

type Handler func(msg *pubsub.Message) error
