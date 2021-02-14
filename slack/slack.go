package slack

import (
	"github.com/Armatorix/SlackUp/config"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
)

func SendMsg(cfg *config.SlackConfig, msg string) error {
	client := slack.New(cfg.AuthToken)
	_, _, err := client.PostMessage(cfg.Channel, slack.MsgOptionText(msg, false))
	if err != nil {
		return errors.Wrap(err, "slack request")
	}
	return nil
}
