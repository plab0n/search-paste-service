package bus

import (
	"errors"
	"github.com/plab0n/search-paste/pkg/logger"
	"time"
)

type MessageBus struct {
	channels map[string]chan interface{}
}

type Bus interface {
	Publish(topic string, message interface{}) error
	Subscribe(topic string) (chan interface{}, error)
	SubscribeWithAction(topic string, action func(message interface{}) error)
}

var bus *MessageBus

func New() *MessageBus {
	if bus != nil {
		return bus
	}
	bus := &MessageBus{
		channels: make(map[string]chan interface{}),
	}
	return bus
}

func (m *MessageBus) Publish(topic string, message interface{}) error {
	ch := m.channels[topic]
	if ch == nil {
		return errors.New("No topic found!")
	}
	ch <- message
	return nil
}

func (m *MessageBus) Subscribe(topic string) (chan interface{}, error) {
	ch := m.channels[topic]
	if ch != nil {
		return ch, nil
	}
	ch = make(chan interface{})
	m.channels[topic] = ch
	return ch, nil
}

func (m *MessageBus) SubscribeWithHandler(topic string, action func(message interface{}) error) error {
	ch := m.channels[topic]
	if ch == nil {
		ch = make(chan interface{})
		m.channels[topic] = ch
	}
	go func() {
		for i := 0; ; i++ {
			message := <-ch
			err := action(message)
			if err != nil {
				logger.Log.Error(err)
			}
			time.Sleep(time.Second)
		}
	}()
	return nil
}
