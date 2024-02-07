package config

import "github.com/mohamadmirzaeidev/phone-book/pkg/logger"

func defaultConfig() *Config{
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level: "debug",
			Encoding: "console",
		},
	}
}