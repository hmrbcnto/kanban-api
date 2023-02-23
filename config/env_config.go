package config

import (
	"errors"
)

type EnvConfig struct {
	SecretKey string
}

func (envConfig EnvConfig) validate() error {
	if envConfig.SecretKey == "" {
		return errors.New("Secret Key not supplied")
	}

	return nil
}
