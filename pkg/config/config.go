package config

import (
	"gopkg.in/yaml.v2"
	"razanaubot/pkg/config/secret"
)

type Telegram struct {
	Token     string `yaml:"token"`
	ChannelID int64  `yaml:"channelID"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Key      string `yaml:"key"`
}

type Config struct {
	Telegram Telegram `yaml:"telegram"`
	Redis    Redis    `yaml:"redis"`
}

func Get() (Config, error) {
	var cfg Config
	err := yaml.Unmarshal(secret.SecretFile, &cfg)
	return cfg, err
}
