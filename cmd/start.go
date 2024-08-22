package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tenbounce/api"
	"tenbounce/repository"
	"tenbounce/util"
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
			var nower = util.NewTimeNower()
			repo = repository.NewMemoryRepository(nower)
		case "postgres":
			var dataSourceName = viper.GetString("postgres.data_source_name")
			repo = repository.NewPostgresRepository(dataSourceName)
		default:
			panic("invalid repository")
		}

		var signingSecret = viper.GetString("signing_secret")
		if signingSecret == "" {
			panic("invalid signing secret")
		}

		var userSecretsJSON = viper.GetString("user_secrets_json")

		var tempHardcodedUsers = []api.UserWithSecretURL{}

		if err := json.Unmarshal([]byte(userSecretsJSON), &tempHardcodedUsers); err != nil {
			panic(fmt.Errorf("unable to unmarshal user secrets: %w", err))
		}

		var apiServer = api.NewTenbounceAPI(repo, signingSecret, tempHardcodedUsers)
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
