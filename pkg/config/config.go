package config

import (
	"os"
)

type Struct struct {
	Namespace    string
	Port         string
	LabelProject string
	LogLevel     string
}

var Config = Struct{
	Namespace:    "events-tracker",
	Port:         "9101",
	LabelProject: "project",
	LogLevel:     "info",
}

func init() {
	if os.Getenv("NAMESPACE") != "" {
		Config.Namespace = os.Getenv("NAMESPACE")
	}
	if os.Getenv("PORT") != "" {
		Config.Port = os.Getenv("PORT")
	}
	if os.Getenv("LABEL_PROJECT") != "" {
		Config.LabelProject = os.Getenv("LABEL_PROJECT")
	}
	if os.Getenv("LOG_LEVEL") != "" {
		Config.LogLevel = os.Getenv("LOG_LEVEL")
	}
}
