package event

import (
	"time"
)

type EventAttributes struct {
	Message   string `json:"message" example:"deployment service lambda version v0.0.0"`
	Source    string `json:"source" example:"github_action"`
	Type      string `json:"type" example:"event"`
	Priority  string `json:"priority" example:"P1"`
	RelatedId string `json:"related_id" example:"53aa35c8-e659-44b2-882f-f6056e443c99"`
	Service   string `json:"service" example:"service-event"`
	Status    string `json:"status" example:"success"`
}

type EventMetadata struct {
	CreatedAt time.Time `json:"created_at"`
	Duration  int       `json:"duration"`
	Id        string    `json:"id"`
}

type EventLinks struct {
	PullRequestLink string `json:"pull_request_link" example:"https://github.com/jplanckeel/events-tracker/pull/240"`
}

type Event struct {
	Title      string `json:"title" example:"Deployment service lambda"`
	Attributes EventAttributes
	Links      EventLinks
	Metadata   EventMetadata
}

type EventList struct {
	Events []Event
}
