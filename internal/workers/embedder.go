package workers

import (
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Embedder struct {
	BaseWorker
}

func (e *Embedder) Start() error {
	h := &workers.WorkerHandler{Bus: e.B}
	topic := workerutils.EmbeddingTopic()
	err := e.B.SubscribeWithHandler(topic, h.EmbeddingHandler)
	if err != nil {
		logger.Log.Error(err)
	}
	return err
}
