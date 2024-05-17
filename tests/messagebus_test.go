package bus_test

import (
	"fmt"
	"github.com/plab0n/search-paste/internal/bus"
	"testing"
	"time"
)

func TestMessageBus_PublishAndSubscribe(t *testing.T) {
	messageBus := bus.New()
	fmt.Println("Starting")
	// Subscribe to a topic
	ch, err := messageBus.Subscribe("test_topic")
	if err != nil {
		t.Errorf("Error subscribing to topic: %v", err)
	}
	expectedMessage := "Hello, World!"

	go func() {
		for i := 0; ; i++ {
			receivedMessage := <-ch
			fmt.Println(receivedMessage)
			if receivedMessage != expectedMessage {
				t.Errorf("Received unexpected message. Expected: %v, Received: %v", expectedMessage, receivedMessage)
			}
			time.Sleep(time.Second)
		}
	}()
	//select {
	//case receivedMessage := <-ch:
	//	if receivedMessage != expectedMessage {
	//		t.Errorf("Received unexpected message. Expected: %v, Received: %v", expectedMessage, receivedMessage)
	//	}
	//}
	// Publish a message to the topic
	go func() {
		err := messageBus.Publish("test_topic", expectedMessage)
		if err != nil {
			t.Errorf("Error publishing message: %v", err)
		}
	}()
	time.Sleep(time.Second * 10)
	// Wait for the message to be received

}

func TestMessageBus_PublishAndSubscribeWithHandler(t *testing.T) {
	messageBus := bus.New()
	expectedMessage := "Hello, World!"

	// Subscribe to a topic
	err := messageBus.SubscribeWithHandler("test_topic", func(message interface{}) error {
		fmt.Println(message)
		if message != expectedMessage {
			t.Errorf("Received unexpected message. Expected: %v, Received: %v", expectedMessage, message)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Error subscribing to topic: %v", err)
	}
	go func() {
		err := messageBus.Publish("test_topic", expectedMessage)
		if err != nil {
			t.Errorf("Error publishing message: %v", err)
		}
	}()
	time.Sleep(time.Second * 3)
	// Wait for the message to be received

}

func TestMessageBus_PubSubMultipleGoroutine(t *testing.T) {
	messageBus := bus.New()
	fmt.Println("Starting")
	// Subscribe to a topic
	ch, err := messageBus.Subscribe("test_topic")
	if err != nil {
		t.Errorf("Error subscribing to topic: %v", err)
	}
	expectedMessage := "Hello, World!"

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; ; j++ {
				receivedMessage := <-ch
				fmt.Printf("%d %s\n", i, receivedMessage)
				if receivedMessage != expectedMessage {
					t.Errorf("Received unexpected message. Expected: %v, Received: %v", expectedMessage, receivedMessage)
				}
				time.Sleep(time.Second)
			}
		}()
	}
	//select {
	//case receivedMessage := <-ch:
	//	if receivedMessage != expectedMessage {
	//		t.Errorf("Received unexpected message. Expected: %v, Received: %v", expectedMessage, receivedMessage)
	//	}
	//}
	// Publish a message to the topic
	for i := 0; i < 10; i++ {
		go func() {
			err := messageBus.Publish("test_topic", expectedMessage)
			if err != nil {
				t.Errorf("Error publishing message: %v", err)
			}
		}()
	}
	time.Sleep(time.Second * 10)
	// Wait for the message to be received

}
