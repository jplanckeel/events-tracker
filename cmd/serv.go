package cmd

import (
	v1 "github.com/jplanckeel/events-tracker/pkg/apis/event/v1"
	"github.com/jplanckeel/events-tracker/pkg/handlers"
	"github.com/jplanckeel/events-tracker/pkg/server"
	"github.com/jplanckeel/events-tracker/pkg/services"
	"github.com/spf13/cobra"
)

var serv = &cobra.Command{
	Use:   "serv",
	Short: "Run event-tracker server",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := v1.NewClient()

		collection := client.EventsTracker("events")

		eventService := services.NewEventService(collection)
		eventHandler := handlers.NewEventHandlers(eventService)
		statusHandler := handlers.NewStatusHandlers()

		return server.NewServer(statusHandler, eventHandler).Initialize()

	},
}

func init() {
	rootCmd.AddCommand(serv)

}
