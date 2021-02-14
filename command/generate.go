package command

import (
	"log"

	"github.com/Armatorix/SlackUp/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate random standup note.",
	Long:    "Generate your totally random standup note.",
	Run: func(cmd *cobra.Command, args []string) {
		report := generator.Report()
		log.Println(report)
	},
}

// TODO add PM mode: sorry, cannot talk right now:
