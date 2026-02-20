package mongodb

import (
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func NewMongoDBClient(cfg *MongoCfg) (*mongo.Client, error) {
	const (
		connectionTimeout = 60 * time.Second
		maxPoolSize       = 100
		minPoolSize       = 10
		minIdleTime       = 30 * time.Second
		maxIdleTime       = 30 * time.Second
	)

	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	opts := options.Client().ApplyURI(uri).
		SetConnectTimeout(connectionTimeout).
		SetMaxConnIdleTime(maxIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.WrapIf(
			err,
			fmt.Sprintf("failed to connect to MongoDB at %s", cfg.URI),
		)
	}

	if cfg.EnableTracing {
		opts.Monitor = otelmongo.NewMonitor()
	}

	return client, nil
}
