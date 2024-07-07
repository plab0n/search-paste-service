package workers

import (
	"github.com/plab0n/search-paste/internal/vector_storage"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Indexer struct {
	BaseWorker
	vector_storage.VectorStorage
}

func (c *Indexer) Start() error {
	topic := workerutils.PasteIndexerTopic()
	h := &workers.IndexHandler{Bus: c.B, VectorStorage: c.VectorStorage}
	err := c.B.SubscribeWithHandler(topic, h.IndexingHandler)
	return err
}
