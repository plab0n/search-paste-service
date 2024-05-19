package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Scrapper struct {
	Worker
}

var b *bus.MessageBus

func (c *Scrapper) Start() error {
	b = bus.New()
	topic := workerutils.PasteCrawlTopic()
	h := &workers.WorkerHandler{Bus: b}
	err := b.SubscribeWithHandler(topic, h.ScrapeUrl)
	return err
}
