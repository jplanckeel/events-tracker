package event

import (
	"time"
)

type EventAttributes struct {
	Message   string `json:"message" example:"deployment service lambda version v0.0.0" validate:"required"`
	Source    string `json:"source" example:"github_action" validate:"required"`
	Type      string `json:"type" example:"deployment" validate:"required"`
	Priority  string `json:"priority" example:"P1" validate:"required"`
	RelatedId string `json:"related_id" example:"53aa35c8-e659-44b2-882f-f6056e443c99"`
	Service   string `json:"service" example:"service-event" validate:"required"`
	Status    string `json:"status" example:"success" validate:"required"`
}

type EventMetadata struct {
	CreatedAt time.Time `json:"created_at"`
	Duration  float64   `json:"duration"`
	Id        string    `json:"id"`
}

type EventLinks struct {
	PullRequestLink string `json:"pull_request_link" example:"https://github.com/jplanckeel/events-tracker/pull/240"`
}

type Event struct {
	Title      string `json:"title" example:"Deployment service lambda" validate:"required"`
	Attributes EventAttributes
	Links      EventLinks
	Metadata   EventMetadata
}

type EventSearch struct {
	Source    string
	EventType string
	Priority  string
	Status    string
	Service   string
	StartDate string
	EndDate   string
}
