package config

import (
	"io/ioutil"
	"log"

	"github.com/Armatorix/SlackUp/x/xos"
	"github.com/Armatorix/SlackUp/x/xpath"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Slack SlackConfig `yaml:"slack"`
}

type SlackConfig struct {
	ClientId          string `yaml:"client_id"`
	ClientSecret      string `yaml:"client_secret"`
	VerificationToken string `yaml:"verification_token"`
	AuthToken         string `yaml:"auth_token"`
	Channel           string `yaml:"channel"`
}

func New(cfgPath string) (*Config, error) {
	cfgPath = xos.ResolveEnvs(cfgPath)

	cfgPath, err := xpath.HomeToAbsolute(cfgPath)
	if err != nil {
		log.Fatalf("failed to create absolute path: %v", err)
	}

	cfgFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, errors.Wrap(err, "read config file")
	}

	cfg := &Config{}
	err = yaml.Unmarshal(cfgFile, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal config file")
	}
	return cfg, nil
}
