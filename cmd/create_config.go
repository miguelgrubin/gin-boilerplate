// Package cmd provides command line interface on the app.
package cmd

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/spf13/cobra"
)

// createConfigCmd represents the create command
var createConfigCmd = &cobra.Command{
	Use:   "create-config",
	Short: "Creates a default config file in the current folder",
	Long:  `Creates a default config file in the current folder`,
	Run: func(_ *cobra.Command, _ []string) {
		configService := services.NewConfigService()
		err := configService.WriteConfig()
		if err != nil {
			log.Println("Error writing config", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
