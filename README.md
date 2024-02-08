<p align="center" style="margin-top: 120px">

  <h3 align="center">EventsTracker</h3>

  <p align="center">
    An Open-Source monitoring events solution.
    <br />
  </p>
</p>


## About EventsTracker 
 
EventsTracker is open-source alternative to Datadog events or Newrelic custom events. The solution is composed of an api and a cli that enable the creation and retrieval of events in a Mongo or FeretDB database.

The idea behind this solution is to provide a simple way of keeping track of everything that happens on your platform, especially in a world of distributed services. Track the start and end of a deployment incident or the opening of an incident.

Each time an event is created, we create a log in json format, which enables EventsTracker to be coupled with a logging solution such as Opensearch or Loki to correlate with logs and metrics.  

## Features

- [x] Swagger docs in /docs
- [x] Create and search event in API
- [x] Linked event in attributes
- [x] Link a pull_request to an event
- [x] Calculates the time between two linked events
- [x] Cli to create and search event
- [ ] Lock deployment endpoint
- [ ] Add to cli lock and unlock function
- [ ] Config file for cli
- [ ] function search event of the day on cli
- [ ] Github Action to add event in CD pipeline
- [ ] Gitlab example to add event in CD pipeline

## CLI

### Create events
```bash
event create --help
Cli to create events in EventsTracker

Usage:
  event create [flags]

Flags:
  -h, --help                  help for create
      --host string           host for EventsTracker api (default "localhost:9101")
      --message string        message of event to be created
  -p, --priority string       priority of event to be created: 'P1|P2|P3|P4' (default "P4")
  -r, --pull_request string   pull_request of event to be created
      --related_id string     id of related event of event to be created
      --service string        service of event to be created
      --source string         source of event to be created (default "cli")
      --status string         status of event to be created: 'start|success|error|failed|...'
      --title string          title of event to be created
  -t, --type string           type of event to be created: 'deployment|incident' (default "deployment")
```
### Get events
```bash
event get --help
Cli to get events in EventsTracker

Usage:
  event get [flags]

Flags:
      --end_date string     filter events before after end_date ex:2024-01-01
  -h, --help                help for get
      --host string         host for EventsTracker api (default "localhost:9101")
  -o, --output string       output format 'wide'
      --priority string     filter events with priority: 'P1|P2|P3|P4'
      --service string      filter events with service name
      --source string       filter events with source
      --start_date string   filter events created after start_date ex:2024-01-01
      --status string       filter events with status: 'start|success|error|failed|...'
      --type string         filter events with type: 'deployment|incident' (default "deployment")
```

## Getting Started 🚀

### Requirements

- [golang](https://go.dev/) >= 1.21
- [swagger](https://swagger.io/)
- [swag](https://github.com/swaggo/swag) >=1.16.2

### Build 

To compile EventsTracker run this command, output a binnary in bin/event

```bash
make build
```

### Update Swagger Docs 

To updates swagger docs and EventsTracker client run this command: 

```bash
make generate
```


### Test

To run test: 

```bash
make test
```
