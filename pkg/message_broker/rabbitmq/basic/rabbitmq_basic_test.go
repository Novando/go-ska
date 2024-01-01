package basic_test

import (
	"fmt"
	"github.com/ciazhar/zhar/pkg/message_broker/rabbitmq/basic"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"testing"
)

var (
	conn      *amqp.Connection
	ch        *amqp.Channel
	queueName = "test_queue"
)

func TestMain(m *testing.M) {
	// Set up the shared connection and channel before running tests
	conn, ch = basic.New("guest", "guest", "localhost", "5672")
	defer conn.Close()
	defer ch.Close()

	// Create the queue
	basic.CreateQueue(ch, queueName)

	// Run the tests
	exitCode := m.Run()

	// Exit with the status code from tests
	os.Exit(exitCode)
}

func TestConsumeMessages(t *testing.T) {

	// Consume messages
	basic.ConsumeMessages(ch, queueName, HandleMessage)

	// Keep the application running
	forever := make(chan bool)
	<-forever
}

func HandleMessage(string2 string) {
	fmt.Printf("Received a message: %s\n", string2)
}

func TestPublishMessage(t *testing.T) {
	basic.PublishMessage(ch, queueName, "Hello, RabbitMQ!")
}