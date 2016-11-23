package cmd

import "github.com/spf13/cobra"

var cmdGetDatastreams = &cobra.Command{
	Use:   "datastreams",
	Short: "Get SensorThing DataStreams: sti get datastreams",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"Iot_id", "Name"}

		getSTEntitys("Datastreams", fields)
	},
}
