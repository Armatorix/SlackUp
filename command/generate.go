package command

import (
	"log"

	"github.com/Armatorix/SlackUp/generator"
	"github.com/Armatorix/SlackUp/slack"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate random standup note.",
	Long:    "Generate your totally random standup note.",
	Run: func(cmd *cobra.Command, args []string) {
		report := generator.Report()
		err := slack.SendMsg(&cfg.Slack, report.String())
		if err != nil {
			log.Printf("send message failed %v\n", err)
		}
	},
}
