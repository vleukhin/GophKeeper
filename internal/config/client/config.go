package client

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App     `yaml:"app"`
		Server  `yaml:"server"`
		Log     `yaml:"logger"`
		Storage `yaml:"sqlite"`
	}

	App struct {
		Name    string `yaml:"name"    env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	Server struct {
		URL string `yaml:"url" env:"SERVER_URL"`
	}

	Log struct {
		Level string `yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Storage struct {
		Driver   string `yaml:"driver"`
		Postgres struct {
			DSN string `yaml:"DSN"`
		}
	}
)

var (
	currentConfig *Config   //nolint:gochecknoglobals // pattern singleton
	once          sync.Once //nolint:gochecknoglobals // pattern singleton
)

// LoadConfig returns app config.
func LoadConfig() *Config {
	var err error

	once.Do(func() {
		cfg := Config{}
		err = cleanenv.ReadConfig("./internal/config/client/config.yml", &cfg)
		if err != nil {
			log.Panicln("LoadConfig - %w", err)
		}

		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Panicln("LoadConfig - %w", err)
		}
		currentConfig = &cfg
	})

	return currentConfig
}
