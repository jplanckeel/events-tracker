package services

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/jplanckeel/events-tracker/pkg/apis/event/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var metricEvents = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "events_tracker_services",
	Help: "Number of events in events-tracker by service and type.",
},
	[]string{"service", "type", "priority", "status"})

type EventOutput struct {
	Title      string `json:"title" example:"Deployment service lambda"`
	Attributes v1.EventAttributes
	Links      v1.EventLinks
	Metadata   v1.EventMetadata
}

type EventListOutput []EventOutput

type IEventService interface {
	Create(event *v1.Event) (*EventOutput, error)
	List() (*EventListOutput, error)
	Count() (int64, error)
	GetId(id string) (*v1.Event, error)
}

type EventService struct {
	EventRepository v1.EventInterface
}

func NewEventService(EventRepository v1.EventInterface) IEventService {
	return &EventService{
		EventRepository: EventRepository,
	}
}

func (e *EventService) Create(event *v1.Event) (*EventOutput, error) {
	if event.Attributes.RelatedId != "" {
		// attributes.relatedId is present
		relatedEvent, err := e.EventRepository.Get(context.Background(), map[string]interface{}{"metadata.id": &event.Attributes.RelatedId})
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				return nil, fmt.Errorf("no event found in events-tracker for attributes.related_id %s", event.Attributes.RelatedId)
			}
			return nil, err
		}
		event.Metadata.Duration = time.Since(relatedEvent.Metadata.CreatedAt).Seconds()

	}

	create, err := e.EventRepository.Create(context.Background(), event)
	if err != nil {
		return nil, err
	}

	go metricCreate(event)

	r := &EventOutput{
		Title: create.Title,
		Attributes: v1.EventAttributes{
			Message:   create.Attributes.Message,
			Source:    create.Attributes.Source,
			Type:      create.Attributes.Type,
			Priority:  create.Attributes.Priority,
			RelatedId: create.Attributes.RelatedId,
			Service:   create.Attributes.Service,
			Status:    create.Attributes.Status,
		},
		Links: v1.EventLinks{
			PullRequestLink: create.Links.PullRequestLink,
		},
		Metadata: v1.EventMetadata{
			CreatedAt: create.Metadata.CreatedAt,
			Duration:  create.Metadata.Duration,
			Id:        create.Metadata.Id,
		},
	}

	return r, nil
}

func (e *EventService) List() (*EventListOutput, error) {
	events, err := e.EventRepository.List(context.Background())
	if err != nil {
		return nil, err
	}

	result := make(EventListOutput, 0)
	for _, event := range events {
		result = append(result, EventOutput(event))
	}

	return &result, nil
}

func (e *EventService) Count() (int64, error) {
	count, err := e.EventRepository.Count(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (e *EventService) GetId(id string) (*v1.Event, error) {
	event, err := e.EventRepository.Get(context.Background(), map[string]interface{}{"metadata.id": id})
	if err != nil {
		return nil, fmt.Errorf("no event found in events-tracker for id %s", id)
	}

	return event, nil
}

func metricCreate(e *v1.Event) {
	prometheus.Register(metricEvents)
	metricEvents.WithLabelValues(e.Attributes.Service, e.Attributes.Type, e.Attributes.Priority, e.Attributes.Status).Add(1)
}
