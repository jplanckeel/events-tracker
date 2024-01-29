package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/jplanckeel/events-tracker/docs"

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

	var listenAddress string = "0.0.0.0:9101"

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/status", s.statusHandler.Index)

	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	apiv1.HandleFunc("/event/", s.eventHandlers.Create).Methods(http.MethodPost)
	apiv1.HandleFunc("/event/{id}", s.eventHandlers.GetId).Methods(http.MethodGet)
	apiv1.HandleFunc("/events/list", s.eventHandlers.List).Methods(http.MethodGet)
	apiv1.HandleFunc("/events/search", s.eventHandlers.Search).Methods(http.MethodGet)

	//define logger for http server error
	handler := slog.NewJSONHandler(os.Stdout, nil)
	httplogger := slog.NewLogLogger(handler, slog.LevelError)

	server := &http.Server{
		Addr:              listenAddress,
		ReadHeaderTimeout: 2 * time.Second, // Fix CWE-400 Potential Slowloris Attack because ReadHeaderTimeout is not configured in the http.Server
		Handler:           router,
		ErrorLog:          httplogger,
	}

	slog.Info(fmt.Sprintf("events-tracker server start and listen on %s", listenAddress))

	return server.ListenAndServe()
}

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
