package event

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/jplanckeel/events-tracker/pkg/config"
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

	var uri string
	c := config.Config

	// prepare the uri for the connection
	if c.DatabaseUsername != "" && c.DatabasePassword != "" {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s?maxPoolSize=20&tls=true&authMechanism=PLAIN",
			c.DatabaseUsername,
			c.DatabasePassword,
			c.Database,
			c.DatabasePort,
			c.DatabaseName,
		)
	} else {
		uri = fmt.Sprintf(
			"mongodb://%s:%s/%s?maxPoolSize=20",
			c.Database,
			c.DatabasePort,
			c.DatabaseName,
		)
	}

	var caCertPool *x509.CertPool
	var cert tls.Certificate
	var err error

	// Loads CA certificate file
	if c.DatabaseCAFile != "" {
		caCert, err := os.ReadFile(c.DatabaseCAFile)
		if err != nil {
			panic(err)
		}
		caCertPool = x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			panic("Error: CA file must be in PEM format")
		}
	}

	// Loads client certificate files
	if c.DatabaseCertFile != "" && c.DatabaseKeyFile != "" {
		cert, err = tls.LoadX509KeyPair(c.DatabaseCertFile, c.DatabaseKeyFile)
		if err != nil {
			panic(err)
		}
	}

	var m *mongo.Client
	ctx := context.Background()
	if c.DatabaseCAFile != "" {

		tlsConfig := &tls.Config{
			RootCAs:      caCertPool,
			Certificates: []tls.Certificate{cert},
		}
		m, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig))
	} else {
		m, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	}

	var cs EventsV1Client
	db := m.Database(c.DatabaseName)
	db.CreateCollection(ctx, "events")
	cs.mongoCollection = db.Collection("events")

	return cs
}
