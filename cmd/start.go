package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tenbounce/api"
	"tenbounce/repository"
)

// TODO(bruce): Update description
// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var repo api.Repository

		switch viper.GetString("repository") {
		case "memory":
			repo = repository.NewMemoryRepository()
		case "postgres":
			var dataSourceName = viper.GetString("postgres.data_source_name")
			repo = repository.NewPostgresRepository(dataSourceName)
			fmt.Println(dataSourceName)
		default:
			panic("invalid repository")
		}

		var apiServer = api.NewTenbounceAPI(repo)
		apiServer.Logger.Fatal(apiServer.Start(":1323"))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
