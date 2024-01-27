package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	v1 "github.com/jplanckeel/events-tracker/pkg/apis/event/v1"
	"github.com/jplanckeel/events-tracker/pkg/services"
	"github.com/jplanckeel/events-tracker/pkg/utils"
)

type CreateDTO struct {
	Title      string `json:"title" example:"Deployment service lambda"`
	Attributes v1.EventAttributes
	Links      v1.EventLinks
}

type UpdateDTO struct {
	Title string `json:"title" example:"Deployement service event"`
	v1.EventAttributes
	v1.EventLinks
}

type IEventHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	GetId(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type EventHandlers struct {
	eventService services.IEventService
}

func NewEventHandlers(userService services.IEventService) IEventHandlers {
	return &EventHandlers{
		eventService: userService,
	}
}

// Create
// @Summary      Create an Event
// @Description  Create an Event
// @Tags         event
// @Accept       json
// @Produce      json
// @Param 		 data body CreateDTO true "The Request body"
// @Success      200  {object}  services.EventOutput
// @Failure      500  {object}  string
// @Router       /api/v1/event/ [post]
func (e *EventHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var query CreateDTO
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event := &v1.Event{
		Title: query.Title,
		Attributes: v1.EventAttributes{
			Message:   query.Attributes.Message,
			Source:    query.Attributes.Source,
			Type:      query.Attributes.Type,
			Priority:  query.Attributes.Priority,
			Service:   query.Attributes.Service,
			Status:    query.Attributes.Status,
			RelatedId: query.Attributes.RelatedId,
		},
		Links: v1.EventLinks{
			PullRequestLink: query.Links.PullRequestLink,
		},
		Metadata: v1.EventMetadata{},
	}

	v := validator.New()
	err = v.Struct(event)
	if err != nil {
		// Validation failed, handle the error
		err = fmt.Errorf("validation %s", err.(validator.ValidationErrors))
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = utils.ValidateValues(event)
	if err != nil {
		err = fmt.Errorf("validation %s", err)
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventCreate, err := e.eventService.Create(event)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info(
		"event created",
		"title", eventCreate.Title,
		"id", eventCreate.Metadata.Id,
		"event_message", eventCreate.Attributes.Message,
		"source", eventCreate.Attributes.Source,
		"type", eventCreate.Attributes.Type,
		"priority", eventCreate.Attributes.Priority,
		"service", eventCreate.Attributes.Service,
		"status", eventCreate.Attributes.Status,
		"pull-request", eventCreate.Links.PullRequestLink,
	)

	utils.WriteResponse(w, http.StatusCreated, eventCreate)
}

// List
// @Summary      List all Events
// @Description  List all Events
// @Tags         event
// @Accept       json
// @Produce      json
// @Success      200  {object}  services.EventListOutput
// @Failure      400  {object}  string
// @Router       /api/v1/events/list [get]
func (e *EventHandlers) List(w http.ResponseWriter, r *http.Request) {
	Events, err := e.eventService.List()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	utils.WriteResponse(w, http.StatusOK, Events)
}

// Get
// @Summary      Get an Events by id
// @Description  Get an Events by id
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id path string true  "xxxxx-xxxxxx-xxxxx-xxxxxx"
// @Success      200  {object}  services.EventOutput
// @Failure      404  {object}  string
// @Router       /api/v1/event/{id} [get]
func (e *EventHandlers) GetId(w http.ResponseWriter, r *http.Request) {
	var EventId = mux.Vars(r)["id"]

	event, err := e.eventService.GetId(EventId)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteResponse(w, http.StatusOK, event)
}

// Get
// @Summary      Search Events
// @Description  Search Events
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        source query string false  "github_action"
// @Param        type query string false  "deployment"
// @Param        priority query string false  "P1"
// @Param        status query string false  "success"
// @Param        start_date query string false  "2024-01-21T12:09"
// @Param        end_date query string false  "2024-01-21T12:09"
// @Success      200  {object}  services.EventListOutput
// @Failure      404  {object}  string
// @Router       /api/v1/events/search [get]
func (e *EventHandlers) Search(w http.ResponseWriter, r *http.Request) {

	p := &v1.EventSearch{
		Source:    r.FormValue("source"),
		EventType: r.FormValue("type"),
		Priority:  r.FormValue("priority"),
		Status:    r.FormValue("status"),
		StartDate: r.FormValue("start_date"),
		EndDate:   r.FormValue("end_date"),
	}

	event, err := e.eventService.Search(p)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteResponse(w, http.StatusOK, event)
}
