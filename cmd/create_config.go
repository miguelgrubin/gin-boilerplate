// Package cmd provides command line interface on the app.
package cmd

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg"
	"github.com/spf13/cobra"
)

// createConfigCmd represents the create command
var createConfigCmd = &cobra.Command{
	Use:   "create-config",
	Short: "Creates a default config file in the current folder",
	Long:  `Creates a default config file in the current folder`,
	Run: func(_ *cobra.Command, _ []string) {
		app, err := pkg.NewApp()
		if err != nil {
			log.Println("Error creating app:")
			log.Fatal(err)
		}
		err = app.WriteConfig()
		if err != nil {
			log.Println("Config cloud not be created")
		} else {
			log.Println("Config file created...")
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
