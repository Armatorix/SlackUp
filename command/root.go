package command

import (
	"log"
	"os"

	"github.com/Armatorix/SlackUp/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	cfg     *config.Config
	rootCmd = &cobra.Command{Use: "slackup"}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$HOME/.slackup/config.yaml", "config file path")

	rootCmd.AddCommand(generateCmd)
}

func initConfig() {
	var err error
	cfg, err = config.New(cfgFile)
	if err != nil {
		log.Fatalf("failed config load: %v", err)
	}
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("failed cmd exec: %v", err)
		os.Exit(1)
	}
}
