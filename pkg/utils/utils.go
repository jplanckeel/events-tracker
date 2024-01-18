package utils

import (
	"fmt"

	v1 "github.com/jplanckeel/events-tracker/pkg/apis/event/v1"
)

var metadataPriority = map[string]bool{
	"P1": true,
	"P2": true,
	"P3": true,
	"P4": true,
}

var metadataType = map[string]bool{
	"deployment": true,
	"incident":   true,
	"operation":  true,
}

var metadataStatus = map[string]bool{
	"start":          true,
	"failure":        true,
	"error":          true,
	"warning":        true,
	"success":        true,
	"user_update":    true,
	"recommendation": true,
	"snapshot":       true,
}

func ValidateValues(event *v1.Event) error {

	if _, ok := metadataPriority[event.Attributes.Priority]; !ok {
		return fmt.Errorf("error value %s in not allowed in attributes.priority", event.Attributes.Priority)
	}

	if _, ok := metadataType[event.Attributes.Type]; !ok {
		return fmt.Errorf("error value %s in not allowed in attributes.type", event.Attributes.Type)
	}

	if _, ok := metadataStatus[event.Attributes.Status]; !ok {
		return fmt.Errorf("error value %s in not allowed in attributes.status", event.Attributes.Status)
	}
	return nil
}
