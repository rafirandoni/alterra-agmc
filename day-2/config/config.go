package config

type Config struct {
	AppPort string
}

func Load() *Config {
	return &Config{
		AppPort: "8080",
	}
}
