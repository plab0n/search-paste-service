package bus

type MessageBus struct {
	channels map[string]chan interface{}
}

type Bus interface {
	Publish(topic string, message interface{}) bool
	Subscribe(topic string)
}

var bus MessageBus

func NewMessageBus() *MessageBus {
	bus := &MessageBus{
		channels: make(map[string]chan interface{}),
	}
	return bus
}

func (m *MessageBus) Publish(topic string, message interface{}) bool {
	return false
}

func (m *MessageBus) Subscribe(topic string) {

}
