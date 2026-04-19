package config

import "github.com/ilyakaznacheev/cleanenv"

type SnowobsConfig struct {
	BaseURL string `env:"SNOWOBS_BASE_URL" env-required:"true"`
	Token   string `env:"SNOWOBS_TOKEN" env-required:"true"`
	Source  string `env:"SNOWOBS_SOURCE" env-required:"true"`
}

type Config struct {
	Snowobs SnowobsConfig
	APIKey  string `env:"API_KEY" env-required:"true"`
}

func New() (Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
