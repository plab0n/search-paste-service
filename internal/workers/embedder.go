package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Embedder struct {
	*BaseWorker
}

func (e *Embedder) Start() error {
	b := bus.New() // same as before such as scrapper
	h := &workers.WorkerHandler{Bus: b}
	topic := workerutils.EmbeddingTopic()
	err := b.SubscribeWithHandler(topic, h.EmbeddingHandler)
	if err != nil {
		logger.Log.Error(err)
	}
	return err
}
