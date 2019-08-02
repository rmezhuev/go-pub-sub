package commands

import (
	"github.com/adiachenko/golang-demo/googlemq"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish random message for specified topic and routing key",
	Run: func(cmd *cobra.Command, args []string) {
		googlemq.Publish(args[0], "Lorem ipsum dolor", &googlemq.MessageAttributes{RoutingKey: args[1]})
	},
}

func init() {
	googlemq.Register(handlers)
	rootCmd.AddCommand(publishCmd)
}
