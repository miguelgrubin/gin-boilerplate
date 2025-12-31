package cmd

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/spf13/cobra"
)

// generateKeysCmd represents the generateKeys command
var generateKeysCmd = &cobra.Command{
	Use:   "generate-keys",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		configService := services.NewConfigService()
		_, err := configService.ReadConfig()
		if err != nil {
			log.Println("Error reading config", err)
			return
		}
		config := configService.GetConfig()
		rsaService := services.NewRSAService(config.Jwt.Keys.Private, config.Jwt.Keys.Public)
		rsaService.GenerateKeyPair()
		err = rsaService.Write()
		if err != nil {
			log.Println("Error generating keys:")
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateKeysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateKeysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateKeysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
