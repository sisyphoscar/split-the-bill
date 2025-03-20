package configs

import (
	"os"
)

type AppConfig struct {
	Domain string
	Port   string
}

var App AppConfig

// LoadAppConfig 載入應用程式設定
func LoadAppConfig() {
	App = AppConfig{
		Domain: os.Getenv("APP_DOMAIN"),
		Port:   os.Getenv("APP_PORT"),
	}
}
