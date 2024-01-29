package cmd

import (
	"fmt"
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/spf13/cobra"

	"github.com/jplanckeel/events-tracker/client"
	"github.com/jplanckeel/events-tracker/client/event"
	"github.com/jplanckeel/events-tracker/models"
)

type message struct {
	source    string
	priority  string
	eventType string
	status    string
	service   string
	message   string
	title     string
	pr        string
	relatedId string
}

var m message

var create = &cobra.Command{
	Use:   "create",
	Short: "Cli to create events in events-tracker",
	Run: func(cmd *cobra.Command, args []string) {

		if os.Getenv("EVENTS_TRACKER_HOST") != "" {
			config.host = os.Getenv("EVENTS_TRACKER_HOST")
		}

		transport := httptransport.New(
			config.host, "/", []string{"http"})
		c := client.New(transport, nil)

		if m.message == "" {
			m.message = m.title
		}

		e := &event.PostAPIV1EventParams{
			Data: &models.HandlersCreateDTO{
				Attributes: &models.EventEventAttributes{
					Message:   &m.message,
					Priority:  &m.priority,
					RelatedID: m.relatedId,
					Service:   &m.service,
					Source:    &m.source,
					Status:    &m.status,
					Type:      &m.eventType,
				},
				Links: &models.EventEventLinks{
					PullRequestLink: m.pr,
				},
				Title: m.title,
			},
		}

		r, err := c.Event.PostAPIV1Event(e)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("event %s created\n", r.Payload.Metadata.ID)
	},
}

func init() {
	rootCmd.AddCommand(create)
	create.PersistentFlags().StringVar(&config.host, "host", "localhost:9101", "host for events-tracker api")
	create.PersistentFlags().StringVar(&m.title, "title", "", "title of event to be created")
	create.PersistentFlags().StringVar(&m.message, "message", "", "message of event to be created")
	create.PersistentFlags().StringVar(&m.relatedId, "related_id", "", "id of related event of event to be created")
	create.PersistentFlags().StringVarP(&m.eventType, "type", "t", "deployment", "type of event to be created: 'deployment|incident")
	create.PersistentFlags().StringVar(&m.service, "service", "", "service of event to be created")
	create.PersistentFlags().StringVar(&m.status, "status", "", "status of event to be created: 'start|success|error|failed|...")
	create.PersistentFlags().StringVar(&m.source, "source", "cli", "source of event to be created")
	create.PersistentFlags().StringVarP(&m.priority, "priority", "p", "P4", "priority of event to be created: 'P1|P2|P3|P4")
	create.PersistentFlags().StringVarP(&m.pr, "pull_request", "r", "", "pull_request of event to be created")

}
