package cmd

import (
	"github.com/adamgoose/gosocat/lib"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Version specifies the command version
var Version = "dev"

var rootCmd = &cobra.Command{
	Use:     "gosocat {ws-url}",
	Version: Version,
	Short:   "Writes lines sent to stdin to the given websocket connection",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		gsc, err := lib.New(args[0])
		if err != nil {
			return err
		}
		defer gsc.Close()

		return <-gsc.Start()
	},
}

// Execute executes the command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	godotenv.Load()
	viper.AutomaticEnv()
}
