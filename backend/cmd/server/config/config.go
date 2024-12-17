package config

import "os"

type DatabaseConfig struct {
	UsernameEnv string
	PasswordEnv string
	HostEnv     string
	NameEnv     string
}

type Secrets struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	EncryptTokenSecret string
}

type Config struct {
	DB      *DatabaseConfig
	Secrets *Secrets
}

func GetConfig() *Config {
	return &Config{
		DB: &DatabaseConfig{
			UsernameEnv: os.Getenv("DB_USERNAME"),
			PasswordEnv: os.Getenv("DB_PASSWORD"),
			HostEnv:     os.Getenv("DB_HOST"),
			NameEnv:     os.Getenv("DB_NAME"),
		},
		Secrets: &Secrets{
			AccessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
			RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
			EncryptTokenSecret: os.Getenv("ENCRYPT_TOKEN_SECRET"),
		},
	}
}
