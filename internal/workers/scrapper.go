package workers

import (
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Scrapper struct {
	BaseWorker
}

func (c *Scrapper) Start() error {
	topic := workerutils.PasteCrawlTopic()
	h := &workers.WorkerHandler{Bus: c.B}
	err := c.B.SubscribeWithHandler(topic, h.Scrapper)
	return err
}
