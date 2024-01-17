package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	_ "github.com/jplanckeel/events-tracker/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jplanckeel/events-tracker/pkg/config"
	"github.com/jplanckeel/events-tracker/pkg/handlers"
)

// @title        Events-Tracker API
// @version      1.0.0
// @description  This API was built to stock all events in infrastructure

type IServer interface {
	Initialize() error
}

type Server struct {
	statusHandler handlers.IStatusHandlers
	eventHandlers handlers.IEventHandlers
}

func NewServer(statusHandler handlers.IStatusHandlers, eventHandlers handlers.IEventHandlers) IServer {
	return &Server{
		statusHandler: statusHandler,
		eventHandlers: eventHandlers,
	}
}

func (s *Server) Initialize() error {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/status", s.statusHandler.Index)

	apiv1 := router.PathPrefix("/api/v1/events").Subrouter()
	apiv1.HandleFunc("/", s.eventHandlers.List).Methods(http.MethodGet)
	apiv1.HandleFunc("/", s.eventHandlers.Create).Methods(http.MethodPost)

	log.Info("Listen on port: ", config.Config.Port)

	return http.ListenAndServe(":"+config.Config.Port, router)
}
