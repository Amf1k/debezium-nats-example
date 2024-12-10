package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/jetstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"log/slog"
	"os"
	"products-cdc/domain"
	"time"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	natsURL := "nats://nats:4222"
	subscriber, err := jetstream.NewSubscriber(jetstream.SubscriberConfig{
		URL:                 natsURL,
		Logger:              watermill.NewSlogLogger(logger),
		AckWaitTimeout:      5 * time.Second,
		ResourceInitializer: jetstream.GroupedConsumer("test"),
		AckAsync:            true,
	})

	if err != nil {
		logger.Error("Failed to create subscriber", slog.Any("error", err))
		panic(err)
	}

	router, err := message.NewRouter(message.RouterConfig{}, watermill.NewSlogLogger(logger))
	if err != nil {
		logger.Error("Failed to create router", slog.Any("error", err))
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)

	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Recoverer,
	)

	router.AddNoPublisherHandler(
		"products-handler",
		"DebeziumStream",
		subscriber,
		handleProductMessage(logger),
	)

	if rErr := router.Run(ctx); rErr != nil {
		logger.Error("Failed to run router", slog.Any("error", rErr))
		panic(rErr)
	}
	logger.Info("Consumer stopped")
}

func handleProductMessage(logger *slog.Logger) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		var product domain.Product

		logger.Info("Received message", slog.Any("payload", string(msg.Payload)))
		if err := json.NewDecoder(bytes.NewReader(msg.Payload)).Decode(&product); err != nil {

			logger.Error("Failed to decode message", slog.Any("error", err))
			return err
		}

		logger.Info("Received message", slog.Any("product", product))
		return nil
	}
}
