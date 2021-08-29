package config

import (
	"gopkg.in/yaml.v2"
	"razanaubot/pkg/config/secret"
)

type Telegram struct {
	Token     string `yaml:"token"`
	ChannelID int64  `yaml:"channelID"`
}

type Config struct {
	Telegram Telegram `yaml:"telegram"`
}

func Get() (Config, error) {
	var cfg Config
	err := yaml.Unmarshal(secret.SecretFile, &cfg)
	return cfg, err
}
