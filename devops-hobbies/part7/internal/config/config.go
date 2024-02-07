package config

import "github.com/mohamadmirzaeidev/phone-book/pkg/logger"

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
