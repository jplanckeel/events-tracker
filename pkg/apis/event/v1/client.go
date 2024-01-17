package event

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EventsV1Client is used to interact with features
type EventsV1Client struct {
	mongoCollection *mongo.Collection
}

func (c *EventsV1Client) EventsTracker(collection string) EventInterface {
	return newEventsTracker(c, collection)
}

func NewClient() EventsV1Client {

	host := "127.0.0.1"
	port := 27017
	//caCertPath := "<instance_certificate.crt>"
	// prepare the uri for the connection
	uri := fmt.Sprintf(
		"mongodb://%s:%d",
		host,
		port,
	)

	ctx := context.Background()
	m, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	var cs EventsV1Client
	db := m.Database("ferretdb")
	db.CreateCollection(ctx, "events")
	cs.mongoCollection = db.Collection("events")

	return cs
}
