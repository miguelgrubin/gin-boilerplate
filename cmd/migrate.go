package cmd

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrates the database to the latest version",
	Long: `Migrates the database to the latest version. This command applies all pending migrations
to ensure the database schema is up to date.`,
	Run: func(_ *cobra.Command, _ []string) {
		err := pkg.MigrateAll()
		if err != nil {
			log.Println("Error applying migration:")
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
