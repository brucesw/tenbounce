package cmd

import (
	"github.com/spf13/cobra"

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
		// TODO(bruce): read repository type from config
		if true {
			repo = repository.NewInMemoryRepository()
		} else {
			// TODO(bruce): read string from config
			var psqlInfo = "host=127.0.0.1 port=5455 user=postgresUser password=postgresPW dbname=postgresDB sslmode=disable"
			repo = repository.NewPostgresRepository(psqlInfo)
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
