package googlemq

import (
	"fmt"
	"log"
)

func HandleError(err error, msg string)  {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func FatalError(err error, msg string) {
	log.Panicf("%s: %s", msg, err)
}

func LogError(err error, msg string) {
	log.Output(1, fmt.Sprintf("%s: %s", msg, err))
}
