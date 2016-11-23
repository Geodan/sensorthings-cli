package cmd

import "github.com/spf13/cobra"

var cmdGetLocations = &cobra.Command{
	Use:   "locations",
	Short: "Get SensorThing Locations: sti get locations",

	Run: func(cmd *cobra.Command, args []string) {
		getSTEntitys("Locations")
	},
}
