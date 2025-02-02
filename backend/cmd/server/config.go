package main

import (
	"errors"
	"log/slog"

	"github.com/BurntSushi/toml"
)

type envConfig struct {
	Server map[string]*envConf
}

type envConf struct {
	DB struct {
		Name        string `toml:"name"`
		Host        string `toml:"host"`
		PasswordEnv string `toml:"password-env"`
		User        string `toml:"username"`
	} `toml:"db"`
	Secrets struct {
		AccessTokenEnv  string `toml:"access-token-env"`
		RefreshTokenEnv string `toml:"refresh-token-env"`
		EncryptTokenEnv string `toml:"encrypt-token-env"`
	}
}

type config struct {
	cfg *envConf
	env string
}

func NewConfig(env string) (*config, error) {
	if env != "prod" {
		slog.Info("Using development environment!")
		env = "development"
	} else {
		slog.Info("Using production environment!")
	}

	var environmentsConfig envConfig
	_, err := toml.DecodeFile("config.toml", &environmentsConfig)
	if err != nil {
		return nil, err
	}

	if environmentsConfig.Server[env] == nil {
		return nil, errors.New("invalid config file")
	}

	return &config{
		cfg: environmentsConfig.Server[env],
		env: env,
	}, nil
}
