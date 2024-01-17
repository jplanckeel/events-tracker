package services

import (
	"context"
	"fmt"

	v1 "github.com/jplanckeel/events-tracker/pkg/apis/event/v1"
)

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
	create, err := e.EventRepository.Create(context.Background(), event)
	if err != nil {
		return nil, err
	}

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
	for _, event := range events.Events {
		fmt.Println(event)
		result = append(result, EventOutput(event))
	}

	return &result, nil
}

/*
func (e *EventService) Get(name string) (*EventOutput, error) {
	event, err := e.EventRepository.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	result := &EventOutput{
		Name: event.Name,
		EventSpec: v1.EventSpec{
		},
	}

	return result, nil
}
*/