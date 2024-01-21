package exporter

import (
	"log/slog"

	"github.com/jplanckeel/events-tracker/pkg/services"
	"github.com/prometheus/client_golang/prometheus"
)

type eventsTrackerCollector struct {
	countEvents  *prometheus.Desc
	eventService services.IEventService
}

func NewEventsTrackerCollector(s services.IEventService) prometheus.Collector {
	return &eventsTrackerCollector{
		countEvents: prometheus.NewDesc(
			"events_tracker_count",
			"count of events in events-tracker",
			[]string{},
			nil,
		),
		eventService: s,
	}
}

func (collector *eventsTrackerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.countEvents
}

func (collector *eventsTrackerCollector) Collect(ch chan<- prometheus.Metric) {
	count, err := collector.eventService.Count()
	if err != nil {
		slog.Error("events_tracker_metric count error %s", err)
	}
	ch <- prometheus.MustNewConstMetric(collector.countEvents, prometheus.GaugeValue, float64(count))
}
