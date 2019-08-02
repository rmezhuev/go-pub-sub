package commands

import (
	"github.com/adiachenko/golang-demo/features"
	"github.com/adiachenko/golang-demo/googlemq"

	"github.com/spf13/cobra"
)

// each service will have handlers map here
// handlers["topic"]["routingKey"] = handler
var handlers = map[string]map[string]googlemq.Handler{
	"events.talent": {
		"routing-key-1": features.Feature1,
		"routing-key-2": features.Feature2,
	},
	"another.topic": {
		"routing-key-3": features.Feature3,
		"routing-key-4": features.Feature4,
	},
}

var subscriberCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "Run subscriber for specified topic subscription",
	Run: func(cmd *cobra.Command, args []string) {
		googlemq.Pull(args[0])
	},
}

func init() {
	googlemq.Register(handlers)
	rootCmd.AddCommand(subscriberCmd)
}
