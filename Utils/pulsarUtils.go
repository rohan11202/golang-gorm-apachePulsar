package pulsarutils

import (
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

var Client pulsar.Client

var LogsBookProducer pulsar.Producer
var LogsBookConsumer pulsar.Consumer
var LogsAuthorProducer pulsar.Producer
var LogsAuthorConsumer pulsar.Consumer

func SetupPulsar() {

	var err error

	Client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://pulsar:6650",
	})
	if err != nil {
		log.Fatalf("Could not create Pulsar client: %v", err)
	}
	log.Println("Created  a pulsar client ")

	//------------

	LogsBookProducer, err = Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "book-logs",
	})
	if err != nil {
		log.Println("Error Creating book Producer Logs ")
	}
	log.Println("Created book Producer logs success")

	//------------

	LogsBookConsumer, err = Client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "book-logs",
		SubscriptionName: "logs-books",
	})

	if err != nil {
		log.Println("Error Creating book Logs Consumer ")
	}
	log.Println("Created book Logs Consumer success")

	//------------

	LogsAuthorProducer, err = Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "author-logs",
	})
	if err != nil {
		log.Println("Error Creating author logs")
	}

	log.Println("Created author Producer logs success")

	//------------

	LogsAuthorConsumer, err = Client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "author-logs",
		SubscriptionName: "logs-author",
	})
	if err != nil {
		log.Println("Erorr Creating author Consumer ")
	}

}

func Close() {
	if LogsBookProducer != nil {
		LogsBookProducer.Close()
	}

	if LogsBookConsumer != nil {
		LogsBookConsumer.Close()
	}

	if LogsAuthorProducer != nil {
		LogsAuthorProducer.Close()
	}

	if LogsAuthorConsumer != nil {
		LogsAuthorConsumer.Close()
	}

	if Client != nil {
		Client.Close()
	}
}
