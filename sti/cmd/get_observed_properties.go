package cmd

import "github.com/spf13/cobra"

var cmdGetObservedProperties = &cobra.Command{
	Use:   "observedproperties",
	Short: "Get SensorThing ObservedProperties: sti get observedproperties",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"Iot_id","Name"}
		getSTEntitys("ObservedProperties", fields)
	},
}