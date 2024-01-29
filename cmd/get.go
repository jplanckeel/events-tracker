package cmd

import (
	"log"
	"os"

	"text/tabwriter"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/spf13/cobra"

	"github.com/jplanckeel/events-tracker/client"
	"github.com/jplanckeel/events-tracker/client/event"
	"github.com/jplanckeel/events-tracker/models"
	"github.com/jplanckeel/events-tracker/pkg/print"
	"github.com/jplanckeel/events-tracker/pkg/utils"
)

type cfg struct {
	host   string
	output string
}

type filter struct {
	source    string
	priority  string
	eventType string
	status    string
	service   string
	startDate string
	endDate   string
}

var config cfg
var f filter

var get = &cobra.Command{
	Use:   "get",
	Short: "Cli to get events in events-tracker",
	Run: func(cmd *cobra.Command, args []string) {

		transport := httptransport.New(
			config.host, "/", []string{"http"})
		c := client.New(transport, nil)

		e := &event.GetAPIV1EventsSearchParams{
			EndDate:   &f.endDate,
			Priority:  &f.priority,
			Source:    &f.source,
			StartDate: &f.startDate,
			Status:    &f.status,
			Service:   &f.service,
			Type:      &f.eventType,
		}

		r, err := c.Event.GetAPIV1EventsSearch(e)
		if err != nil {
			log.Fatal(err)
		}

		getPrint(r.Payload)

	},
}

func init() {
	rootCmd.AddCommand(get)
	get.PersistentFlags().StringVar(&config.host, "host", "localhost:9101", "host for events-tracker api")
	get.PersistentFlags().StringVarP(&config.output, "output", "o", "", "output format 'wide'")

}

func getPrint(payload []*models.ServicesEventOutput) {
	//print table
	w := tabwriter.NewWriter(os.Stdout, 00, 0, 1, ' ', 0)

	if config.output == "wide" {
		header := print.FormatLine("SERVICE", 30) + print.FormatLine("STATUS", 15) + print.FormatLine("PRIORITY", 15) + print.FormatLine("TYPE", 15) + print.FormatLine("CREATED_DATE", 30) + print.FormatLine("PR_ID", 10) + "\n"
		w.Write([]byte(header))
		for _, p := range payload {
			pr, _ := utils.CatchPullRequestId(p.Links.PullRequestLink)
			line := print.FormatLine(*p.Attributes.Service, 30) + print.FormatLine(*p.Attributes.Status, 15) + print.FormatLine(*p.Attributes.Priority, 15) + print.FormatLine(*p.Attributes.Type, 15) + print.FormatLine(p.Metadata.CreatedAt, 30) + print.FormatLine("PR"+pr, 10) + "\n"
			w.Write([]byte(line))
		}
	} else {
		header := print.FormatLine("SERVICE", 30) + print.FormatLine("STATUS", 15) + print.FormatLine("PRIORITY", 15) + "\n"
		w.Write([]byte(header))
		for _, p := range payload {
			line := print.FormatLine(*p.Attributes.Service, 30) + print.FormatLine(*p.Attributes.Status, 15) + print.FormatLine(*p.Attributes.Priority, 15) + "\n"
			w.Write([]byte(line))
		}
	}

	// Flush the Writer to ensure all data is written to the output.
	w.Flush()
}
