package config

import (
	"os"
)

type Struct struct {
	Database         string
	DatabasePort     string
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseCAFile   string
	DatabaseCertFile string
	DatabaseKeyFile  string
	Port             string
	LogLevel         string
}

var Config = Struct{
	Database:     "127.0.0.1",
	DatabasePort: "27017",
	DatabaseName: "eventstracker",
	Port:         "9101",
	LogLevel:     "info",
}

func init() {
	if os.Getenv("DB_HOST") != "" {
		Config.Database = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PORT") != "" {
		Config.DatabasePort = os.Getenv("DB_PORT")
	}
	if os.Getenv("DB_NAME") != "" {
		Config.DatabaseName = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_USERNAME") != "" {
		Config.DatabaseUsername = os.Getenv("DB_USERNAME")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		Config.DatabasePassword = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_CA_FILE") != "" {
		Config.DatabaseCAFile = os.Getenv("DB_CA_FILE")
	}
	if os.Getenv("DB_CERT_FILE") != "" {
		Config.DatabaseCertFile = os.Getenv("DB_CERT_FILE")
	}
	if os.Getenv("DB_KEY_FILE") != "" {
		Config.DatabaseKeyFile = os.Getenv("DB_KEY_FILE")
	}
	if os.Getenv("PORT") != "" {
		Config.Port = os.Getenv("PORT")
	}
	if os.Getenv("LOG_LEVEL") != "" {
		Config.LogLevel = os.Getenv("LOG_LEVEL")
	}
}
