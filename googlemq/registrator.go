package googlemq

var handlers map[string]map[string]Handler

func Register(config map[string]map[string]Handler) {
	handlers = config
}

func GetHandler(subscription string, routingKey string) Handler {
	return handlers[subscription][routingKey]
}
