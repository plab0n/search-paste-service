package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Scrapper struct {
	*BaseWorker
	// you can introduce bus in here so that, you don't have call the new method every time as you invoked the start method
}

var b *bus.MessageBus

func (c *Scrapper) Start() error {
	b = bus.New() // referring to this line
	topic := workerutils.PasteCrawlTopic()
	h := &workers.WorkerHandler{Bus: b}
	err := b.SubscribeWithHandler(topic, h.Scrapper)
	return err
}
