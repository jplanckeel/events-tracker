package event

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventGetter interface {
	Event() EventInterface
}

// EventInterface has methods to work with Event resources.
type EventInterface interface {
	Create(ctx context.Context, Event *Event) (*Event, error)
	List(ctx context.Context) ([]Event, error)
	Get(ctx context.Context, EventGet map[string]interface{}) (*Event, error)
	Count(ctx context.Context) (int64, error)
}

// eventsTracker implements EventInterface
type eventsTracker struct {
	collection mongo.Collection
}

// newEventsTracker returns a VectorStacks
func newEventsTracker(c *EventsV1Client, collection string) *eventsTracker {
	return &eventsTracker{
		collection: *c.mongoCollection.Database().Collection(collection),
	}
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *eventsTracker) List(ctx context.Context) (results []Event, err error) {
	cursor, err := c.collection.Find(context.TODO(), bson.D{})

	//var results []Event
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err)
		//panic(err)
	}

	return
}

// Create takes the representation of an Event and creates it.  Returns the server's representation of the Event, and an error, if there is any.
func (c *eventsTracker) Create(ctx context.Context, EventInsert *Event) (result *Event, err error) {
	result = &Event{}
	id := uuid.New()
	EventInsert.Metadata.Id = id.String()
	EventInsert.Metadata.CreatedAt = time.Now()

	_, err = c.collection.InsertOne(context.TODO(), EventInsert)
	if err != nil {
		log.Fatalln("Error Inserting Document", err)
	}

	err = c.collection.FindOne(context.TODO(), map[string]interface{}{"metadata.id": &EventInsert.Metadata.Id}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return
}

// Get an Event and creates it.  Returns the server's representation of the Event, and an error, if there is any.
func (c *eventsTracker) Get(ctx context.Context, EventGet map[string]interface{}) (result *Event, err error) {
	result = &Event{}

	err = c.collection.FindOne(context.TODO(), EventGet).Decode(&result)
	return
}

// Get an Event and creates it.  Returns the server's representation of the Event, and an error, if there is any.
func (c *eventsTracker) Count(ctx context.Context) (count int64, err error) {
	opts := options.Count().SetHint("_id_")
	count, err = c.collection.CountDocuments(context.TODO(), bson.D{}, opts)
	return
}
