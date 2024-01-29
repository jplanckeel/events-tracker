package utils

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"

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

func CreateFilter(p *v1.EventSearch) (map[string]interface{}, error) {

	filter := make(map[string]interface{})

	if p.Source != "" {
		filter["attributes.source"] = p.Source
	}
	if p.EventType != "" {
		filter["attributes.type"] = p.EventType
	}
	if p.Priority != "" {
		filter["attributes.priority"] = p.Priority
	}
	if p.Status != "" {
		filter["attributes.status"] = p.Status
	}
	if p.Service != "" {
		filter["attributes.service"] = p.Service
	}
	if p.StartDate != "" && p.EndDate != "" {
		s, err := parseDate(p.StartDate)
		if err != nil {
			return nil, err
		}
		e, err := parseDate(p.EndDate)
		if err != nil {
			return nil, err
		}
		err = checkDateInverted(s, e)
		if err != nil {
			return nil, err
		}
		filter["metadata.createdat"] = bson.D{{Key: "$gte", Value: s}, {Key: "$lte", Value: e}}
	}
	if p.StartDate != "" && p.EndDate == "" {
		d, err := parseDate(p.StartDate)
		if err != nil {
			return nil, err
		}
		filter["metadata.createdat"] = bson.D{{Key: "$gte", Value: d}}
	}
	if p.StartDate == "" && p.EndDate != "" {
		d, err := parseDate(p.EndDate)
		if err != nil {
			return nil, err
		}
		filter["metadata.createdat"] = bson.D{{Key: "$lte", Value: d}}
	}
	if filter == nil {
		err := errors.New("no filter for search events")
		return nil, err
	}
	return filter, nil
}

func parseDate(date string) (t time.Time, err error) {

	layouts := []string{
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05.999Z",
		"2006-01-02T15:04:05.999",
		"2006-01-02T15:04:05.999-07:00",
		"2006-01-02T15:04:05Z0700",
		"2006-01-02T15:04:05.999Z0700",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999Z07:00",
		"2006-01-02T15:04:05-07",
		"2006-01-02T15:04",
		"2006-01-02T15",
		"2006-01-02",
	}

	for _, layout := range layouts {
		t, err = time.Parse(layout, date)
		if err == nil {
			break
		}
	}
	return
}

func checkDateInverted(startDate time.Time, endDate time.Time) (err error) {
	if endDate.Before(startDate) {
		err = fmt.Errorf("start_date %s and end_date %s are inversed", startDate, endDate)
	}
	return
}

func CatchPullRequestId(input string) (id string, err error) {
	pattern := `.*\/(\d*)$`
	regex := regexp.MustCompile(pattern)
	// Trouver des correspondances dans la chaîne d'entrée
	matches := regex.FindStringSubmatch(input)

	// Vérifier si des correspondances ont été trouvées
	if len(matches) > 0 {
		id = matches[1]
	} else {
		err = fmt.Errorf("no pull request id found in %s", input)
	}
	return
}
