package bus

import "errors"

type MessageBus struct {
	channels map[string]chan interface{}
}

type Bus interface {
	Publish(topic string, message interface{}) error
	Subscribe(topic string) (chan interface{}, error)
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
