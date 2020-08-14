package cmd

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"lion/usecase"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run http server",
	Run:   runHttpServer,
}

func runHttpServer(command *cobra.Command, args []string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	logger := watermill.NewStdLogger(true, true)
	// Set up subscriber, brokers, consumer group, subscriber, publisher, router

	// Subscriber
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Brokers
	brokers := []string{":9092"}

	// Consumer group
	consumerGroup := "lion-consumer-group"

	// Subscriber
	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         consumerGroup,
		},
		logger)
	if err != nil {
		panic(err)
	}

	// Publisher
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger)
	if err != nil {
		panic(err)
	}

	// Router
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      2,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,
		middleware.Recoverer,
	)

	// Set up lion
	handlerName := "lion handler"
	subscribeTopic := "lion-subscribe-topic"
	publishTopic := "lion-publish-topic"
	usecase.RegisterLionHandler(router, subscriber, publisher, handlerName, subscribeTopic, publishTopic)

	//sample := &model.Prey{
	//	Name:  "kangaroo",
	//	Color: "blue",
	//}
	//data, err := json.Marshal(sample)
	//if err != nil {
	//	panic(err)
	//}
	//err = publisher.Publish(subscribeTopic, &message.Message{
	//	UUID:    watermill.NewUUID(),
	//	Payload: data,
	//})
	//if err != nil {
	//	panic(err)
	//}
	ctx := context.Background()
	if err := router.Run(ctx); err != nil {
		panic(err)
	}
	<-c

}
