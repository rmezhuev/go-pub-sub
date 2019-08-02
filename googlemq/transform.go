package googlemq

import (
	"encoding/json"
	"github.com/imdario/mergo"
)

func composeSubscriptionName(subscription string) string {
	return  config.Service + "_" + config.Version + "." + subscription
}

func compileMessageAttributes(attr *MessageAttributes) map[string]string {

	if attr == nil {
		attr = &defaultAttributes
	}

	mergo.Merge(attr, defaultAttributes)
	attrMap := make(map[string]string)
	attrStr, _ := json.Marshal(attr)
	json.Unmarshal(attrStr, &attrMap)

	return attrMap
}

func extractMessageAttributes(attrMap map[string]string) MessageAttributes {
	str, _ := json.Marshal(&attrMap)
	messageAttr := MessageAttributes{}
	json.Unmarshal(str, &messageAttr)

	return messageAttr
}
