package config

import "os"

type Config struct {
	SnowobsBaseURL string
	SnowobsToken  string
	SnowobsSource string
	APIKey        string
}

func New() Config {
	return Config{
		SnowobsBaseURL: os.Getenv("SNOWOBS_BASE_URL"),
		SnowobsToken:  os.Getenv("SNOWOBS_TOKEN"),
		SnowobsSource: os.Getenv("SNOWOBS_SOURCE"),
		APIKey:        os.Getenv("API_KEY"),
	}
}
